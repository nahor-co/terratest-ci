

test: 
	go mod init $$(basename "$$PWD") || true
	go mod tidy
	go test -v 
