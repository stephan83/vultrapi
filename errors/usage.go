package errors

type ErrUsage struct{}

func (_ ErrUsage) Error() string {
	return "Wrong usage."
}
