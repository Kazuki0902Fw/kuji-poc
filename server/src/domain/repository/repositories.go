package repository

type Repositories struct {
	User                  UserRepository
	IntellectualProperty IntellectualPropertyRepository
}

func NewRepositories(
	userRepository UserRepository,
	intellectualPropertyRepository IntellectualPropertyRepository,
) *Repositories {
	return &Repositories{
		User:                  userRepository,
		IntellectualProperty: intellectualPropertyRepository,
	}
}
