## wsadnlroxy

WebSockets <-> TCP(ADNL) Proxy

## fast setup with ssl (ubuntu linux)

*Prerequisites*
```text
GNU bash 5.1 or newer
GNU make 3.81 or newer
docker 20.10.17 or newer
docker-compose 1.29.2 or newer
```

1. Create an `A` record for your domain
```
example.com. IN A xxx.xxx.xxx.xxx
```
2. Create a `.env` file similar to `.env.example`
3. Pass your domain into `.env`
2. Run `make setup-wtih-ssl`

## License

The main license of this repository is `GNU GENERAL PUBLIC LICENSE Version 3`, but the repository contains an `init-ssl.sh` file its license is `MIT`.
