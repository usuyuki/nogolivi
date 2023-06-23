test:
	go test -shuffle=on -p=1 ./...
ex1:
	go run _examples/go_living.go
ex2:
	go run _examples/go_living_recursively.go
ex3:
	go run _examples/no_go_living.go
ex4:
	go run _examples/go_living_panic.go
ex5:
	go run _examples/no_go_living_panic.go
