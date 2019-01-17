package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	store := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {

		store[input.Text()]++

	}

	for k, v := range store {

		if v > 1 {

			fmt.Printf("d%\ts%\n", v, k)

		}

	}

}
