package model

type RoleRight struct {
	RoleID  string
	Section string
	Route   string
	RCreate bool
	RRead   bool
	RUpdate bool
	RDelete bool
}