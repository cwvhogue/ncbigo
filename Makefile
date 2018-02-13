PROJECT=github.com/cwvhogue/ncbigo
VENDOR=$(PROJECT)/vendor

all : install

build:
	go build $(PROJECT)/pkg/...

install : build verifiers 

test: install
	go test $(PROJECT)/pkg/...

verifiers:
	@go vet $(PROJECT)/pkg/...
	@golint $(PROJECT)/pkg/...
