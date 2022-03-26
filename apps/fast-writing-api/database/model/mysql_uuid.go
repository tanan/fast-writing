package model

import (
	"database/sql/driver"
	"github.com/google/uuid"
)

type MysqlUUID uuid.UUID

func StringToMysqlUUID(s string) (MysqlUUID, error) {
	id, err := uuid.Parse(s)
	return MysqlUUID(id), err
}

func (u MysqlUUID) String() string {
	return uuid.UUID(u).String()
}

func (u MysqlUUID) GormDataType() string {
	return "binary(16)"
}

func (u MysqlUUID) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(u)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

func (u *MysqlUUID) UnmarshalJSON(by []byte) error {
	s, err := uuid.ParseBytes(by)
	*u = MysqlUUID(s)
	return err
}

func (u *MysqlUUID) Scan(value interface{}) error {

	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*u = MysqlUUID(parseByte)
	return err
}

func (u MysqlUUID) Value() (driver.Value, error) {
	return uuid.UUID(u).MarshalBinary()
}
