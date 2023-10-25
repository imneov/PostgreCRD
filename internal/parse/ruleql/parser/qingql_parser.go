// Generated from QingQL.g4 by ANTLR 4.7.

package parser // QingQL

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 58, 299,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5,
	2, 68, 10, 2, 3, 2, 5, 2, 71, 10, 2, 3, 3, 3, 3, 3, 3, 7, 3, 76, 10, 3,
	12, 3, 14, 3, 79, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 86, 10, 4,
	3, 5, 5, 5, 89, 10, 5, 3, 5, 3, 5, 3, 5, 7, 5, 94, 10, 5, 12, 5, 14, 5,
	97, 11, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	7, 9, 109, 10, 9, 12, 9, 14, 9, 112, 11, 9, 3, 10, 3, 10, 3, 10, 7, 10,
	117, 10, 10, 12, 10, 14, 10, 120, 11, 10, 3, 11, 5, 11, 123, 10, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 7, 13, 134,
	10, 13, 12, 13, 14, 13, 137, 11, 13, 3, 14, 3, 14, 5, 14, 141, 10, 14,
	3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 5, 18, 154, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5,
	23, 194, 10, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 7, 23, 205, 10, 23, 12, 23, 14, 23, 208, 11, 23, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 217, 10, 24, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 7, 25, 230, 10,
	25, 12, 25, 14, 25, 233, 11, 25, 3, 25, 3, 25, 5, 25, 237, 10, 25, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 7, 26, 244, 10, 26, 12, 26, 14, 26, 247, 11,
	26, 5, 26, 249, 10, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28,
	6, 28, 258, 10, 28, 13, 28, 14, 28, 259, 3, 28, 6, 28, 263, 10, 28, 13,
	28, 14, 28, 264, 3, 28, 3, 28, 5, 28, 269, 10, 28, 3, 29, 3, 29, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 5, 30, 286, 10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 5, 31, 297, 10, 31, 3, 31, 2, 3, 44, 32, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 58, 60, 2, 8, 6, 2, 4, 4, 45, 45, 50, 51, 54,
	54, 3, 2, 37, 41, 3, 2, 42, 44, 3, 2, 45, 46, 4, 2, 17, 17, 19, 23, 4,
	2, 50, 50, 55, 55, 2, 310, 2, 62, 3, 2, 2, 2, 4, 72, 3, 2, 2, 2, 6, 85,
	3, 2, 2, 2, 8, 88, 3, 2, 2, 2, 10, 98, 3, 2, 2, 2, 12, 100, 3, 2, 2, 2,
	14, 103, 3, 2, 2, 2, 16, 105, 3, 2, 2, 2, 18, 113, 3, 2, 2, 2, 20, 122,
	3, 2, 2, 2, 22, 126, 3, 2, 2, 2, 24, 130, 3, 2, 2, 2, 26, 140, 3, 2, 2,
	2, 28, 142, 3, 2, 2, 2, 30, 144, 3, 2, 2, 2, 32, 146, 3, 2, 2, 2, 34, 153,
	3, 2, 2, 2, 36, 155, 3, 2, 2, 2, 38, 162, 3, 2, 2, 2, 40, 171, 3, 2, 2,
	2, 42, 178, 3, 2, 2, 2, 44, 193, 3, 2, 2, 2, 46, 216, 3, 2, 2, 2, 48, 218,
	3, 2, 2, 2, 50, 238, 3, 2, 2, 2, 52, 252, 3, 2, 2, 2, 54, 268, 3, 2, 2,
	2, 56, 270, 3, 2, 2, 2, 58, 285, 3, 2, 2, 2, 60, 296, 3, 2, 2, 2, 62, 63,
	7, 27, 2, 2, 63, 64, 5, 4, 3, 2, 64, 65, 7, 18, 2, 2, 65, 67, 5, 8, 5,
	2, 66, 68, 5, 12, 7, 2, 67, 66, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 70,
	3, 2, 2, 2, 69, 71, 5, 22, 12, 2, 70, 69, 3, 2, 2, 2, 70, 71, 3, 2, 2,
	2, 71, 3, 3, 2, 2, 2, 72, 77, 5, 6, 4, 2, 73, 74, 7, 3, 2, 2, 74, 76, 5,
	6, 4, 2, 75, 73, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77,
	78, 3, 2, 2, 2, 78, 5, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 81, 5, 44, 23,
	2, 81, 82, 7, 12, 2, 2, 82, 83, 5, 54, 28, 2, 83, 86, 3, 2, 2, 2, 84, 86,
	5, 54, 28, 2, 85, 80, 3, 2, 2, 2, 85, 84, 3, 2, 2, 2, 86, 7, 3, 2, 2, 2,
	87, 89, 7, 43, 2, 2, 88, 87, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 90, 3,
	2, 2, 2, 90, 95, 5, 10, 6, 2, 91, 92, 7, 43, 2, 2, 92, 94, 5, 10, 6, 2,
	93, 91, 3, 2, 2, 2, 94, 97, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 96, 3,
	2, 2, 2, 96, 9, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 98, 99, 9, 2, 2, 2, 99,
	11, 3, 2, 2, 2, 100, 101, 7, 29, 2, 2, 101, 102, 5, 14, 8, 2, 102, 13,
	3, 2, 2, 2, 103, 104, 5, 16, 9, 2, 104, 15, 3, 2, 2, 2, 105, 110, 5, 18,
	10, 2, 106, 107, 7, 13, 2, 2, 107, 109, 5, 18, 10, 2, 108, 106, 3, 2, 2,
	2, 109, 112, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111,
	17, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 113, 118, 5, 20, 11, 2, 114, 115,
	7, 26, 2, 2, 115, 117, 5, 20, 11, 2, 116, 114, 3, 2, 2, 2, 117, 120, 3,
	2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 19, 3, 2, 2,
	2, 120, 118, 3, 2, 2, 2, 121, 123, 7, 24, 2, 2, 122, 121, 3, 2, 2, 2, 122,
	123, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 125, 5, 44, 23, 2, 125, 21,
	3, 2, 2, 2, 126, 127, 7, 31, 2, 2, 127, 128, 7, 32, 2, 2, 128, 129, 5,
	24, 13, 2, 129, 23, 3, 2, 2, 2, 130, 135, 5, 26, 14, 2, 131, 132, 7, 3,
	2, 2, 132, 134, 5, 26, 14, 2, 133, 131, 3, 2, 2, 2, 134, 137, 3, 2, 2,
	2, 135, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 25, 3, 2, 2, 2, 137,
	135, 3, 2, 2, 2, 138, 141, 5, 34, 18, 2, 139, 141, 5, 54, 28, 2, 140, 138,
	3, 2, 2, 2, 140, 139, 3, 2, 2, 2, 141, 27, 3, 2, 2, 2, 142, 143, 9, 3,
	2, 2, 143, 29, 3, 2, 2, 2, 144, 145, 7, 51, 2, 2, 145, 31, 3, 2, 2, 2,
	146, 147, 7, 51, 2, 2, 147, 33, 3, 2, 2, 2, 148, 154, 5, 36, 19, 2, 149,
	154, 5, 38, 20, 2, 150, 154, 5, 40, 21, 2, 151, 154, 5, 42, 22, 2, 152,
	154, 5, 38, 20, 2, 153, 148, 3, 2, 2, 2, 153, 149, 3, 2, 2, 2, 153, 150,
	3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 153, 152, 3, 2, 2, 2, 154, 35, 3, 2,
	2, 2, 155, 156, 7, 33, 2, 2, 156, 157, 7, 5, 2, 2, 157, 158, 5, 28, 15,
	2, 158, 159, 7, 3, 2, 2, 159, 160, 5, 32, 17, 2, 160, 161, 7, 6, 2, 2,
	161, 37, 3, 2, 2, 2, 162, 163, 7, 34, 2, 2, 163, 164, 7, 5, 2, 2, 164,
	165, 5, 28, 15, 2, 165, 166, 7, 3, 2, 2, 166, 167, 5, 32, 17, 2, 167, 168,
	7, 3, 2, 2, 168, 169, 5, 30, 16, 2, 169, 170, 7, 6, 2, 2, 170, 39, 3, 2,
	2, 2, 171, 172, 7, 35, 2, 2, 172, 173, 7, 5, 2, 2, 173, 174, 5, 28, 15,
	2, 174, 175, 7, 3, 2, 2, 175, 176, 5, 32, 17, 2, 176, 177, 7, 6, 2, 2,
	177, 41, 3, 2, 2, 2, 178, 179, 7, 36, 2, 2, 179, 180, 7, 5, 2, 2, 180,
	181, 5, 28, 15, 2, 181, 182, 7, 3, 2, 2, 182, 183, 5, 32, 17, 2, 183, 184,
	7, 6, 2, 2, 184, 43, 3, 2, 2, 2, 185, 186, 8, 23, 1, 2, 186, 194, 5, 46,
	24, 2, 187, 188, 7, 5, 2, 2, 188, 189, 5, 44, 23, 2, 189, 190, 7, 6, 2,
	2, 190, 194, 3, 2, 2, 2, 191, 194, 5, 50, 26, 2, 192, 194, 5, 48, 25, 2,
	193, 185, 3, 2, 2, 2, 193, 187, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 193,
	192, 3, 2, 2, 2, 194, 206, 3, 2, 2, 2, 195, 196, 12, 7, 2, 2, 196, 197,
	9, 4, 2, 2, 197, 205, 5, 44, 23, 8, 198, 199, 12, 6, 2, 2, 199, 200, 9,
	5, 2, 2, 200, 205, 5, 44, 23, 7, 201, 202, 12, 5, 2, 2, 202, 203, 9, 6,
	2, 2, 203, 205, 5, 44, 23, 6, 204, 195, 3, 2, 2, 2, 204, 198, 3, 2, 2,
	2, 204, 201, 3, 2, 2, 2, 205, 208, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 206,
	207, 3, 2, 2, 2, 207, 45, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 209, 217, 7,
	48, 2, 2, 210, 217, 7, 49, 2, 2, 211, 217, 7, 51, 2, 2, 212, 217, 7, 52,
	2, 2, 213, 217, 7, 53, 2, 2, 214, 217, 7, 57, 2, 2, 215, 217, 5, 54, 28,
	2, 216, 209, 3, 2, 2, 2, 216, 210, 3, 2, 2, 2, 216, 211, 3, 2, 2, 2, 216,
	212, 3, 2, 2, 2, 216, 213, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 216, 215,
	3, 2, 2, 2, 217, 47, 3, 2, 2, 2, 218, 219, 7, 14, 2, 2, 219, 220, 5, 44,
	23, 2, 220, 221, 7, 30, 2, 2, 221, 222, 5, 44, 23, 2, 222, 223, 7, 28,
	2, 2, 223, 231, 5, 44, 23, 2, 224, 225, 7, 30, 2, 2, 225, 226, 5, 44, 23,
	2, 226, 227, 7, 28, 2, 2, 227, 228, 5, 44, 23, 2, 228, 230, 3, 2, 2, 2,
	229, 224, 3, 2, 2, 2, 230, 233, 3, 2, 2, 2, 231, 229, 3, 2, 2, 2, 231,
	232, 3, 2, 2, 2, 232, 236, 3, 2, 2, 2, 233, 231, 3, 2, 2, 2, 234, 235,
	7, 15, 2, 2, 235, 237, 5, 44, 23, 2, 236, 234, 3, 2, 2, 2, 236, 237, 3,
	2, 2, 2, 237, 49, 3, 2, 2, 2, 238, 239, 7, 50, 2, 2, 239, 248, 7, 5, 2,
	2, 240, 245, 5, 44, 23, 2, 241, 242, 7, 3, 2, 2, 242, 244, 5, 44, 23, 2,
	243, 241, 3, 2, 2, 2, 244, 247, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 245,
	246, 3, 2, 2, 2, 246, 249, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 248, 240,
	3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250, 251, 7, 6,
	2, 2, 251, 51, 3, 2, 2, 2, 252, 253, 7, 42, 2, 2, 253, 53, 3, 2, 2, 2,
	254, 269, 5, 52, 27, 2, 255, 269, 5, 56, 29, 2, 256, 258, 7, 7, 2, 2, 257,
	256, 3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 259, 260,
	3, 2, 2, 2, 260, 262, 3, 2, 2, 2, 261, 263, 5, 56, 29, 2, 262, 261, 3,
	2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 264, 265, 3, 2, 2,
	2, 265, 266, 3, 2, 2, 2, 266, 267, 7, 7, 2, 2, 267, 269, 3, 2, 2, 2, 268,
	254, 3, 2, 2, 2, 268, 255, 3, 2, 2, 2, 268, 257, 3, 2, 2, 2, 269, 55, 3,
	2, 2, 2, 270, 271, 9, 7, 2, 2, 271, 57, 3, 2, 2, 2, 272, 273, 7, 55, 2,
	2, 273, 274, 7, 8, 2, 2, 274, 286, 7, 9, 2, 2, 275, 276, 7, 55, 2, 2, 276,
	277, 7, 8, 2, 2, 277, 278, 7, 51, 2, 2, 278, 286, 7, 9, 2, 2, 279, 280,
	7, 55, 2, 2, 280, 281, 7, 8, 2, 2, 281, 282, 7, 4, 2, 2, 282, 286, 7, 9,
	2, 2, 283, 286, 7, 55, 2, 2, 284, 286, 7, 53, 2, 2, 285, 272, 3, 2, 2,
	2, 285, 275, 3, 2, 2, 2, 285, 279, 3, 2, 2, 2, 285, 283, 3, 2, 2, 2, 285,
	284, 3, 2, 2, 2, 286, 59, 3, 2, 2, 2, 287, 288, 7, 50, 2, 2, 288, 297,
	7, 10, 2, 2, 289, 290, 7, 50, 2, 2, 290, 291, 7, 8, 2, 2, 291, 292, 7,
	51, 2, 2, 292, 297, 7, 9, 2, 2, 293, 294, 7, 50, 2, 2, 294, 297, 7, 11,
	2, 2, 295, 297, 7, 50, 2, 2, 296, 287, 3, 2, 2, 2, 296, 289, 3, 2, 2, 2,
	296, 293, 3, 2, 2, 2, 296, 295, 3, 2, 2, 2, 297, 61, 3, 2, 2, 2, 27, 67,
	70, 77, 85, 88, 95, 110, 118, 122, 135, 140, 153, 193, 204, 206, 216, 231,
	236, 245, 248, 259, 264, 268, 285, 296,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'#'", "'('", "')'", "'\"'", "'['", "']'", "'[]'", "'[#]'",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "'*'", "'/'", "'%'", "'+'",
	"'-'", "'.'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "AS", "AND", "CASE", "ELSE", "END",
	"EQ", "FROM", "GT", "GTE", "LT", "LTE", "NE", "NOT", "NULL", "OR", "SELECT",
	"THEN", "WHERE", "WHEN", "GROUP", "BY", "TUMBLINGWINDOW", "HOPPINGWINDOW",
	"SLIDINGWINDOW", "SESSIONWINDOW", "DD", "HH", "MI", "SS", "MS", "MUL",
	"DIV", "MOD", "ADD", "SUB", "DOT", "TRUE", "FALSE", "INDENTIFIER", "NUMBER",
	"INTEGER", "FLOAT", "TOPICITEM", "PATHITEM", "ARRAYITEM", "STRING", "WHITESPACE",
}

var ruleNames = []string{
	"root", "fields", "field_elem", "topic", "topic_item", "where", "filter",
	"filter_condition", "filter_condition_or", "filter_condition_not", "group",
	"group_by_elements", "group_by_element", "group_window_unit", "group_window_length",
	"group_window_interval", "group_windows", "group_tumbling_window", "group_hopping_window",
	"group_sliding_window", "group_session_window", "expr", "constant", "switch_stmt",
	"call_expr", "asterisk", "xpath_name", "dotnotation", "identifierWithTOPICITEM",
	"identifierWithQualifier",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type QingQLParser struct {
	*antlr.BaseParser
}

func NewQingQLParser(input antlr.TokenStream) *QingQLParser {
	this := new(QingQLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "QingQL.g4"

	return this
}

// QingQLParser tokens.
const (
	QingQLParserEOF            = antlr.TokenEOF
	QingQLParserT__0           = 1
	QingQLParserT__1           = 2
	QingQLParserT__2           = 3
	QingQLParserT__3           = 4
	QingQLParserT__4           = 5
	QingQLParserT__5           = 6
	QingQLParserT__6           = 7
	QingQLParserT__7           = 8
	QingQLParserT__8           = 9
	QingQLParserAS             = 10
	QingQLParserAND            = 11
	QingQLParserCASE           = 12
	QingQLParserELSE           = 13
	QingQLParserEND            = 14
	QingQLParserEQ             = 15
	QingQLParserFROM           = 16
	QingQLParserGT             = 17
	QingQLParserGTE            = 18
	QingQLParserLT             = 19
	QingQLParserLTE            = 20
	QingQLParserNE             = 21
	QingQLParserNOT            = 22
	QingQLParserNULL           = 23
	QingQLParserOR             = 24
	QingQLParserSELECT         = 25
	QingQLParserTHEN           = 26
	QingQLParserWHERE          = 27
	QingQLParserWHEN           = 28
	QingQLParserGROUP          = 29
	QingQLParserBY             = 30
	QingQLParserTUMBLINGWINDOW = 31
	QingQLParserHOPPINGWINDOW  = 32
	QingQLParserSLIDINGWINDOW  = 33
	QingQLParserSESSIONWINDOW  = 34
	QingQLParserDD             = 35
	QingQLParserHH             = 36
	QingQLParserMI             = 37
	QingQLParserSS             = 38
	QingQLParserMS             = 39
	QingQLParserMUL            = 40
	QingQLParserDIV            = 41
	QingQLParserMOD            = 42
	QingQLParserADD            = 43
	QingQLParserSUB            = 44
	QingQLParserDOT            = 45
	QingQLParserTRUE           = 46
	QingQLParserFALSE          = 47
	QingQLParserINDENTIFIER    = 48
	QingQLParserNUMBER         = 49
	QingQLParserINTEGER        = 50
	QingQLParserFLOAT          = 51
	QingQLParserTOPICITEM      = 52
	QingQLParserPATHITEM       = 53
	QingQLParserARRAYITEM      = 54
	QingQLParserSTRING         = 55
	QingQLParserWHITESPACE     = 56
)

// QingQLParser rules.
const (
	QingQLParserRULE_root                    = 0
	QingQLParserRULE_fields                  = 1
	QingQLParserRULE_field_elem              = 2
	QingQLParserRULE_topic                   = 3
	QingQLParserRULE_topic_item              = 4
	QingQLParserRULE_where                   = 5
	QingQLParserRULE_filter                  = 6
	QingQLParserRULE_filter_condition        = 7
	QingQLParserRULE_filter_condition_or     = 8
	QingQLParserRULE_filter_condition_not    = 9
	QingQLParserRULE_group                   = 10
	QingQLParserRULE_group_by_elements       = 11
	QingQLParserRULE_group_by_element        = 12
	QingQLParserRULE_group_window_unit       = 13
	QingQLParserRULE_group_window_length     = 14
	QingQLParserRULE_group_window_interval   = 15
	QingQLParserRULE_group_windows           = 16
	QingQLParserRULE_group_tumbling_window   = 17
	QingQLParserRULE_group_hopping_window    = 18
	QingQLParserRULE_group_sliding_window    = 19
	QingQLParserRULE_group_session_window    = 20
	QingQLParserRULE_expr                    = 21
	QingQLParserRULE_constant                = 22
	QingQLParserRULE_switch_stmt             = 23
	QingQLParserRULE_call_expr               = 24
	QingQLParserRULE_asterisk                = 25
	QingQLParserRULE_xpath_name              = 26
	QingQLParserRULE_dotnotation             = 27
	QingQLParserRULE_identifierWithTOPICITEM = 28
	QingQLParserRULE_identifierWithQualifier = 29
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) SELECT() antlr.TerminalNode {
	return s.GetToken(QingQLParserSELECT, 0)
}

func (s *RootContext) Fields() IFieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsContext)
}

func (s *RootContext) FROM() antlr.TerminalNode {
	return s.GetToken(QingQLParserFROM, 0)
}

func (s *RootContext) Topic() ITopicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITopicContext)
}

func (s *RootContext) Where() IWhereContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereContext)
}

func (s *RootContext) Group() IGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (p *QingQLParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QingQLParserRULE_root)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(QingQLParserSELECT)
	}
	{
		p.SetState(61)
		p.Fields()
	}
	{
		p.SetState(62)
		p.Match(QingQLParserFROM)
	}
	{
		p.SetState(63)
		p.Topic()
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QingQLParserWHERE {
		{
			p.SetState(64)
			p.Where()
		}

	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QingQLParserGROUP {
		{
			p.SetState(67)
			p.Group()
		}

	}

	return localctx
}

// IFieldsContext is an interface to support dynamic dispatch.
type IFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsContext differentiates from other interfaces.
	IsFieldsContext()
}

type FieldsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsContext() *FieldsContext {
	var p = new(FieldsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_fields
	return p
}

func (*FieldsContext) IsFieldsContext() {}

func NewFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsContext {
	var p = new(FieldsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_fields

	return p
}

func (s *FieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsContext) AllField_elem() []IField_elemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IField_elemContext)(nil)).Elem())
	var tst = make([]IField_elemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IField_elemContext)
		}
	}

	return tst
}

func (s *FieldsContext) Field_elem(i int) IField_elemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_elemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IField_elemContext)
}

func (s *FieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFields(s)
	}
}

func (s *FieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFields(s)
	}
}

func (p *QingQLParser) Fields() (localctx IFieldsContext) {
	localctx = NewFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, QingQLParserRULE_fields)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Field_elem()
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QingQLParserT__0 {
		{
			p.SetState(71)
			p.Match(QingQLParserT__0)
		}
		{
			p.SetState(72)
			p.Field_elem()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IField_elemContext is an interface to support dynamic dispatch.
type IField_elemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_elemContext differentiates from other interfaces.
	IsField_elemContext()
}

type Field_elemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_elemContext() *Field_elemContext {
	var p = new(Field_elemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_field_elem
	return p
}

func (*Field_elemContext) IsField_elemContext() {}

func NewField_elemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_elemContext {
	var p = new(Field_elemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_field_elem

	return p
}

func (s *Field_elemContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_elemContext) CopyFrom(ctx *Field_elemContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Field_elemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_elemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprFieldContext struct {
	*Field_elemContext
}

func NewExprFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprFieldContext {
	var p = new(ExprFieldContext)

	p.Field_elemContext = NewEmptyField_elemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Field_elemContext))

	return p
}

func (s *ExprFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprFieldContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprFieldContext) AS() antlr.TerminalNode {
	return s.GetToken(QingQLParserAS, 0)
}

func (s *ExprFieldContext) Xpath_name() IXpath_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXpath_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IXpath_nameContext)
}

func (s *ExprFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterExprField(s)
	}
}

func (s *ExprFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitExprField(s)
	}
}

type XpathFieldContext struct {
	*Field_elemContext
}

func NewXpathFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *XpathFieldContext {
	var p = new(XpathFieldContext)

	p.Field_elemContext = NewEmptyField_elemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Field_elemContext))

	return p
}

func (s *XpathFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XpathFieldContext) Xpath_name() IXpath_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXpath_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IXpath_nameContext)
}

func (s *XpathFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterXpathField(s)
	}
}

func (s *XpathFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitXpathField(s)
	}
}

func (p *QingQLParser) Field_elem() (localctx IField_elemContext) {
	localctx = NewField_elemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QingQLParserRULE_field_elem)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.expr(0)
		}
		{
			p.SetState(79)
			p.Match(QingQLParserAS)
		}
		{
			p.SetState(80)
			p.Xpath_name()
		}

	case 2:
		localctx = NewXpathFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.Xpath_name()
		}

	}

	return localctx
}

// ITopicContext is an interface to support dynamic dispatch.
type ITopicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopicContext differentiates from other interfaces.
	IsTopicContext()
}

type TopicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopicContext() *TopicContext {
	var p = new(TopicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_topic
	return p
}

func (*TopicContext) IsTopicContext() {}

func NewTopicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopicContext {
	var p = new(TopicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_topic

	return p
}

func (s *TopicContext) GetParser() antlr.Parser { return s.parser }

func (s *TopicContext) AllTopic_item() []ITopic_itemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITopic_itemContext)(nil)).Elem())
	var tst = make([]ITopic_itemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITopic_itemContext)
		}
	}

	return tst
}

func (s *TopicContext) Topic_item(i int) ITopic_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopic_itemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITopic_itemContext)
}

func (s *TopicContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(QingQLParserDIV)
}

func (s *TopicContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(QingQLParserDIV, i)
}

func (s *TopicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterTopic(s)
	}
}

func (s *TopicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitTopic(s)
	}
}

func (p *QingQLParser) Topic() (localctx ITopicContext) {
	localctx = NewTopicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, QingQLParserRULE_topic)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QingQLParserDIV {
		{
			p.SetState(85)
			p.Match(QingQLParserDIV)
		}

	}
	{
		p.SetState(88)
		p.Topic_item()
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QingQLParserDIV {
		{
			p.SetState(89)
			p.Match(QingQLParserDIV)
		}
		{
			p.SetState(90)
			p.Topic_item()
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITopic_itemContext is an interface to support dynamic dispatch.
type ITopic_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopic_itemContext differentiates from other interfaces.
	IsTopic_itemContext()
}

type Topic_itemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopic_itemContext() *Topic_itemContext {
	var p = new(Topic_itemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_topic_item
	return p
}

func (*Topic_itemContext) IsTopic_itemContext() {}

func NewTopic_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Topic_itemContext {
	var p = new(Topic_itemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_topic_item

	return p
}

func (s *Topic_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Topic_itemContext) TOPICITEM() antlr.TerminalNode {
	return s.GetToken(QingQLParserTOPICITEM, 0)
}

func (s *Topic_itemContext) INDENTIFIER() antlr.TerminalNode {
	return s.GetToken(QingQLParserINDENTIFIER, 0)
}

func (s *Topic_itemContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(QingQLParserNUMBER, 0)
}

func (s *Topic_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Topic_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Topic_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterTopic_item(s)
	}
}

func (s *Topic_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitTopic_item(s)
	}
}

func (p *QingQLParser) Topic_item() (localctx ITopic_itemContext) {
	localctx = NewTopic_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, QingQLParserRULE_topic_item)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(96)
	_la = p.GetTokenStream().LA(1)

	if !(_la == QingQLParserT__1 || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(QingQLParserADD-43))|(1<<(QingQLParserINDENTIFIER-43))|(1<<(QingQLParserNUMBER-43))|(1<<(QingQLParserTOPICITEM-43)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IWhereContext is an interface to support dynamic dispatch.
type IWhereContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhereContext differentiates from other interfaces.
	IsWhereContext()
}

type WhereContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereContext() *WhereContext {
	var p = new(WhereContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_where
	return p
}

func (*WhereContext) IsWhereContext() {}

func NewWhereContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereContext {
	var p = new(WhereContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_where

	return p
}

func (s *WhereContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereContext) WHERE() antlr.TerminalNode {
	return s.GetToken(QingQLParserWHERE, 0)
}

func (s *WhereContext) Filter() IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *WhereContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterWhere(s)
	}
}

func (s *WhereContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitWhere(s)
	}
}

func (p *QingQLParser) Where() (localctx IWhereContext) {
	localctx = NewWhereContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, QingQLParserRULE_where)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(QingQLParserWHERE)
	}
	{
		p.SetState(99)
		p.Filter()
	}

	return localctx
}

// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_filter
	return p
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) Filter_condition() IFilter_conditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilter_conditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilter_conditionContext)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFilter(s)
	}
}

func (p *QingQLParser) Filter() (localctx IFilterContext) {
	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, QingQLParserRULE_filter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Filter_condition()
	}

	return localctx
}

// IFilter_conditionContext is an interface to support dynamic dispatch.
type IFilter_conditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilter_conditionContext differentiates from other interfaces.
	IsFilter_conditionContext()
}

type Filter_conditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilter_conditionContext() *Filter_conditionContext {
	var p = new(Filter_conditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_filter_condition
	return p
}

func (*Filter_conditionContext) IsFilter_conditionContext() {}

func NewFilter_conditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Filter_conditionContext {
	var p = new(Filter_conditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_filter_condition

	return p
}

func (s *Filter_conditionContext) GetParser() antlr.Parser { return s.parser }

func (s *Filter_conditionContext) AllFilter_condition_or() []IFilter_condition_orContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilter_condition_orContext)(nil)).Elem())
	var tst = make([]IFilter_condition_orContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilter_condition_orContext)
		}
	}

	return tst
}

func (s *Filter_conditionContext) Filter_condition_or(i int) IFilter_condition_orContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilter_condition_orContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilter_condition_orContext)
}

func (s *Filter_conditionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(QingQLParserAND)
}

func (s *Filter_conditionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(QingQLParserAND, i)
}

func (s *Filter_conditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Filter_conditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Filter_conditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFilter_condition(s)
	}
}

func (s *Filter_conditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFilter_condition(s)
	}
}

func (p *QingQLParser) Filter_condition() (localctx IFilter_conditionContext) {
	localctx = NewFilter_conditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, QingQLParserRULE_filter_condition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Filter_condition_or()
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QingQLParserAND {
		{
			p.SetState(104)
			p.Match(QingQLParserAND)
		}
		{
			p.SetState(105)
			p.Filter_condition_or()
		}

		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFilter_condition_orContext is an interface to support dynamic dispatch.
type IFilter_condition_orContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilter_condition_orContext differentiates from other interfaces.
	IsFilter_condition_orContext()
}

type Filter_condition_orContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilter_condition_orContext() *Filter_condition_orContext {
	var p = new(Filter_condition_orContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_filter_condition_or
	return p
}

func (*Filter_condition_orContext) IsFilter_condition_orContext() {}

func NewFilter_condition_orContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Filter_condition_orContext {
	var p = new(Filter_condition_orContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_filter_condition_or

	return p
}

func (s *Filter_condition_orContext) GetParser() antlr.Parser { return s.parser }

func (s *Filter_condition_orContext) AllFilter_condition_not() []IFilter_condition_notContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilter_condition_notContext)(nil)).Elem())
	var tst = make([]IFilter_condition_notContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilter_condition_notContext)
		}
	}

	return tst
}

func (s *Filter_condition_orContext) Filter_condition_not(i int) IFilter_condition_notContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilter_condition_notContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilter_condition_notContext)
}

func (s *Filter_condition_orContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(QingQLParserOR)
}

func (s *Filter_condition_orContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(QingQLParserOR, i)
}

func (s *Filter_condition_orContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Filter_condition_orContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Filter_condition_orContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFilter_condition_or(s)
	}
}

func (s *Filter_condition_orContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFilter_condition_or(s)
	}
}

func (p *QingQLParser) Filter_condition_or() (localctx IFilter_condition_orContext) {
	localctx = NewFilter_condition_orContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, QingQLParserRULE_filter_condition_or)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Filter_condition_not()
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QingQLParserOR {
		{
			p.SetState(112)
			p.Match(QingQLParserOR)
		}
		{
			p.SetState(113)
			p.Filter_condition_not()
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFilter_condition_notContext is an interface to support dynamic dispatch.
type IFilter_condition_notContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilter_condition_notContext differentiates from other interfaces.
	IsFilter_condition_notContext()
}

type Filter_condition_notContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilter_condition_notContext() *Filter_condition_notContext {
	var p = new(Filter_condition_notContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_filter_condition_not
	return p
}

func (*Filter_condition_notContext) IsFilter_condition_notContext() {}

func NewFilter_condition_notContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Filter_condition_notContext {
	var p = new(Filter_condition_notContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_filter_condition_not

	return p
}

func (s *Filter_condition_notContext) GetParser() antlr.Parser { return s.parser }

func (s *Filter_condition_notContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Filter_condition_notContext) NOT() antlr.TerminalNode {
	return s.GetToken(QingQLParserNOT, 0)
}

func (s *Filter_condition_notContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Filter_condition_notContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Filter_condition_notContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFilter_condition_not(s)
	}
}

func (s *Filter_condition_notContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFilter_condition_not(s)
	}
}

func (p *QingQLParser) Filter_condition_not() (localctx IFilter_condition_notContext) {
	localctx = NewFilter_condition_notContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, QingQLParserRULE_filter_condition_not)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QingQLParserNOT {
		{
			p.SetState(119)
			p.Match(QingQLParserNOT)
		}

	}
	{
		p.SetState(122)
		p.expr(0)
	}

	return localctx
}

// IGroupContext is an interface to support dynamic dispatch.
type IGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupContext() *GroupContext {
	var p = new(GroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) GROUP() antlr.TerminalNode {
	return s.GetToken(QingQLParserGROUP, 0)
}

func (s *GroupContext) BY() antlr.TerminalNode {
	return s.GetToken(QingQLParserBY, 0)
}

func (s *GroupContext) Group_by_elements() IGroup_by_elementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_by_elementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_by_elementsContext)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (p *QingQLParser) Group() (localctx IGroupContext) {
	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, QingQLParserRULE_group)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(QingQLParserGROUP)
	}
	{
		p.SetState(125)
		p.Match(QingQLParserBY)
	}
	{
		p.SetState(126)
		p.Group_by_elements()
	}

	return localctx
}

// IGroup_by_elementsContext is an interface to support dynamic dispatch.
type IGroup_by_elementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_by_elementsContext differentiates from other interfaces.
	IsGroup_by_elementsContext()
}

type Group_by_elementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_by_elementsContext() *Group_by_elementsContext {
	var p = new(Group_by_elementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_group_by_elements
	return p
}

func (*Group_by_elementsContext) IsGroup_by_elementsContext() {}

func NewGroup_by_elementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_by_elementsContext {
	var p = new(Group_by_elementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_group_by_elements

	return p
}

func (s *Group_by_elementsContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_by_elementsContext) AllGroup_by_element() []IGroup_by_elementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGroup_by_elementContext)(nil)).Elem())
	var tst = make([]IGroup_by_elementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGroup_by_elementContext)
		}
	}

	return tst
}

func (s *Group_by_elementsContext) Group_by_element(i int) IGroup_by_elementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_by_elementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGroup_by_elementContext)
}

func (s *Group_by_elementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_by_elementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_by_elementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterGroup_by_elements(s)
	}
}

func (s *Group_by_elementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitGroup_by_elements(s)
	}
}

func (p *QingQLParser) Group_by_elements() (localctx IGroup_by_elementsContext) {
	localctx = NewGroup_by_elementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, QingQLParserRULE_group_by_elements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Group_by_element()
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QingQLParserT__0 {
		{
			p.SetState(129)
			p.Match(QingQLParserT__0)
		}
		{
			p.SetState(130)
			p.Group_by_element()
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IGroup_by_elementContext is an interface to support dynamic dispatch.
type IGroup_by_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_by_elementContext differentiates from other interfaces.
	IsGroup_by_elementContext()
}

type Group_by_elementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_by_elementContext() *Group_by_elementContext {
	var p = new(Group_by_elementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_group_by_element
	return p
}

func (*Group_by_elementContext) IsGroup_by_elementContext() {}

func NewGroup_by_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_by_elementContext {
	var p = new(Group_by_elementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_group_by_element

	return p
}

func (s *Group_by_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_by_elementContext) CopyFrom(ctx *Group_by_elementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Group_by_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_by_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GroupWindowsContext struct {
	*Group_by_elementContext
}

func NewGroupWindowsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupWindowsContext {
	var p = new(GroupWindowsContext)

	p.Group_by_elementContext = NewEmptyGroup_by_elementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Group_by_elementContext))

	return p
}

func (s *GroupWindowsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupWindowsContext) Group_windows() IGroup_windowsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_windowsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_windowsContext)
}

func (s *GroupWindowsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterGroupWindows(s)
	}
}

func (s *GroupWindowsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitGroupWindows(s)
	}
}

type GroupFieldContext struct {
	*Group_by_elementContext
}

func NewGroupFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupFieldContext {
	var p = new(GroupFieldContext)

	p.Group_by_elementContext = NewEmptyGroup_by_elementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Group_by_elementContext))

	return p
}

func (s *GroupFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupFieldContext) Xpath_name() IXpath_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXpath_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IXpath_nameContext)
}

func (s *GroupFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterGroupField(s)
	}
}

func (s *GroupFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitGroupField(s)
	}
}

func (p *QingQLParser) Group_by_element() (localctx IGroup_by_elementContext) {
	localctx = NewGroup_by_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, QingQLParserRULE_group_by_element)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(138)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QingQLParserTUMBLINGWINDOW, QingQLParserHOPPINGWINDOW, QingQLParserSLIDINGWINDOW, QingQLParserSESSIONWINDOW:
		localctx = NewGroupWindowsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.Group_windows()
		}

	case QingQLParserT__4, QingQLParserMUL, QingQLParserINDENTIFIER, QingQLParserPATHITEM:
		localctx = NewGroupFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.Xpath_name()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGroup_window_unitContext is an interface to support dynamic dispatch.
type IGroup_window_unitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_window_unitContext differentiates from other interfaces.
	IsGroup_window_unitContext()
}

type Group_window_unitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_window_unitContext() *Group_window_unitContext {
	var p = new(Group_window_unitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_group_window_unit
	return p
}

func (*Group_window_unitContext) IsGroup_window_unitContext() {}

func NewGroup_window_unitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_window_unitContext {
	var p = new(Group_window_unitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_group_window_unit

	return p
}

func (s *Group_window_unitContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_window_unitContext) DD() antlr.TerminalNode {
	return s.GetToken(QingQLParserDD, 0)
}

func (s *Group_window_unitContext) HH() antlr.TerminalNode {
	return s.GetToken(QingQLParserHH, 0)
}

func (s *Group_window_unitContext) MI() antlr.TerminalNode {
	return s.GetToken(QingQLParserMI, 0)
}

func (s *Group_window_unitContext) SS() antlr.TerminalNode {
	return s.GetToken(QingQLParserSS, 0)
}

func (s *Group_window_unitContext) MS() antlr.TerminalNode {
	return s.GetToken(QingQLParserMS, 0)
}

func (s *Group_window_unitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_window_unitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_window_unitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterGroup_window_unit(s)
	}
}

func (s *Group_window_unitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitGroup_window_unit(s)
	}
}

func (p *QingQLParser) Group_window_unit() (localctx IGroup_window_unitContext) {
	localctx = NewGroup_window_unitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, QingQLParserRULE_group_window_unit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(140)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(QingQLParserDD-35))|(1<<(QingQLParserHH-35))|(1<<(QingQLParserMI-35))|(1<<(QingQLParserSS-35))|(1<<(QingQLParserMS-35)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IGroup_window_lengthContext is an interface to support dynamic dispatch.
type IGroup_window_lengthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_window_lengthContext differentiates from other interfaces.
	IsGroup_window_lengthContext()
}

type Group_window_lengthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_window_lengthContext() *Group_window_lengthContext {
	var p = new(Group_window_lengthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_group_window_length
	return p
}

func (*Group_window_lengthContext) IsGroup_window_lengthContext() {}

func NewGroup_window_lengthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_window_lengthContext {
	var p = new(Group_window_lengthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_group_window_length

	return p
}

func (s *Group_window_lengthContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_window_lengthContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(QingQLParserNUMBER, 0)
}

func (s *Group_window_lengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_window_lengthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_window_lengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterGroup_window_length(s)
	}
}

func (s *Group_window_lengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitGroup_window_length(s)
	}
}

func (p *QingQLParser) Group_window_length() (localctx IGroup_window_lengthContext) {
	localctx = NewGroup_window_lengthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, QingQLParserRULE_group_window_length)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(QingQLParserNUMBER)
	}

	return localctx
}

// IGroup_window_intervalContext is an interface to support dynamic dispatch.
type IGroup_window_intervalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_window_intervalContext differentiates from other interfaces.
	IsGroup_window_intervalContext()
}

type Group_window_intervalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_window_intervalContext() *Group_window_intervalContext {
	var p = new(Group_window_intervalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_group_window_interval
	return p
}

func (*Group_window_intervalContext) IsGroup_window_intervalContext() {}

func NewGroup_window_intervalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_window_intervalContext {
	var p = new(Group_window_intervalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_group_window_interval

	return p
}

func (s *Group_window_intervalContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_window_intervalContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(QingQLParserNUMBER, 0)
}

func (s *Group_window_intervalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_window_intervalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_window_intervalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterGroup_window_interval(s)
	}
}

func (s *Group_window_intervalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitGroup_window_interval(s)
	}
}

func (p *QingQLParser) Group_window_interval() (localctx IGroup_window_intervalContext) {
	localctx = NewGroup_window_intervalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, QingQLParserRULE_group_window_interval)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(QingQLParserNUMBER)
	}

	return localctx
}

// IGroup_windowsContext is an interface to support dynamic dispatch.
type IGroup_windowsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_windowsContext differentiates from other interfaces.
	IsGroup_windowsContext()
}

type Group_windowsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_windowsContext() *Group_windowsContext {
	var p = new(Group_windowsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_group_windows
	return p
}

func (*Group_windowsContext) IsGroup_windowsContext() {}

func NewGroup_windowsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_windowsContext {
	var p = new(Group_windowsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_group_windows

	return p
}

func (s *Group_windowsContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_windowsContext) Group_tumbling_window() IGroup_tumbling_windowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_tumbling_windowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_tumbling_windowContext)
}

func (s *Group_windowsContext) Group_hopping_window() IGroup_hopping_windowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_hopping_windowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_hopping_windowContext)
}

func (s *Group_windowsContext) Group_sliding_window() IGroup_sliding_windowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_sliding_windowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_sliding_windowContext)
}

func (s *Group_windowsContext) Group_session_window() IGroup_session_windowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_session_windowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_session_windowContext)
}

func (s *Group_windowsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_windowsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_windowsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterGroup_windows(s)
	}
}

func (s *Group_windowsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitGroup_windows(s)
	}
}

func (p *QingQLParser) Group_windows() (localctx IGroup_windowsContext) {
	localctx = NewGroup_windowsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, QingQLParserRULE_group_windows)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(146)
			p.Group_tumbling_window()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)
			p.Group_hopping_window()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(148)
			p.Group_sliding_window()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(149)
			p.Group_session_window()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(150)
			p.Group_hopping_window()
		}

	}

	return localctx
}

// IGroup_tumbling_windowContext is an interface to support dynamic dispatch.
type IGroup_tumbling_windowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_tumbling_windowContext differentiates from other interfaces.
	IsGroup_tumbling_windowContext()
}

type Group_tumbling_windowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_tumbling_windowContext() *Group_tumbling_windowContext {
	var p = new(Group_tumbling_windowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_group_tumbling_window
	return p
}

func (*Group_tumbling_windowContext) IsGroup_tumbling_windowContext() {}

func NewGroup_tumbling_windowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_tumbling_windowContext {
	var p = new(Group_tumbling_windowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_group_tumbling_window

	return p
}

func (s *Group_tumbling_windowContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_tumbling_windowContext) TUMBLINGWINDOW() antlr.TerminalNode {
	return s.GetToken(QingQLParserTUMBLINGWINDOW, 0)
}

func (s *Group_tumbling_windowContext) Group_window_unit() IGroup_window_unitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_window_unitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_window_unitContext)
}

func (s *Group_tumbling_windowContext) Group_window_interval() IGroup_window_intervalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_window_intervalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_window_intervalContext)
}

func (s *Group_tumbling_windowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_tumbling_windowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_tumbling_windowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterGroup_tumbling_window(s)
	}
}

func (s *Group_tumbling_windowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitGroup_tumbling_window(s)
	}
}

func (p *QingQLParser) Group_tumbling_window() (localctx IGroup_tumbling_windowContext) {
	localctx = NewGroup_tumbling_windowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, QingQLParserRULE_group_tumbling_window)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(QingQLParserTUMBLINGWINDOW)
	}
	{
		p.SetState(154)
		p.Match(QingQLParserT__2)
	}
	{
		p.SetState(155)
		p.Group_window_unit()
	}
	{
		p.SetState(156)
		p.Match(QingQLParserT__0)
	}
	{
		p.SetState(157)
		p.Group_window_interval()
	}
	{
		p.SetState(158)
		p.Match(QingQLParserT__3)
	}

	return localctx
}

// IGroup_hopping_windowContext is an interface to support dynamic dispatch.
type IGroup_hopping_windowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_hopping_windowContext differentiates from other interfaces.
	IsGroup_hopping_windowContext()
}

type Group_hopping_windowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_hopping_windowContext() *Group_hopping_windowContext {
	var p = new(Group_hopping_windowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_group_hopping_window
	return p
}

func (*Group_hopping_windowContext) IsGroup_hopping_windowContext() {}

func NewGroup_hopping_windowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_hopping_windowContext {
	var p = new(Group_hopping_windowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_group_hopping_window

	return p
}

func (s *Group_hopping_windowContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_hopping_windowContext) HOPPINGWINDOW() antlr.TerminalNode {
	return s.GetToken(QingQLParserHOPPINGWINDOW, 0)
}

func (s *Group_hopping_windowContext) Group_window_unit() IGroup_window_unitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_window_unitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_window_unitContext)
}

func (s *Group_hopping_windowContext) Group_window_interval() IGroup_window_intervalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_window_intervalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_window_intervalContext)
}

func (s *Group_hopping_windowContext) Group_window_length() IGroup_window_lengthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_window_lengthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_window_lengthContext)
}

func (s *Group_hopping_windowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_hopping_windowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_hopping_windowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterGroup_hopping_window(s)
	}
}

func (s *Group_hopping_windowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitGroup_hopping_window(s)
	}
}

func (p *QingQLParser) Group_hopping_window() (localctx IGroup_hopping_windowContext) {
	localctx = NewGroup_hopping_windowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, QingQLParserRULE_group_hopping_window)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(QingQLParserHOPPINGWINDOW)
	}
	{
		p.SetState(161)
		p.Match(QingQLParserT__2)
	}
	{
		p.SetState(162)
		p.Group_window_unit()
	}
	{
		p.SetState(163)
		p.Match(QingQLParserT__0)
	}
	{
		p.SetState(164)
		p.Group_window_interval()
	}
	{
		p.SetState(165)
		p.Match(QingQLParserT__0)
	}
	{
		p.SetState(166)
		p.Group_window_length()
	}
	{
		p.SetState(167)
		p.Match(QingQLParserT__3)
	}

	return localctx
}

// IGroup_sliding_windowContext is an interface to support dynamic dispatch.
type IGroup_sliding_windowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_sliding_windowContext differentiates from other interfaces.
	IsGroup_sliding_windowContext()
}

type Group_sliding_windowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_sliding_windowContext() *Group_sliding_windowContext {
	var p = new(Group_sliding_windowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_group_sliding_window
	return p
}

func (*Group_sliding_windowContext) IsGroup_sliding_windowContext() {}

func NewGroup_sliding_windowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_sliding_windowContext {
	var p = new(Group_sliding_windowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_group_sliding_window

	return p
}

func (s *Group_sliding_windowContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_sliding_windowContext) SLIDINGWINDOW() antlr.TerminalNode {
	return s.GetToken(QingQLParserSLIDINGWINDOW, 0)
}

func (s *Group_sliding_windowContext) Group_window_unit() IGroup_window_unitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_window_unitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_window_unitContext)
}

func (s *Group_sliding_windowContext) Group_window_interval() IGroup_window_intervalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_window_intervalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_window_intervalContext)
}

func (s *Group_sliding_windowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_sliding_windowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_sliding_windowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterGroup_sliding_window(s)
	}
}

func (s *Group_sliding_windowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitGroup_sliding_window(s)
	}
}

func (p *QingQLParser) Group_sliding_window() (localctx IGroup_sliding_windowContext) {
	localctx = NewGroup_sliding_windowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, QingQLParserRULE_group_sliding_window)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(QingQLParserSLIDINGWINDOW)
	}
	{
		p.SetState(170)
		p.Match(QingQLParserT__2)
	}
	{
		p.SetState(171)
		p.Group_window_unit()
	}
	{
		p.SetState(172)
		p.Match(QingQLParserT__0)
	}
	{
		p.SetState(173)
		p.Group_window_interval()
	}
	{
		p.SetState(174)
		p.Match(QingQLParserT__3)
	}

	return localctx
}

// IGroup_session_windowContext is an interface to support dynamic dispatch.
type IGroup_session_windowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_session_windowContext differentiates from other interfaces.
	IsGroup_session_windowContext()
}

type Group_session_windowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_session_windowContext() *Group_session_windowContext {
	var p = new(Group_session_windowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_group_session_window
	return p
}

func (*Group_session_windowContext) IsGroup_session_windowContext() {}

func NewGroup_session_windowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_session_windowContext {
	var p = new(Group_session_windowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_group_session_window

	return p
}

func (s *Group_session_windowContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_session_windowContext) SESSIONWINDOW() antlr.TerminalNode {
	return s.GetToken(QingQLParserSESSIONWINDOW, 0)
}

func (s *Group_session_windowContext) Group_window_unit() IGroup_window_unitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_window_unitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_window_unitContext)
}

func (s *Group_session_windowContext) Group_window_interval() IGroup_window_intervalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_window_intervalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_window_intervalContext)
}

func (s *Group_session_windowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_session_windowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_session_windowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterGroup_session_window(s)
	}
}

func (s *Group_session_windowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitGroup_session_window(s)
	}
}

func (p *QingQLParser) Group_session_window() (localctx IGroup_session_windowContext) {
	localctx = NewGroup_session_windowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, QingQLParserRULE_group_session_window)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(QingQLParserSESSIONWINDOW)
	}
	{
		p.SetState(177)
		p.Match(QingQLParserT__2)
	}
	{
		p.SetState(178)
		p.Group_window_unit()
	}
	{
		p.SetState(179)
		p.Match(QingQLParserT__0)
	}
	{
		p.SetState(180)
		p.Group_window_interval()
	}
	{
		p.SetState(181)
		p.Match(QingQLParserT__3)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionContext struct {
	*ExprContext
}

func NewFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionContext {
	var p = new(FunctionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) Call_expr() ICall_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICall_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICall_exprContext)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFunction(s)
	}
}

type BracesContext struct {
	*ExprContext
}

func NewBracesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracesContext {
	var p = new(BracesContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BracesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracesContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *BracesContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BracesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterBraces(s)
	}
}

func (s *BracesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitBraces(s)
	}
}

type SwitchContext struct {
	*ExprContext
}

func NewSwitchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchContext {
	var p = new(SwitchContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchContext) Switch_stmt() ISwitch_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitch_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwitch_stmtContext)
}

func (s *SwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterSwitch(s)
	}
}

func (s *SwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitSwitch(s)
	}
}

type BinaryContext struct {
	*ExprContext
	op antlr.Token
}

func NewBinaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryContext {
	var p = new(BinaryContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BinaryContext) GetOp() antlr.Token { return s.op }

func (s *BinaryContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *BinaryContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BinaryContext) EQ() antlr.TerminalNode {
	return s.GetToken(QingQLParserEQ, 0)
}

func (s *BinaryContext) GT() antlr.TerminalNode {
	return s.GetToken(QingQLParserGT, 0)
}

func (s *BinaryContext) LT() antlr.TerminalNode {
	return s.GetToken(QingQLParserLT, 0)
}

func (s *BinaryContext) GTE() antlr.TerminalNode {
	return s.GetToken(QingQLParserGTE, 0)
}

func (s *BinaryContext) LTE() antlr.TerminalNode {
	return s.GetToken(QingQLParserLTE, 0)
}

func (s *BinaryContext) NE() antlr.TerminalNode {
	return s.GetToken(QingQLParserNE, 0)
}

func (s *BinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterBinary(s)
	}
}

func (s *BinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitBinary(s)
	}
}

func (p *QingQLParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *QingQLParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, QingQLParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBracesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(184)
			p.Constant()
		}

	case 2:
		localctx = NewBracesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(185)
			p.Match(QingQLParserT__2)
		}
		{
			p.SetState(186)
			p.expr(0)
		}
		{
			p.SetState(187)
			p.Match(QingQLParserT__3)
		}

	case 3:
		localctx = NewFunctionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(189)
			p.Call_expr()
		}

	case 4:
		localctx = NewSwitchContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(190)
			p.Switch_stmt()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(202)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QingQLParserRULE_expr)
				p.SetState(193)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				p.SetState(194)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*BinaryContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(QingQLParserMUL-40))|(1<<(QingQLParserDIV-40))|(1<<(QingQLParserMOD-40)))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*BinaryContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(195)
					p.expr(6)
				}

			case 2:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QingQLParserRULE_expr)
				p.SetState(196)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				p.SetState(197)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*BinaryContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == QingQLParserADD || _la == QingQLParserSUB) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*BinaryContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(198)
					p.expr(5)
				}

			case 3:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QingQLParserRULE_expr)
				p.SetState(199)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				p.SetState(200)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*BinaryContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QingQLParserEQ)|(1<<QingQLParserGT)|(1<<QingQLParserGTE)|(1<<QingQLParserLT)|(1<<QingQLParserLTE)|(1<<QingQLParserNE))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*BinaryContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(201)
					p.expr(4)
				}

			}

		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) CopyFrom(ctx *ConstantContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IntegerContext struct {
	*ConstantContext
}

func NewIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerContext {
	var p = new(IntegerContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(QingQLParserNUMBER, 0)
}

func (s *IntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(QingQLParserINTEGER, 0)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitInteger(s)
	}
}

type XPathContext struct {
	*ConstantContext
}

func NewXPathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *XPathContext {
	var p = new(XPathContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *XPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XPathContext) Xpath_name() IXpath_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXpath_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IXpath_nameContext)
}

func (s *XPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterXPath(s)
	}
}

func (s *XPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitXPath(s)
	}
}

type FloatContext struct {
	*ConstantContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(QingQLParserFLOAT, 0)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFloat(s)
	}
}

type StringContext struct {
	*ConstantContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(QingQLParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitString(s)
	}
}

type BooleanContext struct {
	*ConstantContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) TRUE() antlr.TerminalNode {
	return s.GetToken(QingQLParserTRUE, 0)
}

func (s *BooleanContext) FALSE() antlr.TerminalNode {
	return s.GetToken(QingQLParserFALSE, 0)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (p *QingQLParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, QingQLParserRULE_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(214)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QingQLParserTRUE:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(207)
			p.Match(QingQLParserTRUE)
		}

	case QingQLParserFALSE:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(208)
			p.Match(QingQLParserFALSE)
		}

	case QingQLParserNUMBER:
		localctx = NewIntegerContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(209)
			p.Match(QingQLParserNUMBER)
		}

	case QingQLParserINTEGER:
		localctx = NewIntegerContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(210)
			p.Match(QingQLParserINTEGER)
		}

	case QingQLParserFLOAT:
		localctx = NewFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(211)
			p.Match(QingQLParserFLOAT)
		}

	case QingQLParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(212)
			p.Match(QingQLParserSTRING)
		}

	case QingQLParserT__4, QingQLParserMUL, QingQLParserINDENTIFIER, QingQLParserPATHITEM:
		localctx = NewXPathContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(213)
			p.Xpath_name()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISwitch_stmtContext is an interface to support dynamic dispatch.
type ISwitch_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitch_stmtContext differentiates from other interfaces.
	IsSwitch_stmtContext()
}

type Switch_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_stmtContext() *Switch_stmtContext {
	var p = new(Switch_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_switch_stmt
	return p
}

func (*Switch_stmtContext) IsSwitch_stmtContext() {}

func NewSwitch_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_stmtContext {
	var p = new(Switch_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_switch_stmt

	return p
}

func (s *Switch_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_stmtContext) CASE() antlr.TerminalNode {
	return s.GetToken(QingQLParserCASE, 0)
}

func (s *Switch_stmtContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Switch_stmtContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Switch_stmtContext) AllWHEN() []antlr.TerminalNode {
	return s.GetTokens(QingQLParserWHEN)
}

func (s *Switch_stmtContext) WHEN(i int) antlr.TerminalNode {
	return s.GetToken(QingQLParserWHEN, i)
}

func (s *Switch_stmtContext) AllTHEN() []antlr.TerminalNode {
	return s.GetTokens(QingQLParserTHEN)
}

func (s *Switch_stmtContext) THEN(i int) antlr.TerminalNode {
	return s.GetToken(QingQLParserTHEN, i)
}

func (s *Switch_stmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(QingQLParserELSE, 0)
}

func (s *Switch_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterSwitch_stmt(s)
	}
}

func (s *Switch_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitSwitch_stmt(s)
	}
}

func (p *QingQLParser) Switch_stmt() (localctx ISwitch_stmtContext) {
	localctx = NewSwitch_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, QingQLParserRULE_switch_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(QingQLParserCASE)
	}
	{
		p.SetState(217)
		p.expr(0)
	}
	{
		p.SetState(218)
		p.Match(QingQLParserWHEN)
	}
	{
		p.SetState(219)
		p.expr(0)
	}
	{
		p.SetState(220)
		p.Match(QingQLParserTHEN)
	}
	{
		p.SetState(221)
		p.expr(0)
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(222)
				p.Match(QingQLParserWHEN)
			}
			{
				p.SetState(223)
				p.expr(0)
			}
			{
				p.SetState(224)
				p.Match(QingQLParserTHEN)
			}
			{
				p.SetState(225)
				p.expr(0)
			}

		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(232)
			p.Match(QingQLParserELSE)
		}
		{
			p.SetState(233)
			p.expr(0)
		}

	}

	return localctx
}

// ICall_exprContext is an interface to support dynamic dispatch.
type ICall_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// IsCall_exprContext differentiates from other interfaces.
	IsCall_exprContext()
}

type Call_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
}

func NewEmptyCall_exprContext() *Call_exprContext {
	var p = new(Call_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_call_expr
	return p
}

func (*Call_exprContext) IsCall_exprContext() {}

func NewCall_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_exprContext {
	var p = new(Call_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_call_expr

	return p
}

func (s *Call_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_exprContext) GetKey() antlr.Token { return s.key }

func (s *Call_exprContext) SetKey(v antlr.Token) { s.key = v }

func (s *Call_exprContext) INDENTIFIER() antlr.TerminalNode {
	return s.GetToken(QingQLParserINDENTIFIER, 0)
}

func (s *Call_exprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Call_exprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Call_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterCall_expr(s)
	}
}

func (s *Call_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitCall_expr(s)
	}
}

func (p *QingQLParser) Call_expr() (localctx ICall_exprContext) {
	localctx = NewCall_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, QingQLParserRULE_call_expr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)

		var _m = p.Match(QingQLParserINDENTIFIER)

		localctx.(*Call_exprContext).key = _m
	}
	{
		p.SetState(237)
		p.Match(QingQLParserT__2)
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QingQLParserT__2)|(1<<QingQLParserT__4)|(1<<QingQLParserCASE))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(QingQLParserMUL-40))|(1<<(QingQLParserTRUE-40))|(1<<(QingQLParserFALSE-40))|(1<<(QingQLParserINDENTIFIER-40))|(1<<(QingQLParserNUMBER-40))|(1<<(QingQLParserINTEGER-40))|(1<<(QingQLParserFLOAT-40))|(1<<(QingQLParserPATHITEM-40))|(1<<(QingQLParserSTRING-40)))) != 0) {
		{
			p.SetState(238)
			p.expr(0)
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QingQLParserT__0 {
			{
				p.SetState(239)
				p.Match(QingQLParserT__0)
			}
			{
				p.SetState(240)
				p.expr(0)
			}

			p.SetState(245)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(248)
		p.Match(QingQLParserT__3)
	}

	return localctx
}

// IAsteriskContext is an interface to support dynamic dispatch.
type IAsteriskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAsteriskContext differentiates from other interfaces.
	IsAsteriskContext()
}

type AsteriskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsteriskContext() *AsteriskContext {
	var p = new(AsteriskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_asterisk
	return p
}

func (*AsteriskContext) IsAsteriskContext() {}

func NewAsteriskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsteriskContext {
	var p = new(AsteriskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_asterisk

	return p
}

func (s *AsteriskContext) GetParser() antlr.Parser { return s.parser }
func (s *AsteriskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsteriskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsteriskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterAsterisk(s)
	}
}

func (s *AsteriskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitAsterisk(s)
	}
}

func (p *QingQLParser) Asterisk() (localctx IAsteriskContext) {
	localctx = NewAsteriskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, QingQLParserRULE_asterisk)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.Match(QingQLParserMUL)
	}

	return localctx
}

// IXpath_nameContext is an interface to support dynamic dispatch.
type IXpath_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsXpath_nameContext differentiates from other interfaces.
	IsXpath_nameContext()
}

type Xpath_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXpath_nameContext() *Xpath_nameContext {
	var p = new(Xpath_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_xpath_name
	return p
}

func (*Xpath_nameContext) IsXpath_nameContext() {}

func NewXpath_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Xpath_nameContext {
	var p = new(Xpath_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_xpath_name

	return p
}

func (s *Xpath_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Xpath_nameContext) Asterisk() IAsteriskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsteriskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsteriskContext)
}

func (s *Xpath_nameContext) AllDotnotation() []IDotnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDotnotationContext)(nil)).Elem())
	var tst = make([]IDotnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDotnotationContext)
		}
	}

	return tst
}

func (s *Xpath_nameContext) Dotnotation(i int) IDotnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDotnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDotnotationContext)
}

func (s *Xpath_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Xpath_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Xpath_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterXpath_name(s)
	}
}

func (s *Xpath_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitXpath_name(s)
	}
}

func (p *QingQLParser) Xpath_name() (localctx IXpath_nameContext) {
	localctx = NewXpath_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, QingQLParserRULE_xpath_name)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(266)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QingQLParserMUL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(252)
			p.Asterisk()
		}

	case QingQLParserINDENTIFIER, QingQLParserPATHITEM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
			p.Dotnotation()
		}

	case QingQLParserT__4:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QingQLParserT__4 {
			{
				p.SetState(254)
				p.Match(QingQLParserT__4)
			}

			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QingQLParserINDENTIFIER || _la == QingQLParserPATHITEM {
			{
				p.SetState(259)
				p.Dotnotation()
			}

			p.SetState(262)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(264)
			p.Match(QingQLParserT__4)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDotnotationContext is an interface to support dynamic dispatch.
type IDotnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDotnotationContext differentiates from other interfaces.
	IsDotnotationContext()
}

type DotnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDotnotationContext() *DotnotationContext {
	var p = new(DotnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_dotnotation
	return p
}

func (*DotnotationContext) IsDotnotationContext() {}

func NewDotnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DotnotationContext {
	var p = new(DotnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_dotnotation

	return p
}

func (s *DotnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *DotnotationContext) INDENTIFIER() antlr.TerminalNode {
	return s.GetToken(QingQLParserINDENTIFIER, 0)
}

func (s *DotnotationContext) PATHITEM() antlr.TerminalNode {
	return s.GetToken(QingQLParserPATHITEM, 0)
}

func (s *DotnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DotnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterDotnotation(s)
	}
}

func (s *DotnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitDotnotation(s)
	}
}

func (p *QingQLParser) Dotnotation() (localctx IDotnotationContext) {
	localctx = NewDotnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, QingQLParserRULE_dotnotation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(268)
	_la = p.GetTokenStream().LA(1)

	if !(_la == QingQLParserINDENTIFIER || _la == QingQLParserPATHITEM) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IIdentifierWithTOPICITEMContext is an interface to support dynamic dispatch.
type IIdentifierWithTOPICITEMContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierWithTOPICITEMContext differentiates from other interfaces.
	IsIdentifierWithTOPICITEMContext()
}

type IdentifierWithTOPICITEMContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierWithTOPICITEMContext() *IdentifierWithTOPICITEMContext {
	var p = new(IdentifierWithTOPICITEMContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_identifierWithTOPICITEM
	return p
}

func (*IdentifierWithTOPICITEMContext) IsIdentifierWithTOPICITEMContext() {}

func NewIdentifierWithTOPICITEMContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierWithTOPICITEMContext {
	var p = new(IdentifierWithTOPICITEMContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_identifierWithTOPICITEM

	return p
}

func (s *IdentifierWithTOPICITEMContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierWithTOPICITEMContext) PATHITEM() antlr.TerminalNode {
	return s.GetToken(QingQLParserPATHITEM, 0)
}

func (s *IdentifierWithTOPICITEMContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(QingQLParserNUMBER, 0)
}

func (s *IdentifierWithTOPICITEMContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(QingQLParserFLOAT, 0)
}

func (s *IdentifierWithTOPICITEMContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierWithTOPICITEMContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierWithTOPICITEMContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterIdentifierWithTOPICITEM(s)
	}
}

func (s *IdentifierWithTOPICITEMContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitIdentifierWithTOPICITEM(s)
	}
}

func (p *QingQLParser) IdentifierWithTOPICITEM() (localctx IIdentifierWithTOPICITEMContext) {
	localctx = NewIdentifierWithTOPICITEMContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, QingQLParserRULE_identifierWithTOPICITEM)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(270)
			p.Match(QingQLParserPATHITEM)
		}
		{
			p.SetState(271)
			p.Match(QingQLParserT__5)
		}
		{
			p.SetState(272)
			p.Match(QingQLParserT__6)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(273)
			p.Match(QingQLParserPATHITEM)
		}
		{
			p.SetState(274)
			p.Match(QingQLParserT__5)
		}
		{
			p.SetState(275)
			p.Match(QingQLParserNUMBER)
		}
		{
			p.SetState(276)
			p.Match(QingQLParserT__6)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(277)
			p.Match(QingQLParserPATHITEM)
		}
		{
			p.SetState(278)
			p.Match(QingQLParserT__5)
		}
		{
			p.SetState(279)
			p.Match(QingQLParserT__1)
		}
		{
			p.SetState(280)
			p.Match(QingQLParserT__6)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(281)
			p.Match(QingQLParserPATHITEM)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(282)
			p.Match(QingQLParserFLOAT)
		}

	}

	return localctx
}

// IIdentifierWithQualifierContext is an interface to support dynamic dispatch.
type IIdentifierWithQualifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierWithQualifierContext differentiates from other interfaces.
	IsIdentifierWithQualifierContext()
}

type IdentifierWithQualifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierWithQualifierContext() *IdentifierWithQualifierContext {
	var p = new(IdentifierWithQualifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_identifierWithQualifier
	return p
}

func (*IdentifierWithQualifierContext) IsIdentifierWithQualifierContext() {}

func NewIdentifierWithQualifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierWithQualifierContext {
	var p = new(IdentifierWithQualifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_identifierWithQualifier

	return p
}

func (s *IdentifierWithQualifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierWithQualifierContext) INDENTIFIER() antlr.TerminalNode {
	return s.GetToken(QingQLParserINDENTIFIER, 0)
}

func (s *IdentifierWithQualifierContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(QingQLParserNUMBER, 0)
}

func (s *IdentifierWithQualifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierWithQualifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierWithQualifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterIdentifierWithQualifier(s)
	}
}

func (s *IdentifierWithQualifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitIdentifierWithQualifier(s)
	}
}

func (p *QingQLParser) IdentifierWithQualifier() (localctx IIdentifierWithQualifierContext) {
	localctx = NewIdentifierWithQualifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, QingQLParserRULE_identifierWithQualifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
			p.Match(QingQLParserINDENTIFIER)
		}
		{
			p.SetState(286)
			p.Match(QingQLParserT__7)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(287)
			p.Match(QingQLParserINDENTIFIER)
		}
		{
			p.SetState(288)
			p.Match(QingQLParserT__5)
		}
		{
			p.SetState(289)
			p.Match(QingQLParserNUMBER)
		}
		{
			p.SetState(290)
			p.Match(QingQLParserT__6)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(291)
			p.Match(QingQLParserINDENTIFIER)
		}
		{
			p.SetState(292)
			p.Match(QingQLParserT__8)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(293)
			p.Match(QingQLParserINDENTIFIER)
		}

	}

	return localctx
}

func (p *QingQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 21:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *QingQLParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
