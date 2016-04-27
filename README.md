# monagent version 1.0
- For goal of learning golang. I have written this mini tool. it's collections cpu, network, ram information to elasticsearch. After
grafana will display metric. I hope help everyone monitor owner system.
- This tool have written golang.If you want to use the tool, you can download source and build: You have to install golang http://golang.org


1. Get tool to local:
```
go get github.com/bienkma/monagent
```
2. Build tool
```
cd $GOROOT/src/github.com/bienkma/monagent; go build; go install 
```
3. Pack monagent.
```
mkdir monagent; cd monagent
cp $GOROOT/src/github.com/bienkma/monagen/config.json .
cp $GOROOT/bin/monagent .
cd ..; tar -cvzf monagent_v1.0.tar.gz monagent/
```
4. move package monagent_v1.0.tar.gz to server you want mon. And fill information in config.json file is correct.
