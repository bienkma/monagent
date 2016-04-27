# monagent version 1.0
For goal of learning golang. I write mini tool. it's collections cpu, network, ram information to elasticsearch. After
grafana will display metric. I hope help everyone monitor owner system.

This tool have written golang.If you want to use the tool, you can download source and build:
1. You have to install golang http://golang.org
2. Run command:
```
go get github.com/bienkma/monagent
```
3. Run command:
```
cd $GOROOT/src/github.com/bienkma/monagent
```
4. Run command:
```
go build; go install
```
5. copy monagent to all server linux. You want to monitor