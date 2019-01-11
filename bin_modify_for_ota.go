package main

import (
    "fmt"
    "os"
    "io"
    "bufio"
    "time"
)

/*
 * 从配置文件获取目标bin文件的名字和版本
 */
func get_config_file() (string, string) {
    f, err := os.Open("config.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    line_count := 0
    var bin_file_name string
    var bin_file_version string
    rd := bufio.NewReader(f)
    for {
        line, err := rd.ReadString('\n') // 以'\n'为结束符读入一行
        if err != nil || io.EOF == err {
            break
        }
        //fmt.Println(line) // 打印读取到的一行内容
        line_count++
        if line_count == 1 {
            // 第一行是bin文件名字
            bin_file_name = line
        } else if line_count == 2 {
            // 第二行是bin文件版本
            bin_file_version = line
        }
    }
    //fmt.Printf("line_count = %d \r\n", line_count)    // 打印读取文件的行数
    //fmt.Print("bin_file_name = ", bin_file_name)      // 打印bin文件名字
    //fmt.Print("bin_file_version = ", bin_file_version)// 打印bin文件版本
    return bin_file_name, bin_file_version
}

func get_bin_file(name string) {
    fmt.Println("get_bin_file: ", name)
}

func main() {
    bin_file_name, bin_file_version := get_config_file()
    fmt.Print("bin_file_name = ", bin_file_name)      // 打印bin文件名字
    fmt.Print("bin_file_version = ", bin_file_version)// 打印bin文件版本
    get_bin_file(bin_file_name)



    currentTime := time.Now()   //获取当前时间，类型是Go的时间类型Time
    time.Sleep(time.Duration(1) * time.Second)
    fmt.Println(currentTime)
}
