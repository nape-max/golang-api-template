package handler

import (
	"context"

	"maxnap/platform/internal/generated/schema"
)

func (s Server) PostDelete(
	ctx context.Context,
	request schema.PostDeleteRequestObject,
) (schema.PostDeleteResponseObject, error) {
	return schema.PostDelete200JSONResponse{}, nil
}
