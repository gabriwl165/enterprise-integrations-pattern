package main

import (
	"fmt"
	"sync"
)

type Message struct {
	Type string
	Body string
}

func RouterFilter(in <-chan Message, outs map[string]chan Message) {
	defer func() {
		// Close all output channels when input channel is closed
		for _, out := range outs {
			close(out)
		}
	}()

	for msg := range in {
		if out, ok := outs[msg.Type]; ok {
			out <- msg
		} else {
			fmt.Printf("No output channel for message type: %s\n", msg.Type)
		}
	}
}

func main() {
	in := make(chan Message)

	production := make(chan Message)
	marketing := make(chan Message)
	finance := make(chan Message)

	outs := map[string]chan Message{
		"production": production,
		"marketing":  marketing,
		"finance":    finance,
	}

	var wg sync.WaitGroup

	// Start RouterFilter
	wg.Add(1)
	go func() {
		defer wg.Done()
		RouterFilter(in, outs)
	}()

	// Start consumers
	wg.Add(3)
	go func() {
		defer wg.Done()
		for msg := range production {
			fmt.Printf("Production: %s\n", msg.Body)
		}
	}()
	go func() {
		defer wg.Done()
		for msg := range marketing {
			fmt.Printf("Marketing: %s\n", msg.Body)
		}
	}()
	go func() {
		defer wg.Done()
		for msg := range finance {
			fmt.Printf("Finance: %s\n", msg.Body)
		}
	}()

	// Send messages
	in <- Message{Type: "production", Body: "Production message"}
	in <- Message{Type: "marketing", Body: "Marketing message"}
	in <- Message{Type: "finance", Body: "Finance message"}

	// Close input channel to signal completion
	close(in)

	// Wait for all goroutines to finish
	wg.Wait()
}
