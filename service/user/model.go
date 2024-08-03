package user

import (
	"github.com/alaa2amz/g1/service/model"
)

var (
Path = "/user"
	DroppedColumns        = []string{"publish_at", "afloat"}
)

type  User model.User
func Proto() (p User) { return }
func Protos() (p []User) { return }
