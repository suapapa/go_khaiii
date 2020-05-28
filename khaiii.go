// Copyright 2020 Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package khaiii

// #cgo LDFLAGS: -lkhaiii
// #include <khaiii/khaiii_api.h>
import "C"

import (
	"encoding/json"
	"fmt"
)

// Version returns version string
func Version() string {
	return C.GoString(C.khaiii_version())
}

// Morph represents morpheme data struct
type Morph struct {
	cptr *C.khaiii_morph_t
}

// Lex returns word's lexical
func (m *Morph) Lex() string {
	return C.GoString(m.cptr.lex)
}

// Tag returns word's part-of-speech tag
func (m *Morph) Tag() string {
	return C.GoString(m.cptr.tag)
}

// Word represents word data structure
type Word struct {
	origStr string
	cptr    *C.khaiii_word_t
}

// Val returns current word value
func (w *Word) Val() string {
	return w.origStr[w.cptr.begin : w.cptr.begin+w.cptr.length]
}

// Morphs return Morph stream for the Word
func (w *Word) Morphs() chan *Morph {
	retCh := make(chan *Morph)
	m := w.cptr.morphs
	go func() {
		for m != nil {
			retCh <- &Morph{cptr: m}
			m = m.next
		}
		close(retCh)
	}()

	return retCh
}

// Options has options for make new Khaiii instance
type Options struct {
	Preanal      bool   `json:"preanal"`  // 기분석사전
	Errpatch     bool   `json:"errpatch"` // 오분석패치
	Restore      bool   `json:"restore"`  // 원형복원
	ResourcePath string `jsong:"-"`       // 리소스경로
}

var (
	// DefaultOptions store default options for new khaiii instance
	DefaultOptions = &Options{
		Preanal:      true,
		Errpatch:     true,
		Restore:      true,
		ResourcePath: "/usr/local/share/khaiii",
	}
)

// Khaiii represent khaiii engine
type Khaiii struct {
	handle    C.int
	firstWord *C.khaiii_word_t
}

// New opens khaiii with default options
func New() (*Khaiii, error) {
	return NewWithOptions(nil)
}

// NewWithOptions opens khaiii with given options
func NewWithOptions(opt *Options) (*Khaiii, error) {
	if opt == nil {
		opt = DefaultOptions
	}
	rscDir := opt.ResourcePath
	optJSON, err := json.Marshal(opt)
	if err != nil {
		return nil, fmt.Errorf("someing woring while paring options")
	}
	h := C.khaiii_open(C.CString(rscDir), C.CString(string(optJSON)))
	if h < 0 {
		return nil, fmt.Errorf("fail to open khaiii")
	}

	return &Khaiii{handle: h}, nil
}

// Close closes khaiii resource
func (k *Khaiii) Close() {
	if k.firstWord != nil {
		k.FreeAnalyzeResult()
	}
	C.khaiii_close(k.handle)
}

// Analyze analyzes given input and return Word stream
func (k *Khaiii) Analyze(input, opt string) chan *Word {
	if k.firstWord != nil {
		k.FreeAnalyzeResult()
	}
	cWord := C.khaiii_analyze(k.handle, C.CString(input), C.CString(opt))
	k.firstWord = cWord

	retCh := make(chan *Word)
	go func() {
		for cWord != nil {
			w := &Word{origStr: input, cptr: cWord}
			retCh <- w
			cWord = cWord.next
		}
		close(retCh)
	}()

	return retCh
}

// FreeAnalyzeResult free memories of analyzed results
func (k *Khaiii) FreeAnalyzeResult() {
	C.khaiii_free_results(k.handle, k.firstWord)
	k.firstWord = nil
}
