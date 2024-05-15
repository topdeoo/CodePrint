package client

import "os/exec"

func PrintCode(filename string, printer string) error {
	cmd := exec.Command("cat " + filename + " | paps --fonts=10 --top-margin=10 --bottom-margin=10 --lpi=10 | lp -d " + printer)
	err := cmd.Run()
	return err
}
