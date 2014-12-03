package rand
// http://golang.org/cmd/cgo/

// #include <stdio.h>
// #include <errno.h>
//
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "fmt"

func PrintHello() {
  s := "Hello Cgo"
	cs := C.CString(s)
	// C.print(cs)
	// C.free(unsafe.Pointer(cs))

	fmt.Println(cs)
	C.puts(C.CString("Hello, world\n"))
}