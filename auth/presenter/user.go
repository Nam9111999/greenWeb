package presenter

import "green.env.com/auth/model"

type CreateResponseWrapper struct {
	User *model.User `json:"user"`
}
