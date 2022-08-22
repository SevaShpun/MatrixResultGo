package main

import (
	"fmt"
	"strconv"
	"os"
	"os/exec"
	"strings"
	"bufio"
)

var cell = make([]float64, 9)

func template_text() {
	fmt.Println("\t\t\t\t=================")
	fmt.Println("\t\t\t\t   Template 3x3")
	fmt.Println("\t\t\t\t=================")
	fmt.Println("\t\t\t\t    1 | 2 | 3")
	fmt.Println("\t\t\t\t    4 | 5 | 6")
	fmt.Println("\t\t\t\t    7 | 8 | 9")
	fmt.Println("\t\t\t\t=================")
}

func cleaner() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func matrix_filling_text() {
	fmt.Println("\t\t\t\tFill in the cells")
	fmt.Println("\t\t\t\t=================")
	fmt.Printf("\t\t\t\t%v\t%v\t%v\n", cell[0], cell[1], cell[2])
	fmt.Printf("\t\t\t\t%v\t%v\t%v\n", cell[3], cell[4], cell[5])
	fmt.Printf("\t\t\t\t%v\t%v\t%v\n", cell[6], cell[7], cell[8])
	fmt.Println("\t\t\t\t=================")
	fmt.Println("")
}


func main() {
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(cell); i++ {
		template_text();
		matrix_filling_text();

		fmt.Printf("cell %d > ", i+1)
		send, _ := reader.ReadString('\n')

		n, _ := strconv.ParseFloat(strings.TrimSpace(send), 64)

		cell[i] = n
		cleaner()
	}
	template_text()
	matrix_filling_text();

	r := cell[0]*(cell[4]*cell[8]-cell[5]*cell[7])-cell[1]*(cell[3]*cell[8]-cell[5]*cell[6])+cell[2]*(cell[3]*cell[7]-cell[4]*cell[6]);
	
	fmt.Printf("\t\t\t\tResult: %v\n", r)
	fmt.Print("\t\t\t\tPress any key for exit")

	reader.ReadString('\n')
}
