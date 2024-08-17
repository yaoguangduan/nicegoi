package util

import (
	"fmt"
	"github.com/gofrs/uuid"
	"log"
	"strconv"
	"strings"
	"sync/atomic"
)

var seq = atomic.Uint64{}

func AllocEID() string {
	id := fmt.Sprintf("E%d", seq.Add(1))

	return id
}

func GenUUID() string {
	v4, err := uuid.NewGen().NewV4()
	if err != nil {
		log.Println("new v4 err:", err)
		return "uuid" + strconv.FormatUint(seq.Add(1), 2)
	}
	v4s := v4.String()
	uid := v4s[0:strings.Index(v4s, "-")]
	return uid
}
