package observer

import (
	"fmt"

	"github.com/ivanmatyash/design-patterns/observer/pkg/event"
	"github.com/ivanmatyash/design-patterns/observer/pkg/observable/discounts"
)

// Ensures Observer interface implementation
var _ Observer = ClientSMS{}

type ClientSMS struct{}

func (c ClientSMS) Update(e event.Event) {
	updates := e.GetUpdates()
	switch v := updates.(type) {
	case discounts.List:
		fmt.Printf("[SMS] discounts update:\n%s", v)
	case discounts.Shops:
		fmt.Printf("[SMS] shops update:\n%s", v)
	default:
		fmt.Println("Error in Parse UpdateEvent.\n")
		return
	}
	fmt.Println()
}
