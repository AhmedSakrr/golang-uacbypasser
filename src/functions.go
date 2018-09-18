package guacbypasser

import (
	"log"
	"syscall"

	w32 "github.com/JamesHovious/w32"
	"golang.org/x/sys/windows/registry"
)

var (
	_ErrCode_RegistryGlobal = 0x1
	_ErrCode_ExitProcess    = 0x1
	_ErrCode_ShellExecute   = 0x1
)

func W32_RegistryChecker(key32 registry.Key, keyPath, valuePath string) (err error, code int) {
	var wkey32 registry.Key
	wkey32, err = registry.OpenKey(
		key32, keyPath, registry.QUERY_VALUE|registry.ALL_ACCESS)
	if err != nil {
		log.Println(err)
	}

	_, _, err = wkey32.GetStringValue(valuePath)
	if err != nil {
		log.Println(err)
	}

	return err, _ErrCode_RegistryGlobal
}

func W32_Runas(exePath, key32, key64 string) (err error, code int) {
	err = w32.ShellExecute(
		w32.HWND(0),
		string("runas"),
		string(exePath),
		string(""),
		string(""),
		int(1),
	)
	if err != nil {
		log.Println(err)
	}

	return err, _ErrCode_ShellExecute
}

func W32_Terminate() (err error, code int) {
	var kernel32 syscall.Handle
	kernel32, err = syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		log.Println(err)
	}
	defer syscall.FreeLibrary(kernel32)

	var proc32 uintptr
	proc32, err = syscall.GetProcAddress(kernel32, "ExitProcess")
	if err != nil {
		log.Println(err)
	}

	_, _, err = syscall.Syscall(
		uintptr(proc32),
		1,
		0,
		0,
		0,
	)

	if err != nil {
		_ErrCode_ShellExecute = 0x0
	}

	return err, _ErrCode_ShellExecute
}


func recovery() {
	var (
		rType   string = "_0rec"
		rMethod string = "recovery"
	)

	var err = recover()
	fmt.Printf("<&recovery: type=%s, method=%s, recovery=true, err=%s>",
		rType,
		rMethod,
		err,
	)
}
