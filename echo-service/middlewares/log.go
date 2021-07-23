package middlewares

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"os"
)

func LogMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		w := log.NewSyncWriter(os.Stderr)
		logger := log.NewLogfmtLogger(w)
		err = logger.Log("req",request)
		if err != nil {
			return nil, err
		}
		resp,err := next(ctx,request)
		return resp,err
	}
}