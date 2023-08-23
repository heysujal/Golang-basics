package main
import(
	"fmt"
	"reflect"
)
// notice type comes after the variable name
func add1(x int,y int) int {
	return x + y
}
// When two or more consecutive named function parameters share a type, 
// you can omit the type from all but the last.
func add2(x,y int) int {
	return x + y
}
// A function can return any number of results.
func swap(x,y string) (string, string){
	return y,x
}

// Named return values
// this is called naked return and you should avoid it in longer functions
func split(sum int) (x,y int){
	sum = sum*2 + 4
	x = sum*sum
	y = x-2
	return
}

func main(){
	fmt.Println(add1(2,3))
	fmt.Println(add2(10,90))
	var a = 2
	var b = 3
	fmt.Println(a,b)
	c,d := swap("Sujal", "Gupta")
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("This is for split()")
	fmt.Println(split(0))
}