# Sleep on Lan Service
[![Build Status](https://travis-ci.org/charliemaiors/sleep-on-lan.svg?branch=master)](https://travis-ci.org/charliemaiors/sleep-on-lan)
## WARNING: Do not expose this service on public network
This service allow to remotely turn off, restart, hibernate, sleep your computer inside your lan network.
For example:

```bash
curl -X POST http://your-pc-ip:7740/poweroff #This will shutdown your pc
```
## Installation

You can download binaries from [release](https://github.com/charliemaiors/sleep-on-lan/releases), unzip it and run from a powershell or run:

```bash
go get github.com/charliemaiors/sleep-on-lan/
```
and run:
```bash
sudo $GOPATH/bin/sleeponlan
```
## Service

Sleep on lan could be installed also as a service in your os using this syntax:

```bash
./sleeponlan install [--port <custom-port>]
```
on linux, currently on windows the service is installed but it will not start so there is a script inside the script folder (if you get the executable using ```go get```) or inside the windows binary archive from [releases](https://github.com/charliemaiors/sleep-on-lan/releases) page in order to install and run it.
In order to uninstall it you could run 
```bash
./sleeponlan uninstall
```
on linux systems, there is also a temporary (hopefully) uninstall script in windows binary archive or script folder.