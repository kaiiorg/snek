OS=linux
ARCH=amd64

build:
	export GOOS=$(OS)
	export GOARCH=$(ARCH)
	go build \
		-o ./bin/snek-$(OS)-$(ARCH) \
		./cmd/snek

bubble-playground:
	go build -o ./bin ./cmd/bubble-playground
