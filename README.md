# ncbigo

WIP - Towards a Go (Golang) Asn.1 compiler and toolkit.


### Objectives
	- Support the NCBI biological data Asn.1 specifications and sequence extensions.
	- Support the NCBI Asn.1 text and BER encoded data dialects.
	- Support the Blueprint/BIND biological Asn.1 specification and data.
	- Support telecommunications Asn.1 data and specifications where possible.
	
### Contributors
	Christopher Hogue (Ericsson)
	Johnathan Kans (NCBI)
	Andrew Hume (Ericsson)


### Go environment requirements

	Install go
	Install golint from https://github.com/golang/lint
	Ensure golint is in your PATH

### Get & make

	go get github.com/cwvhogue/ncbigo
	cd $GOPATH/src/github.com/cwvhogue/ncbigo
	make
	make test
	

