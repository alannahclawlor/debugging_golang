package main

import (
	"testing"
)

func TestHasItem(t *testing.T) {
	sweet := Product{Name: "Sweet", Price: 0.1, Stock: 1}
	grape := Product{Name: "Grape", Price: 0.2, Stock: 1}

	store := Store{
		Name: "John's emporium",
		inventory: map[string]Product{
			"sweet": sweet,
			"grape": grape,
		},
	}

	result, expected := store.HasItem("grape"), true
	if result != expected {
		t.Errorf("Result is %v when %v is expected", result, expected)
	}
}

func TestDoesNotHaveItem(t *testing.T) {
	sweet := Product{Name: "Sweet", Price: 0.1, Stock: 1}
	grape := Product{Name: "Grape", Price: 0.2, Stock: 1}

	store := Store{
		Name: "John's emporium",
		inventory: map[string]Product{
			"sweet": sweet,
			"grape": grape,
		},
	}

	result, expected := store.HasItem("orange"), false
	if result != expected {
		t.Errorf("Result is %v when %v is expected", result, expected)
	}
}

func TestSellProduct(t *testing.T) {
	sweet := Product{Name: "Sweet", Price: 0.1, Stock: 1}
	grape := Product{Name: "Grape", Price: 0.2, Stock: 1}

	store := Store{
		Name: "John's emporium",
		inventory: map[string]Product{
			"sweet": sweet,
			"grape": grape,
		},
	}

	result, expected := store.HasItem("sweet"), true
	if result != expected {
		t.Errorf("Result is %v when %v is expected", result, expected)
	}
	store.SellProduct("sweet", 1)
	result, expected = store.HasItem("sweet"), false
	if result != expected {
		t.Errorf("Result is %v when %v is expected", result, expected)
	}
}

func TestStockValue(t *testing.T) {
	sweet := Product{Name: "Sweet", Price: 0.1, Stock: 1}
	grape := Product{Name: "Grape", Price: 0.2, Stock: 1}

	store := Store{
		Name: "John's emporium",
		inventory: map[string]Product{
			"sweet": sweet,
			"grape": grape,
		},
	}

	result, expected := store.StockValue(), 0.3
	if result != expected {
		t.Errorf("Result is %v when %v is expected", result, expected)
	}
}
