//Main package
package main

import (
	"fmt" //Text output, format and control
	//"io/ioutil" //Used for some lower level stuff
	"log" //Standard Go log function
	//"net/http" //Standard Go module for making HTTP GET requests, if we get a reply with content #winning

	"github.com/fatih/color" //Custom module used for Terminal Coloration
  "os"
  "net"
  "time"

  "github.com/tatsushid/go-fastping" //Ping module from Tatsushid on Github, used for Ping functionality
)

func pingTest(){

  p := fastping.NewPinger()
  ra, err := net.ResolveIPAddr("ip4:icmp", os.Args[1])
  if err != nil {log.Fatal(err)}

  p.AddIPAddr(ra)
  p.OnRecv = func(addr *net.IPAddr, rtt time.Duration){
    fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
  }
  p.OnIdle = func(){
    fmt.Println(color.GreenString("Finish"))
  }
  err = p.Run()
  if err != nil{log.Fatal(err)}
}


//Add a function that reads the configuration file
/*Currently we put the configuration file next to the executable
but in the final thing it will live in /etc/{program_name}/config*/

//Add a function that takes an array of

func main() {

	fmt.Println("Hello, this compiles")
	/*resp, err := http.Get("http://samcater.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	color.Blue("Output %s", body)


  fmt.Println("All systems are:", color.GreenString("-GO-"))
  fmt.Printf("... here is another way of doing it %s\n", color.BlueString("BOO"))
  */
  pingTest()
  }
