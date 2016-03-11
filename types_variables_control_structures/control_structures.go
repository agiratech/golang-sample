


package main

import (
        "fmt"
        )
const start = "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~%v~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n\n"
const end = "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n\n"

// For loop example

func for_loop(n int){
  //
  var pre int = 0
  var pre1 int = 1
  var feb_series int
  fmt.Println("0\n1")
  for i := 2; i < n; i++ {
    feb_series = pre + pre1
    fmt.Println(feb_series)
    pre = pre1
    pre1 = feb_series
  }
}


func if_else_statement(n int) {
  if (n < 2000){
    fmt.Printf("The year %d is not exsist in 21st century\n",n)
  }else if(n % 4 == 0) {
    fmt.Printf("The year %d is a leap year\n",n)
  }else {
   fmt.Printf("The year %d is not a leap year\n",n)
  }

}

func switch_statement(month int) {
  switch month {
    case 1: fmt.Println("Jan")
    case 2: fmt.Println("Feb")
    case 3: fmt.Println("March")
    case 4: fmt.Println("Apr")
    case 5: fmt.Println("May")
    case 6: fmt.Println("June")
    case 7: fmt.Println("July")
    case 8: fmt.Println("Aug")
    case 9: fmt.Println("Sep")
    case 10: fmt.Println("Oct")
    case 11: fmt.Println("Nov")
    case 12: fmt.Println("Dec")
    default: fmt.Println("The number Should with in 1 to 12")
  }
}

func main() {
  fmt.Printf(start,"Fibonacci Sequence using For Loop")
  fmt.Println("Enter Last Number Of Series:")
  var n int
  fmt.Scanf("%d",&n)
  for_loop(n)
  fmt.Printf(end)
  fmt.Printf(start,"Year is leap year or not")
  fmt.Print("Enter the year :")
  fmt.Scanf("%d",&n)
  if_else_statement(n)
  fmt.Printf(end)
  fmt.Printf(start,"number to Month Name")
  fmt.Print("Enter the Number :")
  fmt.Scanf("%d",&n)
  switch_statement(n)
  fmt.Printf(end)

}