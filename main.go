/*
wstcproxy – open-source WebSockets <==> TCP Proxy

Copyright (C) 2022 tonstack (github.com/tonstack)

This file is part of wstcproxy.

wstcproxy is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

wstcproxy is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with wstcproxy.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"wstcproxy/config"
	"wstcproxy/proxy"
)

func main() {
	config.Configure()
	proxy.RunServer()
}
