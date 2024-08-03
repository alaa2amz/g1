package login

import (
	"github.com/alaa2amz/g1/service/model"
)

var (
Path = "/login"
	DroppedColumns        = []string{"publish_at", "afloat"}
)

type  Login model.Login
func Proto() (p Login) { return }
func Protos() (p []Login) { return }
