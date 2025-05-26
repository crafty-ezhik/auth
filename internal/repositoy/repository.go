package repositoy

type TokenStorage interface {
	IncrUserVersion(userID string) error
	isBlackList(refreshToken string) (bool, error)
	AddToBlackList(refreshToken string) error
}
