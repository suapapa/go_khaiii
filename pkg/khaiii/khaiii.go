package khaiii

import (
	"github.com/suapapa/go_khaiii/internal/c_khaiii"
	"github.com/suapapa/go_khaiii/pkg/khaiiitype"
)

type Khaiii struct {
	Version string
	cKhaiii *c_khaiii.Khaiii
}

type KhaiiiOptions c_khaiii.Options

func New(options *KhaiiiOptions) (*Khaiii, error) {
	var opt *c_khaiii.Options
	if options != nil {
		opt = &c_khaiii.Options{
			Preanal:      options.Preanal,
			Errpatch:     options.Errpatch,
			Restore:      options.Restore,
			ResourcePath: options.ResourcePath,
		}
	}

	cKhaiii, err := c_khaiii.NewWithOptions(opt)
	if err != nil {
		return nil, err
	}

	ver := c_khaiii.Version()
	return &Khaiii{Version: ver, cKhaiii: cKhaiii}, nil
}

func (k *Khaiii) Close() {
	if k.cKhaiii != nil {
		k.cKhaiii.Close()
	}
}

func (k *Khaiii) Analyze(input, opt string) *khaiiitype.AnalyzeResult {
	cWordCh := k.cKhaiii.AnalyzeCh(input, opt)
	if cWordCh == nil {
		return nil
	}
	defer k.cKhaiii.FreeAnalyzeResult()

	result := &khaiiitype.AnalyzeResult{
		OrigText: input,
	}

	for w := range cWordCh {
		// w := &c_khaiii.Word{OrigStr: input, CWord: cWord}
		wc := &khaiiitype.WordChunk{
			Word:  w.Val(),
			Begin: w.Begin(),
			Len:   w.Length(),
		}

		for m := range w.CMorphs() {
			wc.Morphs = append(wc.Morphs, khaiiitype.Morph{
				Lex: m.Lex(),
				Tag: m.Tag(),
			})
		}

		result.WordChunks = append(result.WordChunks, wc)
	}

	return result
}
