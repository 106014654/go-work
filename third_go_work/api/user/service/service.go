package service

import (
	"database/sql"
	"fmt"
	"go-work/third_go_work/api/user/data"
)

type Info struct {
	label string
	wcnid string
	storeid string
}



type QueryS interface {
	Query() string
}

type conn struct {
	db *sql.DB
}

func (con *conn) Query(db *data.MysqlConn) *Info {
	con.db = db.Conn()
	return con.queryDB()
}


func (con *conn)queryDB() *Info {
	var info Info
	qerror := con.db.QueryRow(`select * from t_test where id > ? limit 1`,1000).Scan(
		&info.label,&info.wcnid,&info.storeid)

	switch  {
	case qerror == sql.ErrNoRows: //属于正常错误
		fmt.Printf("result has no rows")
	case qerror != nil:
		fmt.Printf("result has faild")
	}
	return &info
}