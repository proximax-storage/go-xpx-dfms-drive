package drive

type Event uint8

const (
	NotStarted Event = iota
	Pending
	InProgress
	Finished
	AnyType
)

func (e Event) Loggable() map[string]interface{} {
	return map[string]interface{}{
		"event": e.String(),
	}
}

func (e Event) String() string {
	return eventNames[e]
}

var eventNames = map[Event]string{
	NotStarted: "NotStarted",
	Pending:    "Pending",
	InProgress: "InProgress",
	Finished:   "Finished",
	AnyType:    "AnyType",
}
