package function

import (
	"fmt"
	"testing"
	"time"
)

func getAdd() func() int {
	i := 0
	fmt.Printf("init %p\n", &i) //打印变量地址
	return func() int {
		i += 1
		fmt.Printf("%p\n", &i) //打印变量地址
		return i
	}
}

type test struct {
	Value int
}

func TestFunction(t *testing.T) {
	t.Run("test0", func(t *testing.T) {
		add := getAdd()

		fmt.Println(add())
		fmt.Println(add())
		fmt.Println(add())

		add2 := getAdd()

		fmt.Println(add2())
		fmt.Println(add2())
	})

	t.Run("test1", func(t *testing.T) {
		x, y := 1, 2

		defer func(a int) {
			fmt.Println("defer x, y = ", a, y) //y为闭包引用
		}(x) //x值拷贝 调用时传入参数

		x += 100
		y += 200

		fmt.Println(x, y)
	})

	//go test -v -run=TestFunction/test2
	t.Run("test2", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			//多次注册延迟调用，相反顺序执行
			defer func(a int) {
				fmt.Printf("a=%d i=%d \n", a, i) //闭包引用局部变量
			}(i)

			fmt.Print(i)
			if i == 2 {
				fmt.Printf("\n")
			}
		}
	})

	t.Run("test3", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			go func() {
				//i变量值也是引用, 创建10个线程执行函数， for循环执行过程中可能执行完的时候，线程刚好处于i的某个值。
				fmt.Println(i)
			}()
		}
		time.Sleep(time.Second * 1)
	})

	t.Run("test4", func(t *testing.T) {
		ch := make(chan int, 1)
		for i := 0; i < 10; i++ {
			go func() {
				//ch <- 1 //用来控制串行的chan不能放在前面，要放最后

				time.Sleep(time.Millisecond * 100) //业务处理

				fmt.Println(i)

				ch <- 1
			}()
			<-ch
		}
		time.Sleep(time.Second * 3)
	})

	fmt.Printf("\n\n")
}
