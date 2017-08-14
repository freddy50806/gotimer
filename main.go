package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type event struct {
	Id         int64
	Context    string
	Datetime   string
	Excutetime string
}

type Job struct {
	newjob   event
	jobtimer *time.Timer
}

// type JobInfo struct {
// 	Id       int64  `json:"id" validate:"required"`
// 	JobType  int64  `json:"type" validate:"required,max=1"`
// 	Context  string `json:"context" validate:"required"`
// 	Datetime string `json:"datatime" validate:"required"`
// }
func printls(abc string) {
	println(abc)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("add cmd-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		strArray := strings.Split(text, " ")

		if strings.Compare("exit", text) == 0 {
			fmt.Println("Bye~have great time")
			break
		}

		binary, lookErr := exec.LookPath(strArray[0])
		if lookErr == nil {
			args := strArray[1:]
			cmd := exec.Command(binary, args...)

			fmt.Print("add time-> ")
			var time_input int64
			fmt.Scan(&time_input)

			f := func() {
				stdout, err := cmd.StdoutPipe()
				defer stdout.Close()
				if err != nil {
					log.Fatal(err)
				}
				// 保证关闭输出流
				defer stdout.Close()
				// 运行命令
				if err := cmd.Start(); err != nil {
					log.Fatal(err)
				}
				// 读取输出结果
				opBytes, err := ioutil.ReadAll(stdout)
				if err != nil {
					log.Fatal(err)
				}
				log.Println(string(opBytes))
			}
			fmt.Println(time_input)
			time.AfterFunc(time.Duration(time_input)*time.Second, f)
		} else {
			fmt.Print("input error")
			reader.ReadString('\n')
		}

	}
	println("GG")
	// compeltech := make(chan int64)
	// case1 := event{1, "GG1", "", ""}
	// job1 := make{Job}
	// case2 := event{2, "GG2", "", ""}

}
