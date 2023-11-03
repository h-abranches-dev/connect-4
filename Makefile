SYSTEM ?= "undefined"

GE_BUILD_VERSION = "ge-v0.0.1"
GS_BUILD_VERSION = "gs-v0.0.1"
GC_BUILD_VERSION = "gc-v0.0.1"

ifeq ($(SYSTEM), gc)
	BUILD_VERSION := $(GC_BUILD_VERSION)
endif

ifeq ($(SYSTEM), gs)
	BUILD_VERSION := $(GS_BUILD_VERSION)
endif

ifeq ($(SYSTEM), ge)
	BUILD_VERSION := $(GE_BUILD_VERSION)
endif

LDFLAGS = "-X 'github.com/version-go/ldflags.buildVersion=$(BUILD_VERSION)'"

build:
	go build -ldflags $(LDFLAGS) -o bin/$(SYSTEM) ./entrypoints/$(SYSTEM)

tests:
	# run tests in current directory and all of its subdirectories
	go test -v ./...

run_gc:
	./bin/gc

run_gs:
	./bin/gs --port=50052

run_ge:
	./bin/ge --port=50051
