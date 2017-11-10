# Sleep on Lan Service
[![Build Status](https://travis-ci.org/charliemaiors/sleep-on-lan.svg?branch=master)](https://travis-ci.org/charliemaiors/sleep-on-lan)
## WARNING: Do not expose this service on public network
This service allow to remotely turn off, restart, hibernate, sleep your computer inside your lan network.
For example:

```bash
curl -X POST http://your-pc-ip:7740/poweroff #This will shutdown your pc
```
## Installation

You can download binaries from [release](https://github.com/charliemaiors/sleep-on-lan/releases) or run:

```bash
go get github.com/charliemaiors/sleep-on-lan/
```

and run:
```bash
sudo $GOPATH/bin/sleeponlan
```