package main

import (
	"fmt"
	"strings"
)

func main(){
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0] , a[1])
	fmt.Println((a))

	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)

	main2()
	main3()

	a1 := make([]int , 5) // make 
	main4("a1" , a1)
	
	b := make([]int , 0 ,5)
	main4("b" , b)

	c := b[:2]
	main4("c" ,c)

	d := c[2:5]
	main4("d" , d)

	main5()

	var s []int
	main6(s)

	s= append(s , 0) // append works on nil slices
	main6(s)
	s = append(s , 1) // The slice grows as needed
	main6(s)
	s = append(s , 2 , 3 ,4) // We can add more than one element at a time
	main6(s)

	main7()

}

// Slice -> Slice in go is similar to vector in c++ , and it also has similar functions like that we do in c++/c/python operations on string , append() , s.len() all like this 

func main2(){
	primes := [6]int{2 , 5 , 7 , 11 , 13}

	var s []int = primes[1:4]

	fmt.Println(s)
}

func main3(){
	var s []int
	fmt.Println(s , len(s) , cap(s))
	if s == nil{
		fmt.Println("nil!")
	}
}

func main4(s string , x[]int){
	fmt.Printf("%s len=%d cap =%d %v\n", s , len(x) , cap(x) , x)
}

func main5(){
	board := [][]string{
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i :=0 ; i < len(board) ; i++{
	fmt.Printf("%s\n" , strings.Join(board[i], " "))
	}
}

func main6(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s) , cap(s) , s)

}

func main7(){ // one could also use normal for loop but using range it benefit us by less code.
	var pow = []int{1 , 2 ,4 , 8 , 16 , 64 , 128}

	for i, v := range pow{
		fmt.Printf("2**%d = %d\n" , i , v)
	}
}
func Pic(dx, dy int) [][]uint8 {  // image creation
    picture := make([][]uint8, dy)

    for y := 0; y < dy; y++ {
        picture[y] = make([]uint8, dx)

        for x := 0; x < dx; x++ {
            picture[y][x] = uint8((x + y) / 2)
        }
    }

    return picture
}