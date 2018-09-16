package guacbypasser

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

type __W32_Slui__ struct {
	Type   string
	Method string
	Scode  [2]uintptr

	Location string
	_K_Slui  byte
}

type __W32_SilentCleanUp__ struct {
	Type   string
	Method string
	Scode  [2]uintptr

	Location         string
	_K_SilentCleanUp byte
}

type __W32_Userinit__ struct {
	Type   string
	Method string
	Scode  [2]uintptr

	Location    string
	_K_Userinit byte
}
