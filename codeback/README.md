# Code Print Service Back End

[English](README.md) | [简体中文](README_zh.md)

## TODO

- [] Make the server a true distributed server, i.e. the clients can bind with the printers and finish the task by themselves. 

## Usage

Since it needs a redis to store tasks, we need a install a redis server.

You can choose install a redis in your server or install a redis in docker.

Then, you should modify or create a new configuration file called `setting.prod.toml` in the `config` directory.

An example is as follows:

```toml
[http]
host = "localhost"
port = "8081"

[redis]
host = "localhost"
port = "6379"
password = ""
db = 0
pool = 10
maxIdle = 3
idleTimeout = 240
readTimeout = 15
writeTimeout = 15
connectTimeout = 15
normalTasksPollPeriod = 1000
delayedTasksPollPeriod = 500

[machinery]
queue = "machinery_queue"
expired = 3600

[printer]
name = ["HP_1106P", "HP_1107P"]


[secret]
key = "xcpc@nenu"

[code]
path = "./code"
```

`http` is the http server:
    - `host` is the server host
    - `port` is the server port, users can send requests to this port

`redis` is the redis server
    - `host` is the server host
    - `port` is the server port, which is used to connect to the redis server
    - `password` is the password of the redis server, defalut is empty
    - ... is default configuration of redis.conf

`machinery` is the task queue
    - `queue` is the name of the queue
    - `expired` is the expired time of tasks

`printer` is the printer
    - `name` is the name of the printers

`secret` is the token secret key
    - `key` is the value of key

`code` is the code file path
    - `path` is the path of the code file

After the configuration file is created, you should change `mode` in the `setting.toml` into `prod` like this:

```toml
[env]
mode = "prod"
```

At last, you can run the server as follows:

```bash
go run main &
```

## Change Account and Password

Put your `accounts.csv` into `db` folder, the format is in `example.csv` as follow:

```bash
team_name   password    location
team001     123456      A01
```
