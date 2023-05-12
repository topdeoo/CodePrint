# 一个简单丑陋的XCPC打印服务

采用 `flask` 框架搭建，并且运行在 `Windows` 系统上

## 安装

1. 安装 `python` 环境

请上网搜索如何下载 `python` 并安装（版本暂无要求，`3.8` 及以上均可）

2. 下载或克隆本项目到本地
```bash
git clone https://github.com/topdeoo/CodePrint.git 
```

3. 安装所需要的依赖包
```bash
pip install flask
pip install pywin32
```

4. 运行
```bash
cd CodePrint
python main.py
```

## 配置

对于文件中的 `team_info.csv`，请注意是需要裁判提供的， `password` 为10位密码（此为 `domjudge` 要求）

裁判一般都会生成 `account.tsv` 用来导入到 `domnjudge` 中，这里需要的 `team_info.csv` 也只是 `account.tsv` 的一个子集而已，只需要保留 `team_name` （登录的帐号），`password` 与 `location` 即可

其格式在此项目的 `team_info.csv` 中有示例

## FAQ

1. 如果启动失败（显示无法创建套接字）类似这种错误，一般解决方法为更改 `main.py` 中 `app.run(host="0.0.0.0", port="8080") 的 `port` 为其他值，例如 `8888`

2. 如果场内有部分机器无法访问打印服务（一般为无线网出现此情况），原因是因为信号较差，所以直接把打印服务的机器连接到路由器上就可以解决
