// multiple functions
// _ is blank identifier
// int, int

package main
import "fmt"
func main() {
    _, a:= multi()
    fmt.Println(a)
}
func multi() (int, int) {
    return 10,5
}
