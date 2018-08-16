package helpers

import "os"

func HasSources() bool {
	_, err := os.Stat(GetCalledFileNameSkip(2, true))
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}
