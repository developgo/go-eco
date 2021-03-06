package main

import (
	"image/color"
)

// Gold palette, 1000 colors, sorted order.
var Gold1000 = color.Palette{
	color.RGBA{0, 0, 0, 255},
	color.RGBA{3, 3, 3, 255},
	color.RGBA{5, 5, 5, 255},
	color.RGBA{6, 6, 6, 255},
	color.RGBA{7, 7, 7, 255},
	color.RGBA{8, 8, 8, 255},
	color.RGBA{10, 9, 9, 255},
	color.RGBA{10, 10, 10, 255},
	color.RGBA{11, 11, 11, 255},
	color.RGBA{12, 12, 12, 255},
	color.RGBA{13, 13, 13, 255},
	color.RGBA{14, 13, 13, 255},
	color.RGBA{15, 14, 14, 255},
	color.RGBA{16, 15, 15, 255},
	color.RGBA{16, 15, 15, 255},
	color.RGBA{17, 16, 16, 255},
	color.RGBA{18, 17, 17, 255},
	color.RGBA{19, 17, 17, 255},
	color.RGBA{19, 18, 18, 255},
	color.RGBA{20, 18, 18, 255},
	color.RGBA{21, 19, 19, 255},
	color.RGBA{21, 20, 20, 255},
	color.RGBA{22, 20, 20, 255},
	color.RGBA{23, 21, 21, 255},
	color.RGBA{23, 21, 21, 255},
	color.RGBA{24, 22, 22, 255},
	color.RGBA{24, 22, 22, 255},
	color.RGBA{25, 23, 23, 255},
	color.RGBA{26, 23, 23, 255},
	color.RGBA{26, 24, 23, 255},
	color.RGBA{27, 24, 24, 255},
	color.RGBA{27, 25, 24, 255},
	color.RGBA{28, 25, 25, 255},
	color.RGBA{28, 26, 25, 255},
	color.RGBA{29, 26, 26, 255},
	color.RGBA{30, 26, 26, 255},
	color.RGBA{30, 27, 27, 255},
	color.RGBA{31, 27, 27, 255},
	color.RGBA{31, 28, 27, 255},
	color.RGBA{32, 28, 28, 255},
	color.RGBA{32, 29, 28, 255},
	color.RGBA{33, 29, 29, 255},
	color.RGBA{33, 29, 29, 255},
	color.RGBA{34, 30, 29, 255},
	color.RGBA{34, 30, 30, 255},
	color.RGBA{35, 30, 30, 255},
	color.RGBA{35, 31, 30, 255},
	color.RGBA{36, 31, 31, 255},
	color.RGBA{36, 32, 31, 255},
	color.RGBA{37, 32, 31, 255},
	color.RGBA{37, 32, 32, 255},
	color.RGBA{38, 33, 32, 255},
	color.RGBA{38, 33, 32, 255},
	color.RGBA{39, 33, 33, 255},
	color.RGBA{39, 34, 33, 255},
	color.RGBA{40, 34, 33, 255},
	color.RGBA{40, 34, 34, 255},
	color.RGBA{40, 35, 34, 255},
	color.RGBA{41, 35, 34, 255},
	color.RGBA{41, 35, 35, 255},
	color.RGBA{42, 36, 35, 255},
	color.RGBA{42, 36, 35, 255},
	color.RGBA{43, 36, 36, 255},
	color.RGBA{43, 37, 36, 255},
	color.RGBA{44, 37, 36, 255},
	color.RGBA{44, 37, 36, 255},
	color.RGBA{44, 38, 37, 255},
	color.RGBA{45, 38, 37, 255},
	color.RGBA{45, 38, 37, 255},
	color.RGBA{46, 39, 38, 255},
	color.RGBA{46, 39, 38, 255},
	color.RGBA{47, 39, 38, 255},
	color.RGBA{47, 40, 38, 255},
	color.RGBA{47, 40, 39, 255},
	color.RGBA{48, 40, 39, 255},
	color.RGBA{48, 40, 39, 255},
	color.RGBA{49, 41, 39, 255},
	color.RGBA{49, 41, 40, 255},
	color.RGBA{49, 41, 40, 255},
	color.RGBA{50, 42, 40, 255},
	color.RGBA{50, 42, 40, 255},
	color.RGBA{51, 42, 41, 255},
	color.RGBA{51, 42, 41, 255},
	color.RGBA{51, 43, 41, 255},
	color.RGBA{52, 43, 41, 255},
	color.RGBA{52, 43, 42, 255},
	color.RGBA{53, 44, 42, 255},
	color.RGBA{53, 44, 42, 255},
	color.RGBA{53, 44, 42, 255},
	color.RGBA{54, 44, 42, 255},
	color.RGBA{54, 45, 43, 255},
	color.RGBA{55, 45, 43, 255},
	color.RGBA{55, 45, 43, 255},
	color.RGBA{55, 45, 43, 255},
	color.RGBA{56, 46, 44, 255},
	color.RGBA{56, 46, 44, 255},
	color.RGBA{57, 46, 44, 255},
	color.RGBA{57, 46, 44, 255},
	color.RGBA{57, 47, 44, 255},
	color.RGBA{58, 47, 45, 255},
	color.RGBA{58, 47, 45, 255},
	color.RGBA{58, 47, 45, 255},
	color.RGBA{59, 48, 45, 255},
	color.RGBA{59, 48, 45, 255},
	color.RGBA{60, 48, 46, 255},
	color.RGBA{60, 48, 46, 255},
	color.RGBA{60, 49, 46, 255},
	color.RGBA{61, 49, 46, 255},
	color.RGBA{61, 49, 46, 255},
	color.RGBA{61, 49, 47, 255},
	color.RGBA{62, 50, 47, 255},
	color.RGBA{62, 50, 47, 255},
	color.RGBA{62, 50, 47, 255},
	color.RGBA{63, 50, 47, 255},
	color.RGBA{63, 50, 48, 255},
	color.RGBA{64, 51, 48, 255},
	color.RGBA{64, 51, 48, 255},
	color.RGBA{64, 51, 48, 255},
	color.RGBA{65, 51, 48, 255},
	color.RGBA{65, 52, 48, 255},
	color.RGBA{65, 52, 49, 255},
	color.RGBA{66, 52, 49, 255},
	color.RGBA{66, 52, 49, 255},
	color.RGBA{66, 53, 49, 255},
	color.RGBA{67, 53, 49, 255},
	color.RGBA{67, 53, 49, 255},
	color.RGBA{67, 53, 50, 255},
	color.RGBA{68, 53, 50, 255},
	color.RGBA{68, 54, 50, 255},
	color.RGBA{68, 54, 50, 255},
	color.RGBA{69, 54, 50, 255},
	color.RGBA{69, 54, 50, 255},
	color.RGBA{69, 54, 51, 255},
	color.RGBA{70, 55, 51, 255},
	color.RGBA{70, 55, 51, 255},
	color.RGBA{70, 55, 51, 255},
	color.RGBA{71, 55, 51, 255},
	color.RGBA{71, 56, 51, 255},
	color.RGBA{71, 56, 51, 255},
	color.RGBA{72, 56, 52, 255},
	color.RGBA{72, 56, 52, 255},
	color.RGBA{72, 56, 52, 255},
	color.RGBA{73, 57, 52, 255},
	color.RGBA{73, 57, 52, 255},
	color.RGBA{73, 57, 52, 255},
	color.RGBA{74, 57, 52, 255},
	color.RGBA{74, 57, 53, 255},
	color.RGBA{74, 58, 53, 255},
	color.RGBA{75, 58, 53, 255},
	color.RGBA{75, 58, 53, 255},
	color.RGBA{75, 58, 53, 255},
	color.RGBA{76, 58, 53, 255},
	color.RGBA{76, 59, 53, 255},
	color.RGBA{76, 59, 53, 255},
	color.RGBA{77, 59, 54, 255},
	color.RGBA{77, 59, 54, 255},
	color.RGBA{77, 59, 54, 255},
	color.RGBA{78, 60, 54, 255},
	color.RGBA{78, 60, 54, 255},
	color.RGBA{78, 60, 54, 255},
	color.RGBA{79, 60, 54, 255},
	color.RGBA{79, 60, 54, 255},
	color.RGBA{79, 60, 55, 255},
	color.RGBA{79, 61, 55, 255},
	color.RGBA{80, 61, 55, 255},
	color.RGBA{80, 61, 55, 255},
	color.RGBA{80, 61, 55, 255},
	color.RGBA{81, 61, 55, 255},
	color.RGBA{81, 62, 55, 255},
	color.RGBA{81, 62, 55, 255},
	color.RGBA{82, 62, 56, 255},
	color.RGBA{82, 62, 56, 255},
	color.RGBA{82, 62, 56, 255},
	color.RGBA{83, 62, 56, 255},
	color.RGBA{83, 63, 56, 255},
	color.RGBA{83, 63, 56, 255},
	color.RGBA{83, 63, 56, 255},
	color.RGBA{84, 63, 56, 255},
	color.RGBA{84, 63, 56, 255},
	color.RGBA{84, 64, 56, 255},
	color.RGBA{85, 64, 57, 255},
	color.RGBA{85, 64, 57, 255},
	color.RGBA{85, 64, 57, 255},
	color.RGBA{86, 64, 57, 255},
	color.RGBA{86, 64, 57, 255},
	color.RGBA{86, 65, 57, 255},
	color.RGBA{87, 65, 57, 255},
	color.RGBA{87, 65, 57, 255},
	color.RGBA{87, 65, 57, 255},
	color.RGBA{87, 65, 57, 255},
	color.RGBA{88, 66, 58, 255},
	color.RGBA{88, 66, 58, 255},
	color.RGBA{88, 66, 58, 255},
	color.RGBA{89, 66, 58, 255},
	color.RGBA{89, 66, 58, 255},
	color.RGBA{89, 66, 58, 255},
	color.RGBA{89, 67, 58, 255},
	color.RGBA{90, 67, 58, 255},
	color.RGBA{90, 67, 58, 255},
	color.RGBA{90, 67, 58, 255},
	color.RGBA{91, 67, 58, 255},
	color.RGBA{91, 67, 59, 255},
	color.RGBA{91, 68, 59, 255},
	color.RGBA{92, 68, 59, 255},
	color.RGBA{92, 68, 59, 255},
	color.RGBA{92, 68, 59, 255},
	color.RGBA{92, 68, 59, 255},
	color.RGBA{93, 68, 59, 255},
	color.RGBA{93, 69, 59, 255},
	color.RGBA{93, 69, 59, 255},
	color.RGBA{94, 69, 59, 255},
	color.RGBA{94, 69, 59, 255},
	color.RGBA{94, 69, 59, 255},
	color.RGBA{94, 69, 59, 255},
	color.RGBA{95, 70, 60, 255},
	color.RGBA{95, 70, 60, 255},
	color.RGBA{95, 70, 60, 255},
	color.RGBA{96, 70, 60, 255},
	color.RGBA{96, 70, 60, 255},
	color.RGBA{96, 70, 60, 255},
	color.RGBA{96, 71, 60, 255},
	color.RGBA{97, 71, 60, 255},
	color.RGBA{97, 71, 60, 255},
	color.RGBA{97, 71, 60, 255},
	color.RGBA{97, 71, 60, 255},
	color.RGBA{98, 71, 60, 255},
	color.RGBA{98, 72, 60, 255},
	color.RGBA{98, 72, 60, 255},
	color.RGBA{99, 72, 60, 255},
	color.RGBA{99, 72, 61, 255},
	color.RGBA{99, 72, 61, 255},
	color.RGBA{99, 72, 61, 255},
	color.RGBA{100, 72, 61, 255},
	color.RGBA{100, 73, 61, 255},
	color.RGBA{100, 73, 61, 255},
	color.RGBA{101, 73, 61, 255},
	color.RGBA{101, 73, 61, 255},
	color.RGBA{101, 73, 61, 255},
	color.RGBA{101, 73, 61, 255},
	color.RGBA{102, 74, 61, 255},
	color.RGBA{102, 74, 61, 255},
	color.RGBA{102, 74, 61, 255},
	color.RGBA{102, 74, 61, 255},
	color.RGBA{103, 74, 61, 255},
	color.RGBA{103, 74, 61, 255},
	color.RGBA{103, 74, 61, 255},
	color.RGBA{104, 75, 62, 255},
	color.RGBA{104, 75, 62, 255},
	color.RGBA{104, 75, 62, 255},
	color.RGBA{104, 75, 62, 255},
	color.RGBA{105, 75, 62, 255},
	color.RGBA{105, 75, 62, 255},
	color.RGBA{105, 76, 62, 255},
	color.RGBA{105, 76, 62, 255},
	color.RGBA{106, 76, 62, 255},
	color.RGBA{106, 76, 62, 255},
	color.RGBA{106, 76, 62, 255},
	color.RGBA{107, 76, 62, 255},
	color.RGBA{107, 76, 62, 255},
	color.RGBA{107, 77, 62, 255},
	color.RGBA{107, 77, 62, 255},
	color.RGBA{108, 77, 62, 255},
	color.RGBA{108, 77, 62, 255},
	color.RGBA{108, 77, 62, 255},
	color.RGBA{108, 77, 62, 255},
	color.RGBA{109, 78, 62, 255},
	color.RGBA{109, 78, 62, 255},
	color.RGBA{109, 78, 62, 255},
	color.RGBA{109, 78, 62, 255},
	color.RGBA{110, 78, 63, 255},
	color.RGBA{110, 78, 63, 255},
	color.RGBA{110, 78, 63, 255},
	color.RGBA{110, 79, 63, 255},
	color.RGBA{111, 79, 63, 255},
	color.RGBA{111, 79, 63, 255},
	color.RGBA{111, 79, 63, 255},
	color.RGBA{112, 79, 63, 255},
	color.RGBA{112, 79, 63, 255},
	color.RGBA{112, 79, 63, 255},
	color.RGBA{112, 80, 63, 255},
	color.RGBA{113, 80, 63, 255},
	color.RGBA{113, 80, 63, 255},
	color.RGBA{113, 80, 63, 255},
	color.RGBA{113, 80, 63, 255},
	color.RGBA{114, 80, 63, 255},
	color.RGBA{114, 80, 63, 255},
	color.RGBA{114, 81, 63, 255},
	color.RGBA{114, 81, 63, 255},
	color.RGBA{115, 81, 63, 255},
	color.RGBA{115, 81, 63, 255},
	color.RGBA{115, 81, 63, 255},
	color.RGBA{115, 81, 63, 255},
	color.RGBA{116, 81, 63, 255},
	color.RGBA{116, 82, 63, 255},
	color.RGBA{116, 82, 63, 255},
	color.RGBA{116, 82, 63, 255},
	color.RGBA{117, 82, 63, 255},
	color.RGBA{117, 82, 63, 255},
	color.RGBA{117, 82, 63, 255},
	color.RGBA{117, 82, 63, 255},
	color.RGBA{118, 83, 63, 255},
	color.RGBA{118, 83, 63, 255},
	color.RGBA{118, 83, 63, 255},
	color.RGBA{118, 83, 63, 255},
	color.RGBA{119, 83, 63, 255},
	color.RGBA{119, 83, 63, 255},
	color.RGBA{119, 83, 63, 255},
	color.RGBA{119, 84, 63, 255},
	color.RGBA{120, 84, 64, 255},
	color.RGBA{120, 84, 64, 255},
	color.RGBA{120, 84, 64, 255},
	color.RGBA{120, 84, 64, 255},
	color.RGBA{121, 84, 64, 255},
	color.RGBA{121, 84, 64, 255},
	color.RGBA{121, 85, 64, 255},
	color.RGBA{121, 85, 64, 255},
	color.RGBA{122, 85, 64, 255},
	color.RGBA{122, 85, 64, 255},
	color.RGBA{122, 85, 64, 255},
	color.RGBA{122, 85, 64, 255},
	color.RGBA{123, 85, 64, 255},
	color.RGBA{123, 86, 64, 255},
	color.RGBA{123, 86, 64, 255},
	color.RGBA{123, 86, 64, 255},
	color.RGBA{124, 86, 64, 255},
	color.RGBA{124, 86, 64, 255},
	color.RGBA{124, 86, 64, 255},
	color.RGBA{124, 86, 64, 255},
	color.RGBA{125, 87, 64, 255},
	color.RGBA{125, 87, 64, 255},
	color.RGBA{125, 87, 64, 255},
	color.RGBA{125, 87, 64, 255},
	color.RGBA{126, 87, 64, 255},
	color.RGBA{126, 87, 64, 255},
	color.RGBA{126, 87, 64, 255},
	color.RGBA{126, 88, 64, 255},
	color.RGBA{127, 88, 64, 255},
	color.RGBA{127, 88, 64, 255},
	color.RGBA{127, 88, 64, 255},
	color.RGBA{127, 88, 64, 255},
	color.RGBA{128, 88, 64, 255},
	color.RGBA{128, 88, 64, 255},
	color.RGBA{128, 89, 64, 255},
	color.RGBA{128, 89, 64, 255},
	color.RGBA{128, 89, 64, 255},
	color.RGBA{129, 89, 64, 255},
	color.RGBA{129, 89, 64, 255},
	color.RGBA{129, 89, 64, 255},
	color.RGBA{129, 89, 64, 255},
	color.RGBA{130, 90, 64, 255},
	color.RGBA{130, 90, 64, 255},
	color.RGBA{130, 90, 64, 255},
	color.RGBA{130, 90, 64, 255},
	color.RGBA{131, 90, 64, 255},
	color.RGBA{131, 90, 64, 255},
	color.RGBA{131, 90, 64, 255},
	color.RGBA{131, 90, 64, 255},
	color.RGBA{132, 91, 64, 255},
	color.RGBA{132, 91, 64, 255},
	color.RGBA{132, 91, 64, 255},
	color.RGBA{132, 91, 64, 255},
	color.RGBA{133, 91, 64, 255},
	color.RGBA{133, 91, 64, 255},
	color.RGBA{133, 91, 64, 255},
	color.RGBA{133, 92, 64, 255},
	color.RGBA{133, 92, 64, 255},
	color.RGBA{134, 92, 64, 255},
	color.RGBA{134, 92, 64, 255},
	color.RGBA{134, 92, 64, 255},
	color.RGBA{134, 92, 64, 255},
	color.RGBA{135, 92, 64, 255},
	color.RGBA{135, 93, 64, 255},
	color.RGBA{135, 93, 64, 255},
	color.RGBA{135, 93, 64, 255},
	color.RGBA{136, 93, 63, 255},
	color.RGBA{136, 93, 63, 255},
	color.RGBA{136, 93, 63, 255},
	color.RGBA{136, 93, 63, 255},
	color.RGBA{136, 94, 63, 255},
	color.RGBA{137, 94, 63, 255},
	color.RGBA{137, 94, 63, 255},
	color.RGBA{137, 94, 63, 255},
	color.RGBA{137, 94, 63, 255},
	color.RGBA{138, 94, 63, 255},
	color.RGBA{138, 94, 63, 255},
	color.RGBA{138, 94, 63, 255},
	color.RGBA{138, 95, 63, 255},
	color.RGBA{139, 95, 63, 255},
	color.RGBA{139, 95, 63, 255},
	color.RGBA{139, 95, 63, 255},
	color.RGBA{139, 95, 63, 255},
	color.RGBA{139, 95, 63, 255},
	color.RGBA{140, 95, 63, 255},
	color.RGBA{140, 96, 63, 255},
	color.RGBA{140, 96, 63, 255},
	color.RGBA{140, 96, 63, 255},
	color.RGBA{141, 96, 63, 255},
	color.RGBA{141, 96, 63, 255},
	color.RGBA{141, 96, 63, 255},
	color.RGBA{141, 96, 63, 255},
	color.RGBA{142, 97, 63, 255},
	color.RGBA{142, 97, 63, 255},
	color.RGBA{142, 97, 63, 255},
	color.RGBA{142, 97, 63, 255},
	color.RGBA{142, 97, 63, 255},
	color.RGBA{143, 97, 63, 255},
	color.RGBA{143, 97, 63, 255},
	color.RGBA{143, 97, 63, 255},
	color.RGBA{143, 98, 63, 255},
	color.RGBA{144, 98, 63, 255},
	color.RGBA{144, 98, 63, 255},
	color.RGBA{144, 98, 63, 255},
	color.RGBA{144, 98, 63, 255},
	color.RGBA{144, 98, 63, 255},
	color.RGBA{145, 98, 63, 255},
	color.RGBA{145, 99, 63, 255},
	color.RGBA{145, 99, 63, 255},
	color.RGBA{145, 99, 62, 255},
	color.RGBA{146, 99, 62, 255},
	color.RGBA{146, 99, 62, 255},
	color.RGBA{146, 99, 62, 255},
	color.RGBA{146, 99, 62, 255},
	color.RGBA{147, 100, 62, 255},
	color.RGBA{147, 100, 62, 255},
	color.RGBA{147, 100, 62, 255},
	color.RGBA{147, 100, 62, 255},
	color.RGBA{147, 100, 62, 255},
	color.RGBA{148, 100, 62, 255},
	color.RGBA{148, 100, 62, 255},
	color.RGBA{148, 101, 62, 255},
	color.RGBA{148, 101, 62, 255},
	color.RGBA{149, 101, 62, 255},
	color.RGBA{149, 101, 62, 255},
	color.RGBA{149, 101, 62, 255},
	color.RGBA{149, 101, 62, 255},
	color.RGBA{149, 101, 62, 255},
	color.RGBA{150, 101, 62, 255},
	color.RGBA{150, 102, 62, 255},
	color.RGBA{150, 102, 62, 255},
	color.RGBA{150, 102, 62, 255},
	color.RGBA{151, 102, 62, 255},
	color.RGBA{151, 102, 62, 255},
	color.RGBA{151, 102, 62, 255},
	color.RGBA{151, 102, 62, 255},
	color.RGBA{151, 103, 62, 255},
	color.RGBA{152, 103, 61, 255},
	color.RGBA{152, 103, 61, 255},
	color.RGBA{152, 103, 61, 255},
	color.RGBA{152, 103, 61, 255},
	color.RGBA{152, 103, 61, 255},
	color.RGBA{153, 103, 61, 255},
	color.RGBA{153, 104, 61, 255},
	color.RGBA{153, 104, 61, 255},
	color.RGBA{153, 104, 61, 255},
	color.RGBA{154, 104, 61, 255},
	color.RGBA{154, 104, 61, 255},
	color.RGBA{154, 104, 61, 255},
	color.RGBA{154, 104, 61, 255},
	color.RGBA{154, 104, 61, 255},
	color.RGBA{155, 105, 61, 255},
	color.RGBA{155, 105, 61, 255},
	color.RGBA{155, 105, 61, 255},
	color.RGBA{155, 105, 61, 255},
	color.RGBA{156, 105, 61, 255},
	color.RGBA{156, 105, 61, 255},
	color.RGBA{156, 105, 61, 255},
	color.RGBA{156, 106, 61, 255},
	color.RGBA{156, 106, 60, 255},
	color.RGBA{157, 106, 60, 255},
	color.RGBA{157, 106, 60, 255},
	color.RGBA{157, 106, 60, 255},
	color.RGBA{157, 106, 60, 255},
	color.RGBA{157, 106, 60, 255},
	color.RGBA{158, 107, 60, 255},
	color.RGBA{158, 107, 60, 255},
	color.RGBA{158, 107, 60, 255},
	color.RGBA{158, 107, 60, 255},
	color.RGBA{159, 107, 60, 255},
	color.RGBA{159, 107, 60, 255},
	color.RGBA{159, 107, 60, 255},
	color.RGBA{159, 108, 60, 255},
	color.RGBA{159, 108, 60, 255},
	color.RGBA{160, 108, 60, 255},
	color.RGBA{160, 108, 60, 255},
	color.RGBA{160, 108, 60, 255},
	color.RGBA{160, 108, 60, 255},
	color.RGBA{160, 108, 59, 255},
	color.RGBA{161, 108, 59, 255},
	color.RGBA{161, 109, 59, 255},
	color.RGBA{161, 109, 59, 255},
	color.RGBA{161, 109, 59, 255},
	color.RGBA{161, 109, 59, 255},
	color.RGBA{162, 109, 59, 255},
	color.RGBA{162, 109, 59, 255},
	color.RGBA{162, 109, 59, 255},
	color.RGBA{162, 110, 59, 255},
	color.RGBA{163, 110, 59, 255},
	color.RGBA{163, 110, 59, 255},
	color.RGBA{163, 110, 59, 255},
	color.RGBA{163, 110, 59, 255},
	color.RGBA{163, 110, 59, 255},
	color.RGBA{164, 110, 59, 255},
	color.RGBA{164, 111, 59, 255},
	color.RGBA{164, 111, 59, 255},
	color.RGBA{164, 111, 58, 255},
	color.RGBA{164, 111, 58, 255},
	color.RGBA{165, 111, 58, 255},
	color.RGBA{165, 111, 58, 255},
	color.RGBA{165, 111, 58, 255},
	color.RGBA{165, 112, 58, 255},
	color.RGBA{165, 112, 58, 255},
	color.RGBA{166, 112, 58, 255},
	color.RGBA{166, 112, 58, 255},
	color.RGBA{166, 112, 58, 255},
	color.RGBA{166, 112, 58, 255},
	color.RGBA{167, 112, 58, 255},
	color.RGBA{167, 113, 58, 255},
	color.RGBA{167, 113, 58, 255},
	color.RGBA{167, 113, 58, 255},
	color.RGBA{167, 113, 58, 255},
	color.RGBA{168, 113, 57, 255},
	color.RGBA{168, 113, 57, 255},
	color.RGBA{168, 113, 57, 255},
	color.RGBA{168, 114, 57, 255},
	color.RGBA{168, 114, 57, 255},
	color.RGBA{169, 114, 57, 255},
	color.RGBA{169, 114, 57, 255},
	color.RGBA{169, 114, 57, 255},
	color.RGBA{169, 114, 57, 255},
	color.RGBA{169, 114, 57, 255},
	color.RGBA{170, 115, 57, 255},
	color.RGBA{170, 115, 57, 255},
	color.RGBA{170, 115, 57, 255},
	color.RGBA{170, 115, 57, 255},
	color.RGBA{170, 115, 57, 255},
	color.RGBA{171, 115, 56, 255},
	color.RGBA{171, 115, 56, 255},
	color.RGBA{171, 116, 56, 255},
	color.RGBA{171, 116, 56, 255},
	color.RGBA{171, 116, 56, 255},
	color.RGBA{172, 116, 56, 255},
	color.RGBA{172, 116, 56, 255},
	color.RGBA{172, 116, 56, 255},
	color.RGBA{172, 116, 56, 255},
	color.RGBA{173, 117, 56, 255},
	color.RGBA{173, 117, 56, 255},
	color.RGBA{173, 117, 56, 255},
	color.RGBA{173, 117, 56, 255},
	color.RGBA{173, 117, 56, 255},
	color.RGBA{174, 117, 55, 255},
	color.RGBA{174, 117, 55, 255},
	color.RGBA{174, 118, 55, 255},
	color.RGBA{174, 118, 55, 255},
	color.RGBA{174, 118, 55, 255},
	color.RGBA{175, 118, 55, 255},
	color.RGBA{175, 118, 55, 255},
	color.RGBA{175, 118, 55, 255},
	color.RGBA{175, 118, 55, 255},
	color.RGBA{175, 119, 55, 255},
	color.RGBA{176, 119, 55, 255},
	color.RGBA{176, 119, 55, 255},
	color.RGBA{176, 119, 55, 255},
	color.RGBA{176, 119, 54, 255},
	color.RGBA{176, 119, 54, 255},
	color.RGBA{177, 119, 54, 255},
	color.RGBA{177, 120, 54, 255},
	color.RGBA{177, 120, 54, 255},
	color.RGBA{177, 120, 54, 255},
	color.RGBA{177, 120, 54, 255},
	color.RGBA{178, 120, 54, 255},
	color.RGBA{178, 120, 54, 255},
	color.RGBA{178, 120, 54, 255},
	color.RGBA{178, 121, 54, 255},
	color.RGBA{178, 121, 54, 255},
	color.RGBA{179, 121, 54, 255},
	color.RGBA{179, 121, 53, 255},
	color.RGBA{179, 121, 53, 255},
	color.RGBA{179, 121, 53, 255},
	color.RGBA{179, 121, 53, 255},
	color.RGBA{180, 122, 53, 255},
	color.RGBA{180, 122, 53, 255},
	color.RGBA{180, 122, 53, 255},
	color.RGBA{180, 122, 53, 255},
	color.RGBA{180, 122, 53, 255},
	color.RGBA{181, 122, 53, 255},
	color.RGBA{181, 123, 53, 255},
	color.RGBA{181, 123, 53, 255},
	color.RGBA{181, 123, 52, 255},
	color.RGBA{181, 123, 52, 255},
	color.RGBA{182, 123, 52, 255},
	color.RGBA{182, 123, 52, 255},
	color.RGBA{182, 123, 52, 255},
	color.RGBA{182, 124, 52, 255},
	color.RGBA{182, 124, 52, 255},
	color.RGBA{183, 124, 52, 255},
	color.RGBA{183, 124, 52, 255},
	color.RGBA{183, 124, 52, 255},
	color.RGBA{183, 124, 52, 255},
	color.RGBA{183, 124, 52, 255},
	color.RGBA{184, 125, 51, 255},
	color.RGBA{184, 125, 51, 255},
	color.RGBA{184, 125, 51, 255},
	color.RGBA{184, 125, 51, 255},
	color.RGBA{184, 125, 51, 255},
	color.RGBA{185, 125, 51, 255},
	color.RGBA{185, 125, 51, 255},
	color.RGBA{185, 126, 51, 255},
	color.RGBA{185, 126, 51, 255},
	color.RGBA{185, 126, 51, 255},
	color.RGBA{185, 126, 51, 255},
	color.RGBA{186, 126, 50, 255},
	color.RGBA{186, 126, 50, 255},
	color.RGBA{186, 127, 50, 255},
	color.RGBA{186, 127, 50, 255},
	color.RGBA{186, 127, 50, 255},
	color.RGBA{187, 127, 50, 255},
	color.RGBA{187, 127, 50, 255},
	color.RGBA{187, 127, 50, 255},
	color.RGBA{187, 127, 50, 255},
	color.RGBA{187, 128, 50, 255},
	color.RGBA{188, 128, 50, 255},
	color.RGBA{188, 128, 49, 255},
	color.RGBA{188, 128, 49, 255},
	color.RGBA{188, 128, 49, 255},
	color.RGBA{188, 128, 49, 255},
	color.RGBA{189, 128, 49, 255},
	color.RGBA{189, 129, 49, 255},
	color.RGBA{189, 129, 49, 255},
	color.RGBA{189, 129, 49, 255},
	color.RGBA{189, 129, 49, 255},
	color.RGBA{190, 129, 49, 255},
	color.RGBA{190, 129, 49, 255},
	color.RGBA{190, 130, 48, 255},
	color.RGBA{190, 130, 48, 255},
	color.RGBA{190, 130, 48, 255},
	color.RGBA{191, 130, 48, 255},
	color.RGBA{191, 130, 48, 255},
	color.RGBA{191, 130, 48, 255},
	color.RGBA{191, 130, 48, 255},
	color.RGBA{191, 131, 48, 255},
	color.RGBA{192, 131, 48, 255},
	color.RGBA{192, 131, 48, 255},
	color.RGBA{192, 131, 47, 255},
	color.RGBA{192, 131, 47, 255},
	color.RGBA{192, 131, 47, 255},
	color.RGBA{192, 132, 47, 255},
	color.RGBA{193, 132, 47, 255},
	color.RGBA{193, 132, 47, 255},
	color.RGBA{193, 132, 47, 255},
	color.RGBA{193, 132, 47, 255},
	color.RGBA{193, 132, 47, 255},
	color.RGBA{194, 132, 47, 255},
	color.RGBA{194, 133, 47, 255},
	color.RGBA{194, 133, 46, 255},
	color.RGBA{194, 133, 46, 255},
	color.RGBA{194, 133, 46, 255},
	color.RGBA{195, 133, 46, 255},
	color.RGBA{195, 133, 46, 255},
	color.RGBA{195, 134, 46, 255},
	color.RGBA{195, 134, 46, 255},
	color.RGBA{195, 134, 46, 255},
	color.RGBA{196, 134, 46, 255},
	color.RGBA{196, 134, 46, 255},
	color.RGBA{196, 134, 45, 255},
	color.RGBA{196, 135, 45, 255},
	color.RGBA{196, 135, 45, 255},
	color.RGBA{196, 135, 45, 255},
	color.RGBA{197, 135, 45, 255},
	color.RGBA{197, 135, 45, 255},
	color.RGBA{197, 135, 45, 255},
	color.RGBA{197, 135, 45, 255},
	color.RGBA{197, 136, 45, 255},
	color.RGBA{198, 136, 44, 255},
	color.RGBA{198, 136, 44, 255},
	color.RGBA{198, 136, 44, 255},
	color.RGBA{198, 136, 44, 255},
	color.RGBA{198, 136, 44, 255},
	color.RGBA{199, 137, 44, 255},
	color.RGBA{199, 137, 44, 255},
	color.RGBA{199, 137, 44, 255},
	color.RGBA{199, 137, 44, 255},
	color.RGBA{199, 137, 44, 255},
	color.RGBA{199, 137, 43, 255},
	color.RGBA{200, 138, 43, 255},
	color.RGBA{200, 138, 43, 255},
	color.RGBA{200, 138, 43, 255},
	color.RGBA{200, 138, 43, 255},
	color.RGBA{200, 138, 43, 255},
	color.RGBA{201, 138, 43, 255},
	color.RGBA{201, 138, 43, 255},
	color.RGBA{201, 139, 43, 255},
	color.RGBA{201, 139, 42, 255},
	color.RGBA{201, 139, 42, 255},
	color.RGBA{202, 139, 42, 255},
	color.RGBA{202, 139, 42, 255},
	color.RGBA{202, 139, 42, 255},
	color.RGBA{202, 140, 42, 255},
	color.RGBA{202, 140, 42, 255},
	color.RGBA{202, 140, 42, 255},
	color.RGBA{203, 140, 42, 255},
	color.RGBA{203, 140, 41, 255},
	color.RGBA{203, 140, 41, 255},
	color.RGBA{203, 141, 41, 255},
	color.RGBA{203, 141, 41, 255},
	color.RGBA{204, 141, 41, 255},
	color.RGBA{204, 141, 41, 255},
	color.RGBA{204, 141, 41, 255},
	color.RGBA{204, 141, 41, 255},
	color.RGBA{204, 142, 41, 255},
	color.RGBA{205, 142, 40, 255},
	color.RGBA{205, 142, 40, 255},
	color.RGBA{205, 142, 40, 255},
	color.RGBA{205, 142, 40, 255},
	color.RGBA{205, 142, 40, 255},
	color.RGBA{205, 143, 40, 255},
	color.RGBA{206, 143, 40, 255},
	color.RGBA{206, 143, 40, 255},
	color.RGBA{206, 143, 40, 255},
	color.RGBA{206, 143, 39, 255},
	color.RGBA{206, 143, 39, 255},
	color.RGBA{207, 144, 39, 255},
	color.RGBA{207, 144, 39, 255},
	color.RGBA{207, 144, 39, 255},
	color.RGBA{207, 144, 39, 255},
	color.RGBA{207, 144, 39, 255},
	color.RGBA{207, 144, 39, 255},
	color.RGBA{208, 145, 39, 255},
	color.RGBA{208, 145, 38, 255},
	color.RGBA{208, 145, 38, 255},
	color.RGBA{208, 145, 38, 255},
	color.RGBA{208, 145, 38, 255},
	color.RGBA{209, 145, 38, 255},
	color.RGBA{209, 146, 38, 255},
	color.RGBA{209, 146, 38, 255},
	color.RGBA{209, 146, 38, 255},
	color.RGBA{209, 146, 37, 255},
	color.RGBA{210, 146, 37, 255},
	color.RGBA{210, 146, 37, 255},
	color.RGBA{210, 147, 37, 255},
	color.RGBA{210, 147, 37, 255},
	color.RGBA{210, 147, 37, 255},
	color.RGBA{210, 147, 37, 255},
	color.RGBA{211, 147, 37, 255},
	color.RGBA{211, 147, 37, 255},
	color.RGBA{211, 148, 36, 255},
	color.RGBA{211, 148, 36, 255},
	color.RGBA{211, 148, 36, 255},
	color.RGBA{212, 148, 36, 255},
	color.RGBA{212, 148, 36, 255},
	color.RGBA{212, 148, 36, 255},
	color.RGBA{212, 149, 36, 255},
	color.RGBA{212, 149, 36, 255},
	color.RGBA{212, 149, 35, 255},
	color.RGBA{213, 149, 35, 255},
	color.RGBA{213, 149, 35, 255},
	color.RGBA{213, 149, 35, 255},
	color.RGBA{213, 150, 35, 255},
	color.RGBA{213, 150, 35, 255},
	color.RGBA{214, 150, 35, 255},
	color.RGBA{214, 150, 35, 255},
	color.RGBA{214, 150, 34, 255},
	color.RGBA{214, 150, 34, 255},
	color.RGBA{214, 151, 34, 255},
	color.RGBA{214, 151, 34, 255},
	color.RGBA{215, 151, 34, 255},
	color.RGBA{215, 151, 34, 255},
	color.RGBA{215, 151, 34, 255},
	color.RGBA{215, 151, 34, 255},
	color.RGBA{215, 152, 34, 255},
	color.RGBA{216, 152, 33, 255},
	color.RGBA{216, 152, 33, 255},
	color.RGBA{216, 152, 33, 255},
	color.RGBA{216, 152, 33, 255},
	color.RGBA{216, 152, 33, 255},
	color.RGBA{216, 153, 33, 255},
	color.RGBA{217, 153, 33, 255},
	color.RGBA{217, 153, 33, 255},
	color.RGBA{217, 153, 32, 255},
	color.RGBA{217, 153, 32, 255},
	color.RGBA{217, 154, 32, 255},
	color.RGBA{217, 154, 32, 255},
	color.RGBA{218, 154, 32, 255},
	color.RGBA{218, 154, 32, 255},
	color.RGBA{218, 154, 32, 255},
	color.RGBA{218, 154, 31, 255},
	color.RGBA{218, 155, 31, 255},
	color.RGBA{219, 155, 31, 255},
	color.RGBA{219, 155, 31, 255},
	color.RGBA{219, 155, 31, 255},
	color.RGBA{219, 155, 31, 255},
	color.RGBA{219, 155, 31, 255},
	color.RGBA{219, 156, 31, 255},
	color.RGBA{220, 156, 30, 255},
	color.RGBA{220, 156, 30, 255},
	color.RGBA{220, 156, 30, 255},
	color.RGBA{220, 156, 30, 255},
	color.RGBA{220, 157, 30, 255},
	color.RGBA{221, 157, 30, 255},
	color.RGBA{221, 157, 30, 255},
	color.RGBA{221, 157, 30, 255},
	color.RGBA{221, 157, 29, 255},
	color.RGBA{221, 157, 29, 255},
	color.RGBA{221, 158, 29, 255},
	color.RGBA{222, 158, 29, 255},
	color.RGBA{222, 158, 29, 255},
	color.RGBA{222, 158, 29, 255},
	color.RGBA{222, 158, 29, 255},
	color.RGBA{222, 158, 29, 255},
	color.RGBA{222, 159, 28, 255},
	color.RGBA{223, 159, 28, 255},
	color.RGBA{223, 159, 28, 255},
	color.RGBA{223, 159, 28, 255},
	color.RGBA{223, 159, 28, 255},
	color.RGBA{223, 160, 28, 255},
	color.RGBA{224, 160, 28, 255},
	color.RGBA{224, 160, 27, 255},
	color.RGBA{224, 160, 27, 255},
	color.RGBA{224, 160, 27, 255},
	color.RGBA{224, 160, 27, 255},
	color.RGBA{224, 161, 27, 255},
	color.RGBA{225, 161, 27, 255},
	color.RGBA{225, 161, 27, 255},
	color.RGBA{225, 161, 27, 255},
	color.RGBA{225, 161, 26, 255},
	color.RGBA{225, 162, 26, 255},
	color.RGBA{225, 162, 26, 255},
	color.RGBA{226, 162, 26, 255},
	color.RGBA{226, 162, 26, 255},
	color.RGBA{226, 162, 26, 255},
	color.RGBA{226, 162, 26, 255},
	color.RGBA{226, 163, 25, 255},
	color.RGBA{227, 163, 25, 255},
	color.RGBA{227, 163, 25, 255},
	color.RGBA{227, 163, 25, 255},
	color.RGBA{227, 163, 25, 255},
	color.RGBA{227, 164, 25, 255},
	color.RGBA{227, 164, 25, 255},
	color.RGBA{228, 164, 24, 255},
	color.RGBA{228, 164, 24, 255},
	color.RGBA{228, 164, 24, 255},
	color.RGBA{228, 164, 24, 255},
	color.RGBA{228, 165, 24, 255},
	color.RGBA{228, 165, 24, 255},
	color.RGBA{229, 165, 24, 255},
	color.RGBA{229, 165, 24, 255},
	color.RGBA{229, 165, 23, 255},
	color.RGBA{229, 166, 23, 255},
	color.RGBA{229, 166, 23, 255},
	color.RGBA{229, 166, 23, 255},
	color.RGBA{230, 166, 23, 255},
	color.RGBA{230, 166, 23, 255},
	color.RGBA{230, 167, 23, 255},
	color.RGBA{230, 167, 22, 255},
	color.RGBA{230, 167, 22, 255},
	color.RGBA{231, 167, 22, 255},
	color.RGBA{231, 167, 22, 255},
	color.RGBA{231, 167, 22, 255},
	color.RGBA{231, 168, 22, 255},
	color.RGBA{231, 168, 22, 255},
	color.RGBA{231, 168, 21, 255},
	color.RGBA{232, 168, 21, 255},
	color.RGBA{232, 168, 21, 255},
	color.RGBA{232, 169, 21, 255},
	color.RGBA{232, 169, 21, 255},
	color.RGBA{232, 169, 21, 255},
	color.RGBA{232, 169, 21, 255},
	color.RGBA{233, 169, 20, 255},
	color.RGBA{233, 170, 20, 255},
	color.RGBA{233, 170, 20, 255},
	color.RGBA{233, 170, 20, 255},
	color.RGBA{233, 170, 20, 255},
	color.RGBA{233, 170, 20, 255},
	color.RGBA{234, 170, 20, 255},
	color.RGBA{234, 171, 19, 255},
	color.RGBA{234, 171, 19, 255},
	color.RGBA{234, 171, 19, 255},
	color.RGBA{234, 171, 19, 255},
	color.RGBA{235, 171, 19, 255},
	color.RGBA{235, 172, 19, 255},
	color.RGBA{235, 172, 19, 255},
	color.RGBA{235, 172, 18, 255},
	color.RGBA{235, 172, 18, 255},
	color.RGBA{235, 172, 18, 255},
	color.RGBA{236, 173, 18, 255},
	color.RGBA{236, 173, 18, 255},
	color.RGBA{236, 173, 18, 255},
	color.RGBA{236, 173, 18, 255},
	color.RGBA{236, 173, 17, 255},
	color.RGBA{236, 174, 17, 255},
	color.RGBA{237, 174, 17, 255},
	color.RGBA{237, 174, 17, 255},
	color.RGBA{237, 174, 17, 255},
	color.RGBA{237, 174, 17, 255},
	color.RGBA{237, 175, 17, 255},
	color.RGBA{237, 175, 16, 255},
	color.RGBA{238, 175, 16, 255},
	color.RGBA{238, 175, 16, 255},
	color.RGBA{238, 175, 16, 255},
	color.RGBA{238, 176, 16, 255},
	color.RGBA{238, 176, 16, 255},
	color.RGBA{238, 176, 15, 255},
	color.RGBA{239, 176, 15, 255},
	color.RGBA{239, 176, 15, 255},
	color.RGBA{239, 176, 15, 255},
	color.RGBA{239, 177, 15, 255},
	color.RGBA{239, 177, 15, 255},
	color.RGBA{239, 177, 15, 255},
	color.RGBA{240, 177, 14, 255},
	color.RGBA{240, 177, 14, 255},
	color.RGBA{240, 178, 14, 255},
	color.RGBA{240, 178, 14, 255},
	color.RGBA{240, 178, 14, 255},
	color.RGBA{240, 178, 14, 255},
	color.RGBA{241, 178, 14, 255},
	color.RGBA{241, 179, 13, 255},
	color.RGBA{241, 179, 13, 255},
	color.RGBA{241, 179, 13, 255},
	color.RGBA{241, 179, 13, 255},
	color.RGBA{242, 179, 13, 255},
	color.RGBA{242, 180, 13, 255},
	color.RGBA{242, 180, 12, 255},
	color.RGBA{242, 180, 12, 255},
	color.RGBA{242, 180, 12, 255},
	color.RGBA{242, 180, 12, 255},
	color.RGBA{243, 181, 12, 255},
	color.RGBA{243, 181, 12, 255},
	color.RGBA{243, 181, 12, 255},
	color.RGBA{243, 181, 11, 255},
	color.RGBA{243, 181, 11, 255},
	color.RGBA{243, 182, 11, 255},
	color.RGBA{244, 182, 11, 255},
	color.RGBA{244, 182, 11, 255},
	color.RGBA{244, 182, 11, 255},
	color.RGBA{244, 182, 10, 255},
	color.RGBA{244, 183, 10, 255},
	color.RGBA{244, 183, 10, 255},
	color.RGBA{245, 183, 10, 255},
	color.RGBA{245, 183, 10, 255},
	color.RGBA{245, 183, 10, 255},
	color.RGBA{245, 184, 10, 255},
	color.RGBA{245, 184, 9, 255},
	color.RGBA{245, 184, 9, 255},
	color.RGBA{246, 184, 9, 255},
	color.RGBA{246, 184, 9, 255},
	color.RGBA{246, 185, 9, 255},
	color.RGBA{246, 185, 9, 255},
	color.RGBA{246, 185, 8, 255},
	color.RGBA{246, 185, 8, 255},
	color.RGBA{247, 186, 8, 255},
	color.RGBA{247, 186, 8, 255},
	color.RGBA{247, 186, 8, 255},
	color.RGBA{247, 186, 8, 255},
	color.RGBA{247, 186, 8, 255},
	color.RGBA{247, 187, 7, 255},
	color.RGBA{248, 187, 7, 255},
	color.RGBA{248, 187, 7, 255},
	color.RGBA{248, 187, 7, 255},
	color.RGBA{248, 187, 7, 255},
	color.RGBA{248, 188, 7, 255},
	color.RGBA{248, 188, 6, 255},
	color.RGBA{249, 188, 6, 255},
	color.RGBA{249, 188, 6, 255},
	color.RGBA{249, 188, 6, 255},
	color.RGBA{249, 189, 6, 255},
	color.RGBA{249, 189, 6, 255},
	color.RGBA{249, 189, 5, 255},
	color.RGBA{250, 189, 5, 255},
	color.RGBA{250, 189, 5, 255},
	color.RGBA{250, 190, 5, 255},
	color.RGBA{250, 190, 5, 255},
	color.RGBA{250, 190, 5, 255},
	color.RGBA{250, 190, 5, 255},
	color.RGBA{251, 190, 4, 255},
	color.RGBA{251, 191, 4, 255},
	color.RGBA{251, 191, 4, 255},
	color.RGBA{251, 191, 4, 255},
	color.RGBA{251, 191, 4, 255},
	color.RGBA{251, 192, 4, 255},
	color.RGBA{252, 192, 3, 255},
	color.RGBA{252, 192, 3, 255},
	color.RGBA{252, 192, 3, 255},
	color.RGBA{252, 192, 3, 255},
	color.RGBA{252, 193, 3, 255},
	color.RGBA{252, 193, 3, 255},
	color.RGBA{253, 193, 2, 255},
	color.RGBA{253, 193, 2, 255},
	color.RGBA{253, 193, 2, 255},
	color.RGBA{253, 194, 2, 255},
	color.RGBA{253, 194, 2, 255},
	color.RGBA{253, 194, 2, 255},
	color.RGBA{254, 194, 1, 255},
	color.RGBA{254, 194, 1, 255},
	color.RGBA{254, 195, 1, 255},
	color.RGBA{254, 195, 1, 255},
	color.RGBA{254, 195, 1, 255},
	color.RGBA{254, 195, 1, 255},
	color.RGBA{255, 196, 0, 255},
	color.RGBA{255, 196, 0, 255},
	color.RGBA{255, 196, 0, 255},
	color.RGBA{255, 196, 0, 255},
}
