package handler

import (
	"context"
	"fmt"
	"log/slog"

	"maxnap/platform/internal/generated/schema"
)

func (s Server) PostPublish(
	ctx context.Context,
	request schema.PostPublishRequestObject,
) (schema.PostPublishResponseObject, error) {
	s.logger.WithFields(slog.Group("request",
		slog.String("user_id", request.Body.UserId),
		slog.String("id", request.Body.Id),
		slog.String("body", request.Body.Body),
		slog.String("title", request.Body.Title),
	))

	if request.Body.Title != "Good task" {
		s.logger.Error("not supported title")
		return nil, fmt.Errorf("not supported title: %s", request.Body.Title)
	}

	// post, err := ParseBody[schema.Post](r.Body, s.logger)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	//
	//
	// w.WriteHeader(http.StatusOK)
	// w.Header().Add("Content-Type", "application/json")
	// w.Write([]byte("Okay"))

	return schema.PostPublish201JSONResponse{
		Result: &schema.PostPublishResponseResult{
			PostId: "123",
		},
	}, nil
}
