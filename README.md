# go-HoneyPot
A honeypot server written in Go. 

go-HoneyPot listens on specified ports for any communication. When an attacker attempts to send data on one of these ports it will log relevant details to a database.

## Running go-HoneyPot

1. `git clone https://github.com/Mojachieee/go-HoneyPot`
2. `cd go-HoneyPot`
3. `go mod tidy`
4. `go build .`
4. Create a config.json file. Formatted as follows:
```json
{
    "tcp": {
        "ports": [
            "1220", "5777"
        ]
    }
}
```

5. `./go-HoneyPot`

## Log Sample
```
$ ./go-honeypot             
Listening on tcp port: 5777
Listening on tcp port: 1220
Handle Conection received: connection
Date: 20230727141424, InIp: 127.0.0.1, InPort: 51934, DestIP: 127.0.0.1, DestPort: 5777
Received data from 127.0.0.1:51934, of length 41 data is SSH-2.0-OpenSSH_8.2p1 Ubuntu-4ubuntu0.8
Handle Conection received: connection
Date: 20230727141438, InIp: 127.0.0.1, InPort: 35098, DestIP: 127.0.0.1, DestPort: 5777
2023/07/27 14:14:48 EOF
```
