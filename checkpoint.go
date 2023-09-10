
package main
import(
    "fmt"
    "sync"
    "time"
)

func main(){
    var wg sync.WaitGroup
    start:=make(chan struct{})
    var num int
    fmt.Printf("enter the no of workers")
    fmt.Scan(&num)
    
    for i:=1;i<=num;i++{
        wg.Add(1)
        go func(id int){
            defer wg.Done()
            <-start
            fmt.Printf("Worker %d is starting to work\n",id)
            time.Sleep(time.Second)
            
            fmt.Printf("Worker %d has reached checkpoint\n",id)
            time.Sleep(time.Second)
            time.Sleep(time.Second)
            
            fmt.Printf("Worker %d resuming the work\n",id)
            time.Sleep(time.Second)
            
            
        }(i)
    }
    
    fmt.Println("Starting workers.....")
    close(start)
    wg.Wait()
    fmt.Println("All workers completed their work")
}
