package main

// fmt -> is basically for input/output (ln ->output , f -> input)

//math -> performing mathematical operations

// math/rand -> for generating a random number

import(
	"fmt"
	"math"
	"math/rand"
)

func main(){
	fmt.Println("Hello world")
	fmt.Println("Random number is", rand.Intn(100))
	fmt.Printf("Now you have %g problems \n", math.Sqrt(7))
	fmt.Printf("New number is %g " , math.Sqrt(8))
	fmt.Println(math.Pi)
}