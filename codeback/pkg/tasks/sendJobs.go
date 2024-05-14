package tasks

import (
	"fmt"

	"github.com/RichardKnop/machinery/v2/tasks"
)

func PrintCode(filename string) {
	signature := &tasks.Signature{
		Name: "Print",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: filename,
			},
		},
	}

	_, err := AsyncTaskCenter.SendTask(signature)

	if err != nil {
		fmt.Println(err)
	}
}
