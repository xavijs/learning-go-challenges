package exception

type AdDescriptionTooLongException struct{}

func (e AdDescriptionTooLongException) Error() string {
	return "Ad description is too long"
}
