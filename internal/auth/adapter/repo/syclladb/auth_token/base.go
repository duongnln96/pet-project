package auth_token

import (
	"time"

	"github.com/duongnln96/blog-realworld/internal/auth/core/port"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/google/wire"

	scylladbAdapter "github.com/duongnln96/blog-realworld/pkg/adapter/scylladb"
)

var _ port.AuthTokenRepoI = (*repoManager)(nil)

var RepoManagerSet = wire.NewSet(NewRepoManager)

type repoManager struct {
	db scylladbAdapter.ScyllaDBAdaterI
}

func NewRepoManager(
	db scylladbAdapter.ScyllaDBAdaterI,
) port.AuthTokenRepoI {
	return &repoManager{
		db: db,
	}
}

// GoogleUUIDMarshaler is a custom marshaller for github.com/google/uuid.UUID
type GoogleUUIDMarshaler struct {
	uuid.UUID
}

// MarshalCQL converts the custom UUID to its CQL representation
func (u GoogleUUIDMarshaler) MarshalCQL(info gocql.TypeInfo) ([]byte, error) {
	return u.UUID.MarshalBinary()
}

// GoogleUUIDUnmarshaller is a custom unmarshaller for github.com/google/uuid.UUID
type GoogleUUIDUnmarshaller struct {
	uuid.UUID
}

// UnmarshalCQL converts the CQL representation to the custom UUID type
func (u *GoogleUUIDUnmarshaller) UnmarshalCQL(info gocql.TypeInfo, data []byte) error {
	parsedUUID, err := uuid.FromBytes(data)
	if err != nil {
		return err
	}
	u.UUID = parsedUUID
	return nil
}

type AuthToken struct {
	ID          GoogleUUIDMarshaler
	UserID      GoogleUUIDMarshaler
	DeviceID    GoogleUUIDMarshaler
	UserAgent   string
	JwtToken    string
	RemoteIP    string
	Status      string
	ExpiredDate time.Time

	CreatedDate *time.Time
	UpdatedDate *time.Time
}

type AuthTokenRow struct {
	ID          GoogleUUIDUnmarshaller
	UserID      GoogleUUIDUnmarshaller
	DeviceID    GoogleUUIDUnmarshaller
	UserAgent   string
	JwtToken    string
	RemoteIP    string
	Status      string
	ExpiredDate time.Time

	CreatedDate *time.Time
	UpdatedDate *time.Time
}
