package main

import (
  "encoding/csv"
  "fmt"
  "io"
  "os"
  "strings"
  "./psqlConn"
  // "strconv"
)

func init() {
  psqlConn.InitDB()
}

type CsvHeader struct {
  header []string
}

func (c *CsvHeader) assignValue(a []string)  {
  c.header = strings.Split(a[0],",")
}

func (c *CsvHeader) returnValue() []string {
  return c.header
}

// func CheckType(x interface{}) bool {
//   switch x.(type) {
//     case string:
//           return  false
//     default:
//           return  true
//   }
// }

// func ParseString(x int) string {
//   return strconv.Itoa(x)
// }

func main() {
  file, err := os.Open("sample.csv")
  c := &CsvHeader{}
  var x int =0
  if err != nil {
    // err is printable
    fmt.Println("Error:", err)
    return
  }
  // automatically call Close() at the end of current method
  defer file.Close()
  //
  reader := csv.NewReader(file)
  reader.Comma = '\001'
  lineCount := 0
  for {
    // read just one record, but we could ReadAll() as well
    record, err := reader.Read()
    // end-of-file is fitted into err
    if err == io.EOF {
      break
    } else if err != nil {
      fmt.Println("Error:", err)
      return
    }

    //skip header line in csv
    if x == 0 {
      x++
      c.assignValue(record)
      psqlConn.CreateTable(c.returnValue())
      continue
    }
    stringVal := ""
    var str string
    for i,rec := range strings.Split(record[0],",") {
      rec = strings.Replace(rec, "\"", "'", -1)
      if str  = ","; i == 0 { str = "" }
      if rec = rec; rec == "" { rec = "NULL"}
      stringVal += (str+rec)
    }
    psqlConn.InsertRec(stringVal) // fmt.Println(" ", record)

    lineCount += 1
  }
}