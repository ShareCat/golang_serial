/*
    Author:     Cat(孙关平)
    Version:    V1.0
    Date:       2019-02-18
    E-mail:     843553493@qq.com
*/

package main

import (
    "fmt"
    "os"
    "io"
    "io/ioutil"
    "bufio"
    "strings"
)

/**
 * 错误检查
 */
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func printSlice(x []byte){
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
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

var config_file_name string = "bootloader_application_merge_config.txt"    // 配置文件名
var new_bin_file_name string = "CMCC-NG1A.bin"
var app_bin_file_name, boot_bin_file_name string
var boot_area_max int = 0x5000

/*
 * 从配置文件获取目标bin文件的文件名
 */
func get_config_file() (string, string) {
    f, err := os.Open(config_file_name)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    line_count := 0
    var app_bin_file_name string
    var boot_bin_file_name string
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
            app_bin_file_name = line
        } else if line_count == 2 {
            // 第二行是bin文件版本
            boot_bin_file_name = line
            break; // 后面的不需要读取了
        }
    }
    //fmt.Printf("line_count = %d \r\n", line_count)        // 打印读取文件的行数
    //fmt.Print("app_bin_file_name = ", app_bin_file_name)  // 打印bin文件名字
    //fmt.Print("boot_bin_file_name = ", boot_bin_file_name)// 打印bin文件名字
    // 去掉每行末尾的"\r\n"得到真实的字符串
    app_bin_file_name = strings.Replace(app_bin_file_name, "\r", "", -1)
    app_bin_file_name = strings.Replace(app_bin_file_name, "\n", "", -1)
    boot_bin_file_name = strings.Replace(boot_bin_file_name, "\r", "", -1)
    boot_bin_file_name = strings.Replace(boot_bin_file_name, "\n", "", -1)
    return app_bin_file_name, boot_bin_file_name
}

var app_bin_file_byte []byte    // 保存读取的目标bin文件，十六进制
var app_bin_file_len int        // 目标bin文件的大小
var boot_bin_file_byte []byte   // 保存读取的目标bin文件，十六进制
var boot_bin_file_len int       // 目标bin文件的大小
var new_bin_file_byte []byte    // 保存读取的目标bin文件，十六进制

/*
 * 读取bin文件到内存
 * name:        要读取的文件名字
 * buff_byte:   读取的文件的内容，十六进制
 * buff_len:    读取的文件的字节数
 */
func get_bin_file(name string) ([]byte, int) {
    //fmt.Println("get_bin_file: ", name)
    var buff_byte []byte
    var buff_len int

    f, err := os.Open(name)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    buff_byte, err = ioutil.ReadAll(f)
    buff_len = len(buff_byte)

    //fmt.Println("\r\n")
    //fmt.Println("get_bin_file OK: ", name)
    //fmt.Printf("buff_len = %d \r\n", buff_len)
    //fmt.Printf("buff_byte[0] = 0x%02x \r\n", buff_byte[0])
    //fmt.Printf("buff_byte[1] = 0x%02x \r\n", buff_byte[1])
    //fmt.Printf("buff_byte[%d] = 0x%02x \r\n", buff_len - 2, buff_byte[buff_len - 2])
    //fmt.Printf("buff_byte[%d] = 0x%02x \r\n", buff_len - 1, buff_byte[buff_len - 1])
    //fmt.Println("\r\n")
    return buff_byte, buff_len
}

/*
 * 合成boot和app文件
 */
func get_new_bin_file(name string) {
    fmt.Println("get_new_bin_file")
    var exist bool
    var err error
    exist, err = PathExists(name)
    if exist == true {
        //fmt.Println("file exist")
        // 如果已经发现当前目录有一个同名的文件，则先删除
        var err1 error
        err1 = os.Remove(name)
        check(err1)
        //fmt.Println("file delete \r\n")
    } else {
        //fmt.Println("file not exist")
        //fmt.Println(err)
        err = err // 为了防止报错，提示err这个定义了但是没有使用
    }

    // 可以创建目标文件了
    var f *os.File
    f, err = os.Create(name) // 创建文件
    check(err)

    // 复制boot_bin文件到目标文件
    var count int = 0
    for _, v := range boot_bin_file_byte {
        count++
        new_bin_file_byte = append(new_bin_file_byte, v)
    }
    //printSlice(new_bin_file_byte) // 打印new_bin_file_byte里面的内容，十进制形式

    // boot区域大小固定为0x5000，因此多余部分补充0xff
    count = boot_area_max - boot_bin_file_len
    if (count != 0) {
        for ; count > 0; count -= 1 {
            new_bin_file_byte = append(new_bin_file_byte, 0xff)
        }
    }
    //printSlice(new_bin_file_byte) // 打印new_bin_file_byte里面的内容，十进制形式

    // 复制app_bin文件到目标文件
    count = 0
    for _, v := range app_bin_file_byte {
        count++
        new_bin_file_byte = append(new_bin_file_byte, v)
    }
    //printSlice(new_bin_file_byte) // 打印new_bin_file_byte里面的内容，十进制形式

    // 文件写入保存
    err = ioutil.WriteFile(name, new_bin_file_byte, 0666) //写入文件(字节数组)
    check(err)

    f.Close()
}

func main() {
    // 读取配置文件，得到目标app和boot的文件名
    app_bin_file_name, boot_bin_file_name = get_config_file()
    fmt.Print("app_bin_file_name = ", app_bin_file_name, "\r\n")    // 打印bin文件名字
    fmt.Print("boot_bin_file_name = ", boot_bin_file_name, "\r\n")  // 打印bin文件名字
    fmt.Print("new_bin_file_name = ", new_bin_file_name, "\r\n")    // 打印bin文件名字

    // 读取目标bin文件到内存
    app_bin_file_byte, app_bin_file_len = get_bin_file(app_bin_file_name)
    boot_bin_file_byte, boot_bin_file_len = get_bin_file(boot_bin_file_name)
    fmt.Println("\r\n")
    fmt.Printf("boot_bin_file_len = %d \r\n", boot_bin_file_len)
    fmt.Printf("boot_bin_file_byte[0] = 0x%02x \r\n", boot_bin_file_byte[0])
    fmt.Printf("boot_bin_file_byte[1] = 0x%02x \r\n", boot_bin_file_byte[1])
    fmt.Printf("boot_bin_file_byte[%d] = 0x%02x \r\n", boot_bin_file_len - 2, boot_bin_file_byte[boot_bin_file_len - 2])
    fmt.Printf("boot_bin_file_byte[%d] = 0x%02x \r\n", boot_bin_file_len - 1, boot_bin_file_byte[boot_bin_file_len - 1])
    fmt.Println("\r\n")
    fmt.Printf("app_bin_file_len = %d \r\n", app_bin_file_len)
    fmt.Printf("app_bin_file_byte[0] = 0x%02x \r\n", app_bin_file_byte[0])
    fmt.Printf("app_bin_file_byte[1] = 0x%02x \r\n", app_bin_file_byte[1])
    fmt.Printf("app_bin_file_byte[%d] = 0x%02x \r\n", app_bin_file_len - 2, app_bin_file_byte[app_bin_file_len - 2])
    fmt.Printf("app_bin_file_byte[%d] = 0x%02x \r\n", app_bin_file_len - 1, app_bin_file_byte[app_bin_file_len - 1])
    fmt.Println("\r\n")

    if (boot_area_max < boot_bin_file_len) {
        // 提示boot文件超出范围，并退出
        panic("boot_file > 0x5000 kbyte")
    }

    // 生成目标文件
    get_new_bin_file(new_bin_file_name)
}
