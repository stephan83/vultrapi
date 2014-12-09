package errors

type ErrNotFound struct{}

func (_ ErrNotFound) Error() string {
	return "Not found."
}
