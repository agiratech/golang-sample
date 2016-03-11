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
  stringValue += strconv.Itoa(v2) // Type casting int to string
  if (v1 > 0) { Base_10_to_Base_2(v1)}
  return reverseString(stringValue)
}

func reverseString(s string) string {
  chars :=[]rune(s)
  for i,j :=0,len(chars) -1; i<j;i,j = i+1,j-1 { chars[i], chars[j] = chars[j], chars[i]}
  return string(chars)
}

func list_numbers_divided_by_two_numbers(x,y,z int){
  if (x < 1 || y < 1){
    fmt.Println("Enter number greater than 0")
  }else if( x == y) {
    fmt.Println("Both are same Numbers")
  }else{
    for i:=2; i<=z; i++ {
      if (i % y == 0 && i % x == 0) { fmt.Println(i)}
    }
  }
}

func min(m,n int) int {
 if m > n {
  return n
 }else{
  return m
 }
}


strconv.Itoa(v2) // Type casting int to string

func main() {
  fmt.Print(start,"Bigest Number")
  fmt.Print("Enter the length of the digit:")
  var i float64
  var j,k,n int
  fmt.Scanf("%g",&i)
  fmt.Println(largestBase10Value(i))
  fmt.Print(end)
  fmt.Print(start,"Base 10 --> Base 2")
  fmt.Print("Enter the :")
  fmt.Scanf("%d",&j)
  fmt.Println(Base_10_to_Base_2(j))
  fmt.Print(end)
  fmt.Print(start,"list Numbers commany divied by 2 numbers")
  fmt.Print("Enter the Lenth of List :")
  fmt.Scanf("%d",&n)
  fmt.Print("Enter the 1st number :")
  fmt.Scanf("%d",&k)
  fmt.Print("Enter the 2nd number :")
  fmt.Scanf("%d",&j)
  list_numbers_divided_by_two_numbers(j,k,n)
  fmt.Print(end)

}