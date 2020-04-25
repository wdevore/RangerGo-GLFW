// Package window defines an Event object send to listeners.
package window

const (
	// MouseEvent of type mouse
	MouseEvent = 0
	// KeyboardEvent of type keyboard
	KeyboardEvent = 1
	// JoystickEvent of type joystick
	JoystickEvent = 2
)

const (
	// PressAction indicates a key or button was pressed
	PressAction = 0

	// ReleaseAction indicates a key or button was released
	ReleaseAction = 1
)

const (
	// KEY0 = 48
	// KEY1 = 49
	// KEY2 = 50
	// KEY3 = 51
	// KEY4 = 52
	// KEY5 = 53
	// KEY6 = 54
	// KEY7 = 55
	// KEY8 = 56
	// KEY9 = 57

	// KeyA -
	KeyA = 65

	// KEYB
	// KEYC
	// KEYD
	// KEYE
	// KEYF
	// KEYG
	// KEYH
	// KEYI
	// KEYJ
	// KEYK
	// KEYL
	// KEYM
	// KEYN
	// KEYO
	// KEYP
	// KEYQ
	// KEYR
	// KEYS
	// KEYT
	// KEYU
	// KEYV
	// KEYW
	// KEYX
	// KEYY
	// KEYZ
)

// Event is an object sent to listeners.
type Event struct {
	// Type indicates the event type.
	Type int

	// Action indicates an action from the window event system.
	Action int

	// bubble indicates if the event should bubble "downwards" or stop
	// at the top most listener.
	bubble bool

	// handled indicates that a listener has recognized an "used" the event
	// bubbling could still occur unless a listener prevents it.
	handled bool

	// --------------------------------------------------------------
	// Keyboard
	// --------------------------------------------------------------
	Key         int
	ScanCode    int
	ModifierKey int
}

// NewEvent constructs a new Event and initializes it.
func NewEvent() *Event {
	e := new(Event)
	e.bubble = true
	e.handled = false
	return e
}

// PreventDefault stop bubbling and indicates that the event was handled.
func (e *Event) PreventDefault() {
	e.bubble = false
	e.handled = true
}

// Reset resets to defaults: bubbling and not handled (aka opposite of PreventDefault)
func (e *Event) Reset() {
	e.bubble = true
	e.handled = false
}
