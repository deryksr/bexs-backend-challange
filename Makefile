all: clean build
	$(info Starting application ...)
	@./mysolution input.csv

build:
	@go build -o mysolution main.go
	$(info The project has been built with success)

clean:
	@rm -rf mysolution
	$(info The executable has been deleted)