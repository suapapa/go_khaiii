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
	ch := k.Analyze("세상 안녕", "")
	for v := range ch {
		log.Printf("%+v", v)
		log.Println(v.Val())
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
