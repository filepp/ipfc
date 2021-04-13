SHELL=/usr/bin/env bash

all: ipfc inspector

unexport GOFLAGS

.PHONY: all ipfc inspector

ipfc:
	rm -f ipfc
	go build $(GOFLAGS) -o ipfc ./cmd/ipfc

inspector:
	rm -f inspector
	go build $(GOFLAGS) -o inspector ./cmd/inspector


.PHONY: clean
clean:
	rm -f ipfc
	rm -f inspector

