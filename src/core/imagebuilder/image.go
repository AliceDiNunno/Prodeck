package imagebuilder

import (
	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"image/png"
	"os"
)

func CreateImage(height int, width int, color color.RGBA) *image.RGBA {
	basicImage := image.NewRGBA(image.Rect(0, 0, width, height))
	backgroundColor := image.NewUniform(color)

	draw.Draw(basicImage, basicImage.Bounds(), backgroundColor, image.Point{}, draw.Src)

	return basicImage
}

func LoadPng(path string) *image.RGBA {
	existingImageFile, err := os.Open(path)
	if err != nil {
		// Handle error
	}
	defer existingImageFile.Close()

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	_, _, err = image.Decode(existingImageFile)
	if err != nil {
		// Handle error
	}

	// We only need this because we already read from the file
	// We have to reset the file pointer back to beginning
	existingImageFile.Seek(0, 0)

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	loadedImage, err := png.Decode(existingImageFile)

	if err != nil {
		log.Errorln(err)
		return nil
	}

	return ImageToRGBA(loadedImage)
}

func ImageToRGBA(src image.Image) *image.RGBA {
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

func ResizeImage(image *image.RGBA, width int, height int) *image.RGBA {
	m := resize.Resize(uint(width), uint(height), image, resize.Lanczos3)

	return ImageToRGBA(m)
}

func AddPngToBaseImage(baseImage *image.RGBA, png *image.RGBA) *image.RGBA {
	if png.Bounds().Max.X > baseImage.Bounds().Max.X || png.Bounds().Max.Y > baseImage.Bounds().Max.Y {
		png = ResizeImage(png, baseImage.Bounds().Max.X, baseImage.Bounds().Max.Y)
	}

	rgba := image.NewRGBA(baseImage.Bounds())
	draw.Draw(rgba, baseImage.Bounds(), baseImage, image.Point{0, 0}, draw.Src)

	png = ResizeImage(png, 86, 86)

	draw.Draw(rgba, baseImage.Bounds(), png, image.Point{-5, -5}, draw.Over)

	rgba = ResizeImage(rgba, 96, 96)

	//draw.Draw(baseImage, baseImage.Bounds(), png, image.Point{x, y}, draw.Src)

	return rgba
}

func Icon(path string) *image.RGBA {
	return IconWithBackground(path, color.RGBA{0, 0, 0, 0})
}

func IconWithBackground(path string, color color.RGBA) *image.RGBA {
	img := CreateImage(96, 96, color)
	png := LoadPng("./resources/" + path + ".png")

	if img == nil {
		log.Errorln("img is nil")
		return nil
	}

	if png == nil {
		log.Errorln("png is nil")
		return nil
	}

	newimg := AddPngToBaseImage(img, png)

	return newimg
}

func AddTextToImage(img *image.RGBA, text string) *image.RGBA {
	const S = 96
	im := img

	dc := gg.NewContext(S, S)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.SetRGB(1, 1, 1)
	if err := dc.LoadFontFace("./resources/arial.ttf", 36); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored(text, S/2, S/2, 0.5, 0.5)

	dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringAnchored(text, S/2, S/2, 0.5, 0.5)
	dc.Clip()

	return dc.Image().(*image.RGBA)
}

func IconWithText(path string, text string) *image.RGBA {
	img := Icon(path)
	img = AddTextToImage(img, text)

	return img
}
