package client

// #cgo CFLAGS: -I../bft/gmp -I../bft/libbyz -I../bft/sfs/include/sfslite -O3 -fno-exceptions -DNDEBUG
// #cgo LDFLAGS: -L../bft/gmp -L../bft/libbyz -L../bft/sfs/lib/sfslite -lbyz -lsfscrypt -lasync -lgmp -lstdc++
// #include<stdio.h>
// #include<stdlib.h>
// #include<string.h>
// #include<signal.h>
// #include<unistd.h>
// #include<sys/param.h>
// #include"libbyz.h"
import "C"
import (
	"log"
	"unsafe"
)

const simpleSize int = 4096

var option = 0

// ByzInitClient : Init client
func ByzInitClient(configPath string, configPrivPath string) {
	// var configPath = "../bft/config"
	var port = 0
	config := C.CString(configPath)
	configPriv := C.CString(configPrivPath)
	defer C.free(unsafe.Pointer(config))
	defer C.free(unsafe.Pointer(configPriv))

	C.Byz_init_client(config, configPriv, C.short(port))
}

// ByzRunClient : Alloc req and wait for reply from replica
func ByzRunClient() {
	var	readOnly = 0
	req := C.struct__Byz_buffer{}
	rep := C.struct__Byz_buffer{}
	C.Byz_alloc_request(&req, C.int(simpleSize))
	for i := 0; i < simpleSize; i++ {
		*(*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(req.contents)) + uintptr(i))) = C.char(option)
	}
	if option != 2 {
		req.size = 8
	} else {
		req.size = C.int(simpleSize)
	}

	// invoke request
	C.Byz_invoke(&req, &rep, C.bool(readOnly))

	// check reply
	if !(((option == 2 || option == 0) && rep.size == 8) || (option == 1 && rep.size == C.int(simpleSize))) {
		log.Fatal("invalid reply")
	}

	// free reply
	C.Byz_free_reply(&rep)

	C.Byz_free_request(&req)
}
