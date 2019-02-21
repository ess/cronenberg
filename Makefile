# This how we want to name the binary output
BINARY=cronenberg

# These are the values we want to pass for VERSION and BUILD
# git tag 1.0.1
# git commit -am "One more change after the tags"
VERSION=`./scripts/genver`
BUILD=`date +%FT%T%z`
PACKAGE="github.com/ess/cronenberg/cmd/cronenberg"
TARGET="builds/${BINARY}-${VERSION}"
PREFIX="${TARGET}/${BINARY}-${VERSION}"
TESTFILES=`go list ./... | grep -v /vendor/`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s \
				-X ${PACKAGE}/cmd.Version=${VERSION} \
				-X ${PACKAGE}/cmd.Build=${BUILD} \
				-extldflags '-static'"

# Build for the current platform
all: clean build

# Build a new release
release: distclean distbuild linux

# Builds the project
build:
	go build ${LDFLAGS} -o ${BINARY} ${PACKAGE}

# Builds the project for all possible platforms
distbuild:
	mkdir -p ${TARGET}

# Installs our project: copies binaries
install:
	go install ${LDFLAGS}

# Cleans our project: deletes binaries
clean:
	rm -rf ${BINARY}

# Cleans release files
distclean:
	rm -rf ${TARGET}

test:
	./scripts/blanket

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build ${LDFLAGS} -o ${TARGET}/${BINARY}-${VERSION}-linux-386 ${PACKAGE}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${TARGET}/${BINARY}-${VERSION}-linux-amd64 ${PACKAGE}

darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build ${LDFLAGS} -o ${TARGET}/${BINARY}-${VERSION}-darwin-386 ${PACKAGE}
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${TARGET}/${BINARY}-${VERSION}-darwin-amd64 ${PACKAGE}
	
windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build ${LDFLAGS} -o ${TARGET}/${BINARY}-${VERSION}-windows-386.exe ${PACKAGE}
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${TARGET}/${BINARY}-${VERSION}-windows-amd64.exe ${PACKAGE}

freebsd:
	CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build ${LDFLAGS} -o ${TARGET}/${BINARY}-${VERSION}-freebsd-386 ${PACKAGE}
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build ${LDFLAGS} -o ${TARGET}/${BINARY}-${VERSION}-freebsd-amd64 ${PACKAGE}
