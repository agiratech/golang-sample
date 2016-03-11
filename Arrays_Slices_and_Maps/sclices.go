// Arrays have their place, but they're a bit inflexible,
// so you don't see them too often in Go code. Slices, though, are everywhere.
// They build on arrays to provide great power and convenience.

// Unlike an array type, a slice type has no specified length.

// A slice literal is declared just like an array literal, except you leave out the element count:

// letters
//  []type{elemnt1,ele2,...} ex ---> a:= []string{"a", "b", "c", "d"}
//
//
// make([],len,cap) -->s := make([]byte, 5)

// Go includes two built-in functions to assist with slices: append and copy.

package main

import "fmt"

const end = "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n\n"
const start = "~~~~~~~~~~~~~~~~~~~~~~~~~~%v~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n\n"


func main(){
  mySlice := []int{2,3,5,7,11,13}
  mynewSlice := make([]int,10,10)
  myarray := [...]int{2,3,5,7,11,13}
  fmt.Printf(start,"Slice example")
  fmt.Println(mySlice)
  fmt.Printf("slice for [1:4] ----------> %v\n",mySlice[1:4])
  fmt.Printf("slice for [:3] -----------> %v\n",mySlice[:3])
  fmt.Printf("slice for [4:] -----------> %v\n",mySlice[4:])
  fmt.Printf(start,"Passing Array Function")
  ArrayFunction(myarray[:])
  fmt.Printf(start,"Passing Sclice Function")
  SliceFunction(mySlice)
  fmt.Printf(start,"Copy Slice into New Slice")
  SliceCopy(mySlice,mynewSlice)
  fmt.Printf(start,"Append 100 with  Slice")
  SliceAppend(mySlice,100)
  fmt.Println(end)
}

func ArrayFunction(a []int) {
  fmt.Println(a)
}

func SliceFunction(a []int) {
  fmt.Println(a)
}

func SliceCopy(a,b []int) {
  copy(b,a)
  fmt.Println(b)
}


func SliceAppend(a []int,b int) {
  c := append(a,b)
  fmt.Println(c)
}
