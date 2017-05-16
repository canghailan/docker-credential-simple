# 一个简单的Docker身份认证存储工具

Docker官方文档：[docker login](https://docs.docker.com/engine/reference/commandline/login/#credentials-store)

Docker引擎可以使用外置的身份认证存储，比如操作系统的钥匙串。

Docker提供了3个工具，分别对应Linux、Apple macOS、Microsoft Windows，但是Linux下的工具需要X11环境，导致大多数服务器环境无法使用。

不过外置的身份认证存储的协议是公开的，而且非常简单，可以自行实现一个简单的身份认证存储工具。所有的工作只需要实现一个程序，它的第一个参数支持store，get，erase即可。



## 实现

1. store

store命令从标准输入读取一个JSON，并将其保存即可。

如果出现错误，直接将错误输出到标准输出。

```shell
docker-credential-simple store
{
	"ServerURL": "https://index.docker.io/v1",
	"Username": "david",
	"Secret": "passw0rd1"
}
CRTL-D
```



2. get


get命令从标准输入读取服务器地址，在标准输出中返回存储的身份认证信息。

如果出现错误，直接将错误输出到标准输出。

```shell
docker-credential-simple get
https://index.docker.io/v1
CRTL-D
{
	"ServerURL": "https://index.docker.io/v1",
	"Username": "david",
	"Secret": "passw0rd1"
}
```



3. erase

erase命令从标准输入读取服务器地址，并清除存储的身份认证信息。

如果出现错误，直接将错误输出到标准输出。

```shell
docker-credential-simple erase
https://index.docker.io/v1
CRTL-D
```



4. list

list并不是必须的命令，作为一个辅助工具提供，将所有存储的身份认证信息输出到标准输出。

```shell
docker-credential-simple list
[
	{
		"ServerURL": "https://index.docker.io/v1",
		"Username": "david",
		"Secret": "passw0rd1"
	}
]
```



身份认证信息存储在 ```$HOME/.docker/creds.json``` 中。



## 使用说明

1. 将可执行文件拷贝到系统路径中。
2. 修改 ```$HOME/.docker/config.json``` ，添加：

```json
{
  "credsStore": "simple"
}
```

simple是docker-credential-simple去掉前缀。

3. 通过 ```docker-credential-simple store``` 添加身份认证信息，或直接编辑 ```$HOME/.docker/creds.json```。
4. 直接使用docker命令如 ```docker push``` ，不需登录。



## 注意

* 从标准输入读取时需trim
* 出现错误时需调用os.Exit()指定返回值