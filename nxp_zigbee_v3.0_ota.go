package main

import (
    "fmt"
    "os"
    "io"
    "io/ioutil"
    "bufio"
    "strings"
    "time"
)

/*
 * 从配置文件获取目标bin文件的名字和版本
 */
func get_config_file() (string, string) {
    f, err := os.Open("nxp_zigbee_v3.0_ota_config.txt")
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
            break; // 后面的不需要读取了
        }
    }
    //fmt.Printf("line_count = %d \r\n", line_count)    // 打印读取文件的行数
    //fmt.Print("bin_file_name = ", bin_file_name)      // 打印bin文件名字
    //fmt.Print("bin_file_version = ", bin_file_version)// 打印bin文件版本
    // 去掉每行末尾的"\r\n"得到真实的字符串
    bin_file_name = strings.Replace(bin_file_name, "\r", "", -1)
    bin_file_name = strings.Replace(bin_file_name, "\n", "", -1)
    bin_file_version = strings.Replace(bin_file_version, "\r", "", -1)
    bin_file_version = strings.Replace(bin_file_version, "\n", "", -1)
    return bin_file_name, bin_file_version
}

func get_bin_file(name string) {
    //fmt.Println("get_bin_file: ", name)
    f, err := os.Open(name)
     if err != nil {
        panic(err)
     }
     defer f.Close()

     var file_byte []byte
     file_byte, err = ioutil.ReadAll(f)
     fmt.Println("Success Open File")
     fmt.Printf("file_byte[0] = %x \r\n", file_byte[0])
     fmt.Printf("file_byte[0] = %x \r\n", file_byte[1])
     file_len := len(file_byte)
     fmt.Printf("file_len = %d \r\n", file_len)
}

func main() {
    bin_file_name, bin_file_version := get_config_file()
    fmt.Print("bin_file_name = ", bin_file_name, "\r\n")      // 打印bin文件名字
    fmt.Print("bin_file_version = ", bin_file_version, "\r\n")// 打印bin文件版本
    get_bin_file(bin_file_name)



    currentTime := time.Now()   //获取当前时间，类型是Go的时间类型Time
    time.Sleep(time.Duration(1) * time.Second)
    fmt.Println(currentTime)
}
