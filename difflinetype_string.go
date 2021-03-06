// Code generated by "stringer -type DiffLineType -trimprefix DiffLine -tags static"; DO NOT EDIT.

package git

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DiffLineContext-32]
	_ = x[DiffLineAddition-43]
	_ = x[DiffLineDeletion-45]
	_ = x[DiffLineContextEOFNL-61]
	_ = x[DiffLineAddEOFNL-62]
	_ = x[DiffLineDelEOFNL-60]
	_ = x[DiffLineFileHdr-70]
	_ = x[DiffLineHunkHdr-72]
	_ = x[DiffLineBinary-66]
}

const (
	_DiffLineType_name_0 = "Context"
	_DiffLineType_name_1 = "Addition"
	_DiffLineType_name_2 = "Deletion"
	_DiffLineType_name_3 = "DelEOFNLContextEOFNLAddEOFNL"
	_DiffLineType_name_4 = "Binary"
	_DiffLineType_name_5 = "FileHdr"
	_DiffLineType_name_6 = "HunkHdr"
)

var (
	_DiffLineType_index_3 = [...]uint8{0, 8, 20, 28}
)

func (i DiffLineType) String() string {
	switch {
	case i == 32:
		return _DiffLineType_name_0
	case i == 43:
		return _DiffLineType_name_1
	case i == 45:
		return _DiffLineType_name_2
	case 60 <= i && i <= 62:
		i -= 60
		return _DiffLineType_name_3[_DiffLineType_index_3[i]:_DiffLineType_index_3[i+1]]
	case i == 66:
		return _DiffLineType_name_4
	case i == 70:
		return _DiffLineType_name_5
	case i == 72:
		return _DiffLineType_name_6
	default:
		return "DiffLineType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
