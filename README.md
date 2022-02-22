this library is a fork of go-ethereum json rpc implementations with 2 main differences:

1. Request headers are passed to the context
2. Named parameter is allowed. In go ethereum only positioned parameters are allowed `{"jsonrpc":"2.0","method":"test","params": {},"id":1}	`