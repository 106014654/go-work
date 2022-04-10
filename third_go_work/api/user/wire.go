package user
import "github.com/google/wire"

func InitializeShop()  {
	wire.Build()
	return   // 返回值shop{}或nil都行 没啥用
}
