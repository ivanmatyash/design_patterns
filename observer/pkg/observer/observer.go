package observer

import "github.com/ivanmatyash/design-patterns/observer/pkg/event"

type Observer interface {
	Update(event event.Event)
}
