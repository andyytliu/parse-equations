package parser

import (
	"bufio"
	"io"
	"log"
	"unicode"
)

func ParseEquations(reader *bufio.Reader,
	writer *bufio.Writer, var_map map[string]string) {

	var (
		err error
		r rune
		v string = ""
		coef string = ""
		readingV bool = false
	)
	
	// Throw away everything before '{'
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
				writer.WriteString(var_map[v] + " ")
				if coef == "" {
					coef = "0"
				}
				writer.WriteString(coef + " ")
				writer.WriteString("\n")
				writer.Flush()
				break outer

			case r == 44 /* comma */ :
				writer.WriteString(var_map[v] + " ")
				if coef == "" {
					coef = "0"
				}
				writer.WriteString(coef + " ")
				v = ""
				coef = ""
				readingV = false
				writer.WriteString("\n")
				writer.Flush()

			case unicode.IsSpace(r) || r == 92 /* backslash */ :
				continue

			case r == 43 /* + */ || r == 45 /* - */ :
				writer.WriteString(var_map[v] + " ")
				if coef == "" {
					coef = "0"
				}
				writer.WriteString(coef + " ")
				v = ""
				coef = string(r)
				readingV = false

			case r == 42 /* * */ :
				readingV = true

			default:
				if readingV {
					v += string(r)
				} else {
					if unicode.IsLetter(r) {
						readingV = true
						coef += "1"
						v += string(r)
					} else {
						coef += string(r)
					}
				}

			}

			

		}
	}

}

