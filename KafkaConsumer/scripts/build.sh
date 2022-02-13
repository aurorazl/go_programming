set -e
export GOROOT=./go
export GOOS=linux
export GOARCH=amd64
go build -o billUpdateTrigger cmd/main.go
chmod +x ./billUpdateTrigger