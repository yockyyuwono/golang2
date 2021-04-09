package model

type User struct {
	UserCode  string
	Passwords string
}

type UserSave struct {
	UserCode        string
	Passwords       string
	BadgeNumber     string
	Ssn             string
	LastLogin       string
	LoginLocation   string
	DepartmentId    string
	Email           string
	RoleId          string
	CreatedBy       string
	CreatedDate     string
	LastUpdatedBy   string
	LastUpdatedDate string
	Active          string
}
