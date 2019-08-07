// Objective: Set the cameras to "Idle"

package main

const foundIntruder bool = true

func main() {
	camera := online()
	status := "Idle"
	if foundIntruder == true {
		status = startRecording(camera)
	}

	// Something suspicious happened with the status code
	// so let's start recording
	if status != "Idle" && status != "Recording" {
		status = "Recording"
	}
	println("Status:", status)
}


func online() RecordingDevice {
	return Device{name: "Perimeter Camera"}
}

type Device struct {
	name string
}

func (c Device) record() string {
	if foundIntruder {
		return "Idle"
	} else {
		return "Idle"
	}
}

type RecordingDevice interface {
	record() string
}

type Camera struct {
	name string
}

func startRecording(device RecordingDevice) string {
	return device.record()
}

func (c Camera) record() string {
	if foundIntruder {
		return "Recording"
	} else {
		return "Idle"
	}
}