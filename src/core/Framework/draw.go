package Framework

import (
	"bytes"
	"fmt"
	"github.com/nfnt/resize"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"
)

func createCompatibleImage() *image.RGBA {
	height := 96
	width := 96

	return image.NewRGBA(image.Rect(0, 0, width, height))
}

func imageToRGBA(src image.Image) *image.RGBA {
	// No conversion needed if image is an *image.RGBA.
	if dst, ok := src.(*image.RGBA); ok {
		return dst
	}

	// Use the image/draw package to convert to *image.RGBA.
	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)
	return dst
}

func rotateImage(toRotate *image.RGBA, angle float64) *image.RGBA {
	if toRotate == nil {
		return nil
	}

	rotated := image.NewRGBA(image.Rect(0, 0, toRotate.Bounds().Max.X, toRotate.Bounds().Max.Y))

	for x := 0; x < toRotate.Bounds().Max.X; x++ {
		for y := 0; y < toRotate.Bounds().Max.Y; y++ {
			rotated.Set(x, y, toRotate.At(toRotate.Bounds().Max.X-x, toRotate.Bounds().Max.Y-y))
		}
	}

	return rotated
}

func resizeImage(image *image.RGBA, width int, height int) *image.RGBA {
	m := resize.Resize(uint(width), uint(height), image, resize.Lanczos3)

	return imageToRGBA(m)
}

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func generateImage(textContent string, fgColorHex string, bgColorHex string, fontSize float64) ([]byte, error) {
	img := image.NewRGBA(image.Rect(0, 0, 96, 96))
	addLabel(img, 20, 30, "Hello Go")

	//encode img to jpeg buffer
	buffer := new(bytes.Buffer)
	err := png.Encode(buffer, img)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func randomColor() color.Color {
	r := rand.Int() % 255
	g := rand.Int() % 255
	b := rand.Int() % 255

	return color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}
}

func (p *ProdeckFramework) creatingBasic96x96png(key int) {
	p.SetButtonColor(key, randomColor())
}

func (p *ProdeckFramework) outlineCurrentButton(key int) {
	//check if the image is cached
	_, ok := p.imageCache.images[key]
	if ok {
		basicImage := createCompatibleImage()
		currentImage := p.imageCache.images[key]

		if currentImage == nil {
			return
		}

		height := basicImage.Bounds().Max.Y
		width := basicImage.Bounds().Max.X

		resized := resizeImage(currentImage, 76, 76)

		//draw resized image in the center of the basic image

		for x := 10; x < width-10; x++ {
			for y := 10; y < height-10; y++ {
				basicImage.Set(x, y, resized.At(x-10, y-10))
			}
		}

		p.setButtonImage(key, basicImage, false)
	}
}

func (p *ProdeckFramework) SetButtonColor(button int, color color.Color) {
	basicImage := createCompatibleImage()

	height := basicImage.Bounds().Max.Y
	width := basicImage.Bounds().Max.X

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			basicImage.Set(x, y, color)
		}
	}

	p.setButtonImage(button, basicImage, true)
}

func (p *ProdeckFramework) SetButtonJpeg(button int, fileName string) {
	path := "./resources/" + fileName

	//open image
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, err := jpeg.Decode(file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	} else {
		p.setButtonImage(button, imageToRGBA(image), true)
	}

}

func (p *ProdeckFramework) setButtonImage(button int, image *image.RGBA, cache bool) {
	if cache {
		p.imageCache.images[button] = image
	}

	p.currentDevice.SetButtonImage(button, rotateImage(image, 180))
}
