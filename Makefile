about:
	@echo 'Test task for hiring on finding the index of an element in a slice'
	@echo 'Author: Vitali Saroka (2024)'

install:
	go mod init tidy	
	go get github.com/gin-gonic/gin	
	go get golang.org/x/exp/slices	
	go get github.com/sirupsen/logrus	
	go get github.com/gin-contrib/logger	
	go get github.com/spf13/viper		

build:
	go build -o bin/main service.go

run:
	go run service.go
	
test:
	go test -v
	

all:
	make install
	make test
	make run

help:
	@echo 'Steps for the first launch:'
	@echo '1) Install Go'
	@echo '2) Optionally create .env and et inside file PORT and LOG_LEVEL'
	@echo '3) Run with make all'
	@echo '4) Check with a web client (for.ex. http://localhost:9999/endpoint/90000)'
	@echo 'More useful information can be found in the readme file.'
