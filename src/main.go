package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
)

func main() {
	file, err := os.Open("./../vcard.json")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 1000) // 数组切片
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	fmt.Printf("read %d bytes:%q\n", count, data[:count])

	fmt.Println("*****************数组 & 切片********************")
	var arr1 = new([5]int) //  生成的空白指针类型
	var arr2 [5]int        //
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	arr2 = *arr1
	arr2[0] = 1
	arr2[1] = 5
	arr2[2] = 3

	fmt.Println(*arr1)
	fmt.Println(arr2)
	fmt.Println("*************************************")

	array := []int{1, 2, 3, 4, 5, 8, 9, 10}
	x := 0
	x = sum(array[:])
	fmt.Println(x)
	fmt.Println(array)
	fmt.Println("*************************************")

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b)
	fmt.Println(b[1:4])
	fmt.Println(b[:2])
	fmt.Println(b[2:])
	fmt.Println(b[:])
	fmt.Println("*************************************")

	slice1 := make([]int, 10)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * (i + 1)
	}
	fmt.Println(slice1)
	fmt.Println("*************************************")
	str := "aa"
	num := 2
	num1 := 2.0
	// 判断类型
	fmt.Println("type:", reflect.TypeOf(str))  // =>string
	fmt.Println("type:", reflect.TypeOf(num))  // =>int
	fmt.Println("type:", reflect.TypeOf(num1)) // => float64
	fmt.Println("type:", reflect.TypeOf(arr1)) // =>*[5]int
	fmt.Println("type:", reflect.TypeOf(arr2))
	fmt.Println("type:", reflect.TypeOf(array))
	fmt.Println("**************************************")
	var ar = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7]
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)
	fmt.Println("**************************************")

	slfrom := []int{1, 2, 3}
	slto := make([]int, 10)

	n := copy(slto, slfrom)
	fmt.Println(slto)
	fmt.Println(n)
	fmt.Println("**************************************")
	fmt.Println("ar type:", reflect.TypeOf(ar))
	fmt.Println("sl_from type:", reflect.TypeOf(slfrom))
	str1 := reflect.TypeOf(slfrom)
	if reflect.TypeOf(ar) == str1 {
		fmt.Println(reflect.TypeOf(str1))
	}
	fmt.Println("**************************************")
	str3 := "a1\u00ff\u754c"
	c := []byte(str3)
	fmt.Println(c)
	fmt.Println(str3)
	str4 := "hello world"
	fmt.Println(str4)
	str5 := str4[2:4]
	fmt.Println(str5)
	fmt.Println("***************map结构***********************")
	var items map[string]int
	var maplist map[string]int
	fmt.Println(maplist)
	items = map[string]int{"one": 1, "two": 2}
	maplist = items
	fmt.Println(reflect.TypeOf(items))
	maplist["one"] = 656
	maplist["oness"] = 656
	fmt.Println(items)
	fmt.Println(maplist)

	fmt.Println("**************************************")

	var maplisted map[string]int
	mapcreated := make(map[string]int)
	maplisted = mapcreated // maplisted 是 mapcreated的引用
	mapcreated["key"] = 000
	maplisted["ddd"] = 22
	fmt.Println(mapcreated)
	fmt.Println(maplisted)

	fmt.Println("**************************************")
	maplistedd := map[string]int{}
	maplistedd["key"] = 2
	fmt.Println(maplistedd)
	key, isOk := maplistedd["key"] // 判断key是否存在
	fmt.Println(key)               // =>2
	fmt.Println(isOk)              // =>true
	key1, isOk1 := maplistedd["key1"]
	fmt.Println(key1)  // =>0
	fmt.Println(isOk1) // => false

	fmt.Println("**************map 类型的切片************************")
	mapitems := make([]map[int]int, 5)
	for i, val := range mapitems {
		fmt.Println(val)
		mapitems[i] = make(map[int]int, 1)
		mapitems[i][2*i] = 2 * i
	}
	fmt.Println(mapitems)
	fmt.Println(mapitems[0][0])
	barval := map[string]map[string]int{}
	barval["keay"] = maplistedd
	fmt.Println(reflect.TypeOf(barval))
	fmt.Println(barval)
	for key, val := range barval {
		fmt.Println(key)
		fmt.Println(val)
		for keyItem, valItem := range val {
			fmt.Println(keyItem)
			fmt.Println(valItem)
		}
	}

	fmt.Println("**************************************")
	var (
		barVal = map[string]int{
			"alpha": 34, "bravo": 56, "charlie": 23,
			"delta": 87, "echo": 56, "foxtrot": 12,
			"golf": 34, "hotel": 16, "indio": 87,
			"juliet": 65, "kili": 43, "lima": 98, "lissma": 98}
	)

	for key, val := range barVal {
		fmt.Println(key)
		fmt.Println(val)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k := range barVal {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)
	for _, k := range keys {
		fmt.Println(barVal[k])
	}
	fmt.Println("**************** map 键值对换**********************")
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		if _, ok := invMap[v]; !ok {
			invMap[v] = k
		}
	}
	fmt.Println(invMap)
	fmt.Println("**************** 标准库**********************")
	fmt.Println("**************** 正则 **********************")
	// 字符串
	// searchIn := "4458"
	fmt.Println("**************** 结构体 **********************")

	type innerS struct {
		in1 int
		in2 int
	}
	type outerS struct {
		b   int
		c   float32
		in1 int // =>{6 3.3 5 60 {0 10}}
		int
		innerS
	}

	outer := new(outerS) // 使用new 创建 指针类型
	outer.b = 6
	outer.c = 3.3
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10
	fmt.Println(*outer)
	outer2 := outerS{6, 8.5, 10, 60, innerS{40, 20}} // 使用结构体字面量 引用类型
	fmt.Println(outer2)
	fmt.Println(outer2.in1)
}
func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
