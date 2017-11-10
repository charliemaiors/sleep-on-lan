if (!$env:GOPATH){
    Write-Host "GOPATH is null, please define it and restart installation"
    exit
}

sc.exe create sleeponlan binPath= "$env:GOPATH\bin\sleeponlan.exe" start= auto DisplayName= "Sleep On Lan Service"