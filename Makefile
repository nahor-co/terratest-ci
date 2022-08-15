

build: 
	go mod init $$(basename "$$PWD") || true
	go mod tidy

test:
	go get
	go test -v 
