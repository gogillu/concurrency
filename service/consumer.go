package service

import (
	"concurrency/item"
	"fmt"
)

func Consume(itemChan chan item.Item) {
	for item := range itemChan {
		itemInvoice := item.GetInvoice()
		fmt.Println("received <- : ", itemInvoice.GetItem().GetName(), itemInvoice.GetTax(), itemInvoice.GetEffectivePrice())
	}
}
