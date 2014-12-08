package errors

type ErrUnknownCommand struct{}

func (_ ErrUnknownCommand) Error() string {
	return "Unknown command."
}
