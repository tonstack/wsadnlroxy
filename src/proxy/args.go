package proxy

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/tonstack/wsadnlroxy/src/utils"
)

func readGetArgs(q url.Values) (*liteServerConfig, error) {
	lsip, err := strconv.Atoi(q.Get("ip"))
	if err != nil {
		return nil, fmt.Errorf("can't read ip: %s", err)
	}

	lsport, err := strconv.Atoi(q.Get("port"))
	if err != nil {
		return nil, fmt.Errorf("can't read port: %s", err)
	}

	pubkey := q.Get("pubkey")
	if pubkey == "" {
		return nil, errors.New("pubkey is empty")
	}

	hostport := fmt.Sprintf("%s:%s", utils.IntIPToString(lsip), strconv.Itoa(lsport))
	return &liteServerConfig{hostport, pubkey}, nil
}
