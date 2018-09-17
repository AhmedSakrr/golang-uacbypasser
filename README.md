# golang-uacbypasser
Collection bypass techiques on Golang<br/>
Author - 0x9ef<br/>

![Alt text](https://media.giphy.com/media/3BZDbv9pe7vMKJrV0f/giphy.gif "Work")

## Payloads list:
  1. **computerdefaults.exe payload** (modifying/creating primary DEFAULT value to payload path, "DelegateExecute" to None) | **(in processing)**
  2. **eventvwr.exe payload** (modifying primary DEFAULT value to payload path, "DelegateExecute" to None, start eventvwr.exe) | **(in processing)**
  3. **fodhelper.exe payload** (modifying primary DEFAULT value to payload path, "DelegateExecute" to None, start finhelper.exe) | **(in processing)**
  4. **hkcu runer (OneDriveUpdate) payload** (creating/modifying ABSOLUTE registry key, seting "OneDriveUpdate" value to payload path) | **(in processing)**
  5. **hklm runer (OneDriveUpdate) payload** (creating/modifying ABSOLUTE registry key, seting "OneDriveUpdate" value to payload path) | **(in processing)**
  6. **ifeo payload** (creating/modifying "magnify"|"Debugger" values, seting "magnify" to "magnifierpane", seting "Debugger" value to payload path) | **(in processing)**registry key) | **(in processing)**
  7. **schtasks.exe payload** (creating xml template on "%appdata%\Microsoft\Temp\elevator.xml", create new xml task for "OneDriveUpdate") | **(in processing)**
  8. **sdkltcontrol.exe** (modifying/creating primary DEFAULT value to payload path, spawning new cmd command and executing it) | **(in processing)**
  9. **silentcleanup payload** (creating new ABSOLUTE registry key, seting "windir" value to "cmd /k %payload%", run schtasks process "/Run /TN \Microsoft\Windows\DiskCleanup\SilentCleanup /I") | **(worked)**
  10. **slui.exe payload** (creating/modifying ABSOLUTE registry key, seting DEFAULT value to payload path, "DelegateExecute" to None, executing "slui.exe") | **(in processing)**
  11. **userinit.exe payload** (modifying primary "Userinit" value to "%systemroot%\system32\userinit.exe, payload path, system rebooting) | **(in processing)**
  12. **wmic.exe payload** (spawn cmd commands and executing it) | **(in processing)**
 
## How to build: 
  1. `set CGO_ENABLED=0`
  2. `go build -v -a -ldflags="-w -s" -o guacbypasser.exe main.go`

# If you find error in the code or you want to support project please commit this changes. 
# **_Support project - BITCOIN: 18YsYvrQhyrtAqUcpTXpHFrQ6RHyd73dS6_**
