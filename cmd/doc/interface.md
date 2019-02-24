# interface 变量

## 什么是interface？

interface 是一组method签名的组合，通过interface来定义对象的一组行为。

1. interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，那么这个对象也就实现了此接口；
2. interface类型可以存实现这个interface的任意类型的对象；
3. interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现。

## 空interface

空interface不包含任何method，空interface对于描述起不到任何的作用，但是空interface在我们需要存储人意类型的睡的时候相当有用，它可以存储任意类型的数据。

**一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，如果一个函数返回interface{},那么也就可以返回任意类型的值。**

## interface 函数参数

interface的变量可以持有任意实现该interface类型的对象，可以通过定义interface参数，让函数接受各种类型的参数。
**例如：fmt.PrintIn 它可以接受人意类型的数据**

实现error接口的对象（即是实现了‘Error() string’对象），适应fmt输出时，会调用Error()方法。

## 反推interface变量存储的类型

Comma-ok断言

Go中有一个语法，可以直接判断是该类型的变量‘value, ok = element.(T)’,value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。

## 嵌入interface 

interface1作为interface2的一个字段，那么interface2就可以调用interface1中的menthod。

## 反射

Go实现反射，反射就是能够检查程序在运行时的状态，一般用到的是reflect包。

```golang
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("type:", v.Type())
fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
fmt.Println("value:", v.Float())
```