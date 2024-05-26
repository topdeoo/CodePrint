# XCPC 打印服务后端

[English](README.md) | [简体中文](README_zh.md)

## TODO

- [] 让打印机能够绑定到每一个消费者上，是得只要有消费者空闲就可以自动打印

## 构建与使用

> 由于系统需要使用 `redis`，因此请先安装 `redis`，在本地安装或者使用 `docker` 均可，并且不需要特地配置太多，使用 `redis` 默认的配置文件即可

随后，可以创建或修改在 `config` 目录下的 `setting.prod.toml`（如果没有请自行创建），然后仿照 `setting.dev.toml` 进行自主化配置，例如如下文件：

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

`http` 表示 `http` 服务器的配置
    - `host` 表示服务器运行的地址
    - `port` 表示运行的端口，用户通过这个端口来访问 `web` 服务

`redis` 表示 `redis` 服务器的配置
    - `host` 表示服务器运行的地址
    - `port` 表示运行的端口，用于连接 `redis` 服务器
    - `password` 表示服务器的密码，默认为空
    - ... 是 `redis.conf` 的默认配置

`machinery` 表示 `machinery` 的配置
    - `queue` 表示 `machinery` 队列的名称
    - `expired` 表示 `machinery` 队列的过期时间

`printer` 表示打印机的配置
    - `name` 表示打印机的名称列表

`secret` 表示生成token的密钥配置
    - `key` 表示密钥的值

`code` 表示打印的代码文件配置
    - `path` 表示打印的代码文件的路径（仅仅是为了留存）

当这些配置都完成后，修改 `config` 文件夹下的 `setting.toml` 中的 `mode` 字段如下：

```toml
[env]
mode = "prod"
```

随后，打开终端运行

```bash
go run main &
```

## 修改账号与密码

> 由于服务器没有做数据的持久化，不存在数据库这个说法，都是直接把账号密码写在 `accounts.csv` 文件中，然后读到内存中直接做查询（反正人少，服务器内存大）

把生成的 `accounts.csv` 放在 `db` 文件夹中，格式如同目录下的 `example.csv`

> 注意，虽然是 `.csv` 但其实是以 `\t` 为分隔符的文件，请不要变成四个空格

```bash
team_name   password    location
team001     123456      A01
```

随着每次修改账号密码，需要重启服务（每次导入账号密码也需要重启服务）


