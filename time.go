package lodash

import "time"

// Now ... Returns the current time in Unix Epoch
func Now() int64 {
	return time.Now().Unix()
}
