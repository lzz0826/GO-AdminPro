package utils

import (
	"github.com/google/uuid"
	"github.com/rs/xid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateId() string {
	return xid.New().String()
}
