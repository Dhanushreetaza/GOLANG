//swap existing values

package main
import "fmt"
func main() {
    subject1, subject2 := "Math" , " English"
    subject1, subject2 = "Science" , "Math"
    fmt.Println(subject1, subject2 )
}
