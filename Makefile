include $(GOROOT)/src/Make.inc

TARG=main
GOFILES=\
	main.go
DEPS=dice

include $(GOROOT)/src/Make.cmd
