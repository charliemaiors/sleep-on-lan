Write-Host "Checking if Chocolatey is installed"

if ((Get-Command "choco" -ErrorAction SilentlyContinue) -eq $null) 
{ 
   Write-Host "Unable to find Chocolatey (choco) in your PATH, please install it from https://chocolatey.org/"
   exit
}

Write-Host "Checking current priviledges"
If (-NOT ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole(`
[Security.Principal.WindowsBuiltInRole] "Administrator"))
{
    Write-Warning "You do not have Administrator rights to run this script!`nPlease re-run this script as an Administrator!"
    exit
}

Write-Host "Check if nssm is installed"

if ((Get-Command "nssm" -ErrorAction SilentlyContinue) -eq $null)
{
    choco install -y nssm
}

if (Test-Path "$env:GOPATH\bin\sleeponlan.exe")
{
    $path="$env:GOPATH\bin\sleeponlan.exe"
} else {
    $path = Read-Host 'Please insert the full path where sleeponlan.exe is installed'
}

$port = Read-Host 'Please insert the port number where you want to map your service'
nssm install sleeponlan "$path"
nssm set sleeponlan AppParameters "--port $port"
nssm set sleeponlan DisplayName "Sleep on Lan service"
nssm set sleeponlan Description "Simple Sleep on Lan service"
nssm start sleeponlan
