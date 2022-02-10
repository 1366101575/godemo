package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	//无缓冲通道
	t.Run("test0", func(t *testing.T) {
		ch := make(chan int)

		go func() {
			for i := 0; i < 10; i++ {
				ch <- i
			}
			close(ch)
			fmt.Println("ch send finish, closed")
		}()

		for i := range ch {
			time.Sleep(time.Second * 1)
			fmt.Println(i)
		}

		fmt.Println("done")
	})

	//有缓冲通道，当通道关闭后 for range 会把通道里的数据全部取出后退出循环
	t.Run("test1", func(t *testing.T) {
		ch := make(chan int, 5)

		go func() {
			for i := 0; i < 10; i++ {
				ch <- i
			}
			close(ch)
			fmt.Println("ch send finish, closed")
		}()

		for i := range ch {
			time.Sleep(time.Second * 1)
			fmt.Println(i)
		}

		fmt.Println("done")
	})

	//goroutine worker pool
	t.Run("test2", func(t *testing.T) {
		jobs := make(chan int, 100)
		results := make(chan string, 100)

		//生成工作池：3个处理任务的goroutine，会自动从任务列表中拿任务进行处理，并将处理结果保存下来
		for w := 0; w < 3; w++ {
			go worker(w, jobs, results)
		}

		//生成任务列表：5个待处理的任务
		for j := 0; j < 5; j++ {
			jobs <- j
		}
		close(jobs)

		//输出工作池处理后的任务结果
		for r := 0; r < 5; r++ {
			fmt.Printf("result : %v \n", <-results)
		}
	})

}

func worker(wid int, jobs <-chan int, results chan<- string) {
	for j := range jobs {
		fmt.Printf("wid:%d job:%d start \n", wid, j)
		time.Sleep(time.Second)
		fmt.Printf("wid:%d job:%d end \n", wid, j)
		results <- fmt.Sprintf("result: %d * %d = %d\n", j, j, j*j)
	}
}
