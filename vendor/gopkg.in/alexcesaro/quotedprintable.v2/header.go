// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file defines encoding and decoding functions for encoded-words
// as defined in RFC 2047.

package quotedprintable

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
	"sync"
	"unicode/utf8"

	"gopkg.in/alexcesaro/quotedprintable.v2/internal"
)

// Encoding represents the encoding used in encoded-words. It must be one of the
// two encodings defined in RFC 2047 ("B" or "Q" encoding).
type Encoding string

const (
	// Q represents the Q-encoding defined in RFC 2047.
	Q Encoding = "Q"
	// B represents the Base64 encoding defined in RFC 2045.
	B Encoding = "B"
)

// A HeaderEncoder is an RFC 2047 encoded-word encoder.
type HeaderEncoder struct {
	charset    string
	encoding   Encoding
	splitWords bool
}

const maxEncodedWordLen = 75 // As defined in RFC 2047, section 2

// NewHeaderEncoder returns a new HeaderEncoder to encode strings in the
// specified charset.
func (enc Encoding) NewHeaderEncoder(charset string) *HeaderEncoder {
	// We automatically split encoded-words only when the charset is UTF-8
	// because since multi-octet character must not be split across adjacent
	// encoded-words (see RFC 2047, section 5) there is no way to split words
	// without knowing how the charset works.
	splitWords := strings.ToUpper(charset) == "UTF-8"

	return &HeaderEncoder{charset, enc, splitWords}
}

// Encode encodes a string to be used as a MIME header value.
func (e *HeaderEncoder) Encode(s string) string {
	return e.encodeWord(s)
}

// NeedsEncoding returns whether the given header content needs to be encoded
// into an encoded-words.
func NeedsEncoding(s string) bool {
	for i := 0; i < len(s); i++ {
		b := s[i]
		if (b > '~' || b < ' ') && b != '\t' {
			return true
		}
	}

	return false
}

type bufPool struct {
	sync.Pool
}

func (b *bufPool) GetBuffer() *bytes.Buffer {
	return bufFree.Get().(*bytes.Buffer)
}

func (b *bufPool) PutBuffer(buf *bytes.Buffer) {
	if buf.Len() > 1024 {
		return
	}
	buf.Reset()
	b.Put(buf)
}

var bufFree = &bufPool{
	sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	},
}

// encodeWord encodes a string into an encoded-word.
func (e *HeaderEncoder) encodeWord(s string) string {
	buf := bufFree.GetBuffer()
	defer bufFree.PutBuffer(buf)

	e.openWord(buf)

	switch {
	case e.encoding == B:
		e.encodeWordB(buf, s)
	default:
		e.encodeWordQ(buf, s)
	}

	e.closeWord(buf)
	return buf.String()
}

func (e *HeaderEncoder) encodeWordB(buf *bytes.Buffer, s string) {
	maxLen := maxEncodedWordLen - buf.Len() - 2
	if !e.splitWords || base64.StdEncoding.EncodedLen(len(s)) <= maxLen {
		buf.WriteString(base64.StdEncoding.EncodeToString([]byte(s)))
		return
	}

	v := []byte(s)
	var n, last, runeSize int
	for i := 0; i < len(s); i += runeSize {
		runeSize = getRuneSize(s, i)

		if base64.StdEncoding.EncodedLen(n+runeSize) <= maxLen {
			n += runeSize
		} else {
			buf.WriteString(base64.StdEncoding.EncodeToString(v[last:i]))
			e.splitWord(buf)
			last = i
			n = runeSize
		}
	}
	buf.WriteString(base64.StdEncoding.EncodeToString(v[last:]))
}

func (e *HeaderEncoder) encodeWordQ(buf *bytes.Buffer, s string) {
	if !e.splitWords {
		for i := 0; i < len(s); i++ {
			writeQ(buf, s[i])
		}
		return
	}

	var runeSize int
	n := buf.Len()
	for i := 0; i < len(s); i += runeSize {
		b := s[i]
		var encLen int
		if b >= ' ' && b <= '~' && b != '=' && b != '?' && b != '_' {
			encLen, runeSize = 1, 1
		} else {
			runeSize = getRuneSize(s, i)
			encLen = 3 * runeSize
		}

		// We remove 2 to let spaces for closing chars "?="
		if n+encLen > maxEncodedWordLen-2 {
			n = e.splitWord(buf)
		}
		writeQString(buf, s[i:i+runeSize])
		n += encLen
	}
}

func (e *HeaderEncoder) openWord(buf *bytes.Buffer) int {
	n := buf.Len()
	buf.WriteString("=?")
	buf.WriteString(e.charset)
	buf.WriteByte('?')
	buf.WriteString(string(e.encoding))
	buf.WriteByte('?')

	return buf.Len() - n
}

func (e *HeaderEncoder) closeWord(buf *bytes.Buffer) {
	buf.WriteString("?=")
}

func (e *HeaderEncoder) splitWord(buf *bytes.Buffer) int {
	e.closeWord(buf)
	buf.WriteString("\r\n ")
	return e.openWord(buf)
}

func getRuneSize(s string, i int) int {
	runeSize := 1
	for i+runeSize < len(s) && !utf8.RuneStart(s[i+runeSize]) {
		runeSize++
	}

	return runeSize
}

func writeQString(buf *bytes.Buffer, s string) {
	for i := 0; i < len(s); i++ {
		writeQ(buf, s[i])
	}
}

func writeQ(buf *bytes.Buffer, b byte) {
	switch {
	case b == ' ':
		buf.WriteByte('_')
	case b >= '!' && b <= '~' && b != '=' && b != '?' && b != '_':
		buf.WriteByte(b)
	default:
		enc := make([]byte, 3)
		internal.EncodeByte(enc, b)
		buf.Write(enc)
	}
}

// DecodeHeader decodes a MIME header by decoding all encoded-words of the
// header. The returned text is encoded in the returned charset. Text is not
// necessarily encoded in UTF-8. Returned charset is always upper case. This
// function does not support decoding headers with multiple encoded-words
// using different charsets.
func DecodeHeader(header string) (text, charset string, err error) {
	buf := bufFree.GetBuffer()
	defer bufFree.PutBuffer(buf)

	for {
		i := strings.IndexByte(header, '=')
		if i == -1 {
			break
		}
		if i > 0 {
			buf.WriteString(header[:i])
			header = header[i:]
		}

		word := getEncodedWord(header)
		if word == "" {
			buf.WriteByte('=')
			header = header[1:]
			continue
		}

		for {
			dec, wordCharset, err := decodeWord(word)
			if err != nil {
				buf.WriteString(word)
				header = header[len(word):]
				break
			}
			if charset == "" {
				charset = wordCharset
			} else if charset != wordCharset {
				return "", "", fmt.Errorf("quotedprintable: multiple charsets in header are not supported: %q and %q used", charset, wordCharset)
			}
			buf.Write(dec)
			header = header[len(word):]

			// White-space and newline characters separating two encoded-words
			// must be deleted.
			var j int
			for j = 0; j < len(header); j++ {
				b := header[j]
				if b != ' ' && b != '\t' && b != '\n' && b != '\r' {
					break
				}
			}
			if j == 0 {
				// If there are no white-space characters following the current
				// encoded-word there is nothing special to do.
				break
			}
			word = getEncodedWord(header[j:])
			if word == "" {
				break
			}
			header = header[j:]
		}
	}
	buf.WriteString(header)

	return buf.String(), charset, nil
}

func getEncodedWord(s string) string {
	if len(s) < 2 {
		return ""
	}
	if s[0] != '=' {
		return ""
	}
	if s[1] != '?' {
		return ""
	}

	n := 2
	for {
		if n >= len(s) {
			return ""
		}

		b := s[n]
		if (b < '0' || b > '9') &&
			(b < 'A' || b > 'Z') &&
			(b < 'a' || b > 'z') &&
			b != '-' {
			break
		}

		n++
	}
	if s[n] != '?' {
		return ""
	}

	if n+2 >= len(s) {
		return ""
	}
	b := s[n+1]
	if b != 'Q' && b != 'B' && b != 'q' && b != 'b' {
		return ""
	}
	if s[n+2] != '?' {
		return ""
	}

	n = n + 3
	for {
		if n >= len(s) {
			return ""
		}

		if s[n] < ' ' || s[n] > '~' {
			return ""
		}
		if s[n] == '?' {
			n++
			break
		}
		n++
	}
	if n >= len(s) || s[n] != '=' {
		return ""
	}

	return s[0 : n+1]
}

func decodeWord(s string) (text []byte, charset string, err error) {
	fields := strings.Split(s, "?")
	if len(fields) != 5 || fields[0] != "=" || fields[4] != "=" || len(fields[2]) != 1 {
		return []byte(s), "", nil
	}

	charset, enc, src := fields[1], fields[2], fields[3]

	var dec []byte
	switch Encoding(strings.ToUpper(enc)) {
	case B:
		if dec, err = base64.StdEncoding.DecodeString(src); err != nil {
			return dec, charset, err
		}
	case Q:
		if dec, err = qDecode(src); err != nil {
			return dec, charset, err
		}
	default:
		return []byte(""), charset, fmt.Errorf("quotedprintable: RFC 2047 encoding not supported: %q", enc)
	}

	return dec, strings.ToUpper(charset), nil
}

// qDecode decodes a Q encoded string.
func qDecode(s string) ([]byte, error) {
	length := len(s)
	for i := 0; i < len(s); i++ {
		if s[i] == '=' {
			length -= 2
			i += 2
		}
	}
	dec := make([]byte, length)

	n := 0
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case c == '_':
			dec[n] = ' '
		case c == '=':
			if i+2 >= len(s) {
				return []byte(""), io.ErrUnexpectedEOF
			}
			buf, err := internal.ReadHexByte([]byte(s[i+1:]))
			if err != nil {
				return []byte(""), err
			}
			dec[n] = buf
			i += 2
		case (c >= ' ' && c <= '~') || c == '\n' || c == '\r' || c == '\t':
			dec[n] = c
		default:
			return []byte(""), fmt.Errorf("quotedprintable: invalid unescaped byte %#02x in Q encoded string at byte %d", c, i)
		}
		n++
	}

	return dec, nil
}
