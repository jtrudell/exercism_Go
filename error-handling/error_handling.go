package erratum

const testVersion = 2

func Use(o ResourceOpener, input string) (err error) {
	var res Resource
	var openError error

	for {
		res, openError = o()
		if openError != nil {
			_, ok := openError.(TransientError)
			if !ok {
				return openError
			}
		} else {
			break
		}
	}

	defer res.Close()

	defer func() {
		if r := recover(); r != nil {
			frobErr, ok := err.(FrobError)
			if ok {
				res.Defrob(frobErr.defrobTag)
			}

			err, ok = res.(error)
			if !ok {
				return
			}
		}
	}()

	res.Frob(input)

	return nil
}
