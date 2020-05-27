package main

import "fmt"

func main() {
	fmt.Println(Version())
	k, err := NewKhaiii("", "")
	chk(err)
	defer k.Close()
	fmt.Printf("%+v", k)
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
