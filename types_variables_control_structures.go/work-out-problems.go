// finding largest number in base 10

package main

import (
          "fmt"
          "math"
          "strconv"
      )

var stringValue string= ""// without type declaration
const start = "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~%v~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n\n"
const end = "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n\n"

func largestBase10Value(val float64) float64 {
  return (math.Pow(10,val) - 1)
}

func Base_10_to_Base_2(val int) string {
  v1 := (val / 2)
  v2 := (val - (v1*2))
  stringValue += strconv.Itoa(v2)
  if (v1 > 0) { Base_10_to_Base_2(v1)}
  return reverseString(stringValue)
}

func reverseString(s string) string {
  chars :=[]rune(s)
  for i,j :=0,len(chars) -1; i<j;i,j = i+1,j-1 { chars[i], chars[j] = chars[j], chars[i]}
  return string(chars)
}


func main() {
  fmt.Print(start,"Bigest Number")
  fmt.Print("Enter the length of the digit:")
  // var i float64
  var j int
  // fmt.Scanf("%g",&i)
  // fmt.Println(largestBase10Value(i))
  // fmt.Print(end)
  // fmt.Print(start,"Base 10 --> Base 2")
  fmt.Print("Enter the :")
  fmt.Scanf("%d",&j)
  fmt.Println(Base_10_to_Base_2(j))
}