package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"bufio"
	"os"
	"encoding/csv"
	"bytes"
)

func main() {

	content, err := ioutil.ReadFile("./problems.csv")

	if err != nil {
		log.Fatal(err)
	}

	contentReader := bytes.NewReader(content)
	csvReader := csv.NewReader(contentReader)
	lines, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	var rightAnswers int

	for i := range lines {
		line := lines[i]
		fmt.Println(line[0])
		
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if string(line[1]) == strings.Trim(string(text), "\n") {
			rightAnswers++
		}
	}

	fmt.Printf("Você acertou o total de %d perguntas! \n", rightAnswers)
}
