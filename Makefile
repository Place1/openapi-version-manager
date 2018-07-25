BINARY_NAME := openapi-version-manager
PLATFORMS := linux/amd64 darwin/amd64 windows/amd64
GO := vgo
temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

all: test build

build: $(PLATFORMS)

$(PLATFORMS):
				GOOS=$(os) GOARCH=$(arch) $(GO) build -o '$(BINARY_NAME)-$(os)-$(arch)'

test:
				$(GO) test -v ./...

clean:
				$(GO) clean

