# gRPC - raw requests

gRPC uses HTTP2 as based transport. The protocol is available [here](https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md).

### Capturing gRPC request with Wireshark

Here are steps to capture request:

* run Wireshark and select an interface to capture traffic
* run both gRPC server and client
* trigger request in client
* find packets sent in HTTP2 protocol that are identified as `Magic`, `SETTINGS`, `HEADERS` and `DATA`.
* copy streams in these packets as a Hex Stream

Remember to choose a proper interface - select `Loopback: lo0` for local server, tunnel interface (e.g. `utun1`) for server with VPN access, or `Wi-Fi: en0` for server with public access.

If you don't see packets in HTTP2 protocol, click "Analyze" -> "Decode As...".
Then, add `TCP port X` with `HTTP2` protocol, where `X` is port of gRPC server (e.g. `8083`).

![wireshark](./images/wireshark.png "Wireshark - capturing HTTP2 traffic")


After the process you should have required data in hex:
```
Magic 505249202a20485454502f322e300d0a0d0a534d0d0a0d0a

SETTINGS 000000040000000000

HEADERS 0000570104000000018386458e629f43accbe8f50ebc8c632d141f418aa0e41d139d09b8f01e675f8b1d75d0620d263d4c4d65647a8d9acac8b4c7602bb6b81690bdff40027465864d833505b11f40899acac8b24d494f6a7f867df7df79d6ed

DATA 00000c00010000000100000000070a05576f726c64
```

### Sending gRPC request

You can use many different tools to send a gRPC request as a client.

##### Node.js

```javascript
var net = require('net');

var client = new net.Socket();

client.connect(8080, 'localhost', function() {
	console.log('Connected to server');

	var magic = new Buffer("505249202a20485454502f322e300d0a0d0a534d0d0a0d0a", "hex");
	client.write(magic);

	var settings = new Buffer("000000040000000000", "hex");
	client.write(settings);

	var headers = new Buffer("0000550104000000018386458c62919aa5e1d7918c65a283ff418aa0e41d139d09b8f01e075f8b1d75d0620d263d4c4d65647a8d9acac8b4c7602bb6b81690bdff40027465864d833505b11f40899acac8b24d494f6a7f867df7dd6dd7ad", "hex");
	client.write(headers);

	var data = new Buffer("00000c00010000000100000000070a05576f726c64", "hex");
	client.write(data);

});

client.on('data', function(data) {
	console.log('Received:\n' + data);
});

client.on('close', function() {
	console.log('Connection closed');
});
```


##### Go

```go
package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
)

const APIURL = "localhost:8080"

func main() {
	conn, err := net.Dial("tcp", APIURL)
	if err != nil {
		fmt.Printf("Error: connect to server: %s\n", err)
		os.Exit(1)
	}

	magic := "505249202a20485454502f322e300d0a0d0a534d0d0a0d0a"
	settings := "000000040000000000"
	headers := "0000550104000000018386458c62919aa5e1d7918c65a283ff418aa0e41d139d09b8f01e075f8b1d75d0620d263d4c4d65647a8d9acac8b4c7602bb6b81690bdff40027465864d833505b11f40899acac8b24d494f6a7f867df7dd6dd7ad"
	data := "00000c00010000000100000000070a05576f726c64"

	packets := []string{magic, settings, headers, data}

	for _, p := range packets {
		raw, err := hex.DecodeString(p)
		if err != nil {
			fmt.Printf("Error: cannot decode hex string: %s\n", err)
			os.Exit(1)
		}

		if _, err := conn.Write(raw); err != nil {
			fmt.Printf("Error: cannot decode hex string: %s\n", err)
			os.Exit(1)
		}
	}

	buf := make([]byte, 1024)

	fmt.Print("Info: received:\n")

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Error: cannot read from socket: %s\n", err)
			os.Exit(1)
		}
		fmt.Print(string(buf[:n]))
	}
}
```

##### netcat command

```
$ (echo -e "\x50\x52\x49\x20\x2a\x20\x48\x54\x54\x50\x2f\x32\x2e\x30\x0d\x0a\x0d\x0a\x53\x4d\x0d\x0a\x0d\x0a\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x55\x01\x04\x00\x00\x00\x01\x83\x86\x45\x8c\x62\x91\x9a\xa5\xe1\xd7\x91\x8c\x65\xa2\x83\xff\x41\x8a\xa0\xe4\x1d\x13\x9d\x09\xb8\xf0\x1e\x07\x5f\x8b\x1d\x75\xd0\x62\x0d\x26\x3d\x4c\x4d\x65\x64\x7a\x8d\x9a\xca\xc8\xb4\xc7\x60\x2b\xb6\xb8\x16\x90\xbd\xff\x40\x02\x74\x65\x86\x4d\x83\x35\x05\xb1\x1f\x40\x89\x9a\xca\xc8\xb2\x4d\x49\x4f\x6a\x7f\x86\x7d\xf7\xdd\x6d\xd7\xad\x00\x00\x0c\x00\x01\x00\x00\x00\x01\x00\x00\x00\x00\x07\x0a\x05\x57\x6f\x72\x6c\x64"; sleep 1) | nc localhost 8081
```


---------

Sources:
* http://www.juhonkoti.net/2015/11/26/using-haproxy-to-do-health-checks-to-grpc-services
* https://linux.die.net/man/1/nc
* https://njp.io/grpcc-a-simple-command-line-client-for-grpc-services/
