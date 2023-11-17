package types

type User struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	PermissionLevel int    `json:"permission_level"`
	ApiKey          string `json:"api_key"`
}

const (
	Admin = 58193
	Base  = 1
)

type Shark struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	MaxLength      int    `json:"max_length"`
	Ocean          string `json:"ocean"`
	TopSpeed       int    `json:"top_speed"`
	AttacksPerYear int    `json:"attacks_per_year"`
}
