// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package internal contains quoted-printable internals shared by mime and
// mime/quotedprintable.
package internal

import "fmt"

// EncodeByte encodes a byte using the quoted-printable encoding.
func EncodeByte(dst []byte, b byte) {
	dst[0] = '='
	dst[1] = upperhex[b>>4]
	dst[2] = upperhex[b&0x0f]
}

const upperhex = "0123456789ABCDEF"

func fromHex(b byte) (byte, error) {
	switch {
	case b >= '0' && b <= '9':
		return b - '0', nil
	case b >= 'A' && b <= 'F':
		return b - 'A' + 10, nil
	// Accept badly encoded bytes
	case b >= 'a' && b <= 'f':
		return b - 'a' + 10, nil
	}
	return 0, fmt.Errorf("quotedprintable: invalid quoted-printable hex byte %#02x", b)
}

// ReadHexByte returns the byte represented by an hexadecimal byte slice of length 2.
func ReadHexByte(v []byte) (b byte, err error) {
	var hb, lb byte
	if hb, err = fromHex(v[0]); err != nil {
		return 0, err
	}
	if lb, err = fromHex(v[1]); err != nil {
		return 0, err
	}
	return hb<<4 | lb, nil
}
