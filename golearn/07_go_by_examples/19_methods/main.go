package main

import (
	"fmt"
)

// Go 支持为结构体类型定义方法(methods) 。

type rect struct {
	width, height int
}

// 这里的 area 方法有一个 _接收者类型_：rect。
func (r *rect) area() int {
	return r.width * r.height
}

// 可以为值类型或者指针类型的接收者定义方法。 这是一个值类型接收者的例子。
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// 这里我们调用上面为结构体定义的两个方法。
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// 调用方法时，Go 会自动处理值和指针之间的转换。 想要避免在调用方法时产生一个拷贝，
	// 或者想让方法可以修改接受结构体的值， 你都可以使用指针来调用方法。
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

/* $ go run methods.go
area:  50
perim: 30
area:  50
perim: 30 */
