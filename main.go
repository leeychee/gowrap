package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	flag "github.com/spf13/pflag"
)

const currentVersion = "v0.1.0"

var (
	input   string
	output  string
	pkgname string
	varname string
	version bool
)

func init() {
	flag.StringVarP(&input, "input", "i", "", "Input file name, default is stdin.")
	flag.StringVarP(&output, "output", "o", "", "Output file name, default is stdout.")
	flag.StringVarP(&pkgname, "pkgname", "p", "", "package name, if not empty, output is in go's format.")
	flag.StringVarP(&varname, "varname", "", "varname", "var name")
	flag.BoolVarP(&version, "version", "v", false, "print version.")
	flag.Usage = func() {
		fmt.Println(`Wrap string in go style with backtick(` + "`)")
		fmt.Printf("Usage of %s\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if version {
		fmt.Println(currentVersion)
		return
	}
	var ifs io.Reader
	var ofs io.Writer
	var err error
	if input == "" {
		ifs = os.Stdin
	} else {
		if ifs, err = os.Open(input); err != nil {
			log.Println(err)
		}
	}
	if output == "" {
		ofs = os.Stdout
	} else {
		// Must Trunc when the origin file is longer than the new one.
		if ofs, err = os.OpenFile(output, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666); err != nil {
			log.Println(err)
		}
	}
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, ifs); err != nil {
		log.Println(err)
	}
	bt := backtick(buf.String())
	if pkgname != "" {
		fmt.Fprintf(ofs, "package %s\n\nvar %s = ", pkgname, varname)
	}
	fmt.Fprint(ofs, bt)
}

// func openOrDefault(f string, )

func backtick(s string) string {
	buf := bytes.NewBufferString("")
	// var inbacktick = false
	var backticks = ""
	bs := []byte(s)
	n := len(bs)
	first := true
	inbacktick := false
	for i := 0; i < n; i++ {
		b := bs[i]
		if b == '`' {
			inbacktick = false
			backticks += "`"
			if i == n-1 || bs[i+1] != '`' {
				if !first {
					buf.WriteString("+\n")
				}
				buf.WriteString(`"` + backticks + `"`)
				backticks = ""
				first = false
			}
		} else {
			if !inbacktick {
				if !first {
					buf.WriteString("+\n")
				}
				buf.WriteByte('`')
				inbacktick = true
			}
			buf.WriteByte(b)
			if i == n-1 || bs[i+1] == '`' {
				buf.WriteByte('`')
			}
			first = false
		}
	}
	return buf.String()
}
