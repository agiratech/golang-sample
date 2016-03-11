// array litrals

// variable := [length]type{v1,v2} ex ---> b := [2]string{"Penn", "Teller"}

// variable := [...]type{v1,v2} ex ----> b := [...]string{"Penn", "Teller"}
// in this example compalier will count the element count

// array dealeration
// var a [4]int; Or var a []int

package main

import "fmt"

func arrayFunction(array []int) {
  fmt.Println(array)

}

func main() {
    // Create an array of three ints.
    array := [...]int{10, 20, 30}

    // Loop over three ints and print them.
    for i := 0; i < len(array); i++ {
      fmt.Println(array[i])
    }

    arrayFunction(array[:])
}