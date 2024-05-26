package client

import (
	"github.com/go-cmd/cmd"
)

func PrintCode(rawFile string, printer string) {

	cmdString := "echo " + rawFile + " | paps --font=10 --top-margin=15 --bottom-margin=15 --lpi=7 --format=pdf | lp -d " + printer
	c := cmd.NewCmd("bash", "-c", cmdString)
	go func() {
		<-c.Start()
	}()
}
