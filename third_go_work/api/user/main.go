package user

import (
	"fmt"
	"github.com/google/wire"
	"go-work/third_go_work/api/user/biz"
	"go-work/third_go_work/api/user/data"
	"go-work/third_go_work/api/user/service"
)

func main() {
	db := data.NewMysql("root:@tcp(127.0.0.1:3306)/mysql?charset=utf8")
	var qs service.QueryS
	qyres := wire.Build(db,qs.Query())

	cvs := biz.NewData(qyres)

	fmt.Println("final byte data:",cvs.ConvertData())
	
}
