package payload

type CreateRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
