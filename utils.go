package main

import (
	"fmt"
	"math/rand"
	"time"
)

func 恐慌(v interface{}) {
	panic(v)
}

func 换行打印(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}
func 格式化打印(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

// 随机从[0,n)中选一个整数，如果 n <= 0 将发生运行时恐慌
var 随机 = rand.New(rand.NewSource(time.Now().UnixNano())).Intn

type 整数 = int
type 字符串 = string
