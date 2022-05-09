package service

import (
	"concurrency/config"
	"concurrency/item"
	"concurrency/repository"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

func Init() error {
	config := config.LoadConfig()

	db, cleanup, err := repository.Init(config)
	if err != nil {
		fmt.Println("error connecting to db")
		return errors.Wrap(err, "error connecting db")
	}
	defer cleanup()

	repo := repository.New(db)
	items, err := ReadFromDB(repo)
	if err != nil {
		fmt.Println("error reading items from db")
		return errors.Wrap(err, "error reading items from db")
	}

	itemChan := make(chan item.Item)
	go Produce(items, itemChan)

	for i := 0; i < ConsumerCount; i++ {
		go Consume(itemChan)
	}

	time.Sleep(time.Second * WaitTime)
	return nil
}
