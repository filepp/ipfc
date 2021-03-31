SHELL=/usr/bin/env bash

all: ipfc

unexport GOFLAGS

.PHONY: all ipfc

ipfc:
	rm -f ipfc
	go build $(GOFLAGS) -o ipfc ./cmd/ipfc

.PHONY: clean
clean:
	rm -f scanner


