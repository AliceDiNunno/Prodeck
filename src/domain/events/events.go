package events

type Event string
type EventData map[string]interface{}
type EventListener func(event Event, data EventData)

const ButtonStateChangedEvent Event = "button.state.changed"
const DeviceConnectedEvent Event = "device.connected"
const DeviceDisconnectedEvent Event = "device.disconnected"
const DeviceWillSleepEvent Event = "device.willsleep"
