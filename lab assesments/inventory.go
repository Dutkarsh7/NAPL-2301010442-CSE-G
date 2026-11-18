package main

import "fmt"

// Struct
type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

var inventory []Item

// Adding my item
func addItem(item Item) {
	inventory = append(inventory, item)
}

// Updating the te
func updateItem(id int, quantity int) {
	for i := 0; i < len(inventory); i++ {
		if inventory[i].ID == id {
			inventory[i].Quantity = quantity
		}
	}
}

// Remove Item
func removeItem(id int) {
	for i := 0; i < len(inventory); i++ {
		if inventory[i].ID == id {
			inventory = append(inventory[:i], inventory[i+1:]...)
			break
		}
	}
}

// Goroutine + Channel
func summary(ch chan string) {
	totalValue := 0.0

	for _, item := range inventory {
		totalValue += float64(item.Quantity) * item.Price
	}

	result := fmt.Sprintf(
		"Items = %d , Total Value = %.2f",
		len(inventory),
		totalValue,
	)

	ch <- result
}

func main() {

	addItem(Item{1, "Pen", 10, 5})
	addItem(Item{2, "Book", 20, 50})
	addItem(Item{3, "Pencil", 15, 3})

	updateItem(1, 12)
	removeItem(3)

	ch := make(chan string)

	go summary(ch)

	fmt.Println("Inventory List")
	for _, item := range inventory {
		fmt.Println(item)
	}

	fmt.Println(<-ch)
}