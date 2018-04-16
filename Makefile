dir = $(shell pwd)
.PHONY: linx macos clean

linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/s5go github.com/horechek/s5go

macos: 
	go build -o ./bin/s5go github.com/horechek/s5go

clean:
	rm -rf ./bin
