// Code generated by "stringer -type=Version"; DO NOT EDIT.

package http

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[HTTP1_0-0]
	_ = x[HTTP1_1-1]
	_ = x[HTTP2-2]
	_ = x[HTTP3-3]
}

const _Version_name = "HTTP1_0HTTP1_1HTTP2HTTP3"

var _Version_index = [...]uint8{0, 7, 14, 19, 24}

func (i Version) String() string {
	if i < 0 || i >= Version(len(_Version_index)-1) {
		return "Version(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Version_name[_Version_index[i]:_Version_index[i+1]]
}
