package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"time"
)

/**
	redis benchmark
byte   request per second
10			11614
20			11325
50			8960
100			8071
200			5243
1K			1316
5K			272

	redis	key-size
num			total byte (1134768	every key byte
1w			2158912			215.891
2w			3183184			159.159
5w			6128784			122.575
10w			11124480		111.244
20w			21109120		105.545
50w			50022592		100.045
 */

var strByte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var strByteLen = len(strByte)

func RandString(length int) []byte {

	bytes := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		bytes[i] = strByte[r.Intn(strByteLen)]
	}

	return bytes
}


func main()  {
	conn , err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil{
		fmt.Println(err)
		return
	}

	defer conn.Close()
	for i:= 0; i< 500000;i++{
		str := fmt.Sprint("key", i)
		_,err = conn.Do("Set",str,RandString(rand.Intn(30)));
	}

	fmt.Println(conn)
}