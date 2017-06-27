# Make gRPC requests

gRPC uses HTTP2 as based transport. The protocol is available [here](https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md).

### Capturing gRPC request with Wireshark

Here are steps to capture request:

* run Wireshark and select Capture `Loopback: (lo0)`
* run both gRPC server and client
* trigger request in client
* find packets sent in HTTP2 protocol that are identified as `Magic`, `SETTINGS`, `HEADERS` and `DATA`.
* copy streams in these packets as a Hex Stream

If you don't see packets in HTTP2 protocol, click "Analyze" -> "Decode As...".
Then, add `TCP port X` with `HTTP2` protocol, where `X` is port of gRPC server (e.g. `8083`).

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

client.connect(8083, 'localhost', function() {
	console.log('Connected to server');

	var magic = new Buffer("505249202a20485454502f322e300d0a0d0a534d0d0a0d0a", "hex");
	client.write(magic);

	var settings = new Buffer("000000040000000000", "hex");
	client.write(settings);

	var headers = new Buffer("0000570104000000018386458e629f43accbe8f50ebc8c632d141f418aa0e41d139d09b8f01e675f8b1d75d0620d263d4c4d65647a8d9acac8b4c7602bb6b81690bdff40027465864d833505b11f40899acac8b24d494f6a7f867df7df79d6ed", "hex");
	client.write(headers);

	var data = new Buffer("00000c00010000000100000000070a05576f726c64", "hex");
	client.write(data);

});

client.on('data', function(data) {
	console.log('Received: ' + data);
});

client.on('close', function() {
	console.log('Connection closed');
});

```


##### Go

##### nc command

##### grpcc


---------

Sources:
* http://www.juhonkoti.net/2015/11/26/using-haproxy-to-do-health-checks-to-grpc-services
* https://linux.die.net/man/1/nc
* https://njp.io/grpcc-a-simple-command-line-client-for-grpc-services/