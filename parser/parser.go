//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/mattn/anko/ast"
)

//line parser.go.y:29
type yySymType struct {
	yys             int
	compstmt        []ast.Stmt
	stmt_if         ast.Stmt
	stmt_default    ast.Stmt
	stmt_case       ast.Stmt
	stmt_cases      []ast.Stmt
	stmts           []ast.Stmt
	stmt            ast.Stmt
	expr            ast.Expr
	exprs           []ast.Expr
	expr_many       []ast.Expr
	expr_lets       ast.Expr
	expr_pair       ast.Expr
	expr_pairs      []ast.Expr
	expr_idents     []string
	expr_type       string
	tok             ast.Token
	term            ast.Token
	terms           ast.Token
	opt_terms       ast.Token
	array_count     ast.ArrayCount
	expr_slice      ast.SliceExpr
	stmt_multi_case []ast.Stmt
}

const IDENT = 57346
const NUMBER = 57347
const STRING = 57348
const ARRAY = 57349
const VARARG = 57350
const FUNC = 57351
const RETURN = 57352
const VAR = 57353
const THROW = 57354
const IF = 57355
const ELSE = 57356
const FOR = 57357
const IN = 57358
const EQEQ = 57359
const NEQ = 57360
const GE = 57361
const LE = 57362
const OROR = 57363
const ANDAND = 57364
const NEW = 57365
const TRUE = 57366
const FALSE = 57367
const NIL = 57368
const MODULE = 57369
const TRY = 57370
const CATCH = 57371
const FINALLY = 57372
const PLUSEQ = 57373
const MINUSEQ = 57374
const MULEQ = 57375
const DIVEQ = 57376
const ANDEQ = 57377
const OREQ = 57378
const BREAK = 57379
const CONTINUE = 57380
const PLUSPLUS = 57381
const MINUSMINUS = 57382
const POW = 57383
const SHIFTLEFT = 57384
const SHIFTRIGHT = 57385
const SWITCH = 57386
const CASE = 57387
const DEFAULT = 57388
const GO = 57389
const DEFER = 57390
const CHAN = 57391
const MAKE = 57392
const OPCHAN = 57393
const TYPE = 57394
const LEN = 57395
const DELETE = 57396
const UNARY = 57397

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"NUMBER",
	"STRING",
	"ARRAY",
	"VARARG",
	"FUNC",
	"RETURN",
	"VAR",
	"THROW",
	"IF",
	"ELSE",
	"FOR",
	"IN",
	"EQEQ",
	"NEQ",
	"GE",
	"LE",
	"OROR",
	"ANDAND",
	"NEW",
	"TRUE",
	"FALSE",
	"NIL",
	"MODULE",
	"TRY",
	"CATCH",
	"FINALLY",
	"PLUSEQ",
	"MINUSEQ",
	"MULEQ",
	"DIVEQ",
	"ANDEQ",
	"OREQ",
	"BREAK",
	"CONTINUE",
	"PLUSPLUS",
	"MINUSMINUS",
	"POW",
	"SHIFTLEFT",
	"SHIFTRIGHT",
	"SWITCH",
	"CASE",
	"DEFAULT",
	"GO",
	"DEFER",
	"CHAN",
	"MAKE",
	"OPCHAN",
	"TYPE",
	"LEN",
	"DELETE",
	"'='",
	"'?'",
	"':'",
	"','",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UNARY",
	"'{'",
	"'}'",
	"';'",
	"'('",
	"')'",
	"'['",
	"']'",
	"'.'",
	"'!'",
	"'^'",
	"'&'",
	"'|'",
	"'\\n'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:842

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 3,
	-2, 143,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	58, 62,
	-2, 1,
	-1, 10,
	58, 63,
	-2, 32,
	-1, 47,
	58, 62,
	-2, 144,
	-1, 90,
	68, 3,
	-2, 143,
	-1, 93,
	58, 63,
	-2, 59,
	-1, 94,
	16, 54,
	58, 54,
	-2, 72,
	-1, 96,
	68, 3,
	-2, 143,
	-1, 107,
	1, 77,
	8, 77,
	45, 77,
	46, 77,
	55, 77,
	57, 77,
	58, 77,
	67, 77,
	68, 77,
	69, 77,
	71, 77,
	73, 77,
	79, 77,
	-2, 72,
	-1, 109,
	1, 79,
	8, 79,
	45, 79,
	46, 79,
	55, 79,
	57, 79,
	58, 79,
	67, 79,
	68, 79,
	69, 79,
	71, 79,
	73, 79,
	79, 79,
	-2, 72,
	-1, 139,
	17, 0,
	18, 0,
	-2, 104,
	-1, 140,
	17, 0,
	18, 0,
	-2, 105,
	-1, 160,
	58, 63,
	-2, 59,
	-1, 162,
	68, 3,
	-2, 143,
	-1, 164,
	68, 3,
	-2, 143,
	-1, 166,
	68, 1,
	-2, 50,
	-1, 169,
	68, 3,
	-2, 143,
	-1, 198,
	68, 3,
	-2, 143,
	-1, 249,
	58, 64,
	-2, 60,
	-1, 250,
	1, 61,
	45, 61,
	46, 61,
	55, 61,
	57, 61,
	58, 65,
	68, 61,
	69, 61,
	79, 61,
	-2, 72,
	-1, 259,
	1, 65,
	8, 65,
	45, 65,
	46, 65,
	58, 65,
	68, 65,
	69, 65,
	71, 65,
	73, 65,
	79, 65,
	-2, 72,
	-1, 261,
	68, 3,
	-2, 143,
	-1, 263,
	68, 3,
	-2, 143,
	-1, 277,
	1, 25,
	45, 25,
	46, 25,
	68, 25,
	69, 25,
	79, 25,
	-2, 123,
	-1, 279,
	1, 27,
	45, 27,
	46, 27,
	68, 27,
	69, 27,
	79, 27,
	-2, 125,
	-1, 281,
	1, 29,
	45, 29,
	46, 29,
	68, 29,
	69, 29,
	79, 29,
	-2, 123,
	-1, 283,
	1, 31,
	45, 31,
	46, 31,
	68, 31,
	69, 31,
	79, 31,
	-2, 125,
	-1, 288,
	68, 3,
	-2, 143,
	-1, 311,
	68, 3,
	-2, 143,
	-1, 315,
	1, 24,
	45, 24,
	46, 24,
	68, 24,
	69, 24,
	79, 24,
	-2, 122,
	-1, 316,
	1, 26,
	45, 26,
	46, 26,
	68, 26,
	69, 26,
	79, 26,
	-2, 124,
	-1, 317,
	1, 28,
	45, 28,
	46, 28,
	68, 28,
	69, 28,
	79, 28,
	-2, 122,
	-1, 318,
	1, 30,
	45, 30,
	46, 30,
	68, 30,
	69, 30,
	79, 30,
	-2, 124,
	-1, 321,
	68, 3,
	-2, 143,
	-1, 322,
	68, 3,
	-2, 143,
	-1, 333,
	68, 3,
	-2, 143,
	-1, 334,
	68, 3,
	-2, 143,
	-1, 338,
	45, 3,
	46, 3,
	68, 3,
	-2, 143,
	-1, 342,
	68, 3,
	-2, 143,
	-1, 350,
	45, 3,
	46, 3,
	68, 3,
	-2, 143,
	-1, 351,
	45, 3,
	46, 3,
	68, 3,
	-2, 143,
	-1, 365,
	68, 3,
	-2, 143,
	-1, 366,
	68, 3,
	-2, 143,
}

const yyPrivate = 57344

const yyLast = 3196

var yyAct = [...]int{

	86, 186, 269, 10, 302, 38, 268, 270, 6, 300,
	190, 293, 49, 1, 236, 11, 87, 301, 7, 93,
	240, 97, 99, 101, 299, 6, 104, 105, 106, 108,
	110, 91, 246, 241, 95, 7, 240, 298, 115, 103,
	240, 102, 175, 119, 102, 191, 122, 173, 10, 102,
	323, 2, 126, 127, 129, 46, 131, 132, 133, 134,
	135, 136, 137, 138, 139, 140, 141, 142, 143, 144,
	145, 146, 147, 148, 149, 150, 44, 318, 151, 152,
	153, 154, 291, 156, 158, 160, 282, 113, 114, 157,
	234, 85, 69, 70, 71, 72, 73, 74, 280, 317,
	159, 6, 60, 177, 165, 192, 316, 278, 194, 276,
	171, 7, 82, 243, 315, 289, 287, 189, 257, 112,
	183, 196, 120, 118, 182, 160, 187, 117, 195, 203,
	116, 81, 125, 52, 232, 54, 161, 227, 79, 77,
	199, 168, 166, 205, 370, 290, 310, 369, 161, 283,
	362, 271, 272, 358, 357, 354, 353, 161, 155, 161,
	349, 281, 339, 332, 209, 331, 305, 10, 213, 214,
	279, 160, 277, 197, 267, 295, 208, 200, 210, 226,
	179, 265, 262, 215, 125, 111, 216, 161, 326, 260,
	184, 217, 211, 161, 163, 366, 233, 231, 249, 365,
	228, 342, 253, 242, 244, 256, 206, 334, 258, 311,
	322, 321, 251, 207, 288, 162, 96, 167, 124, 303,
	161, 125, 275, 121, 219, 284, 273, 274, 337, 314,
	238, 170, 84, 271, 272, 364, 235, 237, 359, 296,
	266, 5, 8, 89, 187, 297, 48, 252, 164, 304,
	220, 221, 222, 223, 245, 191, 50, 230, 229, 125,
	130, 259, 24, 30, 88, 309, 34, 4, 193, 185,
	92, 47, 312, 218, 17, 307, 3, 308, 0, 0,
	40, 31, 32, 33, 0, 0, 292, 313, 294, 48,
	123, 0, 0, 258, 0, 0, 325, 0, 0, 0,
	327, 0, 320, 328, 329, 0, 0, 41, 42, 0,
	39, 43, 0, 0, 0, 0, 0, 0, 0, 25,
	29, 0, 0, 0, 36, 335, 0, 37, 0, 35,
	324, 0, 26, 27, 28, 340, 341, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 356, 347, 348, 0,
	0, 0, 352, 0, 0, 0, 355, 0, 0, 0,
	0, 0, 0, 0, 360, 361, 338, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 23, 24, 30, 367,
	368, 34, 14, 9, 15, 45, 0, 18, 350, 351,
	0, 0, 0, 0, 0, 40, 31, 32, 33, 16,
	19, 0, 0, 0, 0, 0, 0, 0, 0, 12,
	13, 0, 0, 0, 0, 0, 20, 0, 0, 21,
	22, 0, 41, 42, 0, 39, 43, 0, 0, 0,
	0, 0, 0, 0, 25, 29, 0, 0, 0, 36,
	0, 6, 37, 0, 35, 0, 0, 26, 27, 28,
	0, 7, 83, 63, 64, 66, 68, 78, 80, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 70, 71,
	72, 73, 74, 0, 0, 75, 76, 60, 61, 62,
	0, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 53, 0, 345, 65, 67, 55, 56, 57,
	58, 59, 0, 0, 0, 0, 81, 344, 52, 0,
	54, 0, 0, 79, 77, 83, 63, 64, 66, 68,
	78, 80, 0, 0, 0, 0, 0, 0, 0, 0,
	69, 70, 71, 72, 73, 74, 0, 0, 75, 76,
	60, 61, 62, 0, 0, 0, 0, 0, 0, 0,
	82, 0, 0, 0, 0, 53, 0, 248, 65, 67,
	55, 56, 57, 58, 59, 0, 0, 0, 0, 81,
	247, 52, 0, 54, 0, 0, 79, 77, 83, 63,
	64, 66, 68, 78, 80, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 70, 71, 72, 73, 74, 0,
	0, 75, 76, 60, 61, 62, 0, 0, 0, 0,
	0, 0, 0, 82, 0, 0, 0, 0, 53, 224,
	0, 65, 67, 55, 56, 57, 58, 59, 0, 0,
	0, 0, 81, 0, 52, 225, 54, 0, 0, 79,
	77, 83, 63, 64, 66, 68, 78, 80, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 70, 71, 72,
	73, 74, 0, 0, 75, 76, 60, 61, 62, 0,
	0, 0, 0, 0, 0, 0, 82, 0, 0, 0,
	0, 53, 201, 0, 65, 67, 55, 56, 57, 58,
	59, 0, 0, 0, 0, 81, 0, 52, 202, 54,
	0, 0, 79, 77, 83, 63, 64, 66, 68, 78,
	80, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	70, 71, 72, 73, 74, 0, 0, 75, 76, 60,
	61, 62, 0, 0, 0, 0, 0, 0, 0, 82,
	0, 0, 0, 0, 53, 0, 0, 65, 67, 55,
	56, 57, 58, 59, 0, 0, 0, 0, 81, 363,
	52, 0, 54, 0, 0, 79, 77, 83, 63, 64,
	66, 68, 78, 80, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 70, 71, 72, 73, 74, 0, 0,
	75, 76, 60, 61, 62, 0, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 53, 0, 0,
	65, 67, 55, 56, 57, 58, 59, 0, 0, 0,
	0, 81, 346, 52, 0, 54, 0, 0, 79, 77,
	83, 63, 64, 66, 68, 78, 80, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 70, 71, 72, 73,
	74, 0, 0, 75, 76, 60, 61, 62, 0, 0,
	0, 0, 0, 0, 0, 82, 0, 0, 0, 0,
	53, 0, 0, 65, 67, 55, 56, 57, 58, 59,
	0, 0, 0, 0, 81, 343, 52, 0, 54, 0,
	0, 79, 77, 83, 63, 64, 66, 68, 78, 80,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 70,
	71, 72, 73, 74, 0, 0, 75, 76, 60, 61,
	62, 0, 0, 0, 0, 0, 0, 0, 82, 0,
	0, 0, 0, 53, 336, 0, 65, 67, 55, 56,
	57, 58, 59, 0, 0, 0, 0, 81, 0, 52,
	0, 54, 0, 0, 79, 77, 83, 63, 64, 66,
	68, 78, 80, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 70, 71, 72, 73, 74, 0, 0, 75,
	76, 60, 61, 62, 0, 0, 0, 0, 0, 0,
	0, 82, 0, 0, 0, 0, 53, 0, 0, 65,
	67, 55, 56, 57, 58, 59, 0, 333, 0, 0,
	81, 0, 52, 0, 54, 0, 0, 79, 77, 83,
	63, 64, 66, 68, 78, 80, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 70, 71, 72, 73, 74,
	0, 0, 75, 76, 60, 61, 62, 0, 0, 0,
	0, 0, 0, 0, 82, 0, 0, 0, 0, 53,
	0, 0, 65, 67, 55, 56, 57, 58, 59, 0,
	0, 0, 0, 81, 330, 52, 0, 54, 0, 0,
	79, 77, 83, 63, 64, 66, 68, 78, 80, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 70, 71,
	72, 73, 74, 0, 0, 75, 76, 60, 61, 62,
	0, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 53, 0, 0, 65, 67, 55, 56, 57,
	58, 59, 0, 0, 0, 0, 81, 0, 52, 319,
	54, 0, 0, 79, 77, 83, 63, 64, 66, 68,
	78, 80, 0, 0, 0, 0, 0, 0, 0, 0,
	69, 70, 71, 72, 73, 74, 0, 0, 75, 76,
	60, 61, 62, 0, 0, 0, 0, 0, 0, 0,
	82, 0, 0, 0, 0, 53, 0, 0, 65, 67,
	55, 56, 57, 58, 59, 0, 0, 0, 0, 81,
	0, 52, 306, 54, 0, 0, 79, 77, 83, 63,
	64, 66, 68, 78, 80, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 70, 71, 72, 73, 74, 0,
	0, 75, 76, 60, 61, 62, 0, 0, 0, 0,
	0, 0, 0, 82, 0, 0, 0, 0, 53, 0,
	0, 65, 67, 55, 56, 57, 58, 59, 0, 0,
	0, 0, 81, 0, 52, 286, 54, 0, 0, 79,
	77, 83, 63, 64, 66, 68, 78, 80, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 70, 71, 72,
	73, 74, 0, 0, 75, 76, 60, 61, 62, 0,
	0, 0, 0, 0, 0, 0, 82, 0, 0, 0,
	0, 53, 0, 0, 65, 67, 55, 56, 57, 58,
	59, 0, 0, 0, 264, 81, 0, 52, 0, 54,
	0, 0, 79, 77, 83, 63, 64, 66, 68, 78,
	80, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	70, 71, 72, 73, 74, 0, 0, 75, 76, 60,
	61, 62, 0, 0, 0, 0, 0, 0, 0, 82,
	0, 0, 0, 0, 53, 0, 0, 65, 67, 55,
	56, 57, 58, 59, 0, 263, 0, 0, 81, 0,
	52, 0, 54, 0, 0, 79, 77, 83, 63, 64,
	66, 68, 78, 80, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 70, 71, 72, 73, 74, 0, 0,
	75, 76, 60, 61, 62, 0, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 53, 0, 0,
	65, 67, 55, 56, 57, 58, 59, 0, 261, 0,
	0, 81, 0, 52, 0, 54, 0, 0, 79, 77,
	83, 63, 64, 66, 68, 78, 80, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 70, 71, 72, 73,
	74, 0, 0, 75, 76, 60, 61, 62, 0, 0,
	0, 0, 0, 0, 0, 82, 0, 0, 0, 0,
	53, 0, 0, 65, 67, 55, 56, 57, 58, 59,
	0, 0, 0, 0, 81, 0, 52, 255, 54, 0,
	0, 79, 77, 83, 63, 64, 66, 68, 78, 80,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 70,
	71, 72, 73, 74, 0, 0, 75, 76, 60, 61,
	62, 0, 0, 0, 0, 0, 0, 0, 82, 0,
	0, 0, 0, 53, 0, 0, 65, 67, 55, 56,
	57, 58, 59, 0, 0, 0, 0, 81, 239, 52,
	0, 54, 0, 0, 79, 77, 83, 63, 64, 66,
	68, 78, 80, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 70, 71, 72, 73, 74, 0, 0, 75,
	76, 60, 61, 62, 0, 0, 0, 0, 0, 0,
	0, 82, 0, 0, 0, 0, 53, 204, 0, 65,
	67, 55, 56, 57, 58, 59, 0, 0, 0, 0,
	81, 0, 52, 0, 54, 0, 0, 79, 77, 83,
	63, 64, 66, 68, 78, 80, 0, 0, 0, 0,
	0, 0, 0, 0, 69, 70, 71, 72, 73, 74,
	0, 0, 75, 76, 60, 61, 62, 0, 0, 0,
	0, 0, 0, 0, 82, 0, 0, 0, 0, 53,
	0, 0, 65, 67, 55, 56, 57, 58, 59, 0,
	198, 0, 0, 81, 0, 52, 0, 54, 0, 0,
	79, 77, 83, 63, 64, 66, 68, 78, 80, 0,
	0, 0, 0, 0, 0, 0, 0, 69, 70, 71,
	72, 73, 74, 0, 0, 75, 76, 60, 61, 62,
	0, 0, 0, 0, 0, 0, 0, 82, 0, 0,
	0, 0, 53, 0, 0, 65, 67, 55, 56, 57,
	58, 59, 0, 0, 0, 0, 81, 188, 52, 0,
	54, 0, 0, 79, 77, 83, 63, 64, 66, 68,
	78, 80, 0, 0, 0, 0, 0, 0, 0, 0,
	69, 70, 71, 72, 73, 74, 0, 0, 75, 76,
	60, 61, 62, 0, 0, 0, 0, 0, 0, 0,
	82, 0, 0, 0, 0, 53, 0, 0, 65, 67,
	55, 56, 57, 58, 59, 0, 172, 0, 0, 81,
	0, 52, 0, 54, 0, 0, 79, 77, 83, 63,
	64, 66, 68, 78, 80, 0, 0, 0, 0, 0,
	0, 0, 0, 69, 70, 71, 72, 73, 74, 0,
	0, 75, 76, 60, 61, 62, 0, 0, 0, 0,
	0, 0, 0, 82, 0, 0, 0, 0, 53, 0,
	0, 65, 67, 55, 56, 57, 58, 59, 0, 169,
	0, 0, 81, 0, 52, 0, 54, 0, 0, 79,
	77, 83, 63, 64, 66, 68, 78, 80, 0, 0,
	0, 0, 0, 0, 0, 0, 69, 70, 71, 72,
	73, 74, 0, 0, 75, 76, 60, 61, 62, 0,
	0, 0, 0, 0, 0, 0, 82, 0, 0, 0,
	51, 53, 0, 0, 65, 67, 55, 56, 57, 58,
	59, 0, 0, 0, 0, 81, 0, 52, 0, 54,
	0, 0, 79, 77, 83, 63, 64, 66, 68, 78,
	80, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	70, 71, 72, 73, 74, 0, 0, 75, 76, 60,
	61, 62, 0, 0, 0, 0, 0, 0, 0, 82,
	0, 0, 0, 0, 53, 0, 0, 65, 67, 55,
	56, 57, 58, 59, 0, 0, 0, 0, 81, 0,
	52, 0, 54, 0, 0, 79, 77, 83, 63, 64,
	66, 68, 78, 80, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 70, 71, 72, 73, 74, 0, 0,
	75, 76, 60, 61, 62, 0, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 53, 0, 0,
	65, 67, 55, 56, 57, 58, 59, 0, 0, 0,
	0, 81, 0, 52, 0, 181, 0, 0, 79, 77,
	83, 63, 64, 66, 68, 78, 80, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 70, 71, 72, 73,
	74, 0, 0, 75, 76, 60, 61, 62, 0, 0,
	0, 0, 0, 0, 0, 82, 0, 0, 0, 0,
	53, 0, 0, 65, 67, 55, 56, 57, 58, 59,
	0, 0, 0, 0, 81, 0, 52, 0, 180, 0,
	0, 79, 77, 83, 63, 64, 66, 68, 78, 80,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 70,
	71, 72, 73, 74, 0, 0, 75, 76, 60, 61,
	62, 0, 0, 0, 0, 0, 0, 0, 82, 0,
	0, 0, 0, 53, 0, 0, 65, 67, 55, 56,
	57, 58, 59, 0, 0, 0, 0, 176, 0, 52,
	0, 54, 0, 0, 79, 77, 83, 63, 64, 66,
	68, 78, 80, 0, 0, 0, 0, 0, 0, 0,
	0, 69, 70, 71, 72, 73, 74, 0, 0, 75,
	76, 60, 61, 62, 0, 0, 0, 0, 0, 0,
	0, 82, 0, 0, 0, 0, 53, 0, 0, 65,
	67, 55, 56, 57, 58, 59, 0, 0, 0, 0,
	174, 0, 52, 0, 54, 0, 0, 79, 77, 23,
	24, 212, 0, 0, 34, 14, 9, 15, 45, 0,
	18, 0, 0, 0, 0, 0, 0, 0, 40, 31,
	32, 33, 16, 19, 0, 0, 0, 0, 0, 0,
	0, 0, 12, 13, 0, 0, 0, 0, 0, 20,
	0, 0, 21, 22, 0, 41, 42, 0, 39, 43,
	0, 0, 0, 0, 0, 0, 0, 25, 29, 0,
	0, 0, 36, 0, 0, 37, 0, 35, 0, 0,
	26, 27, 28, 23, 24, 30, 0, 0, 34, 14,
	9, 15, 45, 0, 18, 0, 0, 0, 0, 0,
	0, 0, 40, 31, 32, 33, 16, 19, 0, 0,
	0, 0, 0, 0, 0, 0, 12, 13, 0, 0,
	0, 0, 0, 20, 0, 0, 21, 22, 0, 41,
	42, 0, 39, 43, 0, 0, 0, 0, 0, 0,
	0, 25, 29, 0, 0, 0, 36, 0, 0, 37,
	0, 35, 0, 0, 26, 27, 28, 83, 63, 64,
	66, 68, 0, 80, 0, 0, 0, 0, 0, 0,
	0, 0, 69, 70, 71, 72, 73, 74, 0, 0,
	75, 76, 60, 61, 62, 0, 0, 0, 0, 0,
	0, 0, 82, 0, 0, 0, 0, 0, 0, 0,
	65, 67, 55, 56, 57, 58, 59, 0, 0, 0,
	0, 81, 0, 52, 0, 54, 0, 0, 79, 77,
	83, 63, 64, 66, 68, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 70, 71, 72, 73,
	74, 0, 0, 75, 76, 60, 61, 62, 0, 0,
	0, 0, 0, 0, 0, 82, 0, 0, 0, 0,
	0, 0, 0, 65, 67, 55, 56, 57, 58, 59,
	0, 66, 68, 0, 81, 0, 52, 0, 54, 0,
	0, 79, 77, 69, 70, 71, 72, 73, 74, 0,
	0, 75, 76, 60, 61, 62, 0, 0, 0, 23,
	24, 30, 0, 82, 34, 0, 0, 0, 0, 0,
	0, 65, 67, 55, 56, 57, 58, 59, 40, 31,
	32, 33, 81, 0, 52, 0, 54, 0, 0, 79,
	77, 23, 24, 30, 0, 0, 34, 0, 0, 0,
	0, 0, 0, 0, 0, 41, 42, 0, 39, 43,
	40, 31, 32, 33, 0, 0, 0, 25, 29, 0,
	0, 0, 36, 0, 0, 37, 0, 35, 285, 0,
	26, 27, 28, 0, 0, 0, 0, 41, 42, 0,
	39, 43, 0, 0, 0, 0, 0, 0, 0, 25,
	29, 0, 0, 0, 36, 0, 0, 37, 0, 35,
	254, 0, 26, 27, 28, 69, 70, 71, 72, 73,
	74, 0, 0, 75, 76, 60, 0, 0, 0, 0,
	0, 23, 24, 30, 0, 82, 34, 0, 0, 0,
	0, 0, 0, 0, 0, 55, 56, 57, 58, 59,
	40, 31, 32, 33, 81, 0, 52, 0, 54, 0,
	0, 79, 77, 0, 0, 0, 23, 24, 30, 0,
	0, 34, 0, 0, 0, 0, 0, 41, 42, 0,
	39, 43, 0, 0, 178, 40, 31, 32, 33, 25,
	29, 0, 0, 0, 36, 0, 0, 37, 0, 35,
	0, 0, 26, 27, 28, 0, 0, 0, 0, 0,
	0, 0, 41, 42, 0, 39, 43, 0, 0, 128,
	0, 23, 24, 30, 25, 29, 34, 0, 0, 36,
	0, 0, 37, 0, 35, 0, 0, 26, 27, 28,
	40, 31, 32, 33, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 259, 24, 30, 0, 0, 34, 0,
	0, 0, 0, 0, 0, 0, 0, 41, 42, 0,
	39, 43, 40, 31, 32, 33, 0, 0, 0, 25,
	29, 0, 0, 0, 36, 0, 0, 37, 0, 35,
	0, 0, 26, 27, 28, 0, 0, 0, 0, 41,
	42, 0, 39, 43, 0, 0, 0, 0, 250, 24,
	30, 25, 29, 34, 0, 0, 36, 0, 0, 37,
	0, 35, 0, 0, 26, 27, 28, 40, 31, 32,
	33, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	109, 24, 30, 0, 0, 34, 0, 0, 0, 0,
	0, 0, 0, 0, 41, 42, 0, 39, 43, 40,
	31, 32, 33, 0, 0, 0, 25, 29, 0, 0,
	0, 36, 0, 0, 37, 0, 35, 0, 0, 26,
	27, 28, 0, 0, 0, 0, 41, 42, 0, 39,
	43, 0, 0, 0, 0, 107, 24, 30, 25, 29,
	34, 0, 0, 36, 0, 0, 37, 0, 35, 0,
	0, 26, 27, 28, 40, 31, 32, 33, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 100, 24, 30,
	0, 0, 34, 0, 0, 0, 0, 0, 0, 0,
	0, 41, 42, 0, 39, 43, 40, 31, 32, 33,
	0, 0, 0, 25, 29, 0, 0, 0, 36, 0,
	0, 37, 0, 35, 0, 0, 26, 27, 28, 0,
	0, 0, 0, 41, 42, 0, 39, 43, 0, 0,
	0, 0, 98, 24, 30, 25, 29, 34, 0, 0,
	36, 0, 0, 37, 0, 35, 0, 0, 26, 27,
	28, 40, 31, 32, 33, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 94, 24, 30, 0, 0, 34,
	0, 0, 0, 0, 0, 0, 0, 0, 41, 42,
	0, 39, 43, 40, 31, 32, 33, 0, 0, 0,
	25, 29, 0, 0, 0, 36, 0, 0, 37, 0,
	35, 0, 0, 26, 27, 28, 0, 0, 0, 0,
	41, 42, 0, 39, 43, 0, 0, 0, 0, 0,
	0, 0, 25, 29, 0, 0, 0, 90, 0, 0,
	37, 0, 35, 0, 0, 26, 27, 28, 69, 70,
	71, 72, 73, 74, 0, 0, 0, 0, 60, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 82, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	57, 58, 59, 0, 0, 0, 0, 81, 0, 52,
	0, 54, 0, 0, 79, 77,
}
var yyPact = [...]int{

	-61, -1000, 2349, -61, -61, -1000, -1000, -1000, -1000, 252,
	1885, 177, -1000, -1000, 2777, 2777, 260, 229, 3070, 149,
	2777, 3038, 2983, -31, -1000, 2777, 2777, 2777, 2951, 2896,
	-1000, -1000, -1000, -1000, 115, -61, -61, 2777, -1000, 60,
	57, 53, 2777, 52, 165, 2777, -1000, 372, -1000, 163,
	-1000, 2777, 2722, 2777, 256, 2777, 2777, 2777, 2777, 2777,
	2777, 2777, 2777, 2777, 2777, 2777, 2777, 2777, 2777, 2777,
	2777, 2777, 2777, 2777, 2777, -1000, -1000, 2777, 2777, 2777,
	2777, 2777, 2777, 2777, 2777, 162, 1948, 1948, 148, 181,
	-61, 201, 72, 1822, -31, 176, -61, 1759, -23, 2200,
	-28, 2137, 2687, 2777, 61, 61, 61, -31, 2074, -31,
	2011, 252, 50, 2777, 238, 1696, 2777, 251, 56, 1948,
	2777, -61, 1633, -1000, 2777, -61, 1948, 625, 2777, 1570,
	-1000, 3117, 3117, 61, 61, 61, 1948, 2644, 2644, 2522,
	2522, 2644, 2644, 2644, 2644, 1948, 1948, 1948, 1948, 1948,
	1948, 1948, 2411, 1948, 2474, 135, 1948, -1000, 2474, -1000,
	1948, -61, -61, 2777, -61, 124, 2275, 2777, 2777, -61,
	2777, 123, -61, 2777, 2777, 2777, 2777, 562, 2777, 129,
	254, 253, 126, 252, 32, -44, -1000, 173, -1000, 1507,
	-38, -1000, 251, 41, 250, -41, 499, 2864, -61, -1000,
	243, 2597, -1000, 1444, 2777, 47, -1000, 2809, 121, 1381,
	114, -1000, 173, 1318, 1255, 113, -1000, 211, 106, 188,
	101, 99, 90, 78, 2565, -1000, 1192, 45, -1000, -1000,
	-1000, 147, 44, 74, -61, -62, -61, 107, 2777, -1000,
	241, -1000, -34, -64, -54, 161, -1000, -1000, 2777, 1948,
	-31, 98, -1000, 1129, -1000, -1000, 1948, -1000, 1948, -31,
	-1000, -61, -1000, -61, 2777, -1000, 142, -1000, -1000, -1000,
	-1000, 2777, 172, -1000, -1000, -1000, 43, -1000, 35, -1000,
	28, -1000, 6, -1000, 1066, -1000, -1000, -1000, -61, 144,
	143, -21, 257, -1000, 120, -1000, 1948, -1000, -1000, 2777,
	-1000, -1000, 2777, 2777, 1003, -1000, -1000, 97, 95, 940,
	140, -61, 877, 171, -61, -1000, -1000, -1000, -1000, -1000,
	94, -61, -61, 134, -1000, -1000, -1000, 814, 436, 751,
	-1000, -1000, -1000, -61, -61, 92, -61, -61, -61, -1000,
	88, 87, -61, -1000, -1000, 2777, -1000, 86, 85, 208,
	-61, -61, -1000, -1000, -1000, 82, 688, -1000, 205, 132,
	-1000, -1000, -1000, -1000, 128, -61, -61, 79, 76, -1000,
	-1000,
}
var yyPgo = [...]int{

	0, 13, 276, 242, 274, 7, 6, 273, 0, 76,
	15, 270, 1, 269, 12, 10, 268, 5, 2, 51,
	267, 241,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 4, 4, 7, 7, 7, 7,
	7, 7, 7, 6, 18, 5, 16, 16, 16, 12,
	13, 13, 13, 14, 14, 14, 15, 15, 11, 10,
	10, 10, 9, 9, 9, 9, 17, 17, 17, 17,
	17, 17, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 19, 19, 20, 20, 21, 21,
}
var yyR2 = [...]int{

	0, 1, 2, 0, 2, 3, 4, 3, 3, 1,
	1, 2, 2, 5, 1, 4, 7, 9, 5, 13,
	12, 9, 8, 5, 6, 5, 6, 5, 6, 5,
	6, 5, 1, 7, 5, 5, 0, 2, 2, 2,
	2, 2, 2, 5, 5, 4, 0, 2, 3, 3,
	0, 1, 4, 0, 1, 4, 1, 3, 3, 1,
	4, 4, 0, 1, 4, 4, 6, 5, 5, 6,
	5, 5, 1, 1, 2, 2, 2, 2, 4, 2,
	4, 1, 1, 1, 1, 5, 3, 7, 8, 8,
	9, 5, 6, 5, 6, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 2, 3, 3,
	3, 3, 5, 4, 5, 4, 4, 4, 1, 4,
	4, 5, 7, 5, 7, 9, 7, 3, 2, 4,
	6, 3, 3, 0, 1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -19, -2, -20, -21, 69, 79, -3, 11,
	-8, -10, 37, 38, 10, 12, 27, -4, 15, 28,
	44, 47, 48, 4, 5, 62, 75, 76, 77, 63,
	6, 24, 25, 26, 9, 72, 67, 70, -17, 53,
	23, 50, 51, 54, -9, 13, -19, -20, -21, -14,
	4, 55, 72, 56, 74, 61, 62, 63, 64, 65,
	41, 42, 43, 17, 18, 59, 19, 60, 20, 31,
	32, 33, 34, 35, 36, 39, 40, 78, 21, 77,
	22, 70, 51, 16, 55, -9, -8, -8, 4, 14,
	67, -14, -11, -8, 4, -10, 67, -8, 4, -8,
	4, -8, 72, 70, -8, -8, -8, 4, -8, 4,
	-8, 70, 4, -19, -19, -8, 70, 70, 70, -8,
	70, 58, -8, -3, 55, 58, -8, -8, 57, -8,
	4, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -9, -8, -17, -8, -10,
	-8, 58, 67, 13, 67, -1, -19, 16, 69, 67,
	55, -1, 67, 70, 70, 70, 70, -8, 57, -9,
	74, 74, -14, 70, -9, -13, -12, 6, 71, -8,
	-15, 4, 49, -16, 52, 72, -8, -19, 67, -10,
	-19, 57, 73, -8, 57, 8, 71, -19, -1, -8,
	-1, 68, 6, -8, -8, -1, -10, 68, -7, -19,
	-9, -9, -9, -9, 57, 73, -8, 8, 71, 4,
	4, 71, 8, -14, 58, -19, 58, -19, 57, 71,
	74, 71, -15, 72, -15, 4, 73, 71, 58, -8,
	4, -1, 4, -8, 73, 73, -8, 71, -8, 4,
	68, 67, 68, 67, 69, 68, 29, 68, -6, -18,
	-5, 45, 46, -6, -5, -18, 8, 71, 8, 71,
	8, 71, 8, 71, -8, 73, 73, 71, 67, 71,
	71, 8, -19, 73, -19, 68, -8, 4, 71, 58,
	73, 71, 58, 58, -8, 68, 73, -1, -1, -8,
	4, 67, -8, -10, 57, 71, 71, 71, 71, 73,
	-1, 67, 67, 71, 73, -12, 68, -8, -8, -8,
	71, 68, 68, 67, 67, -1, 57, 57, -19, 68,
	-1, -1, 67, 71, 71, 58, 71, -1, -1, 68,
	-19, -19, -1, 68, 68, -1, -8, 68, 68, 30,
	-1, -1, 68, 71, 30, 67, 67, -1, -1, 68,
	68,
}
var yyDef = [...]int{

	-2, -2, -2, 143, 144, 145, 147, 148, 4, 53,
	-2, 0, 9, 10, 62, 0, 0, 14, 53, 0,
	0, 0, 0, 72, 73, 0, 0, 0, 0, 0,
	81, 82, 83, 84, 0, 143, 143, 0, 128, 0,
	0, 0, 0, 0, 0, 0, 2, -2, 146, 0,
	54, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 116, 117, 0, 0, 0,
	0, 62, 0, 0, 62, 11, 63, 12, 0, 0,
	-2, 0, 0, -2, -2, 0, -2, 0, 72, 0,
	72, 0, 0, 62, 74, 75, 76, -2, 0, -2,
	0, 53, 0, 62, 50, 0, 0, 0, 46, 138,
	0, 143, 0, 5, 62, 143, 7, 0, 0, 0,
	86, 96, 97, 98, 99, 100, 101, 102, 103, -2,
	-2, 106, 107, 108, 109, 110, 111, 112, 113, 114,
	115, 118, 119, 120, 121, 0, 137, 128, 142, 8,
	-2, 143, -2, 0, -2, 0, -2, 0, 0, -2,
	62, 0, 36, 62, 62, 62, 62, 0, 0, 0,
	0, 0, 0, 53, 143, 143, 51, 0, 95, 0,
	0, 56, 0, 0, 0, 0, 0, 0, -2, 6,
	0, 0, 127, 0, 0, 0, 125, 0, 0, 0,
	0, 15, 81, 0, 0, 0, 58, 0, 0, 0,
	0, 0, 0, 0, 0, 126, 0, 0, 123, 78,
	80, 0, 0, 0, 143, 0, 143, 0, 0, 129,
	0, 130, 0, 0, 0, 0, 47, 139, 0, -2,
	-2, 0, 55, 0, 70, 71, 85, 124, 64, -2,
	13, -2, 34, -2, 0, 18, 0, 23, 39, 41,
	42, 62, 0, 37, 38, 40, 0, -2, 0, -2,
	0, -2, 0, -2, 0, 67, 68, 122, -2, 0,
	0, 0, 0, 91, 0, 93, 49, 57, 131, 0,
	48, 133, 0, 0, 0, 35, 69, 0, 0, 0,
	0, -2, 63, 0, 143, -2, -2, -2, -2, 66,
	0, -2, -2, 0, 92, 52, 94, 0, 0, 0,
	140, 33, 16, -2, -2, 0, 143, 143, -2, 87,
	0, 0, -2, 132, 134, 0, 136, 0, 0, 22,
	-2, -2, 45, 88, 89, 0, 0, 17, 21, 0,
	43, 44, 90, 135, 0, -2, -2, 0, 0, 20,
	19,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	79, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 75, 3, 3, 3, 65, 77, 3,
	70, 71, 63, 61, 58, 62, 74, 64, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 57, 69,
	60, 55, 59, 56, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 72, 3, 73, 76, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 67, 78, 68,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 66,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:73
		{
			yyVAL.compstmt = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:77
		{
			yyVAL.compstmt = yyDollar[1].stmts
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:82
		{
			yyVAL.stmts = nil
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:89
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
			if l, ok := yylex.(*Lexer); ok {
				l.stmts = yyVAL.stmts
			}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:96
		{
			if yyDollar[3].stmt != nil {
				yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
				if l, ok := yylex.(*Lexer); ok {
					l.stmts = yyVAL.stmts
				}
			}
		}
	case 6:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:107
		{
			yyVAL.stmt = &ast.VarStmt{Names: yyDollar[2].expr_idents, Exprs: yyDollar[4].expr_many}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:112
		{
			yyVAL.stmt = &ast.LetsStmt{Lhss: []ast.Expr{yyDollar[1].expr}, Operator: "=", Rhss: []ast.Expr{yyDollar[3].expr}}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:116
		{
			if len(yyDollar[1].expr_many) == 2 && len(yyDollar[3].expr_many) == 1 {
				if _, ok := yyDollar[3].expr_many[0].(*ast.ItemExpr); ok {
					yyVAL.stmt = &ast.LetMapItemStmt{Lhss: yyDollar[1].expr_many, Rhs: yyDollar[3].expr_many[0]}
				} else {
					yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
				}
			} else {
				yyVAL.stmt = &ast.LetsStmt{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
			}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:128
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:133
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:138
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:143
		{
			yyVAL.stmt = &ast.ThrowStmt{Expr: yyDollar[2].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:148
		{
			yyVAL.stmt = &ast.ModuleStmt{Name: yyDollar[2].tok.Lit, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:153
		{
			yyVAL.stmt = yyDollar[1].stmt_if
			yyVAL.stmt.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:158
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:163
		{
			yyVAL.stmt = &ast.ForStmt{Vars: yyDollar[2].expr_idents, Value: yyDollar[4].expr, Stmts: yyDollar[6].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:168
		{
			yyVAL.stmt = &ast.CForStmt{Expr1: yyDollar[2].expr_lets, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:173
		{
			yyVAL.stmt = &ast.LoopStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[4].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 19:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line parser.go.y:178
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt, Finally: yyDollar[12].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 20:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line parser.go.y:183
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt, Finally: yyDollar[11].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 21:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:188
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Var: yyDollar[6].tok.Lit, Catch: yyDollar[8].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:193
		{
			yyVAL.stmt = &ast.TryStmt{Try: yyDollar[3].compstmt, Catch: yyDollar[7].compstmt}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:198
		{
			yyVAL.stmt = &ast.SwitchStmt{Expr: yyDollar[2].expr, Cases: yyDollar[4].stmt_cases}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:203
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].tok.Position())
		}
	case 25:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:208
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].tok.Position())
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:213
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].expr.Position())
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:218
		{
			yyVAL.stmt = &ast.GoroutineStmt{Expr: &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Go: true}}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 28:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:223
		{
			yyVAL.stmt = &ast.DeferStmt{Expr: &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, VarArg: true, Defer: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].tok.Position())
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:228
		{
			yyVAL.stmt = &ast.DeferStmt{Expr: &ast.CallExpr{Name: yyDollar[2].tok.Lit, SubExprs: yyDollar[4].exprs, Defer: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].tok.Position())
		}
	case 30:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:233
		{
			yyVAL.stmt = &ast.DeferStmt{Expr: &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, VarArg: true, Defer: true}}
			yyVAL.stmt.SetPosition(yyDollar[2].expr.Position())
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:238
		{
			yyVAL.stmt = &ast.DeferStmt{Expr: &ast.AnonCallExpr{Expr: yyDollar[2].expr, SubExprs: yyDollar[4].exprs, Defer: true}}
			yyVAL.stmt.SetPosition(yyDollar[1].tok.Position())
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:243
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
			yyVAL.stmt.SetPosition(yyDollar[1].expr.Position())
		}
	case 33:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:251
		{
			yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf = append(yyDollar[1].stmt_if.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].compstmt})
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:256
		{
			if yyVAL.stmt_if.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyVAL.stmt_if.(*ast.IfStmt).Else = append(yyVAL.stmt_if.(*ast.IfStmt).Else, yyDollar[4].compstmt...)
			}
			yyVAL.stmt_if.SetPosition(yyDollar[1].stmt_if.Position())
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:265
		{
			yyVAL.stmt_if = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].compstmt, Else: nil}
			yyVAL.stmt_if.SetPosition(yyDollar[1].tok.Position())
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:271
		{
			yyVAL.stmt_cases = []ast.Stmt{}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:275
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_case}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:279
		{
			yyVAL.stmt_cases = []ast.Stmt{yyDollar[2].stmt_default}
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:283
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_case)
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:287
		{
			yyVAL.stmt_cases = yyDollar[2].stmt_multi_case
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:291
		{
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_multi_case...)
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:295
		{
			for _, stmt := range yyDollar[1].stmt_cases {
				if _, ok := stmt.(*ast.DefaultStmt); ok {
					yylex.Error("multiple default statement")
				}
			}
			yyVAL.stmt_cases = append(yyDollar[1].stmt_cases, yyDollar[2].stmt_default)
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:306
		{
			yyVAL.stmt_case = &ast.CaseStmt{Expr: yyDollar[2].expr, Stmts: yyDollar[5].compstmt}
		}
	case 44:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:312
		{
			var cases = []ast.Stmt{}
			for _, stmt := range yyDollar[2].expr_many {
				cases = append(cases, &ast.CaseStmt{Expr: stmt, Stmts: yyDollar[5].compstmt})
			}
			yyVAL.stmt_multi_case = cases
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:322
		{
			yyVAL.stmt_default = &ast.DefaultStmt{Stmts: yyDollar[4].compstmt}
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:327
		{
			yyVAL.array_count = ast.ArrayCount{Count: 0}
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:331
		{
			yyVAL.array_count = ast.ArrayCount{Count: 1}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:335
		{
			yyVAL.array_count.Count = yyVAL.array_count.Count + 1
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:341
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].tok.Lit, Value: yyDollar[3].expr}
		}
	case 50:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:346
		{
			yyVAL.expr_pairs = []ast.Expr{}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:350
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:354
		{
			if len(yyDollar[1].expr_pairs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:362
		{
			yyVAL.expr_idents = []string{}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:366
		{
			yyVAL.expr_idents = []string{yyDollar[1].tok.Lit}
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:370
		{
			if len(yyDollar[1].expr_idents) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.expr_idents = append(yyDollar[1].expr_idents, yyDollar[4].tok.Lit)
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:379
		{
			yyVAL.expr_type = yyDollar[1].tok.Lit
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:383
		{
			yyVAL.expr_type = yyVAL.expr_type + "." + yyDollar[3].tok.Lit
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:388
		{
			yyVAL.expr_lets = &ast.LetsExpr{Lhss: yyDollar[1].expr_many, Operator: "=", Rhss: yyDollar[3].expr_many}
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:394
		{
			yyVAL.expr_many = []ast.Expr{yyDollar[1].expr}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:398
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:402
		{
			yyVAL.expr_many = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 62:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:407
		{
			yyVAL.exprs = nil
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:411
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:415
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:422
		{
			if len(yyDollar[1].exprs) == 0 {
				yylex.Error("syntax error: unexpected ','")
			}
			yyVAL.exprs = append(yyDollar[1].exprs, &ast.IdentExpr{Lit: yyDollar[4].tok.Lit})
		}
	case 66:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:431
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
		}
	case 67:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:435
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: yyDollar[3].expr, End: nil}
		}
	case 68:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:439
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Begin: nil, End: yyDollar[4].expr}
		}
	case 69:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:443
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
		}
	case 70:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:447
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:451
		{
			yyVAL.expr_slice = ast.SliceExpr{Value: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:457
		{
			yyVAL.expr = &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:462
		{
			yyVAL.expr = &ast.NumberExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:467
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:472
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:477
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "^", Expr: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:482
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 78:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:487
		{
			yyVAL.expr = &ast.AddrExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:492
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.IdentExpr{Lit: yyDollar[2].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].tok.Position())
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:497
		{
			yyVAL.expr = &ast.DerefExpr{Expr: &ast.MemberExpr{Expr: yyDollar[2].expr, Name: yyDollar[4].tok.Lit}}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:502
		{
			yyVAL.expr = &ast.StringExpr{Lit: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:507
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:512
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:517
		{
			yyVAL.expr = &ast.ConstExpr{Value: yyDollar[1].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 85:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:522
		{
			yyVAL.expr = &ast.TernaryOpExpr{Expr: yyDollar[1].expr, Lhs: yyDollar[3].expr, Rhs: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:527
		{
			yyVAL.expr = &ast.MemberExpr{Expr: yyDollar[1].expr, Name: yyDollar[3].tok.Lit}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 87:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:532
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmts: yyDollar[6].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 88:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:537
		{
			yyVAL.expr = &ast.FuncExpr{Params: yyDollar[3].expr_idents, Stmts: yyDollar[7].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 89:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:542
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmts: yyDollar[7].compstmt}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 90:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:547
		{
			yyVAL.expr = &ast.FuncExpr{Name: yyDollar[2].tok.Lit, Params: yyDollar[4].expr_idents, Stmts: yyDollar[8].compstmt, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 91:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:552
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 92:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:557
		{
			yyVAL.expr = &ast.ArrayExpr{Exprs: yyDollar[3].exprs}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 93:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:562
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyDollar[3].expr_pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 94:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:571
		{
			mapExpr := make(map[string]ast.Expr)
			for _, v := range yyDollar[3].expr_pairs {
				mapExpr[v.(*ast.PairExpr).Key] = v.(*ast.PairExpr).Value
			}
			yyVAL.expr = &ast.MapExpr{MapExpr: mapExpr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:580
		{
			yyVAL.expr = &ast.ParenExpr{SubExpr: yyDollar[2].expr}
			if l, ok := yylex.(*Lexer); ok {
				yyVAL.expr.SetPosition(l.pos)
			}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:585
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "+", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:590
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "-", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:595
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "*", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:600
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "/", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:605
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "%", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:610
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "**", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:615
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:620
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">>", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:625
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "==", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:630
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "!=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:635
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:640
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: ">=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:645
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:650
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "<=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:655
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "+=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:660
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "-=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:665
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "*=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:670
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "/=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:675
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "&=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:680
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "|=", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:685
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "++"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:690
		{
			yyVAL.expr = &ast.AssocExpr{Lhs: yyDollar[1].expr, Operator: "--"}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:695
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "|", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:700
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "||", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:705
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:710
		{
			yyVAL.expr = &ast.BinOpExpr{Lhs: yyDollar[1].expr, Operator: "&&", Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 122:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:715
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 123:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:720
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].tok.Lit, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 124:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:725
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs, VarArg: true}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 125:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:730
		{
			yyVAL.expr = &ast.AnonCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 126:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:735
		{
			yyVAL.expr = &ast.ItemExpr{Value: &ast.IdentExpr{Lit: yyDollar[1].tok.Lit}, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 127:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:740
		{
			yyVAL.expr = &ast.ItemExpr{Value: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:745
		{
			yyVAL.expr = &yyDollar[1].expr_slice
			yyVAL.expr.SetPosition(yyDollar[1].expr_slice.Position())
		}
	case 129:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:750
		{
			yyVAL.expr = &ast.LenExpr{Expr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 130:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:755
		{
			yyVAL.expr = &ast.NewExpr{Type: yyDollar[3].expr_type}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 131:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:760
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr_type, SizeExpr: nil}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 132:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:765
		{
			yyVAL.expr = &ast.MakeChanExpr{Type: yyDollar[4].expr_type, SizeExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 133:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:770
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 134:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:775
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type, LenExpr: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 135:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:780
		{
			yyVAL.expr = &ast.MakeExpr{Dimensions: yyDollar[3].array_count.Count, Type: yyDollar[4].expr_type, LenExpr: yyDollar[6].expr, CapExpr: yyDollar[8].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 136:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:785
		{
			yyVAL.expr = &ast.MakeTypeExpr{Name: yyDollar[4].tok.Lit, Type: yyDollar[6].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:790
		{
			yyVAL.expr = &ast.ChanExpr{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 138:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:795
		{
			yyVAL.expr = &ast.ChanExpr{Rhs: yyDollar[2].expr}
			yyVAL.expr.SetPosition(yyDollar[2].expr.Position())
		}
	case 139:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:800
		{
			yyVAL.expr = &ast.DeleteExpr{WhatExpr: yyDollar[3].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 140:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:805
		{
			yyVAL.expr = &ast.DeleteExpr{WhatExpr: yyDollar[3].expr, KeyExpr: yyDollar[5].expr}
			yyVAL.expr.SetPosition(yyDollar[1].tok.Position())
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:810
		{
			yyVAL.expr = &ast.IncludeExpr{ItemExpr: yyDollar[1].expr, ListExpr: yyDollar[3].expr_slice}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:815
		{
			yyVAL.expr = &ast.IncludeExpr{ItemExpr: yyDollar[1].expr, ListExpr: ast.SliceExpr{Value: yyDollar[3].expr, Begin: nil, End: nil}}
			yyVAL.expr.SetPosition(yyDollar[1].expr.Position())
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:828
		{
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:831
		{
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:836
		{
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:839
		{
		}
	}
	goto yystack /* stack new state and value */
}
