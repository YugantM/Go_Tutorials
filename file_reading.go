package main

import("fmt"
       "io/ioutil"
       "strings"
      // "./frame"
       )

func main(){

  data, err := ioutil.ReadFile("text.txt")
      if err != nil {
          fmt.Println("File reading error", err)
          return
      }
      //fmt.Println("Contents of file:", string(data))

      rows:= strings.Split(string(data),"\n")
      fmt.Println("\n\nContents of file:", rows[1])
/*
      y := frame.Demo{}
      y.Val = 78

      fmt.Println("Printing",y.Val)
    */
}
