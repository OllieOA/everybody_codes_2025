package common

type Context struct {
	Quest    int
	Sample bool
}

type Handler func(Context) error

var (
	handlers = make(map[int]Handler)
	names    = make(map[int]string)
)

// Register registers a handler for a quest number
func Register(quest int, name string, h Handler) {
	handlers[quest] = h
	names[quest] = name
}

// GetHandler returns the handler for a quest, or nil if not registered.
func GetHandler(quest int) Handler {
	return handlers[quest]
}

// Name returns a registered name for a quest (or empty string).
func Name(quest int) string {
	return names[quest]
}