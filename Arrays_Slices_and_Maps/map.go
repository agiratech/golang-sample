// Maps are somewhat similar to what other languages call “dictionaries” or “hashes”.
// A map maps keys to values.

// maps must be created with make (not new) before use.
// The nil map is empty and cannot be assigned to.

// The builtin len returns the number of key/value pairs when called on a map.

// The builtin delete removes key/value pairs from a map.

package main

import "fmt"

const start = "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~%v~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n"
const end = "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n"
const title = "%v ----------->%v\n"
func main() {
    m := make(map[string]int)
    fmt.Printf(start,"Maps")
    fmt.Printf(title,"Empty Map created by make",m)
    m["k1"] = 7
    fmt.Printf(title,"assining Value to map m['k1'] = 7",m)
    m["k2"] = 13
    fmt.Println("map:", m)
    v1 := m["k1"]
    fmt.Println("v1: ", v1)
    fmt.Println("len:", len(m))
    delete(m, "k2")
    fmt.Printf(title,"Fetch the value of k2",m)
    _, prs := m["k2"]
    fmt.Println("prs:", prs)
    fmt.Printf(start,"assining Value to map in a single line")
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
    fmt.Println(end)
}
