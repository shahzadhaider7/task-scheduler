build:
	go build -o bin/main main.go

run:
	go run main.go

format:
	${DOCKRUN} bash ./scripts/format.sh

check: format
	${DOCKRUN} bash ./scripts/check.sh

test: check
	#action
