package main

import (
	"fmt"
)

func main() {
	appleProduct := Product{Name: "Apple", Price: 0.5, Stock: 100}
	bananaProduct := Product{Name: "Banana", Price: 0.25, Stock: 75}
	orangeProduct := Product{Name: "Orange", Price: 0.75, Stock: 50}

	store := Store{
		Name: "John's emporium",
		inventory: map[string]Product{
			"apple":  appleProduct,
			"banana": bananaProduct,
			"orange": orangeProduct,
		},
	}

	// Print the inventory
	store.PrintInventory()

	// Sell some products
	store.SellProduct("apple", 25)
	store.SellProduct("banana", 15)
	store.SellProduct("orange", 10)

	// Print the updated inventory
	store.PrintInventory()
	value := store.StockValue()
	fmt.Println(value)
}

type Product struct {
	Name  string
	Price float64
	Stock int
}

type Store struct {
	Name      string
	inventory map[string]Product
}

func (store Store) HasItem(itemName string) bool {
	if store.inventory[itemName].Stock >= 0 {
		return true
	} else {
		return false
	}
}

func (store Store) PrintInventory() {
	fmt.Println(store.Name, "inventory:")
	for name, product := range store.inventory {
		fmt.Printf("%s: £%.2f, %d in stock\n", name, product.Price, product.Stock)
	}
	fmt.Println()
}

func (s Store) SellProduct(name string, quantity int) (int, error) {
	p, ok := s.inventory[name]
	if !ok {
		return 0, fmt.Errorf("product %v not found", name)
	}

	if p.Stock > quantity {
		return p.Stock, fmt.Errorf("not enough stock. %v %v requested but only %v in stock", quantity, name, p.Stock)
	}

	p.Stock -= quantity

	return p.Stock, nil
}

func (store Store) StockValue() float64 {
	total := 0.0
	for _, price := range store.inventory {
		total += price.Price * float64(price.Stock)
	}
	return total
}
