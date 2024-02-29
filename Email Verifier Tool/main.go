package main

import ("fmt"
        "bufio"
		"log"
		"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")
}