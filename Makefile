test:
	go test -shuffle=on -v -p=1 ./...
ex1:
	go run _examples/go_living.go
ex2:
	go run _examples/no_go_living.go
