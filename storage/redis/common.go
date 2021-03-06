package redis

import (
	"strconv"
	"time"

	"github.com/xh3b4sd/anna/spec"
)

func (s *storage) retryErrorLogger(err error, d time.Duration) {
	s.Log.WithTags(spec.Tags{C: nil, L: "E", O: s, V: 4}, "retry error: %#v", maskAny(err))
}

func (s *storage) withPrefix(keys ...string) string {
	newKey := s.Prefix

	for _, k := range keys {
		newKey += ":" + k
	}

	return newKey
}

func parseMultiBulkReply(reply []interface{}) (int64, []string, error) {
	cursor, err := strconv.ParseInt(string(reply[0].([]uint8)), 10, 64)
	if err != nil {
		return 0, nil, maskAny(err)
	}
	var values []string
	for _, v := range reply[1].([]interface{}) {
		values = append(values, string(v.([]uint8)))
	}

	return cursor, values, nil
}
