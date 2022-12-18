package parser

import (
	"bufio"
	"io"
	"log"
	"os"
	"unicode"
)

func ParseEquations(file_name string,
	var_map map[string]string) [][]string {

	var (
		r rune
		eqs [][]string
		eq []string
		v string = ""
		coef string = ""
		readingV bool = false
	)

	file, err := os.Open(file_name)
	if err != nil {
		log.Println(">>>>>>>>>>> error opening file: " + err.Error())
		return eqs
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	

	for {
		r, _, err = reader.ReadRune()
		if err != nil && err != io.EOF {
			log.Println(">>>>>>>>>>> error: " + err.Error())
			break
		}
		if err == io.EOF {
			break
		}

		if r == 123 /* { */ {
			break
		}
	}

	if r == 123 {
		outer:
		for {

			r, _, err = reader.ReadRune()
			if err != nil && err != io.EOF {
				log.Println(">>>>>>>>>>> error: " + err.Error())
				break
			}
			if err == io.EOF {
				log.Println(">>>>>>>>>>> error: unexpected EOF")
				break
			}

			switch {
			case r == 125 /* } */ :
				eq = append(eq, var_map[v])
				if coef == "" {
					coef = "0"
				}
				eq = append(eq, coef)
				eqs = append(eqs, eq)
				break outer

			case r == 44 /* comma */ :
				eq = append(eq, var_map[v])
				if coef == "" {
					coef = "0"
				}
				eq = append(eq, coef)
				v = ""
				coef = ""
				readingV = false
				eqs = append(eqs, eq)
				eq = []string{}

			case unicode.IsSpace(r) || r == 92 /* backslash */ :
				continue

			case r == 43 /* + */ || r == 45 /* - */ :
				if readingV {
					eq = append(eq, var_map[v])
					if coef == "" {
						coef = "0"
					}
					eq = append(eq, coef)
					v = ""
					coef = ""
					readingV = false
				}
				v = ""
				coef = string(r)

			case r == 42 /* * */ :
				readingV = true

			default:
				if readingV {
					v += string(r)
				} else {
					if unicode.IsLetter(r) {
						readingV = true
						coef = "1"
						v += string(r)
					} else {
						coef += string(r)
					}
				}
				
			}

			

		}
	}

	return eqs
}

