package main

import(
	"fmt"
 	"math/cmplx"
)
// func -> taking 2 arguements in go we also have to define type of the variable similar to c++ but syntax is lil bit changed

// This is how you define variables in Go
var c , python , java bool 
var i int

var j , k int = 1,2
var c1 , python1 , java1 = true , false , "no!"

var(
	ToBe bool = false
	z complex128 = cmplx.Sqrt(-5 + 12i)
	MaxInt uint64 = 1<<64 -1 // 1<<64 = 2^64 => 1<<64 -1 = 2^64 -1
)

const World = "World" // const -> constant variable which means we cannot change the value of this variable

func add(x int , y int) int{
	return x+y
}


// we shortened this func parammeters coz both vars share same var type
func sub(x , y int) int{
	return x-y
}
// func to swap 2 strings and return them in reverse order
func swap(x , y string) (string , string){
	return y,x
}
// A return statement without arguements return the named return values. This is known as "naked" return
func split(sum int) ( x, y int){
	x = sum*4/9
	y = sum -x
	return 
}

func main(){
	fmt.Println(add(10,20)) // func add
	fmt.Println(sub(20,10)) // func sub
	a , b := swap("hello" , "World") // inside a function we can define variables using := operator which is known as short variable declaration but outside a function we have to use var keywod to define variables
	fmt.Println(a , b) // func swap
	fmt.Println(split(17)) // split function
	fmt.Println(i , c , python , java) // var declaration 
	fmt.Println(j , k , c1 , python1 , java1) // var declaration with initialization
	fmt.Printf("Type: %T Value: %v \n", ToBe, ToBe) //  %T means -> print the type of the value , %v -> print the value of the variable
	fmt.Printf("Type: %T Value: %v \n", z, z) // same logic here 
	fmt.Printf("Type: %T Value: %v \n", MaxInt, MaxInt) // same logic here 
	fmt.Println("Hello" , World)
}
