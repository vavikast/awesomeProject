package main

import (
	"awesomeProject/ES/tour/cmd"
	"log"
)

func main()  {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err : %v ",err)
	}

}