package goroutinesandchannels

import (
	"fmt"
	"time"
)

func greet(prompt string, doneChan chan bool) {
	fmt.Println(prompt)
	doneChan <- true
}

func slowGreet(prompt string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(prompt)
	doneChan <- true
	close(doneChan)
}

func goRoutinesIntro() {
	// dones := make([]chan bool, 4)
	// OR
	done := make(chan bool)

	// dones[0] = make(chan bool)
	go greet("Hi", done)
	// dones[1] = make(chan bool)
	go greet("Hi, How are you", done)
	// dones[2] = make(chan bool)
	go slowGreet("Hi,.... How .... are.... you?", done)
	// dones[3] = make(chan bool)
	go greet("How you are doing?", done)

	// for i, _ := range dones {
	// 	fmt.Println(<-dones[i])
	// }

	// OR

	for range done {

	}

}
