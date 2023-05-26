package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	dll, err := syscall.LoadDLL("user32.dll")
	if err != nil {
		panic(err)
	}

	messageBoxW, err := dll.FindProc("MessageBoxW")
	if err != nil {
		panic(err)
	}
	fmt.Println("MessageBoxW found:", messageBoxW.Name)

	executable := "C:\\Windows\\System32\\notepad.exe"
	var si syscall.StartupInfo
	var pi syscall.ProcessInformation

	err = syscall.CreateProcess(
		syscall.StringToUTF16Ptr(executable),
		nil,
		nil,
		nil,
		false,
		0,
		nil,
		nil,
		&si,
		&pi,
	)
	if err != nil {
		panic(err)
	}

	hwnd := pi.Process
	fmt.Println("Process handle:", hwnd)

	title := "YOU HAVE BEEN HACKED"
	message := "JUST KIDDING"
	flags := uint32(0)

	ret, _, err := messageBoxW.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(message))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(flags),
	)
	if ret == 0 {
		panic(err)
	}

	fmt.Println("MessageBox returned:", ret)
}
