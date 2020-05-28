# softetherConfig

SoftEther JSON-RPC API Configuration Go CLI

Golang CLI for SoftEther VPN management JSON-RPC API protocol. Can be used for controlling remote server, automation or statistics.

## Usage example

Build first to get binary

```bash
go build -o secgo
```

```bash
./secgo -h
Usage of ./secgo:
  -H string
      SoftEther Hub Name (default "DEFAULT")
  -P string
      SoftEther Admin Password
  -p int
      SoftEther Port number (default 443)
  -s string
      SoftEther Hostname (default "127.0.0.1")
```

## Credits/References

* Credit: [./softetherApi](./softetherApi) package included in this project is a clone of [kiddnoke/SoftetherGo](https://github.com/kiddnoke/SoftetherGo).
