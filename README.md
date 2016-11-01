## colorizer

### An straightforward way to add color to your command line

You supply the ANSI color codes, `colorizer` hands you back nifty little functions...

install: `go get github.com/ravindersahni/czr`

Simple Example:

	package main

	import (
		"fmt"
		"github.com/ravindersahni/czr"
	)
	
	func main() {
	
	red := czr.Make(9)
	yellow := czr.Make(11)
	blue := czr.Make(12)
    
	fmt.Println(blue("This text is blue."))
	fmt.Println(red("This text is red."))
	fmt.Println(yellow("This text is red too. Just kidding, it's yellow."))
	czr.Reset() // Your terminal is now back to normal
	
	}
	
### A Public Domain Color Chart for your convenience

![256 color chart](Xterm_256color_chart.svg?raw=true)
