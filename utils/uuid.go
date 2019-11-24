package utils

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

func GenerateUUID() string {
	withLineUUID := uuid.NewV4().String()
	uuidList := strings.Split(withLineUUID, "-")
	return strings.Join(uuidList, "")
}
