## colorizer

### A Straightforward Way to Add Color to Your Command Line.

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
	
		red := czr.Make(9) // don't use leading 0s
		yellow := czr.Make(11)
		blue := czr.Make(12)
    
		fmt.Println(blue("This text is blue."))
		fmt.Println(red("This text is red."))
		fmt.Println(yellow("This text is red too. Just kidding -- it's yellow."))
	
		czr.Reset() // Your terminal is now back to normal
	
	}
	
<br>

### A Public Domain Color Chart for Your Convenience

You'll want to use the 3 digit codes in the lower left of each block (0-255).


![256 color chart](Xterm_256color_chart.gif?raw=true)

<br>

### More to Come

I intend to build this out a bit, and you are more than welcome to help. I would, however, like to keep the focus narrow and the interface simple and clean. There are many fantastic existing packages that enable all kinds of command execution, and it is not my intent to compete with or distract from those.

<br>

### Licensing
Do what thou wilt...
