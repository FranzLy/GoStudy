/*
 * @Description: hello world
 * @Author: liyu
 * @Version:
 * @Date: 2023-12-23 17:23:55
 * @LastEditTime: 2023-12-23 18:10:20
 */
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func testArray() {
	/*测试数组*/
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)  实际上s2 := a[4:5]也就是{5},[4:5]是在[3:4]基础上+上面的s在a中的偏移1得来的，
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
}

func testSlice() {
	/*测试切片*/
	var b = make([]string, 5, 10)
	//fmt.Println(b)
	//fmt.Printf("len(b):%v, cap(b):%v\n", len(b), cap(b))
	for i := 0; i < 10; i++ {
		b = append(b, fmt.Sprintf("%v", i))
		//fmt.Println(b)
		//fmt.Printf("len(b):%v, cap(b):%v\n", len(b), cap(b))
	}
	fmt.Println(b)
	//fmt.Printf("len(b):%v, cap(b):%v\n", len(b), cap(b))

	var c = [...]int{3, 7, 8, 9, 1} // c是数组
	slice1 := c[:]
	sort.Ints(slice1) //sort包只能对切片排序
	fmt.Println(slice1)

	var d = [...]int64{3, 7, 8, 9, 1} // c是数组
	slice2 := d[:]
	sort.Slice(slice2, func(i int, j int) bool {
		return int64(d[i]) < int64(d[j])
	}) //sort包只能对切片排序
	fmt.Println(slice2)
}

func testMap() {
	/*测试map*/
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	//遍历key val
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	//也可以只遍历key
	for key := range scoreMap {
		fmt.Println(key)
	}

	//按key删除元素
	scoreMap["娜扎"] = 60
	delete(scoreMap, "张三")
	//遍历key val
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	//根据排序后的Key值遍历
	rand.Seed(time.Now().UnixNano())
	var anyMap = make(map[string]int, 200)
	var keys = make([]string, 0, 100)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		anyMap[key] = value

		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("student:%v : score: %v\n", k, anyMap[k])
	}

}

/*
元素为map类型的切片
值为切片类型的map
*/
func testSliceMap() {
	// 1
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	// 2
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init.")
	key := "China"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0)
		value = append(value, "Beijing")
		value = append(value, "Shanghai")
		value = append(value, "Nanjing")
		fmt.Println(value)
	}
	sliceMap[key] = value
	fmt.Println(sliceMap)

	// 3写一个程序，统计一个字符串中每个单词出现的次数。比如：“how do you do"中how=1 do=2 you=1。
	str := "how do you do"
	words := strings.Split(str, " ")
	wordsMap := make(map[string]int, 0)
	for _, w := range words {
		if _, ok := wordsMap[w]; !ok {
			wordsMap[w] = 1
		} else {
			wordsMap[w]++
		}
	}
	fmt.Println(wordsMap)

	//观察下面代码，写出最终的打印结果
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	//c := make([]int, 3)
	//copy(c, s)
	fmt.Printf("%+v,  %p, %p, %p, cap(s)=%v\n", s, &s[0], &s[1], &s[2], cap(s))
	m["q1mi"] = s //此时m["q1mi"]的结果是[1 2 3]
	fmt.Printf("%+v, %p, %p, %p\n", m["q1mi"], &m["q1mi"][0], &m["q1mi"][1], &m["q1mi"][2])
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v, %p, %p, cap(s)=%v\n", s, &s[0], &s[1], cap(s))                         //s=[1 3]
	fmt.Printf("%+v, %p, %p, %p\n", m["q1mi"], &m["q1mi"][0], &m["q1mi"][1], &m["q1mi"][2]) //m["q1mi"]=[1 3 3]
}

func test() {
	fmt.Println("hello world!")
	fmt.Println("test!")
	fmt.Println("hello world!")
	fmt.Println("hello world!")
	println("rger")

	testArray()

	testSlice()

	testMap()

	testSliceMap()
}

func main() {
	test()
}
