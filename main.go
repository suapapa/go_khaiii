package main

import (
	"fmt"
	"log"
	"os"
)

var (
	exampleStr = "사랑은 모든것을 덮어주고 모든것을 믿으며 모든것을 바라고 모든것을  견디어냅니다"
)

func main() {
	fmt.Println(Version())
	k, err := New("", "")
	chk(err)
	defer k.Close()

	var inputStr string
	if len(os.Args) == 0 {
		inputStr = exampleStr
	} else {
		inputStr = os.Args[1]
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
