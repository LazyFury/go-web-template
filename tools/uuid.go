package tools

import (
	"strings"

	"github.com/google/uuid"
)

// UUID UUID
func UUID() (_uuid string, err error) {
	uuid, err := uuid.NewUUID()
	_uuid = uuid.String()
	_uuid = strings.ToLower(_uuid)
	_uuid = strings.ReplaceAll(_uuid, "-", "")
	return
}
