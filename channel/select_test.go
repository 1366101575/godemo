package channel

import (
	"fmt"
	"testing"
)

//select 多路复用
//在某些场景下我们需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以接收将会发生阻塞。
//为了应对这种场景，Go内置了select关键字，可以同时响应多个通道的操作。
//select的使用类似于switch语句，它有一系列case分支和一个默认的分支。
//每个case会对应一个通道的通信（接收或发送）过程。select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。

/**
可处理一个或多个channel的发送/接收操作。
如果多个case同时满足，select会随机选择一个。
对于没有case的select{}会一直等待，可用于阻塞main函数。
*/
func TestSelect(t *testing.T) {
	t.Run("test0", func(t *testing.T) {
		ch := make(chan int, 1)
		for i := 0; i < 10; i++ {
			select {
			case ch <- i:
				fmt.Printf("ch <- %d\n", i)
			case data := <-ch:
				fmt.Printf("%d <- ch\n", data)
			}
		}

		fmt.Println("done")
	})
	/**
	zego@zegodeMacBook-Pro-83 channel % go test -v --run=TestSelect/test0
	=== RUN   TestSelect
	=== RUN   TestSelect/test0
	ch <- 0
	0 <- ch
	ch <- 2
	2 <- ch
	ch <- 4
	4 <- ch
	ch <- 6
	6 <- ch
	ch <- 8
	8 <- ch
	done
	--- PASS: TestSelect (0.00s)
	    --- PASS: TestSelect/test0 (0.00s)
	PASS
	ok  	godemo/channel	0.618s
	*/

}
