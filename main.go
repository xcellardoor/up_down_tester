//Main package
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
)

//Add a function that reads the configuration file
/*Currently we put the configuration file next to the executable
but in the final thing it will live in /etc/{program_name}/config*/

//Add a function that takes an array of

func main() {
	fmt.Println("Hello, this compiles")
	resp, err := http.Get("http://samcater.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	color.Blue("Output %s", body)
}
