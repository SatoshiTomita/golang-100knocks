package datatype

import "fmt"

func main() {
	list := []string{"a", "b", "c"}
	fmt.Println("list is", list)

	list = append(list, "d")
	fmt.Println("list is", list)

	list[0] = "A"
	fmt.Println("list is", list)

	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println("m is", m)
	m["orange"] = 300
	fmt.Println("m is", m) // m is map[apple:100 banana:200 orange:300]

	m2 := make(map[string]int)
	fmt.Println("m2 is", m2)
	m2["pc"] = 5000
	fmt.Println("m2 is", m2) // m2 is map[pc:5000]

	m3 := make(map[int][]string)
	m3[1] = []string{"A", "B"}
	fmt.Println("m3 is", m3) // m3 is map[1:[A B]]
}
