package main

import (
	"fmt"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/os/gproc"
	"github.com/gogf/gf/util/gconv"
	"runtime"
	"sync"
)

func pingAll(ipscan, startS, endS string) (string, string) {
	gproc.StartTime()
	//ipscan := "101.6.19."
	start := gconv.Int(startS)
	end := gconv.Int(endS)

	wait := sync.WaitGroup{}
	wait.Add(end - start + 1)
	backinfo := make(chan string, end-start+1)
	for i := start; i <= end; i++ {
		go pingIp(ipscan+gconv.String(i), backinfo, &wait)
	}
	wait.Wait()
	//配合wait 此时可以关闭chan
	//必须是带指定缓存的chan才可以这样读写chan
	close(backinfo)

	//关闭之后的chan可以多种安全方式读取
	//method1 直接for循环取
	canpingips := garray.NewStrArray(true)
	for v := range backinfo {
		canpingips.Append(v)
	}
	ips := ""
	canpingips.Iterator(func(k int, v string) bool {
		k++
		if k%3 == 0 {
			ips = ips + "  " + v + "\n"
		} else {
			ips = ips + "  " + v
		}
		return true
	})
	return ips, "total time:" + gproc.Uptime().String() + "\n can ping ip nums:" + gconv.String(canpingips.Len()) + "\n"
}

//检测ip是否可以ping通
func pingIp(ip string, back chan string, w *sync.WaitGroup) {
	//fmt.Println("正检测ip:", ip)
	defer w.Done()
	//重试三次 每次等待2s
	cmd_windows := "ping %s -w 3000 -n 1"
	cmd_linux := "ping %s -w 3 -n 1"
	system := runtime.GOOS
	switch system {
	case "windows":
		{
			_, err := gproc.ShellExec(fmt.Sprintf(cmd_windows, ip))
			if err != nil {
			} else {
				back <- ip
			}
			break
		}
	case "linux":
		{
			_, err := gproc.ShellExec(fmt.Sprintf(cmd_linux, ip))
			if err != nil {
			} else {
				back <- ip
			}
			break
		}
	case "darwin":
		{
			_, err := gproc.ShellExec(fmt.Sprintf(cmd_linux, ip))
			if err != nil {
			} else {
				back <- ip
			}
			break
		}
	}
}
