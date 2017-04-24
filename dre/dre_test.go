package dre

import (
	"crypto/rand"
	"testing"

	"github.com/twtiger/crypto/cramershoup"
	"github.com/twtiger/crypto/curve"
	"github.com/twtiger/crypto/testHelpers"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type DRESuite struct{}

var _ = Suite(&DRESuite{})

var (
	testMessage = []byte{
		0x5d, 0xf1, 0x18, 0xbf, 0x8e, 0x3f, 0xfe, 0xcd,
		0x95, 0xd3, 0x49, 0xda, 0xcd, 0xac, 0x2c, 0xdf,
		0x72, 0x5e, 0xb7, 0x61, 0x44, 0xf1, 0x93, 0xa6,
		0x70, 0x8e, 0x64, 0xff, 0x7c, 0xec, 0x6c, 0xe5,
		0xc6, 0x8d, 0x8f, 0xa0, 0x43, 0x23, 0x45, 0x33,
		0x73, 0x71, 0xe6, 0x2f, 0x57, 0xbb, 0x0f, 0x70,
		0x11, 0x8c, 0x62, 0x26, 0x9e, 0x17, 0x5d, 0x22,
	}

	randDREData = []byte{
		0xc9, 0x21, 0xa6, 0x41, 0xc3, 0x43, 0xb3, 0x4f,
		0x3e, 0x86, 0x99, 0xbf, 0x11, 0x75, 0x2c, 0x40,
		0x05, 0xb9, 0x0e, 0xd1, 0x01, 0xd8, 0x3e, 0xeb,
		0xda, 0xfa, 0x7e, 0x28, 0x94, 0xe8, 0x62, 0x31,
		0xa5, 0x62, 0xfd, 0x27, 0x85, 0x00, 0xdf, 0x4a,
		0xc3, 0xc2, 0x27, 0x2e, 0x11, 0x49, 0xfc, 0x3c,
		0xc0, 0xdf, 0x80, 0x3d, 0x7a, 0x2f, 0x1f, 0x06,

		0xc9, 0x21, 0xa6, 0x41, 0xc3, 0x43, 0xb3, 0x4f,
		0x3e, 0x86, 0x99, 0xbf, 0x11, 0x75, 0x2c, 0x40,
		0x05, 0xb9, 0xff, 0xd1, 0x01, 0xd8, 0x3e, 0xeb,
		0xda, 0xfa, 0x7e, 0x28, 0x20, 0xe8, 0x62, 0x31,
		0xa5, 0x34, 0xfd, 0x27, 0x85, 0x00, 0xdd, 0x4a,
		0xcc, 0xc2, 0x27, 0xee, 0x11, 0x10, 0xfc, 0x3c,
		0xc0, 0xdf, 0x80, 0x3d, 0x7a, 0x2f, 0x1f, 0x06,

		0xc9, 0x21, 0xa6, 0x41, 0xc3, 0x43, 0xb3, 0x4f,
		0x3e, 0x86, 0x99, 0xbf, 0x11, 0x75, 0x2c, 0x40,
		0x05, 0xb9, 0xff, 0xd1, 0x01, 0xd8, 0x3e, 0xeb,
		0xda, 0xfa, 0x7e, 0x28, 0x20, 0xe8, 0x62, 0x31,
		0xa5, 0x34, 0xfd, 0x27, 0x85, 0x00, 0xdd, 0x4a,
		0xcc, 0xc2, 0x27, 0xee, 0x11, 0x10, 0xfc, 0x3c,
		0xc0, 0xdf, 0x80, 0x3d, 0x7a, 0x2f, 0x1f, 0x06,

		0xc9, 0x21, 0xa6, 0x41, 0xc3, 0x43, 0xb3, 0x4f,
		0x3e, 0x86, 0x99, 0xbf, 0x11, 0x75, 0x2c, 0x40,
		0x05, 0xb9, 0xff, 0xd1, 0x01, 0xd8, 0x3e, 0xeb,
		0xda, 0xfa, 0x7e, 0x28, 0x20, 0xe8, 0x62, 0x31,
		0xa5, 0x34, 0xfd, 0x27, 0x85, 0x00, 0xdd, 0x4a,
		0xcc, 0xc2, 0x27, 0xee, 0x11, 0x10, 0xfc, 0x3c,
		0xc0, 0xdf, 0x80, 0x3d, 0x7a, 0x2f, 0x1f, 0x06,
	}

	randNIZKPKData = []byte{
		0xc9, 0x21, 0xa6, 0x41, 0xc3, 0x43, 0xb3, 0x4f,
		0x3e, 0x86, 0x99, 0xbf, 0x11, 0x75, 0x2c, 0x40,
		0x05, 0xb9, 0xff, 0xd1, 0x01, 0xd8, 0x3e, 0xeb,
		0xda, 0xfa, 0x7e, 0x28, 0x20, 0xe8, 0x62, 0x31,
		0xa5, 0x34, 0xfd, 0x27, 0x85, 0x00, 0xdd, 0x4a,
		0xcc, 0xc2, 0x27, 0xee, 0x11, 0x10, 0xfc, 0x3c,
		0xc0, 0xdf, 0x80, 0x3d, 0x7a, 0x2f, 0x1f, 0x06,

		0xc9, 0x21, 0xa6, 0x41, 0xc3, 0x43, 0xb3, 0x4f,
		0x3e, 0x86, 0x99, 0xbf, 0x11, 0x75, 0x2c, 0x40,
		0x05, 0xb9, 0xff, 0xd1, 0x01, 0xd8, 0x3e, 0xeb,
		0xda, 0xfa, 0x7e, 0x28, 0x20, 0xe8, 0x62, 0x31,
		0xa5, 0x34, 0xfd, 0x27, 0x85, 0x00, 0xdd, 0x4a,
		0xcc, 0xc2, 0x27, 0xee, 0x11, 0x10, 0xfc, 0x3c,
		0xc0, 0xdf, 0x80, 0x3d, 0x7a, 0x2f, 0x1f, 0x06,
	}

	testPubA = &cramershoup.PublicKey{
		C: curve.Ed448GoldPoint(
			[16]uint32{
				0x0600d17a, 0x0a9375b0, 0x057841c1, 0x0d174be0,
				0x0011badb, 0x006ef801, 0x02e0a39f, 0x0fecd541,
				0x055b1c78, 0x0895cbd0, 0x072a6628, 0x0bb03485,
				0x036131cf, 0x0a79778a, 0x07e006b9, 0x097fb665,
			},
			[16]uint32{
				0x0f010a4f, 0x03e8789d, 0x047c75cc, 0x07ec8505,
				0x0225156f, 0x07e8e08f, 0x0ac4e95e, 0x065ee99b,
				0x077b97fd, 0x0c851d94, 0x07e2ef48, 0x0004fe3e,
				0x0c3e7fcb, 0x0a9dcbf1, 0x0296ae5f, 0x0b844a1f,
			},
			[16]uint32{
				0x0cb95a0b, 0x079e46f8, 0x03337dcb, 0x00dcdf8d,
				0x0629b80f, 0x0ff224ea, 0x0e697c53, 0x0a5bc379,
				0x03001fb6, 0x026826d2, 0x02954624, 0x07574fd0,
				0x0cf79d59, 0x07396756, 0x016f3154, 0x0009df66,
			},
			[16]uint32{
				0x0bcdc5cd, 0x0492d355, 0x09d9d8a5, 0x01fea4f6,
				0x0081e336, 0x0697e45b, 0x0c712575, 0x0e7b112b,
				0x08fdb73a, 0x06ae6268, 0x0c271c2a, 0x084b467f,
				0x0155f3e6, 0x054cf783, 0x0b453295, 0x042c22ff,
			},
		),
		D: curve.Ed448GoldPoint(
			[16]uint32{
				0x03aee5b7, 0x032b3f7f, 0x06176ed0, 0x056fa571,
				0x04d01a0b, 0x0b382fa3, 0x03d55289, 0x04c8e69c,
				0x0ab17cae, 0x07b995ce, 0x03263f63, 0x08d4efc8,
				0x0382c935, 0x0587fbd5, 0x03d42439, 0x0b6979c9,
			},
			[16]uint32{
				0x070e385a, 0x04b3e06c, 0x0d8e3b40, 0x04064fff,
				0x04c1ae53, 0x0348a758, 0x04fa81f9, 0x06ef7f5a,
				0x0be2c435, 0x0b6c6794, 0x0159c719, 0x0350c7e1,
				0x0a5d1620, 0x0c9e7983, 0x0c90bd0e, 0x01196a72,
			},
			[16]uint32{
				0x0174aff6, 0x0aa1703e, 0x0b3d41e7, 0x0d68f123,
				0x0fcc832b, 0x0c11adbe, 0x0faecb96, 0x0152998d,
				0x0902ed06, 0x03560403, 0x067b0008, 0x0825eef3,
				0x0bb42471, 0x0ae05b98, 0x06c93c9b, 0x0bb36c4b,
			},
			[16]uint32{
				0x0970457b, 0x0bbb1746, 0x065927c4, 0x012c45ac,
				0x03a587e6, 0x095e2a6c, 0x0ec77f11, 0x09226042,
				0x0c304e97, 0x01783fbf, 0x0f4f1dbd, 0x07628142,
				0x0981f1fd, 0x0311dbc5, 0x05c12822, 0x033e76b1,
			},
		),
		H: curve.Ed448GoldPoint(
			[16]uint32{
				0x06808e72, 0x0ce35b11, 0x0e5e2f5c, 0x0b88b4d4,
				0x0869c12a, 0x04414132, 0x0bb898a8, 0x07c1e17c,
				0x0f04e50e, 0x068bad3b, 0x05c8d2b1, 0x0682f5cb,
				0x0a6b80e2, 0x0519b3a5, 0x045b7bec, 0x02b1b0d6,
			},
			[16]uint32{
				0x059d303a, 0x072683d3, 0x01b3a38d, 0x0b73118c,
				0x05dc7e0e, 0x0cd643d7, 0x09575347, 0x0e7653ae,
				0x0c59d3e1, 0x00d2a8d6, 0x0d9d3cb6, 0x0539c8ab,
				0x0d2cdc35, 0x03e95ff4, 0x0ca0a361, 0x0d6b571f,
			},
			[16]uint32{
				0x028916ca, 0x024a5ca9, 0x0ff426c7, 0x093dda43,
				0x0781af41, 0x07ec215e, 0x0e3deaef, 0x05963af4,
				0x0f1db9f4, 0x0018b7b8, 0x020b8cb8, 0x0e497381,
				0x0b98d304, 0x0750e83f, 0x00d61916, 0x0f0809f0,
			},
			[16]uint32{
				0x04b92c3b, 0x0d44025b, 0x09d68237, 0x0efa91f3,
				0x080def8c, 0x08703dcb, 0x0e39b56a, 0x0e3017a0,
				0x05ecb8cc, 0x0cd53123, 0x0c69b8db, 0x0fde3887,
				0x0cb571d9, 0x0e0580f7, 0x0b44788e, 0x087c0443,
			},
		),
	}

	testSecA = &cramershoup.SecretKey{
		X1: curve.Ed448GoldScalar([]byte{
			0x70, 0x27, 0xc3, 0x28, 0x5d, 0xb9, 0x02, 0x11,
			0xbd, 0xc2, 0x6b, 0xef, 0xb3, 0xb2, 0xe7, 0x6d,
			0x6f, 0x2e, 0x20, 0xf6, 0x97, 0xb1, 0xfe, 0xa1,
			0xc6, 0x75, 0x11, 0xc8, 0x24, 0x6d, 0x73, 0x44,
			0x8f, 0x28, 0xeb, 0x1d, 0x15, 0x48, 0x36, 0xac,
			0xc5, 0x5f, 0xbe, 0xc7, 0xa9, 0x04, 0x13, 0x03,
			0x3d, 0x6a, 0xda, 0xc6, 0x7c, 0x36, 0x71, 0x06,
		}),
		X2: curve.Ed448GoldScalar([]byte{
			0x94, 0xf0, 0xd5, 0xed, 0x37, 0x78, 0x8a, 0xdb,
			0x91, 0xdf, 0x1f, 0xb9, 0xab, 0xc4, 0x04, 0x94,
			0x7c, 0xa4, 0x7d, 0x19, 0x5b, 0x5c, 0xc5, 0x88,
			0x57, 0xcf, 0x99, 0x57, 0xa8, 0x9d, 0x08, 0x17,
			0x2e, 0x7b, 0x3e, 0xd4, 0x43, 0xbe, 0x1c, 0x65,
			0xf4, 0x0e, 0x8c, 0x65, 0x08, 0x64, 0xb1, 0x91,
			0x49, 0xd5, 0xfa, 0x15, 0xbd, 0xc4, 0x4f, 0x1b,
		}),
		Y1: curve.Ed448GoldScalar([]byte{
			0xf8, 0xa9, 0x1a, 0xd6, 0x23, 0xb0, 0xbc, 0x94,
			0x3f, 0x9b, 0x84, 0xd4, 0x6a, 0xbf, 0x9e, 0xb8,
			0x9e, 0x03, 0x2e, 0xf6, 0xe0, 0xa7, 0x25, 0x1a,
			0x95, 0xda, 0xd1, 0xa6, 0x3e, 0xa2, 0x9d, 0x46,
			0x66, 0x98, 0x34, 0x58, 0x26, 0x82, 0xf5, 0x52,
			0xc9, 0x86, 0x5e, 0xf7, 0x39, 0x9e, 0x5c, 0x6e,
			0xfb, 0x46, 0xbc, 0xf8, 0x3d, 0xdc, 0x1b, 0x06,
		}),
		Y2: curve.Ed448GoldScalar([]byte{
			0xe0, 0x1d, 0xd8, 0x77, 0x7b, 0x40, 0x92, 0xae,
			0xb1, 0x01, 0x2e, 0x5a, 0xcf, 0x54, 0x1f, 0xc2,
			0x52, 0x89, 0x8d, 0xb4, 0x5c, 0x02, 0xe3, 0x65,
			0xcd, 0x6d, 0x7a, 0x0c, 0xda, 0x35, 0xfc, 0x08,
			0x39, 0xcb, 0xfc, 0x24, 0xd0, 0x92, 0xb4, 0xc5,
			0xc7, 0xb8, 0x46, 0x89, 0xb8, 0xa0, 0xfa, 0x38,
			0x7e, 0x47, 0xd8, 0xe2, 0x23, 0xda, 0x4d, 0x20,
		}),
		Z: curve.Ed448GoldScalar([]byte{
			0xa3, 0xf8, 0xe, 0xb2, 0xa6, 0x99, 0x23, 0x9a,
			0x81, 0x9b, 0x5e, 0xc3, 0x30, 0xce, 0xd7, 0x49,
			0x7b, 0xdb, 0x3b, 0xe7, 0x0d, 0xd0, 0x91, 0xec,
			0x6e, 0xc6, 0xd7, 0xdc, 0xd1, 0xd3, 0xe2, 0x68,
			0xd5, 0xf1, 0xcc, 0xd6, 0x2f, 0x87, 0xb0, 0x27,
			0xd7, 0x59, 0x89, 0x65, 0x02, 0x16, 0xec, 0x5a,
			0x0f, 0x84, 0x1a, 0xbe, 0xda, 0xa1, 0x88, 0x02,
		}),
	}

	testPubB = &cramershoup.PublicKey{
		C: curve.Ed448GoldPoint(
			[16]uint32{
				0x048603ce, 0x08657fa3, 0x0a7ed7d1, 0x00467214,
				0x0781aeaf, 0x06202e8e, 0x0142a539, 0x0c55ab78,
				0x05405585, 0x0c7d68bc, 0x0ffe9cc1, 0x061888b5,
				0x0d994802, 0x05147e54, 0x0f533d6f, 0x023ab901,
			},
			[16]uint32{
				0x06cf7d90, 0x06269f0e, 0x0ab6c4e2, 0x09804fd9,
				0x048bab98, 0x0e33fdcf, 0x0996c34a, 0x0ab80d7e,
				0x0d59830b, 0x058a5d0b, 0x026dbc6a, 0x0e2cb254,
				0x04664218, 0x0b106106, 0x011ced8c, 0x00851398,
			},
			[16]uint32{
				0x0df8cf30, 0x07b60230, 0x076e63ae, 0x06998644,
				0x05abd81b, 0x0ecd45f8, 0x0e03891e, 0x0bb625fa,
				0x09337432, 0x0a2dec2d, 0x0cc75d04, 0x023ff0ef,
				0x0be80126, 0x0cf82eed, 0x057fd1f6, 0x0fa34c11,
			},
			[16]uint32{
				0x03800dae, 0x0c0a6db2, 0x06019eaf, 0x000dda7c,
				0x0162bcd6, 0x0327a6c6, 0x0b772fef, 0x0754b0fe,
				0x00816e60, 0x0626b4f5, 0x01cb5bea, 0x0623ffcd,
				0x0c288061, 0x025fd8c5, 0x08dca40d, 0x05c1517d,
			},
		),
		D: curve.Ed448GoldPoint(
			[16]uint32{
				0x0709010b, 0x02d4ba32, 0x08faaaeb, 0x0d4e3745,
				0x029afa09, 0x069f27d1, 0x0837cb09, 0x0cd2cbd3,
				0x0ac1328c, 0x015ab0b2, 0x064acd6b, 0x08848eb0,
				0x0581b01b, 0x00e7d2fb, 0x07f23ff1, 0x0110a447,
			},
			[16]uint32{
				0x01847b50, 0x066ac557, 0x020a9cac, 0x0372f5ab,
				0x08ce345f, 0x0256b7ad, 0x0f04fe02, 0x00fb185c,
				0x0be9667e, 0x00fc1829, 0x0efeb866, 0x04b46c69,
				0x06ae76a1, 0x0af60c2b, 0x025efe1e, 0x025e4dfe,
			},
			[16]uint32{
				0x00ba52f5, 0x0aaca440, 0x03fcdb56, 0x06197634,
				0x093e91d9, 0x0c8a3675, 0x0420dc6c, 0x0c64e35f,
				0x0d16490f, 0x0a506e5f, 0x0a3c59fe, 0x057a4618,
				0x060ce8ff, 0x0c934b48, 0x09e55d0b, 0x0e715a59,
			},
			[16]uint32{
				0x08c3d764, 0x0ba5f0b8, 0x0975b4a7, 0x0e9ef8b1,
				0x067d9228, 0x0a8efb50, 0x0354f975, 0x0e6648a4,
				0x0d4bd71b, 0x04fb4240, 0x05efd032, 0x054d5c20,
				0x0b0112c1, 0x0a848b3b, 0x04e226c1, 0x097bd4a7,
			},
		),
		H: curve.Ed448GoldPoint(
			[16]uint32{
				0x00ed1e98, 0x096ec1e8, 0x0ef9b54e, 0x013e6f69,
				0x08831655, 0x0a168fe3, 0x07c19a2f, 0x0eef8698,
				0x01b07937, 0x0426d9ab, 0x007df6c1, 0x0705b02d,
				0x04ce6925, 0x088909bd, 0x0d5572a7, 0x0d99fe28,
			},
			[16]uint32{
				0x0c77ea4a, 0x034ccc3f, 0x098aad7c, 0x0524287b,
				0x0f562bc5, 0x0628d945, 0x043f5846, 0x0b504b04,
				0x0c048b25, 0x0af1a6db, 0x03be40bd, 0x03bc81b8,
				0x0d7e910d, 0x0972f700, 0x05667215, 0x061194c8,
			},
			[16]uint32{
				0x01ce0361, 0x09f661d8, 0x0f36b63d, 0x0704c10b,
				0x0374ae03, 0x04dec46c, 0x0c89cd30, 0x0bf0bb8e,
				0x0cc5274c, 0x0d5eeb28, 0x0a1d5e27, 0x07e21c1c,
				0x0731ef5f, 0x063c7eb6, 0x0499a13c, 0x0c28ae8b,
			},
			[16]uint32{
				0x0453f134, 0x0a60461d, 0x007b9436, 0x06c800dc,
				0x0f8e2d2c, 0x086375eb, 0x044921f5, 0x0e204f0f,
				0x01595b82, 0x0687d38f, 0x052e65d9, 0x05ac37bd,
				0x095f0f25, 0x0e09b7d4, 0x063d2df2, 0x052c8028,
			},
		),
	}

	testSecB = &cramershoup.SecretKey{
		X1: curve.Ed448GoldScalar([]byte{
			0xb9, 0x1c, 0xa1, 0xe6, 0x54, 0xb5, 0xdc, 0x03,
			0x11, 0x0e, 0x6f, 0xa8, 0x52, 0x6b, 0x3d, 0x7c,
			0x46, 0xbd, 0xd6, 0x1b, 0x52, 0x8b, 0x18, 0xa4,
			0x46, 0x32, 0x10, 0x1a, 0x57, 0xab, 0x51, 0xbe,
			0x9a, 0x90, 0xdd, 0x76, 0xa7, 0x82, 0x72, 0x0e,
			0x5e, 0x91, 0xda, 0x1b, 0xbd, 0xca, 0x47, 0x46,
			0xfb, 0x49, 0xc1, 0x74, 0x8f, 0xa5, 0x79, 0x2e,
		},
		),
		X2: curve.Ed448GoldScalar([]byte{
			0x27, 0x3b, 0x9e, 0x9c, 0x88, 0x5d, 0x7d, 0xca,
			0x38, 0xba, 0x2e, 0x0c, 0x02, 0xa3, 0xd5, 0x31,
			0x0f, 0x4c, 0x96, 0x26, 0xb1, 0x85, 0xc6, 0x4d,
			0x00, 0x68, 0x3e, 0xb0, 0xb5, 0xe1, 0xcf, 0x9c,
			0xfc, 0x99, 0x9d, 0xe4, 0xc1, 0x97, 0x96, 0x81,
			0x97, 0x7d, 0x85, 0x5f, 0x26, 0x73, 0x09, 0xa9,
			0x1c, 0xde, 0x71, 0xe5, 0x5d, 0x25, 0x23, 0x02,
		},
		),
		Y1: curve.Ed448GoldScalar([]byte{
			0x2b, 0xbb, 0x99, 0x17, 0xa6, 0x86, 0xc6, 0x5b,
			0x80, 0xc3, 0x15, 0xe2, 0x92, 0x96, 0x99, 0x8b,
			0x11, 0x9a, 0x1e, 0x0b, 0x2c, 0x7a, 0x48, 0x0e,
			0xcf, 0x9c, 0x8e, 0x7c, 0x62, 0x9a, 0x9a, 0xaa,
			0x8a, 0x40, 0x3c, 0xec, 0xb6, 0xbf, 0x4a, 0x28,
			0x86, 0x99, 0x72, 0xde, 0x52, 0x53, 0x9c, 0x77,
			0x36, 0xb9, 0x4d, 0xc3, 0x40, 0x3b, 0xf0, 0x1d,
		},
		),
		Y2: curve.Ed448GoldScalar([]byte{
			0x9c, 0xc7, 0x6d, 0xdf, 0xc5, 0x89, 0xc5, 0x24,
			0xcb, 0x90, 0x12, 0xe6, 0xab, 0xaa, 0x73, 0x79,
			0x56, 0x25, 0xf8, 0x6d, 0x13, 0x86, 0xca, 0xdd,
			0xa7, 0x81, 0x44, 0xbb, 0x31, 0x41, 0x7b, 0x50,
			0xae, 0x43, 0x53, 0x2a, 0x3f, 0x3d, 0x8c, 0x7f,
			0xd6, 0x56, 0x70, 0x88, 0x8f, 0x73, 0x24, 0xd6,
			0xfc, 0x22, 0x13, 0x36, 0xb6, 0x15, 0x81, 0x2d,
		},
		),
		Z: curve.Ed448GoldScalar([]byte{
			0x35, 0xb5, 0x48, 0x9c, 0xb2, 0xee, 0xd8, 0xc9,
			0xe9, 0x5b, 0x62, 0x4b, 0x04, 0x93, 0x44, 0x29,
			0xb9, 0x17, 0xf2, 0xc6, 0x23, 0xfc, 0x3d, 0xd4,
			0x6e, 0x57, 0x83, 0x8a, 0xdb, 0x3d, 0xfe, 0x18,
			0x28, 0x8e, 0x31, 0x23, 0x1f, 0x36, 0xf0, 0xff,
			0x6f, 0xb8, 0x4d, 0xa1, 0xae, 0x1f, 0xd4, 0xaf,
			0xe1, 0x1b, 0x6d, 0x49, 0xe7, 0x40, 0xe0, 0x2e,
		},
		),
	}

	testPubC = curve.Ed448GoldPoint(
		[16]uint32{
			0xbaaead5, 0x685c976, 0xb0ca061, 0xba86cd7,
			0xea519fa, 0x57e4ab4, 0xc02f5fc, 0xb82204c,
			0xd78542c, 0x74a01e5, 0x9328d21, 0x3871d1a,
			0x5dcaa1c, 0x0156612, 0x7ea2255, 0xd9d787d,
		},
		[16]uint32{
			0x1a1f403, 0xbfea69f, 0x4a6d633, 0xa60a88d,
			0x4e36fdc, 0x8028db6, 0xc2fe5ba, 0xc58f5c6,
			0x85375d6, 0xbef9d4f, 0x2953a6a, 0xa779d7a,
			0xb729468, 0x9b47792, 0x0ac10fe, 0x434eff1,
		},
		[16]uint32{
			0x803ebe1, 0x18470ae, 0xb80caad, 0x3b777f9,
			0x67a6139, 0x6d9aa39, 0x787cf2e, 0x7c11d72,
			0x0374caf, 0x02cd168, 0x3d6858b, 0xdac0675,
			0xcaae54b, 0xffe21b4, 0x4bb4af8, 0x55a2cce,
		},
		[16]uint32{
			0x384109e, 0x40695a8, 0xa822b8e, 0x6026944,
			0x8e9ae46, 0xaad36b0, 0xa10ec79, 0x505a46c,
			0x4e7c598, 0x0b9daf8, 0xe22fb37, 0xa2aeb13,
			0x126a250, 0x874aa07, 0x0fe3b33, 0x1be9c1d,
		})

	invalidPub = &cramershoup.PublicKey{
		C: curve.Ed448GoldPoint(
			[16]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
			[16]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
			[16]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
			[16]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
		),
		D: curve.Ed448GoldPoint(
			[16]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
			[16]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
			[16]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
			[16]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
		),
		H: curve.Ed448GoldPoint(
			[16]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
			[16]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
			[16]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
			[16]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
		),
	}

	testDRMessage = &drMessage{
		drCipher{
			// u11
			curve.Ed448GoldPoint(
				[16]uint32{
					0x0a17e677, 0x0c72738b, 0x044b91ea, 0x0aeee85a,
					0x097a1724, 0x0944b5a6, 0x0511f65a, 0x0e43351c,
					0x01012f88, 0x0afb7a50, 0x0b351695, 0x00db5165,
					0x0a4208f5, 0x046cb655, 0x04613291, 0x0a634ecc,
				},
				[16]uint32{
					0x0c79617c, 0x0c7fe24e, 0x01e77c1f, 0x0080bbf4,
					0x064227f8, 0x04c633e7, 0x074f4065, 0x0e4bffeb,
					0x08cdef09, 0x01f11868, 0x0819242d, 0x02dc8499,
					0x09544f24, 0x0b49eb65, 0x0a71cab9, 0x00562d85,
				},
				[16]uint32{
					0x005c962f, 0x08a14cc2, 0x09d9e8f5, 0x0d173a26,
					0x0b2be661, 0x036d1bc7, 0x0320d04d, 0x01cf0f71,
					0x07a2db99, 0x0ca1501c, 0x04db7eda, 0x0ffbf0c9,
					0x01850aa5, 0x0c2a6add, 0x0e15a4ec, 0x0a8e75e4,
				},
				[16]uint32{
					0x08480d1d, 0x0ef3dbcc, 0x037eb9f7, 0x03e74b03,
					0x09b0f66c, 0x043e7fb7, 0x00697125, 0x0378e43c,
					0x08740e77, 0x036135fb, 0x03930dca, 0x0964da2e,
					0x09b771fa, 0x00fef489, 0x0de92c8d, 0x0c35584d,
				},
			),
			// u21
			curve.Ed448GoldPoint(
				[16]uint32{
					0x0be81ed1, 0x07cd9250, 0x0f14c9a0, 0x00caddba,
					0x02c9f028, 0x0a8c516c, 0x0f5daf07, 0x0dda6081,
					0x008615e6, 0x00b47acf, 0x02cbfa2e, 0x0650e516,
					0x09ff4e42, 0x0e34bc86, 0x012aa6b4, 0x0da18443,
				},
				[16]uint32{
					0x01c66bf9, 0x05347f55, 0x070733bf, 0x083c3390,
					0x033f645f, 0x07f454a8, 0x020317ca, 0x09628235,
					0x0ea10625, 0x06772961, 0x0c132cbf, 0x0c93a6ae,
					0x0133b800, 0x09ad6c99, 0x0c3708c8, 0x0220d97a,
				},
				[16]uint32{
					0x0dd25779, 0x012fb2bd, 0x021ca196, 0x03f6288b,
					0x05915652, 0x01dbfecf, 0x01514d5c, 0x09ca77ff,
					0x0e8c3200, 0x012ed313, 0x082b0716, 0x0ab2d598,
					0x0882428f, 0x0e00f355, 0x0287d490, 0x0ef93f0a,
				},
				[16]uint32{
					0x076e2640, 0x094802ae, 0x0d7fd074, 0x091d9ef8,
					0x03b650f2, 0x0bddd3e5, 0x00d5f4e3, 0x02e0aa79,
					0x004bad45, 0x00a3b440, 0x09de886d, 0x067d3fc9,
					0x02b8223c, 0x093df563, 0x04ef4ca5, 0x0d82aefc,
				},
			),
			// e1
			curve.Ed448GoldPoint(
				[16]uint32{
					0x033e0cc6, 0x0efef0cf, 0x0f419393, 0x00d09b8f,
					0x028d8004, 0x05f18fd4, 0x0110fb7b, 0x0f252160,
					0x0a51f338, 0x0985567e, 0x0315fd57, 0x09d3c67e,
					0x0b82fe37, 0x0c552ab0, 0x0e58e076, 0x08c01b8b,
				},
				[16]uint32{
					0x057a0710, 0x02c2214f, 0x0fa0e26a, 0x040b0a36,
					0x087cdc54, 0x0ab3381e, 0x0366717c, 0x0e481d05,
					0x08badd65, 0x046dde6a, 0x0b7c9307, 0x0f9f770e,
					0x085ed0c6, 0x09450192, 0x06fcf6c3, 0x02d4515b,
				},
				[16]uint32{
					0x05ee64c7, 0x0f3befd6, 0x06b81cb8, 0x0b1021ef,
					0x032bfd8d, 0x07305134, 0x05cd62bb, 0x032f519b,
					0x01ed8c7b, 0x090c93ee, 0x0aaaaf92, 0x08094eef,
					0x0eb3e08a, 0x025c880c, 0x0ef9b16e, 0x0cbfa732,
				},
				[16]uint32{
					0x0a46fd14, 0x004a34d5, 0x04d28f53, 0x0c17a931,
					0x0711e60c, 0x0d89b8d0, 0x078d7e92, 0x03cf534f,
					0x06287f2a, 0x04b6dd63, 0x08c2893c, 0x023a0afe,
					0x0a707baf, 0x0db13e74, 0x0fff6812, 0x060b8757,
				},
			),
			// v1
			curve.Ed448GoldPoint(
				[16]uint32{
					0x0ad6ecad, 0x0dd931ae, 0x0d7f0667, 0x0bf04cf4,
					0x070e6fd9, 0x08e8f0d6, 0x034bc5a1, 0x0ec0635b,
					0x0ccf2e81, 0x0ec720bd, 0x07ae54fd, 0x085fed13,
					0x01be7c69, 0x07d21258, 0x0a506e85, 0x0f496262,
				},
				[16]uint32{
					0x0a230f14, 0x0d90d0fe, 0x0d49fdaf, 0x088d26bc,
					0x0c098725, 0x03ce24ca, 0x0a9b3b0b, 0x02aeb985,
					0x09b13796, 0x0d689a92, 0x073d320e, 0x01b6e619,
					0x0769cb7e, 0x0c01ce9c, 0x0dbd63a9, 0x0530467d,
				},
				[16]uint32{
					0x08388298, 0x0f924918, 0x08cffb7f, 0x0b7983c0,
					0x03c897ef, 0x066e5cc5, 0x0c0c5f26, 0x0e25e130,
					0x0b4b63cb, 0x080fa248, 0x0c5f32ed, 0x02f7a55d,
					0x085ede11, 0x0216a520, 0x0c889e9f, 0x05168915,
				},
				[16]uint32{
					0x00ae0a66, 0x0f0ce915, 0x054c877c, 0x0b71fc0f,
					0x0bafbe5a, 0x0e609399, 0x057c453b, 0x01a8bae7,
					0x0c1842b4, 0x078a82a3, 0x0d7979d9, 0x09732b3e,
					0x0460ad24, 0x0c653230, 0x08e93554, 0x0818f738,
				},
			),
			// u12
			curve.Ed448GoldPoint(
				[16]uint32{
					0x0c8b8396, 0x03aa75e7, 0x08c3dbbd, 0x018f2180,
					0x006904c9, 0x0c57fd7a, 0x075d14a6, 0x0504e045,
					0x04b6cf4d, 0x0fde7c99, 0x0ed1ed53, 0x096ab4cc,
					0x067993f2, 0x08d5cd2c, 0x0ce72cac, 0x0fba9428,
				},
				[16]uint32{
					0x0824d64e, 0x0e9783c1, 0x02e17d29, 0x0eec032e,
					0x0ff4b999, 0x0f4c526a, 0x00e44ded, 0x0d1915f4,
					0x0174c5c6, 0x07ad3d23, 0x04260041, 0x0944f671,
					0x005b695f, 0x06e26c1d, 0x0eea52d8, 0x030dd784,
				},
				[16]uint32{
					0x040e4131, 0x0f317b4d, 0x00a5f0c4, 0x00b0ffbd,
					0x0f02c1e3, 0x0e9512b1, 0x0ec742e3, 0x099f7b96,
					0x05542fcb, 0x05acd5fe, 0x02246935, 0x03b2dfdd,
					0x099a76ed, 0x0789fad9, 0x0addead7, 0x0985940c,
				},
				[16]uint32{
					0x0eb7d91c, 0x0abc62b8, 0x01607071, 0x0fdb6b8b,
					0x057d218a, 0x0f50e376, 0x04424d3e, 0x0080ab1a,
					0x012d00b9, 0x00187d84, 0x00cbfda6, 0x05684419,
					0x03d6444f, 0x04408f88, 0x02fd2e98, 0x0c00eaae,
				},
			),
			// u22
			curve.Ed448GoldPoint(
				[16]uint32{
					0x057e244f, 0x09842135, 0x07621c02, 0x053c7677,
					0x0b59f119, 0x02bd0778, 0x00946a29, 0x05fb8eba,
					0x02b9bcd1, 0x0cfffa34, 0x00fa277d, 0x06a77894,
					0x05898996, 0x050a7056, 0x0f4e5ba9, 0x02ca34fe,
				},
				[16]uint32{
					0x0894bb48, 0x06364bf6, 0x032bd738, 0x041b580d,
					0x08d7cc58, 0x0b4d8370, 0x05b32011, 0x03ecb176,
					0x0a7c79bf, 0x0f6f0b7c, 0x0a67356c, 0x02e3cf99,
					0x04f66417, 0x023de7e3, 0x06e2e74f, 0x0143841c,
				},
				[16]uint32{
					0x027e6abf, 0x0a146a3f, 0x02fa5fcb, 0x0f52285f,
					0x0e898ab3, 0x043d8f72, 0x077f99ab, 0x066ca58c,
					0x089391d7, 0x0f8e8a79, 0x01625814, 0x00735ff5,
					0x0e2c1e27, 0x03a5882c, 0x0efd15d4, 0x0e93c854,
				},
				[16]uint32{
					0x0752c266, 0x07baee88, 0x09b961dc, 0x073e0898,
					0x06a3f190, 0x0d16def6, 0x05c702d2, 0x01bb3ff9,
					0x0928c817, 0x0139fd2c, 0x0658862a, 0x02004992,
					0x0595d978, 0x030d4ecb, 0x0f5d93f3, 0x051490e8,
				},
			),
			// e2
			curve.Ed448GoldPoint(
				[16]uint32{
					0x053b45f5, 0x04a2f630, 0x018168a1, 0x00fa524e,
					0x0cb80c2d, 0x03d08144, 0x0956ca33, 0x0418f7ce,
					0x029c80d9, 0x0cb17a75, 0x0130ba1d, 0x03077c7a,
					0x0546d429, 0x0e507ddc, 0x03b6b96e, 0x05cc60de,
				},
				[16]uint32{
					0x0576dc13, 0x08231879, 0x08b92bd1, 0x04a544e0,
					0x01bd50b1, 0x05ddff8a, 0x0c6873b5, 0x0a65a34c,
					0x09d7643e, 0x0c182d5a, 0x0ad5312a, 0x0c10c4da,
					0x0380488e, 0x094a7871, 0x07c99e05, 0x0638b295,
				},
				[16]uint32{
					0x0b3b638e, 0x0fb02938, 0x078897a4, 0x0fb4b7aa,
					0x01925e98, 0x094b40d1, 0x080558c2, 0x037b3947,
					0x036db9e6, 0x078f33f2, 0x01d17323, 0x0e1fe494,
					0x0b336ce8, 0x08815289, 0x0ff53f80, 0x0f04d08f,
				},
				[16]uint32{
					0x057e0451, 0x001fbdd4, 0x098e3651, 0x0cd20d94,
					0x0ad4a031, 0x02a942b0, 0x00983b4d, 0x0bdbc793,
					0x053b1894, 0x0a16e402, 0x09d87866, 0x0498d662,
					0x0828ebdf, 0x0043d6e5, 0x0e341588, 0x06433bc8,
				},
			),
			// v2
			curve.Ed448GoldPoint(
				[16]uint32{
					0x03950c5e, 0x00ffa746, 0x0dfe9b29, 0x0047d7db,
					0x08a59302, 0x032753f4, 0x0abc3fca, 0x08d4e54f,
					0x082df960, 0x0311057c, 0x0478f49b, 0x014ada85,
					0x0a469648, 0x02f1d1f5, 0x0a036e28, 0x01a187d5,
				},
				[16]uint32{
					0x00276f5d, 0x0df58242, 0x0e286122, 0x070ec43f,
					0x02d6d47c, 0x00e3eb3e, 0x0b50d769, 0x0a14326a,
					0x069a770a, 0x07836238, 0x0470fe02, 0x091d691d,
					0x015f829d, 0x007d88a6, 0x0c7ac0a8, 0x02cde92f,
				},
				[16]uint32{
					0x06cbf8b3, 0x08d25e21, 0x0f459d9c, 0x0d956496,
					0x0de21486, 0x0d7dab4e, 0x0cc15116, 0x04706519,
					0x08d9dcee, 0x01eb26a7, 0x093107f4, 0x06671827,
					0x0e358da7, 0x0c1a600d, 0x084f71f6, 0x0033595f,
				},
				[16]uint32{
					0x05d7d0b8, 0x0ea3a063, 0x071eafa1, 0x0950f67f,
					0x02d3db40, 0x06dae454, 0x06509291, 0x03adb3e0,
					0x0390680b, 0x004a5059, 0x07ad3ca0, 0x0cc6c068,
					0x01c04860, 0x0cec467a, 0x0fb8ebe1, 0x07ab9076,
				},
			),
		},
		&nIZKProof{
			// l
			curve.Ed448GoldScalar([]byte{
				0xf5, 0x26, 0x1a, 0xbb, 0xe9, 0x4c, 0xad, 0x18,
				0x6e, 0x31, 0x0f, 0x1d, 0x4d, 0x72, 0x59, 0x5c,
				0xd0, 0xbd, 0x56, 0x8e, 0x1a, 0xb0, 0x38, 0x4d,
				0x9a, 0x01, 0xb1, 0xe3, 0xaf, 0x80, 0xb8, 0x80,
				0x25, 0xa0, 0x31, 0x35, 0x4b, 0xd1, 0x1f, 0x52,
				0xe2, 0xd4, 0x10, 0x00, 0x30, 0x17, 0xd1, 0x48,
				0x30, 0x21, 0x29, 0xac, 0xbc, 0x1e, 0x38, 0xe,
			}),
			// n1
			curve.Ed448GoldScalar([]byte{
				0xdd, 0x4e, 0x4f, 0xae, 0x2b, 0xed, 0xff, 0x9b,
				0x68, 0x94, 0xd1, 0x23, 0xd2, 0x5c, 0x25, 0x9f,
				0xac, 0x89, 0xb9, 0x70, 0x21, 0xfa, 0xf6, 0x5d,
				0x56, 0x9e, 0x5b, 0x0b, 0xe2, 0x70, 0x46, 0xcc,
				0x74, 0x11, 0x9c, 0x43, 0x21, 0xb4, 0x2c, 0x08,
				0xdf, 0x04, 0xbe, 0xee, 0x21, 0xb4, 0xa1, 0x92,
				0xdd, 0x82, 0x7d, 0x92, 0x93, 0xb4, 0x4b, 0x3c,
			}),
			// n2
			curve.Ed448GoldScalar([]byte{
				0xd0, 0x23, 0xf6, 0x1d, 0x9a, 0x27, 0x40, 0xee,
				0x0d, 0xfe, 0x8f, 0x52, 0xee, 0x43, 0x53, 0x5b,
				0x4d, 0xeb, 0x69, 0x95, 0x21, 0x26, 0xc7, 0x19,
				0x0a, 0x2b, 0x34, 0x8f, 0xd6, 0x2e, 0x96, 0x24,
				0x9a, 0xcd, 0x6e, 0x82, 0x06, 0xa2, 0xba, 0xa4,
				0x19, 0xe8, 0xe9, 0x8a, 0x69, 0x76, 0x9e, 0x74,
				0x20, 0x57, 0xef, 0x8a, 0xc5, 0x06, 0x67, 0x06,
			}),
		},
	}
)

var d *DRE
var crsh *cramershoup.CramerShoup

func (s *DRESuite) SetUpTest(c *C) {
	d = &DRE{&curve.Ed448Gold{}}
	crsh = &cramershoup.CramerShoup{Curve: &curve.Ed448Gold{}}
}

func (s *DRESuite) Test_DREnc(c *C) {
	m, err := d.drEnc(testMessage, testHelpers.FixedRandReader(randDREData), testPubA, testPubB)
	c.Assert(m.cipher, DeepEquals, testDRMessage.cipher)
	c.Assert(m.proof, DeepEquals, testDRMessage.proof)
	c.Assert(err, IsNil)

	_, err = d.drEnc(testMessage, testHelpers.FixedRandReader(randDREData), invalidPub, testPubB)
	c.Assert(err, ErrorMatches, ".*not a valid public key")

	_, err = d.drEnc(testMessage, testHelpers.FixedRandReader([]byte{0x00}), testPubA, testPubB)
	c.Assert(err, ErrorMatches, ".*cannot source enough entropy")
}

func (s *DRESuite) Test_DRDec(c *C) {
	m, err := d.drDec(testDRMessage, testPubA, testPubB, testSecA, 1)
	c.Assert(m, DeepEquals, testMessage)
	c.Assert(err, IsNil)

	_, err = d.drDec(testDRMessage, invalidPub, testPubB, testSecA, 1)
	c.Assert(err, ErrorMatches, ".*not a valid public key")

	_, err = d.drDec(testDRMessage, testPubA, testPubB, testSecB, 1)
	c.Assert(err, ErrorMatches, ".*cannot decrypt the message")

	_, err = d.drDec(testDRMessage, testPubA, testPubB, testSecA, 2)
	c.Assert(err, ErrorMatches, ".*cannot decrypt the message")
}

func (s *DRESuite) Test_DREncryptAndDecrypt(c *C) {
	message := []byte{
		0xfd, 0xf1, 0x18, 0xbf, 0x8e, 0xc9, 0x64, 0xc7,
		0x94, 0x46, 0x49, 0xda, 0xcd, 0xac, 0x2c, 0xff,
		0x72, 0x5e, 0xb7, 0x61, 0x46, 0xf1, 0x93, 0xa6,
		0x70, 0x81, 0x64, 0x37, 0x7c, 0xec, 0x6c, 0xe5,
		0xc6, 0x8d, 0x8f, 0xa0, 0x43, 0x23, 0x45, 0x33,
		0x73, 0x79, 0xa6, 0x48, 0x57, 0xbb, 0x0f, 0x70,
		0x63, 0x8c, 0x62, 0x26, 0x9e, 0x17, 0x5d, 0x22,
	}

	keyPairA, err := crsh.GenerateKeys(rand.Reader)
	keyPairB, err := crsh.GenerateKeys(rand.Reader)

	drMessage, err := d.drEnc(message, rand.Reader, keyPairA.Pub, keyPairB.Pub)

	expMessage1, err := d.drDec(drMessage, keyPairA.Pub, keyPairB.Pub, keyPairA.Sec, 1)
	c.Assert(err, IsNil)
	c.Assert(expMessage1, DeepEquals, message)
	expMessage2, err := d.drDec(drMessage, keyPairA.Pub, keyPairB.Pub, keyPairB.Sec, 2)
	c.Assert(err, IsNil)
	c.Assert(expMessage2, DeepEquals, message)
}

func (s *DRESuite) Test_GenerationOfNIZKPK(c *C) {
	alpha1 := curve.Ed448GoldScalar([]byte{
		0x1c, 0x51, 0x56, 0x90, 0x17, 0x2d, 0x14, 0x41,
		0x2c, 0x71, 0x8e, 0x1f, 0x1f, 0x2b, 0x38, 0x60,
		0x02, 0x23, 0x42, 0x97, 0xd4, 0x5c, 0x8b, 0x9d,
		0xb9, 0x67, 0xe9, 0x11, 0x9c, 0xe3, 0xbf, 0x14,
		0x99, 0xc2, 0xe1, 0xf3, 0xa1, 0x65, 0xb8, 0x30,
		0xbc, 0x97, 0x8b, 0xa9, 0x98, 0x86, 0x53, 0x6e,
		0x9f, 0x45, 0xbd, 0x44, 0x8a, 0x40, 0x2a, 0x12,
	})

	alpha2 := curve.Ed448GoldScalar([]byte{
		0xeb, 0xb9, 0xdc, 0x1a, 0x38, 0x12, 0xed, 0xe1,
		0xbd, 0x4b, 0x2c, 0xfe, 0x12, 0x1f, 0xf8, 0x01,
		0xc9, 0x52, 0xa0, 0x69, 0xc1, 0x78, 0xa8, 0xd5,
		0xc5, 0xe1, 0x4f, 0x94, 0x01, 0x71, 0x55, 0xf6,
		0xce, 0x4b, 0xc7, 0xba, 0xe5, 0xcf, 0x02, 0xae,
		0xbd, 0xb5, 0x33, 0xd3, 0x5d, 0x5c, 0x2a, 0xf3,
		0xbf, 0xce, 0x5e, 0x4e, 0xc7, 0x4d, 0xa7, 0x3e,
	})

	k1 := curve.Ed448GoldScalar([]byte{
		0xc9, 0x21, 0xa6, 0x41, 0xc3, 0x43, 0xb3, 0x4f,
		0x3e, 0x86, 0x99, 0xbf, 0x11, 0x75, 0x2c, 0x40,
		0x05, 0xb9, 0x0e, 0xd1, 0x01, 0xd8, 0x3e, 0xeb,
		0xda, 0xfa, 0x7e, 0x28, 0x94, 0xe8, 0x62, 0x31,
		0xa5, 0x62, 0xfd, 0x27, 0x85, 0x00, 0xdf, 0x4a,
		0xc3, 0xc2, 0x27, 0x2e, 0x11, 0x49, 0xfc, 0x3c,
		0xc0, 0xdf, 0x80, 0x3d, 0x7a, 0x2f, 0x1f, 0x06,
	})

	k2 := curve.Ed448GoldScalar([]byte{
		0xc9, 0x21, 0xa6, 0x41, 0xc3, 0x43, 0xb3, 0x4f,
		0x3e, 0x86, 0x99, 0xbf, 0x11, 0x75, 0x2c, 0x40,
		0x5, 0xb9, 0xff, 0xd1, 0x1, 0xd8, 0x3e, 0xeb,
		0xda, 0xfa, 0x7e, 0x28, 0x20, 0xe8, 0x62, 0x31,
		0xa5, 0x34, 0xfd, 0x27, 0x85, 0x0, 0xdd, 0x4a,
		0xcc, 0xc2, 0x27, 0xee, 0x11, 0x10, 0xfc, 0x3c,
		0xc0, 0xdf, 0x80, 0x3d, 0x7a, 0x2f, 0x1f, 0x6,
	})

	p, err := d.genNIZKPK(testHelpers.FixedRandReader(randNIZKPKData), &testDRMessage.cipher, testPubA, testPubB, alpha1, alpha2, k1, k2)
	c.Assert(p, DeepEquals, testDRMessage.proof)
	c.Assert(err, IsNil)

	_, err = d.genNIZKPK(testHelpers.FixedRandReader([]byte{0x00}), &testDRMessage.cipher, testPubA, testPubB, alpha1, alpha2, k1, k2)
	c.Assert(err, ErrorMatches, "cannot source enough entropy")
}

func (s *DRESuite) Test_VerificationOfNIZKPK(c *C) {
	alpha1 := curve.Ed448GoldScalar([]byte{
		0x1c, 0x51, 0x56, 0x90, 0x17, 0x2d, 0x14, 0x41,
		0x2c, 0x71, 0x8e, 0x1f, 0x1f, 0x2b, 0x38, 0x60,
		0x02, 0x23, 0x42, 0x97, 0xd4, 0x5c, 0x8b, 0x9d,
		0xb9, 0x67, 0xe9, 0x11, 0x9c, 0xe3, 0xbf, 0x14,
		0x99, 0xc2, 0xe1, 0xf3, 0xa1, 0x65, 0xb8, 0x30,
		0xbc, 0x97, 0x8b, 0xa9, 0x98, 0x86, 0x53, 0x6e,
		0x9f, 0x45, 0xbd, 0x44, 0x8a, 0x40, 0x2a, 0x12,
	})

	alpha2 := curve.Ed448GoldScalar([]byte{
		0xeb, 0xb9, 0xdc, 0x1a, 0x38, 0x12, 0xed, 0xe1,
		0xbd, 0x4b, 0x2c, 0xfe, 0x12, 0x1f, 0xf8, 0x01,
		0xc9, 0x52, 0xa0, 0x69, 0xc1, 0x78, 0xa8, 0xd5,
		0xc5, 0xe1, 0x4f, 0x94, 0x01, 0x71, 0x55, 0xf6,
		0xce, 0x4b, 0xc7, 0xba, 0xe5, 0xcf, 0x02, 0xae,
		0xbd, 0xb5, 0x33, 0xd3, 0x5d, 0x5c, 0x2a, 0xf3,
		0xbf, 0xce, 0x5e, 0x4e, 0xc7, 0x4d, 0xa7, 0x3e,
	})

	valid, err := d.isValid(testDRMessage.proof, &testDRMessage.cipher, testPubA, testPubB, alpha1, alpha2)
	c.Assert(valid, Equals, true)
	c.Assert(err, IsNil)

	invalid, err := d.isValid(testDRMessage.proof, &testDRMessage.cipher, invalidPub, testPubB, alpha1, alpha2)
	c.Assert(invalid, Equals, false)
	c.Assert(err, ErrorMatches, ".*cannot decrypt the message")
}

func (s *DRESuite) Test_VerificationOfDRMessage(c *C) {
	alpha1 := curve.Ed448GoldScalar([]byte{
		0x1c, 0x51, 0x56, 0x90, 0x17, 0x2d, 0x14, 0x41,
		0x2c, 0x71, 0x8e, 0x1f, 0x1f, 0x2b, 0x38, 0x60,
		0x02, 0x23, 0x42, 0x97, 0xd4, 0x5c, 0x8b, 0x9d,
		0xb9, 0x67, 0xe9, 0x11, 0x9c, 0xe3, 0xbf, 0x14,
		0x99, 0xc2, 0xe1, 0xf3, 0xa1, 0x65, 0xb8, 0x30,
		0xbc, 0x97, 0x8b, 0xa9, 0x98, 0x86, 0x53, 0x6e,
		0x9f, 0x45, 0xbd, 0x44, 0x8a, 0x40, 0x2a, 0x12,
	})

	valid, err := d.verifyDRMessage(testDRMessage.cipher.u11, testDRMessage.cipher.u21, testDRMessage.cipher.v1, alpha1, testSecA)
	c.Assert(valid, Equals, true)
	c.Assert(err, IsNil)

	invalid, err := d.verifyDRMessage(testDRMessage.cipher.u22, testDRMessage.cipher.u21, testDRMessage.cipher.v1, alpha1, testSecA)
	c.Assert(invalid, Equals, false)
	c.Assert(err, ErrorMatches, "cannot decrypt the message")
}
