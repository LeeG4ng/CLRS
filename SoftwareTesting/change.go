package main

import "fmt"

func main() {
	for {
		loop()
	}
}

func loop() {
	var price, pay int

	fmt.Print("商品金额(大于0不大于100且为整数):")
	_, err := fmt.Scanf("%d", &price)
	if err != nil || price <= 0 || price > 100 {
		fmt.Println("商品金额输入有误")
		return
	}

	fmt.Print("支付金额(大于0不大于100且为整数):")
	_, err = fmt.Scan(&pay)
	if err != nil || pay < price || pay > 100 {
		fmt.Println("支付金额输入有误")
		return
	}

	change := pay - price
	money := []int{50, 10, 5, 2, 1}
	changeMap := make(map[int]int)

	if change < 0 {
		fmt.Println("支付金额小于商品金额")
		return
	}
	for _, unit := range money {
		changeMap[unit] = change / unit
		change = change % unit
	}
	fmt.Print("找零:")
	for unit, count := range changeMap {
		fmt.Printf(" %d*%d元", count, unit)
	}
	fmt.Println()
}
