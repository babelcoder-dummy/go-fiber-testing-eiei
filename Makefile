build:
    go build -o app main.go

run: build
    ./app

dev:
    air