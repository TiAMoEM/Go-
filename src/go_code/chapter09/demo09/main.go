package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello杯"
	fmt.Println("str len=", len(str))

	//字符串遍历
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("字符=%c, \n", str2[i])
	}

	//字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result =", n)
	}

	//整数转字符串
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v, str=%T\n", str, str)

	//字符串转[]byte
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	//[]byte转字符串
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

	//10进制转换为其他进制 返回是一个字符串
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是： %v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的16进制是： %v\n", str)

	//查找子串是否在指定的字符串中 返回ture / false
	b := strings.Contains("seafood", "oo")
	fmt.Printf("b = %v\n", b)

	//不区分大小写的字符串比较
	a := strings.EqualFold("abc", "Abc")
	fmt.Printf("a = %v\n", a)
	fmt.Println("abc" == "Abc")

	//返回子串在字符串第一次出现的index值，如果没有返回-1
	index := strings.Index("NLT_abcabcabc", "abc")
	fmt.Printf("index = %v\n", index)
	index = strings.Index("NLT_abcabcabc", "hello")

	//返回子串在字符串最后一次出现的index，如果没有返回-1
	index1 := strings.LastIndex("go golang", "go")
	fmt.Printf("index1 =%v\n", index1)

	//将指定的子串替换成另外一个子串,最后的参数n为替换的数量，-1指全部替换
	str3 := "go go hello"
	str = strings.Replace(str3, "go", "world", -1)
	fmt.Printf("str = %v str3 = %v\n", str, str3)

	//按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组
	strArr := strings.Split("hello,world,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v] = %v\n", i, strArr[i])
	}
	fmt.Printf("strArr = %v\n", strArr)

	//将字符串的字母进行大小写的转换
	str = "goLang Hello"
	str = strings.ToLower(str)
	fmt.Printf("str = %v\n", str)
	str = strings.ToUpper(str)
	fmt.Printf("str = %v\n", str)

	//将字符串左右两边的空格去掉
	str = strings.TrimSpace("  holy shit !         ")
	fmt.Printf("str = %q\n", str)

	//将字符串左右两边指定的字符去掉
	str = strings.Trim("!  he!llo! ", " !")
	fmt.Printf("str = %q\n", str)

	//去左字符和去右字符
	str = strings.TrimLeft("!  he!llo! ", " !")
	str = strings.TrimRight("!  he!llo! ", " !")

	//判断字符串是否以指定的字符串开头
	//判断字符串是否以指定的字符串结束
	b = strings.HasPrefix("http://www.baidu.com", "http")
	fmt.Printf("b = %v\n", b)
	b = strings.HasSuffix("http://www.baidu.com", "cn")
	fmt.Printf("b = %v\n", b)

}
