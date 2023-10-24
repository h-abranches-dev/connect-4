SYSTEM ?= "undefined"

GC_BUILD_VERSION = "gc-v0.0.1"

ifeq ($(SYSTEM), gc)
	BUILD_VERSION := $(GC_BUILD_VERSION)
endif

LDFLAGS = "-X 'github.com/version-go/ldflags.buildVersion=$(BUILD_VERSION)'"

build:
	go build -ldflags $(LDFLAGS) -o bin/$(SYSTEM) ./entrypoints/$(SYSTEM)

run_gc:
	./bin/gc
