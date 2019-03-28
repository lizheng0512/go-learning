package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ls1 := new(LockedSlice)
	ls1.info = []string{"你", "好", "啊"}
	go ls1.read(1)
	go ls1.read(2)
	go ls1.read(3)
	go ls1.read(4)
	go ls1.Update(1)
	go ls1.Update(2)

	ls2 := new(LockedSlice)
	ls2.info = []string{"你2", "好2", "啊2"}
	go ls2.read(1)
	go ls2.read(2)
	go ls2.read(3)
	go ls2.read(4)
	go ls2.Update(1)
	go ls2.Update(2)

	time.Sleep(10e9)
}

type LockedSlice struct {
	// 读写锁
	rwLock sync.RWMutex
	info   []string
}

func (ls *LockedSlice) Update(t int) {
	// 写锁
	ls.rwLock.Lock()
	defer ls.rwLock.Unlock()
	fmt.Printf("update... %v sleep %d\n", ls.info, t)
	d := time.Duration(t) * time.Second
	time.Sleep(d)
	ls.info[0] = "给你改啦  哈哈哈!"
	fmt.Println("update done")
}

func (ls *LockedSlice) read(t int) {
	// 读锁, 进入次方法的线程将持有读锁, 其他线程在获取写锁的时候, 必须要等待读锁或者写锁都释放才能获取写锁
	// 写锁权限高于读锁, 有线程等待获取写锁时不会再分配读锁
	ls.rwLock.RLock()
	defer ls.rwLock.RUnlock()
	fmt.Printf("read... %v sleep %d\n", ls.info, t)
	var d = time.Duration(t) * time.Second
	time.Sleep(d)
	str := ""
	for _, v := range ls.info {
		str += v
	}
	fmt.Println(str)
}
