// Copyright 2020 Homin Lee <ff4500@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/suapapa/go_khaiii/internal/c_khaiii"
	"github.com/suapapa/go_khaiii/pkg/krpos"
)

var (
	exampleStr = "사랑은 모든것을 덮어주고 모든것을 믿으며 모든것을 바라고 모든것을 견디어냅니다"
)

func main() {
	fmt.Println(c_khaiii.Version())
	k, err := c_khaiii.New()
	chk(err)
	defer k.Close()

	var inputStr string
	if len(os.Args) > 1 {
		inputStr = os.Args[1]
	} else {
		inputStr = exampleStr
	}

	for v := range k.AnalyzeCh(inputStr, "") {
		log.Println(v.Val())
		for m := range v.CMorphs() {
			// log.Println(m.Lex(), m.Tag(), m.cptr.begin, m.cptr.length)
			log.Printf("   %s %s",
				m.Lex(),               // 형태소
				krpos.POSMap[m.Tag()], // 품사
			)
		}
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
