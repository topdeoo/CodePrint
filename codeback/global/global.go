package global

import (
	"encoding/csv"
	"os"

	"acm.nenu.edu.cn/xcpc/config"
	"acm.nenu.edu.cn/xcpc/model"
)

var MyConfig *config.Config

var SecretKey string

var Database map[string]model.User

func Init() {
	MyConfig = config.ConfigInit()

	SecretKey = MyConfig.SecretKey

	initDatabase()
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
