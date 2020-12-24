build:
	go build -o bin/main main.go

run:
	go run main.go

format:
	#action

check: format
	#action

test: check
	#action
