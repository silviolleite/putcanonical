package main

import (
	"flag"
	"fmt"
	"github.com/silviolleite/putcanonical/pkg/canonical"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	var f, i, c, t string

	flag.StringVar(&f, "file", "", "Path to data file (Required)")
	flag.StringVar(&i, "id", "", "Meli ID eg.: MLB1728494308")
	flag.StringVar(&c, "sku", "", "Canonical SKU eg.: PRDYMUPKZRHCFGW9")
	flag.StringVar(&t, "token", "", "Access Token")

	flag.Parse()

	var ls []string

	if f != "" {
		d, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Printf("File path not found. \n")
			os.Exit(1)
		}
		ls = strings.Split(string(d), "\n")

	} else {
		if i != "" && c != "" {
			ls = []string{fmt.Sprintf("%s,%s", c, i)}
		}
	}

	if len(ls) == 0 {
		fmt.Printf("Error: The data file path or the Meli ID and Canonical SKU is required! \n")
		fmt.Printf("You must use the flag -f <file_path> or the -i and -c to run it :) \n\n")
		fmt.Printf("Tip: Try use the flag -h to see all options\n\n")
		os.Exit(1)
	}

	if t == "" {
		fmt.Printf("Access Token is required! \n\n")
		os.Exit(1)
	}

	qs := canonical.New(&http.Client{}, ls, t)

	qs.Run()
}
