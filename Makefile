#現在的時間點以utc+？的方式 local端的時間台灣就是+8
BUILD_DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

DOCKER_ACCOUNT := hazelnutkiller

DOCKER_IMAGE := $(SERVICE)


ifneq ($(DRONE_TAG),)
VERSION ?=(subst v,,$(DRONE_TAG))
#如果在local端就用 git describe --tags 的方式去取資訊
else
VERSION ?= $(shell git describe --tags --always | sed 's/-/+/' | sed 's/^v//')

LDFLAGS ?= -X github.com/hazelnutkiller/dockertest/version.Version=$(VERSION) -X github.com/hazelnutkiller/dockertest/version.BuildDate=$(BUILD_DATE)

all: build