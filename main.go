package main

import "fmt"

func main() {
	fmt.Println(Version())
	k, err := New("", "")
	chk(err)
	defer k.Close()
	k.Analyze("세상 안녕", "")
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
