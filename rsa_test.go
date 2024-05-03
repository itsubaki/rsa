package rsa_test

import (
	"fmt"

	"github.com/itsubaki/rsa"
	"github.com/itsubaki/rsa/number"
)

func Example() {
	seed := [32]byte([]byte("123456789_123456789_123456789_12"))
	p, q := 17, 19

	e := number.Euler(p, q)
	E := rsa.PublicKey(e, seed)
	D := rsa.PrivateKey(e, E)
	N := p * q
	fmt.Printf("p=%v, q=%v, e=%v, E=%v, D=%v, N=%v\n", p, q, e, E, D, N)

	for msg := 2; msg < N-1; msg++ {
		enc := rsa.Encrypt(msg, E, N)
		dec := rsa.Decrypt(enc, D, N)
		fmt.Printf("msg=%3d, enc=%3d, dec=%3d\n", msg, enc, dec)
	}

	// Output:
	// p=17, q=19, e=288, E=35, D=107, N=323
	// msg=  2, enc=314, dec=  2
	// msg=  3, enc=146, dec=  3
	// msg=  4, enc= 81, dec=  4
	// msg=  5, enc= 23, dec=  5
	// msg=  6, enc=301, dec=  6
	// msg=  7, enc=258, dec=  7
	// msg=  8, enc=240, dec=  8
	// msg=  9, enc=321, dec=  9
	// msg= 10, enc=116, dec= 10
	// msg= 11, enc=311, dec= 11
	// msg= 12, enc=198, dec= 12
	// msg= 13, enc=174, dec= 13
	// msg= 14, enc=262, dec= 14
	// msg= 15, enc=128, dec= 15
	// msg= 16, enc=101, dec= 16
	// msg= 17, enc= 85, dec= 17
	// msg= 18, enc= 18, dec= 18
	// msg= 19, enc= 76, dec= 19
	// msg= 20, enc=248, dec= 20
	// msg= 21, enc=200, dec= 21
	// msg= 22, enc=108, dec= 22
	// msg= 23, enc=233, dec= 23
	// msg= 24, enc=156, dec= 24
	// msg= 25, enc=206, dec= 25
	// msg= 26, enc= 49, dec= 26
	// msg= 27, enc= 31, dec= 27
	// msg= 28, enc=226, dec= 28
	// msg= 29, enc=249, dec= 29
	// msg= 30, enc=140, dec= 30
	// msg= 31, enc=160, dec= 31
	// msg= 32, enc= 60, dec= 32
	// msg= 33, enc=186, dec= 33
	// msg= 34, enc=204, dec= 34
	// msg= 35, enc=120, dec= 35
	// msg= 36, enc=161, dec= 36
	// msg= 37, enc=265, dec= 37
	// msg= 38, enc=285, dec= 38
	// msg= 39, enc=210, dec= 39
	// msg= 40, enc= 29, dec= 40
	// msg= 41, enc=241, dec= 41
	// msg= 42, enc=138, dec= 42
	// msg= 43, enc=270, dec= 43
	// msg= 44, enc=320, dec= 44
	// msg= 45, enc=277, dec= 45
	// msg= 46, enc=164, dec= 46
	// msg= 47, enc= 55, dec= 47
	// msg= 48, enc=211, dec= 48
	// msg= 49, enc= 26, dec= 49
	// msg= 50, enc= 84, dec= 50
	// msg= 51, enc=136, dec= 51
	// msg= 52, enc=205, dec= 52
	// msg= 53, enc=280, dec= 53
	// msg= 54, enc= 44, dec= 54
	// msg= 55, enc= 47, dec= 55
	// msg= 56, enc=227, dec= 56
	// msg= 57, enc=114, dec= 57
	// msg= 58, enc= 20, dec= 58
	// msg= 59, enc=257, dec= 59
	// msg= 60, enc= 32, dec= 60
	// msg= 61, enc=252, dec= 61
	// msg= 62, enc=175, dec= 62
	// msg= 63, enc=130, dec= 63
	// msg= 64, enc=106, dec= 64
	// msg= 65, enc=126, dec= 65
	// msg= 66, enc=264, dec= 66
	// msg= 67, enc=135, dec= 67
	// msg= 68, enc=102, dec= 68
	// msg= 69, enc=103, dec= 69
	// msg= 70, enc=212, dec= 70
	// msg= 71, enc=129, dec= 71
	// msg= 72, enc=166, dec= 72
	// msg= 73, enc=  6, dec= 73
	// msg= 74, enc=199, dec= 74
	// msg= 75, enc= 37, dec= 75
	// msg= 76, enc= 19, dec= 76
	// msg= 77, enc=134, dec= 77
	// msg= 78, enc= 48, dec= 78
	// msg= 79, enc=260, dec= 79
	// msg= 80, enc= 62, dec= 80
	// msg= 81, enc=  4, dec= 81
	// msg= 82, enc= 92, dec= 82
	// msg= 83, enc=315, dec= 83
	// msg= 84, enc= 50, dec= 84
	// msg= 85, enc= 17, dec= 85
	// msg= 86, enc=154, dec= 86
	// msg= 87, enc=178, dec= 87
	// msg= 88, enc= 27, dec= 88
	// msg= 89, enc= 98, dec= 89
	// msg= 90, enc= 91, dec= 90
	// msg= 91, enc=318, dec= 91
	// msg= 92, enc=139, dec= 92
	// msg= 93, enc=104, dec= 93
	// msg= 94, enc=151, dec= 94
	// msg= 95, enc=133, dec= 95
	// msg= 96, enc= 39, dec= 96
	// msg= 97, enc=181, dec= 97
	// msg= 98, enc= 89, dec= 98
	// msg= 99, enc= 24, dec= 99
	// msg=100, enc=213, dec=100
	// msg=101, enc= 16, dec=101
	// msg=102, enc= 68, dec=102
	// msg=103, enc= 69, dec=103
	// msg=104, enc= 93, dec=104
	// msg=105, enc= 78, dec=105
	// msg=106, enc= 64, dec=106
	// msg=107, enc=312, dec=107
	// msg=108, enc=250, dec=108
	// msg=109, enc=224, dec=109
	// msg=110, enc=223, dec=110
	// msg=111, enc=253, dec=111
	// msg=112, enc=218, dec=112
	// msg=113, enc= 56, dec=113
	// msg=114, enc=266, dec=114
	// msg=115, enc=191, dec=115
	// msg=116, enc=143, dec=116
	// msg=117, enc=298, dec=117
	// msg=118, enc=271, dec=118
	// msg=119, enc=289, dec=119
	// msg=120, enc= 35, dec=120
	// msg=121, enc=144, dec=121
	// msg=122, enc=316, dec=122
	// msg=123, enc=302, dec=123
	// msg=124, enc= 40, dec=124
	// msg=125, enc=216, dec=125
	// msg=126, enc=122, dec=126
	// msg=127, enc=155, dec=127
	// msg=128, enc= 15, dec=128
	// msg=129, enc= 14, dec=129
	// msg=130, enc=158, dec=130
	// msg=131, enc= 28, dec=131
	// msg=132, enc=208, dec=132
	// msg=133, enc=228, dec=133
	// msg=134, enc= 77, dec=134
	// msg=135, enc= 67, dec=135
	// msg=136, enc= 51, dec=136
	// msg=137, enc=290, dec=137
	// msg=138, enc= 42, dec=138
	// msg=139, enc=282, dec=139
	// msg=140, enc= 30, dec=140
	// msg=141, enc=278, dec=141
	// msg=142, enc=131, dec=142
	// msg=143, enc=173, dec=143
	// msg=144, enc=121, dec=144
	// msg=145, enc=236, dec=145
	// msg=146, enc=269, dec=146
	// msg=147, enc=243, dec=147
	// msg=148, enc=147, dec=148
	// msg=149, enc=310, dec=149
	// msg=150, enc=313, dec=150
	// msg=151, enc= 94, dec=151
	// msg=152, enc=152, dec=152
	// msg=153, enc=153, dec=153
	// msg=154, enc= 86, dec=154
	// msg=155, enc=127, dec=155
	// msg=156, enc=214, dec=156
	// msg=157, enc=251, dec=157
	// msg=158, enc=244, dec=158
	// msg=159, enc=182, dec=159
	// msg=160, enc= 88, dec=160
	// msg=161, enc= 36, dec=161
	// msg=162, enc=287, dec=162
	// msg=163, enc=235, dec=163
	// msg=164, enc=141, dec=164
	// msg=165, enc= 79, dec=165
	// msg=166, enc= 72, dec=166
	// msg=167, enc=109, dec=167
	// msg=168, enc=196, dec=168
	// msg=169, enc=237, dec=169
	// msg=170, enc=170, dec=170
	// msg=171, enc=171, dec=171
	// msg=172, enc=229, dec=172
	// msg=173, enc= 10, dec=173
	// msg=174, enc= 13, dec=174
	// msg=175, enc=176, dec=175
	// msg=176, enc= 80, dec=176
	// msg=177, enc= 54, dec=177
	// msg=178, enc= 87, dec=178
	// msg=179, enc=202, dec=179
	// msg=180, enc=150, dec=180
	// msg=181, enc=192, dec=181
	// msg=182, enc= 45, dec=182
	// msg=183, enc=293, dec=183
	// msg=184, enc= 41, dec=184
	// msg=185, enc=281, dec=185
	// msg=186, enc= 33, dec=186
	// msg=187, enc=272, dec=187
	// msg=188, enc=256, dec=188
	// msg=189, enc=246, dec=189
	// msg=190, enc= 95, dec=190
	// msg=191, enc=115, dec=191
	// msg=192, enc=295, dec=192
	// msg=193, enc=165, dec=193
	// msg=194, enc=309, dec=194
	// msg=195, enc=308, dec=195
	// msg=196, enc=168, dec=196
	// msg=197, enc=201, dec=197
	// msg=198, enc=107, dec=198
	// msg=199, enc=283, dec=199
	// msg=200, enc= 21, dec=200
	// msg=201, enc=  7, dec=201
	// msg=202, enc=179, dec=202
	// msg=203, enc=288, dec=203
	// msg=204, enc= 34, dec=204
	// msg=205, enc= 52, dec=205
	// msg=206, enc= 25, dec=206
	// msg=207, enc=180, dec=207
	// msg=208, enc=132, dec=208
	// msg=209, enc= 57, dec=209
	// msg=210, enc=267, dec=210
	// msg=211, enc=105, dec=211
	// msg=212, enc= 70, dec=212
	// msg=213, enc=100, dec=213
	// msg=214, enc= 99, dec=214
	// msg=215, enc= 73, dec=215
	// msg=216, enc= 11, dec=216
	// msg=217, enc=259, dec=217
	// msg=218, enc=245, dec=218
	// msg=219, enc=230, dec=219
	// msg=220, enc=254, dec=220
	// msg=221, enc=255, dec=221
	// msg=222, enc=307, dec=222
	// msg=223, enc=110, dec=223
	// msg=224, enc=299, dec=224
	// msg=225, enc=234, dec=225
	// msg=226, enc=142, dec=226
	// msg=227, enc=284, dec=227
	// msg=228, enc=190, dec=228
	// msg=229, enc=172, dec=229
	// msg=230, enc=219, dec=230
	// msg=231, enc=184, dec=231
	// msg=232, enc=  5, dec=232
	// msg=233, enc=232, dec=233
	// msg=234, enc=225, dec=234
	// msg=235, enc=296, dec=235
	// msg=236, enc=145, dec=236
	// msg=237, enc=169, dec=237
	// msg=238, enc=306, dec=238
	// msg=239, enc=273, dec=239
	// msg=240, enc=  8, dec=240
	// msg=241, enc=231, dec=241
	// msg=242, enc=319, dec=242
	// msg=243, enc=261, dec=243
	// msg=244, enc= 63, dec=244
	// msg=245, enc=275, dec=245
	// msg=246, enc=189, dec=246
	// msg=247, enc=304, dec=247
	// msg=248, enc=286, dec=248
	// msg=249, enc=124, dec=249
	// msg=250, enc=317, dec=250
	// msg=251, enc=157, dec=251
	// msg=252, enc=194, dec=252
	// msg=253, enc=111, dec=253
	// msg=254, enc=220, dec=254
	// msg=255, enc=221, dec=255
	// msg=256, enc=188, dec=256
	// msg=257, enc= 59, dec=257
	// msg=258, enc=197, dec=258
	// msg=259, enc=217, dec=259
	// msg=260, enc=193, dec=260
	// msg=261, enc=148, dec=261
	// msg=262, enc= 71, dec=262
	// msg=263, enc=291, dec=263
	// msg=264, enc= 66, dec=264
	// msg=265, enc=303, dec=265
	// msg=266, enc=209, dec=266
	// msg=267, enc= 96, dec=267
	// msg=268, enc=276, dec=268
	// msg=269, enc=279, dec=269
	// msg=270, enc= 43, dec=270
	// msg=271, enc=118, dec=271
	// msg=272, enc=187, dec=272
	// msg=273, enc=239, dec=273
	// msg=274, enc=297, dec=274
	// msg=275, enc=112, dec=275
	// msg=276, enc=268, dec=276
	// msg=277, enc=159, dec=277
	// msg=278, enc= 46, dec=278
	// msg=279, enc=  3, dec=279
	// msg=280, enc= 53, dec=280
	// msg=281, enc=185, dec=281
	// msg=282, enc= 82, dec=282
	// msg=283, enc=294, dec=283
	// msg=284, enc=113, dec=284
	// msg=285, enc= 38, dec=285
	// msg=286, enc= 58, dec=286
	// msg=287, enc=162, dec=287
	// msg=288, enc=203, dec=288
	// msg=289, enc=119, dec=289
	// msg=290, enc=137, dec=290
	// msg=291, enc=263, dec=291
	// msg=292, enc=163, dec=292
	// msg=293, enc=183, dec=293
	// msg=294, enc= 74, dec=294
	// msg=295, enc= 97, dec=295
	// msg=296, enc=292, dec=296
	// msg=297, enc=274, dec=297
	// msg=298, enc=117, dec=298
	// msg=299, enc=167, dec=299
	// msg=300, enc= 90, dec=300
	// msg=301, enc=215, dec=301
	// msg=302, enc=123, dec=302
	// msg=303, enc= 75, dec=303
	// msg=304, enc=247, dec=304
	// msg=305, enc=305, dec=305
	// msg=306, enc=238, dec=306
	// msg=307, enc=222, dec=307
	// msg=308, enc=195, dec=308
	// msg=309, enc= 61, dec=309
	// msg=310, enc=149, dec=310
	// msg=311, enc=125, dec=311
	// msg=312, enc= 12, dec=312
	// msg=313, enc=207, dec=313
	// msg=314, enc=  2, dec=314
	// msg=315, enc= 83, dec=315
	// msg=316, enc= 65, dec=316
	// msg=317, enc= 22, dec=317
	// msg=318, enc=300, dec=318
	// msg=319, enc=242, dec=319
	// msg=320, enc=177, dec=320
	// msg=321, enc=  9, dec=321
}
