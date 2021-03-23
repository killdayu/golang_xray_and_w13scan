package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

var urlfile string

func init() {
	flag.StringVar(&urlfile, "f", "", "Open shadow") //StringVar定义了一个有指定名字，默认值，和用法说明的string标签。 参数p指向一个存储标签解析值的string变量。
}

func main() {
	flag.Parse()
	//打开文件
	urls, err := ioutil.ReadFile(urlfile) //打开passfile
	if err != nil {
		log.Fatalln(err)
	}
	//defer urls.Close()

	targets := strings.Split(string(urls), "\n")
	for _, target := range targets {
		fmt.Println("test")
		go exec_cmd(target)
		proxy_two()
	}
}

func exec_cmd(target string) {
	cmd := exec.Command("./crawlergo", "-c", "google-chrome", "--push-to-proxy", "http://127.0.0.1:9999/", string(target))
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))
}
