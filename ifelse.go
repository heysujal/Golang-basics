package main
import(
	"fmt"
)

func isPrime(num int){
	var flag bool = true
	for i:=2 ; i < num ; i++{
		if num%i == 0{
			flag = false
			fmt.Println("Not a prime")
			break
		}
	}
	if flag{
		fmt.Println("This is a prime number")
	}
}

func isEven(num int) bool{
	return num%2 == 0
}

func main(){
	isPrime(4)
	fmt.Println(isEven(2), isEven(7))
}