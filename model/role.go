package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type Role string

const (
	RoleMember    Role = "MEMBER"
	RoleModerator Role = "MODERATOR"
	RoleAdmin     Role = "ADMIN"
)

func (r Role) Value() (driver.Value, error) {
	switch r {
	case RoleMember, RoleModerator, RoleAdmin:
		return string(r), nil
	default:
		return nil, errors.New("invalid role value")
	}
}

func (r *Role) Scan(value interface{}) error {
	var status Role

	if value == nil {
		*r = RoleMember
		return nil
	}

	v, ok := value.(string)
	if !ok {
		return errors.New("role is not string")
	}

	status = Role(v)
	switch status {
	case RoleMember, RoleModerator, RoleAdmin:
		*r = status
		return nil
	default:
		return fmt.Errorf("invalid role value :%s", v)
	}
}
