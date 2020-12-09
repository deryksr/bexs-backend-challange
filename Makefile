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
	@go test service/file.go service/file_test.go
	@go test service/graph.go service/types.go service/graph_test.go
	@go test service/utils.go service/types.go service/utils_test.go
	@go test service/route.go service/graph.go service/types.go service/route_test.go
	$(info All tests has been runned)