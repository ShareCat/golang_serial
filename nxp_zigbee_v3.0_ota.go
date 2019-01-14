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

/**
 * 错误检查
 */
func check(e error) {
    if e != nil {
        panic(e)
    }
}

/*
 * 判断一个文件夹或文件是否存在
 * 1 如果返回的错误为nil,说明文件或文件夹存在
 * 2 如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
 * 3 如果返回的错误为其它类型,则不确定是否在存在
 */
func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}

var ota_file_name string = "zigbee.ota"
// 写入ota文件的特征头
    var ota_head = []byte{  0x1E, 0xF1, 0xEE, 0x0B, 0x00, 0x01, 0x38, 0x00,
                            0x00, 0x00, 0x37, 0x10, 0x01, 0x01, 0x03, 0x01,
                            0x00, 0x00, 0x02, 0x00, 0x44, 0x52, 0x31, 0x31,
                            0x37, 0x35, 0x72, 0x31, 0x76, 0x32, 0x2D, 0x2D,
                            0x4A, 0x4E, 0x35, 0x31, 0x36, 0x39, 0x30, 0x30,
                            0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
                            0x30, 0x30, 0x30, 0x30, 0x72, 0x1F, 0x03, 0x00,
                            0x00, 0x00, 0xB2, 0x09, 0x03, 0x00}
func make_ota_file() {
    //fmt.Println("make_ota_file")
    var exist bool
    var err error
    exist, err = PathExists(ota_file_name)
    if exist == true {
        //fmt.Println("ota file exist")
        // 如果已经发现当前目录有一个同名的文件，则先删除
        var err1 error
        err1 = os.Remove(ota_file_name)
        check(err1)
        fmt.Println("ota file delete \r\n")
    } else {
        //fmt.Println("ota file not exist")
        //fmt.Println(err)
        err = err // 为了防止报错，提示err这个定义了但是没有使用
    }

    // 可以创建目标ota文件了
    var f *os.File
    f, err = os.Create(ota_file_name) // 创建文件
    check(err)

    // 复制bin文件到ota文件，注意忽略bin文件前4个字节
    var count int = 0
    for _, v := range bin_file_byte {
        count++
        if count > 4 { // 忽略bin文件前4个字节
            ota_head = append(ota_head, v)
        }
    }
    printSlice(ota_head)

    // 文件写入保存
    err = ioutil.WriteFile(ota_file_name, ota_head, 0666) //写入文件(字节数组)
    check(err)

    f.Close()
}

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

var bin_file_byte []byte    // 保存读取的目标bin文件，十六进制
var bin_file_len int        // 目标bin文件的大小
func get_bin_file(name string) {
    //fmt.Println("get_bin_file: ", name)
    f, err := os.Open(name)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    bin_file_byte, err = ioutil.ReadAll(f)
    fmt.Println("Success Open File")
    bin_file_len = len(bin_file_byte)
}

func printSlice(x []byte){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func main() {
    // 读取配置文件，得到目标bin文件名和版本
    bin_file_name, bin_file_version := get_config_file()
    fmt.Print("bin_file_name = ", bin_file_name, "\r\n")      // 打印bin文件名字
    fmt.Print("bin_file_version = ", bin_file_version, "\r\n")// 打印bin文件版本

    // 读取目标bin文件到内存
    get_bin_file(bin_file_name)
    fmt.Printf("bin_file_len = %d \r\n", bin_file_len)
    fmt.Printf("bin_file_byte[0] = 0x%02x \r\n", bin_file_byte[0])
    fmt.Printf("bin_file_byte[1] = 0x%02x \r\n", bin_file_byte[1])
    fmt.Printf("bin_file_byte[%d] = 0x%02x \r\n", bin_file_len - 2, bin_file_byte[bin_file_len - 2])
    fmt.Printf("bin_file_byte[%d] = 0x%02x \r\n", bin_file_len - 1, bin_file_byte[bin_file_len - 1])
    //printSlice(bin_file_byte)
    // 生成目标ota文件
    make_ota_file()



    currentTime := time.Now()   //获取当前时间，类型是Go的时间类型Time
    time.Sleep(time.Duration(1) * time.Second)
    fmt.Println(currentTime)
}