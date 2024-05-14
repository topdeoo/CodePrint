package tasks

import (
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/tasks"

	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"

	"github.com/topdeoo/codeprint/back/global"
	"github.com/topdeoo/codeprint/back/pkg/tasks/client"
)

var AsyncTaskCenter *machinery.Server

func startServer() (*machinery.Server, error) {
	machineryConfig := global.MyConfig.GetMachineryConfig()
	redisConfig := global.MyConfig.GetRedisConfig()
	uri := redisConfig.Host + ":" + redisConfig.Port
	broker := redisbroker.NewGR(machineryConfig, []string{uri}, 0)
	backend := redisbackend.NewGR(machineryConfig, []string{uri}, 0)
	lock := eagerlock.New()

	server := machinery.NewServer(
		machineryConfig, broker, backend, lock,
	)

	tasksMap := initAsyncTaskMap()

	AsyncTaskCenter = server
	return server, server.RegisterTasks(tasksMap)

}

func initAsyncTaskMap() map[string]interface{} {
	tasksMap := map[string]interface{}{
		"Print": client.PrintCode,
	}
	return tasksMap
}

func startWorker(tag string) error {
	server, err := startServer()
	if err != nil {
		return err
	}
	client := server.NewWorker(tag, 0)

	errorhandler := func(err error) {}
	pretaskhandler := func(signature *tasks.Signature) {}
	posttaskhandler := func(signature *tasks.Signature) {}

	client.SetPostTaskHandler(posttaskhandler)
	client.SetErrorHandler(errorhandler)
	client.SetPreTaskHandler(pretaskhandler)
	return client.Launch()
}
