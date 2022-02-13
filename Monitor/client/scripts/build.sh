
export GOOS=linux
export GOARCH=amd64
go build -o alert cmd/main.go
chmod +x ./alert
scp alert root@159.75.16.88:/home/hadoop/chenzenglong/monitor/