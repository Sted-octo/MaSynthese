package domain

type ILobGetter interface {
	Get(accessToken string, lobId int64) (*Lob, error)
}
