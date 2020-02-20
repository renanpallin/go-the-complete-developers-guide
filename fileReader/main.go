package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo", err)
	}
	io.Copy(os.Stdout, f)

	// data, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	fmt.Println("Erro ao abrir arquivo", err)
	// }
	// // fmt.Println(string(data))

}
