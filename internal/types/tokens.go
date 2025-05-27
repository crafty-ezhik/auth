package types

type AccessToken string
type RefreshToken string

type PairToken struct {
	AccessToken  AccessToken
	RefreshToken RefreshToken
}
