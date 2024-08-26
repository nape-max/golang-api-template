package handler

import (
	"context"

	"maxnap/platform/internal/generated/schema"
)

func (s Server) GetPing(
	ctx context.Context,
	request schema.GetPingRequestObject,
) (schema.GetPingResponseObject, error) {
	return schema.GetPing200JSONResponse{
		Ping: "Pong",
	}, nil
}
