PACKAGE = github.com/Sticksman/longform-text-compiler
GOEXE ?= go

dep:
	## No dependency management for now.
	## One day
	${GOEXE} get -u ./...

build: dep
	${GOEXE} build ${PACKAGE}

install: dep
	${GOEXE} install ${PACKAGE}

