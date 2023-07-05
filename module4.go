package main
import(
	"fmt"
	"sync"
	"time"
)

func worker(id int ,checkpoint chan bool,resume chan bool,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Worker %d Starting\n",id)
	time.Sleep(time.Duration(id) *time.Second)
	fmt.Println("Worker %d checkout\n",id)
	checkpoint<- true
	<- resume
	fmt.Println("Worker %d resuming\n",id)
}