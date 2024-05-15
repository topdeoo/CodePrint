package tasks

import "github.com/topdeoo/codeprint/back/global"

func Start() {

	printerConfig := global.MyConfig.GetPrinterConfig()

	if err := startWorker(printerConfig.PrinterName); err != nil {
		panic(err)
	}

}
