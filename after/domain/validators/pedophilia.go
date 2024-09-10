package validators

import (
	"activeobject/after/domain"
	"errors"
)

var PedophiliaDetectedErr = errors.New("pedophilia detected")

type PedophiliaValidator struct {
	ageOfConsent int
}

func NewPedophiliaValidator(ageOfConsent int) *PedophiliaValidator {
	return &PedophiliaValidator{ageOfConsent: ageOfConsent}
}

func (v *PedophiliaValidator) Validate(data domain.ProfileData) error {
	if data.Age < v.ageOfConsent {
		for _, p := range data.Purposes {
			if p == domain.PurposeSex || p == domain.PurposeMarriage {
				return PedophiliaDetectedErr
			}
		}
	}

	return nil
}
