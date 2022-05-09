package item

import (
	"concurrency/item/enum"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Item struct {
	Name     string        `gorm:"column:name"`
	Price    int           `gorm:"column:price"`
	Quantity int           `gorm:"column:quantity"`
	Typ      enum.ItemType `gorm:"column:Type;type:enum('raw','manufactured','imported');"`
}

type Invoice struct {
	itm            Item
	tax            float64
	effectivePrice float64
}

func (inv Invoice) GetItem() Item {
	return inv.itm
}

func (inv Invoice) GetTax() float64 {
	return inv.tax
}

func (inv Invoice) GetEffectivePrice() float64 {
	return inv.effectivePrice
}

func (itm Item) GetName() string {
	return itm.Name
}

func (itm Item) GetPrice() int {
	return itm.Price
}

func (itm Item) GetQuantity() int {
	return itm.Quantity
}

func (itm Item) GetType() enum.ItemType {
	return itm.Typ
}

func New(name string, price int, quantity int, Type string) (Item, error) {

	itmType, err := enum.ItemTypeString(Type)
	if err != nil {
		return Item{}, fmt.Errorf("error : provided item type value is not valid")
	}

	itm := Item{}
	itm.Name = name
	itm.Price = price
	itm.Quantity = quantity
	itm.Typ = itmType

	if err := itm.validate(); err != nil {
		return itm, fmt.Errorf("error : item validation failed")
	}

	return itm, nil
}

func (itm Item) validate() error {
	return validation.ValidateStruct(&itm,
		validation.Field(&itm.Price, validation.By(checkNegativeValue)),
		validation.Field(&itm.Quantity, validation.By(checkNegativeValue)),
	)
}

func checkNegativeValue(val interface{}) error {
	switch x := val.(type) {
	case int:
		if x < 0 {
			return fmt.Errorf("error : %d given value is negative", val)
		}
	}
	return nil
}

func (itm Item) GetInvoice() Invoice {
	return Invoice{
		itm:            itm,
		tax:            itm.CalculateTax(),
		effectivePrice: itm.CalculateTax() + float64(itm.GetPrice()*itm.GetQuantity()),
	}
}

func (itm Item) CalculateTax() float64 {

	var tax float64

	switch itm.Typ {
	case enum.Raw:
		tax = float64(itm.Price) * float64(itm.Quantity) * RawItmTaxRate

	case enum.Manufactured:
		tax = float64(itm.Price) * float64(itm.Quantity) * (ManufacturedItmTaxRate + ManufacturedItmTaxRate*(1+ManufacturedItmExtraTaxRate))

	case enum.Imported:
		tax = float64(itm.Price)*float64(itm.Quantity)*ImportDutyRate + getSurcharge(float64(itm.Price*itm.Quantity)*(1.0+ImportDutyRate))
	}

	return tax
}

func getSurcharge(price float64) float64 {
	if price <= SurchargeCap1MaxAmt {
		return SurchargeAmt1
	} else if price <= SurchargeCap2MaxAmt {
		return SurchargeAmt2
	} else {
		return price * SurchargeRate3
	}
}
