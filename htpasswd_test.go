package htpasswd

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type testUser struct {
	username string
	password string
}

var testUsers = []testUser{
	testUser{"user1", "mickey5"},
	testUser{"user2", "alexandrew"},
	testUser{"user3", "hawaiicats78"},
	testUser{"user4", "DIENOW"},
	testUser{"user5", "e8f685"},
	testUser{"user6", "Rickygirl03"},
	testUser{"user7", "123vb123"},
	testUser{"user8", "sheng060576"},
	testUser{"user9", "hansisme"},
	testUser{"user10", "h4ck3rs311t3"},
	testUser{"user11", "K90JyTGA"},
	testUser{"user12", "aspire5101"},
	testUser{"user13", "553568"},
	testUser{"user14", "SRI"},
	testUser{"user15", "maxmus"},
	testUser{"user16", "a5xp9707"},
	testUser{"user17", "tomasrim"},
	testUser{"user18", "2a0mag"},
	testUser{"user19", "wmsfht"},
	testUser{"user20", "webmaster2364288"},
	testUser{"user21", "121516m"},
	testUser{"user22", "T69228803"},
	testUser{"user23", "qq820221"},
	testUser{"user24", "chenfy"},
	testUser{"user25", "www.debure.net"},
	testUser{"user26", "1333e763"},
	testUser{"user27", "burberries"},
	testUser{"user28", "chanmee14"},
	testUser{"user29", "65432106543210"},
	testUser{"user30", "powernet"},
	testUser{"user31", "a2d8i6a7"},
	testUser{"user32", "gvs9ptc"},
	testUser{"user33", "Pookie"},
	testUser{"user34", "lorissss"},
	testUser{"user35", "ess"},
	testUser{"user36", "sparra"},
	testUser{"user37", "allysson"},
	testUser{"user38", "99128008"},
	testUser{"user39", "evisanne"},
	testUser{"user40", "qfxg7x9l"},
	testUser{"user41", "03415"},
	testUser{"user42", "87832309"},
	testUser{"user43", "816283"},
	testUser{"user44", "banach12"},
	testUser{"user45", "sjdszpsc"},
	testUser{"user46", "changsing"},
	testUser{"user47", "56339388"},
	testUser{"user48", "52114157"},
	testUser{"user49", "jinebimb"},
	testUser{"user50", "erol43"},
	testUser{"user51", "2yagos"},
	testUser{"user52", "habparty!"},
	testUser{"user53", "tangjianhui"},
	testUser{"user54", "serandah"},
	testUser{"user55", "mirrages"},
	testUser{"user56", "mantgaxxl"},
	testUser{"user57", "45738901"},
	testUser{"user58", "g523minna"},
	testUser{"user59", "j202020"},
	testUser{"user60", "g@mmaecho"},
	testUser{"user61", "042380"},
	testUser{"user62", "ASRuin"},
	testUser{"user63", "061990"},
	testUser{"user64", "ysoline"},
	testUser{"user65", "liuzhouzhou"},
	testUser{"user66", "b0000000wind"},
	testUser{"user67", "7913456852"},
	testUser{"user68", "9008"},
	testUser{"user69", "waitlin11"},
	testUser{"user70", "8fdakar"},
	testUser{"user71", "eisball"},
	testUser{"user72", "jenna17"},
	testUser{"user73", "belkadonam"},
	testUser{"user74", "tfyuj9JW"},
	testUser{"user75", "nihaijidema"},
	testUser{"user76", "talapia"},
	testUser{"user77", "7376220"},
	testUser{"user78", "c7m8e1xsc3"},
	testUser{"user79", "84129793"},
	testUser{"user80", "test1000"},
	testUser{"user81", "ecmanhatten"},
	testUser{"user82", "EvanYo3327"},
	testUser{"user83", "269john139"},
	testUser{"user84", "3348159zw"},
	testUser{"user85", "lu184020"},
	testUser{"user86", "aszasw"},
	testUser{"user87", "33059049"},
	testUser{"user88", "li3255265"},
	testUser{"user89", "kerrihayes"},
	testUser{"user90", "0167681809"},
	testUser{"user91", "stefano123"},
	testUser{"user92", "15054652730"},
	testUser{"user93", "natdvd213"},
	testUser{"user94", "680929"},
	testUser{"user95", "steelpad8"},
	testUser{"user96", "374710"},
	testUser{"user97", "394114"},
	testUser{"user98", "24347"},
	testUser{"user99", "krait93"},
	testUser{"user100", "5164794"},
	testUser{"user101", "rswCyJE5"},
	testUser{"user102", "31480019"},
	testUser{"user103", "19830907ok"},
	testUser{"user104", "zlsmhzlsmh"},
	testUser{"user105", "Zengatsu"},
	testUser{"user106", "0127603331"},
	testUser{"user107", "axelle77"},
	testUser{"user108", "password2147"},
	testUser{"user109", "olixkl8b"},
	testUser{"user110", "maiwen"},
	testUser{"user111", "198613"},
	testUser{"user112", "s17kr8wu"},
	testUser{"user113", "biker02"},
	testUser{"user114", "m1399"},
	testUser{"user115", "a2dc6a"},
	testUser{"user116", "zhd8902960"},
	testUser{"user117", "parasuta"},
	testUser{"user118", "the1secret"},
	testUser{"user119", "teddy14"},
	testUser{"user120", "4516388amt"},
	testUser{"user121", "245520"},
	testUser{"user122", "D34dw00d"},
	testUser{"user123", "officiel"},
	testUser{"user124", "36653665"},
	testUser{"user125", "hipol"},
	testUser{"user126", "Nylon0"},
	testUser{"user127", "caitlyne6"},
	testUser{"user128", "dogzilla"},
	testUser{"user129", "lemegaboss"},
	testUser{"user130", "c0valerius"},
	testUser{"user131", "liseczek44"},
	testUser{"user132", "saulosi"},
	testUser{"user133", "53522"},
	testUser{"user134", "ajgebam"},
	testUser{"user135", "freshplayer"},
	testUser{"user136", "logistica1"},
	testUser{"user137", "12calo66"},
	testUser{"user138", "kenno"},
	testUser{"user139", "34639399"},
	testUser{"user140", "0408636405"},
	testUser{"user141", "weezer12"},
	testUser{"user142", "9888735777"},
	testUser{"user143", "7771877"},
	testUser{"user144", "6620852"},
	testUser{"user145", "98billiards"},
	testUser{"user146", "angelik"},
	testUser{"user147", "86815057"},
	testUser{"user148", "p16alfalfa"},
	testUser{"user149", "7236118"},
	testUser{"user150", "glock17l"},
	testUser{"user151", "sigmundm"},
	testUser{"user152", "ltbgeqsd"},
	testUser{"user153", "wqnd8k2m"},
	testUser{"user154", "yangjunjie"},
	testUser{"user155", "manjinder"},
	testUser{"user156", "nick2000"},
	testUser{"user157", "193416"},
	testUser{"user158", "pang168"},
	testUser{"user159", "454016"},
	testUser{"user160", "phair08"},
	testUser{"user161", "10252007cw"},
	testUser{"user162", "zhuzhuzhu"},
	testUser{"user163", "metafunds"},
	testUser{"user164", "smash"},
	testUser{"user165", "76387638"},
	testUser{"user166", "S226811954"},
	testUser{"user167", "mintymoo00"},
	testUser{"user168", "seven711"},
	testUser{"user169", "924414"},
	testUser{"user170", "changchengxu"},
	testUser{"user171", "alaska58"},
	testUser{"user172", "7678208"},
	testUser{"user173", "szazsoo73"},
	testUser{"user174", "3830371"},
	testUser{"user175", "0qdzx66b"},
	testUser{"user176", "09124248099"},
	testUser{"user177", "bachrain"},
	testUser{"user178", "sJsSdFBY"},
	testUser{"user179", "676215000"},
	testUser{"user180", "nimamapwoaini"},
	testUser{"user181", "nitsuj"},
	testUser{"user182", "cukierek2003"},
	testUser{"user183", "seeder"},
	testUser{"user184", "00167148786"},
	testUser{"user185", "ashok198"},
	testUser{"user186", "kt2116"},
	testUser{"user187", "another82"},
	testUser{"user188", "75995794"},
	testUser{"user189", "19901130"},
	testUser{"user190", "gijs010389"},
	testUser{"user191", "26263199"},
	testUser{"user192", "hi1j42x8"},
	testUser{"user193", "6922235"},
	testUser{"user194", "67749330"},
	testUser{"user195", "ccpatrik"},
	testUser{"user196", "summer3011"},
	testUser{"user197", "331516"},
	testUser{"user198", "135745"},
	testUser{"user199", "603762004"},
	testUser{"user200", "29011985"},
}
var textPlain = `user1:mickey5
user2:alexandrew
user3:hawaiicats78
user4:DIENOW
user5:e8f685
user6:Rickygirl03
user7:123vb123
user8:sheng060576
user9:hansisme
user10:h4ck3rs311t3
user11:K90JyTGA
user12:aspire5101
user13:553568
user14:SRI
user15:maxmus
user16:a5xp9707
user17:tomasrim
user18:2a0mag
user19:wmsfht
user20:webmaster2364288
user21:121516m
user22:T69228803
user23:qq820221
user24:chenfy
user25:www.debure.net
user26:1333e763
user27:burberries
user28:chanmee14
user29:65432106543210
user30:powernet
user31:a2d8i6a7
user32:gvs9ptc
user33:Pookie
user34:lorissss
user35:ess
user36:sparra
user37:allysson
user38:99128008
user39:evisanne
user40:qfxg7x9l
user41:03415
user42:87832309
user43:816283
user44:banach12
user45:sjdszpsc
user46:changsing
user47:56339388
user48:52114157
user49:jinebimb
user50:erol43
user51:2yagos
user52:habparty!
user53:tangjianhui
user54:serandah
user55:mirrages
user56:mantgaxxl
user57:45738901
user58:g523minna
user59:j202020
user60:g@mmaecho
user61:042380
user62:ASRuin
user63:061990
user64:ysoline
user65:liuzhouzhou
user66:b0000000wind
user67:7913456852
user68:9008
user69:waitlin11
user70:8fdakar
user71:eisball
user72:jenna17
user73:belkadonam
user74:tfyuj9JW
user75:nihaijidema
user76:talapia
user77:7376220
user78:c7m8e1xsc3
user79:84129793
user80:test1000
user81:ecmanhatten
user82:EvanYo3327
user83:269john139
user84:3348159zw
user85:lu184020
user86:aszasw
user87:33059049
user88:li3255265
user89:kerrihayes
user90:0167681809
user91:stefano123
user92:15054652730
user93:natdvd213
user94:680929
user95:steelpad8
user96:374710
user97:394114
user98:24347
user99:krait93
user100:5164794
user101:rswCyJE5
user102:31480019
user103:19830907ok
user104:zlsmhzlsmh
user105:Zengatsu
user106:0127603331
user107:axelle77
user108:password2147
user109:olixkl8b
user110:maiwen
user111:198613
user112:s17kr8wu
user113:biker02
user114:m1399
user115:a2dc6a
user116:zhd8902960
user117:parasuta
user118:the1secret
user119:teddy14
user120:4516388amt
user121:245520
user122:D34dw00d
user123:officiel
user124:36653665
user125:hipol
user126:Nylon0
user127:caitlyne6
user128:dogzilla
user129:lemegaboss
user130:c0valerius
user131:liseczek44
user132:saulosi
user133:53522
user134:ajgebam
user135:freshplayer
user136:logistica1
user137:12calo66
user138:kenno
user139:34639399
user140:0408636405
user141:weezer12
user142:9888735777
user143:7771877
user144:6620852
user145:98billiards
user146:angelik
user147:86815057
user148:p16alfalfa
user149:7236118
user150:glock17l
user151:sigmundm
user152:ltbgeqsd
user153:wqnd8k2m
user154:yangjunjie
user155:manjinder
user156:nick2000
user157:193416
user158:pang168
user159:454016
user160:phair08
user161:10252007cw
user162:zhuzhuzhu
user163:metafunds
user164:smash
user165:76387638
user166:S226811954
user167:mintymoo00
user168:seven711
user169:924414
user170:changchengxu
user171:alaska58
user172:7678208
user173:szazsoo73
user174:3830371
user175:0qdzx66b
user176:09124248099
user177:bachrain
user178:sJsSdFBY
user179:676215000
user180:nimamapwoaini
user181:nitsuj
user182:cukierek2003
user183:seeder
user184:00167148786
user185:ashok198
user186:kt2116
user187:another82
user188:75995794
user189:19901130
user190:gijs010389
user191:26263199
user192:hi1j42x8
user193:6922235
user194:67749330
user195:ccpatrik
user196:summer3011
user197:331516
user198:135745
user199:603762004
user200:29011985`

var textApr1 = `user1:$apr1$gxNb79DX$6wi9QaGNM5TA0kBKiC4710
user2:$apr1$kv1uUfCO$iEwrWojf92uZ/9uhTQmMo.
user3:$apr1$UQ6GxE7V$OrIqWONGuSV9RfS3B2dfO1
user4:$apr1$OZ.RwYJH$AwfW2h0gJnu2fQi0GegVe1
user5:$apr1$9r9GyMpL$3IiaLNos/tbouLJwsW8ey/
user6:$apr1$0tlsxL/0$cfS6c2JZjwISRTgFvrMWL1
user7:$apr1$/4XFfQuK$bnMIHM0j/Cf8apmbvPzn/.
user8:$apr1$NEJJUzVT$o/CWI9InAMXWAsbl5gx0p1
user9:$apr1$JAOXCriK$gB/Yox3wTae3NujwKUiFv1
user10:$apr1$KmkPgS2r$5qIFMPNVAXzlevkzOQwhj.
user11:$apr1$mM7q5ZHN$03LeGh9D1CujEBwiVRO6B0
user12:$apr1$tlxr3zoa$dQJiJmk4pEtRTssYiLwlv0
user13:$apr1$YI.r2X/w$H/1DtcmTHSgcdkgz8NS1W0
user14:$apr1$StJ5t4wb$tIVEx.MPZR1SqDm5y9VCs1
user15:$apr1$ad29tH08$xEHwr706Yz/3FFGqnVB6l/
user16:$apr1$aH0sN4io$y0heNz5hL67/HA7/7mDRS.
user17:$apr1$SgbYnJV9$7Z.enu6vZ7b6Zo7/lYce60
user18:$apr1$lSOzbc7i$Ae21yFmdTMpSz.aQsjyoE1
user19:$apr1$yicl6/5x$p/dCDdQ0q9lLaZbBJsIDP0
user20:$apr1$PLoY5sMf$KEDmvJskiSNFwiygtWXin1
user21:$apr1$3T5gmyrq$AucgLmXU53aTQJuRKCFo50
user22:$apr1$Aajhupso$/EPFyux8bd7Iw.tLevaVE.
user23:$apr1$G43B4jFl$4TUFaOD7Fz5.lZiq5v8P40
user24:$apr1$mDnux.Mf$vXsdihwaTLCJTHnuk9/cK/
user25:$apr1$bZzoRW4K$DfI3Col55.57HP3FW4L1h.
user26:$apr1$rRvCcrzo$plG5/rpEPSM7uc3bro6P51
user27:$apr1$Qx6JtYcz$10t2dI6u0LyNBjeCAQ.3z1
user28:$apr1$p9t9dUC1$Nlr96oZWIe/VVpYBUgG6q0
user29:$apr1$CBG7TUqG$Olyygy0L6HfSPfkLg24U60
user30:$apr1$ogVPlakG$.SLiqbN/KUECQ6pgdck2/.
user31:$apr1$sNrtmvPF$rvbRuKdcPPvN.dK.mHeYq/
user32:$apr1$gQgMMxVG$5sI4ezBQxqpfh14AvEEVU0
user33:$apr1$x.wVLgoG$HTj5gT.lQ71BpifSlcQVy1
user34:$apr1$O0ySiIf4$AmmYBbHWjfiVcGEbl4wiy/
user35:$apr1$nE19zEmy$Rg3/wMTNMVOkbhez/QhD//
user36:$apr1$By1OjZuF$PRY4G6D8u3aFhruSTgIQC.
user37:$apr1$mI6fsU64$WqCg/f9CpYr4586AVr6nP.
user38:$apr1$LQLXA.du$kazspxn165TFSiDavu75N/
user39:$apr1$AhDCR8bW$2lR137DLMfr1mQ9xLlMsw0
user40:$apr1$ZAGBUFGw$HI3fWMR0Y6Z4U3MSc70sd.
user41:$apr1$nkFIBpLJ$AvABMUIgvoMp0zmOTCwCG1
user42:$apr1$WbCq7Hv8$dxe0LoM3vlD.t/A/3Cfd11
user43:$apr1$PrEjUTNt$vGLTgLqJp9XEtwEJBv5XF.
user44:$apr1$S1G5jLiH$CySeS1zgVlMLLElxG6Dmw0
user45:$apr1$QmuQrgcB$xZk5zcK2QRF8PZ24P9vPr1
user46:$apr1$Z0i29yA5$KnTYiWEZQYzQlH/SxQ7Qp/
user47:$apr1$RZlCHiTm$8mFKCLkRHxoJ2ieVa.K79/
user48:$apr1$3NkMs.IK$02HiBvqlIVA.hLbktlHsD1
user49:$apr1$1ww3avga$haxtp7TGUm9PHPBrBeM9u.
user50:$apr1$.aE1EJya$3zkhvRyNbF.DOOyJSPSJ21
user51:$apr1$L0YlhvFW$R0J.Bk9wYb7sQKXBbP4AN/
user52:$apr1$vveX0m/D$hPoF3j.Ac5zSOAmHBZklT.
user53:$apr1$8Ivzj66d$J4A.NOn6TRk4RYC9oGqIB/
user54:$apr1$v9AJex0e$qn/isKH9e6EG66KCtFdmI1
user55:$apr1$UM0E3yNn$4V4IJI2Q0Bqh0EG8HAHbq0
user56:$apr1$1spakyg4$NwPcxatLI7bWUpeDzAw2h1
user57:$apr1$oepJpf/s$p0F.JGVJCyvUHfWnpF.Wy1
user58:$apr1$yWpavB.B$q4KExAyIKMKWTLq86n0820
user59:$apr1$DTRNSWt7$2At.lEmBM2waU9F2QsDvd.
user60:$apr1$QGA07jk6$U9Uw/dD666GNV60hX6AKM/
user61:$apr1$FDnW17iI$6jkNwkfAi.4LMYkIkNO2v1
user62:$apr1$GKFI0Se3$go4Tko/O9UCA2WtSJBjgc.
user63:$apr1$yJR0EnuF$CzzsiUo2Q5cRhtlptUf7D1
user64:$apr1$7D0hCvVq$HLIRmi013HBi2TgATkgJM.
user65:$apr1$m.MSvKt4$oFYUki/pESjwOfF5YH9VO0
user66:$apr1$qOQrkTXw$PJXv2X.0Efe4VUPcvyxA61
user67:$apr1$lPKDpKzC$q9kt0R9.I4rxhlIcNe2gg1
user68:$apr1$PYsksC92$3oqtOxrMnQc1n3GfSIAJM.
user69:$apr1$x5UDLNO2$yHLWIm/50ORtDhT56f9bi0
user70:$apr1$E9a2XIvt$fcsw4gZfbiDXPywMzwhik1
user71:$apr1$gHg16GuT$DGI/O8HzZemhsQ4o2jA560
user72:$apr1$yzwqt8mS$3QqqiFB9Z6q1fp4z/q1pU.
user73:$apr1$iGU4vuaZ$w3xf5rVAIJYz0dgImL8a2.
user74:$apr1$5cPUmio7$wttScNV7Fk4Njs9QX1yUi.
user75:$apr1$DZW4Gt4h$EXlVFPbqnXGPp2vLQT5TK0
user76:$apr1$61i3ruRm$cNcNvti2hQ8mXjLahFnSb/
user77:$apr1$Z89Ynh0K$A2k6aLQnMOa2uwXX8MJZf1
user78:$apr1$QRn4AsCM$gUztH0RWKuX1Vy0WaYfdC1
user79:$apr1$rghudgt5$XA7QLtfRq84JHtbjdke0I.
user80:$apr1$zwkIVA3j$Iuz7zNyLvIiKWIl2VA8bl.
user81:$apr1$zfVlWDS.$emJhRC3N0SnvZLo5en4zE0
user82:$apr1$VDajAiZs$lMKGzN91BhIX0hHCNqErU1
user83:$apr1$Ryash8LF$u96Rir1Izuwf/oHnaykmS/
user84:$apr1$fdErikUY$.gX/8MNguTOTWT35m4DCy/
user85:$apr1$uabGv1xC$X5NNdH/1dzD0gQUyHwzKB0
user86:$apr1$41WiK.i.$2q1CW/s4oRBLAFxmLESmO1
user87:$apr1$bYPWMY2a$fvKkFR1RRccGtIUhLuvBR0
user88:$apr1$FTGQVCcu$QS/ub5DGLK/wgfkYQ0DBR.
user89:$apr1$cFc9bc86$3cVFy8/qB/fNGNueG65vG0
user90:$apr1$A5TvYYWy$s4HBh0Wum2QQj1c9e0s79.
user91:$apr1$YNrpseN3$Yt52Yo9IEBs2LpX7A/CUb0
user92:$apr1$12CL4km4$NJm8fh.JFi5dE.p6A9g7v/
user93:$apr1$hssJjJTG$dDK3pbBFTLbEigu.eCN7s.
user94:$apr1$iaZlOft5$w7iC6f5BUzuXox9THmHuj1
user95:$apr1$mAoHmdUe$5HePkkuSVu9F2UYgCvn0M.
user96:$apr1$RFR4xs7H$9GH0NjiDIgBD0t.w5/fwt0
user97:$apr1$Jt2syL5H$tJ18tBNlcBEBqphUQc9jm.
user98:$apr1$QnSWI03c$8GC6c0AwpC.c8j4H7/9QU0
user99:$apr1$bwzDGet.$ntnX3fwzi3Zzhy0eHuwA9.
user100:$apr1$gkhv.jfD$2fljug5HHu01vs.6KGJXQ.
user101:$apr1$HzyuhjzZ$pXmWtTfn0/1voBaBkNaRy0
user102:$apr1$ZZc0Ogd8$1TNy1gTG6GLc.P/98kXXT.
user103:$apr1$4t6oHDY9$kFoi2gvPcKMZs.AiGq1yb1
user104:$apr1$cih9diuY$AwNc6TaKzFm9c8.kQxfwN1
user105:$apr1$wuXDXGlS$FXFvRPPs7HHg96sSCFnFM1
user106:$apr1$z3inhAFw$vkfbG7KVT4SYHiUn7Yqrz1
user107:$apr1$jydGNcWd$qz3N5yqg0woVcZ6TN7SHr0
user108:$apr1$GoP2TF8P$c/b36Y.Qg/Grq7b7p.jbl.
user109:$apr1$wxkU6WKQ$IlhCpPwTWvESASvpOToqh.
user110:$apr1$7JgCOFuj$0WVRunftYuoR3o5ktLMdM1
user111:$apr1$Vai72CeM$6WWXwZhxx/EW0IONm7n0A.
user112:$apr1$uNqfw7fr$NAmeX1Mag2xf5lOCxGrcx/
user113:$apr1$.dmc8gVd$ZB4OmwWIeJ5Iy66Ta/7mU0
user114:$apr1$vg1vnQVK$UUqQibheBizuB0JxR1rbz/
user115:$apr1$lsH2FMPS$dBBuRArwOlN/1p1BuncB3/
user116:$apr1$rMGc2ODd$jG6/9kzAkMHFVAYYVEKN60
user117:$apr1$GeWoySy2$WZ9pwqAb72tKP0xob81Ho0
user118:$apr1$7LW61iOz$a9dFA0cRmBIuaxbBqnT/w/
user119:$apr1$GJ9nS.Cn$jwpBiFBLr1XIo.J5klB39.
user120:$apr1$NEgOG19t$CjfmPSbrJqUx6imCL4WPD/
user121:$apr1$rEzCqOtj$sSblCTbLq2XDMTeDjYHMu0
user122:$apr1$Bugn2T/z$gTZ/TZ24SMiL1AVQIPgam1
user123:$apr1$oCnbHp3p$lXVZn0P1qWe7dGRkwiJkj0
user124:$apr1$cCwY3el7$.sx/Uv4UADYdLSGjfI0gD0
user125:$apr1$b0jFoiEY$BELMMlTsgKPQ8jSloicdh.
user126:$apr1$cIw8xXs1$uiDYDxgJsujwuQtU9Rjyr/
user127:$apr1$UffYyvRf$IHrP6qbFVQEFwcl5BNh9j/
user128:$apr1$2wvpCP1I$vudGA0I1SLgEMr6xmmizy.
user129:$apr1$QOdrh1Z.$tFHoBTGKnHwf.MWzX7IBD/
user130:$apr1$z4ckUwmA$hq0/DLKdj/0PaR9uJ67fd1
user131:$apr1$nPnWx0Kv$FF9VO/i4rbKiD8p.Kor0x0
user132:$apr1$Ox3Y2bAv$HBZQJd7esDSp/3StMc4xs1
user133:$apr1$VJn0Rpzz$7CCQCvpxd3vVsBTIQNHmA1
user134:$apr1$3wMf8geF$vyqUHs9babWmAeAIHgcCJ0
user135:$apr1$H6BJsnhE$sdUNxVuP0wbG8GXYaaE3H0
user136:$apr1$ycXMTiTE$8cXiewb9rsL9EuNi.ygaa/
user137:$apr1$.DEY1oqo$TWeDNa7xX7W3sZWNTZKjG/
user138:$apr1$QTq2YDtZ$3b9BdtbYMbObjKa8.Fvy3/
user139:$apr1$qAOAsxTH$2c8ueVqVPiKAN2ihhA/xw.
user140:$apr1$cLdGrOiq$WedaFW4qjBLvBKWNZ98ik/
user141:$apr1$mY8WCPXG$8xEw.ExVVzBOa9u3lJe/W/
user142:$apr1$4l3ZZKUa$Nor5nWfN0h2HaeQwWBL3u.
user143:$apr1$3J0yl1xy$1h9c1aatf.IaVJvkATLhE0
user144:$apr1$UNtXqO0n$Ag6gmPaH1guubjCy4bJHr0
user145:$apr1$4GJSSWxR$wNggaBr4TH94zYGEuDvWX1
user146:$apr1$Wo9Y7PP9$btm.n8EiQMUnAFXtlqMpp/
user147:$apr1$59qG1lpq$C1efDS5Cyz33AEdcqNNjP/
user148:$apr1$VW75OiLp$EeU9NvGQn3l0es.EqOJyt1
user149:$apr1$3mis3uOG$sXNyXtdsWoNUpMaipVw3a.
user150:$apr1$J1Vs.bJ4$AULv/cwYjMeBoMTvEZXvU.
user151:$apr1$.k9ZvRfT$lbGDjiA90kolu9DzQLOvv1
user152:$apr1$WT1wTKP8$UDawOWZ73u8wBBZ7ohlSP0
user153:$apr1$mqiUjAJl$xYZ0sN8LEwKrxU1g1Did30
user154:$apr1$wMWIiKAK$yScptAfXmU8DVl6AVoAWB0
user155:$apr1$dOljUCkA$pEb7LT2zG/qezaTTzd1Nj.
user156:$apr1$9qhbsAfO$.peZB9DgrJqAKlp2R1Nq70
user157:$apr1$Tke5EI49$2suXXCRZuzJvjJ7QcJQMU1
user158:$apr1$goNotyBA$/lhn.zMA5z.a2VF31jaO3.
user159:$apr1$1MdFKwJb$/MBNPsDN66rZdg1SGQeKj1
user160:$apr1$B3uB4Hl/$LUqRKHuzcnb2q6xwqVok11
user161:$apr1$ewVqnTQ1$HkdOCIGKHYg193aUfQuer.
user162:$apr1$BiILrcFo$tqGhsuOrQDvg/JPV00RSd/
user163:$apr1$dLMwXEWa$Hq/WjMSgbxkp.wCelyfRX.
user164:$apr1$aMgvovYi$op3FHJ5OuM2tS93TKnhoc1
user165:$apr1$GanQOcQh$G5qdkoizpSOjWFc3PeL8D.
user166:$apr1$GF9EM5zg$whu07gAcDNRBfRInKdQz2.
user167:$apr1$jDnIOwmz$vBkkiacYuF8kcp1Nw3tf/1
user168:$apr1$mwX.ezPE$58Q31F7jya8UTnrFUzwO41
user169:$apr1$wcsVK7PY$iOErsaSDD8l478QPn/ecp.
user170:$apr1$ON3zxaJ9$4K0aR4n6JwbGM8jiE78eo1
user171:$apr1$KIIvW1ib$ZqJQRoEoDpx30bt4HkZNO0
user172:$apr1$xLTFhFu0$wgkf1zwnwG.rwUGaHlzKK/
user173:$apr1$S8RvlMwv$XKeXw9RfHH163LjG.yQ4/0
user174:$apr1$E1WhUznq$qUOza3gf2ZzUohYpnA/Gt/
user175:$apr1$zSbUMRoi$EJKnTL40qyiKNTWdOkg8K1
user176:$apr1$vkxQrmli$gfLBcPOpLI.x4BHcGgG5o1
user177:$apr1$i74JdOeY$l/rxskCai9U2yu6QAuYiP0
user178:$apr1$Ucs2cgJv$ltZWhw3rvDThU3h4wTiMR0
user179:$apr1$PJ52qkEa$FVxkESgiPU8HVk9CVr5Aw0
user180:$apr1$iJhvvMzV$c11ZLkLbU3oTL0tO4Uc2b0
user181:$apr1$Eg6C/017$PBjnkuRuhfwSMso1of0CU/
user182:$apr1$DtaGU5uw$wj9U6W39HosDe4d20aq9b0
user183:$apr1$Hu7E7fh9$ro5jNBVSUr7P3xXB7bWTs1
user184:$apr1$kIAtp5Qp$0mGyQcPNotlS9PXmD8VLX/
user185:$apr1$yz/u5zIx$TcuTnX2cLRkGGPWuQ1DHe0
user186:$apr1$zIlMHa5m$v.HKzAXRicCxQlNwap5r5/
user187:$apr1$kf0a2hjv$.8kEpY7NyyNfBs4Udeu2T.
user188:$apr1$2AcSlaOt$PdPz3ooJyaCM4rD9AuS4c/
user189:$apr1$4wioa3Us$uaKSWrWjJlqHdsqBdF7Zr.
user190:$apr1$4D9hzr6I$PsnXK455GeQ3NCdOHmoSY1
user191:$apr1$rXOrEHJ9$atQhaNEYAfdzht02mRZcg.
user192:$apr1$i8PdGfO7$Xv.aSLFQjyqbJ1KnM9hCs1
user193:$apr1$I2xWkhl3$oth511sBJphjpr0chWodC1
user194:$apr1$AGNgrF8B$KBcUjzo9d3pXFNsUCD6Ur1
user195:$apr1$zuNtiCs2$54MqesBdp3RoL98/fklXb/
user196:$apr1$ZK2FB9JV$8x8Ug7Jh3oWXgxWrLBuhr.
user197:$apr1$UoqGMAIH$bEG70EwRgt0SC6h5nr1wY1
user198:$apr1$DTVm48a7$KE/H8KTGE0gi9wxM.ZzOs/
user199:$apr1$0B44zHt5$Xsbx3F0DtToD.KHYc5ViP1
user200:$apr1$2YOvrTZM$/n5Fol4IfYqLv9tS/QWWj0`

var textSha = `user1:{SHA}D9rQ8iK6feNAniulHNKdr5V38ok=
user2:{SHA}KS7VQqgAnMUfXgWmFCCa6DVhY+M=
user3:{SHA}mzD9ouM0P06arY0Obdb2KojkFeY=
user4:{SHA}2HrOk971ockoAr1Ct1o7GpvFLdU=
user5:{SHA}IyrjpSzIjrlLT7KjVh1q1LBDCFA=
user6:{SHA}gh5ZdWJ6UypV+CRv8Kd0herEe8U=
user7:{SHA}19j8b+o4MIImvb7zkv4gncUWs1s=
user8:{SHA}u5hjc4kFFwv2QrMa49uCmn7RmPQ=
user9:{SHA}Au92ail1kvpgW+HC4PiuJu1Yays=
user10:{SHA}KvvqIY6j/LU3D/orcNYioTE4yOQ=
user11:{SHA}QftskBtqwp5hsovHb+dGB7azRsI=
user12:{SHA}vjUxr9fJRAYTdO0lZb83C2z0FKo=
user13:{SHA}FTguhILOzfejAav4aWK2X4D51qE=
user14:{SHA}7CPWhF7FKaS6fBC+zAuncRjVDOo=
user15:{SHA}tHLz2JrCsGMOAPUh38Td1yP65rI=
user16:{SHA}qTzg8xc+MFLHwsR9DFVuUTPiFA8=
user17:{SHA}BxtDkfZMLEfKDlqluTcU/UOFknw=
user18:{SHA}cnCoTKX3EK7nIA+zmxcu3ovItrM=
user19:{SHA}pBahyhKlvfj+JoJNul/LNITFz+A=
user20:{SHA}NQb47c348gpwiOL+NqztAYm6LVQ=
user21:{SHA}Oy/RamubYWzVZ1IxxiezV5G8ODc=
user22:{SHA}gN1LmPBMrRxZkzhipSu/vEKgqnY=
user23:{SHA}zMx1exH6oGOtF78Y8ZY7KPp2E5k=
user24:{SHA}5wQp5hYE/N3i/i6Yjz97fBQwh/s=
user25:{SHA}quy6WOPSGssG33/XYI1GjM1EcSw=
user26:{SHA}rmbzU+tedWkSg4z19JgtmpXoO60=
user27:{SHA}lCe4FgTHpkpOhbCp0K5JqA5Ze9Q=
user28:{SHA}AJnAjBWd4U8wze/raohAgH69VJk=
user29:{SHA}CesAlg5KhkapACkKNGlfyDZu1RA=
user30:{SHA}K3p3TiyMMAcQ5vVa+dGNV3x4dXs=
user31:{SHA}w5xupnv3GzEjrZeRTaN/BAIwZG0=
user32:{SHA}5i89qVd8KDVGIj/6WCs42qiLz78=
user33:{SHA}dyfp5ehLdcgD8DV5OTPFMOgapMw=
user34:{SHA}nPBPyIWbALLxg9JZQfsMjS4RcwU=
user35:{SHA}Uk8l/rcR+Og3uJP/+q604N8MY40=
user36:{SHA}hIoJ2Qv4wRE3IOC0T+2PhtuhIRs=
user37:{SHA}zoSpRZj3pbSShQ9O/SQs+72Let8=
user38:{SHA}6LIl/XA9I2XAqwjj2oyYbAGtbOE=
user39:{SHA}+04wmC0awrfmGiRG+HJXjzB3ksw=
user40:{SHA}8ppTyzpxZSWxqsnvxLtFJrvJV0E=
user41:{SHA}7D+UYPYha3+V5/PRn2lmCoLkE3E=
user42:{SHA}DdMdedbyWGN2wXXkFEiaq0aH9r4=
user43:{SHA}apgv+JsqqKiGnwxX93atjb/XRDU=
user44:{SHA}7k3SGIwIJzTmBgFXSlJuO7OvLx8=
user45:{SHA}jTKQDFGg8jCZfA7Z7NTvF9ayFNQ=
user46:{SHA}9Ugo2jB2WaxPh24gPzBGY1zk8XA=
user47:{SHA}VziiTWLQk/9DswnexDFgWnA0i7Q=
user48:{SHA}4DANwobCvUtq/rIntyXbGBRx1Ag=
user49:{SHA}3Hrb7SJv64iKtC6+p/hkSwt7DqI=
user50:{SHA}mhvtao9e4qlisxYyuFvpiDgk9AA=
user51:{SHA}gOUVpYSGfZS2G1e1ETbr9xh3s2w=
user52:{SHA}uwszyh+dmQGG7TTF+og53Ktkoxs=
user53:{SHA}aARgf7klVi/pz8VN86UyQTGzz1Q=
user54:{SHA}zsl2jAqcqruDzXNA8N0mSLUZbzY=
user55:{SHA}UTUesDbSS6/GgPE9iZf3lrCWZTQ=
user56:{SHA}l7my97BzljjfxmBiJckh4YSyb98=
user57:{SHA}BvqybCxPG+FdshJEASmykXCKFJs=
user58:{SHA}8vtQDuXJUEzst96ogzkYvG8uxng=
user59:{SHA}Kdw33gVp/H9JnSZLlWsnPsima8k=
user60:{SHA}pngHweDVwzvWLbkxKI0jpOFMTM8=
user61:{SHA}BVd3k63K6BSYJK1B/hjI8yNXfN0=
user62:{SHA}smmoa8r9gdt2lp7HcRV2K5/IIIc=
user63:{SHA}mQqcr8Htt8OdiJfi6ZATYKrThEY=
user64:{SHA}+r2NvWFQ5X+4vorvec9K2mFDuOE=
user65:{SHA}sEmnRhZegMdBDNnYnyVB47TeLdg=
user66:{SHA}aSdcYgX9Dm3KW/Za/tlT57vxgEs=
user67:{SHA}Mw6kTqFaODDOJj5VpTT2ch/YmeE=
user68:{SHA}VUcPEIUzSnM4hAcg6rxZjmLVDJ0=
user69:{SHA}ZLBhQrXPeeuv1cCPi+7jJPxgXVI=
user70:{SHA}n3eCpjIH9zhOq8ZPlEXJu6p7z+4=
user71:{SHA}rV5kTR4QT6OUb5x34+U4V7KCLaY=
user72:{SHA}7p5fBYB+sDLAAUlXN7mRGpplD68=
user73:{SHA}cTgM/9Qsi2NMwHIVfkRcURDXsdc=
user74:{SHA}yxMuwJUB/2qDA4nryzDaWwDTwkw=
user75:{SHA}eRz741UJnWDv3as1Zr2hXD+V5NY=
user76:{SHA}cgoLdQ4IZQJsWhuAk+JCTQ8t9ro=
user77:{SHA}fySbHcTbxJwfik52qAUS8rI+hzo=
user78:{SHA}lzR3In34RL83gu8qY87y0YHheew=
user79:{SHA}15alyIIv299z3WG9HSMyypwPHrk=
user80:{SHA}d6DZQ826zlJxap75+uEuReJ4jTk=
user81:{SHA}7KJnN2t/mNwfOL+l3iWaxOhxoTM=
user82:{SHA}vM6SgURZ4FWN09Qkhg9GKfYNlvI=
user83:{SHA}Opk05mxS7TatlO37aS/U9IIVQ8g=
user84:{SHA}C9q9Bn45z0QrZmQ39373SmzvDIU=
user85:{SHA}kVVGcThWup+til7v26KfKky9tF0=
user86:{SHA}8vG8fcwUCtnsSxWNTQOZyjmMN9U=
user87:{SHA}rTkBT+vWAX8AkrMCneVIoGWMPIs=
user88:{SHA}4z+IIfZ6CC3UooUH+pntq40RGLU=
user89:{SHA}97ep9vYDA8k3ULenovM1JUP/ZuM=
user90:{SHA}4LkUGAH8SUWCmhpgSR+UoVEuoIU=
user91:{SHA}VMb2w8BIouL+pnejIpju4YvSFrA=
user92:{SHA}TDOwuWHnpr8IWec4G4BV0C1dugY=
user93:{SHA}YU7zF05TgscOFcJyYpoQqYEsYZU=
user94:{SHA}0q/yO/YHEg+nTHaZj9wUMpYXc8U=
user95:{SHA}XRyWxrZ4CYOdZetXUtoBNOYBgkA=
user96:{SHA}B9utQxyZ22KOIV5ucDFQOQBHOmI=
user97:{SHA}pBSxqpqhhW+Gwpq+mPNsstCr8tM=
user98:{SHA}lLeRG7AY/3SvGfvIWXrJLuAFLHo=
user99:{SHA}QzZbE1LP+9XcLOc+2FxJuydCQZY=
user100:{SHA}Ycbe25pEYtUEXhwcEhb4KQjdJmU=
user101:{SHA}bIPx+G5TLj+NPl0MfyVy85moTJ8=
user102:{SHA}ooI9TrVVE7LkLKm5UWwTZT5NDW4=
user103:{SHA}bEIqh9ertcOn4weApT0AoOM04oU=
user104:{SHA}o51ovielNeyLtb8WuWL0r4PNXus=
user105:{SHA}kJDexedKxfHbXeo9MDBUZgagMJA=
user106:{SHA}Q5t2BVGrnc6e0WXfX3dmgPeFdw8=
user107:{SHA}2gUrA8RSDeAsBqTLEjZdaMn+jgo=
user108:{SHA}7/hIGwVORnrIkQ45Kam+1bQAvzg=
user109:{SHA}Ys9f3svSIg9qsti6VK/KO73Whyw=
user110:{SHA}wFrYvAKCXwBKqeLvEN/vpLPDqiY=
user111:{SHA}jCtBfLQYgycvbgwtlsTk0RaBcyk=
user112:{SHA}1ndFd2yaVGQ++IN3o4s7nhmJd6c=
user113:{SHA}3xG++i1ZRHjEAMWIwuLa/lln4Os=
user114:{SHA}lFkJRnfrNTNxWWwLVuS4TDdDT2g=
user115:{SHA}xuNHuDtaTcimJJfoic1iM5V3+74=
user116:{SHA}OZ1xZJTVdEVVKSOYEUWsJInUxoY=
user117:{SHA}LqofIkTgPad+sr0Qfmh5xmd7F7k=
user118:{SHA}tzixFK8WpqP0bM0qS5sLHEurrm8=
user119:{SHA}H3PBIct3rsU9JIPys8flehjPd3A=
user120:{SHA}HINCI2ECClS+2vjbEyE/KBDrz08=
user121:{SHA}T2K08XkQmxOIIFMxN7QN/dHtaXk=
user122:{SHA}mMOZxFazBwasFq0MdAVvMQw5b30=
user123:{SHA}BrQDy1wpNzHm2EYj20X3PWbtGIc=
user124:{SHA}POdeGoaJ4u9dDL7+1aSpKG7u7Jo=
user125:{SHA}J+YDaKhmEbjYz71UaVshKoCxg7c=
user126:{SHA}g9Gphl4ckCSIFal2qP1G+WyIWIs=
user127:{SHA}swnIE5GKm3qcHsCdYAVlR9vfldM=
user128:{SHA}cu8KAzVlpHUeCIDL3dDmhoZwu9k=
user129:{SHA}oDyWybc7jVPlmqA+e+F4t/thfRA=
user130:{SHA}8Y2mGtCTaGznCWYhA/K+XrejBB4=
user131:{SHA}digyAvFugxt+pDECZ/1XMyHysGo=
user132:{SHA}huAd/hD7INmNYaTrOKbS4auA4Yk=
user133:{SHA}JK5NmoqlEpzEtwI8cyILiMh710w=
user134:{SHA}t/RpkDJt44MYi97pB25RABGCTdw=
user135:{SHA}FW+1cwXfHQtLjBjt9EkBNwnuHnY=
user136:{SHA}jTj0t+mHSGoxWFexu7ac+v8nHCQ=
user137:{SHA}cKdPr0oYieZ1qSvGpMk7G6hB7mw=
user138:{SHA}7rCX8u5lLwCcAzDReeSjlU/QkgI=
user139:{SHA}r8kaguSOy20qHsPRkQ2uMzKasM8=
user140:{SHA}V7Hjc5P/zB+sxHcqKPkgyodAFT4=
user141:{SHA}YR60ptREeIpQv0B4cho/t2kzODw=
user142:{SHA}zwfYrgeJ508R2yvS/SdyZPNUuw0=
user143:{SHA}pKvhH74q06+iWSbsIg9jOqLIOuw=
user144:{SHA}LYBr5lIErSckQ/wKA3+YEZin728=
user145:{SHA}eXKk+N0xbstCZ6FV28J/XtVSYiY=
user146:{SHA}k80oUNvldNqQwDvS7n9/U7uqyAA=
user147:{SHA}nOQ2Wjlc5BGlvX5BfN+RMbH+mKg=
user148:{SHA}O6wvUOWKPtROfqtXZ4TMjlJ/Eqw=
user149:{SHA}G1oAkkte9vKp330iZKVnwoyV/FE=
user150:{SHA}6Fys40HEBJYb58Xk4MEa3dYmfh0=
user151:{SHA}V0QFZ76ISexnBTXJNn4egImklsI=
user152:{SHA}O5XnrX44sxaey9uUvGLkJQXthkw=
user153:{SHA}uYqLu0+hCAVauHxHhHj9S8K6dLQ=
user154:{SHA}bNXP45UPcI6hnkXMm9XNE4LncRc=
user155:{SHA}u4yMsvZEFcal38qQSk6pK/YeljY=
user156:{SHA}5xFv7yRZIk605O6D1cn52caIorY=
user157:{SHA}eMiZn8p3ZPwXjrzScxmVZDo5W6s=
user158:{SHA}kjiXjei7Tjszq9RunSYjJ++OUHM=
user159:{SHA}TPcur6LmXg4W+FcHKhC2JVIDnwI=
user160:{SHA}2EckYONFyUcDEqgkPhbfah9cHK8=
user161:{SHA}/3sW+7ineIj1jOU0Dby9nP3ZspA=
user162:{SHA}9mONNF7CArnydjmwkhfQebkGaOo=
user163:{SHA}rGImsJEm0Pd8fbE40OBmP6VGfIY=
user164:{SHA}Mf6CeupM9fas594sId4LX2t4OFg=
user165:{SHA}HxRugr3cGBjQB75CJ9urfWG9YSI=
user166:{SHA}Nh+ikcGOUUgInrJh0iHlVlBmsrw=
user167:{SHA}3BrwZ+qGvcyFB/Lc/XDfOFWwxXQ=
user168:{SHA}8mvvCcUzMnsapzC953/XCxfOT+8=
user169:{SHA}3X/lD6ATcOaRHopH+6m27+LVeIM=
user170:{SHA}Ttiw09bDmyeY/YhY0H/YX6MrHWY=
user171:{SHA}wpd/t9fl65nDq3cSCmRYwv9v65E=
user172:{SHA}pltR0K3LNrS/qOcwZjzTlXpmw+U=
user173:{SHA}mU+OEwW2NBcebLfRPZ9C/Ki0vCI=
user174:{SHA}sOfYGqpAMRqXD4e69F6oH3IrqRc=
user175:{SHA}IjbOYS0WcCo0Op0Sz+ETYSVmHNA=
user176:{SHA}e61WLuSD9xNufpbU1Y8wKlgupXM=
user177:{SHA}3Z41BvXKnswJCMrPj94Xj3Lmd0s=
user178:{SHA}p96d7T+S9V6V8NuYLF679N35qHA=
user179:{SHA}yqC6+2MiJondbEaBtxMBn0X81Zw=
user180:{SHA}cC4wcJtLGRSHKlIqwkr72t9/b9w=
user181:{SHA}hgNZhLC7ZcfxDef0KFz/xEabK+E=
user182:{SHA}ptbHX+v+J4WVRvM9a2tlWTOIjaQ=
user183:{SHA}ls+1/MdSymeKHyVoU6OJ2n9VF2s=
user184:{SHA}oRqSqk2/YvT0NNWY9M9ejSNRWms=
user185:{SHA}mFHUN4xN0yWhXLeHSeRxPFXNIUA=
user186:{SHA}W0GRYxTwoVnywElvbpwdNR0wEmk=
user187:{SHA}CVVbc3vnDGApFxvmTX+/4WUErnE=
user188:{SHA}pKxQI+Q6WmVal5D8XMAy9c3jEyM=
user189:{SHA}C91rFij4vkgcUZisF7TTcpMUFKA=
user190:{SHA}nyzMGhYQfLMBCG7rBtXnbj0bF0Y=
user191:{SHA}hw69n/v0JJf7Z1UK0hn6NcQOiMY=
user192:{SHA}iTz0n/olV3SiPW9zQX2v2/sYZHE=
user193:{SHA}NM/XeJbQkI57FPhUn0pU4pV6tYc=
user194:{SHA}ws12wS7j4dCN+IItHQi8SZS00kc=
user195:{SHA}HhYw12y7WG8xahUNKf4MHlUPitA=
user196:{SHA}KrCi609b3nk3uKVCgw76/cbsAZc=
user197:{SHA}uSyZbnpZu1t1TsvoUyp4TsQvePs=
user198:{SHA}oLBO8TyN9/yjytrlYF6NW0o6He0=
user199:{SHA}axYN5gmBWHFLDMoMf4Oj/t4ND9I=
user200:{SHA}KdFYg6fHUtb8tvbRWak+g+kyk/A=`

var textBcrypt = `user1:$2y$05$fpu.jNd5fPlx3ggfZ2BWR.Wc3/hc7ke7LsIpwZM6/e0B6VniqFRIW
user2:$2y$05$4QmbRfzXERVFyLbUdtCd8ekz1pAfNB5ZsmXevnKgSMc3XHqDYm2wa
user3:$2y$05$.V03HbzL5HAdwq8DYbt/JOVi/crBiqSXvsgNLHucGLLBpApHjK0Di
user4:$2y$05$/jwDvqAoKjNWwRpUzyLvcuhcSloP9tjxAlPfAUlVvVtmMpBPEC9s2
user5:$2y$05$yVjPeTy8/FIUZAJWSSmnAO7GsWHFA2jVeBWFF6Y6RoWEpoxGxtFzS
user6:$2y$05$Uz/8nUK2oWYfFMOoQhLwE.o85Gafb5u7OFEXRNMZkelfsbTtsdi3m
user7:$2y$05$QSeGLMX.BNEWjt8DwcbUYueQFSfVuf0gwcSu/T7zOvpbbRDnwt582
user8:$2y$05$nC1kboym4OaWP8cc3oFEEONouDrId2o40NC/C56c053rOJwGBqdyy
user9:$2y$05$KFX8G2lHpD/LrSTibb/iX.4QX..YIwFj8GnoRWLEwRwa0rZNOnikK
user10:$2y$05$7n20IQEM7RsnxRBn5ds2ReFfkXE0UhkBB5XW.iIA2JLlhZPQTGmOG
user11:$2y$05$EHl3eN3r7/iZKKboeY8y.e9La6ml6JEi8kqL80wOI5f.8/KsT.sDq
user12:$2y$05$Ry11P2IUf7.UGz82PbxGj.oSQaicbeKQPkcvclA5oCb33pkP2Ek4S
user13:$2y$05$Mp5n.vpAZ82i.XVzFDY4Y.dwwrk77L9rk.ZqJ7jYTMLkfuSCohw96
user14:$2y$05$/VPBacHpWoMJaahkq7awD.uarLFUkbNFBfT8wIFrb1zhfg9CpdCum
user15:$2y$05$nB38MWqy2iwUColqIJ.I1OloxpIB3TUNrc4hvufeosSPIMD3qqtpa
user16:$2y$05$5KQoqgOOYc71s4QvTqg8ZO25WO8yDVx3g3Px96JWmu7fc6ocwFCgu
user17:$2y$05$ZZu2PGdDrylM20olrR6j9uQy0555GAzQMPDLZwKxzS4Sh6nasgqA6
user18:$2y$05$YBTVQ1cMDPzI6bHSOTyty.CKWmKMck4Dop01w0HnXO4LmxG503blK
user19:$2y$05$f2yjibPOiBMsNkiXQnRQneMSdvsg7QzCgUkJNL2pAJGeVrJOxFx/S
user20:$2y$05$OXCKpftzXrsbmvgx2zQxlu2EhOpX1KvuVdxhienfxmfY4fZpODjs.
user21:$2y$05$p2hv8h3/12F1VyZPFb.CmO5Hk5lPof/la9fvBwj24sTRBw/3OxNAe
user22:$2y$05$J4Azn9Gi2DoPrYes5G2le.JZzFSJDLCI7RpR8N/uFD9z2A7o8hRkO
user23:$2y$05$Zwd23lXN0R.zEmT9cDGF7.iqxW8NbLVxSrwPzn22N2EZFCjYyq0Ry
user24:$2y$05$VsTetsE8n0hzFj5tORkklu2K3dy2vFWhZz9yBicf36e0p5DbUj6a6
user25:$2y$05$rebFy/Me6yRqeQE0yQiz3uXmGPc16M0rSHVzlv3as.P7WEdrtC1i6
user26:$2y$05$p9RZpnoQb4RqJDTByGjZQupvOeAG63A3gt5OewQWefayCW.aCsx9y
user27:$2y$05$cvqRVup2fO/uygOTy9vTIuWUaBLtT1mXuJrFuJCEjAkjFKo4R8dzC
user28:$2y$05$Q/eCxYpUEH9pTHPGIXIuh..eT1OJ7u44SaQlub1VzM/ZzT4YCLemS
user29:$2y$05$TVU69Aqn3ii9Rr6jKrEq1eMdrkNbdqLwzkBG38UEOoacDmnzWykbu
user30:$2y$05$yRULBrC3eRp.axPNpKXqu.CyaBmXHLKdlLWVs3qEERfs7MaUvLj1W
user31:$2y$05$QG9Cz/BxJMceHX5rh7eqOehwynSjsLvnP7hcx1/CKjbLvKzOSp1AO
user32:$2y$05$oH5mMKeiWfZRMZIooNowI.BA9IynCkq0IXRLzfUNEK2snM8jNFXIK
user33:$2y$05$J3xANq9Au66dzOEzDkbUxuUF/ZnrX51EbkLoaNud6KokVEnrZ1NAm
user34:$2y$05$hyXHnEENS2CXFuRna7UtpeCI2BvkcjqvxDRI1SHylVoagNTLgyHHO
user35:$2y$05$PquIRYSdFTsD/iksRadiy.PHerfiURxHnXx.3WqSCuX1PRZm79LKa
user36:$2y$05$8YhuTDyWXi.BTjXBD9bp9..Pakffo/6Ms/1FYHs1tX9IGXoofFfm.
user37:$2y$05$AHw4z2ftObVNtAqY2ifDQ.SapTREZqpPd3.otexEd7nWwpx4G3WrC
user38:$2y$05$EQMKPMeRKfvweG3LLIjf0.fU.QJIfdckuKx06DCEETGBq7az5MvQ2
user39:$2y$05$7MoTa.ShO5OvRuMHbBaDkuPgF6vCDrm4ZNs3JixBomyxTAEcI/OAK
user40:$2y$05$S96TthOpyBSih2M2/s/Si.S0R/K18Z9WrnU.1Jcvshf7IwJNQFSAm
user41:$2y$05$ofV8eWu/gq7Cl5jeM/g8gOjxagMPBgafEcLokUQjFVjMPOtUeYb7W
user42:$2y$05$oqpzIgFFdmZZclNRzCTctenR3593tQeIB2.Kng2n.77mxAZVofDvC
user43:$2y$05$G4r6paxV3Vnb2TznoUnXUubqluay82LplRZcRfFPsgbrGEP1VvS16
user44:$2y$05$d2J4qIkZwLt.DMuEZnztcu/rza9sfvWJ4x5orOYTl/4e1sbX7AxLO
user45:$2y$05$GvzkFUAI6eVVpqAPSmgy/OyS8WQAzIEGJL1U3nRJGFgUD8LGY9rAC
user46:$2y$05$0VhRbZS7CqiUTDZVwvF.ce9h5WXc2wmVcTyDTUZhVfNB5Ksu71BYC
user47:$2y$05$0Aa6On.Cg8u36U/.GxZVI.hDTxsijKfpfGTQbPL0Ayj7zq6Ve5zVS
user48:$2y$05$Ia.EF4y53dlX098YHRODmO3shsjof0Yg5t11V2JOReJV/EbHNSzNu
user49:$2y$05$UmLUkDgUtFLZ/JSoVTbzWu7cAMwEQ.S5qZU9Tb.y35Jnh0vgNNIta
user50:$2y$05$JgjmiFUGIcr7IWusuoO36egNqkCfihmrzvkYeiZznBDoOesOOK93e
user51:$2y$05$VReKSINlgML6BOqS8jvFjuIj1sS/4mTTgCZXlqX1BW8GVZ9qiKNCm
user52:$2y$05$ow.TPBduyugHa3whAX90ouL05LtRjCN9Ero/2W.qRVXbcDFt.xgSK
user53:$2y$05$wuBbsrneEjs.HloFV0jrF.J4E311zZS2OpUN46Hs8r.jv8rxxBm72
user54:$2y$05$k022XhSDlXPzG91vP6wJk.U.nos5dCfr.nnyxkFkaCPphyYRy8wqG
user55:$2y$05$2.7736YD25ngfQjTw3f.I.dBeXeuelhD.6pDzSs624d9Rhn5wm6zW
user56:$2y$05$vs4S..Iw8DZac0gPbVFlROt9rMZTwQkBDBVeafBPC3QZ.PQ2yrcYy
user57:$2y$05$3U8kq/6QwtAMX8gQcnqGfuDLa9T2.LjXeJasgB6jxEqeejBoolHRK
user58:$2y$05$IQfXyv6F3oOCx0CJTmm.seXeDTjsy2A4dpm8nu4IfcsOGEfulidbm
user59:$2y$05$I711r.W404fDV.BFMf1lcOW1Ui0t.VHP944JBtgHOZLFumPWFhLd6
user60:$2y$05$Y7589tyt.8DCM./Qd7zNke3lqrE2/NVOIwYHUWndeG4b8H/RxfOgm
user61:$2y$05$Ppig7aA.K7QNRFzJSLqtsOMyMOsZ1q2zX2ylaiBHMw7OuBymV2FmC
user62:$2y$05$h5ywSTTkq/OO.qmZCGLO9.u/DpWgtfPmQV9iWuKolim3xxg1FBxe2
user63:$2y$05$2J3ilSS54gUgWd2QHa3Za.AJ2I3U9qwa4Yt5fCOZXUtZC64c6CE1q
user64:$2y$05$hArtMCE/kwCP4A8qpUk7C.dUpBN64BEWh4S3JLVwk7juqyi/TbiQu
user65:$2y$05$rjaN9UndVJkCn28n/IxFC.Z9ZZaXNyKta6IOi.YMfIQneIjRYW1JS
user66:$2y$05$RvjpQJ0c7fr1oqhidGnZhul4xZ4Wz5OOxyoKZS9zIJd04PFxOKKXK
user67:$2y$05$KW7YOYlGCj.xILSDc2qj1ukAbLP8T1j1C8CdvV/fubqojMHuUROQq
user68:$2y$05$yoNxBJYVpko7Tdl4wlUiiO4gaPISGiJ2hYMEGdR.iJ3EcfKLHT7Jm
user69:$2y$05$IRQm9evZEMNLncGSRd0aH.Qsb21GH0B9wvjqiZ3vwxenyM0lwVM1y
user70:$2y$05$FyjSXqjUxZwXF4R5m6c4qOyFKqTVwBv/0ZeRN0jDaUvufm9FL/Lk.
user71:$2y$05$7rNZytgyH2HlBGsF16rZjOMHMELgT05vUidO2WGJTU11Xgmgcf4Vu
user72:$2y$05$emeKtLWD/6e72aIkBMHRROCqx21mN.Gm1x.5EPFn5Lbyms9BfH.0O
user73:$2y$05$R8uipqZiBjQz9fLfskIIT.9KZcPmr8qHY4AEkD14vyB2jXuQmyJPa
user74:$2y$05$T2pEQR1szpQ47GbwDaA5qullwCIeDq1BG1a2ceYBSdAKXnPH7/doe
user75:$2y$05$KF3l/wQFIlMBzieNfFogs.nk4YC8Vw0loswEo2DEE9Syi6WAvf.S6
user76:$2y$05$Nml8AAAQKado2b863fFxR.S9XCFuLuytdIrsIvL/Br5A5mrjJsTBa
user77:$2y$05$ip4kF7zLCcJD9sE7Q9ef7On2AE4R4W8tvanyFEXMG.WLqFGdntJ.e
user78:$2y$05$2ovSeGnlvBUJqcLlus8spex40G2dHaUAN/QarQtEUNn1nbSj/gAs.
user79:$2y$05$p7I29qcE.eQPBecOCDe7keix3IIAbzqUu3Ir4s.g1Oiz57IclsIJy
user80:$2y$05$JQoJ3yuvIlAX/IZBWD8juOdcHFpToKj5nXvNi8u2R07rbdMVTGBzW
user81:$2y$05$npr1gqlBk5MfC1xxMa/81.Fvz1dQClMf5fkX1760ZwBsk9Z1C6JyK
user82:$2y$05$PUxmaztoq6LEvEH7Z1lVhupdFDn8QAC9kZJZhn7jWvjNH3uVlsWIi
user83:$2y$05$JCFzyR2crPQWxAUGbhOmxuCRvPtWh9DBifl6BLqxjtIgYA6t3.Y7.
user84:$2y$05$Zlm7cBEHc82Cr45U6s4mF.yns7SIbQfzvud8zLn4JIzuJrfIN5uD.
user85:$2y$05$vnXd3b7mwQBHvhM03QyROus/aIvNWzlMbTbd2Scaq0jqz8wixqzSC
user86:$2y$05$iZeCnIlKhbGPM1CyNPhPBuSA2N.HfdQ3PptdHxMiV27ztq2mqYWj6
user87:$2y$05$lrZkYmIAS3wi/7qoYDVZf.BGqYjJi5HrHGYJeNCYNrddTBoZbKwbK
user88:$2y$05$xYkdPF.werZJ/BlsCu4zouTSRGlMmx5ZVATJRcFCpi.BwLdIRICMK
user89:$2y$05$cPWffLKZediydMyTmyfWbunBqJKrbiRbg8SZVOGfoGlHII6o3IkcO
user90:$2y$05$8rvpHcG22H4suih5Fg6mHerW8wxYyIgB43ZDq.OMWVxOYfo/nyquu
user91:$2y$05$MLrLmkwjOMAYEtWnCw5TDeWY4Ue8K7pKQK/ZJJZSMKGq6kxl9tLXG
user92:$2y$05$cn0765.8hUJO/1MogfCR7Ol/frYuteDprFpIgJl9J9QQGIYRnjMXq
user93:$2y$05$/aZy1Cry/kN5CMzA96JViOLEwXd7u1Aadmm3iPr5QNuaM72bOjpui
user94:$2y$05$iN2UCtmqV7klZdZOHb.R7e5TYKijn5NDkE.os5hvemoA3nwQoeJT.
user95:$2y$05$NjE.eJP0.jWObEol7gNBYeVQFi1/aovArVQ7x9DeVqmwxWBqt5o5i
user96:$2y$05$41JNXQpdNJPkiOBM3F1PXeTlu/T.hgI1bZEQhe/pf0eCXDIpXS7ei
user97:$2y$05$OAxCZPaPALCAfBQvoSb3B.udhGRpmH7RdCrGKZTXPqXQAry4v6Qfm
user98:$2y$05$zUykSYNsiLVEB.f52zRC.e8sXXIkQPy5Shj/4mPPyYwsODX6GuS/e
user99:$2y$05$51OkkDCIFjrvG/bGBDNo0.bFRx4RzQbnNh5AnN0oNnbp06ueJGSdK
user100:$2y$05$TqRBmY9bVNdwoviMt.2o/eNxvvwIAYx1szbqtlp9sF5FxK8aAsSfW
user101:$2y$05$C5Gfb24e1Vt2p8He.I1pA.vukWSdP8bP8y2wHFxPc9b1YgHLuNGje
user102:$2y$05$RDnVwwP1Lw8zc2qoFb201ONE8RXGhAdcStEhpcoEtD9zjclCCGU3i
user103:$2y$05$2zABVzL9tnfu8Bap4d886OoWLeW7Qu9gkxDT5YvFc8iHW4xexNuEu
user104:$2y$05$vh5bmiqTDrSJ4.FcnPGuLup25v80rTU.KSAfpLCyf3WMQc/mmHZiW
user105:$2y$05$vWB0WGXgQeohXNXwtpp0B.yZ1hU82QWI041sbuzlJTzDqn6dG79y6
user106:$2y$05$WyevYNKgnxiNNMttQOAjVuFuO94BVvlAOqLVoky1IxMrjsgeHhvJK
user107:$2y$05$wxSeMKJpmDZtgLqfnQ09QOT177hEG7yWF.6Gkm9hjj/CoWA7Rg57S
user108:$2y$05$o8M8DdHSULopvLbdOxHgq.8wWfzSFyVIuKg0rPTSd.Kekhh5QPiFK
user109:$2y$05$elje9vjVayPOY9K5WE6WsOVkSB82RQy2ScVbgMCZaZA8DrBKWfjiC
user110:$2y$05$1YFeINEvVqmcHhnpoE7Cx.LAgAML/5UHhdnwPcIueRQkTGcG8Naji
user111:$2y$05$4PAzaOnaEcVCTt3PeJIEsuqjtkV.sqOGQmCyLCH7dxUOoKxvzYtz6
user112:$2y$05$.PgLHC9D6Lx87FD4zMtj8eYE.RXsXqIfBJy4pvUJ3s6F3E55HABGa
user113:$2y$05$/Zl25uha8SuQed5cpnJgB..fy3tGi1b9BrDRfwOFwFe6Xm99ywVwO
user114:$2y$05$vxsYCWHExy97KIr0uqlb/.zQldEGJ9b7tUwxGRdlBnuSK2Ch/9CSi
user115:$2y$05$IiUlNTxdbyATh79DTVkKOeom.mEpCoolj4Hze7HuXE9iyierxW0wu
user116:$2y$05$28uYLgKrXKMoXD//2/DuC.yWORZ7bmvGe/edVT68wGRF5g2tEY/46
user117:$2y$05$Xqw0nSoJcwGOhV/ysuEKG.0MP6PL5MAypXTvlER5hQcN706ci8SNO
user118:$2y$05$7LV5GmyW6OsZO9/2WIWvW.PP6Nudxke6irVZpJ/1XBsn86Np0ya62
user119:$2y$05$fsiHzJCe920m5s7rXm8XdOaXm0//n.B2IzkoP4yMwnWhkyiALfvcO
user120:$2y$05$IX.gWZo.Wqjh9Iovo5sweuPM0ajoEf.RMhYkZyn2JL2lkGi07ft9.
user121:$2y$05$2FYLncq4yV55YOpLP/wJdeoVGAa3BjDR7krPE097eSG8gde/eC0R6
user122:$2y$05$0Nxu3s2mO5vA.BNc0fxvB.KkgXEjFUQjcHvwhg3ZYF.6xJqeBZVKi
user123:$2y$05$aXGniNsIWKP/T67g6Pcx0udcqhNI60a80f.WureeWkoANKt5FgW3G
user124:$2y$05$n.phXNC4bIjaXOwgnIlJ..D2AVGyC5kF7XMdvQxiXntJlHd5vJ1.a
user125:$2y$05$e08FDibQZ0etegypI1VbnesPMDSq2J30usqlHWfvFXgPBcxCXAvRW
user126:$2y$05$XR2gWLZyqll7CLoBBVgcquGJa7daRfuZYu6UmFHR.lm/AH.zacwCO
user127:$2y$05$TBMOzQV3M7oylv3vEEIFZuMUnE8VLPbWCVj2E6ZVTYX0WR3cgP44i
user128:$2y$05$QG16ACPoPgoWCM4x.OVkMuRgJFWmG3ju48kgqxTi3jKXOa2JrU7jO
user129:$2y$05$2zvNfbv3OtAfpserCAHuEOjV5VnjteIpFQL0E4YdDhDjwFyo3BXbu
user130:$2y$05$J71u6qftS8rgZaf2ZbsnHO3fZspP.hxoNwqORBTLCdcUs.O/cXtLm
user131:$2y$05$g986dBMr.1xh7o5CkGIpsOZFzEQXKa8sZ/1zPSNIOLu9xtigzb9Am
user132:$2y$05$Zin4JporBoAOj4IcI68ZmeDrEm40eKq2bBpHx70pL5V8N90wSlocG
user133:$2y$05$MJ2Y1Zyq4KPwcubkZ4OkHOYVu/WHE2I3yS4zeDMVrKw6QScGa5p.O
user134:$2y$05$7D69czCOO.ue7FOqOItLOOkAV4fD53LO560Vh3ZXFdeZI9ztFiAI2
user135:$2y$05$flASl2xIJbu1S6eh6qKsU.8evfNyRhFnD4YHOb3jQG7IkU81IHeWa
user136:$2y$05$HYwFWzdVdy32MFc4925GluQgvFeYYuFFo57sdW2JfXRNjvVUdoz8e
user137:$2y$05$bc123OMzliw3Fs7tKAG52.dpLrXq/pLVa51zzxNTC6aoT2lB0DJrK
user138:$2y$05$lzROdSwfnB1eL.KRcnoxSeHw9rvC3Yzedlapp4Linl5S4Nqp62jnS
user139:$2y$05$tqpPCbpoNRPuvOD7Ft3TsOUSMjkBVCvuoNgJ7H6befM6hSFHXO8Ee
user140:$2y$05$Idrf81gkXnukRSBMoYUCwua/qzATCR6Nvcd2Q7vQjdgPVn/aA27He
user141:$2y$05$ivAiXWgWojZGYZxzQNExi.6rEGnFs8B9LVDly8jbS4N9AK3CvjHlW
user142:$2y$05$aL2JG7sPcfpSxZuRnMVJwuAFmPyiXQRrHgyUBgsTjKSLf018ErI/W
user143:$2y$05$Kqs5Yevpa1uxfCF.LiqIbenCVOeH5jIKG.FnxmnGyBVrDnD.dhRGa
user144:$2y$05$x18HY2fc5kpkMAGC0c17/.7fjI/z4YxsBuR.04UDZIlOJAK6U0yqq
user145:$2y$05$40qblQzCoH7Iuo8kQnkPx.1J0SES.pmtJsLxaHgMDdDe4//iSJU5u
user146:$2y$05$unAahznVcDZpP0hVqAeda.Peqrlrv2deezG3fgo4I6AsBq91J8cda
user147:$2y$05$o6c0UELvAMeu4d60YzffrOL4A4aXjaj1WuppDOhVmWxJScAqHA602
user148:$2y$05$PA4qcZuBBvDmkoy2nltcR.ondB20vQkpuP1n734QSi8uP64Z4z78.
user149:$2y$05$VV5oiqyBLpgFA8Mx3Q73Iu9m14eqecHgLZNXpF70L.IB/1dkxUNQ2
user150:$2y$05$XMZ3A3J09fSaDQVG3nBZReWI.lhvywnY1MoQJeWY7gLxj0NnkV7sO
user151:$2y$05$Foegw9N2dp3KzovQf.VhzuNwKlM596h8E3jooZlvK4sydyq/wXOS6
user152:$2y$05$USV/siqS6k4A2EVZCjItSOOVkiCGJjUTFP5GWuwpEyoSYcOWqtFTS
user153:$2y$05$a6Bme5bWs2jZ5miW/OwWTOeEMnufQke5toRgI3UeylZEh.cfF87uK
user154:$2y$05$slN6WWU1G/OY5U2QmvTW9uBc.e5EUTHZhJVBSqjYwxovxszF1L9fe
user155:$2y$05$wMe2..gX8czpGmZ1w5x1P.UhslzONBdj2ctmeIWb8ouketL5i4RM6
user156:$2y$05$HsbkKhLMPS/4wBJrL429quKQeF/Qdcf/rAZWg9xzz9e1FaRiILTia
user157:$2y$05$ZZKt/HNdOin.hroEcUN16eORw/MZVIkEB9LnI5XZn1yEa2rkwRNCe
user158:$2y$05$XJAZY9u2MoaGjPx/ZXgbjOpIZDxQpKEDZ/Md665Nx1ccuuMPZQaii
user159:$2y$05$j5z75UZHASfKp.NTop0pMeqEVCXyozMm36TJpyOpdtfdvSe.UMhmy
user160:$2y$05$q2ovZcoy..57uh6g2Akw/u4R4DEPU8HZKbObd6UV5w6r/ShNRgzk6
user161:$2y$05$5kkpP0chaCA0Dws11TsgC.qJvU.6kQnTzZgbWgonn6u/QpRYUkYoG
user162:$2y$05$wXfSl2ej.X5ob39Nf6h9Oeqsu.3f4FBi96rd/zmG.LWt3UY/rrHSO
user163:$2y$05$PoRCSHwtuRaIDAPrl4IFu.XQ4BZo82qSrRQ6dfj.1sBDIoh325XBi
user164:$2y$05$VCdqzvFQNYVz351tBHTqWe5MGtqeQ8W6mX0BdN5ejYmDwQDkpO/wa
user165:$2y$05$1ObuLcibu0t5IYEqr59fD.ZGTl71sokNBbe05YesJ1WmrDQEikeBy
user166:$2y$05$VzJrKDUQDwwdLQQy2tjQreGvzJHlsuA.EArBc8E9IDdj7L0849DKa
user167:$2y$05$NZPa7SOEWodCYfSXAqYk0unt4X77OHufus4l1HOF5qlWI0Jy1GlLe
user168:$2y$05$HFlvCvLD/F256wCMZUBlnuopUOUCis0V8rcW5oa2MksRO2sQ60k9O
user169:$2y$05$C6E2pBOz2IbvINEcnD3eZuoihoZVGzRnJ4bkfvW0Z56sPraSXqHtO
user170:$2y$05$AN05MYUjH.uCfnWMcu.qSOZG9ZorKROPlOk7IjgNXKv13peLz1/TG
user171:$2y$05$V4tACKJFPT4j2myrLna7Q.8Wo0rzr3vlCWIouRgz6AKxZ9D09IY1i
user172:$2y$05$u2vl6vGV9M3QGGlOynsuI.eCWOgj.NfEK3a/TBupNGLS/t0fbMUZ2
user173:$2y$05$.gcHT5HClXOfMpvYxjktW.z7dF0OedXuF0kbjl/Z9hq8G6LvMeM56
user174:$2y$05$CvKw6kGjjmz1hr1ek6VvZOi.yEzc7yah5PWmeup.H5fwto2bgKSbu
user175:$2y$05$5JS2jsi.FTkbhMFD2OghbObXKq0e8VlU0k03OQqqXygWCP5Pc3ydu
user176:$2y$05$Otk4kuVJg4N6Sk.ppwVEl.RLzUr/fFdPOhQ68jJVzdZkxE4mTmeH6
user177:$2y$05$ol9zBezfUmheimpVvYAGDuGVK7/q5Q1NMSF1/ph3iA1kn/MJzhv5q
user178:$2y$05$j2ASG7XBQks27NaXQ9UPkOvQwQT1cJks3gDm3c3E6pvxBU2RC/rdG
user179:$2y$05$PVDmQKnaOI//Q22C0A88H.IgwTdSwWBKaacJ91mPy/9l/.t.o17Yy
user180:$2y$05$TEh236aWbGFowb3dDLynZO5x00duGa3xDlNYz3zPTSe/ICp1Led96
user181:$2y$05$idwthhR8KpqzZ.eI.JemmuikpYwjE6rp6gIgPbPbztzWFpiqrrE3W
user182:$2y$05$bgv7RlB1KwdDc82YrY7JqOrPRWh0F5yVg67Xq5epbk1tA4hZM0uJ2
user183:$2y$05$xpv9YpSwREc14RguzkaSquJtQZDZA9mqlp3L.iZEnPbCR41Wjn9r.
user184:$2y$05$XhDVxSxipFBkP00Xdjd8cuBBLbdXIAs8M91NUt29ryvB6ttI1snEe
user185:$2y$05$t9er2Dy5F6zdk4fgXw.lFuJF1X32udmV9aSx7pps3oY9sfKgcFsmm
user186:$2y$05$gDr.uxsBEulmKBIzGdkmMO18YYYz4tpmn6sUofvpek/Hqa4lNbQcS
user187:$2y$05$oInWaewy8ndw/qXBYVEp3e9Wv12uXG6tlYulvPPtlxyCQIJ1w9GdC
user188:$2y$05$16bz5v17uC39Koc.rV97R.AmwRjhOGhCceEtx.1uFtFLFA3XDsY1u
user189:$2y$05$AP55Jluw4VQ3hW5UkrHjz./3iECPnoMFjLGelK0rFtkgkrdflF0T.
user190:$2y$05$ZsDnjCAxhgevU.J5SZKc.eF1lQIJ3gcL.n5IwgOVOFcq5QZfzjmkW
user191:$2y$05$pSMMwCuW/8cUNiOF7LqBDO32bJQE.gdAznWiCfkyYls8LI5wyJIeK
user192:$2y$05$pLa34zaUzmy4DJQ97Qi2Mu/79OnttIYtsU9O09a/qRK4QAVNADB6q
user193:$2y$05$uWKyW.hPmXgVXImS53u8ie.RYxui6so1Frk06kg6zL4.RGDcoff5u
user194:$2y$05$Y9V2kSmgnoXHCoYLxIDzGOq47Hr3frXcwHhYJy9dz/0FYfE4Uhpfq
user195:$2y$05$pA3YGc/uz0Z0wtE84BOn9OhkhqqZ9eYD4jJTuJ7hQPvjEj1J1keuK
user196:$2y$05$Mr/Wy6f850nVpk6DqBwnXefHGX8VbCbNoRzsEOeERkOHCUkxVptKO
user197:$2y$05$0wcbcdEB8/ibQsNnf4euHuBEr.sb9ELtNM/cNddIdCMyq40WgU/la
user198:$2y$05$qt4IdndlY3VjMhi2CkmUUOWhCrkUEi72.DrPIhZqynnlo4AM.M8Ou
user199:$2y$05$wXHemFAraZVila6SRTYCluqzby34ZlAMnTuC76f9ZAKHPpOV.vdYe
user200:$2y$05$m2/Q/xtOckFnDGWbPFfh4.iTVNQlmx0MZP5MURF7adM8qqtBVgC3u`

var textSsha = `user1:{SSHA}KHQzbbDgqjRkfd7li1NBL7kI0D5oMzBM
user2:{SSHA}OFxiAyw1TSNiyGybLyjVg+yewdhoMzBM
user3:{SSHA}FeUKUVnpp9IlolmuqfIMUYaa3/doMzBM
user4:{SSHA}mfDeU9QRfvED1gfBExrJgDsi74xoMzBM
user5:{SSHA}Y3eY1xbgHUOFKOPoiLwluYlsd3FoMzBM
user6:{SSHA}8YIr07hiZ/xCLZuv+tl/PamyNzloMzBM
user7:{SSHA}Lv4FvMGR79gPrNXb6GM0rsYr1kFoMzBM
user8:{SSHA}eI/7z5Fbz6fS55+uY4mNXDyFzVhoMzBM
user9:{SSHA}HhVE6yZC4cGYZJrjp8bNizVNkvtoMzBM
user10:{SSHA}jPcUc3pza/FTKrZwNWY1pXogrc1oMzBM
user11:{SSHA}YWLQZmFWYYRPa8d52wlNqHce+31oMzBM
user12:{SSHA}+QLebXR4nY1xqIYuN5fLMuj/OcFoMzBM
user13:{SSHA}E9DwGVpfo1Wb8pg8FGIsBLMEeYtoMzBM
user14:{SSHA}bA6g5CWRkOa2+HAdqIP53tXPrItoMzBM
user15:{SSHA}Hu10u9ozPmwdJUM9ebqDoZstZ3VoMzBM
user16:{SSHA}ea3lxQVUtjFrm/eTgkWvOJKTVT9oMzBM
user17:{SSHA}aYJ8z6N2No+WSqoJmMpCw5qRk3toMzBM
user18:{SSHA}kvYKi+yiReC4aSfym87aJsj95TRoMzBM
user19:{SSHA}SHVoHbUcvmmXbLc7k4k6O3AZOuZoMzBM
user20:{SSHA}pmu8UohVZRq9mvn6pyFv4jgmQkRoMzBM
user21:{SSHA}7+/sVrBRXm9XZbzhk55JUASDUYNoMzBM
user22:{SSHA}Fp1p+z7cuVujAaohJkgrmnLf9ktoMzBM
user23:{SSHA}U9dfaNzdeJ1u8zKniQ6Z5Lp2779oMzBM
user24:{SSHA}i0/C67599TUT9tUmJgAyFtX9xx9oMzBM
user25:{SSHA}y0Qvlh7EmJBa4m2dnfNEx/j54+ZoMzBM
user26:{SSHA}O80eZv+Dwv9Nb2J0IW84PKm+6aloMzBM
user27:{SSHA}aPYp8DM0dOv4gmtpJgd8BMQBK8xoMzBM
user28:{SSHA}C9srG/+T6iwcBN20hjMiUnaqePhoMzBM
user29:{SSHA}d53nozT1OQQa+atT7SbO4Fj+Cl9oMzBM
user30:{SSHA}6hjz/GEbliQkPC3orkmyD6fPMhxoMzBM
user31:{SSHA}4Rb7U5/frazhx06b3co8j5vqSZZoMzBM
user32:{SSHA}EisC3/gKFV3MGKJ6igxAB7PWldZoMzBM
user33:{SSHA}MtIpYndBk+ncjTvVjpjrx7hw1uBoMzBM
user34:{SSHA}k+vyIDFCycpT3V0ohh/QyraHOFVoMzBM
user35:{SSHA}HPQI9htwR+vaZwyMGd6AE7P7vZdoMzBM
user36:{SSHA}V2jfn+UdRewTpv0lTHi0YsUI7dxoMzBM
user37:{SSHA}nC8ZWb4NJBFXaJBJ2SELqt9NjItoMzBM
user38:{SSHA}I2BhFvLcOHu3kpJ4ArSAL8r+R4NoMzBM
user39:{SSHA}krXqIY0wTLcGf8wZcPTRtAWApYxoMzBM
user40:{SSHA}4l4OVj6kaQ3dSjJspefFtkLs98doMzBM
user41:{SSHA}w7UKTyr7fYAv9EP5BdF5MwzhmOFoMzBM
user42:{SSHA}RS7w5ZTbFZ4cK8WSX98yTG++NVtoMzBM
user43:{SSHA}3IV3vZpo/3Bfyn6MZvxbf2wF+yRoMzBM
user44:{SSHA}bmWZRmCB2pf0lAXTKgZiYVec96hoMzBM
user45:{SSHA}hlLhgGiThrJDeYTiSQK/wCzJqVloMzBM
user46:{SSHA}po82UhfKIkGikuNPovsI72sbgntoMzBM
user47:{SSHA}hetH7xSE8xI7zLsBuEO9ufALkE5oMzBM
user48:{SSHA}FV4R3Z4Uf8znN8C2eaFi6E7QchZoMzBM
user49:{SSHA}ca/9mXQqLEoVRSnzQR9vc/fcVKFoMzBM
user50:{SSHA}TppiH0dqE4KbmOOsCWbsJJowHcFoMzBM
user51:{SSHA}wyzYVvmUtlVYDYl6JicC5RotqsloMzBM
user52:{SSHA}Y+cjO4EZgMvWKAo+WgyAX2Q8SExoMzBM
user53:{SSHA}gHwpYMGY8pH6wEwRVw/Cm5w2lHNoMzBM
user54:{SSHA}poxdGbfnmdKvpLcQt71kZNH/D7poMzBM
user55:{SSHA}khBK9AqcZTTHHf4zLGducfz5TLxoMzBM
user56:{SSHA}b1AhE9NpyPkp2xT3yzX+86SMwmpoMzBM
user57:{SSHA}1Tq/UtIv8HVRiy3SlLYOKSpwQUFoMzBM
user58:{SSHA}5ZbyAJuunxfFxoiBJuuS3ZU2V0VoMzBM
user59:{SSHA}qAvT9MYSoXwHsO14aZxxP9CJd3poMzBM
user60:{SSHA}/0xhnl4FEcqWxVbXltMflVMFpRVoMzBM
user61:{SSHA}Xy0VSSr2Xx9FFj0s/GiGVOwcwm9oMzBM
user62:{SSHA}FnkOooLy9YklOp0WHTJRmXcTv6xoMzBM
user63:{SSHA}aBYmNsX0QeeMmCtuRq0aIJzVB0loMzBM
user64:{SSHA}kSZWJj9SJcWW2opW69EPjFfjoI9oMzBM
user65:{SSHA}8w8T0Tx4ocX7JxhgeukJPhKGB6loMzBM
user66:{SSHA}vX9Crz2xv1l4Kza/R+/VESELpoFoMzBM
user67:{SSHA}pv0YVoC9d992+Siu9VegZol7vUFoMzBM
user68:{SSHA}tXggGK/I9Bvk68PJmLOjF4RqiwhoMzBM
user69:{SSHA}GnWz4il+uLx0e/JnNklEj3ZwJeBoMzBM
user70:{SSHA}BoEW8ObIfMxBCJT7Ok3bi8+ZSyhoMzBM
user71:{SSHA}RU+1lSihMmimuwMTDNvv1fCeHCJoMzBM
user72:{SSHA}Hiw4EAZ2goue4XaSFuyXdqqfxM1oMzBM
user73:{SSHA}2tLAOEJTpsWsN2JUwKU/mRXxP6xoMzBM
user74:{SSHA}BWQhdxiP2xUJXIjJU3LeZR3V9xhoMzBM
user75:{SSHA}dCbng3IitpL7A0zPBXBcrbtBjA5oMzBM
user76:{SSHA}htqbeTdR8MBN0CmHWEOqslvZRKhoMzBM
user77:{SSHA}baNLStg1x+cmgzCe7SDJ6rpk4vVoMzBM
user78:{SSHA}yKnoi2OHHgBWw15tMDblv6xzeSpoMzBM
user79:{SSHA}Le/e17YF4WJPnJ0YIVUuzRMZaLRoMzBM
user80:{SSHA}b8eOOUjZ6jdMKkGbKlfZw6lQPIhoMzBM
user81:{SSHA}JbCEc1aY2BzLkmCD+/qM0jvEYe1oMzBM
user82:{SSHA}BNOxlQQkbBEy1HjrWxKar7QsdUpoMzBM
user83:{SSHA}eeLU7ug80NHe8LfPY+oRQLAvdwRoMzBM
user84:{SSHA}irfuv5ljgC4QOt1R3sOoXAEyhy1oMzBM
user85:{SSHA}YRDBXhlaWsA0CBRt+myBzpCj2aZoMzBM
user86:{SSHA}bZ22uNBHwnPcMyXCDbuz3WiZI4toMzBM
user87:{SSHA}egkSIVwc59/vfe0j1LZ2HJTcpc1oMzBM
user88:{SSHA}KBCe6qUEmAf8DW0a7gCAdgDSsqpoMzBM
user89:{SSHA}t4tKWxCQnHBOKeyXOef1mBTXbMFoMzBM
user90:{SSHA}Ypn8CHHGEp/UxeDe0+VEhRAVmwNoMzBM
user91:{SSHA}Rmqn/fp3ZYI6pFQf1kHZnMAAgJhoMzBM
user92:{SSHA}LuQObAdJBOOwJZaj6gjWKDGPXw5oMzBM
user93:{SSHA}eBWmJJcITESmsI0KvqhVxyGWzHNoMzBM
user94:{SSHA}JMomFFPTD5mKAMX5RESy5Gzn2L9oMzBM
user95:{SSHA}2IDJMQu6/qH9vx69lbz1/CAr7E1oMzBM
user96:{SSHA}0USP7ALYOJbSOA+8dQsPcEB2UAloMzBM
user97:{SSHA}1kFW8+ZOQaKBYgSjr/+Xg5gSeo1oMzBM
user98:{SSHA}82K3okDKowglUX07wHUVze194NtoMzBM
user99:{SSHA}lTGN0Tzz+IMY4SJ9Qrqxa2vVfgdoMzBM
user100:{SSHA}RQ7u4Q8nsyIP6rTQIwJzVs9BsMVoMzBM
user101:{SSHA}4TXcpqrGY1cFv8NCXz7MyTI/08FoMzBM
user102:{SSHA}mjPo5L+y8LsZa9UM6ZAjU1jVsoBoMzBM
user103:{SSHA}BdnavrD58G4dxZe1T4UEJgmrCMdoMzBM
user104:{SSHA}1JS0MKf3wbTJO+0NAX95IM+gLoBoMzBM
user105:{SSHA}hxXVWQY9olJjJrQDUQs1ZvO2mQxoMzBM
user106:{SSHA}04r3p/bEDY1RFFy4HzViPRwMY8FoMzBM
user107:{SSHA}ueaYWKFrV5MkkDKz6o0OxDTD2ploMzBM
user108:{SSHA}Kk5ZWv4cG1/7OkQxbfeX0NC7ZixoMzBM
user109:{SSHA}jz9UsGxB+WWSa5t/+yO3vEXxTsZoMzBM
user110:{SSHA}7+mH+eOgnHYRdj3sgCtMZXXbuI5oMzBM
user111:{SSHA}S/bxP+7k1MFoOtnMI/+jRVEZEHdoMzBM
user112:{SSHA}DP3/KiVPZUeRKHLqd615V9zheYdoMzBM
user113:{SSHA}xSBpP/Ks2bHkkNw43ZoWgZwyLXxoMzBM
user114:{SSHA}Jelh8enXbfVgowFrGAezz7eRU6VoMzBM
user115:{SSHA}kh/CqFaDa8/qZRefquWGcndqkPtoMzBM
user116:{SSHA}VI2vj1Up9E4fueC5rBdnqRBxQrVoMzBM
user117:{SSHA}phVsSEcVMiO8zmw+nA6G9fWS55FoMzBM
user118:{SSHA}PeEC0PzZzpQG0yZI0GirnjzukEVoMzBM
user119:{SSHA}fOZPdW6w9c2S18275BNDDVsftGFoMzBM
user120:{SSHA}3Jrv1Fbonf8IZTNdQJ/B1QLkmvxoMzBM
user121:{SSHA}tCwozhADmF/uW5bgPNL1VsiCuQZoMzBM
user122:{SSHA}8pUP8G/zhGTtbGRO1kOR+7d1TUloMzBM
user123:{SSHA}d/GUjH4rTUr0DZBwUBTJKr7EtCdoMzBM
user124:{SSHA}1FVlOEVaFsiF0Z9DS8trJf8JSnFoMzBM
user125:{SSHA}N9yPyRsGOAudEKRb3BWDoceOVbtoMzBM
user126:{SSHA}5RcrJM9HGxqOgmaLH55GBILygapoMzBM
user127:{SSHA}J5mAnDnDUHs6c5P0owInZTSUye9oMzBM
user128:{SSHA}l0ghIM+c2Pzw8A6yTOv3JyGjeRBoMzBM
user129:{SSHA}w9csbivGGj2ljyQKyb29udb3HhJoMzBM
user130:{SSHA}rB4TRiWPoN3lpRLb4MLpIZjMWvBoMzBM
user131:{SSHA}yqEnK3v8994I5+c367UdJE8HfyNoMzBM
user132:{SSHA}qJc/rZQnzYNlMiadFNhWvOHBvINoMzBM
user133:{SSHA}PFploR84iEbdtF/x+dwWwpRPokxoMzBM
user134:{SSHA}5Zt7YJh5RZ+ok0k9tIWpe7byQeloMzBM
user135:{SSHA}O/kFdl4oTuR7b/hW6v9nBAXZJq5oMzBM
user136:{SSHA}Z0/jL7eP4UCCz4Nor5Zwv3/7rGVoMzBM
user137:{SSHA}s3q0nzSa5Kwr4uria6hPkJc+fxpoMzBM
user138:{SSHA}99o7YV6tFOjsaSJ5Zp61qY40YZ1oMzBM
user139:{SSHA}DJ9376eJUCmxD4qCgkEeGpL6exFoMzBM
user140:{SSHA}TiaqzWkzs0yIMVx1jPHfNcC1y59oMzBM
user141:{SSHA}Iov2I0xZjPDldvg3TyLsbVzu6RFoMzBM
user142:{SSHA}SicIuNHFniKXqatyD5BqmXQla9doMzBM
user143:{SSHA}rWqfhuQ4lqvppOtzH85mk23wjZ5oMzBM
user144:{SSHA}N4k4oB7kXsgBTY5CzozHC6UQcR5oMzBM
user145:{SSHA}byB/hau+e8lzLKH4hJ7v0AC9FG1oMzBM
user146:{SSHA}+yHFBPw6dMQNc8nG2Cf/3CwsLcxoMzBM
user147:{SSHA}PSsVEdn71YdUbTfeNZUdWqZOqjFoMzBM
user148:{SSHA}dIN7+mJ98Mx3tTN4pUytF/bUyVdoMzBM
user149:{SSHA}Lpf8jkaR6O6UX9Eo7/6adnlplY1oMzBM
user150:{SSHA}lSEACsysI/QDB1flO17SfIqLRTNoMzBM
user151:{SSHA}PyiqiXVdIAUU2vWuWmYluC7/1oloMzBM
user152:{SSHA}XnoIQTiT7iX4p/LiYaI8/GeidW5oMzBM
user153:{SSHA}eN4eA+XbcgUJT9yE/0LEfHBCbkJoMzBM
user154:{SSHA}29JQsa2RAHBzBWvTQYj6xPnmfqNoMzBM
user155:{SSHA}GKVmTEZjUx9yb9awOj/ytwlJpDZoMzBM
user156:{SSHA}CQ6EWkVYt2xcPxVTwRG3HIV86lFoMzBM
user157:{SSHA}odr4eK00XOv39LonI8c2kIZ0AHtoMzBM
user158:{SSHA}O6SjDJZHsZlK98P2dlpxEvffEPZoMzBM
user159:{SSHA}3tSTU4yAUjnaD+Yc9N4FGu1myBFoMzBM
user160:{SSHA}kIxaiwca+r740x+MAg1wUeuho4doMzBM
user161:{SSHA}lq55WI87UsAETQgRsTq85FJ/+1toMzBM
user162:{SSHA}M9/mVMsfq9mQp6CR6csTNJ1/OiRoMzBM
user163:{SSHA}WysDSR2VVHt3q3RaVwTLTu1T0RBoMzBM
user164:{SSHA}r5Ld2ldElAxot9JhUlicrL4cooNoMzBM
user165:{SSHA}c7YhRwWq1uUgMV88KPLUTfRHEqhoMzBM
user166:{SSHA}Qya9WpbcXm4CGe+74snT6+wNvuRoMzBM
user167:{SSHA}LY1sCSerjc/3yibSTtWoKfaa8YpoMzBM
user168:{SSHA}GX0QU5/+f0Ay5qztDkNH3uxzk6RoMzBM
user169:{SSHA}SGVZtFmKPsx6kYD5aBIpKYieFxloMzBM
user170:{SSHA}Cy0ONo8OzwKV/ARm/84tQp0Y6U9oMzBM
user171:{SSHA}rcRRLAU6TK3T1Jsl76YX6VOv/MhoMzBM
user172:{SSHA}5w2j1pCMcEKJh+MzK+d2rTQMiYtoMzBM
user173:{SSHA}8Pidv4ozJZfpN4eBFfGPuiPgEM5oMzBM
user174:{SSHA}+BARKI8oCv8CCMVDGC4SPN6gATdoMzBM
user175:{SSHA}rTeutNoSgHrCpcaJaIgrikrZjTtoMzBM
user176:{SSHA}25OEOVO23FKntU5+RDo9S6Dkr9JoMzBM
user177:{SSHA}MKLCUfJGoZBKu9krvcTseI0goLZoMzBM
user178:{SSHA}plzb6TMfairoikuiP+KK+BBPm3toMzBM
user179:{SSHA}yZ20fANjvHSNsxWdkG8K56DKc/hoMzBM
user180:{SSHA}I4jSmOBFOZIY2UHnRw62pYNnBMNoMzBM
user181:{SSHA}RDvaroyOH+U7n2Q64DQNAN2FStdoMzBM
user182:{SSHA}aJ+PiZ43i2GyVORiyTc8qmqehVJoMzBM
user183:{SSHA}CWum4640+nzpvXZsQ1ehiwvc3/RoMzBM
user184:{SSHA}zq+mHI0UFb801Xqi09kBtnTymJNoMzBM
user185:{SSHA}wYm1qGO5B0zQX8Y3WQ/3QEYeLfVoMzBM
user186:{SSHA}e9vu5c+YM4Kve+LoZuVXizjnA35oMzBM
user187:{SSHA}edOCLF3ylXjmaclBdWmuuRV4IeNoMzBM
user188:{SSHA}axoiuTXPOrinEInJ+xJZh2wRr25oMzBM
user189:{SSHA}HvXH6BqE4cItgD5vUNA6c/C0lRZoMzBM
user190:{SSHA}0qpTjPnBN4/8rv48WF1sRv9/WgNoMzBM
user191:{SSHA}49TJjIbq8751PplEYqGmiCF/SOFoMzBM
user192:{SSHA}4mN9pnH3mNN1tw5SbRR60ZW4e1poMzBM
user193:{SSHA}vha97emDO5P6FQI5LHLzRpphNhpoMzBM
user194:{SSHA}95fd56CSiwAX1tXtFzKh1NR1y2VoMzBM
user195:{SSHA}Afyedh7hf2PnO+sByzce5QdXsaJoMzBM
user196:{SSHA}YrrIbz4n//wQpycwYWWvdaXsmSFoMzBM
user197:{SSHA}lO9vZbWoLYBXqg7jEKfmWszliFpoMzBM
user198:{SSHA}Mu4MUUt+JNiV+wC0zqnxjO0Lb0hoMzBM
user199:{SSHA}e+9w+4XNZlPmc0/c2lyLaIbCWPBoMzBM
user200:{SSHA}y7fnpcvJRDlQAoU/0zgFLNqr5OxoMzBM`

var textMd5Crypt = `
user1:$1$D89ubl/e$dJ8XW4DfrJHTrnwCdx3Ji1
user2:$1$D89ubl/e$xuQ74IxhM3J10sv0QHVgA/
user3:$1$D89ubl/e$Y07COBJSUbNDlYlFyRYUp.
user4:$1$D89ubl/e$4IZ.tBiqvtxt7Dpt1MkgE1
user5:$1$D89ubl/e$mLrBtDw8UTdAX7jDZLQIB0
user6:$1$D89ubl/e$gDHt/53o.SrKMB2Ts06ll1
user7:$1$D89ubl/e$rRYNMP4siiirAmmukKbLH1
user8:$1$D89ubl/e$eZSBBncGvTB0M0FC5y25f/
user9:$1$D89ubl/e$ulbSP3eTh0fe.Xi.1yZUK0
user10:$1$D89ubl/e$.eWY6zxlyQZR2H/Wl1oBv.
user11:$1$D89ubl/e$z92umtR06bhr7asKE0qys/
user12:$1$D89ubl/e$TAcr1BfieDgQrINZM5ob4/
user13:$1$D89ubl/e$/hZIUw5diSqk2M17H8ya2/
user14:$1$D89ubl/e$sQaL5tfFD85YivTet25eA0
user15:$1$D89ubl/e$d/DZx6g1NFNtzU4.7a8zz1
user16:$1$D89ubl/e$Ss.xxvcZk.hyZhfw.uHu./
user17:$1$D89ubl/e$0P5vwLE5BrTDStMNUSVvy.
user18:$1$D89ubl/e$6.cSxEJwebnwRhE3Hzn0p0
user19:$1$D89ubl/e$C/xkXc8Y1I58TITNT5B000
user20:$1$D89ubl/e$C793xtKVyHtLgE9kv9F3G/
user21:$1$D89ubl/e$l77x/.bF/SKs9j1fzZsxT0
user22:$1$D89ubl/e$Kyx0nL8mqLVQSTEoLozXf0
user23:$1$D89ubl/e$/K9HdFFaZs6fQtHeg50i2/
user24:$1$D89ubl/e$5KXGSkIs4eVf9.vv784h2/
user25:$1$D89ubl/e$Unn6gxvier6BbNqXXggcC1
user26:$1$D89ubl/e$y6uoa9UFNh89NqUYl52rU1
user27:$1$D89ubl/e$/hj9e6a9Ka4A6PLtscUAJ/
user28:$1$D89ubl/e$FaTM181xgGpilXHhzKspp/
user29:$1$D89ubl/e$juw3fg2cPTz96Styto/mD/
user30:$1$D89ubl/e$VPgIKROD.HnIac7efgPqp/
user31:$1$D89ubl/e$YHiBHNKujH.hOoduR6yI30
user32:$1$D89ubl/e$EZatY2FiaT38R7pJM28Ta1
user33:$1$D89ubl/e$nXBo5NIufmx/azSj5c7xn0
user34:$1$D89ubl/e$4c4jEBluS1Z/5gwpP5IHP1
user35:$1$D89ubl/e$OTaO6MXXx/9Kzgdax4IFT1
user36:$1$D89ubl/e$fyLMarn1NGkBlrVvnOkWe/
user37:$1$D89ubl/e$BSEg4wMYeZZrqQRbgLZQ40
user38:$1$D89ubl/e$09Q0RgEq0luH52Q5D1r5h0
user39:$1$D89ubl/e$0CFkGtb.nFnwY0AsvIZ/p1
user40:$1$D89ubl/e$mxD6Tg8iiXGn2iuap9dkZ1
user41:$1$D89ubl/e$jTHXwK8wWOmCWMhV/Nqq6.
user42:$1$D89ubl/e$eQQM1N3kzqYaJsIHAuma3.
user43:$1$D89ubl/e$VPIMSuAdaNExEPg1o6BR21
user44:$1$D89ubl/e$tMKU2NslMUsgBVMJ4Z8Bw1
user45:$1$D89ubl/e$Zo9sQwDO9imDRzQjAvSit.
user46:$1$D89ubl/e$BnGW9hkQkwO/Fpj6lLGKK/
user47:$1$D89ubl/e$Gc5RbOS3wD6GwQ8rbSJDB.
user48:$1$D89ubl/e$5lIEMaiA6epkGKq3ZJJr./
user49:$1$D89ubl/e$tFt1GFVE6wxs5UhiRIPJo0
user50:$1$D89ubl/e$WNTQoeaBUI6P1ypLfFBGz1
user51:$1$D89ubl/e$U/zmE2HZ9arX1CFysF48F0
user52:$1$D89ubl/e$p.bG9zn.rvJB7E2nGhleK0
user53:$1$D89ubl/e$zal5AqxTqbMoo36DXwEQi0
user54:$1$D89ubl/e$s8mCg6jFym06sqZWoOXZr/
user55:$1$D89ubl/e$80YBEXI4zCLlO9bOld9ey/
user56:$1$D89ubl/e$coCMG.asGxvMHowtOtD/p.
user57:$1$D89ubl/e$ZSf.3PkORJgZTiW3WmC0S0
user58:$1$D89ubl/e$Q8Ti7BlgMjNwXCjzsvGz0.
user59:$1$D89ubl/e$GplTzL8mgIki4Grxkmsnn0
user60:$1$D89ubl/e$SMrOpSVsLdDZh0fgSv3RQ.
user61:$1$D89ubl/e$L7c4tHvipozBibsgDlIae.
user62:$1$D89ubl/e$0kSaA27U/3XSYr8ysXYk00
user63:$1$D89ubl/e$h0QjZLtAyLFIS0RHzUC0g0
user64:$1$D89ubl/e$m/HGZlMaYri6dyglbebGq1
user65:$1$D89ubl/e$ZCeDrrlRD4z0GuWUhkRIp/
user66:$1$D89ubl/e$mrunNmbLqlK6GBjHAyXme0
user67:$1$D89ubl/e$LAleKOBmjHO.JkMqxZJva0
user68:$1$D89ubl/e$PsZG8hCSDwlu61oK/xVwY/
user69:$1$D89ubl/e$FI4Kdq4kuDrE.02FKhu7l.
user70:$1$D89ubl/e$vM6r7Y64Zfo7Tw0YkPT.T.
user71:$1$D89ubl/e$AiLLO.VR/Z6HBLIsg59yg/
user72:$1$D89ubl/e$NZ7SaJ8LroRQacSqrP7fp0
user73:$1$D89ubl/e$gfKqHDjTyhd5guuVyAg/i/
user74:$1$D89ubl/e$Z45Dy1hE/XhDlnLJh63w80
user75:$1$D89ubl/e$MmNfgZpCBgohV/RiCqEeh0
user76:$1$D89ubl/e$yx7VuwWX2kJ3fgelcamZp/
user77:$1$D89ubl/e$qSMTwpxirNVKQmnv0LKGt/
user78:$1$D89ubl/e$i1mU8nOLP85sWe.XN35Cw/
user79:$1$D89ubl/e$/MSgRW0uYjW5NdqyBJdPB.
user80:$1$D89ubl/e$4pXO/YHKqJ51rx2dWolr.1
user81:$1$D89ubl/e$xLJthQG.pQzgK8oVOGzNB1
user82:$1$D89ubl/e$l1Q0BEAEDlXPW44kj5xhR0
user83:$1$D89ubl/e$qnjT/g8iIRqgJjEdutYMI/
user84:$1$D89ubl/e$5zxhC1KC7L1bfJPiZRzXp1
user85:$1$D89ubl/e$IaTFU11RpmJaS9o4H6Zbd0
user86:$1$D89ubl/e$pi5D8IwHVYgPyiXuIuW/Z.
user87:$1$D89ubl/e$mQukEUOQ3Y67xULFLxenN1
user88:$1$D89ubl/e$pwTwGBzTj/f4qRvrXWDeH1
user89:$1$D89ubl/e$eSRqyebfyJIf.wGwNWJiF/
user90:$1$D89ubl/e$aaKVVJe1BPZbXaKqWDhJH.
user91:$1$D89ubl/e$myjNwn1J/Xur2sGPb2KDW/
user92:$1$D89ubl/e$Sk9rTdAjux9EEWjIi983s.
user93:$1$D89ubl/e$49aMem462jQkqqZ1YCYtu1
user94:$1$D89ubl/e$SRje2pCzH8NjNInSkec5p/
user95:$1$D89ubl/e$WYHvQD3.AucnP.fo2kgLP/
user96:$1$D89ubl/e$MMK0xP5K1qHKsFsERRYCR.
user97:$1$D89ubl/e$sVVqstvqjaibnGEvVOYci1
user98:$1$D89ubl/e$VFIhSi8R9JmFPoZA4qDHI1
user99:$1$D89ubl/e$BlA2Z9tvvxDmV2dhtW14I0
user100:$1$D89ubl/e$52FunD0ymhST8S2IxFDwB.
user101:$1$D89ubl/e$3GLf6CNVmDH48iwOIVdYI/
user102:$1$D89ubl/e$d2eAdEkBFPJ/CBOM.MHET0
user103:$1$D89ubl/e$TbhdN5Ec9YD8bplz1A/ce/
user104:$1$D89ubl/e$ZCOXuzvrKNSVFOIlo8NWa0
user105:$1$D89ubl/e$5XqFN00zJn8eZ6mxIlNDn.
user106:$1$D89ubl/e$Scb3wMWrVNXDw0FtBU2Bx.
user107:$1$D89ubl/e$NKpZ4pHNMRByBtpN43qaR/
user108:$1$D89ubl/e$z3O5D9UC4h8DqIND7HI8e0
user109:$1$D89ubl/e$ex2DzBcd0OK0xjWCk9gXn/
user110:$1$D89ubl/e$nU.oBq2dK8u1/WYTb.qN./
user111:$1$D89ubl/e$0DiN4Qgg78qXCKmCgn1v.0
user112:$1$D89ubl/e$0M452BZ4RsUFJlWmAzC/r0
user113:$1$D89ubl/e$kon/OQAJaw1jWzUVGF696/
user114:$1$D89ubl/e$rffe2sXyO8D4i8b/Zz4fN.
user115:$1$D89ubl/e$hZxrfcVVudoH/vlLjHbO/0
user116:$1$D89ubl/e$RXh0ZI2skEE34Okz7ZFkS.
user117:$1$D89ubl/e$VdO1d5wKNz.082DfQBKdw/
user118:$1$D89ubl/e$0O7aLsmdpnkds69gC8F3D0
user119:$1$D89ubl/e$HWdm5snaFBk8FiDKaaIEQ1
user120:$1$D89ubl/e$xqw/ieFswTtoXsUl.qFiE/
user121:$1$D89ubl/e$SzKFu0pIWn3aYOU4n8vD80
user122:$1$D89ubl/e$lEOrH/zkoSs9gEPhWFzPY0
user123:$1$D89ubl/e$12DkA0BGmzbnFiXKQXOoO/
user124:$1$D89ubl/e$B4vLgohydtKjuwc3PDmLJ.
user125:$1$D89ubl/e$jANv18j9YRx4dSRG3e9dR0
user126:$1$D89ubl/e$5V/VlZE2LSVL/FXTfXSEY/
user127:$1$D89ubl/e$iFyY3.Bcpd2jcEQLV5wSf.
user128:$1$D89ubl/e$DHPXon3QnoN6NBaTv2Mg5/
user129:$1$D89ubl/e$DsPKwJLYCgtQGgiV1vsJW1
user130:$1$D89ubl/e$6DGzvs.WKNeB.s4iVzHL..
user131:$1$D89ubl/e$FwjNgb634KqRDONd9mWn10
user132:$1$D89ubl/e$njQdR835SwHOdC9yOv7t3/
user133:$1$D89ubl/e$0wHxUj/ReVsKDnUnPR/Su1
user134:$1$D89ubl/e$WAD9PbSIhz206tHvVsLIX0
user135:$1$D89ubl/e$LDUiW43DNTTs3M.00nnTs1
user136:$1$D89ubl/e$PHYo0bAH5mMVAyjIC8piI0
user137:$1$D89ubl/e$UyR/Od.i9HQ.st4.tpEI.0
user138:$1$D89ubl/e$5UD3WvZ/FwruX9YFVaIUT/
user139:$1$D89ubl/e$48eSiBnmw5e9w5giOn3ye1
user140:$1$D89ubl/e$4X3PcQgOz0JVEvHTuFRh9.
user141:$1$D89ubl/e$I1uauxPkvEIyk8UVXsDAQ.
user142:$1$D89ubl/e$Ya0d5paah4ZnA8MbLTvqo/
user143:$1$D89ubl/e$1kY4BuTtTHHAEqBpDnh2B.
user144:$1$D89ubl/e$cfbSrTlsCwlAD3l7NWcjg.
user145:$1$D89ubl/e$gKcTYvUdSIfR9/wxxVKh00
user146:$1$D89ubl/e$dr3IRAptFJZJihDhfJM6L1
user147:$1$D89ubl/e$W7HymXzTiMw19qnRxMFQz.
user148:$1$D89ubl/e$IDzuywgAHby2yNESi/cEu/
user149:$1$D89ubl/e$Fj8gIA/n0wzQQeFow.C7X1
user150:$1$D89ubl/e$bbWZK6Y64mVxUFx/KiX1B1
user151:$1$D89ubl/e$p/oVop6W/YiSDKNOOKi6D1
user152:$1$D89ubl/e$Qo8SMcE2S9QFp/zsRZeNG0
user153:$1$D89ubl/e$q8ae/TEYgiU80vniBKIqo.
user154:$1$D89ubl/e$gYAm1Win2xeG8VSEwezd30
user155:$1$D89ubl/e$SFZgFYUGzipDnBs9gWwbr/
user156:$1$D89ubl/e$J0Amxd2tBuLkMwRfT.9.l.
user157:$1$D89ubl/e$s13zGQ31soCmTTRI6dtuO0
user158:$1$D89ubl/e$sROc1K.SSkTKzz17qcjkR.
user159:$1$D89ubl/e$RyRsl5a2AMQomSbXxJwKd.
user160:$1$D89ubl/e$UPDj0.FApBTp/Mf8Pa5tG.
user161:$1$D89ubl/e$yDOxOyZ8Na.utOWkk19s01
user162:$1$D89ubl/e$2upo6jW4F5OLgMvJv3PuU1
user163:$1$D89ubl/e$JpaGvclLj2SHhVAFBOqtT0
user164:$1$D89ubl/e$XeFap1wOn3vppPM5HYTVp0
user165:$1$D89ubl/e$hRcYGGfn3HcDZcxeZtL.P1
user166:$1$D89ubl/e$4t.pGGDRnT.epiS4XemUZ0
user167:$1$D89ubl/e$MYyxqzymgUVhXkjLDPoeq1
user168:$1$D89ubl/e$Oes0cPGc7hQF.2nQgHThF0
user169:$1$D89ubl/e$laB87gR5mmN/aNaQW37.E1
user170:$1$D89ubl/e$wvY4dNj.Q3U7pnEQgf/je1
user171:$1$D89ubl/e$jRk3i8CF1C7Sh43qytD.31
user172:$1$D89ubl/e$0JCRZtaPgouKCrcAiF3290
user173:$1$D89ubl/e$iZjmNH7GBYKN19Yl5IWF9.
user174:$1$D89ubl/e$QwDfjzDOqFgbsP52awbrP.
user175:$1$D89ubl/e$gwfYHJCUhkaGXuCTv9WYO1
user176:$1$D89ubl/e$w98gHUCztuXegDue.AOTV.
user177:$1$D89ubl/e$ROjdv8BY.9e3naiqMgt94/
user178:$1$D89ubl/e$u0zmiD8YSiNfBZMgjhUkw/
user179:$1$D89ubl/e$FvH6jFxbT3BNYSJLNEkRE.
user180:$1$D89ubl/e$50h2tbqHW9vagG7wL9YP71
user181:$1$D89ubl/e$BUr/iUuPZAJMn3LC7qzkq/
user182:$1$D89ubl/e$70s6Is6uEwtFldtmotDhG.
user183:$1$D89ubl/e$k7PCEkvp2Pdm3HWWOBgmE0
user184:$1$D89ubl/e$Cknm.cI/57fbdNnztPqVv/
user185:$1$D89ubl/e$z1CwDc9dhBLK/3nPyrjN4/
user186:$1$D89ubl/e$vA4kvwltG58Twovn9ia./1
user187:$1$D89ubl/e$E75YhMUix/GZ6sGKMzKcH.
user188:$1$D89ubl/e$7Vjp3/aQgDkvrvihk.v020
user189:$1$D89ubl/e$ZrqfkcX5deYA89WLAiDsW.
user190:$1$D89ubl/e$s7khoFO35J03/IWR0Au9a/
user191:$1$D89ubl/e$oWHcMVYWZe.30EjvdVhSz/
user192:$1$D89ubl/e$TqtgsDnP6LyrpDriHI3g1.
user193:$1$D89ubl/e$BBVL9.H9p0kabiBuuGCP5.
user194:$1$D89ubl/e$nGYG2N2nmEZA7Xv0gpsUa.
user195:$1$D89ubl/e$g.ihTAfr9HICtitiMed6U/
user196:$1$D89ubl/e$BJ7gzGwNQmCy8WLTAzEcG0
user197:$1$D89ubl/e$oroQkvFnt80PX9.DFORsH0
user198:$1$D89ubl/e$SdHoMvPduS1kS3KVqEw9W.
user199:$1$D89ubl/e$1FQtoOElFQQCBL53IT2LL0
user200:$1$D89ubl/e$xO3.z/20nsNEXnaWJdsfB/`

func testSystemReader(t *testing.T, name string, contents string) {
	r := strings.NewReader(contents)

	htp, err := NewFromReader(r, DefaultSystems, nil)
	if err != nil {
		t.Fatalf("Failed to read htpasswd reader")
	}

	for _, u := range testUsers {
		if good := htp.Match(u.username, u.password); !good {
			t.Errorf("%s user %s, password %s failed to authenticate: %t", name, u.username, u.password, good)
		}

		notPass := u.password + "not"
		if bad := htp.Match(u.username, notPass); bad {
			t.Errorf("%s user %s, password %s erroneously authenticated: %t", name, u.username, notPass, bad)
		}
	}
}

func testSystem(t *testing.T, name string, contents string) {
	f, err := ioutil.TempFile("", "??")
	if err != nil {
		t.Fatalf("Failed to make temp file: %s", err.Error())
	}
	defer os.Remove(f.Name())

	if _, err := f.WriteString(contents); err != nil {
		t.Fatalf("Failed to write temporary file: %s", err.Error())
	}
	if err := f.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %s", err.Error())
	}

	htp, err := New(f.Name(), DefaultSystems, nil)
	if err != nil {
		t.Fatalf("Failed to read htpasswd file")
	}

	for _, u := range testUsers {

		if good := htp.Match(u.username, u.password); !good {
			t.Errorf("%s user %s, password %s failed to authenticate: %t", name, u.username, u.password, good)
		}

		notPass := u.password + "not"
		if bad := htp.Match(u.username, notPass); bad {
			t.Errorf("%s user %s, password %s erroneously authenticated: %t", name, u.username, notPass, bad)
		}
	}
}

func Test_PlainReader(t *testing.T) { testSystemReader(t, "plain", textPlain) }
func Test_PlainFile(t *testing.T)   { testSystem(t, "plain", textPlain) }

func Test_ShaReader(t *testing.T) { testSystemReader(t, "sha", textSha) }
func Test_ShaFile(t *testing.T)   { testSystem(t, "sha", textSha) }

func Test_Apr1Reader(t *testing.T) { testSystemReader(t, "md5", textApr1) }
func Test_Apr1File(t *testing.T)   { testSystem(t, "md5", textApr1) }

func Test_Md5Reader(t *testing.T) { testSystemReader(t, "md5", textMd5Crypt) }
func Test_Md5File(t *testing.T)   { testSystem(t, "md5", textMd5Crypt) }

func Test_BcryptReader(t *testing.T) { testSystemReader(t, "bcrypt", textBcrypt) }
func Test_BcryptFile(t *testing.T)   { testSystem(t, "bcrypt", textBcrypt) }

func Test_SshaReader(t *testing.T) { testSystemReader(t, "ssha", textSsha) }
func Test_SshaFile(t *testing.T)   { testSystem(t, "ssha", textSsha) }
