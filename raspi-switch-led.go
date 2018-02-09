/*

*/
package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
)

var (
	ledPin = rpio.Pin(15)
	buttonPin = rpio.Pin(16)
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sum := 0
	for {
		
		// check if button is pressed
		// toggle led
			// if on toggle off
			// else if off toggle on
		// if button is not presse, but led is on and button was pressed 2 seconds ago
		// shut off led
			// sleep 10 ms
		
	}
}
