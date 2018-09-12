package main

import (
	"fmt"
	"reflect" // 这里引入reflect模块
	"testing"
)

type testConst int

//const C1 testConst = 1 "C1"

type UserName string

type User struct {
	Name   UserName "user name" //这引号里面的就是tag
	Passwd string   "user passsword"
	Age    int      "user age"
}

func Test1(t *testing.T) {
	user := &User{
		Name:   "chronos",
		Passwd: "pass",
		Age:    100,
	}

	tp := reflect.TypeOf(user)
	t.Logf("reflect.TypeOf(user) = %v", tp)

	s := tp.Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag) //将tag输出出来
	}

	tp = reflect.TypeOf(user.Name)
	t.Logf("reflect.TypeOf(user.Name) = %v", tp)
	t.Logf("tp.Name() = %v", tp.Name())
	t.Logf("tp.String() = %v", tp.String())
}
