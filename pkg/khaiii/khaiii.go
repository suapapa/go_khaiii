// Copyright 2020 Homin Lee <ff4500@gmail.com>. All rights reserved.
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

// CMorph represents morpheme data struct
type CMorph struct {
	cptr *C.khaiii_morph_t
}

// Lex returns word's lexical
func (m *CMorph) Lex() string {
	return C.GoString(m.cptr.lex)
}

// Tag returns word's part-of-speech tag
func (m *CMorph) Tag() string {
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
func (w *Word) CMorphs() chan *CMorph {
	retCh := make(chan *CMorph)
	m := w.cptr.morphs
	go func() {
		for m != nil {
			retCh <- &CMorph{cptr: m}
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
	k.FreeAnalyzeResult()

	C.khaiii_close(k.handle)
}

// AnalyzeCh analyzes given input and return Word stream
func (k *Khaiii) AnalyzeCh(input, opt string) chan *Word {
	k.FreeAnalyzeResult()

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
	if k.firstWord == nil {
		return
	}

	C.khaiii_free_results(k.handle, k.firstWord)
	k.firstWord = nil
}

func (k *Khaiii) Analyze(input, opt string) *AnalyzeResult {
	k.FreeAnalyzeResult()

	cWord := C.khaiii_analyze(k.handle, C.CString(input), C.CString(opt))
	k.firstWord = cWord

	result := &AnalyzeResult{
		OrigText: input,
	}

	for cWord != nil {
		w := &Word{origStr: input, cptr: cWord}
		wc := &WordChunk{
			Word:  w.Val(),
			Begin: int(w.cptr.begin),
			Len:   int(w.cptr.length),
		}

		for m := range w.CMorphs() {
			wc.Morphs = append(wc.Morphs, Morph{
				Lex: m.Lex(),
				Tag: m.Tag(),
			})
		}

		result.WordChunks = append(result.WordChunks, wc)

		cWord = cWord.next
	}

	k.FreeAnalyzeResult()
	return result
}

type AnalyzeResult struct {
	OrigText   string       `json:"orig_text" yaml:"orig_text"`
	WordChunks []*WordChunk `json:"word_chunks" yaml:"word_chunks"`
}

type WordChunk struct {
	Word   string  `json:"word" yaml:"word"`
	Begin  int     `json:"begin" yaml:"begin"`
	Len    int     `json:"len" yaml:"len"`
	Morphs []Morph `json:"morphs" yaml:"morphs"`
}

type Morph struct {
	Lex string `json:"lex" yaml:"lex"`
	Tag string `json:"tag" yaml:"tag"`
}
