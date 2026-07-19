package main

// import "fmt"

// type Animal interface {
// 	MakeSound() string
// 	GetName() string
// 	GetInfo() string
// }
// //
// type animal struct {
// 	name    string
// 	species string
// 	age     int
// 	sound   string
// }

// func (a *animal) MakeSound() string {
// 	return a.sound
// }

// func (a *animal) GetName() string {
// 	return a.name
// }

// func (a *animal) GetInfo() string {
// 	return fmt.Sprintf("Имя: %v, Вид: %v, Возраст: %v", a.name, a.species, a.age)
// }

// func NewAnimal(name, species string, age int, sound string) Animal {
// 	return &animal{name, species, age, sound}
// }

// func ZooShow(animals []Animal) {
// 	for _, animal := range animals {
// 		fmt.Println(animal.GetInfo())
// 		fmt.Println(animal.MakeSound())
// 	}
// }

// type ZooKeeper struct {
// }

// func (z ZooKeeper) Feed(animal Animal) {
// 	fmt.Println(fmt.Sprintf("Смотритель зоопарка кормит %v. %v!", animal.GetName(), animal.MakeSound()))
// }
