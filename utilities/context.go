package utilities

import "golang.org/x/net/context"

var Ctx context.Context

func ProvideContext(ctx context.Context) {
	Ctx = ctx
}
