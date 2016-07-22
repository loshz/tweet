setup:
	cp config.json.example config.json

build:
	go install

test:
	go test ./... -cover -race
