package main

import (
    "fmt"
    "reflect" // 这里引入reflect模块
)

type testConst int
//const C1 testConst = 1 "C1"

type User struct {
    Name   string "user name" //这引号里面的就是tag
    Passwd string "user passsword"
}

func main() {
    user := &User{"chronos", "pass"}
    s := reflect.TypeOf(user).Elem() //通过反射获取type定义
    for i := 0; i < s.NumField(); i++ {
        fmt.Println(s.Field(i).Tag) //将tag输出出来
    }
	
	bin := []byte{0,1,2,3}
	
	var i interface{} = bin
	
	if v, ok := i.([]byte); ok {
		fmt.Println("is slice", v, len(v))
	}
}