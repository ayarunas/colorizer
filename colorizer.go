package czr

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

type colorizer func(string) string

func Make(c int) colorizer {
	c_bytes, err := exec.Command("tput", "setaf", strconv.Itoa(c)).Output()

	if err != nil {
		log.Fatal("ERROR:", err)
	}

	return func(s string) string {
		var b bytes.Buffer
		b.WriteString(string(c_bytes))
		b.WriteString(s)
		color_str := b.String()
		b.Reset()

		return color_str
	}
}

func Reset() {
	reset, err := exec.Command("tput", "sgr0").Output()

	if err != nil {
		log.Fatal("ERROR:", err)
	}

	fmt.Printf(string(reset))
}
