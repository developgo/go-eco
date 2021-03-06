package main

import (
	"image/color"
)

// Hypsometric palette, 1000 colors, sorted order.
var Hypso1000 = color.Palette{
	color.RGBA{81, 17, 126, 255},
	color.RGBA{80, 17, 127, 255},
	color.RGBA{80, 17, 127, 255},
	color.RGBA{79, 16, 128, 255},
	color.RGBA{78, 16, 129, 255},
	color.RGBA{78, 16, 129, 255},
	color.RGBA{77, 16, 130, 255},
	color.RGBA{77, 16, 131, 255},
	color.RGBA{76, 16, 132, 255},
	color.RGBA{75, 16, 132, 255},
	color.RGBA{75, 15, 133, 255},
	color.RGBA{74, 15, 134, 255},
	color.RGBA{73, 15, 134, 255},
	color.RGBA{73, 15, 135, 255},
	color.RGBA{72, 15, 136, 255},
	color.RGBA{71, 15, 136, 255},
	color.RGBA{71, 15, 137, 255},
	color.RGBA{70, 15, 138, 255},
	color.RGBA{69, 14, 138, 255},
	color.RGBA{69, 14, 139, 255},
	color.RGBA{68, 14, 140, 255},
	color.RGBA{67, 14, 141, 255},
	color.RGBA{67, 14, 141, 255},
	color.RGBA{66, 14, 142, 255},
	color.RGBA{66, 14, 143, 255},
	color.RGBA{65, 13, 143, 255},
	color.RGBA{64, 13, 144, 255},
	color.RGBA{64, 13, 145, 255},
	color.RGBA{63, 13, 145, 255},
	color.RGBA{62, 13, 146, 255},
	color.RGBA{62, 13, 147, 255},
	color.RGBA{61, 13, 147, 255},
	color.RGBA{60, 13, 148, 255},
	color.RGBA{60, 12, 149, 255},
	color.RGBA{59, 12, 150, 255},
	color.RGBA{58, 12, 150, 255},
	color.RGBA{58, 12, 151, 255},
	color.RGBA{57, 12, 152, 255},
	color.RGBA{57, 12, 152, 255},
	color.RGBA{56, 12, 153, 255},
	color.RGBA{55, 11, 154, 255},
	color.RGBA{55, 11, 154, 255},
	color.RGBA{54, 11, 155, 255},
	color.RGBA{53, 11, 156, 255},
	color.RGBA{53, 11, 156, 255},
	color.RGBA{52, 11, 157, 255},
	color.RGBA{51, 11, 158, 255},
	color.RGBA{51, 11, 159, 255},
	color.RGBA{50, 10, 159, 255},
	color.RGBA{49, 10, 160, 255},
	color.RGBA{49, 10, 161, 255},
	color.RGBA{48, 10, 161, 255},
	color.RGBA{47, 10, 162, 255},
	color.RGBA{47, 10, 163, 255},
	color.RGBA{46, 10, 163, 255},
	color.RGBA{46, 9, 164, 255},
	color.RGBA{45, 9, 165, 255},
	color.RGBA{44, 9, 166, 255},
	color.RGBA{44, 9, 166, 255},
	color.RGBA{43, 9, 167, 255},
	color.RGBA{42, 9, 168, 255},
	color.RGBA{42, 9, 168, 255},
	color.RGBA{41, 9, 169, 255},
	color.RGBA{40, 8, 170, 255},
	color.RGBA{40, 8, 170, 255},
	color.RGBA{39, 8, 171, 255},
	color.RGBA{39, 8, 171, 255},
	color.RGBA{39, 8, 172, 255},
	color.RGBA{38, 8, 172, 255},
	color.RGBA{38, 8, 173, 255},
	color.RGBA{37, 8, 173, 255},
	color.RGBA{37, 8, 174, 255},
	color.RGBA{36, 8, 174, 255},
	color.RGBA{36, 7, 175, 255},
	color.RGBA{35, 7, 175, 255},
	color.RGBA{35, 7, 175, 255},
	color.RGBA{34, 7, 176, 255},
	color.RGBA{34, 7, 176, 255},
	color.RGBA{34, 7, 177, 255},
	color.RGBA{33, 7, 177, 255},
	color.RGBA{33, 7, 178, 255},
	color.RGBA{32, 7, 178, 255},
	color.RGBA{32, 7, 179, 255},
	color.RGBA{31, 6, 179, 255},
	color.RGBA{31, 6, 180, 255},
	color.RGBA{30, 6, 180, 255},
	color.RGBA{30, 6, 181, 255},
	color.RGBA{29, 6, 181, 255},
	color.RGBA{29, 6, 182, 255},
	color.RGBA{29, 6, 182, 255},
	color.RGBA{28, 6, 183, 255},
	color.RGBA{28, 6, 183, 255},
	color.RGBA{27, 6, 184, 255},
	color.RGBA{27, 6, 184, 255},
	color.RGBA{26, 5, 185, 255},
	color.RGBA{26, 5, 185, 255},
	color.RGBA{25, 5, 186, 255},
	color.RGBA{25, 5, 186, 255},
	color.RGBA{24, 5, 187, 255},
	color.RGBA{24, 5, 187, 255},
	color.RGBA{24, 5, 188, 255},
	color.RGBA{23, 5, 188, 255},
	color.RGBA{23, 5, 189, 255},
	color.RGBA{22, 5, 189, 255},
	color.RGBA{22, 5, 190, 255},
	color.RGBA{21, 4, 190, 255},
	color.RGBA{21, 4, 191, 255},
	color.RGBA{20, 4, 191, 255},
	color.RGBA{20, 4, 192, 255},
	color.RGBA{19, 4, 192, 255},
	color.RGBA{19, 4, 193, 255},
	color.RGBA{19, 4, 193, 255},
	color.RGBA{18, 4, 194, 255},
	color.RGBA{18, 4, 194, 255},
	color.RGBA{17, 4, 195, 255},
	color.RGBA{17, 3, 195, 255},
	color.RGBA{16, 3, 196, 255},
	color.RGBA{16, 3, 196, 255},
	color.RGBA{15, 3, 197, 255},
	color.RGBA{15, 3, 197, 255},
	color.RGBA{14, 3, 198, 255},
	color.RGBA{14, 3, 198, 255},
	color.RGBA{14, 3, 198, 255},
	color.RGBA{13, 3, 199, 255},
	color.RGBA{13, 3, 199, 255},
	color.RGBA{12, 3, 200, 255},
	color.RGBA{12, 2, 200, 255},
	color.RGBA{11, 2, 201, 255},
	color.RGBA{11, 2, 201, 255},
	color.RGBA{10, 2, 202, 255},
	color.RGBA{10, 2, 202, 255},
	color.RGBA{9, 2, 203, 255},
	color.RGBA{9, 2, 203, 255},
	color.RGBA{9, 2, 204, 255},
	color.RGBA{8, 2, 204, 255},
	color.RGBA{8, 2, 205, 255},
	color.RGBA{7, 1, 205, 255},
	color.RGBA{7, 1, 206, 255},
	color.RGBA{6, 1, 206, 255},
	color.RGBA{6, 1, 207, 255},
	color.RGBA{5, 1, 207, 255},
	color.RGBA{5, 1, 208, 255},
	color.RGBA{4, 1, 208, 255},
	color.RGBA{4, 1, 209, 255},
	color.RGBA{3, 1, 209, 255},
	color.RGBA{3, 1, 210, 255},
	color.RGBA{3, 1, 210, 255},
	color.RGBA{2, 0, 211, 255},
	color.RGBA{2, 0, 211, 255},
	color.RGBA{1, 0, 212, 255},
	color.RGBA{1, 0, 212, 255},
	color.RGBA{0, 0, 213, 255},
	color.RGBA{24, 25, 207, 255},
	color.RGBA{24, 26, 206, 255},
	color.RGBA{24, 28, 206, 255},
	color.RGBA{24, 29, 206, 255},
	color.RGBA{24, 31, 206, 255},
	color.RGBA{23, 32, 206, 255},
	color.RGBA{23, 33, 206, 255},
	color.RGBA{23, 35, 206, 255},
	color.RGBA{23, 36, 206, 255},
	color.RGBA{23, 38, 206, 255},
	color.RGBA{23, 39, 206, 255},
	color.RGBA{22, 41, 206, 255},
	color.RGBA{22, 42, 206, 255},
	color.RGBA{22, 43, 206, 255},
	color.RGBA{22, 45, 206, 255},
	color.RGBA{22, 46, 205, 255},
	color.RGBA{22, 48, 205, 255},
	color.RGBA{22, 49, 205, 255},
	color.RGBA{21, 51, 205, 255},
	color.RGBA{21, 52, 205, 255},
	color.RGBA{21, 53, 205, 255},
	color.RGBA{21, 55, 205, 255},
	color.RGBA{21, 56, 205, 255},
	color.RGBA{21, 58, 205, 255},
	color.RGBA{20, 59, 205, 255},
	color.RGBA{20, 61, 205, 255},
	color.RGBA{20, 62, 205, 255},
	color.RGBA{20, 63, 205, 255},
	color.RGBA{20, 65, 205, 255},
	color.RGBA{20, 66, 204, 255},
	color.RGBA{19, 68, 204, 255},
	color.RGBA{19, 69, 204, 255},
	color.RGBA{19, 71, 204, 255},
	color.RGBA{19, 72, 204, 255},
	color.RGBA{19, 73, 204, 255},
	color.RGBA{19, 75, 204, 255},
	color.RGBA{18, 76, 204, 255},
	color.RGBA{18, 78, 204, 255},
	color.RGBA{18, 79, 204, 255},
	color.RGBA{18, 81, 204, 255},
	color.RGBA{18, 82, 204, 255},
	color.RGBA{18, 83, 204, 255},
	color.RGBA{17, 85, 204, 255},
	color.RGBA{17, 86, 203, 255},
	color.RGBA{17, 88, 203, 255},
	color.RGBA{17, 89, 203, 255},
	color.RGBA{17, 91, 203, 255},
	color.RGBA{17, 92, 203, 255},
	color.RGBA{16, 94, 203, 255},
	color.RGBA{16, 95, 203, 255},
	color.RGBA{16, 96, 203, 255},
	color.RGBA{16, 98, 203, 255},
	color.RGBA{16, 99, 203, 255},
	color.RGBA{16, 101, 203, 255},
	color.RGBA{15, 102, 203, 255},
	color.RGBA{15, 104, 203, 255},
	color.RGBA{15, 105, 203, 255},
	color.RGBA{15, 106, 202, 255},
	color.RGBA{15, 108, 202, 255},
	color.RGBA{15, 109, 202, 255},
	color.RGBA{14, 111, 202, 255},
	color.RGBA{14, 112, 202, 255},
	color.RGBA{14, 114, 202, 255},
	color.RGBA{14, 115, 202, 255},
	color.RGBA{14, 116, 202, 255},
	color.RGBA{14, 118, 202, 255},
	color.RGBA{13, 119, 202, 255},
	color.RGBA{13, 121, 202, 255},
	color.RGBA{13, 122, 202, 255},
	color.RGBA{13, 124, 202, 255},
	color.RGBA{13, 125, 202, 255},
	color.RGBA{13, 126, 201, 255},
	color.RGBA{12, 128, 201, 255},
	color.RGBA{12, 129, 201, 255},
	color.RGBA{12, 130, 201, 255},
	color.RGBA{12, 131, 201, 255},
	color.RGBA{12, 132, 201, 255},
	color.RGBA{12, 133, 201, 255},
	color.RGBA{12, 134, 201, 255},
	color.RGBA{12, 135, 201, 255},
	color.RGBA{11, 136, 201, 255},
	color.RGBA{11, 137, 201, 255},
	color.RGBA{11, 137, 201, 255},
	color.RGBA{11, 138, 201, 255},
	color.RGBA{11, 139, 201, 255},
	color.RGBA{11, 140, 201, 255},
	color.RGBA{11, 141, 201, 255},
	color.RGBA{11, 142, 201, 255},
	color.RGBA{11, 143, 201, 255},
	color.RGBA{11, 144, 201, 255},
	color.RGBA{10, 145, 201, 255},
	color.RGBA{10, 146, 200, 255},
	color.RGBA{10, 146, 200, 255},
	color.RGBA{10, 147, 200, 255},
	color.RGBA{10, 148, 200, 255},
	color.RGBA{10, 149, 200, 255},
	color.RGBA{10, 150, 200, 255},
	color.RGBA{10, 151, 200, 255},
	color.RGBA{10, 152, 200, 255},
	color.RGBA{10, 153, 200, 255},
	color.RGBA{9, 154, 200, 255},
	color.RGBA{9, 154, 200, 255},
	color.RGBA{9, 155, 200, 255},
	color.RGBA{9, 156, 200, 255},
	color.RGBA{9, 157, 200, 255},
	color.RGBA{9, 158, 200, 255},
	color.RGBA{9, 159, 200, 255},
	color.RGBA{9, 160, 200, 255},
	color.RGBA{9, 161, 200, 255},
	color.RGBA{8, 162, 200, 255},
	color.RGBA{8, 162, 200, 255},
	color.RGBA{8, 163, 200, 255},
	color.RGBA{8, 164, 200, 255},
	color.RGBA{8, 165, 199, 255},
	color.RGBA{8, 166, 199, 255},
	color.RGBA{8, 167, 199, 255},
	color.RGBA{8, 168, 199, 255},
	color.RGBA{8, 169, 199, 255},
	color.RGBA{8, 170, 199, 255},
	color.RGBA{7, 170, 199, 255},
	color.RGBA{7, 171, 199, 255},
	color.RGBA{7, 172, 199, 255},
	color.RGBA{7, 173, 199, 255},
	color.RGBA{7, 174, 199, 255},
	color.RGBA{7, 175, 199, 255},
	color.RGBA{7, 176, 199, 255},
	color.RGBA{7, 177, 199, 255},
	color.RGBA{7, 178, 199, 255},
	color.RGBA{7, 179, 199, 255},
	color.RGBA{6, 179, 199, 255},
	color.RGBA{6, 180, 199, 255},
	color.RGBA{6, 181, 199, 255},
	color.RGBA{6, 182, 199, 255},
	color.RGBA{6, 183, 199, 255},
	color.RGBA{6, 184, 199, 255},
	color.RGBA{6, 185, 199, 255},
	color.RGBA{6, 186, 198, 255},
	color.RGBA{6, 187, 198, 255},
	color.RGBA{6, 187, 198, 255},
	color.RGBA{5, 188, 198, 255},
	color.RGBA{5, 189, 198, 255},
	color.RGBA{5, 190, 198, 255},
	color.RGBA{5, 191, 198, 255},
	color.RGBA{5, 192, 198, 255},
	color.RGBA{5, 193, 198, 255},
	color.RGBA{5, 194, 198, 255},
	color.RGBA{5, 195, 198, 255},
	color.RGBA{5, 195, 198, 255},
	color.RGBA{4, 196, 198, 255},
	color.RGBA{4, 197, 198, 255},
	color.RGBA{4, 198, 198, 255},
	color.RGBA{4, 199, 198, 255},
	color.RGBA{4, 200, 198, 255},
	color.RGBA{4, 201, 198, 255},
	color.RGBA{4, 202, 198, 255},
	color.RGBA{4, 203, 198, 255},
	color.RGBA{4, 204, 198, 255},
	color.RGBA{4, 204, 198, 255},
	color.RGBA{3, 205, 197, 255},
	color.RGBA{3, 206, 197, 255},
	color.RGBA{3, 207, 197, 255},
	color.RGBA{3, 208, 197, 255},
	color.RGBA{3, 209, 197, 255},
	color.RGBA{3, 210, 197, 255},
	color.RGBA{3, 211, 197, 255},
	color.RGBA{3, 212, 197, 255},
	color.RGBA{3, 212, 197, 255},
	color.RGBA{3, 213, 197, 255},
	color.RGBA{2, 214, 197, 255},
	color.RGBA{2, 215, 197, 255},
	color.RGBA{2, 216, 197, 255},
	color.RGBA{2, 217, 197, 255},
	color.RGBA{2, 218, 197, 255},
	color.RGBA{2, 219, 197, 255},
	color.RGBA{2, 220, 197, 255},
	color.RGBA{2, 220, 197, 255},
	color.RGBA{2, 221, 197, 255},
	color.RGBA{1, 222, 197, 255},
	color.RGBA{1, 223, 197, 255},
	color.RGBA{1, 224, 197, 255},
	color.RGBA{1, 225, 197, 255},
	color.RGBA{1, 226, 196, 255},
	color.RGBA{1, 227, 196, 255},
	color.RGBA{1, 228, 196, 255},
	color.RGBA{1, 229, 196, 255},
	color.RGBA{1, 229, 196, 255},
	color.RGBA{1, 230, 196, 255},
	color.RGBA{0, 231, 196, 255},
	color.RGBA{0, 232, 196, 255},
	color.RGBA{0, 233, 196, 255},
	color.RGBA{0, 234, 196, 255},
	color.RGBA{0, 235, 196, 255},
	color.RGBA{0, 235, 196, 255},
	color.RGBA{0, 235, 195, 255},
	color.RGBA{0, 235, 194, 255},
	color.RGBA{0, 235, 194, 255},
	color.RGBA{0, 235, 193, 255},
	color.RGBA{0, 235, 192, 255},
	color.RGBA{0, 235, 192, 255},
	color.RGBA{0, 235, 191, 255},
	color.RGBA{0, 235, 190, 255},
	color.RGBA{0, 235, 190, 255},
	color.RGBA{0, 235, 189, 255},
	color.RGBA{0, 235, 188, 255},
	color.RGBA{0, 235, 187, 255},
	color.RGBA{0, 235, 187, 255},
	color.RGBA{0, 235, 186, 255},
	color.RGBA{0, 235, 185, 255},
	color.RGBA{0, 235, 185, 255},
	color.RGBA{0, 235, 184, 255},
	color.RGBA{0, 235, 183, 255},
	color.RGBA{0, 235, 183, 255},
	color.RGBA{0, 235, 182, 255},
	color.RGBA{0, 235, 181, 255},
	color.RGBA{0, 235, 180, 255},
	color.RGBA{0, 235, 180, 255},
	color.RGBA{0, 235, 179, 255},
	color.RGBA{0, 234, 178, 255},
	color.RGBA{0, 234, 178, 255},
	color.RGBA{0, 234, 177, 255},
	color.RGBA{0, 234, 176, 255},
	color.RGBA{0, 234, 176, 255},
	color.RGBA{0, 234, 175, 255},
	color.RGBA{0, 234, 174, 255},
	color.RGBA{0, 234, 173, 255},
	color.RGBA{0, 234, 173, 255},
	color.RGBA{0, 234, 172, 255},
	color.RGBA{0, 234, 171, 255},
	color.RGBA{0, 234, 171, 255},
	color.RGBA{0, 234, 170, 255},
	color.RGBA{0, 234, 169, 255},
	color.RGBA{0, 234, 169, 255},
	color.RGBA{0, 234, 168, 255},
	color.RGBA{0, 234, 167, 255},
	color.RGBA{0, 234, 167, 255},
	color.RGBA{0, 234, 166, 255},
	color.RGBA{0, 234, 165, 255},
	color.RGBA{0, 234, 164, 255},
	color.RGBA{0, 234, 164, 255},
	color.RGBA{0, 234, 163, 255},
	color.RGBA{0, 234, 162, 255},
	color.RGBA{0, 234, 162, 255},
	color.RGBA{0, 234, 161, 255},
	color.RGBA{0, 234, 160, 255},
	color.RGBA{0, 234, 160, 255},
	color.RGBA{0, 234, 159, 255},
	color.RGBA{0, 234, 158, 255},
	color.RGBA{0, 233, 157, 255},
	color.RGBA{0, 233, 157, 255},
	color.RGBA{0, 233, 156, 255},
	color.RGBA{0, 233, 155, 255},
	color.RGBA{0, 233, 155, 255},
	color.RGBA{0, 233, 154, 255},
	color.RGBA{0, 233, 153, 255},
	color.RGBA{0, 233, 153, 255},
	color.RGBA{0, 233, 152, 255},
	color.RGBA{0, 233, 151, 255},
	color.RGBA{0, 233, 150, 255},
	color.RGBA{0, 233, 150, 255},
	color.RGBA{0, 233, 149, 255},
	color.RGBA{0, 233, 148, 255},
	color.RGBA{0, 233, 148, 255},
	color.RGBA{0, 233, 147, 255},
	color.RGBA{0, 233, 146, 255},
	color.RGBA{0, 233, 146, 255},
	color.RGBA{0, 233, 145, 255},
	color.RGBA{0, 233, 144, 255},
	color.RGBA{0, 233, 144, 255},
	color.RGBA{0, 233, 143, 255},
	color.RGBA{0, 233, 142, 255},
	color.RGBA{0, 233, 141, 255},
	color.RGBA{0, 233, 141, 255},
	color.RGBA{0, 233, 140, 255},
	color.RGBA{0, 233, 139, 255},
	color.RGBA{0, 233, 139, 255},
	color.RGBA{0, 233, 138, 255},
	color.RGBA{0, 232, 137, 255},
	color.RGBA{0, 232, 137, 255},
	color.RGBA{0, 232, 136, 255},
	color.RGBA{0, 232, 135, 255},
	color.RGBA{0, 232, 134, 255},
	color.RGBA{0, 232, 134, 255},
	color.RGBA{0, 232, 133, 255},
	color.RGBA{0, 232, 132, 255},
	color.RGBA{0, 232, 132, 255},
	color.RGBA{0, 232, 131, 255},
	color.RGBA{0, 232, 130, 255},
	color.RGBA{0, 232, 130, 255},
	color.RGBA{0, 232, 129, 255},
	color.RGBA{0, 232, 128, 255},
	color.RGBA{0, 232, 127, 255},
	color.RGBA{0, 232, 127, 255},
	color.RGBA{0, 232, 126, 255},
	color.RGBA{0, 232, 125, 255},
	color.RGBA{0, 232, 125, 255},
	color.RGBA{0, 232, 124, 255},
	color.RGBA{0, 232, 123, 255},
	color.RGBA{0, 232, 123, 255},
	color.RGBA{0, 232, 122, 255},
	color.RGBA{0, 232, 121, 255},
	color.RGBA{0, 232, 121, 255},
	color.RGBA{0, 232, 120, 255},
	color.RGBA{0, 232, 119, 255},
	color.RGBA{0, 232, 118, 255},
	color.RGBA{0, 232, 118, 255},
	color.RGBA{0, 232, 117, 255},
	color.RGBA{0, 231, 116, 255},
	color.RGBA{0, 231, 116, 255},
	color.RGBA{0, 231, 115, 255},
	color.RGBA{0, 231, 114, 255},
	color.RGBA{0, 231, 114, 255},
	color.RGBA{0, 231, 113, 255},
	color.RGBA{0, 231, 112, 255},
	color.RGBA{0, 231, 111, 255},
	color.RGBA{0, 231, 111, 255},
	color.RGBA{0, 231, 110, 255},
	color.RGBA{0, 231, 109, 255},
	color.RGBA{0, 231, 109, 255},
	color.RGBA{0, 231, 108, 255},
	color.RGBA{0, 231, 107, 255},
	color.RGBA{0, 231, 107, 255},
	color.RGBA{0, 231, 106, 255},
	color.RGBA{0, 231, 105, 255},
	color.RGBA{0, 231, 104, 255},
	color.RGBA{0, 231, 104, 255},
	color.RGBA{0, 231, 103, 255},
	color.RGBA{0, 231, 102, 255},
	color.RGBA{0, 231, 102, 255},
	color.RGBA{0, 231, 101, 255},
	color.RGBA{0, 231, 100, 255},
	color.RGBA{0, 231, 99, 255},
	color.RGBA{0, 231, 98, 255},
	color.RGBA{0, 231, 98, 255},
	color.RGBA{0, 231, 97, 255},
	color.RGBA{0, 230, 96, 255},
	color.RGBA{0, 230, 95, 255},
	color.RGBA{0, 230, 94, 255},
	color.RGBA{0, 230, 93, 255},
	color.RGBA{0, 230, 92, 255},
	color.RGBA{0, 230, 91, 255},
	color.RGBA{0, 230, 90, 255},
	color.RGBA{0, 230, 89, 255},
	color.RGBA{0, 230, 88, 255},
	color.RGBA{0, 230, 87, 255},
	color.RGBA{0, 230, 87, 255},
	color.RGBA{0, 230, 86, 255},
	color.RGBA{0, 230, 85, 255},
	color.RGBA{0, 230, 84, 255},
	color.RGBA{0, 230, 83, 255},
	color.RGBA{0, 230, 82, 255},
	color.RGBA{0, 230, 81, 255},
	color.RGBA{0, 230, 80, 255},
	color.RGBA{0, 230, 79, 255},
	color.RGBA{0, 230, 78, 255},
	color.RGBA{0, 230, 77, 255},
	color.RGBA{0, 230, 76, 255},
	color.RGBA{0, 229, 76, 255},
	color.RGBA{0, 229, 75, 255},
	color.RGBA{0, 229, 74, 255},
	color.RGBA{0, 229, 73, 255},
	color.RGBA{0, 229, 72, 255},
	color.RGBA{0, 229, 71, 255},
	color.RGBA{0, 229, 70, 255},
	color.RGBA{0, 229, 69, 255},
	color.RGBA{0, 229, 68, 255},
	color.RGBA{0, 229, 67, 255},
	color.RGBA{0, 229, 66, 255},
	color.RGBA{0, 229, 65, 255},
	color.RGBA{0, 229, 64, 255},
	color.RGBA{0, 229, 64, 255},
	color.RGBA{0, 229, 63, 255},
	color.RGBA{0, 229, 62, 255},
	color.RGBA{0, 229, 61, 255},
	color.RGBA{0, 229, 60, 255},
	color.RGBA{0, 229, 59, 255},
	color.RGBA{0, 229, 58, 255},
	color.RGBA{0, 229, 57, 255},
	color.RGBA{0, 229, 56, 255},
	color.RGBA{0, 228, 55, 255},
	color.RGBA{0, 228, 54, 255},
	color.RGBA{0, 228, 53, 255},
	color.RGBA{0, 228, 53, 255},
	color.RGBA{0, 228, 52, 255},
	color.RGBA{0, 228, 51, 255},
	color.RGBA{0, 228, 50, 255},
	color.RGBA{0, 228, 49, 255},
	color.RGBA{0, 228, 48, 255},
	color.RGBA{0, 228, 47, 255},
	color.RGBA{0, 228, 46, 255},
	color.RGBA{0, 228, 45, 255},
	color.RGBA{0, 228, 44, 255},
	color.RGBA{0, 228, 43, 255},
	color.RGBA{0, 228, 42, 255},
	color.RGBA{0, 228, 42, 255},
	color.RGBA{0, 228, 41, 255},
	color.RGBA{0, 228, 40, 255},
	color.RGBA{0, 228, 39, 255},
	color.RGBA{0, 228, 38, 255},
	color.RGBA{0, 228, 37, 255},
	color.RGBA{0, 228, 36, 255},
	color.RGBA{0, 228, 35, 255},
	color.RGBA{0, 227, 34, 255},
	color.RGBA{0, 227, 33, 255},
	color.RGBA{0, 227, 32, 255},
	color.RGBA{0, 227, 31, 255},
	color.RGBA{0, 227, 30, 255},
	color.RGBA{0, 227, 30, 255},
	color.RGBA{0, 227, 29, 255},
	color.RGBA{0, 227, 28, 255},
	color.RGBA{0, 227, 27, 255},
	color.RGBA{0, 227, 26, 255},
	color.RGBA{0, 227, 25, 255},
	color.RGBA{0, 227, 24, 255},
	color.RGBA{0, 227, 23, 255},
	color.RGBA{0, 227, 22, 255},
	color.RGBA{0, 227, 21, 255},
	color.RGBA{0, 227, 20, 255},
	color.RGBA{0, 227, 19, 255},
	color.RGBA{0, 227, 19, 255},
	color.RGBA{0, 227, 18, 255},
	color.RGBA{0, 227, 17, 255},
	color.RGBA{0, 227, 16, 255},
	color.RGBA{0, 227, 15, 255},
	color.RGBA{0, 226, 14, 255},
	color.RGBA{0, 226, 13, 255},
	color.RGBA{0, 226, 12, 255},
	color.RGBA{0, 226, 11, 255},
	color.RGBA{0, 226, 10, 255},
	color.RGBA{0, 226, 9, 255},
	color.RGBA{0, 226, 8, 255},
	color.RGBA{0, 226, 7, 255},
	color.RGBA{0, 226, 7, 255},
	color.RGBA{0, 226, 6, 255},
	color.RGBA{0, 226, 5, 255},
	color.RGBA{0, 226, 4, 255},
	color.RGBA{2, 226, 4, 255},
	color.RGBA{4, 226, 4, 255},
	color.RGBA{6, 226, 4, 255},
	color.RGBA{7, 226, 4, 255},
	color.RGBA{9, 226, 4, 255},
	color.RGBA{11, 226, 4, 255},
	color.RGBA{12, 226, 4, 255},
	color.RGBA{14, 226, 4, 255},
	color.RGBA{16, 226, 4, 255},
	color.RGBA{18, 226, 4, 255},
	color.RGBA{19, 226, 4, 255},
	color.RGBA{21, 226, 4, 255},
	color.RGBA{23, 226, 4, 255},
	color.RGBA{25, 226, 4, 255},
	color.RGBA{26, 226, 4, 255},
	color.RGBA{28, 226, 4, 255},
	color.RGBA{30, 226, 4, 255},
	color.RGBA{32, 226, 4, 255},
	color.RGBA{33, 226, 4, 255},
	color.RGBA{35, 226, 4, 255},
	color.RGBA{37, 226, 4, 255},
	color.RGBA{39, 226, 4, 255},
	color.RGBA{40, 226, 4, 255},
	color.RGBA{42, 226, 4, 255},
	color.RGBA{44, 226, 4, 255},
	color.RGBA{45, 226, 4, 255},
	color.RGBA{47, 226, 4, 255},
	color.RGBA{49, 226, 4, 255},
	color.RGBA{51, 226, 4, 255},
	color.RGBA{52, 226, 4, 255},
	color.RGBA{54, 226, 4, 255},
	color.RGBA{56, 226, 4, 255},
	color.RGBA{58, 226, 4, 255},
	color.RGBA{59, 226, 4, 255},
	color.RGBA{61, 226, 4, 255},
	color.RGBA{63, 226, 4, 255},
	color.RGBA{65, 226, 4, 255},
	color.RGBA{66, 226, 4, 255},
	color.RGBA{68, 226, 4, 255},
	color.RGBA{70, 226, 4, 255},
	color.RGBA{71, 226, 4, 255},
	color.RGBA{73, 226, 4, 255},
	color.RGBA{75, 226, 4, 255},
	color.RGBA{77, 226, 4, 255},
	color.RGBA{78, 226, 4, 255},
	color.RGBA{80, 226, 4, 255},
	color.RGBA{82, 226, 4, 255},
	color.RGBA{84, 226, 4, 255},
	color.RGBA{85, 226, 4, 255},
	color.RGBA{87, 226, 4, 255},
	color.RGBA{89, 226, 4, 255},
	color.RGBA{91, 226, 4, 255},
	color.RGBA{92, 226, 4, 255},
	color.RGBA{94, 226, 4, 255},
	color.RGBA{96, 226, 4, 255},
	color.RGBA{97, 226, 4, 255},
	color.RGBA{99, 226, 4, 255},
	color.RGBA{101, 226, 4, 255},
	color.RGBA{103, 226, 4, 255},
	color.RGBA{104, 226, 4, 255},
	color.RGBA{106, 226, 4, 255},
	color.RGBA{108, 226, 4, 255},
	color.RGBA{110, 226, 4, 255},
	color.RGBA{111, 226, 4, 255},
	color.RGBA{113, 226, 4, 255},
	color.RGBA{115, 226, 4, 255},
	color.RGBA{117, 226, 4, 255},
	color.RGBA{118, 226, 4, 255},
	color.RGBA{120, 226, 4, 255},
	color.RGBA{122, 226, 4, 255},
	color.RGBA{123, 226, 4, 255},
	color.RGBA{125, 226, 4, 255},
	color.RGBA{127, 226, 4, 255},
	color.RGBA{128, 226, 4, 255},
	color.RGBA{129, 226, 4, 255},
	color.RGBA{130, 226, 4, 255},
	color.RGBA{131, 226, 4, 255},
	color.RGBA{132, 226, 4, 255},
	color.RGBA{133, 226, 4, 255},
	color.RGBA{134, 226, 4, 255},
	color.RGBA{134, 226, 4, 255},
	color.RGBA{135, 226, 4, 255},
	color.RGBA{136, 226, 4, 255},
	color.RGBA{137, 226, 4, 255},
	color.RGBA{138, 226, 4, 255},
	color.RGBA{139, 226, 4, 255},
	color.RGBA{140, 226, 4, 255},
	color.RGBA{141, 226, 4, 255},
	color.RGBA{142, 226, 4, 255},
	color.RGBA{143, 226, 4, 255},
	color.RGBA{144, 226, 4, 255},
	color.RGBA{145, 226, 4, 255},
	color.RGBA{146, 226, 4, 255},
	color.RGBA{147, 226, 4, 255},
	color.RGBA{148, 226, 4, 255},
	color.RGBA{148, 226, 4, 255},
	color.RGBA{149, 226, 4, 255},
	color.RGBA{150, 226, 4, 255},
	color.RGBA{151, 226, 4, 255},
	color.RGBA{152, 226, 4, 255},
	color.RGBA{153, 226, 4, 255},
	color.RGBA{154, 226, 4, 255},
	color.RGBA{155, 226, 4, 255},
	color.RGBA{156, 226, 4, 255},
	color.RGBA{157, 226, 4, 255},
	color.RGBA{158, 226, 4, 255},
	color.RGBA{159, 226, 4, 255},
	color.RGBA{160, 226, 4, 255},
	color.RGBA{161, 226, 4, 255},
	color.RGBA{162, 226, 4, 255},
	color.RGBA{162, 226, 4, 255},
	color.RGBA{163, 226, 4, 255},
	color.RGBA{164, 226, 4, 255},
	color.RGBA{165, 226, 4, 255},
	color.RGBA{166, 226, 4, 255},
	color.RGBA{167, 226, 4, 255},
	color.RGBA{168, 226, 4, 255},
	color.RGBA{169, 226, 4, 255},
	color.RGBA{170, 226, 4, 255},
	color.RGBA{171, 226, 4, 255},
	color.RGBA{172, 226, 4, 255},
	color.RGBA{173, 226, 4, 255},
	color.RGBA{174, 226, 4, 255},
	color.RGBA{175, 226, 4, 255},
	color.RGBA{176, 226, 4, 255},
	color.RGBA{176, 226, 4, 255},
	color.RGBA{177, 226, 4, 255},
	color.RGBA{178, 226, 4, 255},
	color.RGBA{179, 226, 4, 255},
	color.RGBA{180, 226, 4, 255},
	color.RGBA{181, 226, 4, 255},
	color.RGBA{182, 226, 4, 255},
	color.RGBA{183, 226, 4, 255},
	color.RGBA{184, 226, 4, 255},
	color.RGBA{185, 226, 4, 255},
	color.RGBA{186, 226, 4, 255},
	color.RGBA{187, 226, 4, 255},
	color.RGBA{188, 226, 4, 255},
	color.RGBA{189, 226, 4, 255},
	color.RGBA{190, 226, 4, 255},
	color.RGBA{190, 226, 4, 255},
	color.RGBA{191, 226, 4, 255},
	color.RGBA{192, 226, 4, 255},
	color.RGBA{193, 226, 4, 255},
	color.RGBA{194, 226, 4, 255},
	color.RGBA{195, 226, 4, 255},
	color.RGBA{196, 226, 4, 255},
	color.RGBA{197, 226, 4, 255},
	color.RGBA{198, 226, 4, 255},
	color.RGBA{199, 226, 4, 255},
	color.RGBA{200, 226, 4, 255},
	color.RGBA{201, 226, 4, 255},
	color.RGBA{202, 226, 4, 255},
	color.RGBA{203, 226, 4, 255},
	color.RGBA{204, 226, 4, 255},
	color.RGBA{204, 226, 4, 255},
	color.RGBA{205, 226, 4, 255},
	color.RGBA{206, 226, 4, 255},
	color.RGBA{207, 226, 4, 255},
	color.RGBA{208, 226, 4, 255},
	color.RGBA{209, 226, 4, 255},
	color.RGBA{210, 226, 4, 255},
	color.RGBA{211, 226, 4, 255},
	color.RGBA{212, 226, 4, 255},
	color.RGBA{213, 226, 4, 255},
	color.RGBA{214, 226, 4, 255},
	color.RGBA{215, 226, 4, 255},
	color.RGBA{216, 226, 4, 255},
	color.RGBA{217, 226, 4, 255},
	color.RGBA{218, 226, 4, 255},
	color.RGBA{218, 226, 4, 255},
	color.RGBA{219, 226, 4, 255},
	color.RGBA{220, 226, 4, 255},
	color.RGBA{221, 226, 4, 255},
	color.RGBA{222, 226, 4, 255},
	color.RGBA{223, 226, 4, 255},
	color.RGBA{224, 226, 4, 255},
	color.RGBA{225, 226, 4, 255},
	color.RGBA{226, 226, 4, 255},
	color.RGBA{227, 226, 4, 255},
	color.RGBA{228, 226, 4, 255},
	color.RGBA{229, 226, 4, 255},
	color.RGBA{230, 226, 4, 255},
	color.RGBA{231, 226, 4, 255},
	color.RGBA{232, 226, 4, 255},
	color.RGBA{233, 226, 4, 255},
	color.RGBA{233, 226, 4, 255},
	color.RGBA{234, 226, 4, 255},
	color.RGBA{235, 226, 4, 255},
	color.RGBA{236, 226, 4, 255},
	color.RGBA{237, 226, 4, 255},
	color.RGBA{238, 226, 4, 255},
	color.RGBA{239, 226, 4, 255},
	color.RGBA{240, 226, 4, 255},
	color.RGBA{241, 226, 4, 255},
	color.RGBA{242, 226, 4, 255},
	color.RGBA{243, 226, 4, 255},
	color.RGBA{244, 226, 4, 255},
	color.RGBA{245, 226, 4, 255},
	color.RGBA{246, 226, 4, 255},
	color.RGBA{247, 226, 4, 255},
	color.RGBA{247, 226, 4, 255},
	color.RGBA{248, 226, 4, 255},
	color.RGBA{249, 226, 4, 255},
	color.RGBA{250, 226, 4, 255},
	color.RGBA{251, 226, 4, 255},
	color.RGBA{252, 226, 4, 255},
	color.RGBA{253, 226, 4, 255},
	color.RGBA{254, 226, 4, 255},
	color.RGBA{254, 225, 4, 255},
	color.RGBA{254, 225, 4, 255},
	color.RGBA{254, 224, 4, 255},
	color.RGBA{254, 224, 4, 255},
	color.RGBA{254, 223, 4, 255},
	color.RGBA{254, 223, 4, 255},
	color.RGBA{254, 222, 4, 255},
	color.RGBA{254, 221, 4, 255},
	color.RGBA{254, 221, 4, 255},
	color.RGBA{254, 220, 4, 255},
	color.RGBA{254, 220, 4, 255},
	color.RGBA{254, 219, 4, 255},
	color.RGBA{254, 219, 4, 255},
	color.RGBA{254, 218, 4, 255},
	color.RGBA{254, 217, 4, 255},
	color.RGBA{254, 217, 4, 255},
	color.RGBA{254, 216, 4, 255},
	color.RGBA{254, 216, 4, 255},
	color.RGBA{254, 215, 4, 255},
	color.RGBA{254, 215, 4, 255},
	color.RGBA{254, 214, 4, 255},
	color.RGBA{254, 213, 4, 255},
	color.RGBA{254, 213, 4, 255},
	color.RGBA{254, 212, 4, 255},
	color.RGBA{254, 212, 4, 255},
	color.RGBA{254, 211, 4, 255},
	color.RGBA{254, 211, 4, 255},
	color.RGBA{254, 210, 4, 255},
	color.RGBA{254, 210, 4, 255},
	color.RGBA{254, 209, 4, 255},
	color.RGBA{254, 208, 4, 255},
	color.RGBA{254, 208, 4, 255},
	color.RGBA{254, 207, 4, 255},
	color.RGBA{254, 207, 4, 255},
	color.RGBA{254, 206, 4, 255},
	color.RGBA{254, 206, 4, 255},
	color.RGBA{254, 205, 4, 255},
	color.RGBA{254, 204, 4, 255},
	color.RGBA{254, 204, 4, 255},
	color.RGBA{254, 203, 4, 255},
	color.RGBA{254, 203, 4, 255},
	color.RGBA{254, 202, 4, 255},
	color.RGBA{254, 202, 4, 255},
	color.RGBA{254, 201, 4, 255},
	color.RGBA{254, 200, 4, 255},
	color.RGBA{254, 200, 4, 255},
	color.RGBA{254, 199, 4, 255},
	color.RGBA{254, 199, 4, 255},
	color.RGBA{254, 198, 4, 255},
	color.RGBA{254, 198, 4, 255},
	color.RGBA{254, 197, 4, 255},
	color.RGBA{254, 196, 4, 255},
	color.RGBA{254, 196, 4, 255},
	color.RGBA{254, 195, 4, 255},
	color.RGBA{254, 195, 4, 255},
	color.RGBA{254, 194, 4, 255},
	color.RGBA{254, 194, 4, 255},
	color.RGBA{254, 193, 4, 255},
	color.RGBA{254, 192, 4, 255},
	color.RGBA{254, 192, 4, 255},
	color.RGBA{254, 191, 4, 255},
	color.RGBA{254, 191, 4, 255},
	color.RGBA{254, 190, 4, 255},
	color.RGBA{254, 190, 4, 255},
	color.RGBA{254, 189, 4, 255},
	color.RGBA{254, 188, 4, 255},
	color.RGBA{254, 188, 4, 255},
	color.RGBA{254, 187, 4, 255},
	color.RGBA{254, 187, 4, 255},
	color.RGBA{254, 186, 4, 255},
	color.RGBA{254, 186, 4, 255},
	color.RGBA{254, 185, 4, 255},
	color.RGBA{254, 185, 4, 255},
	color.RGBA{254, 184, 4, 255},
	color.RGBA{254, 183, 4, 255},
	color.RGBA{254, 183, 4, 255},
	color.RGBA{254, 182, 4, 255},
	color.RGBA{254, 182, 4, 255},
	color.RGBA{254, 181, 4, 255},
	color.RGBA{254, 181, 4, 255},
	color.RGBA{254, 180, 4, 255},
	color.RGBA{254, 179, 4, 255},
	color.RGBA{254, 179, 4, 255},
	color.RGBA{254, 178, 4, 255},
	color.RGBA{254, 178, 4, 255},
	color.RGBA{254, 177, 4, 255},
	color.RGBA{254, 177, 4, 255},
	color.RGBA{254, 176, 4, 255},
	color.RGBA{254, 175, 4, 255},
	color.RGBA{254, 175, 4, 255},
	color.RGBA{254, 174, 4, 255},
	color.RGBA{254, 174, 4, 255},
	color.RGBA{254, 173, 4, 255},
	color.RGBA{254, 173, 4, 255},
	color.RGBA{254, 172, 4, 255},
	color.RGBA{254, 171, 4, 255},
	color.RGBA{254, 171, 4, 255},
	color.RGBA{254, 170, 4, 255},
	color.RGBA{254, 170, 4, 255},
	color.RGBA{254, 169, 4, 255},
	color.RGBA{254, 169, 4, 255},
	color.RGBA{254, 168, 4, 255},
	color.RGBA{254, 167, 4, 255},
	color.RGBA{254, 167, 4, 255},
	color.RGBA{254, 166, 4, 255},
	color.RGBA{254, 166, 4, 255},
	color.RGBA{254, 165, 4, 255},
	color.RGBA{254, 165, 4, 255},
	color.RGBA{254, 164, 4, 255},
	color.RGBA{254, 163, 4, 255},
	color.RGBA{254, 163, 4, 255},
	color.RGBA{254, 162, 4, 255},
	color.RGBA{254, 162, 4, 255},
	color.RGBA{254, 161, 4, 255},
	color.RGBA{254, 161, 4, 255},
	color.RGBA{254, 160, 4, 255},
	color.RGBA{254, 159, 4, 255},
	color.RGBA{254, 159, 4, 255},
	color.RGBA{254, 158, 4, 255},
	color.RGBA{254, 158, 4, 255},
	color.RGBA{254, 157, 4, 255},
	color.RGBA{254, 157, 4, 255},
	color.RGBA{254, 156, 4, 255},
	color.RGBA{254, 156, 4, 255},
	color.RGBA{254, 155, 4, 255},
	color.RGBA{254, 154, 4, 255},
	color.RGBA{254, 154, 4, 255},
	color.RGBA{254, 153, 4, 255},
	color.RGBA{254, 153, 4, 255},
	color.RGBA{254, 152, 4, 255},
	color.RGBA{254, 152, 4, 255},
	color.RGBA{254, 151, 4, 255},
	color.RGBA{254, 150, 4, 255},
	color.RGBA{254, 150, 4, 255},
	color.RGBA{254, 149, 4, 255},
	color.RGBA{254, 149, 4, 255},
	color.RGBA{254, 148, 4, 255},
	color.RGBA{254, 148, 4, 255},
	color.RGBA{254, 147, 4, 255},
	color.RGBA{254, 146, 4, 255},
	color.RGBA{254, 146, 4, 255},
	color.RGBA{254, 145, 4, 255},
	color.RGBA{254, 143, 4, 255},
	color.RGBA{254, 142, 4, 255},
	color.RGBA{254, 141, 4, 255},
	color.RGBA{254, 140, 4, 255},
	color.RGBA{254, 138, 4, 255},
	color.RGBA{254, 137, 4, 255},
	color.RGBA{254, 136, 4, 255},
	color.RGBA{254, 134, 4, 255},
	color.RGBA{254, 133, 4, 255},
	color.RGBA{254, 132, 4, 255},
	color.RGBA{254, 131, 4, 255},
	color.RGBA{254, 129, 4, 255},
	color.RGBA{254, 128, 4, 255},
	color.RGBA{254, 127, 4, 255},
	color.RGBA{254, 126, 4, 255},
	color.RGBA{254, 124, 4, 255},
	color.RGBA{254, 123, 4, 255},
	color.RGBA{254, 122, 4, 255},
	color.RGBA{254, 120, 4, 255},
	color.RGBA{254, 119, 4, 255},
	color.RGBA{254, 118, 4, 255},
	color.RGBA{254, 117, 4, 255},
	color.RGBA{254, 115, 4, 255},
	color.RGBA{254, 114, 4, 255},
	color.RGBA{254, 113, 4, 255},
	color.RGBA{254, 111, 4, 255},
	color.RGBA{254, 110, 4, 255},
	color.RGBA{254, 109, 4, 255},
	color.RGBA{254, 108, 4, 255},
	color.RGBA{254, 106, 4, 255},
	color.RGBA{254, 105, 4, 255},
	color.RGBA{254, 104, 4, 255},
	color.RGBA{254, 103, 4, 255},
	color.RGBA{254, 101, 4, 255},
	color.RGBA{254, 100, 4, 255},
	color.RGBA{254, 99, 4, 255},
	color.RGBA{254, 97, 4, 255},
	color.RGBA{254, 96, 4, 255},
	color.RGBA{254, 95, 4, 255},
	color.RGBA{254, 94, 4, 255},
	color.RGBA{254, 92, 4, 255},
	color.RGBA{254, 91, 4, 255},
	color.RGBA{254, 90, 4, 255},
	color.RGBA{254, 88, 4, 255},
	color.RGBA{254, 87, 4, 255},
	color.RGBA{254, 86, 4, 255},
	color.RGBA{254, 85, 4, 255},
	color.RGBA{254, 83, 4, 255},
	color.RGBA{254, 82, 4, 255},
	color.RGBA{254, 81, 4, 255},
	color.RGBA{254, 79, 4, 255},
	color.RGBA{254, 78, 4, 255},
	color.RGBA{254, 77, 4, 255},
	color.RGBA{254, 76, 4, 255},
	color.RGBA{254, 74, 4, 255},
	color.RGBA{254, 73, 4, 255},
	color.RGBA{254, 72, 4, 255},
	color.RGBA{254, 71, 4, 255},
	color.RGBA{254, 69, 4, 255},
	color.RGBA{254, 68, 4, 255},
	color.RGBA{254, 67, 4, 255},
	color.RGBA{254, 65, 4, 255},
}
