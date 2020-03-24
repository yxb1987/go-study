package main

import (
	"fmt"
	"testing"
)

// 想要写出好的 Go 程序，单元测试是很重要的一部分。
// testing 包为提供了编写单元测试所需的工具，写好单元测试后，
// 我们可以通过 go test 命令运行测试。

// 为方便演示，例子的代码位于 main 包，实际上，单元测试的代码可以位于任何包下。
// 测试代码通常与需要被测试的代码位于同一个包下。

// 我们要测试下面这个简单的函数——返回最小值。
// 一般地，需要被测试的代码应该在类似于 intutils.go 的文件下，
// 其对应的测试文件应该被命名为 intutils_test.go。

func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// 通常编写一个名称以 Test 开头的函数来创建测试。

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Error* 会报告测试失败的信息，然后继续运行测试。
		// t.Fail* 会报告测试失败的信息，然后立即终止测试。
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// 单元测试可以重复，所以会经常使用 表驱动 风格编写单元测试，
// 表中列出了输入数据，预期输出，使用循环，遍历并执行测试逻辑。

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	// t.Run 可以运行一个 “subtests” 子测试，一个子测试对应表中一行数据。
	// 运行 go test -v 时，他们会分开显示。
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// 以啰嗦模式运行当前项目下的所有测试。

/* $ go test -v
== RUN   TestIntMinBasic
--- PASS: TestIntMinBasic (0.00s)
=== RUN   TestIntMinTableDriven
=== RUN   TestIntMinTableDriven/0,1
=== RUN   TestIntMinTableDriven/1,0
=== RUN   TestIntMinTableDriven/2,-2
=== RUN   TestIntMinTableDriven/0,-1
=== RUN   TestIntMinTableDriven/-1,0
--- PASS: TestIntMinTableDriven (0.00s)
    --- PASS: TestIntMinTableDriven/0,1 (0.00s)
    --- PASS: TestIntMinTableDriven/1,0 (0.00s)
    --- PASS: TestIntMinTableDriven/2,-2 (0.00s)
    --- PASS: TestIntMinTableDriven/0,-1 (0.00s)
    --- PASS: TestIntMinTableDriven/-1,0 (0.00s)
PASS
ok      examples/testing    0.023s */
