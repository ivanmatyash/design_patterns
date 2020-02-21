package observable

import (
	"reflect"
	"testing"

	"github.com/ivanmatyash/design-patterns/observer/pkg/event"

	"github.com/ivanmatyash/design-patterns/observer/pkg/observer"
)

type fakeObserver struct{}

func (fakeObserver) Update(event.Event) {}

func TestDiscountCenter_AddRemoveDiscountObserver(t *testing.T) {
	dCenter := NewDiscountCenter()
	checkLen(t, 0, len(dCenter.discountObservers))

	o1 := fakeObserver{}
	dCenter.AddDiscountObserver(o1)
	checkLen(t, 1, len(dCenter.discountObservers))
	reflect.DeepEqual(dCenter.discountObservers, []observer.Observer{o1})

	dCenter.RemoveDiscountObserver(o1)
	checkLen(t, 0, len(dCenter.discountObservers))
	dCenter.RemoveDiscountObserver(o1)
	checkLen(t, 0, len(dCenter.discountObservers))

	o2 := fakeObserver{}
	dCenter.AddDiscountObserver(o1)
	dCenter.AddDiscountObserver(o2)
	checkLen(t, 2, len(dCenter.discountObservers))
	reflect.DeepEqual(dCenter.discountObservers, []observer.Observer{o1, o2})

	dCenter.RemoveDiscountObserver(o1)
	checkLen(t, 1, len(dCenter.discountObservers))
	reflect.DeepEqual(dCenter.discountObservers, []observer.Observer{o2})

	dCenter.RemoveDiscountObserver(o2)
	checkLen(t, 0, len(dCenter.discountObservers))
}

func TestDiscountCenter_AddRemoveShopObserver(t *testing.T) {
	dCenter := NewDiscountCenter()
	checkLen(t, 0, len(dCenter.shopObservers))

	o1 := fakeObserver{}
	dCenter.AddShopObserver(o1)
	checkLen(t, 1, len(dCenter.shopObservers))
	reflect.DeepEqual(dCenter.shopObservers, []observer.Observer{o1})

	dCenter.RemoveShopObserver(o1)
	checkLen(t, 0, len(dCenter.shopObservers))
	dCenter.RemoveShopObserver(o1)
	checkLen(t, 0, len(dCenter.shopObservers))

	o2 := fakeObserver{}
	dCenter.AddShopObserver(o1)
	dCenter.AddShopObserver(o2)
	checkLen(t, 2, len(dCenter.shopObservers))
	reflect.DeepEqual(dCenter.shopObservers, []observer.Observer{o1, o2})

	dCenter.RemoveShopObserver(o1)
	checkLen(t, 1, len(dCenter.shopObservers))
	reflect.DeepEqual(dCenter.shopObservers, []observer.Observer{o2})

	dCenter.RemoveShopObserver(o2)
	checkLen(t, 0, len(dCenter.shopObservers))
}

func TestDiscountCenter_AddRemoveDiscount(t *testing.T) {
	dCenter := NewDiscountCenter()

	dCenter.AddDiscount("ball", 50)
	checkLen(t, 1, len(dCenter.discounts))

	dCenter.RemoveDiscount("ball")
	checkLen(t, 0, len(dCenter.discounts))
}

func checkLen(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Expected len %d != actual len %d.", expected, actual)
	}
}
