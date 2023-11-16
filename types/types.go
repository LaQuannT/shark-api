package types

type User struct {
	Id              int
	Name            string
	Email           string
	PermissionLevel int
	ApiKey          string
}

const (
	Admin = 10
	Base  = 1
)

type Shark struct {
	Name           string
	Type           string
	MaxLength      int
	Ocean          string
	TopSpeed       int
	AttacksPerYear int
}
