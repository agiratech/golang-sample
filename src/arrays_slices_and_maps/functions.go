package main

import "fmt"

const start = "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~%v~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n"
const end = "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n"
const title = "%v ----------->%v\n"

func main() {
  a := "arun"
  x,y := 1,2
  variadicFunctoin(a,x,y)
}
// function get multiple params by ... this symbol
// variadic param also in the last
func variadicFunctoin(b string,a ...int) {
  fmt.Printf(title,"The values of a is :",a)
  fmt.Printf(title,"The values of b is :",b)
}
//Function inside a function
func closureFunction() {

}