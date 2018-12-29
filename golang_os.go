package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    var err int
    count := len(os.Args)
    fmt.Printf("args count = %d ", count)
    fmt.Println(" ")

    if count == 1 {
        /* 没有带任何参数 */
        fmt.Println("no args")
        //fmt.Println(os.Args[0])// args 第一个片 是文件路径
        os.Exit(err)
    }

    /*
        打印所有的参数
    */
    i := 1
    for i < count {
        fmt.Println(os.Args[i])
        i++
    }

    currentTime := time.Now()   //获取当前时间，类型是Go的时间类型Time
    time.Sleep(time.Duration(1) * time.Second)
    fmt.Println(currentTime)
}