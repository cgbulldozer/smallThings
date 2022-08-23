/*
package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {

    content, err := ioutil.ReadFile("test.hip")

     if err != nil {
          log.Fatal(err)
     }

    fmt.Println(string(content))
}

*/





package main
import (
    "bufio"
    "fmt"
    "log"
    "os"
    _"io"
)
func main() {

    filename := os.Args[1]
    fmt.Println(filename)
    f, err := os.Open(filename)

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {

        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}


/*

func main() {
    filename := os.Args[1]
    fmt.Println("Open file: ",filename)

    f, err := os.Open(filename)
      if err != nil {
        fmt.Println("error opening file ", err)
        os.Exit(1)
      }
      defer f.Close()
      r := bufio.NewReader(f)
      for {
        path, err := r.ReadString('\n') // 0x0A separator = newline
        if err == io.EOF {
          // do something here
           fmt.Println("Path:", path)

        } else if err != nil {
           fmt.Println("err:", path)
           // return err // if you return error
        }
      }


}

*/

/*
// It work for reading file:
import (
    "os"
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    filename := os.Args[1]
    fmt.Println("Open file: ",filename)

    bytesRead, _ := ioutil.ReadFile(filename)
    file_content := string(bytesRead)
    lines := strings.Split(file_content, "\n")
    fmt.Println(lines)

}
*/

/*
package main
import (
    "fmt"
    "io/ioutil"
     "log"
     "os"
)

func main() {
    filename := os.Args[1]
    fmt.Println("Open file: ",filename)

    files, err := ioutil.ReadDir(filename)
    if err != nil {
        log.Fatal(err)
    }
 
    for _, f := range files {
            fmt.Println(f.Name())
    }
}

*/

