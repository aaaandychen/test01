package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// // 定义缓冲区大小
// const BUFLEN = 5
//
// // 全局位置定义全局变量
// var cond *sync.Cond = sync.NewCond(&sync.Mutex{})
//
// // 生产者
//
//	func producer(ch chan<- int) {
//		for {
//			cond.L.Lock()           // 给条件变量对应的互斥锁加锁
//			for len(ch) == BUFLEN { // 缓冲区满，则等待消费者消费，这里不能是if
//				cond.Wait()
//			}
//			ch <- rand.Intn(1000) // 写入缓冲区一个随机数
//			cond.L.Unlock()       // 生产结束，解锁互斥锁
//			cond.Signal()         // 一旦生产后，就唤醒其他被阻塞的消费者
//			time.Sleep(time.Second * 2)
//		}
//	}
//
// // 消费者
//
//	func consumer(ch <-chan int) {
//		for {
//			cond.L.Lock()      // 全局条件变量加锁
//			for len(ch) == 0 { // 如果缓冲区为空，则等待生产者生产，，这里不能是if
//				cond.Wait() // 挂起当前协程，等待条件变量满足，唤醒生产者
//			}
//			fmt.Println("Receive:", <-ch)
//			cond.L.Unlock()
//			cond.Signal()
//			time.Sleep(time.Second * 1)
//		}
//	}
var db *sql.DB

func initMysql() (err error) {

	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/demo"
	//Open  函数只是校验   dsn  的查数是否正确，  并不会连接数据库
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	fmt.Println("连接成功")

	//尝试与数据库进行连接
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	return
}

type User struct {
	id   int
	age  int
	name string
}

// 单行查询
func queryRowDemo() {
	sqlStr := "select  id,  age, name  from student  where id = ?"
	var u User

	//执行查询语句, QueryRow执行完之后一定要调用  Scan 方法（会自动关闭  连接）
	row := db.QueryRow(sqlStr, 2)
	//将数据取出赋值到  user  结构体中的变量中
	err := row.Scan(&u.id, &u.age, &u.name)
	if err != nil {
		fmt.Println("scan  filed  fail", err)
		return
	}
	fmt.Printf("id: %d,   age: %d ,  name:%s", u.id, u.age, u.name)
}

func main() {
	if err := initMysql(); err != nil {
		fmt.Println("数据库连接失败", err)
	}

	// 注意这个  defer 关闭的  需要  拿出来
	defer db.Close()
	fmt.Println("数据库连接成功....")
	queryRowDemo()
}
