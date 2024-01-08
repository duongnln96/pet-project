package auth_token

import (
	"github.com/scylladb/gocqlx/v2/table"
)

var (
	authenTokenTable = table.New(
		table.Metadata{
			Name:    "auth_token",
			Columns: []string{"id", "user_id", "device_id", "user_agent", "jwt_token", "remote_ip", "status", "expired_date", "created_date", "updated_date"},
			PartKey: []string{"id"},
			SortKey: []string{"expired_date"},
		},
	)
)
