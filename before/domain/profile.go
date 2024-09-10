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

type Profile struct {
	Name             string
	Gender           string
	Photo            string
	Age              int
	Purposes         []string
	LookingForGender []string
	LookingForAge    AgeRange
}

type AgeRange struct {
	From *int
	To   *int
}
