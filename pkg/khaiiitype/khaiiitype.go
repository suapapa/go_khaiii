package khaiiitype

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
