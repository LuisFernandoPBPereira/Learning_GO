package main

import "fmt"

func main() {
  var x, y, divisao float32
  var n int

  fmt.Scanf("%d", &n);

  for i:=0; i<n; i++{
      fmt.Scanf("%f", &x)
      fmt.Scanf("%f", &y)

      if y==0 {
          fmt.Println("divisao impossivel")
      }else{
        divisao = x / y
        fmt.Printf("%.1f\n", divisao)
      }

  }
}
