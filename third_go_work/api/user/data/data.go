package data


import (
	"database/sql"
)

type MysqlConn struct {
	str string
}

func (msql *MysqlConn) Conn() *sql.DB  {
	db , _ := sql.Open("mysql",msql.str)
	if err:= db.Ping(); err != nil {
		return nil
	}
	defer db.Close()
	return db
}

func NewMysql(str string) *MysqlConn  {
	return &MysqlConn{
		str : str,
	}
}
