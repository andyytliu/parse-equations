package parser

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

// Empty string is always mapped to 0 in var_map
func ReadVariables(file_name string, var_map map[string]string) {

	file, err := os.Open(file_name)
	if err != nil {
		log.Println(">>>>>>>>>>> error opening file: " + err.Error())
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var v []rune
	var vs []string

	for {
		var (
			err error
			r rune
		)

		r, _, err = reader.ReadRune()
		if err != nil && err != io.EOF {
			log.Println(">>>>>>>>>>> error: " + err.Error())
			break
		}
		if err == io.EOF {
			break
		}

		if unicode.IsSpace(r) || r == 92 /* backslash */ {
			continue
		} else if r == 44 /* comma */ {
			vs = append(vs, string(v))
			v = []rune{}
		} else {
			v = append(v, r)
		}
	}
	vs = append(vs, string(v))

	for i, s := range vs {
		var_map[s] = fmt.Sprint(i + 1)
	}
	var_map[""] = "0"
}
