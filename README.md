# nimiq-miners-bot

Bot for nimiq miners Channel <br>
<https://t.me/NimiqMiners>

# Requirements

1. Bot token @botFather
2. Golang (Dev/Compile)
3. Docker

# Run bot

## Binary

```
export BOT_TOKEN="BOT_TOKEN"
go build .
./nimiq-miners-bot
```

## Golang

```
export BOT_TOKEN="BOT_TOKEN"
go run main.go
```

## Docker

Run bot with docker-compose <br> Don't forget copy and add bot token in `docker-compose.yml`

``` docker-compose build && docker-compose up -d ```
