## colorizer

### An straightforward way to add color to your command line.

You supply the ANSI color codes, `colorizer` hands you back nifty little functions...

<br>

install: `go get github.com/ravindersahni/czr`

<br>

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
	
<br>

### A Public Domain Color Chart for your convenience

You'll want to use the 3 digit codes in the lower left of each block (0-255).


![256 color chart](Xterm_256color_chart.gif?raw=true)
