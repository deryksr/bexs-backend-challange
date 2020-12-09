all: clean build
	$(info Starting app ...)
	@./mysolution input.csv

build:
	@go build -o mysolution main.go
	$(info The project has been built with success)
	$(info -----------------------------------------------------------------)

clean:
	@rm -rf mysolution
	$(info -----------------------------------------------------------------)
	$(info The executable has been deleted)

test:
	$(info Starting running tests ....)
	@go test service/*
	$(info All tests has been runned)