/*
รับ input เป็น array 2 ค่า
รวมข้อมูลที่นำมารวมกันโดยที่ข้อมูลไม่ซ้ำกัน
*/
package main

import "fmt"

func main() {
	array1 := []int{1, 2, 3, 4, 5, 6, 63, 68, 65, 9, 12, 4, 7, 4}
	array2 := []int{7, 4, 5, 9, 12, 8, 67, 5, 34, 69, 65, 85, 89}

	result := sumArray(array1, array2)

	fmt.Println(result)
}

func sumArray(a []int, b []int) []int {
	a = append(a, b...)
	data := []int{a[0]}

	for _, v := range a[1:] {
		isFound := false
		for _, d := range data {
			if d == v {
				isFound = true
			}
		}

		if !isFound {
			data = append(data, v)
		}
	}
	return data
}
