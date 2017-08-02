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
