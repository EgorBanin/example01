package before

import (
	"activeobject/before/domain"
	"activeobject/common"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"
)

var (
	BadRequestErr     = errors.New("bad request")
	InternalServerErr = errors.New("internal server error")
)

func Handle(cnf common.Config, rq common.CreateProfileRequest) (any, error) {
	p := domain.Profile{
		Name: rq.Name,
	}

	if !cnf.GayEnabled && slices.Contains([]string{domain.GenderTransF, domain.GenderTransM}, rq.Gender) {
		return nil, BadRequestErr
	}

	p.Gender = rq.Gender

	if !cnf.DickPicAllowed && isDickPic(rq.Photo) {
		return nil, BadRequestErr
	}

	p.Photo = rq.Photo

	if rq.Age < cnf.AgeOfConsent && (slices.Contains(rq.Purposes, domain.PurposeSex) ||
		slices.Contains(rq.Purposes, domain.PurposeMarriage)) {
		return nil, BadRequestErr
	}

	p.Age = rq.Age
	p.Purposes = rq.Purposes

	if !cnf.GayEnabled &&
		!slices.Contains(p.Purposes, domain.PurposeFriendship) &&
		slices.Contains(rq.LookingForGender, p.Gender) {
		return nil, BadRequestErr
	}

	p.LookingForGender = rq.LookingForGender

	if !slices.Contains(rq.Purposes, domain.PurposeFriendship) && rq.LookingForAge.From != nil {
		f := *rq.LookingForAge.From
		if f < cnf.AgeOfConsent {
			return nil, BadRequestErr
		}
	}

	p.LookingForAge = domain.AgeRange{
		From: rq.LookingForAge.From,
		To:   rq.LookingForAge.To,
	}

	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("%w: json.MarshalIndent: %s", InternalServerErr, err.Error())
	}

	return string(b), nil
}

func isDickPic(pic string) bool {
	return strings.Contains(pic, "penis")
}
