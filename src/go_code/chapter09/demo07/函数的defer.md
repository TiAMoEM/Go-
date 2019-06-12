1) 当go执行到一个defer时，不会立即执行defer后的语句，而是将defer后的语句压入到一个栈中
2) 当函数执行完毕后，再从defer栈中，依次从栈顶取出语句执行，先入后出原则


defer 的最佳实践:
例如打开文件或者连接上数据库时，可以及时释放资源

```
func test() {
    file = openfile(xxxx)
    defer file.close()
}

func test() {
    connect = openDatebse()
    defer connect.close()
}
```
在defer后仍然可以使用创建的资源
