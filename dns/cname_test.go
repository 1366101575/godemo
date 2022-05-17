package array

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
	"testing"
	"time"
)

/**
版权声明：本文为CSDN博主「huowuyou」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/huowuyou/article/details/111027521
*/
func TestCname(t *testing.T) {

	/*
		官方net包调用的是系统API, 所以在不同的系统上可能有不同的结果, 我测试当一个域名的解析记录如下时, linux和windows返回的结果不一致.

		bysir.xyz.           297     IN      CNAME   blog.bysir.top.
		blog.bysir.top.      60      IN      CNAME   bysir.top.
		bysir.top.           60      IN      A       78.178.25.149

		windows下net.LookupCNAME会返回第一层cname, 即blog.bysir.top., linux下会返回第二层cname, 即bysir.top..
	*/
	t.Run("test0", func(t *testing.T) {

		domain := "tieba.baidu.com"
		info, err := net.LookupCNAME(domain)
		fmt.Printf("%v CNAME %v  err:%v\n", domain, info, err)

		fmt.Println("done")
	})

	//获取完整的解析记录, 或者是精确的第一层cname
	t.Run("allCname", func(t *testing.T) {
		domain := "tieba.baidu.com"
		domain = "publish-tencent1.zego.im"
		cname, err := miekgDnsCname(domain, "114.114.114.114")
		fmt.Println(err)
		for _, c := range cname {
			fmt.Println(c)
		}

		fmt.Println("done")
	})

	//查找DNS A记录
	//使用 Go 语言的标准库 net.LookupIP() 接受域名的字符串参数,返回 net.IP的切片. 这个 net.IP 对象包含IPv4地址和IPv6地址.
	t.Run("test1", func(t *testing.T) {
		domain := "tieba.baidu.com"
		iprecords, _ := net.LookupIP(domain)

		for _, ip := range iprecords {
			fmt.Println(ip)
		}

		fmt.Println("done")
	})
}

// CNAME2 返回所有层的cname github.com/miekg/dns包
// src: 域名
// dnsService: dns服务器, 如114.114.114.114
func miekgDnsCname(src string, dnsService string) (dst []string, err error) {
	c := dns.Client{
		Timeout: 5 * time.Second,
	}

	var lastErr error
	// retry 3 times
	for i := 0; i < 3; i++ {
		m := dns.Msg{}
		// 最终都会指向一个ip 也就是typeA, 这样就可以返回所有层的cname.
		m.SetQuestion(src+".", dns.TypeA)
		r, _, err := c.Exchange(&m, dnsService+":53")
		if err != nil {
			lastErr = err
			time.Sleep(1 * time.Second * time.Duration(i+1))
			continue
		}

		dst = []string{}
		for _, ans := range r.Answer {
			record, isType := ans.(*dns.CNAME)
			if isType {
				dst = append(dst, record.Target)
			}
		}
		lastErr = nil
		break
	}

	err = lastErr

	return
}
