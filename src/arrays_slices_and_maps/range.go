// range on arrays and slices provides both the index and value for each entry.
// range on strings iterates over Unicode code points.
// The first value is the starting byte index of the string and the second the string itself.
// range on map iterates over key/value pairs.

package main

import "fmt"
const start = "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~%v~~~~~~~~~~~~~~~~~~~~~~~~~~\n"
const end = "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n"
func main() {
  fmt.Printf(start,"Range")
  array := []int {1,2,3,4,5,6,7,8,9,10}
  hashTable := map[int]string { 1:"arun",2:"kumar"}
  stringVal := "Arun"
  fmt.Printf("The defined Map Is ~~~~~~~~~~~~~~~~~~~~> %v\n",hashTable)
  fmt.Printf("The defined Array Is ~~~~~~~~~~~~~~~~~~~~> %v\n",array)
  fmt.Printf("The defined String Is ~~~~~~~~~~~~~~~~~~~~> %v\n",stringVal)
  map_with_range(hashTable)
  array_with_range(array)
  string_with_range(stringVal)

}

func map_with_range(m map[int]string) {
  fmt.Printf(start,"Map part executing...........")
  for a,m := range m {
    fmt.Printf("The value of the _ ---> %v\n",a)
    fmt.Printf("The value of the m ---> %v\n",m)
  }
  fmt.Printf(end)
}

func array_with_range(array []int) {
  fmt.Printf(start,"Array part executing...........")
  for a,num := range array {
    fmt.Printf("The value of the _ ---> %v\n",a)
    fmt.Printf("The value of the num ---> %v\n",num)
  }
  fmt.Printf(end)
}

func string_with_range(s string) {
  fmt.Printf(start,"String part executing...........")
  for a,str := range s {
    fmt.Printf("The value of the _ ---> %v\n",a)
    fmt.Printf("The value of the s ---> %v\n",str)
  }
  fmt.Printf(end)
}