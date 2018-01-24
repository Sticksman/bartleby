PACKAGE = github.com/Sticksman/bartleby
GOEXE ?= go

dep:
	## No dependency management for now.
	## One day
	${GOEXE} get -u ./...

build:
	${GOEXE} build ${PACKAGE}

all: dep install

install:
	${GOEXE} install ${PACKAGE}

