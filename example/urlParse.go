package example

import (
	"fmt"
	"net/url"
	"strings"
)

func UrlParse() {
	p := fmt.Println
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	p(u.Scheme)

	// User包含了所有的验证信息，使用
	// Username和Password来获取单独的信息
	p(u.User)
	p(u.User.Username())
	pwd, _ := u.User.Password()
	p(pwd)

	// Host包含了主机名和端口，如果需要可以
	// 手动分解主机名和端口
	p(u.Host)
	h := strings.Split(u.Host, ":")
	p(h[0])
	p(h[1])

	// 这里我们解析出路径和`#`后面的片段
	p(u.Path)
	p(u.Fragment)

	// 为了得到`k=v`格式的查询参数，使用RawQuery。你可以将
	// 查询参数解析到一个map里面。这个map为字符串作为key，
	// 字符串切片作为value。
	p(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	p(m["k"][0])
}
