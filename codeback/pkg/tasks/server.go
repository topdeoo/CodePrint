package tasks

import (
	"acm.nenu.edu.cn/xcpc/global"
	"acm.nenu.edu.cn/xcpc/pkg/tasks/client"
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/tasks"

	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
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

func startWorker(tags []string) error {
	server, err := startServer()
	if err != nil {
		return err
	}

	for _, tag := range tags {
		client := server.NewWorker(tag, 0)

		errorhandler := func(err error) {}
		pretaskhandler := func(signature *tasks.Signature) {
			signature.Args = append(signature.Args, tasks.Arg{Name: "printer", Type: "string", Value: client.ConsumerTag})
		}
		posttaskhandler := func(signature *tasks.Signature) {}

		client.SetPostTaskHandler(posttaskhandler)
		client.SetErrorHandler(errorhandler)
		client.SetPreTaskHandler(pretaskhandler)

		err := client.Launch()
		if err != nil {
			return err
		}
	}
	return err
}
