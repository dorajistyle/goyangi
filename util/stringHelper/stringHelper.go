package stringHelper

import "bytes"

// Concat concatenates string to buffer.
// According to 'Efficient String Concatenation in Go(http://herman.asia/efficient-string-concatenation-in-go)',
// bytes.Buffer is best choice for heavy-duty case.
// You should call buffer.String() to get a concatenated string after all concaternating finished.
func Concat(buffer *bytes.Buffer, str string) {
	buffer.WriteString(str)
}

// ConcatExist concatenates string to string array.
// According to 'Efficient String Concatenation in Go(http://herman.asia/efficient-string-concatenation-in-go)',
// When str is already exist, it's faster than buffer concatenation.
// You should call strings.Join(strs, "") to get a concatenated string after all concaternating finished.
func ConcatExist(strs []string, str string) []string {
	return append(strs, str)
}
