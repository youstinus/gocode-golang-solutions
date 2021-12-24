// Objective: Fix the race condition

package main

import (
	"math/rand"
	"sync"
	"time"
)

var buttons = []string{"red", "blue", "green", "yellow", "purple"}

func main() {
	rand.Seed(111009)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(buttons))

	sequence := make([]Button, 0)

	for x := 0; x < len(buttons); x++ {
		go addButton(x, &sequence, &waitGroup)
	}

	waitGroup.Wait()

	println("Sending Code Sequence...")
	for i, button := range sequence {
		println(i+1, ":", button.color)
	}
}


var mu sync.Mutex

func setButton(x int, sequence *[]Button) {
	mu.Lock()
	newButton := Button{buttons[x]}
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	*sequence = append(*sequence, newButton)
	mu.Unlock()
	//time.Sleep(time.Millisecond * time.Duration((6-x) * 20))
}

func setButton2(x int, sequence *[]Button) {

	newButton := Button{buttons[x]}
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	*sequence = append(*sequence, newButton)
}

func addButton(x int, sequence *[]Button, waitGroup *sync.WaitGroup) {
	setButton(x, sequence)
	waitGroup.Done()
}

// Button represents a colored button
type Button struct {
	color string
}