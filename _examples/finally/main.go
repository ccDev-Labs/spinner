package main

// Example application using the github.com/briandowns/spinner API

import (
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.Color("red")                                              // Set the spinner color to red
	s.Finally = "Complete!\nNew line!\nAnother one!\n"
	s.Start()                   // Start the spinner
	time.Sleep(4 * time.Second) // Run for some time to simulate work
	s.Stop()                    // Stop the spinner
}