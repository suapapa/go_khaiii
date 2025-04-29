package main

import (
	"encoding/csv"
	"html/template"
	"io"
	"os"
)

const (
	posStruct = `	// {{.Pos}} is {{.Desc}}
	{{.Pos}} = POS{
		Pos:         "{{.Pos}}",
		Category:    "{{.Cat}}",
		Description: "{{.Desc}}",
	}
`
	posMap = `	"{{.Pos}}": {{.Pos}},
`
)

type Line struct {
	Cat, Pos, Desc string
}

func main() {
	f, err := os.Open("pos.csv")
	chk(err)
	r := csv.NewReader(f)
	r.Comma = ';'
	r.Read() // skip title line
	t := template.Must(template.New("posStruct").Parse(posStruct))
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		chk(err)
		l := Line{
			Cat:  line[0],
			Pos:  line[1],
			Desc: line[2],
		}
		// log.Printf("%+v", l)
		err = t.Execute(os.Stdout, l)
		chk(err)
	}
	f.Close()

	f, err = os.Open("pos.csv")
	chk(err)
	r = csv.NewReader(f)
	r.Comma = ';'
	r.Read() // skip title line
	t = template.Must(template.New("posMap").Parse(posMap))
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		chk(err)
		l := Line{
			Cat:  line[0],
			Pos:  line[1],
			Desc: line[2],
		}
		// log.Printf("%+v", l)
		err = t.Execute(os.Stdout, l)
		chk(err)
	}
	f.Close()

}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
