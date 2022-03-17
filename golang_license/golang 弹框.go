package main

import (
    "syscall"
    "time"
    "unsafe"
)

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

func main() {
    go func() {
        for {
            ShowMessage2("windows下的另一种DLL方法调用", "HELLO !")
            time.Sleep(3 * time.Second)
        }
    }()
    select {}
}