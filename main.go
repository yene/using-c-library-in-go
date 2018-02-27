package main

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR} -lhiya
#include "hiya.h"
*/
import "C"

type Message C.Message

func CreateMessage(msg string) *C.Message {
	cMsg := C.CString(msg)
	return C.create_message(cMsg)
}

func DisplayMessage(msg *C.Message) {
	C.display_message(msg)
}

func FreeMessage(msg *C.Message) {
	C.free_message(msg)
}

func main() {
	msg := CreateMessage("Hello, world!")
	DisplayMessage(msg)
	FreeMessage(msg)
}
