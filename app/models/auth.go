package models

// LoginDTO defined the /login payload
type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"password"`
}

// SignupDTO defined the /login payload
type SignupDTO struct {
	Rank           string `json:"rank" validate:"required,min=2"`
	FirstName      string `json:"firstName" validate:"required,min=1"`
	LastName       string `json:"lastName" validate:"required,min=1"`
	OrganizationID uint   `json:"organizationID" validate:"required,gte=1"`
	Trade          string `json:"trade" validate:"required,min=3"`
	Email          string `json:"email" validate:"required,email,endswith=@ecn.forces.gc.ca|endswith=forces.gc.ca|endswith=@rcafinnovation.ca"`
}


// AccessResponse todo
type AccessResponse struct {
	Token string `json:"token"`
}

// AuthResponse todo
type AuthResponse struct {
	User *UserResponse   `json:"user"`
	Auth *AccessResponse `json:"auth"`
}
