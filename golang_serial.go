package main

import (
    "fmt"
    "log"
    "io"
    //"io/ioutil"
    "bufio"
    "serial"
    "os"
    "strings"
)

import "time"

/* log文件名字 */
var filename = "./log1.txt"

func init() {
    fmt.Println("go init!")
}

/**
 * 错误检查
 */
func check(e error) {
    if e != nil {
        panic(e)
    }
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
    var exist = true
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        exist = false
    }
    return exist
}

/**
 * 在文件最后写入
 * fileName:文件名字(带全路径)
 * content: 写入的内容
 */
func appendToFile(fileName string, content string) error {
    // 以只写的模式，打开文件
    f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
    } else {
        // 查找文件末尾的偏移量
        n, _ := f.Seek(0, os.SEEK_END)
        // 从末尾的偏移量开始写入内容
        _, err = f.WriteAt([]byte(content), n)
    }   
    defer f.Close()
    return err
}

/**
 * 检查数据
 */
func check_temperature_humidity() {
    //fmt.Println("check_temperature_humidity")
    var temperature bool = false
    var humidity bool = false

    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    rd := bufio.NewReader(f)
    /* 按行读取log文件 */
    for {
        line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

        if err != nil || io.EOF == err {
            break
        }

        /* 查看温度是否上报 */
        tmp := strings.Contains(line, "ZHA Sensor Task handled!")
        if true == tmp {
            temperature = true
            fmt.Println(line)
        }

        /* 查看湿度是否上报 */
        tmp = strings.Contains(line, "vReportRelativeHumidity is worked!")
        if true == tmp {
            humidity = true
            fmt.Println("vReportRelativeHumidity is worked!")
        }
    }

    /* 查看是否漏报 */
    if temperature != true || humidity != true {
        currentTime := time.Now()   //获取当前时间，类型是Go的时间类型Time
        fmt.Println(currentTime)
        for true {
            /* 停在这里，提示有漏报 */
            fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
            time.Sleep(time.Duration(3600)*time.Second)
        }
    }
}

/**
 * 删除log文件
 */
func log_file_delete() {
    var err1 error

    err1 = os.Remove(filename)
    check(err1)
    fmt.Println("log_file_delete")
}

/**
 * 测试程序
 */
func test_com() {

    /*
        开启串口
    */
    c := &serial.Config{Name: "COM3", Baud: 115200}
    s, err := serial.OpenPort(c)
    if err != nil {
            log.Fatal(err)
    }
    
    n, err := s.Write([]byte("test"))
    if err != nil {
            log.Fatal(err)
    }

    /*
        接收串口数据
    */
    for true {
        buf := make([]byte, 1024)
        n, err = s.Read(buf)
        if err != nil {
            log.Fatal(err)
        }
        //log.Printf("%q", buf[:n])
        fmt.Printf("收到 %d 个字节", n)
        fmt.Println(" ")

        if n == 0 {
            /* 收到数据才需要保存 */
            continue
        }

        /*
            写入文件保存
        */
        var f *os.File
        var err1 error

        /*
            如果log文件不存在就创建
        */
        if checkFileIsExist(filename) { //如果文件存在
            f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
            //fmt.Println(filename + "文件存在")
        } else {
            f, err1 = os.Create(filename) //创建文件
            //fmt.Println(filename + "文件不存在")
        }
        check(err1)

        var str2 = string(buf)  /* 转化成字符串 */
        appendToFile(filename, str2)
        f.Close()

        if true == strings.Contains(str2, "Sleeping") {
            //fmt.Printf("Got Sleeping")
            check_temperature_humidity()
            log_file_delete()
        }

        //var str3 = string("haha\r\n")  /* 转化成字符串 */
        //n, err1 = io.WriteString(f, str3) //写入文件(字符串)
        //check(err1)
        //fmt.Printf("写入 %d 个字节n", n)
        f.Sync()
    }
}

func main() {
    test_com()
}
