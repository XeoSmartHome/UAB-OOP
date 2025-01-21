package main

/*
#include <stdio.h>
#include <stdlib.h>  // For free function

// C function that prints a message
void helloFromC() {
    printf("Hello from C!\n");
}

int add(int a, int b) {
	return a + b;
}

void uppercase(char *str) {
	while (*str) {
		if (*str >= 'a' && *str <= 'z'){
			*str = *str - 32;
		}
		str++;
	}
}
*/
import "C"
import (
	"unsafe"
)

func main() {
	// Call the C function
	C.helloFromC()

	s := "hello, world"

	// Convert Go string to C string
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	// Call the C function with the C string
	C.uppercase(cs)

	// Convert C string to Go string
	result := C.GoString(cs)

	// Print the result
	println(result)
}
