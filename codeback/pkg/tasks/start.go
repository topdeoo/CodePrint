package tasks

import "acm.nenu.edu.cn/xcpc/global"

func Start() {

	printerConfig := global.MyConfig.GetPrinterConfig()

	if err := startWorker(printerConfig.PrinterName); err != nil {
		panic(err)
	}

}
