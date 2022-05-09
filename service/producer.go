package service

import (
	"fmt"

	"concurrency/item"
	"concurrency/repository"
)

func ReadFromDB(repo *repository.Repository) ([]item.Item, error) {
	items, err := repo.ReadItems()
	if err != nil {
		fmt.Println(err)
		return []item.Item{}, err
	}

	return items, nil
}

func Produce(items []item.Item, itemChan chan item.Item) {

	for _, item := range items {
		itemChan <- item
	}

	fmt.Println("all item sent")
}
