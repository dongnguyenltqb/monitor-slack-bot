package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"strings"
)

func UUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func IsUUID(uuid string) (err error) {
	s := strings.Split(uuid, "-")
	if len(s) != 5 {
		err = errors.New("wrong file id")
		return
	}
	if len(s[0]) != 8 || len(s[1]) != 4 || len(s[2]) != 4 || len(s[3]) != 4 || len(s[4]) != 12 {
		err = errors.New("wrong uuid")
		return
	}
	return
}
