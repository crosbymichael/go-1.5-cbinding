package main

import "C"
import "fmt"

//export PrintName
func PrintName(c *C.char) {
	name := C.GoString(c)
	fmt.Println(name)
}

func main() {

}
