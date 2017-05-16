# Simple Credentials Store

[docker login credentials-store](https://docs.docker.com/engine/reference/commandline/login/#credentials-store)

```$HOME/.docker/config.json```
```
{
	"credsStore": "simple"
}
```

```$HOME/.docker/creds.json```
```
{
	"ServerURL": "https://index.docker.io/v1",
	"Username": "david",
	"Secret": "passw0rd1"
}
```

```
docker-credential-simple store < test/cred.json
docker-credential-simple get < test/url.txt
docker-credential-simple erase < test/url.txt
docker-credential-simple list
```