package main

import "fmt"

func main() {
	var nota, media float32

	for i := 0; i < 2; i++ {
		fmt.Scanf("%f", &nota)
		if nota < 0 || nota > 10 {
			fmt.Println("nota invalida")
			i--
		}else{
		  media += nota
    }
	}

	media = media / 2
  fmt.Printf("media = %.1f\n", media)
}
