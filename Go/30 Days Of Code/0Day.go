package main
import "fmt"
import "bufio"
import "os"

func main() {
 text, err := bufio.NewReader(os.Stdin).ReadString('\n')
 fmt.Println("Hello, World.")
 fmt.Println(text)
 if err != nil && err.Error() != "EOF" {
     fmt.Println(err)
 }
}
