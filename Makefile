NAME = tweet 

setup:
	mkdir -p ~/.tweet/
	cp config.json.example ~/.tweet/config.json

build:
	go build -o $(NAME)

install:
	go install

test:
	go test ./... -cover
