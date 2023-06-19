package main  

import (
	"fmt"	
	"github.com/gofrs/uuid"	
)

func main(){
	fmt.Println("assignment")
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Println("failed to generate UUID: %v", err)
	}
	fmt.Println("generated Version 4 UUID %v", u2)
}
