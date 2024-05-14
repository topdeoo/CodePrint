package tasks

import "github.com/topdeoo/codeprint/back/global"

func Start() {

	printerConfig := global.MyConfig.GetPrinterConfig()

	for _, printer := range printerConfig.PrinterName {
		if err := startWorker(printer); err != nil {
			panic(err)
		}
	}
}
