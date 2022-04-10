package biz

import (
	"encoding/json"
)

type SendData interface {
	convertData() []byte
}

type data struct {
	res string
}

func (dt *data)ConvertData() []byte  {
	dat , _ := json.Marshal(dt.res)
	return dat
}

func NewData(str string) *data {
	return &data{res:str}

}
