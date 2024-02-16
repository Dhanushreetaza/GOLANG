//split folder,file and discard the file
// import path

package main
import (
	"fmt"
	"path"
)
func main() {
	dir, _ := path.Split("folder/file.txt")
	fmt.Println(dir)
}


