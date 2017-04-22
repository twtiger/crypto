package utils

import (
	"github.com/twstrike/ed448"
	"github.com/twtiger/crypto/curve"

	. "gopkg.in/check.v1"
)

var (
	randData = []byte{
		// x1
		0x40, 0x80, 0x66, 0x2d, 0xd8, 0xe7, 0xf0, 0x9c,
		0xdf, 0xb0, 0x4e, 0x1c, 0x6e, 0x12, 0x62, 0xa3,
		0x7c, 0x31, 0x9a, 0xe1, 0xe7, 0x86, 0x87, 0xcc,
		0x82, 0x05, 0x78, 0xe6, 0x44, 0x2f, 0x4f, 0x77,
		0x0e, 0xd1, 0xb4, 0x48, 0xa6, 0x05, 0x90, 0x5e,
		0xe7, 0xba, 0xfc, 0x25, 0x99, 0x99, 0xb8, 0xc3,
		0x90, 0x3e, 0xf4, 0xa3, 0x75, 0xee, 0x85, 0x32,
		// x2
		0x16, 0xb1, 0x06, 0x5b, 0x81, 0xea, 0xac, 0xb3,
		0x69, 0x47, 0x6d, 0xa2, 0xaa, 0x86, 0x0b, 0xe5,
		0xcd, 0xac, 0x43, 0xd7, 0xb7, 0xe3, 0xb0, 0x85,
		0xd8, 0x66, 0xf9, 0xb6, 0x45, 0x2e, 0x81, 0x43,
		0xc2, 0x6f, 0x61, 0xc4, 0xdd, 0x65, 0x35, 0xa4,
		0xa4, 0xf9, 0x55, 0xf0, 0xf9, 0xd2, 0xf4, 0xb7,
		0xa4, 0xf9, 0x55, 0xf0, 0xf9, 0xd2, 0xf4, 0xb7,
		// y1
		0x52, 0x18, 0x41, 0x48, 0x60, 0x2d, 0x67, 0x8a,
		0xd3, 0xf3, 0xd2, 0xa4, 0xfd, 0x6f, 0x64, 0xf3,
		0x72, 0x82, 0xb0, 0x6a, 0x4d, 0xea, 0x9c, 0xef,
		0x99, 0x05, 0xe1, 0x8d, 0xaf, 0x2d, 0xdb, 0x52,
		0x57, 0x00, 0xac, 0x45, 0x24, 0x24, 0xb4, 0x79,
		0x02, 0x5f, 0x99, 0x70, 0x95, 0x2a, 0x90, 0x08,
		0x02, 0x5f, 0x99, 0x70, 0x95, 0x2a, 0x90, 0x08,
		// y2
		0x51, 0x5b, 0x69, 0x03, 0xd5, 0x77, 0xb0, 0x77,
		0x35, 0x1f, 0x1b, 0x2d, 0xb1, 0x26, 0xf1, 0x69,
		0x3b, 0xcc, 0x4b, 0x0a, 0x95, 0x83, 0xd7, 0xec,
		0xfa, 0x8c, 0xf7, 0x80, 0xbe, 0x9b, 0x6d, 0xb4,
		0xc3, 0x24, 0x3c, 0x94, 0x9b, 0x63, 0xbc, 0x89,
		0xbc, 0x09, 0x39, 0xb8, 0xbf, 0xa2, 0x9b, 0xf4,
		0x3a, 0xa2, 0x9b, 0xbe, 0x6e, 0x78, 0x7b, 0x11,
		// z
		0x66, 0x60, 0x01, 0xb9, 0x83, 0x10, 0xd5, 0x7d,
		0xe4, 0x86, 0x58, 0x0a, 0x42, 0xd2, 0x2a, 0x74,
		0xe9, 0x5d, 0x77, 0xc4, 0x08, 0x46, 0x31, 0xb4,
		0x75, 0x1b, 0xf2, 0x67, 0x23, 0x19, 0x5e, 0xb6,
		0xfc, 0xe8, 0xd1, 0x38, 0x81, 0xa3, 0x98, 0x41,
		0xdf, 0xdf, 0x5d, 0x8d, 0x41, 0xb4, 0x66, 0x0f,
		0x39, 0xe1, 0x6f, 0x8c, 0x89, 0xed, 0xf6, 0x11,
	}
)

func (s *UtilsSuite) Test_RandomScalar(c *C) {
	scalar, err := RandScalar(FixedRand(randData))

	exp := curve.Ed448GoldScalar(
		[]byte{
			0x40, 0x80, 0x66, 0x2d, 0xd8, 0xe7, 0xf0, 0x9c,
			0xdf, 0xb0, 0x4e, 0x1c, 0x6e, 0x12, 0x62, 0xa3,
			0x7c, 0x31, 0x9a, 0xe1, 0xe7, 0x86, 0x87, 0xcc,
			0x82, 0x05, 0x78, 0xe6, 0x44, 0x2f, 0x4f, 0x77,
			0x0e, 0xd1, 0xb4, 0x48, 0xa6, 0x05, 0x90, 0x5e,
			0xe7, 0xba, 0xfc, 0x25, 0x99, 0x99, 0xb8, 0xc3,
			0x90, 0x3e, 0xf4, 0xa3, 0x75, 0xee, 0x85, 0x32,
		},
	)

	c.Assert(err, IsNil)
	c.Assert(scalar, DeepEquals, exp)
}

func (s *UtilsSuite) Test_RandomLongTermScalar(c *C) {
	scalar, err := RandLongTermScalar(FixedRand(randData))

	exp := ed448.NewScalar(
		[]byte{
			0xc6, 0xd0, 0x98, 0x2e, 0xe4, 0xe5, 0x81, 0xe4,
			0x61, 0x3c, 0x46, 0x99, 0x0a, 0x37, 0x79, 0xc3,
			0xfa, 0xe5, 0xd5, 0x29, 0x27, 0x31, 0xa3, 0x55,
			0x9f, 0x34, 0x91, 0xd1, 0x0c, 0x7f, 0x88, 0x56,
			0x8c, 0x62, 0xe1, 0x86, 0xb7, 0xef, 0xd6, 0xcb,
			0x1b, 0x14, 0x88, 0x3b, 0xc0, 0xfb, 0xac, 0x46,
			0x0c, 0xc7, 0x20, 0x82, 0x3e, 0xd0, 0xdc, 0x2c,
		},
	)

	c.Assert(err, IsNil)
	c.Assert(scalar, DeepEquals, exp)
}
