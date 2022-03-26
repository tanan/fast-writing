package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm/clause"
)

type SQLHandler struct {
	Conn *gorm.DB
}

func NewSQLHandler(connInfo string, verbose bool) (*SQLHandler, error) {
	db, err := gorm.Open(mysql.Open(connInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	//if verbose {
	//	db.LogMode(true)
	//}

	//db.DB().SetConnMaxLifetime(time.Minute * 3)
	//db.DB().SetMaxOpenConns(10)
	//db.DB().SetMaxIdleConns(10)

	return &SQLHandler{
		Conn: db,
	}, nil
}

func (h *SQLHandler) Close() error {
	//return h.Conn.Close()
	return nil
}
