package main

// import (
// 	"fmt"
// )

// func main() {
// 	var p1 string = "-"
// 	var p2 string = "-"
// 	var p3 string = "-"
// 	var p4 string = "-"
// 	var p5 string = "-"

// 	var com string

// 	var count int
// 	for {
// 		fmt.Scan(&com)

// 		if com == "очередь" {
// 			fmt.Printf("1. %s \n", p1)
// 			fmt.Printf("2. %s \n", p2)
// 			fmt.Printf("3. %s \n", p3)
// 			fmt.Printf("4. %s \n", p4)
// 			fmt.Printf("5. %s \n", p5)
// 			continue
// 		}

// 		if com == "количество" {
// 			fmt.Printf("Осталось свободных мест: %d\n", 5-count)
// 			fmt.Printf("Всего человек в очереди: %d\n", count)
// 			continue
// 		}
// 		if com == "конец" {
// 			fmt.Printf("1. %s \n", p1)
// 			fmt.Printf("2. %s \n", p2)
// 			fmt.Printf("3. %s \n", p3)
// 			fmt.Printf("4. %s \n", p4)
// 			fmt.Printf("5. %s \n", p5)
// 			break
// 		}
// 		var number int
// 		fmt.Scan(&number)
// 		if com == string(com) && number == int(number) {
// 			if number < 1 || number > 5 {
// 				fmt.Println(fmt.Sprintf("Запись на место номер %d невозможна: некорректный ввод", number))

// 				continue
// 			}
// 			if count >= 5 {

// 				fmt.Println(fmt.Sprintf("Запись на место номер %d невозможна: очередь переполнена", number))
// 				continue
// 			}
// 			if number == 1 {
// 				if p1 == "-" {
// 					p1 = com
// 					count += 1
// 					continue
// 				} else {

// 					fmt.Println(fmt.Sprintf("Запись на место номер %d невозможна: место уже занято", number))
// 					continue
// 				}
// 			}
// 			if number == 2 {
// 				if p2 == "-" {
// 					p2 = com
// 					count += 1
// 					continue
// 				} else {
// 					fmt.Println(fmt.Sprintf("Запись на место номер %d невозможна: место уже занято", number))
// 					continue
// 				}
// 			}

// 			if number == 3 {
// 				if p3 == "-" {
// 					p3 = com
// 					count += 1
// 					continue
// 				} else {
// 					fmt.Println(fmt.Sprintf("Запись на место номер %d невозможна: место уже занято", number))
// 					continue
// 				}
// 			}

// 			if number == 4 {
// 				if p4 == "-" {
// 					p4 = com
// 					count += 1
// 					continue
// 				} else {
// 					fmt.Println(fmt.Sprintf("Запись на место номер %d невозможна: место уже занято", number))
// 					continue
// 				}
// 			}

// 			if number == 5 {
// 				if p5 == "-" {
// 					p5 = com
// 					count += 1
// 					continue
// 				} else {
// 					fmt.Println(fmt.Sprintf("Запись на место номер %d невозможна: место уже занято", number))
// 					continue
// 				}
// 			}

// 		}

// 	}

// }
