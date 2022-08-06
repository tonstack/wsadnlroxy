## wsadnlroxy

WebSockets <-> TCP(ADNL) Proxy

## fast setup with ssl (ubuntu linux)

**Prerequisites**
```text
GNU bash 5.1 or newer
GNU make 3.81 or newer
docker 20.10.17 or newer
docker-compose 1.29.2 or newer
```

1. create an `A` record for your domain
```
example.com. IN A xxx.xxx.xxx.xxx
```
2. create a `.env` file similar to `.env.example`
3. pass your domain into `.env`
2. run `make setup-wtih-ssl`

## how to upgrade

1. run `git pull`
2. remove all containers `make docker-rm-all`
3. run `make setup-wtih-ssl` or `make setup-no-ssl`

## connection string

```
wss://host:port/?ip={}&port={}&pubkey={}
```

- `ip` – liteserver signed integer IP
- `port` – liteserver unsigned integer TCP port
- `pubkey` – liteserver base64 pubkey [(percent-encoding)](https://en.wikipedia.org/wiki/Percent-encoding)

## License

The main license of this repository is `GNU GENERAL PUBLIC LICENSE Version 3`, but the repository contains an `init-ssl.sh` file its license is `MIT`.
