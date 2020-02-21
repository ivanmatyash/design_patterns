package main

import (
	"github.com/ivanmatyash/design-patterns/observer/pkg/observable"
	"github.com/ivanmatyash/design-patterns/observer/pkg/observer"
)

func main() {
	dc := observable.NewDiscountCenter()
	clientSMS := observer.ClientSMS{}
	clientTV := observer.TVChannel{}

	dc.AddDiscountObserver(clientSMS)

	dc.AddDiscount("ball", 50) // SMS client will receive a notification about new discounts.
	dc.AddShop("Shop")         // There are no subscribers to these notifications, no one will receive a message.

	dc.AddShopObserver(clientTV)
	dc.AddShop("Shop2") // Only TV client will receive a notification about new shop.

	dc.AddShopObserver(clientSMS) // SMS client has subscribed to notifications of adding shops.
	dc.AddShop("Shop3")           // Both TV and SMS clients will receive a notification about new shop.
}
