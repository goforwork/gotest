package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	for _, v := range os.Args[1:] {

		if !strings.HasPrefix(v, "http://") {

			v = "http://" + v

		}

		res, err := http.Get(v)

		if err != nil {

			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)

		}

		defer res.Body.Close()

		io.Copy(os.Stdout, res.Body)

		if err != nil {

			fmt.Fprintf(os.Stderr, "reading: %v\n", err)

		}

		fmt.Printf("%s", res.Status)

	}

}
