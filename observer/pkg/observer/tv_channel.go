package observer

import (
	"fmt"

	"github.com/ivanmatyash/design-patterns/observer/pkg/event"
	"github.com/ivanmatyash/design-patterns/observer/pkg/observable/discounts"
)

// Ensures Observer interface implementation
var _ Observer = TVChannel{}

type TVChannel struct{}

func (c TVChannel) Update(e event.Event) {
	updates := e.GetUpdates()
	switch v := updates.(type) {
	case discounts.List:
		fmt.Printf("[TV] discounts update:\n%s", v)
	case discounts.Shops:
		fmt.Printf("[TV] shops update:\n%s", v)
	default:
		fmt.Println("Error in Parse UpdateEvent.\n")
		return
	}
	fmt.Println()
}
