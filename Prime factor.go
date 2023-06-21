package main
import ( 
    "fmt"
        "math"
)
func largestPrimeFactor(n int) int {
    largest:=1
    for n%2==0 {
        largest =2
        n/=2
    }
    for i:=3;i<= int(math.Sqrt(float64(n))); i+= 2 {
        for n%i ==0 {
            largest =i
            n/=i
        }
    }
    if n>2 {
        largest = n
    }
    return largest
    
}


func main() {
    number :=14
    largest := largestPrimeFactor(number)
    fmt.Printf("The largest Prime Factor of %d is %d\n",number,largest)
  }
