package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

func md5_test() {
	str := "123456"
	//MD5方法一
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	fmt.Println(md5str1)
	//MD5方法二
	w := md5.New()
	io.WriteString(w, str)
	//将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))

	fmt.Println(md5str2)

	//结果
	//e10adc3949ba59abbe56e057f20f883e
	//e10adc3949ba59abbe56e057f20f883e
}

func IntPtr(n int) uintptr {
	return uintptr(n)
}

func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

// windows下的另一种DLL方法调用
func ShowMessage2(title, text string) {
	user32dll, _ := syscall.LoadLibrary("user32.dll")
	user32 := syscall.NewLazyDLL("user32.dll")
	MessageBoxW := user32.NewProc("MessageBoxW")
	MessageBoxW.Call(IntPtr(0), StrPtr(text), StrPtr(title), IntPtr(0))
	defer syscall.FreeLibrary(user32dll)
}

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

func main() {
	magic_string_1 := "iA3dC0fM4cF1bF2h"
	magic_string_2 := "oH6)bD1_mE2&sC1)"
	fileName := "license_code.txt"

	var exist bool
	var err error
	exist, err = PathExists(fileName)
	if exist == true {
		// 如果已经发现当前目录有一个同名的文件，则先删除
		var err1 error
		err1 = os.Remove(fileName)
		check(err1)
		fmt.Println(fileName, " delete ok \r\n")
	}

	// 读取注册码
	bytes, err := ioutil.ReadFile("register_code.txt")
	if err != nil {
		str_err := fmt.Sprintf("%s", err)
		ShowMessage2("ERROR!", str_err)
		log.Fatal(err)
	}
	register_code := string(bytes)
	//fmt.Println("Bytes read: ", len(bytes))
	//fmt.Println("String read: ", register_code)

	str_1 := magic_string_1 + register_code
	str_2 := magic_string_2 + register_code
	//fmt.Println("str_1: ", str_1)
	//fmt.Println("str_2: ", str_2)

	w_1 := md5.New()
	io.WriteString(w_1, str_1)
	//将str写入到w中
	md5str1 := fmt.Sprintf("%x", w_1.Sum(nil))

	w_2 := md5.New()
	io.WriteString(w_2, str_2)
	//将str写入到w中
	md5str2 := fmt.Sprintf("%x", w_2.Sum(nil))
	license_code := md5str1 + md5str2
	license_code = strings.ToUpper(license_code)
	fmt.Println("license_code: ", license_code)

	// 生成授权码
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	dstFile.WriteString(license_code)
}
