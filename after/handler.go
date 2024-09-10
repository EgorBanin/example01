package after

import (
	"activeobject/after/domain"
	"activeobject/after/domain/validators"
	"activeobject/common"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	BadRequestErr     = errors.New("bad request")
	InternalServerErr = errors.New("internal server error")
)

func Handle(cnf common.Config, rq common.CreateProfileRequest) (any, error) {
	p := domain.NewProfile(domain.ProfileData{
		Name:             rq.Name,
		Gender:           rq.Gender,
		Photo:            rq.Photo,
		Age:              rq.Age,
		Purposes:         rq.Purposes,
		LookingForGender: rq.LookingForGender,
		LookingForAge: domain.AgeRange{
			From: rq.LookingForAge.From,
			To:   rq.LookingForAge.To,
		},
	})
	if err := p.Validate(
		validators.NewDickPicValidator(cnf.DickPicAllowed),
		validators.NewGayValidator(cnf.GayEnabled),
		validators.NewPedophiliaValidator(cnf.AgeOfConsent),
	); err != nil {
		return nil, fmt.Errorf("%w: profile.Validate: %s", BadRequestErr, err.Error())
	}

	b, err := json.MarshalIndent(p.ProfileData, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("%w: json.MarshalIndent: %s", InternalServerErr, err.Error())
	}

	return string(b), nil
}
