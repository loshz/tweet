NAME = tweet 

setup:
	mkdir -p ~/.config/tweet && cp config.example.json ~/.config/tweet/config.json

build:
	go build -o $(NAME)

install:
	go install

test:
	go test ./... -cover
