# Sleep on Lan Service
--

## WARNING: Do not expose this service on public network
This service allow to remotely turn off, restart, hibernate, sleep your computer inside your lan network.
For example:

```bash
curl -X POST http://your-pc-ip:7740/poweroff #This will shutdown your pc
```
