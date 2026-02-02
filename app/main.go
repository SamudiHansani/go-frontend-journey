package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Getting started with Go for Frontend!")
	/**  Passing data from an HTML input field */
	keepAlive := make(chan bool)
	// Create a Go function that retrieves input data
	processInput := js.FuncOf(func(this js.Value, args []js.Value) any {
		document := js.Global().Get("document")
		
		// 1. Get the value from the input field
		userInput := document.Call("getElementById", "user-name").Get("value").String()
		
		// 2. Process it in Go
		message := fmt.Sprintf("Gopher says: Hello, %s! Your name is %d characters long.", userInput, len(userInput))
		
		// 3. Send it back to the UI
		document.Call("getElementById", "response-text").Set("innerHTML", message)
		
		return nil
	})

	js.Global().Set("processInputInGo", processInput)

	fmt.Println("WASM bridge is active!")
	<-keepAlive

	//---------------------------------------------------------------------------------------------------

	/** Event Handling in Go WASM */

	// 1. Create a channel to keep our app from exiting
	// keepAlive := make(chan bool)

	// // 2. Define the Go function that the button will trigger
	// onButtonClick := js.FuncOf(func(this js.Value, args []js.Value) any {
	// 	fmt.Println("Button clicked! Go is handling the event. 🚀")
		
	// 	document := js.Global().Get("document")
	// 	heading := document.Call("getElementById", "main-heading")
	// 	heading.Set("innerHTML", "The Gopher just reacted to you!")
	// 	heading.Get("style").Set("color", "orange")
		
	// 	return nil
	// })

	// // 3. Attach this function to the global 'window' object so JS can see it
	// js.Global().Set("myGoButtonHandler", onButtonClick)

	// fmt.Println("Go WASM fully loaded and waiting for clicks...")

	//---------------------------------------------------------------------------------------------------

	/** Manipulating the DOM */

	// 1. Get the global 'document' object from the browser
	// document := js.Global().Get("document")

	// // 2. Find our <h1> tag (we need to make sure it has an ID in index.html)
	// heading := document.Call("getElementById", "main-heading")

	// // 3. Change the text content of that heading
	// heading.Set("innerHTML", "Hello from the Go Speed Demon! 🐹⚡")

	// // 4. Let's even change the style!
	// heading.Get("style").Set("color", "deepskyblue")

	// // 4. Wait on the channel (this prevents the program from finishing)
	// <-keepAlive

	//---------------------------------------------------------------------------------------------------

	/** Performance Testing in Go WASM */

	// Let's do a "Speed Demon" test: Calculating a large sum
	// start := time.Now()
	// sum := 0
	// for i := 0; i < 100000000; i++ {
	// 	sum += i
	// }
	// elapsed := time.Since(start)

	// fmt.Printf("Calculation result: %d\n", sum)
	// fmt.Printf("Time taken in Go WASM: %s\n", elapsed)
}
