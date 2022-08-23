package main
 
import (
	_"archive/zip"
	"fmt"
	"log"
	"os"
	"io/ioutil"
)
 
 
func main() {
    filename := os.Args[1]
    fmt.Println("Open file: ",filename)
	content, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	byteContainer, err := ioutil.ReadAll(content) // you may want to handle the error
	fmt.Printf("size:%d", len(byteContainer))
	fmt.Printf("size:%d", byteContainer )
	fmt.Printf("\n")
}

