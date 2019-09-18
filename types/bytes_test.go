// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
// Copyright (C) 2019  Centrifuge GmbH
//
// This file is part of Go Substrate RPC Client (GSRPC).
//
// GSRPC is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// GSRPC is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package types_test

import (
	"testing"

	. "github.com/centrifuge/go-substrate-rpc-client/types"
)

func TestBytes_EncodeDecode(t *testing.T) {
	assertRoundtrip(t, NewBytes(mustDecodeHexString("0x00")))
	assertRoundtrip(t, NewBytes(mustDecodeHexString("0xab1234")))
	assertRoundtrip(t, NewBytes(mustDecodeHexString("0x0001")))
}

func TestBytes_EncodedLength(t *testing.T) {
	assertEncodedLength(t, []encodedLengthAssert{
		{NewBytes(mustDecodeHexString("0x00")), 2},
		{NewBytes(mustDecodeHexString("0xab1234")), 4},
		{NewBytes(mustDecodeHexString("0x0001")), 3},
	})
}

func TestBytes_Encode(t *testing.T) {
	assertEncode(t, []encodingAssert{
		{NewBytes([]byte{0, 0, 0}), mustDecodeHexString("0x0c000000")},
		{NewBytes([]byte{171, 18, 52}), mustDecodeHexString("0x0cab1234")},
		{NewBytes([]byte{0, 1}), mustDecodeHexString("0x080001")},
		{NewBytes([]byte{18, 52, 86}), mustDecodeHexString("0x0c123456")},
	})
}

func TestBytes_Hash(t *testing.T) {
	assertHash(t, []hashAssert{
		{NewBytes([]byte{0, 42, 254}), mustDecodeHexString(
			"0xabf7fe6eb94e0816bf2db57abb296d012f7cb9ddfe59ebf52f9c2770f49a0a46")},
		{NewBytes([]byte{0, 0}), mustDecodeHexString(
			"0xd1200120e01c48b4bbf7e1cd7ebab20087b34ea11e1e9e4ebc2f207aea77139d")},
	})
}

func TestBytes_Hex(t *testing.T) {
	assertEncodeToHex(t, []encodeToHexAssert{
		{NewBytes([]byte{0, 0, 0}), "0x0c000000"},
		{NewBytes([]byte{171, 18, 52}), "0x0cab1234"},
		{NewBytes([]byte{0, 1}), "0x080001"},
		{NewBytes([]byte{18, 52, 86}), "0x0c123456"},
	})
}

func TestBytes_String(t *testing.T) {
	assertString(t, []stringAssert{
		{NewBytes([]byte{0, 0, 0}), "[0 0 0]"},
		{NewBytes([]byte{171, 18, 52}), "[171 18 52]"},
		{NewBytes([]byte{0, 1}), "[0 1]"},
		{NewBytes([]byte{18, 52, 86}), "[18 52 86]"},
	})
}

func TestBytes_Eq(t *testing.T) {
	assertEq(t, []eqAssert{
		{NewBytes([]byte{1, 0, 0}), NewBytes([]byte{1, 0}), false},
		{NewBytes([]byte{0, 0, 1}), NewBytes([]byte{0, 1}), false},
		{NewBytes([]byte{0, 0, 0}), NewBytes([]byte{0, 0}), false},
		{NewBytes([]byte{12, 48, 255}), NewBytes([]byte{12, 48, 255}), true},
		{NewBytes([]byte{0}), NewBytes([]byte{0}), true},
		{NewBytes([]byte{1}), NewBool(true), false},
		{NewBytes([]byte{0}), NewBool(false), false},
	})
}

func TestBytes8_EncodeDecode(t *testing.T) {
	assertRoundtrip(t, NewBytes8([8]byte{}))
	assertRoundtrip(t, NewBytes8([8]byte{0, 1, 2, 3, 4, 5, 6, 7}))
}

func TestBytes8_EncodedLength(t *testing.T) {
	assertEncodedLength(t, []encodedLengthAssert{
		{NewBytes8([8]byte{}), 9},
		{NewBytes8([8]byte{7, 6, 5, 4, 3, 2, 1, 0}), 9},
	})
}

func TestBytes8_Encode(t *testing.T) {
	assertEncode(t, []encodingAssert{
		{NewBytes8([8]byte{0, 0, 0}), mustDecodeHexString("0x200000000000000000")},
		{NewBytes8([8]byte{171, 18, 52}), mustDecodeHexString("0x20ab12340000000000")},
	})
}

func TestBytes8_Hash(t *testing.T) {
	assertHash(t, []hashAssert{
		{NewBytes8([8]byte{0, 42, 254}), mustDecodeHexString(
			"0xaaafe19720d6258bc2186e1b127f35471b4c53257ca648367674887aa5abd19f")},
		{NewBytes8([8]byte{0, 0}), mustDecodeHexString(
			"0xf4630e4005914ba92fa80ad507cacdf1f7df5dc3bccb0ec4eba7fed3e257a000")},
	})
}

func TestBytes8_Hex(t *testing.T) {
	assertEncodeToHex(t, []encodeToHexAssert{
		{NewBytes8([8]byte{0, 0, 0}), "0x200000000000000000"},
		{NewBytes8([8]byte{171, 18, 52}), "0x20ab12340000000000"},
	})
}

func TestBytes8_String(t *testing.T) {
	assertString(t, []stringAssert{
		{NewBytes8([8]byte{0, 0, 0}), "[0 0 0 0 0 0 0 0]"},
		{NewBytes8([8]byte{171, 18, 52}), "[171 18 52 0 0 0 0 0]"},
	})
}

func TestBytes8_Eq(t *testing.T) {
	assertEq(t, []eqAssert{
		{NewBytes8([8]byte{1, 0, 0}), NewBytes8([8]byte{1, 0}), true},
		{NewBytes8([8]byte{0, 0, 1}), NewBytes8([8]byte{0, 1}), false},
		{NewBytes8([8]byte{0, 0, 0}), NewBytes8([8]byte{0, 0}), true},
		{NewBytes8([8]byte{12, 48, 255}), NewBytes8([8]byte{12, 48, 255}), true},
		{NewBytes8([8]byte{0}), NewBytes8([8]byte{0}), true},
		{NewBytes8([8]byte{1}), NewBool(true), false},
		{NewBytes8([8]byte{0}), NewBool(false), false},
	})
}

func TestBytes16_EncodeDecode(t *testing.T) {
	assertRoundtrip(t, NewBytes16([16]byte{0, 255, 0, 42}))
}

func TestBytes32_EncodeDecode(t *testing.T) {
	assertRoundtrip(t, NewBytes32([32]byte{0, 255, 0, 42}))
}

func TestBytes64_EncodeDecode(t *testing.T) {
	assertRoundtrip(t, NewBytes64([64]byte{0, 255, 0, 42}))
}

func TestBytes128_EncodeDecode(t *testing.T) {
	assertRoundtrip(t, NewBytes128([128]byte{0, 255, 0, 42}))
}

func TestBytes256_EncodeDecode(t *testing.T) {
	assertRoundtrip(t, NewBytes256([256]byte{0, 255, 0, 42}))
}

func TestBytes512_EncodeDecode(t *testing.T) {
	assertRoundtrip(t, NewBytes512([512]byte{0, 255, 0, 42}))
}

func TestBytes1024_EncodeDecode(t *testing.T) {
	assertRoundtrip(t, NewBytes1024([1024]byte{0, 255, 0, 42}))
}

func TestBytes2048_EncodeDecode(t *testing.T) {
	assertRoundtrip(t, NewBytes2048([2048]byte{0, 255, 0, 42}))
}