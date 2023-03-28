package _map

import "fmt"

func MapHandle() {
	//map 不能使用slice、map、function做key
	//struct做key的时候struct里也不能有以上三种成员变量

	m1 := map[string]string{
		"A": "1",
		"B": "2",
		"C": "3",
		"D": "4",
	}
	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Printf("m1 = %v,len = %d\n", m1, len(m1))
	fmt.Printf("m2 = %v,len = %d\n", m2, len(m2))
	fmt.Printf("m3 = %v,len = %d\n", m3, len(m3))

	fmt.Println("Traversing map")
	for k, v := range m1 {
		//遍历顺序随机
		fmt.Println(k, "=", v)
	}

	fmt.Println("adding value to map")
	fmt.Println(m1)
	m1["E"] = "5"
	fmt.Println(m1)

	fmt.Println("getting value from map")
	v1, exist := m1["A"]
	fmt.Println("m1[\"A\"]=", v1, exist)
	v2, exist := m1["Z"] //不存在的数据会返回空串或0
	fmt.Println("m1[\"Z\"]=", v2, exist)
	if v3, exist := m2["X"]; !exist {
		fmt.Println("m2[\"X\"]=", m2["X"], " key does not have exist!")
	} else {
		fmt.Println("m1[\"X\"]=", v3)
	}

	fmt.Println("delete element from map")
	fmt.Println(m1)
	delete(m1, "A")
	fmt.Println(m1)
	delete(m1, "Z")
	fmt.Println(m1)

}
