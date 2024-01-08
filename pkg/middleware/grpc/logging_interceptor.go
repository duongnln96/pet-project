package grpc

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"google.golang.org/grpc"
)

func NewLoggingUnaryInterceptor(logger *slog.Logger) *loggingUnaryInterceptor {
	return &loggingUnaryInterceptor{
		logger: logger,
	}
}

type loggingUnaryInterceptor struct {
	logger *slog.Logger
}

func (m *loggingUnaryInterceptor) LoggingUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	startTime := time.Now()
	// Log the request body
	reqJSON, _ := json.Marshal(req)
	// Call the actual handler to process the request
	resp, err := handler(ctx, req)
	// Log the response body
	respJSON, _ := json.Marshal(resp)

	latency := time.Since(startTime).Milliseconds()

	attrs := []slog.Attr{
		{
			Key:   "is_keep",
			Value: slog.BoolValue(true),
		},
		{
			Key:   "latency",
			Value: slog.Int64Value(latency),
		},
		{
			Key:   "body-request",
			Value: slog.StringValue(string(reqJSON)),
		},
		{
			Key:   "body-response",
			Value: slog.StringValue(string(respJSON)),
		},
	}

	if err != nil {
		attrs = append(attrs, slog.Attr{
			Key:   "err_info",
			Value: slog.StringValue(err.Error()),
		})
	}

	m.logger.LogAttrs(ctx, slog.LevelInfo, info.FullMethod, attrs...)

	return resp, err
}
