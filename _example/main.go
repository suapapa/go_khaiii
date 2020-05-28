package main

import (
	"fmt"
	"log"
	"os"

	"github.com/suapapa/go/khaiii"
)

var (
	exampleStr = "사랑은 모든것을 덮어주고 모든것을 믿으며 모든것을 바라고 모든것을 견디어냅니다"
)

func main() {
	fmt.Println(khaiii.Version())
	k, err := khaiii.New()
	chk(err)
	defer k.Close()

	var inputStr string
	if len(os.Args) > 1 {
		inputStr = os.Args[1]
	} else {
		inputStr = exampleStr
	}

	for v := range k.Analyze(inputStr, "") {
		log.Println(v.Val())
		for m := range v.Morphs() {
			// log.Println(m.Lex(), m.Tag(), m.cptr.begin, m.cptr.length)
			log.Println("   ", m.Lex(), m.Tag())
		}
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
