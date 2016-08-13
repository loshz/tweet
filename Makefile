setup:
	mkdir -p ~/.tweet/
	cp config.json.example ~/.tweet/config.json

build:
	go install

test:
	go test ./... -cover -race
