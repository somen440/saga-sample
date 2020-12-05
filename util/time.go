package util

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func IsValidTimestamp(ts *timestamppb.Timestamp) bool {
	return time.Now().Before(ts.AsTime())
}
