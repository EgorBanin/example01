package common

type CreateProfileRequest struct {
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
