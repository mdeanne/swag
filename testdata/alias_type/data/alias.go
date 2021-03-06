package data

import (
	"github.com/integration-system/swag/testdata/alias_type/types"
	"time"
)

type TimeContainer struct {
	Name      types.StringAlias `json:"name"`
	Timestamp time.Time         `json:"timestamp"`
	CreatedAt types.DateOnly    `json:"created_at"`
}
