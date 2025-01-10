package main

func AddUpper(n int) int {
	res := 0
	for i := 1; i < n+1; i++ {
		res += i
	}
	return res
}

// func addUpper(n int) int {
// 	start_time := time.Now()
// 	fmt.Println(start_time)
// 	var total_number int
// 	for i := 0; i < n; i++ {
// 		total_number += i
// 	}
// 	fmt.Println("total_number =", total_number)
// 	end_time := time.Now()

// 	fmt.Printf("cost time: %s\n", end_time.Sub(start_time))
// 	return total_number
// }
