package htpasswd

import (
	"fmt"
	"testing"
)

type md5Datum struct {
	password string
	salt     string
	hashed   string
}

var md5TestData = []md5Datum{
	md5Datum{"mickey5", "gxNb79DX", "6wi9QaGNM5TA0kBKiC4710"},
	md5Datum{"alexandrew", "kv1uUfCO", "iEwrWojf92uZ/9uhTQmMo."},
	md5Datum{"hawaiicats78", "UQ6GxE7V", "OrIqWONGuSV9RfS3B2dfO1"},
	md5Datum{"DIENOW", "OZ.RwYJH", "AwfW2h0gJnu2fQi0GegVe1"},
	md5Datum{"e8f685", "9r9GyMpL", "3IiaLNos/tbouLJwsW8ey/"},
	md5Datum{"Rickygirl03", "0tlsxL/0", "cfS6c2JZjwISRTgFvrMWL1"},
	md5Datum{"123vb123", "/4XFfQuK", "bnMIHM0j/Cf8apmbvPzn/."},
	md5Datum{"sheng060576", "NEJJUzVT", "o/CWI9InAMXWAsbl5gx0p1"},
	md5Datum{"hansisme", "JAOXCriK", "gB/Yox3wTae3NujwKUiFv1"},
	md5Datum{"h4ck3rs311t3", "KmkPgS2r", "5qIFMPNVAXzlevkzOQwhj."},
	md5Datum{"K90JyTGA", "mM7q5ZHN", "03LeGh9D1CujEBwiVRO6B0"},
	md5Datum{"aspire5101", "tlxr3zoa", "dQJiJmk4pEtRTssYiLwlv0"},
	md5Datum{"553568", "YI.r2X/w", "H/1DtcmTHSgcdkgz8NS1W0"},
	md5Datum{"SRI", "StJ5t4wb", "tIVEx.MPZR1SqDm5y9VCs1"},
	md5Datum{"maxmus", "ad29tH08", "xEHwr706Yz/3FFGqnVB6l/"},
	md5Datum{"a5xp9707", "aH0sN4io", "y0heNz5hL67/HA7/7mDRS."},
	md5Datum{"tomasrim", "SgbYnJV9", "7Z.enu6vZ7b6Zo7/lYce60"},
	md5Datum{"2a0mag", "lSOzbc7i", "Ae21yFmdTMpSz.aQsjyoE1"},
	md5Datum{"wmsfht", "yicl6/5x", "p/dCDdQ0q9lLaZbBJsIDP0"},
	md5Datum{"webmaster2364288", "PLoY5sMf", "KEDmvJskiSNFwiygtWXin1"},
	md5Datum{"121516m", "3T5gmyrq", "AucgLmXU53aTQJuRKCFo50"},
	md5Datum{"T69228803", "Aajhupso", "/EPFyux8bd7Iw.tLevaVE."},
	md5Datum{"qq820221", "G43B4jFl", "4TUFaOD7Fz5.lZiq5v8P40"},
	md5Datum{"chenfy", "mDnux.Mf", "vXsdihwaTLCJTHnuk9/cK/"},
	md5Datum{"www.debure.net", "bZzoRW4K", "DfI3Col55.57HP3FW4L1h."},
	md5Datum{"1333e763", "rRvCcrzo", "plG5/rpEPSM7uc3bro6P51"},
	md5Datum{"burberries", "Qx6JtYcz", "10t2dI6u0LyNBjeCAQ.3z1"},
	md5Datum{"chanmee14", "p9t9dUC1", "Nlr96oZWIe/VVpYBUgG6q0"},
	md5Datum{"65432106543210", "CBG7TUqG", "Olyygy0L6HfSPfkLg24U60"},
	md5Datum{"powernet", "ogVPlakG", ".SLiqbN/KUECQ6pgdck2/."},
	md5Datum{"a2d8i6a7", "sNrtmvPF", "rvbRuKdcPPvN.dK.mHeYq/"},
	md5Datum{"gvs9ptc", "gQgMMxVG", "5sI4ezBQxqpfh14AvEEVU0"},
	md5Datum{"Pookie", "x.wVLgoG", "HTj5gT.lQ71BpifSlcQVy1"},
	md5Datum{"lorissss", "O0ySiIf4", "AmmYBbHWjfiVcGEbl4wiy/"},
	md5Datum{"ess", "nE19zEmy", "Rg3/wMTNMVOkbhez/QhD//"},
	md5Datum{"sparra", "By1OjZuF", "PRY4G6D8u3aFhruSTgIQC."},
	md5Datum{"allysson", "mI6fsU64", "WqCg/f9CpYr4586AVr6nP."},
	md5Datum{"99128008", "LQLXA.du", "kazspxn165TFSiDavu75N/"},
	md5Datum{"evisanne", "AhDCR8bW", "2lR137DLMfr1mQ9xLlMsw0"},
	md5Datum{"qfxg7x9l", "ZAGBUFGw", "HI3fWMR0Y6Z4U3MSc70sd."},
	md5Datum{"03415", "nkFIBpLJ", "AvABMUIgvoMp0zmOTCwCG1"},
	md5Datum{"87832309", "WbCq7Hv8", "dxe0LoM3vlD.t/A/3Cfd11"},
	md5Datum{"816283", "PrEjUTNt", "vGLTgLqJp9XEtwEJBv5XF."},
	md5Datum{"banach12", "S1G5jLiH", "CySeS1zgVlMLLElxG6Dmw0"},
	md5Datum{"sjdszpsc", "QmuQrgcB", "xZk5zcK2QRF8PZ24P9vPr1"},
	md5Datum{"changsing", "Z0i29yA5", "KnTYiWEZQYzQlH/SxQ7Qp/"},
	md5Datum{"56339388", "RZlCHiTm", "8mFKCLkRHxoJ2ieVa.K79/"},
	md5Datum{"52114157", "3NkMs.IK", "02HiBvqlIVA.hLbktlHsD1"},
	md5Datum{"jinebimb", "1ww3avga", "haxtp7TGUm9PHPBrBeM9u."},
	md5Datum{"erol43", ".aE1EJya", "3zkhvRyNbF.DOOyJSPSJ21"},
	md5Datum{"2yagos", "L0YlhvFW", "R0J.Bk9wYb7sQKXBbP4AN/"},
	md5Datum{"habparty!", "vveX0m/D", "hPoF3j.Ac5zSOAmHBZklT."},
	md5Datum{"tangjianhui", "8Ivzj66d", "J4A.NOn6TRk4RYC9oGqIB/"},
	md5Datum{"serandah", "v9AJex0e", "qn/isKH9e6EG66KCtFdmI1"},
	md5Datum{"mirrages", "UM0E3yNn", "4V4IJI2Q0Bqh0EG8HAHbq0"},
	md5Datum{"mantgaxxl", "1spakyg4", "NwPcxatLI7bWUpeDzAw2h1"},
	md5Datum{"45738901", "oepJpf/s", "p0F.JGVJCyvUHfWnpF.Wy1"},
	md5Datum{"g523minna", "yWpavB.B", "q4KExAyIKMKWTLq86n0820"},
	md5Datum{"j202020", "DTRNSWt7", "2At.lEmBM2waU9F2QsDvd."},
	md5Datum{"g@mmaecho", "QGA07jk6", "U9Uw/dD666GNV60hX6AKM/"},
	md5Datum{"042380", "FDnW17iI", "6jkNwkfAi.4LMYkIkNO2v1"},
	md5Datum{"ASRuin", "GKFI0Se3", "go4Tko/O9UCA2WtSJBjgc."},
	md5Datum{"061990", "yJR0EnuF", "CzzsiUo2Q5cRhtlptUf7D1"},
	md5Datum{"ysoline", "7D0hCvVq", "HLIRmi013HBi2TgATkgJM."},
	md5Datum{"liuzhouzhou", "m.MSvKt4", "oFYUki/pESjwOfF5YH9VO0"},
	md5Datum{"b0000000wind", "qOQrkTXw", "PJXv2X.0Efe4VUPcvyxA61"},
	md5Datum{"7913456852", "lPKDpKzC", "q9kt0R9.I4rxhlIcNe2gg1"},
	md5Datum{"9008", "PYsksC92", "3oqtOxrMnQc1n3GfSIAJM."},
	md5Datum{"waitlin11", "x5UDLNO2", "yHLWIm/50ORtDhT56f9bi0"},
	md5Datum{"8fdakar", "E9a2XIvt", "fcsw4gZfbiDXPywMzwhik1"},
	md5Datum{"eisball", "gHg16GuT", "DGI/O8HzZemhsQ4o2jA560"},
	md5Datum{"jenna17", "yzwqt8mS", "3QqqiFB9Z6q1fp4z/q1pU."},
	md5Datum{"belkadonam", "iGU4vuaZ", "w3xf5rVAIJYz0dgImL8a2."},
	md5Datum{"tfyuj9JW", "5cPUmio7", "wttScNV7Fk4Njs9QX1yUi."},
	md5Datum{"nihaijidema", "DZW4Gt4h", "EXlVFPbqnXGPp2vLQT5TK0"},
	md5Datum{"talapia", "61i3ruRm", "cNcNvti2hQ8mXjLahFnSb/"},
	md5Datum{"7376220", "Z89Ynh0K", "A2k6aLQnMOa2uwXX8MJZf1"},
	md5Datum{"c7m8e1xsc3", "QRn4AsCM", "gUztH0RWKuX1Vy0WaYfdC1"},
	md5Datum{"84129793", "rghudgt5", "XA7QLtfRq84JHtbjdke0I."},
	md5Datum{"test1000", "zwkIVA3j", "Iuz7zNyLvIiKWIl2VA8bl."},
	md5Datum{"ecmanhatten", "zfVlWDS.", "emJhRC3N0SnvZLo5en4zE0"},
	md5Datum{"EvanYo3327", "VDajAiZs", "lMKGzN91BhIX0hHCNqErU1"},
	md5Datum{"269john139", "Ryash8LF", "u96Rir1Izuwf/oHnaykmS/"},
	md5Datum{"3348159zw", "fdErikUY", ".gX/8MNguTOTWT35m4DCy/"},
	md5Datum{"lu184020", "uabGv1xC", "X5NNdH/1dzD0gQUyHwzKB0"},
	md5Datum{"aszasw", "41WiK.i.", "2q1CW/s4oRBLAFxmLESmO1"},
	md5Datum{"33059049", "bYPWMY2a", "fvKkFR1RRccGtIUhLuvBR0"},
	md5Datum{"li3255265", "FTGQVCcu", "QS/ub5DGLK/wgfkYQ0DBR."},
	md5Datum{"kerrihayes", "cFc9bc86", "3cVFy8/qB/fNGNueG65vG0"},
	md5Datum{"0167681809", "A5TvYYWy", "s4HBh0Wum2QQj1c9e0s79."},
	md5Datum{"stefano123", "YNrpseN3", "Yt52Yo9IEBs2LpX7A/CUb0"},
	md5Datum{"15054652730", "12CL4km4", "NJm8fh.JFi5dE.p6A9g7v/"},
	md5Datum{"natdvd213", "hssJjJTG", "dDK3pbBFTLbEigu.eCN7s."},
	md5Datum{"680929", "iaZlOft5", "w7iC6f5BUzuXox9THmHuj1"},
	md5Datum{"steelpad8", "mAoHmdUe", "5HePkkuSVu9F2UYgCvn0M."},
	md5Datum{"374710", "RFR4xs7H", "9GH0NjiDIgBD0t.w5/fwt0"},
	md5Datum{"394114", "Jt2syL5H", "tJ18tBNlcBEBqphUQc9jm."},
	md5Datum{"24347", "QnSWI03c", "8GC6c0AwpC.c8j4H7/9QU0"},
	md5Datum{"krait93", "bwzDGet.", "ntnX3fwzi3Zzhy0eHuwA9."},
	md5Datum{"5164794", "gkhv.jfD", "2fljug5HHu01vs.6KGJXQ."},
	md5Datum{"rswCyJE5", "HzyuhjzZ", "pXmWtTfn0/1voBaBkNaRy0"},
	md5Datum{"31480019", "ZZc0Ogd8", "1TNy1gTG6GLc.P/98kXXT."},
	md5Datum{"19830907ok", "4t6oHDY9", "kFoi2gvPcKMZs.AiGq1yb1"},
	md5Datum{"zlsmhzlsmh", "cih9diuY", "AwNc6TaKzFm9c8.kQxfwN1"},
	md5Datum{"Zengatsu", "wuXDXGlS", "FXFvRPPs7HHg96sSCFnFM1"},
	md5Datum{"0127603331", "z3inhAFw", "vkfbG7KVT4SYHiUn7Yqrz1"},
	md5Datum{"axelle77", "jydGNcWd", "qz3N5yqg0woVcZ6TN7SHr0"},
	md5Datum{"password2147", "GoP2TF8P", "c/b36Y.Qg/Grq7b7p.jbl."},
	md5Datum{"olixkl8b", "wxkU6WKQ", "IlhCpPwTWvESASvpOToqh."},
	md5Datum{"maiwen", "7JgCOFuj", "0WVRunftYuoR3o5ktLMdM1"},
	md5Datum{"198613", "Vai72CeM", "6WWXwZhxx/EW0IONm7n0A."},
	md5Datum{"s17kr8wu", "uNqfw7fr", "NAmeX1Mag2xf5lOCxGrcx/"},
	md5Datum{"biker02", ".dmc8gVd", "ZB4OmwWIeJ5Iy66Ta/7mU0"},
	md5Datum{"m1399", "vg1vnQVK", "UUqQibheBizuB0JxR1rbz/"},
	md5Datum{"a2dc6a", "lsH2FMPS", "dBBuRArwOlN/1p1BuncB3/"},
	md5Datum{"zhd8902960", "rMGc2ODd", "jG6/9kzAkMHFVAYYVEKN60"},
	md5Datum{"parasuta", "GeWoySy2", "WZ9pwqAb72tKP0xob81Ho0"},
	md5Datum{"the1secret", "7LW61iOz", "a9dFA0cRmBIuaxbBqnT/w/"},
	md5Datum{"teddy14", "GJ9nS.Cn", "jwpBiFBLr1XIo.J5klB39."},
	md5Datum{"4516388amt", "NEgOG19t", "CjfmPSbrJqUx6imCL4WPD/"},
	md5Datum{"245520", "rEzCqOtj", "sSblCTbLq2XDMTeDjYHMu0"},
	md5Datum{"D34dw00d", "Bugn2T/z", "gTZ/TZ24SMiL1AVQIPgam1"},
	md5Datum{"officiel", "oCnbHp3p", "lXVZn0P1qWe7dGRkwiJkj0"},
	md5Datum{"36653665", "cCwY3el7", ".sx/Uv4UADYdLSGjfI0gD0"},
	md5Datum{"hipol", "b0jFoiEY", "BELMMlTsgKPQ8jSloicdh."},
	md5Datum{"Nylon0", "cIw8xXs1", "uiDYDxgJsujwuQtU9Rjyr/"},
	md5Datum{"caitlyne6", "UffYyvRf", "IHrP6qbFVQEFwcl5BNh9j/"},
	md5Datum{"dogzilla", "2wvpCP1I", "vudGA0I1SLgEMr6xmmizy."},
	md5Datum{"lemegaboss", "QOdrh1Z.", "tFHoBTGKnHwf.MWzX7IBD/"},
	md5Datum{"c0valerius", "z4ckUwmA", "hq0/DLKdj/0PaR9uJ67fd1"},
	md5Datum{"liseczek44", "nPnWx0Kv", "FF9VO/i4rbKiD8p.Kor0x0"},
	md5Datum{"saulosi", "Ox3Y2bAv", "HBZQJd7esDSp/3StMc4xs1"},
	md5Datum{"53522", "VJn0Rpzz", "7CCQCvpxd3vVsBTIQNHmA1"},
	md5Datum{"ajgebam", "3wMf8geF", "vyqUHs9babWmAeAIHgcCJ0"},
	md5Datum{"freshplayer", "H6BJsnhE", "sdUNxVuP0wbG8GXYaaE3H0"},
	md5Datum{"logistica1", "ycXMTiTE", "8cXiewb9rsL9EuNi.ygaa/"},
	md5Datum{"12calo66", ".DEY1oqo", "TWeDNa7xX7W3sZWNTZKjG/"},
	md5Datum{"kenno", "QTq2YDtZ", "3b9BdtbYMbObjKa8.Fvy3/"},
	md5Datum{"34639399", "qAOAsxTH", "2c8ueVqVPiKAN2ihhA/xw."},
	md5Datum{"0408636405", "cLdGrOiq", "WedaFW4qjBLvBKWNZ98ik/"},
	md5Datum{"weezer12", "mY8WCPXG", "8xEw.ExVVzBOa9u3lJe/W/"},
	md5Datum{"9888735777", "4l3ZZKUa", "Nor5nWfN0h2HaeQwWBL3u."},
	md5Datum{"7771877", "3J0yl1xy", "1h9c1aatf.IaVJvkATLhE0"},
	md5Datum{"6620852", "UNtXqO0n", "Ag6gmPaH1guubjCy4bJHr0"},
	md5Datum{"98billiards", "4GJSSWxR", "wNggaBr4TH94zYGEuDvWX1"},
	md5Datum{"angelik", "Wo9Y7PP9", "btm.n8EiQMUnAFXtlqMpp/"},
	md5Datum{"86815057", "59qG1lpq", "C1efDS5Cyz33AEdcqNNjP/"},
	md5Datum{"p16alfalfa", "VW75OiLp", "EeU9NvGQn3l0es.EqOJyt1"},
	md5Datum{"7236118", "3mis3uOG", "sXNyXtdsWoNUpMaipVw3a."},
	md5Datum{"glock17l", "J1Vs.bJ4", "AULv/cwYjMeBoMTvEZXvU."},
	md5Datum{"sigmundm", ".k9ZvRfT", "lbGDjiA90kolu9DzQLOvv1"},
	md5Datum{"ltbgeqsd", "WT1wTKP8", "UDawOWZ73u8wBBZ7ohlSP0"},
	md5Datum{"wqnd8k2m", "mqiUjAJl", "xYZ0sN8LEwKrxU1g1Did30"},
	md5Datum{"yangjunjie", "wMWIiKAK", "yScptAfXmU8DVl6AVoAWB0"},
	md5Datum{"manjinder", "dOljUCkA", "pEb7LT2zG/qezaTTzd1Nj."},
	md5Datum{"nick2000", "9qhbsAfO", ".peZB9DgrJqAKlp2R1Nq70"},
	md5Datum{"193416", "Tke5EI49", "2suXXCRZuzJvjJ7QcJQMU1"},
	md5Datum{"pang168", "goNotyBA", "/lhn.zMA5z.a2VF31jaO3."},
	md5Datum{"454016", "1MdFKwJb", "/MBNPsDN66rZdg1SGQeKj1"},
	md5Datum{"phair08", "B3uB4Hl/", "LUqRKHuzcnb2q6xwqVok11"},
	md5Datum{"10252007cw", "ewVqnTQ1", "HkdOCIGKHYg193aUfQuer."},
	md5Datum{"zhuzhuzhu", "BiILrcFo", "tqGhsuOrQDvg/JPV00RSd/"},
	md5Datum{"metafunds", "dLMwXEWa", "Hq/WjMSgbxkp.wCelyfRX."},
	md5Datum{"smash", "aMgvovYi", "op3FHJ5OuM2tS93TKnhoc1"},
	md5Datum{"76387638", "GanQOcQh", "G5qdkoizpSOjWFc3PeL8D."},
	md5Datum{"S226811954", "GF9EM5zg", "whu07gAcDNRBfRInKdQz2."},
	md5Datum{"mintymoo00", "jDnIOwmz", "vBkkiacYuF8kcp1Nw3tf/1"},
	md5Datum{"seven711", "mwX.ezPE", "58Q31F7jya8UTnrFUzwO41"},
	md5Datum{"924414", "wcsVK7PY", "iOErsaSDD8l478QPn/ecp."},
	md5Datum{"changchengxu", "ON3zxaJ9", "4K0aR4n6JwbGM8jiE78eo1"},
	md5Datum{"alaska58", "KIIvW1ib", "ZqJQRoEoDpx30bt4HkZNO0"},
	md5Datum{"7678208", "xLTFhFu0", "wgkf1zwnwG.rwUGaHlzKK/"},
	md5Datum{"szazsoo73", "S8RvlMwv", "XKeXw9RfHH163LjG.yQ4/0"},
	md5Datum{"3830371", "E1WhUznq", "qUOza3gf2ZzUohYpnA/Gt/"},
	md5Datum{"0qdzx66b", "zSbUMRoi", "EJKnTL40qyiKNTWdOkg8K1"},
	md5Datum{"09124248099", "vkxQrmli", "gfLBcPOpLI.x4BHcGgG5o1"},
	md5Datum{"bachrain", "i74JdOeY", "l/rxskCai9U2yu6QAuYiP0"},
	md5Datum{"sJsSdFBY", "Ucs2cgJv", "ltZWhw3rvDThU3h4wTiMR0"},
	md5Datum{"676215000", "PJ52qkEa", "FVxkESgiPU8HVk9CVr5Aw0"},
	md5Datum{"nimamapwoaini", "iJhvvMzV", "c11ZLkLbU3oTL0tO4Uc2b0"},
	md5Datum{"nitsuj", "Eg6C/017", "PBjnkuRuhfwSMso1of0CU/"},
	md5Datum{"cukierek2003", "DtaGU5uw", "wj9U6W39HosDe4d20aq9b0"},
	md5Datum{"seeder", "Hu7E7fh9", "ro5jNBVSUr7P3xXB7bWTs1"},
	md5Datum{"00167148786", "kIAtp5Qp", "0mGyQcPNotlS9PXmD8VLX/"},
	md5Datum{"ashok198", "yz/u5zIx", "TcuTnX2cLRkGGPWuQ1DHe0"},
	md5Datum{"kt2116", "zIlMHa5m", "v.HKzAXRicCxQlNwap5r5/"},
	md5Datum{"another82", "kf0a2hjv", ".8kEpY7NyyNfBs4Udeu2T."},
	md5Datum{"75995794", "2AcSlaOt", "PdPz3ooJyaCM4rD9AuS4c/"},
	md5Datum{"19901130", "4wioa3Us", "uaKSWrWjJlqHdsqBdF7Zr."},
	md5Datum{"gijs010389", "4D9hzr6I", "PsnXK455GeQ3NCdOHmoSY1"},
	md5Datum{"26263199", "rXOrEHJ9", "atQhaNEYAfdzht02mRZcg."},
	md5Datum{"hi1j42x8", "i8PdGfO7", "Xv.aSLFQjyqbJ1KnM9hCs1"},
	md5Datum{"6922235", "I2xWkhl3", "oth511sBJphjpr0chWodC1"},
	md5Datum{"67749330", "AGNgrF8B", "KBcUjzo9d3pXFNsUCD6Ur1"},
	md5Datum{"ccpatrik", "zuNtiCs2", "54MqesBdp3RoL98/fklXb/"},
	md5Datum{"summer3011", "ZK2FB9JV", "8x8Ug7Jh3oWXgxWrLBuhr."},
	md5Datum{"331516", "UoqGMAIH", "bEG70EwRgt0SC6h5nr1wY1"},
	md5Datum{"135745", "DTVm48a7", "KE/H8KTGE0gi9wxM.ZzOs/"},
	md5Datum{"603762004", "0B44zHt5", "Xsbx3F0DtToD.KHYc5ViP1"},
	md5Datum{"29011985", "2YOvrTZM", "/n5Fol4IfYqLv9tS/QWWj0"},
}

func Test_apr1Md5(t *testing.T) {
	for _, v := range md5TestData {
		if r := apr1Md5(v.password, v.salt); r != v.hashed {
			t.Errorf("apr1Md5(%v,%v) is wrong: %v != %v", v.password, v.salt, r, v.hashed)
		}
	}
}

func Test_Md5(t *testing.T) {
	for _, v := range md5TestData {
		text := fmt.Sprintf("$apr1$%s$%s", v.salt, v.hashed)
		testParserGood(t, "md5", AcceptMd5, RejectMd5, text, v.password)
	}
	testParserBad(t, "md5", AcceptMd5, RejectMd5, "$apr1$nosalt")
	testParserNot(t, "md5", AcceptMd5, RejectMd5, "plain")
	testParserNot(t, "md5", AcceptMd5, RejectMd5, "{SHA}plain")
}
