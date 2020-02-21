package observable

import "github.com/ivanmatyash/design-patterns/observer/pkg/observer"

type Observable interface {
	AddDiscountObserver(observer observer.Observer)
	RemoveDiscountObserver(observer observer.Observer)
	AddShopObserver(observer observer.Observer)
	RemoveShopObserver(observer observer.Observer)
	NotifyDiscountObservers()
	NotifyShopObservers()
}
