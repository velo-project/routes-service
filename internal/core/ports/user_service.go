package ports

type GetUserIdByEmailPort interface {
	Execute(email string) (*int, error)
}
