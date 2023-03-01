package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	test_cmdRun()
	//test_combinedOutput()
	//test_StdoutPipe()
}

/**
 * 1、执行命令基本操作
 */
func test_cmdRun() {
	//声明变量类型
	var (
		cmd *exec.Cmd //可以跳转到exec.Command()方法查看返回类型
		err error     //可以到cmd.Run()方法查看返回类型
	)
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序，第二个参数可以无限指定携带参数
	cmd = exec.Command("ls", "-l", "-a")

	/**
	func (c *Cmd) Run() error
	1、Run方法执行命令并阻塞直到完成。
	2、如果命令成功执行，stdin、stdout、stderr的转交没有问题，并且返回状态码为0，方法的返回值为nil。
	3、如果函数没有执行或者执行失败，会返回*ExitError类型的错误；否则返回的error可能是表示I/O问题。
	*/
	err = cmd.Run()
	fmt.Println(err) //返回<nil>
}

/**
 * 2、获取执行命令的结果
 */
func test_combinedOutput() {
	cmd := exec.Command("ls", "-l", "-a")

	//执行命令并返回标准输出和错误输出合并的切片。
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

/*
*

	3、通过stdout管道读取执行命令后的结果：
	func (c *Cmd) StdoutPipe() (io.ReadCloser, error)　　　　//err 返回的是执行函数时的错误
*/
func test_StdoutPipe() {
	cmd := exec.Command("ls", "-l", "-a")
	stdout, err := cmd.StdoutPipe() //指向cmd命令的stdout，然后就可以从stdout读出信息

	/**
	1、Start 非阻塞执行，，不会等待该命令完成就直接返回
		func (c *Cmd) Start() error
	2、Wait Wait会阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的。
		func (c *Cmd) Wait() error
		2.1 如果命令成功执行，stdin、stdout、stderr的转交没有问题，并且返回状态码为0，方法的返回值为nil；
		2.2 如果命令没有执行或者执行失败，会返回*ExitError类型的错误；否则返回的error可能是表示I/O问题。Wait方法会在命令返回后释放相关的资源。
	*/
	startErr := cmd.Start()
	fmt.Println(startErr) //<nil>

	//处理中间任务
	content, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content)) //输出ls命令查看到的内容

	waitErr := cmd.Wait()
	fmt.Println(waitErr) //<nil>
}
