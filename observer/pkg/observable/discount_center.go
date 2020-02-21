package observable

import (
	"github.com/ivanmatyash/design-patterns/observer/pkg/observable/discounts"
	"github.com/ivanmatyash/design-patterns/observer/pkg/observer"
)

// Ensures Observable interface implementation
var _ Observable = &DiscountCenter{}

// Observable subject
type DiscountCenter struct {
	discountObservers []observer.Observer
	shopObservers     []observer.Observer

	discounts discounts.List
	shops     discounts.Shops
}

func NewDiscountCenter() *DiscountCenter {
	d := &DiscountCenter{}
	d.discounts = make(map[string]int)
	d.shops = make([]string, 0)
	return d
}

func (d *DiscountCenter) AddDiscountObserver(o observer.Observer) {
	d.discountObservers = append(d.discountObservers, o)
}

func (d *DiscountCenter) RemoveDiscountObserver(o observer.Observer) {
	for i, obj := range d.discountObservers {
		if obj == o {
			d.discountObservers[i] = d.discountObservers[len(d.discountObservers)-1]
			d.discountObservers = d.discountObservers[:len(d.discountObservers)-1]
			return
		}
	}
}

func (d *DiscountCenter) AddShopObserver(o observer.Observer) {
	d.shopObservers = append(d.shopObservers, o)
}

func (d *DiscountCenter) RemoveShopObserver(o observer.Observer) {
	for i, obj := range d.shopObservers {
		if obj == o {
			d.shopObservers[i] = d.shopObservers[len(d.shopObservers)-1]
			d.shopObservers = d.shopObservers[:len(d.shopObservers)-1]
			return
		}
	}
}

func (d *DiscountCenter) AddDiscount(name string, discount int) {
	d.discounts[name] = discount
	d.NotifyDiscountObservers()
}

func (d *DiscountCenter) RemoveDiscount(name string) {
	delete(d.discounts, name)
	d.NotifyDiscountObservers()
}

func (d *DiscountCenter) AddShop(name string) {
	d.shops = append(d.shops, name)
	d.NotifyShopObservers()
}

func (d *DiscountCenter) NotifyDiscountObservers() {
	for _, o := range d.discountObservers {
		o.Update(d.discounts)
	}
}

func (d *DiscountCenter) NotifyShopObservers() {
	for _, o := range d.shopObservers {
		o.Update(d.shops)
	}
}
