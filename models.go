package main

// Device Struct
type Device struct {
	ID         string // Id number of the device
	TypeNumber string // Type of the device
	Version    string // Version of the devices firmware
	Status     string // e.g. "in production", "in sales", "operational", ...
	Enabled    bool   // enabled / disabled
}

// App struct
type App struct {
	ID      string // id of the app instance
	AppName string // name of the app
	Version string // version of the app
	Enabled bool   // enabled / disabled
}

// User struct
type User struct {
	ID        string // id of the user
	Login     string // login name of the user
	Firstname string // last name
	Lastname  string // first name
	Email     string // users email
	Enabled   bool   // enabled / disabled
}

// OwnerMapping struct
type OwnerMapping struct {
	ID        string // id of the mapping
	OwnerID   string // id of the owner (could be a User, App or Device)
	SubjectID string // id of the owned subject (could be a User, App or Device)
}

// AccessRight struct
type AccessRight struct {
	ID                    string        // id of the AccessRight
	DisplayName           string        // display name
	Description           string        // description to explain results of assigning this access right
	InheritedAccessRights []AccessRight // list of inherited access rights (optional)
}

// AccessRightsMapping struct
type AccessRightMapping struct {
	ID           string        // id of the mapping
	SubjectID    string        // id of the subject (who is granted access rights to the object)
	ObjectID     string        // id of the object (to whom access rights are granted)
	AccessRights []AccessRight // mapped AccessRights
}
