

build: 
	go mod init $$(basename "$$PWD") || true
	go mod tidy

test:
	go mod tidy
	go test -v 
