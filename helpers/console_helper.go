package helpers

import "fmt"

func WrapError(customMsg string, originalErr error) error {
	return fmt.Errorf("%s : %v", customMsg, originalErr)
}
