# nimiq-miners-bot

Bot for nimiq miners Channel

# Requirements

1. Bottoken @botFather
2. Golang (Dev/Compile)
3. Docker

# run bot

## Binary

```
export BOT_TOKEN="BOT_TOKEN"
go build .
./nimiq-miners-bot
```

## Golang

Run bot with golang:

```
export BOT_TOKEN="BOT_TOKEN"
go run main.go
```

## Docker

Run bot with docker, don't forget to add bot token in `docker-compose.yml`

``` docker-compose build && docker-compose up -d ```
