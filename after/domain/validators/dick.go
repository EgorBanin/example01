package validators

import (
	"activeobject/after/domain"
	"errors"
	"strings"
)

var DickDetectedErr = errors.New("dick detected")

type DickPicValidator struct {
	dickPicAllowed bool
}

func NewDickPicValidator(dickPicAllowed bool) *DickPicValidator {
	return &DickPicValidator{dickPicAllowed: dickPicAllowed}
}

func (v *DickPicValidator) Validate(data domain.ProfileData) error {
	if !v.dickPicAllowed && strings.Contains(data.Photo, "penis") {
		return DickDetectedErr
	}

	return nil
}
