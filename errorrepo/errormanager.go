package errorrepo

// DieIf manage error
func DieIf(err error) {
	if err != nil {
		panic(err)
	}
}
