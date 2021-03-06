package main

import (
	"image/color"
)

// Heat palette, 256 colors, sorted order.
var Heat256 = color.Palette{
	color.RGBA{0, 0, 0, 255},
	color.RGBA{4, 4, 4, 255},
	color.RGBA{8, 7, 7, 255},
	color.RGBA{11, 10, 10, 255},
	color.RGBA{14, 13, 13, 255},
	color.RGBA{17, 15, 15, 255},
	color.RGBA{19, 18, 18, 255},
	color.RGBA{22, 20, 20, 255},
	color.RGBA{25, 22, 22, 255},
	color.RGBA{27, 24, 24, 255},
	color.RGBA{30, 26, 26, 255},
	color.RGBA{32, 28, 28, 255},
	color.RGBA{35, 30, 30, 255},
	color.RGBA{37, 32, 32, 255},
	color.RGBA{39, 34, 33, 255},
	color.RGBA{42, 35, 35, 255},
	color.RGBA{44, 37, 37, 255},
	color.RGBA{46, 38, 38, 255},
	color.RGBA{49, 40, 39, 255},
	color.RGBA{51, 41, 41, 255},
	color.RGBA{53, 43, 42, 255},
	color.RGBA{56, 44, 43, 255},
	color.RGBA{58, 46, 45, 255},
	color.RGBA{60, 47, 46, 255},
	color.RGBA{62, 48, 47, 255},
	color.RGBA{64, 49, 48, 255},
	color.RGBA{67, 51, 49, 255},
	color.RGBA{69, 52, 50, 255},
	color.RGBA{71, 53, 51, 255},
	color.RGBA{73, 54, 52, 255},
	color.RGBA{75, 55, 53, 255},
	color.RGBA{77, 56, 54, 255},
	color.RGBA{79, 57, 55, 255},
	color.RGBA{81, 58, 55, 255},
	color.RGBA{83, 59, 56, 255},
	color.RGBA{86, 60, 57, 255},
	color.RGBA{88, 61, 58, 255},
	color.RGBA{90, 62, 58, 255},
	color.RGBA{92, 62, 59, 255},
	color.RGBA{94, 63, 59, 255},
	color.RGBA{96, 64, 60, 255},
	color.RGBA{98, 65, 60, 255},
	color.RGBA{100, 65, 61, 255},
	color.RGBA{102, 66, 61, 255},
	color.RGBA{104, 67, 62, 255},
	color.RGBA{106, 67, 62, 255},
	color.RGBA{108, 68, 62, 255},
	color.RGBA{110, 69, 63, 255},
	color.RGBA{112, 69, 63, 255},
	color.RGBA{114, 70, 63, 255},
	color.RGBA{116, 70, 63, 255},
	color.RGBA{118, 71, 63, 255},
	color.RGBA{120, 72, 64, 255},
	color.RGBA{121, 72, 64, 255},
	color.RGBA{123, 73, 64, 255},
	color.RGBA{125, 73, 64, 255},
	color.RGBA{127, 73, 64, 255},
	color.RGBA{129, 74, 64, 255},
	color.RGBA{131, 74, 64, 255},
	color.RGBA{133, 75, 64, 255},
	color.RGBA{135, 75, 64, 255},
	color.RGBA{137, 75, 63, 255},
	color.RGBA{139, 76, 63, 255},
	color.RGBA{141, 76, 63, 255},
	color.RGBA{142, 76, 63, 255},
	color.RGBA{144, 77, 63, 255},
	color.RGBA{146, 77, 62, 255},
	color.RGBA{148, 77, 62, 255},
	color.RGBA{150, 78, 62, 255},
	color.RGBA{152, 78, 61, 255},
	color.RGBA{154, 78, 61, 255},
	color.RGBA{155, 78, 61, 255},
	color.RGBA{157, 79, 60, 255},
	color.RGBA{159, 79, 60, 255},
	color.RGBA{161, 79, 59, 255},
	color.RGBA{163, 79, 59, 255},
	color.RGBA{165, 79, 58, 255},
	color.RGBA{167, 80, 58, 255},
	color.RGBA{168, 80, 57, 255},
	color.RGBA{170, 80, 57, 255},
	color.RGBA{172, 80, 56, 255},
	color.RGBA{174, 80, 55, 255},
	color.RGBA{176, 80, 55, 255},
	color.RGBA{177, 80, 54, 255},
	color.RGBA{179, 80, 53, 255},
	color.RGBA{181, 80, 53, 255},
	color.RGBA{183, 80, 52, 255},
	color.RGBA{185, 81, 51, 255},
	color.RGBA{186, 81, 50, 255},
	color.RGBA{188, 81, 49, 255},
	color.RGBA{190, 81, 48, 255},
	color.RGBA{192, 81, 48, 255},
	color.RGBA{194, 81, 47, 255},
	color.RGBA{195, 81, 46, 255},
	color.RGBA{197, 81, 45, 255},
	color.RGBA{199, 81, 44, 255},
	color.RGBA{201, 81, 43, 255},
	color.RGBA{202, 81, 42, 255},
	color.RGBA{204, 81, 41, 255},
	color.RGBA{206, 81, 40, 255},
	color.RGBA{208, 81, 39, 255},
	color.RGBA{209, 81, 37, 255},
	color.RGBA{211, 81, 36, 255},
	color.RGBA{213, 81, 35, 255},
	color.RGBA{215, 80, 34, 255},
	color.RGBA{216, 80, 33, 255},
	color.RGBA{218, 80, 32, 255},
	color.RGBA{220, 80, 30, 255},
	color.RGBA{222, 80, 29, 255},
	color.RGBA{223, 80, 28, 255},
	color.RGBA{225, 80, 26, 255},
	color.RGBA{227, 80, 25, 255},
	color.RGBA{229, 80, 24, 255},
	color.RGBA{230, 80, 22, 255},
	color.RGBA{232, 80, 21, 255},
	color.RGBA{234, 80, 20, 255},
	color.RGBA{235, 79, 18, 255},
	color.RGBA{237, 79, 17, 255},
	color.RGBA{239, 79, 15, 255},
	color.RGBA{241, 79, 14, 255},
	color.RGBA{242, 79, 12, 255},
	color.RGBA{244, 79, 11, 255},
	color.RGBA{246, 79, 9, 255},
	color.RGBA{247, 79, 7, 255},
	color.RGBA{249, 78, 6, 255},
	color.RGBA{251, 78, 4, 255},
	color.RGBA{252, 78, 3, 255},
	color.RGBA{254, 78, 1, 255},
	color.RGBA{255, 79, 0, 255},
	color.RGBA{255, 80, 0, 255},
	color.RGBA{255, 81, 0, 255},
	color.RGBA{255, 82, 0, 255},
	color.RGBA{255, 83, 0, 255},
	color.RGBA{255, 84, 0, 255},
	color.RGBA{255, 86, 0, 255},
	color.RGBA{255, 87, 0, 255},
	color.RGBA{255, 88, 0, 255},
	color.RGBA{255, 89, 0, 255},
	color.RGBA{255, 90, 0, 255},
	color.RGBA{255, 91, 0, 255},
	color.RGBA{255, 93, 0, 255},
	color.RGBA{255, 94, 0, 255},
	color.RGBA{255, 95, 0, 255},
	color.RGBA{255, 96, 0, 255},
	color.RGBA{255, 97, 0, 255},
	color.RGBA{255, 98, 0, 255},
	color.RGBA{255, 100, 0, 255},
	color.RGBA{255, 101, 0, 255},
	color.RGBA{255, 102, 0, 255},
	color.RGBA{255, 103, 0, 255},
	color.RGBA{255, 104, 0, 255},
	color.RGBA{255, 105, 0, 255},
	color.RGBA{255, 106, 0, 255},
	color.RGBA{255, 108, 0, 255},
	color.RGBA{255, 109, 0, 255},
	color.RGBA{255, 110, 0, 255},
	color.RGBA{255, 111, 0, 255},
	color.RGBA{255, 112, 0, 255},
	color.RGBA{255, 113, 0, 255},
	color.RGBA{255, 115, 0, 255},
	color.RGBA{255, 116, 0, 255},
	color.RGBA{255, 117, 0, 255},
	color.RGBA{255, 118, 0, 255},
	color.RGBA{255, 119, 0, 255},
	color.RGBA{255, 120, 0, 255},
	color.RGBA{255, 122, 0, 255},
	color.RGBA{255, 123, 0, 255},
	color.RGBA{255, 124, 0, 255},
	color.RGBA{255, 125, 0, 255},
	color.RGBA{255, 126, 0, 255},
	color.RGBA{255, 127, 0, 255},
	color.RGBA{255, 129, 0, 255},
	color.RGBA{255, 130, 0, 255},
	color.RGBA{255, 131, 0, 255},
	color.RGBA{255, 132, 0, 255},
	color.RGBA{255, 133, 0, 255},
	color.RGBA{255, 134, 0, 255},
	color.RGBA{255, 136, 0, 255},
	color.RGBA{255, 137, 0, 255},
	color.RGBA{255, 138, 0, 255},
	color.RGBA{255, 139, 0, 255},
	color.RGBA{255, 140, 0, 255},
	color.RGBA{255, 141, 0, 255},
	color.RGBA{255, 143, 0, 255},
	color.RGBA{255, 144, 0, 255},
	color.RGBA{255, 145, 0, 255},
	color.RGBA{255, 146, 0, 255},
	color.RGBA{255, 147, 0, 255},
	color.RGBA{255, 148, 0, 255},
	color.RGBA{255, 149, 0, 255},
	color.RGBA{255, 151, 0, 255},
	color.RGBA{255, 152, 0, 255},
	color.RGBA{255, 153, 0, 255},
	color.RGBA{255, 154, 0, 255},
	color.RGBA{255, 155, 0, 255},
	color.RGBA{255, 156, 0, 255},
	color.RGBA{255, 158, 0, 255},
	color.RGBA{255, 159, 0, 255},
	color.RGBA{255, 160, 0, 255},
	color.RGBA{255, 161, 0, 255},
	color.RGBA{255, 162, 0, 255},
	color.RGBA{255, 163, 0, 255},
	color.RGBA{255, 165, 0, 255},
	color.RGBA{255, 166, 0, 255},
	color.RGBA{255, 167, 0, 255},
	color.RGBA{255, 168, 0, 255},
	color.RGBA{255, 169, 0, 255},
	color.RGBA{255, 170, 0, 255},
	color.RGBA{255, 172, 0, 255},
	color.RGBA{255, 173, 0, 255},
	color.RGBA{255, 174, 0, 255},
	color.RGBA{255, 175, 0, 255},
	color.RGBA{255, 176, 0, 255},
	color.RGBA{255, 177, 0, 255},
	color.RGBA{255, 179, 0, 255},
	color.RGBA{255, 180, 0, 255},
	color.RGBA{255, 181, 0, 255},
	color.RGBA{255, 182, 0, 255},
	color.RGBA{255, 183, 0, 255},
	color.RGBA{255, 184, 0, 255},
	color.RGBA{255, 186, 0, 255},
	color.RGBA{255, 187, 0, 255},
	color.RGBA{255, 188, 0, 255},
	color.RGBA{255, 189, 0, 255},
	color.RGBA{255, 190, 0, 255},
	color.RGBA{255, 191, 0, 255},
	color.RGBA{255, 192, 0, 255},
	color.RGBA{255, 194, 0, 255},
	color.RGBA{255, 195, 0, 255},
	color.RGBA{255, 196, 0, 255},
	color.RGBA{255, 197, 0, 255},
	color.RGBA{255, 198, 0, 255},
	color.RGBA{255, 199, 0, 255},
	color.RGBA{255, 201, 0, 255},
	color.RGBA{255, 202, 0, 255},
	color.RGBA{255, 203, 0, 255},
	color.RGBA{255, 204, 0, 255},
	color.RGBA{255, 205, 0, 255},
	color.RGBA{255, 206, 0, 255},
	color.RGBA{255, 208, 0, 255},
	color.RGBA{255, 209, 0, 255},
	color.RGBA{255, 210, 0, 255},
	color.RGBA{255, 211, 0, 255},
	color.RGBA{255, 212, 0, 255},
	color.RGBA{255, 213, 0, 255},
	color.RGBA{255, 215, 0, 255},
	color.RGBA{255, 216, 0, 255},
	color.RGBA{255, 217, 0, 255},
	color.RGBA{255, 218, 0, 255},
	color.RGBA{255, 219, 0, 255},
	color.RGBA{255, 220, 0, 255},
	color.RGBA{255, 222, 0, 255},
	color.RGBA{255, 223, 0, 255},
	color.RGBA{255, 224, 0, 255},
	color.RGBA{255, 225, 0, 255},
	color.RGBA{255, 226, 0, 255},
}
