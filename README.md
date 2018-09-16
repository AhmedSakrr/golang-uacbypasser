# golang-uacbypasser
Collection bypass techiques on Golang<br/>
Author - 0x9ef<br/>

![Alt text](https://media.giphy.com/media/3BZDbv9pe7vMKJrV0f/giphy.gif "Work")

## Payloads list:
  1. **eventvwr.exe payload** (modifying primary DEFAULT registry key to payload value, start eventvwr.exe) | **(worked)**
  2. **userinit.exe payload** (modifying primary Userinit registry key to "%systemroot%\system32\userinit.exe, %payload%", system rebooting) | **(worked)**
  3. **schtasks.exe payload** (creating xml template on "%appdata%\Microsoft\Temp\elevator.xml", create new xml task for "OneDriveUpdate") | **(worked)**
  4. **silentcleanup payload** (creating new ABSOLUTE registry key, seting "windir" value to "cmd /k %payload%", run schtasks process "/Run /TN \Microsoft\Windows\DiskCleanup\SilentCleanup /I") | **(in developing)**
  5. **hkcu runer (OneDriveUpdate)** (creating/modifying ABSOLUTE registry key, seting "OneDriveUpdate" value to payload path) | **(worked)**
  6. **hklm runer (OneDriveUpdate)** (creating/modifying ABSOLUTE registry key, seting "OneDriveUpdate" value to payload path) | **(worked)**
  7. **computerdefaults.exe** (loading kernel32.dll library and seting Wow64DisableWow64FsRedirection to 1, creating/modifying ABSOLUTE registry key, seting DEFAULT to payload path, DelegateExecute to None, executing "start computerdefaults.exe" and deleting registry key) | **(worked)**
  8. **slui.exe** (creating/modifying ABSOLUTE registry key, seting DEFAULT value to payload path, "DelegateExecute" to None, executing "slui.exe") | **(worked)**
 
## How to build: 
  1. `set CGO_ENABLED=0`
  2. `go build -v -a -ldflags="-w -s" -o guacbypasser.exe main.go`

# If you find error in the code or you want to support project please commit this changes. 
# **_Support project - BITCOIN: 18YsYvrQhyrtAqUcpTXpHFrQ6RHyd73dS6_**
