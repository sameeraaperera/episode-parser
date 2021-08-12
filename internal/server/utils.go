package server

import (
	"fmt"
	"strconv"
)

func GetListenPort(port string) (int, error) {
	if port == "" {
		return -1, fmt.Errorf("listen port not defined")
	}
	i, err := strconv.Atoi(port)
	if err != nil {
		return -1, fmt.Errorf("error parsing port: %v", err)
	}
	return i, nil
}
