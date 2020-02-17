package main

import (
	"fmt"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/os/gproc"
	"github.com/gogf/gf/util/gconv"
	"github.com/sparrc/go-ping"
	"runtime"
	"sync"
	"time"
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
		go pingIpV2(ipscan+gconv.String(i), backinfo, &wait)
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
	// gproc.Uptime().String() +
	return ips, "can ping ip nums:" + gconv.String(canpingips.Len()) + "\n"
}

//检测ip是否可以ping通
func pingIpV2(ip string, back chan string, w *sync.WaitGroup) {
	//fmt.Println("正检测ip:", ip)
	defer w.Done()
	canPing := ServerPing(ip)

	if canPing == true {
		back <- ip
	} else {
	}
}

//检测ip是否可以ping通
//结合fyne ui，会显示cmd窗口
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
func ServerPing(ip string) bool {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		panic(err)
	}
	pinger.Count = gconv.Int(1)
	pinger.Timeout = time.Duration(3 * time.Millisecond)
	pinger.SetPrivileged(true)
	pinger.Run() // blocks until finished
	stats := pinger.Statistics()
	// 有回包，就是说明IP是可用的
	if stats.PacketsRecv >= 1 {
		return true
	}
	return false
}
