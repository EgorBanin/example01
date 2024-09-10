package validators

import (
	"activeobject/after/domain"
	"errors"
	"slices"
)

var GayDetectedErr = errors.New("gay detected")

type GayValidator struct {
	gayEnabled bool
}

func NewGayValidator(gayEnabled bool) *GayValidator {
	return &GayValidator{gayEnabled: gayEnabled}
}

func (v *GayValidator) Validate(data domain.ProfileData) error {
	if !v.gayEnabled {
		if slices.Contains([]string{domain.GenderTransF, domain.GenderTransM}, data.Gender) {
			return GayDetectedErr
		}

		if slices.Contains(data.LookingForGender, data.Gender) && !slices.Contains(data.Purposes, domain.PurposeFriendship) {
			return GayDetectedErr
		}
	}

	return nil
}
