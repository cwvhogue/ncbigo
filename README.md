# ncbigo

WIP - Towards a Go (Golang) Asn.1 compiler and toolkit.


### Objectives
	- Support the NCBI biological data Asn.1 specifications and sequence extensions.
	- Support the NCBI Asn.1 text and BER encoded data dialects.
	- Support the Blueprint/BIND biological Asn.1 specification and data.
	- Support telecommunications Asn.1 data and specifications where possible.

### Status

Nothing much to see here yet... Just getting repo organized and adding simple print form (.prt) and BER (.val) Asn.1 data & specifications for testing.
	
### Go environment requirements

	Install go
	Install golint from https://github.com/golang/lint
	Ensure golint is in your PATH

### Get & make

	go get github.com/cwvhogue/ncbigo
	cd $GOPATH/src/github.com/cwvhogue/ncbigo
	make
	make test
	
### Contributors
	Christopher Hogue (Ericsson)
	Jonathan Kans (NCBI)
	Andrew Hume (Ericsson)


