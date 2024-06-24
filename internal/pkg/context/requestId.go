package context

import "context"

const RequestIdKey string = "request-id"

func GetRequestID(ctx context.Context) string {
	rid, ok := ctx.Value(RequestIdKey).(string)
	if !ok || rid == "" {
		return "N/A"
	}

	return rid
}

func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIdKey, requestID)
}
