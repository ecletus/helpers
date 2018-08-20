package helpers

import "github.com/moisespsena/go-error-wrap"

func CheckReturnError(funcs ...func() (key string, err error)) (name string, index int, err error) {
	for i, f := range funcs {
		name, err = f()
		if err != nil {
			return name, i, err
		}
	}
	return "", 0, nil
}

func CheckReturnE(funcs ...func() (key string, err error)) error {
	name, index, err := CheckReturnError(funcs...)
	if err != nil {
		return errwrap.Wrap(err, "Func %d %q", index, name)
	}
	return nil
}
