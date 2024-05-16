package tasks

import (
	"fmt"

	"acm.nenu.edu.cn/xcpc/global"
	"github.com/RichardKnop/machinery/v2/tasks"
)

func PrintCode(filename string) {
	printer := global.GetNextPrinter()
	signature := &tasks.Signature{
		Name: "Print",
		Args: []tasks.Arg{
			{
				Name:  "filename",
				Type:  "string",
				Value: filename,
			},
			{
				Name:  "printer",
				Type:  "string",
				Value: printer,
			},
		},
	}

	_, err := AsyncTaskCenter.SendTask(signature)

	if err != nil {
		fmt.Println(err)
	}
}
