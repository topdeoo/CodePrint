package global

import (
	"encoding/csv"
	"os"
	"sync"

	"acm.nenu.edu.cn/xcpc/config"
	"acm.nenu.edu.cn/xcpc/model"
)

type Counter struct {
	Index    uint64
	Spinlock sync.Mutex
}

var MyConfig *config.Config

var SecretKey string

var Database map[string]model.User

var counter *Counter

func Init() {
	MyConfig = config.ConfigInit()

	SecretKey = MyConfig.SecretKey

	counter = &Counter{Index: 0, Spinlock: sync.Mutex{}}

	initDatabase()
}

func GetNextPrinter() string {
	counter.Spinlock.Lock()
	defer counter.Spinlock.Unlock()
	printer := MyConfig.GetPrinterConfig()
	result := printer.PrinterName[counter.Index]
	counter.Index = (counter.Index + 1) % uint64(len(printer.PrinterName))
	return result
}

func initDatabase() {
	file, err := os.Open("db/example.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	records, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	Database = make(map[string]model.User)

	for idx, record := range records {
		if idx == 0 {
			continue
		}
		Database[record[0]] = model.User{TeamName: record[0], Password: record[1], Location: record[2]}
	}
}
