package main

import ("fmt"
        "bufio"
        "log"
        "os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

	for scanner.scan(){
		checkDomain(scanner.Text())
		}
	if err := scanner.Err(); err != nil {
		log.Fatal ( "Error: could not read from input : %v\n ", err) 
		}
}
