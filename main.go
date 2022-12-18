package main

import (
	. "github.com/andyytliu/parse-equations/parser"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	in_file string
	vars_file string
	out_file string
)

func main() {
	var_map := make(map[string]string)

	flag.StringVar(&in_file, "input", "input.txt", "Input file")
	flag.StringVar(&out_file, "output", "output.txt", "Output file")
	flag.StringVar(&vars_file, "vars", "vars.txt", "Variables file")
	flag.Parse()

	fmt.Println("Welcome!")

	ReadVariables(vars_file, var_map)
	eqs := ParseEquations(in_file, var_map)


	file, err := os.Create(out_file)
	if err != nil {
		log.Println(">>>>>>>>>>> error: " + err.Error())
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, eq := range eqs {
		for _, term := range eq {
			_, err := writer.WriteString(term + " ")
			if err != nil {
				log.Println(">>>>>>>>>>> error: " + err.Error())
			}
		}
		_, err = writer.WriteString("\n")
		if err != nil {
			log.Println(">>>>>>>>>>> error: " + err.Error())
		}
		writer.Flush()
	}

}
