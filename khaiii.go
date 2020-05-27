package main

// #cgo LDFLAGS: -lkhaiii
// #include <khaiii/khaiii_api.h>
import "C"

import (
	"fmt"
	"log"
)

// Version returns version string
func Version() string {
	return C.GoString(C.khaiii_version())
}

// Khaiii represent khaiii engine
type Khaiii struct {
	handle C.int
}

// New opens khaiii resources
func New(rscDir, opt string) (*Khaiii, error) {
	if rscDir == "" {
		rscDir = "/usr/local/share/khaiii"
	}
	h := C.khaiii_open(C.CString(rscDir), C.CString(opt))
	if h < 0 {
		return nil, fmt.Errorf("fail to open khaiii")
	}
	return &Khaiii{handle: h}, nil
}

// Close closes khaiii resource
func (k *Khaiii) Close() {
	C.khaiii_close(k.handle)
}

func (k *Khaiii) Analyze(input, opt string) {
	words := C.khaiii_analyze(k.handle, C.CString(input), C.CString(opt))
	log.Printf("%+v", words)
	log.Println(input[words.begin:words.length])
}
