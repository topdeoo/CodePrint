package client

import (
	"github.com/go-cmd/cmd"
)

func PrintCode(filename string, printer string) {
	cmdString := "cat " + filename + " | paps --font=10 --top-margin=10 --bottom-margin=10 --lpi=10 | lp -d " + printer
	c := cmd.NewCmd("bash", "-c", cmdString)
	go func() {
		<-c.Start()
	}()
}
