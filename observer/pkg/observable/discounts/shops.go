package discounts

import (
	"fmt"

	"github.com/ivanmatyash/design-patterns/observer/pkg/event"
)

// Ensures Event interface implementation
var _ event.Event = Shops{}

type Shops []string

func (s Shops) GetUpdates() interface{} {
	return s
}

func (s Shops) String() string {
	result := "Shops:\n"
	for k, v := range s {
		result += fmt.Sprintf("%d:\t%s\n", k, v)
	}
	return result
}
