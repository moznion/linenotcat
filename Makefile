GOVERSION=$(shell go version)
GOOS=$(word 1,$(subst /, ,$(lastword $(GOVERSION))))
GOARCH=$(word 2,$(subst /, ,$(lastword $(GOVERSION))))
RELEASE_DIR=bin
DEVTOOL_DIR=devtools
PACKAGE=github.com/moznion/linenotcat
REVISION=$(shell git rev-parse --verify HEAD)
HAVE_GLIDE:=$(shell which glide > /dev/null 2>&1)

.PHONY: clean build build-linux-amd64 build-linux-386 build-darwin-amd64 build-darwin-386 $(RELEASE_DIR)/linenotcat_$(GOOS)_$(GOARCH) all

all: installdeps build-linux-amd64 build-linux-386 build-darwin-amd64 build-darwin-386

build: $(RELEASE_DIR)/linenotcat_$(GOOS)_$(GOARCH)

build-linux-amd64:
	@$(MAKE) build GOOS=linux GOARCH=amd64

build-linux-386:
	@$(MAKE) build GOOS=linux GOARCH=386

build-darwin-amd64:
	@$(MAKE) build GOOS=darwin GOARCH=amd64

build-darwin-386:
	@$(MAKE) build GOOS=darwin GOARCH=386

$(RELEASE_DIR)/linenotcat_$(GOOS)_$(GOARCH):
ifndef VERSION
	@echo '[ERROR] $$VERSION must be specified'
	exit 255
endif
	go build -ldflags "-X $(PACKAGE).rev=$(REVISION) -X $(PACKAGE).ver=$(VERSION)" \
		-o $(RELEASE_DIR)/linenotcat_$(GOOS)_$(GOARCH)_$(VERSION) cmd/linenotcat/linenotcat.go

$(DEVTOOL_DIR)/$(GOOS)/$(GOARCH)/glide:
ifndef HAVE_GLIDE
	@echo "Installing glide for $(GOOS)/$(GOARCH)..."
	mkdir -p $(DEVTOOL_DIR)/$(GOOS)/$(GOARCH)
	wget -q -O - https://github.com/Masterminds/glide/releases/download/v0.12.3/glide-v0.12.3-$(GOOS)-$(GOARCH).tar.gz | tar xvz
	mv $(GOOS)-$(GOARCH)/glide $(DEVTOOL_DIR)/$(GOOS)/$(GOARCH)/glide
	rm -rf $(GOOS)-$(GOARCH)
endif

glide: $(DEVTOOL_DIR)/$(GOOS)/$(GOARCH)/glide

installdeps: glide
	@PATH=$(DEVTOOL_DIR)/$(GOOS)/$(GOARCH):$(PATH) glide install

clean:
	rm -rf $(RELEASE_DIR)/linenotcat_*

