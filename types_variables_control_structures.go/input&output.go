package main

import (
    "fmt"
)

func arun(i int) string {

    fmt.Println(i)

    return "20"
}

func main(){
    var j int
    fmt.Scanf("%d",&j)
    fmt.Println(j)
    arun(j)
}

