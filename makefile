benchmark:
	go test ./pkg -bench=.

b64c:
	go run examples/b64c.go

rsa:
	go run examples/rsa.go

aes:
	go run examples/aes.go

serialize:
	go run examples/serialize.go