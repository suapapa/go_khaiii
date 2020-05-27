package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(Version())
	k, err := New("", "")
	chk(err)
	defer k.Close()
	ch := k.Analyze("세상아 안녕", "")
	for v := range ch {
		log.Println(v.Val())
		for m := range v.Morphs() {
			log.Println(m.Lex(), m.Tag(), m.cptr.begin, m.cptr.length)
		}
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
