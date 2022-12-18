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
	in_file_name string
	vars_file_name string
	out_file_name string
)

func main() {
	

	flag.StringVar(&in_file_name, "input", "input.txt", "Input file")
	flag.StringVar(&out_file_name, "output", "output.txt", "Output file")
	flag.StringVar(&vars_file_name, "vars", "vars.txt", "Variables file")
	flag.Parse()

	fmt.Println("Welcome!")

	var_map := make(map[string]string)
	ReadVariables(vars_file_name, var_map)

	in_file, err := os.Open(in_file_name)
	if err != nil {
		log.Println(">>>>>>>>>>> error opening file: " + err.Error())
		return
	}
	defer in_file.Close()

	out_file, err := os.Create(out_file_name)
	if err != nil {
		log.Println(">>>>>>>>>>> error opening file: " + err.Error())
		return
	}
	defer out_file.Close()

	reader := bufio.NewReader(in_file)
	writer := bufio.NewWriter(out_file)

	ParseEquations(reader, writer, var_map)


	// file, err := os.Create(out_file_name)
	// if err != nil {
	// 	log.Println(">>>>>>>>>>> error: " + err.Error())
	// 	return
	// }
	// defer file.Close()

	// writer := bufio.NewWriter(file)

	// for _, eq := range eqs {
	// 	for _, term := range eq {
	// 		_, err := writer.WriteString(term + " ")
	// 		if err != nil {
	// 			log.Println(">>>>>>>>>>> error: " + err.Error())
	// 		}
	// 	}
	// 	_, err = writer.WriteString("\n")
	// 	if err != nil {
	// 		log.Println(">>>>>>>>>>> error: " + err.Error())
	// 	}
	// 	writer.Flush()
	// }

}
