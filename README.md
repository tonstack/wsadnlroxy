## wstcproxy

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tonstack/wstcproxy)
[![GitHub license](https://img.shields.io/github/license/tonstack/wstcproxy)](https://github.com/tonstack/wstcproxy/blob/main/LICENSE)

> :warning: **WARN:**: This isn't ready for production environment

`WebSockets <==> TCP` Proxy

Allows you to send messages via WebSockets, which will then be sent to the required TCP connection. This proxy was developed because of the inability to send a raw TCP packet from the browser, but you can use it for other tasks as well.

### Connecting

Create an active websocket connection to the proxy server, specify in the query parameters `dest_host` of the destination server that supports TCP. Например `?dest_host=127.0.0.1:2035`




