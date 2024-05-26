package tasks

import (
	"fmt"

	"acm.nenu.edu.cn/xcpc/global"
	"github.com/RichardKnop/machinery/v2/tasks"
)

func PrintCode(rawFile string) {
	printer := global.GetNextPrinter()
	signature := &tasks.Signature{
		Name: "Print",
		Args: []tasks.Arg{
			{
				Name:  "rawFile",
				Type:  "string",
				Value: rawFile,
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
