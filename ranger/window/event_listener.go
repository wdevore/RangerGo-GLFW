// Package window defines an interface for eventlistners
package window

// EventListener is an interface for objects that want notification from
// the window such as keyboard and mouse.
type EventListener interface {
	Receive(e *Event)
}
