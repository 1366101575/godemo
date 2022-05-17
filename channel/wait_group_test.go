package channel

import (
	"fmt"
	"sync"
	"testing"
)

/**
对 https://www.liwenzhou.com/posts/Go/14_concurrence/#autoid-1-6-0 的学习笔记
并发安全和锁

有时候在Go代码中可能会存在多个goroutine同时操作一个资源（临界区），这种情况会发生竞态问题（数据竞态）。

Go语言中可以使用sync.WaitGroup来实现并发任务的同步。
sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了N 个并发任务时，就将计数器值增加N。
每个任务完成时通过调用Done()方法将计数器减1。通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。
*/

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func TestWaitGroup(t *testing.T) {
	//开启了两个goroutine去累加变量x的值，这两个goroutine在访问和修改x变量的时候就会存在数据竞争，导致最后的结果与期待的不符。
	t.Run("test0", func(t *testing.T) {
		wg.Add(2)
		go add()
		go add()
		wg.Wait()
		fmt.Println(x)
	})

	//互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。Go语言中使用sync包的Mutex类型来实现互斥锁。
	//使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；
	//当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。
	t.Run("test1", func(t *testing.T) {
		wg.Add(2)
		go add2()
		go add2()
		wg.Wait()
		fmt.Println(x)
	})

	/**
	读写互斥锁
	互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，
	这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用sync包中的RWMutex类型。

	读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
	当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。

	需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。
	*/
}

func add() {
	defer wg.Done()

	for i := 0; i < 5000; i++ {
		x = x + 1
	}
}

func add2() {
	defer wg.Done()

	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
}
