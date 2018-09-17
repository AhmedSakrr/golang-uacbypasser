package guacbypasser

import "fmt"

type __W32_ComputerDefaults__ struct {
	Type   string
	Method string
	Scode  [8]uintptr

	Location            string
	_K_ComputerDefaults byte
}

type __W32_Eventvwr__ struct {
	Type   string
	Method string
	Scode  [4]uintptr

	Location    string
	_K_Eventvwr byte
}

type __W32_Fodhelper__ struct {
	Type   string
	Method string
	Scode  [4]uintptr

	Location     string
	_K_Fodhelper byte
}

type __W32_HKCU_Runer__ struct {
	Type   string
	Method string
	Scode  [4]uintptr

	Location      string
	_K_HKCU_Runer byte
}

type __W32_HKLM_Runer__ struct {
	Type   string
	Method string
	Scode  [4]uintptr

	Location      string
	_K_HKLM_Runer byte
}

type __W32_Schtasks__ struct {
	Type   string
	Method string
	Scode  [4]uintptr

	Location    string
	_K_Schtasks byte
}

type __W32_Sdcltcontrol__ struct {
	Type   string
	Method string
	Scode  [4]uintptr

	Location        string
	_K_Sdcltcontrol byte
}

type __W32_SilentCleanUp__ struct {
	Type   string
	Method string
	Scode  [2]uintptr

	Location         string
	_K_SilentCleanUp byte
}

type __W32_Slui__ struct {
	Type   string
	Method string
	Scode  [2]uintptr

	Location string
	_K_Slui  byte
}

type __W32_Userinit__ struct {
	Type   string
	Method string
	Scode  [2]uintptr

	Location    string
	_K_Userinit byte
}

type __W32_WMIC__ struct {
	Type   string
	Method string
	Scode  [2]uintptr

	Location string
	_K_Wmic  byte
}

func recovery() {
	var err = recover()
	fmt.Printf("%s\r\n",
		err,
	)
}
