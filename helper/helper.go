package helper

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

func SepIPPort(hostport string, ip *string, port *string) (err error) {
	arr := strings.Split(hostport, ":")
	if len(arr) != 2 {
		return errors.New("sepIPPort: can't split ip and port")
	}

	if net.ParseIP(arr[0]) == nil {
		return errors.New("sepIPPort: invalid ip")
	}

	intIp, err := strconv.Atoi(arr[1])
	if err != nil {
		return errors.New("sepIPPort: port in not integer")
	}

	if !(intIp >= 1 && intIp <= 65_535) {
		return errors.New("sepIPPort: port must be int 1 - 65535 range")
	}

	*ip = arr[0]
	*port = arr[1]

	return nil
}
