package main

import (
	"bytes"
	"container/list"
	"fmt"
	"strconv"
)

func thousandSeparator(n int) string {
	res := ""
	c := 0
	l := list.New()
	for ; n >= 10; n /= 10 {
		l.PushBack(strconv.Itoa(n % 10))
		c += 1
		if c == 3 {
			l.PushBack(".")
			c = 0
		}
	}
	l.PushBack(strconv.Itoa(n))
	for l.Len() > 0 {
		e := l.Back()
		l.Remove(e)
		res += e.Value.(string)
	}
	return res
}

func thousandSeparator2(n int) string {
	if n <= 100 {
		return strconv.Itoa(n)
	}
	buf := bytes.Buffer{}
	c := 0
	for ; n >= 10; n /= 10 {
		buf.WriteString(strconv.Itoa(n % 10))
		c += 1
		if c == 3 {
			buf.WriteByte('.')
			c = 0
		}
	}
	buf.WriteString(strconv.Itoa(n))
	res := buf.Bytes()
	reverse(res)
	return string(res)
}

func reverse(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}

func main() {
	// 3.87 6.67
	fmt.Println(thousandSeparator2(136590))
}
