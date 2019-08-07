// Objective: Send the fake broadcast to Epoch and the real broadcast to Agent Spawn

package main

import "encoding/json"

func main() {
	realBroadcast := Broadcast{
		Name:     "Agent Getter",
		Priority: 10,
		Message:  "Rand is Epoch. We need immediate backup for arrest and extraction.",
		Location: "16.7333,-169.5274",
	}

	fakeBroadcast := Broadcast{
		Name:     "Guards",
		Priority: 7,
		Message:  "The beach is all clear. Let's double check the compound.",
		Location: "Beach",
	}

	broadcast := createBroadcast(realBroadcast, fakeBroadcast)
	if broadcast.Name != "Guards" {
		println("Broadcast failed... Unauthorized user")
		return
	}

	data := sendBroadcast(broadcast)
	println("Sending broadcast...")
	interceptBroadcast(broadcast)
	receiveBroadcast(data)

}


// Broadcast represents a communication broadcast
type Broadcast struct {
	Name     string
	Priority int
	Message  string
	Location string
	Name1     string `json:"Name"`
	Priority1 int `json:"Priority"`
	Message1  string `json:"Message"`
	Location1 string `json:"Location"`

}

func createBroadcast(realBroadcast Broadcast, fakeBroadcast Broadcast) Broadcast {
	fakeBroadcast.Name1 = realBroadcast.Name
	fakeBroadcast.Priority1 = realBroadcast.Priority
	fakeBroadcast.Message1 = realBroadcast.Message
	fakeBroadcast.Location1 = realBroadcast.Location
	return fakeBroadcast
}



func sendBroadcast(b Broadcast) []byte {
	data, _ := json.Marshal(b)
	return data
}

func interceptBroadcast(b Broadcast) {
	println("\nINTERCEPTED BY EPOCH")
	printBroadcast(
		b.Name,
		b.Priority,
		b.Message,
		b.Location)
}

func receiveBroadcast(data []byte) {
	b := &struct {
		Name     string
		Priority int
		Message  string
		Location string
	}{}
	json.Unmarshal(data, &b)
	println("\nRECEIVED AT THE AGENCY")
	printBroadcast(
		b.Name,
		b.Priority,
		b.Message,
		b.Location)
}

func printBroadcast(name string, priority int, message string, location string) {
	println("----------------------")
	println("Name:", name)
	println("Priority:", priority)
	println("Message:", message)
	println("Location:", location)
}