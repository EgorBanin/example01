package domain

const (
	GenderM      = "m"
	GenderF      = "f"
	GenderTransM = "tm"
	GenderTransF = "tf"

	PurposeSex        = "sex"
	PurposeMarriage   = "marriage"
	PurposeFriendship = "friendship"
)

type AgeRange struct {
	From *int
	To   *int
}

type ProfileData struct {
	Name             string
	Gender           string
	Photo            string
	Age              int
	Purposes         []string
	LookingForGender []string
	LookingForAge    AgeRange
}

type Validator interface {
	Validate(data ProfileData) error
}

type Profile struct {
	ProfileData
}

func NewProfile(data ProfileData) *Profile {
	return &Profile{ProfileData: data}
}

func (p *Profile) Validate(validators ...Validator) error {
	for i := range validators {
		if err := validators[i].Validate(p.ProfileData); err != nil {
			return err
		}
	}

	return nil
}
