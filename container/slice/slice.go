package main

import (
	"fmt"
	"golang_study/container/map"
)

func updateSlice(s []int) {
	s[0] = 100
}

func printSlice(s []int) {
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
}

func SliceHandle() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] : ", arr[2:6])
	fmt.Println("arr[:6] : ", arr[:6])
	s1 := arr[2:6]
	s2 := arr[2:]
	fmt.Println("arr[2:] : ", s1)
	fmt.Println("arr[:] : ", s2)

	fmt.Println("after updateSlice ")
	updateSlice(s1)
	fmt.Println("s1 : ", s1)
	fmt.Println("arr : ", arr)
	fmt.Println("s2 : ", s2)

	fmt.Println("reSlice")
	fmt.Println(s1)
	s1 = s1[:5]
	fmt.Println(s1)
	s1 = s1[2:]
	fmt.Println(s1)

	fmt.Println("extend slice")
	arr[2] = 2
	s1 = arr[2:6] //2,3,4,5
	s2 = s1[3:5]  //	  5,(6),(7)
	//slice 底层其实还是arr，多了三个参数，ptr 指向切片开头元素、len 切片的长度、cap 从ptr开始到arr结束的长度
	//所以s2取s1的第4位，其实取的是s1切片后的一位，仍然可见
	fmt.Println("s1 : ", s1)
	fmt.Println("s2 : ", s2)
	//slice取值不能大于等于len，切片扩展不能大于cap
	fmt.Printf("s2=%v,len=%d,cap=%d\n", s2, len(s2), cap(s2))
	//fmt.Println("s2[2] : ", s2[2]) wrong
	//fmt.Println("s2[1:4] : ", s2[1:4]) wrong
	fmt.Println("s2[1:3] : ", s2[1:3])
}

func SliceAppend() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[3:7]
	//如果append没有超过arr的len则会覆盖
	s2 := append(s1, 100)
	//如果append超过arr的len则会生成新的更大容量的数组，此时s3，s4的切片不再是arr的view了
	s3 := append(s2, 1000)
	s4 := append(s3, 10000)
	fmt.Printf("s1=%v\ns2=%v\ns3=%v\ns4=%v\narr=%v\n", s1, s2, s3, s4, arr)
}

func SliceOperation() {
	var s []int //元素全为0
	printSlice(s)

	for i := 0; i < 100; i++ {
		//len=cap时候，cap^2扩容
		s = append(s, 2*i+1)
		printSlice(s)
	}
	fmt.Println(s)

	s1 := []int{1, 2, 3, 4}

	//新建len=16，值都为0的切片
	s2 := make([]int, 16)
	printSlice(s2)
	//新建len=10，cap=32，值都为0的切片
	s3 := make([]int, 10, 32)
	printSlice(s3)

	fmt.Println("Copy(target,source) slice")
	copy(s2, s1)
	fmt.Println(s2)
	//超过不copy
	s11 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	copy(s2, s11)
	fmt.Println(s2)

	fmt.Println("Deleting element from slice")
	//删掉第[1,2)个数据
	s2 = append(s2[:1], s2[2:]...)
	fmt.Println(s2)

	fmt.Println("Popping from front")
	//拿去第一个数据
	front := s2[0]
	//删除第一个数据
	s2 = s2[1:]
	fmt.Println(front, "\n", s2)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail, "\n", s2)
	printSlice(s2)
}

func main() {
	fmt.Println("===SliceAppend===")
	SliceAppend()
	fmt.Println("===SliceOperation===")
	SliceOperation()
	fmt.Println("===MapHandle===")
	_map.MapHandle()
}
