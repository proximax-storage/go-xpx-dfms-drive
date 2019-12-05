package drive

type Event uint8

const (
	NotStarted Event = iota
	Pending
	InProgress
	Finished
	AnyType
)
