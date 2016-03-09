// Numeric types:

// uint        either 32 or 64 bits
// int         same size as uint
// uintptr     an unsigned integer large enough to store the uninterpreted bits of
//             a pointer value
// uint8       the set of all unsigned  8-bit integers (0 to 255)
// uint16      the set of all unsigned 16-bit integers (0 to 65535)
// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

// int8        the set of all signed  8-bit integers (-128 to 127)
// int16       the set of all signed 16-bit integers (-32768 to 32767)
// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
// int64       the set of all signed 64-bit integers
//             (-9223372036854775808 to 9223372036854775807)

// float32     the set of all IEEE-754 32-bit floating-point numbers
// float64     the set of all IEEE-754 64-bit floating-point numbers

// complex64   the set of all complex numbers with float32 real and imaginary parts
// complex128  the set of all complex numbers with float64 real and imaginary parts

// byte        alias for uint8
// rune        alias for int32 (represents a Unicode code point)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


package main

import (
  "fmt"
  "math/cmplx"
)

func main() {

// declarations

  const f = "%T(%v)\n"
  const printline = "~~~~~~~~~~~~~~~~~~~~~~~~~%v~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n"
  const endline = "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n"
  var unitVariable uint
  var intVariable int
  var uintptrVariable uintptr
  var uint8Variable uint8
  var float32Variable float32
  var float64Variable float64
  var complex64Varibale complex64
  var complex128Varibale complex128
  var byteVariable byte
  var runeVariable rune
  var strtingVariable string
  var booleanVariable bool
// check default values
  fmt.Printf(printline,"Numeric Types Default Values")
  fmt.Printf(f,unitVariable,unitVariable)
  fmt.Printf(f,intVariable,intVariable)
  fmt.Printf(f,uintptrVariable,uintptrVariable)
  fmt.Printf(f,uint8Variable,uint8Variable)
  fmt.Printf(f,float32Variable,float32Variable)
  fmt.Printf(f,float64Variable,float64Variable)
  fmt.Printf(f,complex64Varibale,complex64Varibale)
  fmt.Printf(f,complex128Varibale,complex128Varibale)
  fmt.Printf("byte variable alias for "+f,byteVariable,byteVariable)
  fmt.Printf("rune variable alias for "+f,runeVariable,runeVariable)
  fmt.Print(endline)
  fmt.Printf(printline,"String Types Default Values")
  fmt.Printf(f,strtingVariable,strtingVariable)
  fmt.Print(endline)
  fmt.Printf(printline,"Boolean Types Default Values")
  fmt.Printf(f,booleanVariable,booleanVariable)
  fmt.Print(endline)
// complex number example
  fmt.Printf(printline,"Numeric Types -> Complex type Example")
  fmt.Printf("cmplx.Sqrt(-5 + 12i) = %v\n",cmplx.Sqrt(-5 + 12i))
  fmt.Print(endline)
  fmt.Printf(printline,"Booelan Types -> Operations")
  fmt.Printf("(true && true) = %v\n",true && true)
  fmt.Printf("(true && false) = %v\n",true && false)
  fmt.Printf("(false&& true) = %v\n",false && true)
  fmt.Printf("(false&& false) = %v\n",false && false)
  fmt.Print(endline)
  fmt.Printf("(true || true) = %v\n",true || true)
  fmt.Printf("(true || false) = %v\n",true || false)
  fmt.Printf("(false|| true) = %v\n",false || true)
  fmt.Printf("(false|| false) = %v\n",false|| false)
  fmt.Print(endline)

}