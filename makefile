

bubble-playground:
	go build -o ./bin ./cmd/bubble-playground

bubble-playground-windows:
	GOOS="windows" go build -o ./bin ./cmd/bubble-playground
