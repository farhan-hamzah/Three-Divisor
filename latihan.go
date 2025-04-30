package main
import "fmt"

func main(){
	var n, i, jum int 
	fmt.Scan(&n)
	for i = 1; i <= n; i++{
		if n%i == 0{
			jum++
		}
	}
	if jum == 3{
		fmt.Print(true)
	}else{
		fmt.Print(false)
	}
}

