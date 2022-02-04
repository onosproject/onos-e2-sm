// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

import (
	"gotest.tools/assert"
	"testing"
)

func Test_newAsnCodecsPrim(t *testing.T) {

	msg := []byte{0x02}

	byteArrayC := newAsnCodecsPrim(msg)
	assert.Equal(t, 1, int(byteArrayC.size), "unexpected size")
}

func Test_decodeAsnCodecsPrim(t *testing.T) {

	msg := []byte{0x02}

	byteArrayC := newAsnCodecsPrim(msg)

	result := decodeAsnCodecsPrim(byteArrayC)
	assert.Equal(t, len(result), len(msg), "Something went wrong")
}
