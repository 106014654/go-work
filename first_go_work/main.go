package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	xerros "github.com/pkg/errors"
	"os"
	"runtime"
	"strconv"
)

type MyError struct {
	file string
	line int
	err error
}

func (e *MyError)Error() string  {
	return  "file:"+ e.file + " line:" +strconv.Itoa(e.line)
}

var myError *MyError

func recordErrorInfo(err error) *MyError  {
	if _ ,file , line ,ok := runtime.Caller(0);ok{
		return  &MyError{file: file,line: line,err: err}
	}
	return nil
}

type Info struct {
	label string
	wcnid string
	storeid string
}

var info Info

func conSql() (*sql.DB,error) {
	db , _ := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/mysql?charset=utf8")
	if err:= db.Ping(); err != nil {
		return nil,xerros.Wrap(err,"connect fail")
	}
	defer db.Close()
	return db,nil
}

func queryDB(db *sql.DB) error  {
	qerror := db.QueryRow(`select * from t_test where id > ? limit 1`,1000).Scan(
		&info.label,&info.wcnid,&info.storeid)

	switch  {
	case qerror == sql.ErrNoRows: //属于正常错误
		fmt.Printf("result has no rows")
	case qerror != nil:
		myError =  recordErrorInfo(qerror)
		return xerros.WithMessage(qerror,myError.Error())
	}
	return nil
}

func main()  {
	db ,err := conSql()
	if err != nil{
		fmt.Printf("xxxxxxxxx connect fail %T\n %v\n",xerros.Cause(err),xerros.Cause(err))
		os.Exit(1)
	}

	err = queryDB(db)

	if err != nil{
		fmt.Printf("vvvvvvvvvvvv connect fail  %v\n",err)
		os.Exit(1)
	}

}