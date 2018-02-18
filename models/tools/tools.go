package tools

import (
	"github.com/rs/xid"
)

func GenerateId(letter string) string {
	newId := xid.New()
	return letter + newId.String()
}
