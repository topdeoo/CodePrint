package global

import (
	"encoding/csv"
	"os"

	"github.com/topdeoo/codeprint/back/config"
	"github.com/topdeoo/codeprint/back/model"
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
	file, err := os.Open("data/example.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	for _, record := range records {
		Database[record[0]] = model.User{TeamName: record[0], Password: record[1], Location: record[2]}
	}
}
