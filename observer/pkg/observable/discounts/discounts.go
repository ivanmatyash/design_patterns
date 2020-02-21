package discounts

import (
	"fmt"

	"github.com/ivanmatyash/design-patterns/observer/pkg/event"
)

// Ensures Event interface implementation
var _ event.Event = List{}

type List map[string]int

func (l List) GetUpdates() interface{} {
	return l
}

func (l List) String() string {
	result := "Discounts:\n"
	for k, v := range l {
		result += fmt.Sprintf("%s:\t%d\n", k, v)
	}
	return result
}
