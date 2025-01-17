# 3a6ccb03c434f719c04ccfcc2c9ce35013424b90 (2021-05-21)

Don't reduce modulo m unnecessarily.

```
[safenum] → go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        7046594               184.3 ns/op
BenchmarkModAddBig-4                      984085              1531 ns/op
BenchmarkLargeModAddBig-4                 363100              2769 ns/op
BenchmarkMulBig-4                         388844              2784 ns/op
BenchmarkModMulBig-4                      314078              3786 ns/op
BenchmarkLargeModMulBig-4                 226860              6252 ns/op
BenchmarkModBig-4                        1158972               986.0 ns/op
BenchmarkLargeModBig-4                    415216              2623 ns/op
BenchmarkModInverseBig-4                  832084              1403 ns/op
BenchmarkLargeModInverseBig-4             289676              4534 ns/op
BenchmarkExpBig-4                           7690            140706 ns/op
BenchmarkLargeExpBig-4                       105          11499733 ns/op
BenchmarkSetBytesBig-4                   5206584               234.5 ns/op
BenchmarkAddNat-4                        5770597               183.7 ns/op
BenchmarkModAddNat-4                    11950340               103.5 ns/op
BenchmarkLargeModAddNat-4                2921572               406.5 ns/op
BenchmarkMulNat-4                         137380              8128 ns/op
BenchmarkModMulNat-4                     2682211               450.0 ns/op
BenchmarkLargeModMulNat-4                  49249             25011 ns/op
BenchmarkModNat-4                          54798             20895 ns/op
BenchmarkLargeModNat-4                     63405             19135 ns/op
BenchmarkModInverseNat-4                  172198              5939 ns/op
BenchmarkLargeModInverseNat-4                636           1880771 ns/op
BenchmarkModInverseEvenNat-4                 157           7773234 ns/op
BenchmarkLargeModInverseEvenNat-4            159           7483443 ns/op
BenchmarkExpNat-4                         387169              3243 ns/op
BenchmarkLargeExpNat-4                        57          21061245 ns/op
BenchmarkSetBytesNat-4                    785596              1796 ns/op
PASS
ok      github.com/cronokirby/safenum   43.576s

[safenum] → go test -bench=. -tags math_big_pure_go
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        5793404               214.7 ns/op
BenchmarkModAddBig-4                      917923              1171 ns/op
BenchmarkLargeModAddBig-4                 170028              7783 ns/op
BenchmarkMulBig-4                         197192              5113 ns/op
BenchmarkModMulBig-4                      219387              5886 ns/op
BenchmarkLargeModMulBig-4                 101758             11690 ns/op
BenchmarkModBig-4                        1164991               931.1 ns/op
BenchmarkLargeModBig-4                    154022              6841 ns/op
BenchmarkModInverseBig-4                  746452              1447 ns/op
BenchmarkLargeModInverseBig-4             137312              8781 ns/op
BenchmarkExpBig-4                           8520            138144 ns/op
BenchmarkLargeExpBig-4                        60          19625440 ns/op
BenchmarkSetBytesBig-4                   6022134               200.6 ns/op
BenchmarkAddNat-4                        4837562               221.7 ns/op
BenchmarkModAddNat-4                    12762744                95.37 ns/op
BenchmarkLargeModAddNat-4                2299520               455.9 ns/op
BenchmarkMulNat-4                          77550             15464 ns/op
BenchmarkModMulNat-4                     2675755               453.0 ns/op
BenchmarkLargeModMulNat-4                  44835             25523 ns/op
BenchmarkModNat-4                          56588             21005 ns/op
BenchmarkLargeModNat-4                     59696             20052 ns/op
BenchmarkModInverseNat-4                  206808              5787 ns/op
BenchmarkLargeModInverseNat-4                427           2663714 ns/op
BenchmarkModInverseEvenNat-4                  93          11266719 ns/op
BenchmarkLargeModInverseEvenNat-4            100          11274026 ns/op
BenchmarkExpNat-4                         436466              2709 ns/op
BenchmarkLargeExpNat-4                        54          22100573 ns/op
BenchmarkSetBytesNat-4                    775238              1528 ns/op
PASS
ok      github.com/cronokirby/safenum   37.968s
```

# 036c5a428268cec0fd63887d75a4fe53c50a350b (2021-05-21)

Benchmarks using 2048 bit modulus:

```
[safenum] → go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        7074402               164.0 ns/op
BenchmarkModAddBig-4                     1025569              1105 ns/op
BenchmarkLargeModAddBig-4                 449142              2932 ns/op
BenchmarkMulBig-4                         429507              2975 ns/op
BenchmarkModMulBig-4                      228830              4452 ns/op
BenchmarkLargeModMulBig-4                 221468              5562 ns/op
BenchmarkModBig-4                        1210222               957.8 ns/op
BenchmarkLargeModBig-4                    444782              2693 ns/op
BenchmarkModInverseBig-4                  836184              1478 ns/op
BenchmarkLargeModInverseBig-4             295900              4466 ns/op
BenchmarkExpBig-4                           8374            140578 ns/op
BenchmarkLargeExpBig-4                        98          11188430 ns/op
BenchmarkSetBytesBig-4                   5052667               232.9 ns/op
BenchmarkAddNat-4                        6121905               208.5 ns/op
BenchmarkModAddNat-4                       53536             21612 ns/op
BenchmarkLargeModAddNat-4                  61086             19195 ns/op
BenchmarkMulNat-4                         134721              8297 ns/op
BenchmarkModMulNat-4                       27606             43755 ns/op
BenchmarkLargeModMulNat-4                  19029             62319 ns/op
BenchmarkModNat-4                          57270             21360 ns/op
BenchmarkLargeModNat-4                     61735             19143 ns/op
BenchmarkModInverseNat-4                   43052             27536 ns/op
BenchmarkLargeModInverseNat-4                610           1939277 ns/op
BenchmarkModInverseEvenNat-4                 151           7830059 ns/op
BenchmarkLargeModInverseEvenNat-4            148           7759581 ns/op
BenchmarkExpNat-4                           5539            196737 ns/op
BenchmarkLargeExpNat-4                        27          42940694 ns/op
BenchmarkSetBytesNat-4                    738734              1484 ns/op
PASS
ok      github.com/cronokirby/safenum   45.238s

[safenum] → go test -bench=. -tags math_big_pure_go
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        4995013               255.2 ns/op
BenchmarkModAddBig-4                      882956              1198 ns/op
BenchmarkLargeModAddBig-4                 175794              7042 ns/op
BenchmarkMulBig-4                         260041              4783 ns/op
BenchmarkModMulBig-4                      204600              5758 ns/op
BenchmarkLargeModMulBig-4                 101822             11598 ns/op
BenchmarkModBig-4                        1265289               960.7 ns/op
BenchmarkLargeModBig-4                    178405              6966 ns/op
BenchmarkModInverseBig-4                  755259              1506 ns/op
BenchmarkLargeModInverseBig-4             132333              8900 ns/op
BenchmarkExpBig-4                           8341            141174 ns/op
BenchmarkLargeExpBig-4                        57          20066576 ns/op
BenchmarkSetBytesBig-4                   4863494               275.2 ns/op
BenchmarkAddNat-4                        4241242               335.4 ns/op
BenchmarkModAddNat-4                       49152             22361 ns/op
BenchmarkLargeModAddNat-4                  59218             21233 ns/op
BenchmarkMulNat-4                          76137             16409 ns/op
BenchmarkModMulNat-4                       26935             44656 ns/op
BenchmarkLargeModMulNat-4                  18156             66906 ns/op
BenchmarkModNat-4                          54295             21300 ns/op
BenchmarkLargeModNat-4                     60390             20335 ns/op
BenchmarkModInverseNat-4                   44236             26966 ns/op
BenchmarkLargeModInverseNat-4                435           2753945 ns/op
BenchmarkModInverseEvenNat-4                  99          11446948 ns/op
BenchmarkLargeModInverseEvenNat-4            103          11365429 ns/op
BenchmarkExpNat-4                           7198            165886 ns/op
BenchmarkLargeExpNat-4                        26          44880442 ns/op
BenchmarkSetBytesNat-4                    801037              1479 ns/op
PASS
ok      github.com/cronokirby/safenum   42.140s
```

# 92874261776e63721c16f50b476383a6c2a1818b (2021-04-10)

Various small improvements, namely free limb injection for even modular inversion.

```
[safenum] → go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        7626368               174.2 ns/op
BenchmarkModAddBig-4                      860586              1749 ns/op
BenchmarkLargeModAddBig-4                 447950              2914 ns/op
BenchmarkMulBig-4                         435757              3613 ns/op
BenchmarkModMulBig-4                      176112              6350 ns/op
BenchmarkLargeModMulBig-4                 261256              5036 ns/op
BenchmarkModBig-4                        1238294               958.5 ns/op
BenchmarkLargeModBig-4                    511916              1963 ns/op
BenchmarkModInverseBig-4                  753504              1396 ns/op
BenchmarkLargeModInverseBig-4              96507             12173 ns/op
BenchmarkExpBig-4                           7321            143193 ns/op
BenchmarkLargeExpBig-4                        46          25072028 ns/op
BenchmarkSetBytesBig-4                   4645417               229.7 ns/op
BenchmarkAddNat-4                        7302646               168.2 ns/op
BenchmarkModAddNat-4                       57741             21532 ns/op
BenchmarkLargeModAddNat-4                 116892             11512 ns/op
BenchmarkMulNat-4                         141757              8153 ns/op
BenchmarkModMulNat-4                       27888             44462 ns/op
BenchmarkLargeModMulNat-4                  16935             71999 ns/op
BenchmarkModNat-4                          52398             20922 ns/op
BenchmarkLargeModNat-4                    114846             10610 ns/op
BenchmarkModInverseNat-4                   45645             26265 ns/op
BenchmarkLargeModInverseNat-4                228           4966108 ns/op
BenchmarkModInverseEvenNat-4                 159           7317795 ns/op
BenchmarkLargeModInverseEvenNat-4            154           7610100 ns/op
BenchmarkExpNat-4                           5998            186802 ns/op
BenchmarkLargeExpNat-4                        10         102149566 ns/op
BenchmarkSetBytesNat-4                    746450              2190 ns/op
PASS
ok      github.com/cronokirby/safenum   44.172s

[safenum] → go test -bench=. -tags math_big_pure_go
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        5950827               242.9 ns/op
BenchmarkModAddBig-4                     1000000              1464 ns/op
BenchmarkLargeModAddBig-4                 227496              5226 ns/op
BenchmarkMulBig-4                         244974              4665 ns/op
BenchmarkModMulBig-4                      204112              6142 ns/op
BenchmarkLargeModMulBig-4                 125530             10708 ns/op
BenchmarkModBig-4                        1205188              1031 ns/op
BenchmarkLargeModBig-4                    238192              6104 ns/op
BenchmarkModInverseBig-4                  768084              1936 ns/op
BenchmarkLargeModInverseBig-4              53173             26462 ns/op
BenchmarkExpBig-4                           8479            138762 ns/op
BenchmarkLargeExpBig-4                        22          48573421 ns/op
BenchmarkSetBytesBig-4                   4857422               234.5 ns/op
BenchmarkAddNat-4                        4293802               306.1 ns/op
BenchmarkModAddNat-4                       57752             23735 ns/op
BenchmarkLargeModAddNat-4                  89205             13366 ns/op
BenchmarkMulNat-4                          75943             15527 ns/op
BenchmarkModMulNat-4                       28130             43486 ns/op
BenchmarkLargeModMulNat-4                  15502             76571 ns/op
BenchmarkModNat-4                          57678             19955 ns/op
BenchmarkLargeModNat-4                    105062             12205 ns/op
BenchmarkModInverseNat-4                   43506             26133 ns/op
BenchmarkLargeModInverseNat-4                178           6637800 ns/op
BenchmarkModInverseEvenNat-4                 100          10713290 ns/op
BenchmarkLargeModInverseEvenNat-4             99          11205415 ns/op
BenchmarkExpNat-4                           6625            163308 ns/op
BenchmarkLargeExpNat-4                        10         106379810 ns/op
BenchmarkSetBytesNat-4                    594848              1728 ns/op
PASS
ok      github.com/cronokirby/safenum   40.055s
```

# f51bc7910016e703d1389250ee07a90eabcceac3 (2021-04-08)

Implement a more streamlined modular inversion routine.

```
[safenum] → go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        5783803               316.4 ns/op
BenchmarkModAddBig-4                     1000000              1103 ns/op
BenchmarkLargeModAddBig-4                 694466              1777 ns/op
BenchmarkMulBig-4                         438339              2668 ns/op
BenchmarkModMulBig-4                      276226              3624 ns/op
BenchmarkLargeModMulBig-4                 257265              4327 ns/op
BenchmarkModBig-4                        1267786               959.5 ns/op
BenchmarkLargeModBig-4                    670425              1718 ns/op
BenchmarkModInverseBig-4                  779264              1466 ns/op
BenchmarkLargeModInverseBig-4             103576             12414 ns/op
BenchmarkExpBig-4                           8476            139698 ns/op
BenchmarkLargeExpBig-4                        42          26590436 ns/op
BenchmarkSetBytesBig-4                   4994770               233.0 ns/op
BenchmarkAddNat-4                        5213209               193.6 ns/op
BenchmarkModAddNat-4                       57008             21120 ns/op
BenchmarkLargeModAddNat-4                 117346             10567 ns/op
BenchmarkMulNat-4                         149577              8031 ns/op
BenchmarkModMulNat-4                       54616             21473 ns/op
BenchmarkLargeModMulNat-4                  19770             61052 ns/op
BenchmarkModNat-4                          56593             21178 ns/op
BenchmarkLargeModNat-4                    115038             10264 ns/op
BenchmarkModInverseNat-4                   43377             27282 ns/op
BenchmarkLargeModInverseNat-4                232           5132395 ns/op
BenchmarkModInverseEvenNat-4                 152           7839481 ns/op
BenchmarkLargeModInverseEvenNat-4            130           9155569 ns/op
BenchmarkExpNat-4                           6046            195910 ns/op
BenchmarkLargeExpNat-4                        10         107854567 ns/op
BenchmarkSetBytesNat-4                    667095              1520 ns/op
PASS
ok      github.com/cronokirby/safenum   39.583s

[safenum] → go test -bench=. -tags math_big_pure_go
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        5109685               229.7 ns/op
BenchmarkModAddBig-4                     1000000              1172 ns/op
BenchmarkLargeModAddBig-4                 212607              5168 ns/op
BenchmarkMulBig-4                         243620              5084 ns/op
BenchmarkModMulBig-4                      205398              5749 ns/op
BenchmarkLargeModMulBig-4                 121932              9739 ns/op
BenchmarkModBig-4                        1000000              1031 ns/op
BenchmarkLargeModBig-4                    236596              5281 ns/op
BenchmarkModInverseBig-4                  683730              1575 ns/op
BenchmarkLargeModInverseBig-4              54470             23258 ns/op
BenchmarkExpBig-4                           8566            139349 ns/op
BenchmarkLargeExpBig-4                        22          50228266 ns/op
BenchmarkSetBytesBig-4                   5612074               268.1 ns/op
BenchmarkAddNat-4                        4814877               233.6 ns/op
BenchmarkModAddNat-4                       55846             21443 ns/op
BenchmarkLargeModAddNat-4                 108088             11674 ns/op
BenchmarkMulNat-4                          74908             16143 ns/op
BenchmarkModMulNat-4                       54472             21836 ns/op
BenchmarkLargeModMulNat-4                  18576             64639 ns/op
BenchmarkModNat-4                          55894             20642 ns/op
BenchmarkLargeModNat-4                    109884             10816 ns/op
BenchmarkModInverseNat-4                   44550             26488 ns/op
BenchmarkLargeModInverseNat-4                160           7306197 ns/op
BenchmarkModInverseEvenNat-4                 100          11287731 ns/op
BenchmarkLargeModInverseEvenNat-4             90          13300550 ns/op
BenchmarkExpNat-4                           6496            164679 ns/op
BenchmarkLargeExpNat-4                        10         110866574 ns/op
BenchmarkSetBytesNat-4                    764706              1486 ns/op
PASS
ok      github.com/cronokirby/safenum   40.311s
```

# 6768b30cbd9284b75aa387717f286b1d81edcf4f (2021-03-31)

Added even modular inversion

```
[safenum] → go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        8544007               187.3 ns/op
BenchmarkModAddBig-4                     1000000              1164 ns/op
BenchmarkLargeModAddBig-4                 660194              1883 ns/op
BenchmarkMulBig-4                         429873              2810 ns/op
BenchmarkModMulBig-4                      291237              3534 ns/op
BenchmarkLargeModMulBig-4                 273622              4221 ns/op
BenchmarkModBig-4                        1233248               957.5 ns/op
BenchmarkLargeModBig-4                    668458              1706 ns/op
BenchmarkModInverseBig-4                  881304              1400 ns/op
BenchmarkLargeModInverseBig-4             101865             11527 ns/op
BenchmarkExpBig-4                           8653            136691 ns/op
BenchmarkLargeExpBig-4                        45          25358294 ns/op
BenchmarkSetBytesBig-4                   6028752               204.8 ns/op
BenchmarkAddNat-4                        6311770               212.2 ns/op
BenchmarkModAddNat-4                       55579             21493 ns/op
BenchmarkLargeModAddNat-4                 116384             10377 ns/op
BenchmarkMulNat-4                         147770              7869 ns/op
BenchmarkModMulNat-4                       56871             21316 ns/op
BenchmarkLargeModMulNat-4                  20094             59612 ns/op
BenchmarkModNat-4                          58078             20815 ns/op
BenchmarkLargeModNat-4                    122284             10707 ns/op
BenchmarkModInverseNat-4                   44977             26167 ns/op
BenchmarkLargeModInverseNat-4                278           4266387 ns/op
BenchmarkModInverseEvenNat-4                 192           6378308 ns/op
BenchmarkLargeModInverseEvenNat-4            168           7169524 ns/op
BenchmarkExpNat-4                           6134            197088 ns/op
BenchmarkLargeExpNat-4                        10         108836798 ns/op
BenchmarkSetBytesNat-4                    785342              1525 ns/op
PASS
ok      github.com/cronokirby/safenum   42.185s

[safenum] → go test -bench=. -tags math_big_pure_go
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        6626785               206.2 ns/op
BenchmarkModAddBig-4                     1000000              1168 ns/op
BenchmarkLargeModAddBig-4                 241417              5376 ns/op
BenchmarkMulBig-4                         269402              4368 ns/op
BenchmarkModMulBig-4                      210955              5503 ns/op
BenchmarkLargeModMulBig-4                 124407              9295 ns/op
BenchmarkModBig-4                        1315934              1054 ns/op
BenchmarkLargeModBig-4                    244891              4897 ns/op
BenchmarkModInverseBig-4                  887767              1470 ns/op
BenchmarkLargeModInverseBig-4              56559             19833 ns/op
BenchmarkExpBig-4                           8506            136205 ns/op
BenchmarkLargeExpBig-4                        22          49501706 ns/op
BenchmarkSetBytesBig-4                   6449314               188.5 ns/op
BenchmarkAddNat-4                        5942666               206.2 ns/op
BenchmarkModAddNat-4                       58060             20800 ns/op
BenchmarkLargeModAddNat-4                 113953             11640 ns/op
BenchmarkMulNat-4                          79896             14701 ns/op
BenchmarkModMulNat-4                       57115             20863 ns/op
BenchmarkLargeModMulNat-4                  19510             62708 ns/op
BenchmarkModNat-4                          58936             20227 ns/op
BenchmarkLargeModNat-4                    115466             10372 ns/op
BenchmarkModInverseNat-4                   48084             25082 ns/op
BenchmarkLargeModInverseNat-4                195           5999599 ns/op
BenchmarkModInverseEvenNat-4                 122           9402313 ns/op
BenchmarkLargeModInverseEvenNat-4            100          10791384 ns/op
BenchmarkExpNat-4                           7232            158975 ns/op
BenchmarkLargeExpNat-4                        10         108437946 ns/op
BenchmarkSetBytesNat-4                    773128              1483 ns/op
PASS
ok      github.com/cronokirby/safenum   41.718s
```

# b89445f7bada17baf2db88b52f2a39a8a168ceea (2021-03-31)

Various minor optimizations around aliasing when multiplying

```
[safenum] → go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        7508224               162.0 ns/op
BenchmarkModAddBig-4                     1000000              1112 ns/op
BenchmarkLargeModAddBig-4                 605701              2081 ns/op
BenchmarkMulBig-4                         448010              3001 ns/op
BenchmarkModMulBig-4                      293226              3936 ns/op
BenchmarkLargeModMulBig-4                 277750              4489 ns/op
BenchmarkModBig-4                        1279603               920.3 ns/op
BenchmarkLargeModBig-4                    719839              1712 ns/op
BenchmarkModInverseBig-4                  867788              1480 ns/op
BenchmarkLargeModInverseBig-4             105806             12099 ns/op
BenchmarkExpBig-4                           8371            137640 ns/op
BenchmarkLargeExpBig-4                        43          25984116 ns/op
BenchmarkSetBytesBig-4                   6138981               200.8 ns/op
BenchmarkAddNat-4                        6275767               168.9 ns/op
BenchmarkModAddNat-4                       57129             20985 ns/op
BenchmarkLargeModAddNat-4                 115216             10546 ns/op
BenchmarkMulNat-4                         148363              8168 ns/op
BenchmarkModMulNat-4                       54138             21883 ns/op
BenchmarkLargeModMulNat-4                  19264             63472 ns/op
BenchmarkModNat-4                          54705             20608 ns/op
BenchmarkLargeModNat-4                    119592             10277 ns/op
BenchmarkModInverseNat-4                   41794             26500 ns/op
BenchmarkLargeModInverseNat-4                276           4266428 ns/op
BenchmarkExpNat-4                           6061            192362 ns/op
BenchmarkLargeExpNat-4                        10         106890226 ns/op
BenchmarkSetBytesNat-4                    690368              1506 ns/op
PASS
ok      github.com/cronokirby/safenum   35.207s

[safenum] → go test -bench=. -tags math_big_pure_go
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        5036264               221.7 ns/op
BenchmarkModAddBig-4                     1000000              1173 ns/op
BenchmarkLargeModAddBig-4                 235732              5102 ns/op
BenchmarkMulBig-4                         256430              4607 ns/op
BenchmarkModMulBig-4                      214002              5646 ns/op
BenchmarkLargeModMulBig-4                 126057              9631 ns/op
BenchmarkModBig-4                        1261798               948.9 ns/op
BenchmarkLargeModBig-4                    241308              5021 ns/op
BenchmarkModInverseBig-4                  883976              1482 ns/op
BenchmarkLargeModInverseBig-4              58646             20443 ns/op
BenchmarkExpBig-4                           8444            140072 ns/op
BenchmarkLargeExpBig-4                        21          50785149 ns/op
BenchmarkSetBytesBig-4                   4901631               264.0 ns/op
BenchmarkAddNat-4                        5093500               259.5 ns/op
BenchmarkModAddNat-4                       53115             21281 ns/op
BenchmarkLargeModAddNat-4                 108441             11162 ns/op
BenchmarkMulNat-4                          77780             15389 ns/op
BenchmarkModMulNat-4                       54999             21898 ns/op
BenchmarkLargeModMulNat-4                  18313             64847 ns/op
BenchmarkModNat-4                          56948             20848 ns/op
BenchmarkLargeModNat-4                    110329             10730 ns/op
BenchmarkModInverseNat-4                   45327             26150 ns/op
BenchmarkLargeModInverseNat-4                181           6601408 ns/op
BenchmarkExpNat-4                           7132            166127 ns/op
BenchmarkLargeExpNat-4                        10         111626393 ns/op
BenchmarkSetBytesNat-4                    773126              1467 ns/op
PASS
ok      github.com/cronokirby/safenum   36.167s
```

# c0e31da784cec1419655a732a96f03c42bc3d97f (2021-03-31)

Implement basic Montgomery multiplication

```
[safenum] → go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        7037252               225.2 ns/op
BenchmarkModAddBig-4                     1000000              1158 ns/op
BenchmarkLargeModAddBig-4                 686259              1750 ns/op
BenchmarkMulBig-4                         448657              2704 ns/op
BenchmarkModMulBig-4                      339421              3771 ns/op
BenchmarkLargeModMulBig-4                 258805              4402 ns/op
BenchmarkModBig-4                        1264462               944.2 ns/op
BenchmarkLargeModBig-4                    685807              1785 ns/op
BenchmarkModInverseBig-4                  843180              1446 ns/op
BenchmarkLargeModInverseBig-4              98060             11984 ns/op
BenchmarkExpBig-4                           8458            138325 ns/op
BenchmarkLargeExpBig-4                        40          26431885 ns/op
BenchmarkSetBytesBig-4                   4958301               270.1 ns/op
BenchmarkAddNat-4                        5462413               223.7 ns/op
BenchmarkModAddNat-4                       55483             21294 ns/op
BenchmarkLargeModAddNat-4                 113378             10548 ns/op
BenchmarkMulNat-4                         149622              7999 ns/op
BenchmarkModMulNat-4                       28458             41601 ns/op
BenchmarkLargeModMulNat-4                  16954             70324 ns/op
BenchmarkModNat-4                          57303             20695 ns/op
BenchmarkLargeModNat-4                    116892             10228 ns/op
BenchmarkModInverseNat-4                   45306             26319 ns/op
BenchmarkLargeModInverseNat-4                277           4395341 ns/op
BenchmarkExpNat-4                           4910            238846 ns/op
BenchmarkLargeExpNat-4                        10         111907965 ns/op
BenchmarkSetBytesNat-4                    754429              1452 ns/op
PASS
ok      github.com/cronokirby/safenum   36.098s

[safenum] → go test -bench=. -tags math_big_pure_go
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        5705341               217.9 ns/op
BenchmarkModAddBig-4                      982310              1138 ns/op
BenchmarkLargeModAddBig-4                 218103              5170 ns/op
BenchmarkMulBig-4                         239662              4615 ns/op
BenchmarkModMulBig-4                      212642              5477 ns/op
BenchmarkLargeModMulBig-4                 128007              9509 ns/op
BenchmarkModBig-4                        1297141               938.8 ns/op
BenchmarkLargeModBig-4                    239312              5124 ns/op
BenchmarkModInverseBig-4                  826834              1389 ns/op
BenchmarkLargeModInverseBig-4              60763             20150 ns/op
BenchmarkExpBig-4                           8456            138709 ns/op
BenchmarkLargeExpBig-4                        22          49983921 ns/op
BenchmarkSetBytesBig-4                   5739404               220.6 ns/op
BenchmarkAddNat-4                        4374946               257.3 ns/op
BenchmarkModAddNat-4                       56361             21243 ns/op
BenchmarkLargeModAddNat-4                 110377             10957 ns/op
BenchmarkMulNat-4                          78637             15193 ns/op
BenchmarkModMulNat-4                       28364             41493 ns/op
BenchmarkLargeModMulNat-4                  16376             73648 ns/op
BenchmarkModNat-4                          56200             21209 ns/op
BenchmarkLargeModNat-4                    111158             10637 ns/op
BenchmarkModInverseNat-4                   45657             25961 ns/op
BenchmarkLargeModInverseNat-4                183           6448992 ns/op
BenchmarkExpNat-4                           5218            217985 ns/op
BenchmarkLargeExpNat-4                         9         112084020 ns/op
BenchmarkSetBytesNat-4                    746052              1433 ns/op
PASS
ok      github.com/cronokirby/safenum   35.794s
```

# a24e618bccbc040c7121394c889e8bdd0dca2d01 (2021-03-29)

Implement free limb injection

```
[safenum] → go test -bench=.                                                                                                                                                     
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        6953644               212.4 ns/op
BenchmarkModAddBig-4                     1100115              1040 ns/op
BenchmarkLargeModAddBig-4                 585422              1841 ns/op
BenchmarkMulBig-4                         429370              3033 ns/op
BenchmarkModMulBig-4                      339429              3708 ns/op
BenchmarkLargeModMulBig-4                 284680              4709 ns/op
BenchmarkModBig-4                        1257049               936.7 ns/op
BenchmarkLargeModBig-4                    650443              1853 ns/op
BenchmarkModInverseBig-4                  823682              1468 ns/op
BenchmarkLargeModInverseBig-4              98150             12050 ns/op
BenchmarkExpBig-4                           8720            135014 ns/op
BenchmarkLargeExpBig-4                        43          25975861 ns/op
BenchmarkSetBytesBig-4                   5997328               225.9 ns/op
BenchmarkAddNat-4                        5293556               211.6 ns/op
BenchmarkModAddNat-4                       55104             21435 ns/op
BenchmarkLargeModAddNat-4                 108148             10385 ns/op
BenchmarkMulNat-4                         140707              8281 ns/op
BenchmarkModMulNat-4                       24109             51049 ns/op
BenchmarkLargeModMulNat-4                  17731             63715 ns/op
BenchmarkModNat-4                          56545             20935 ns/op
BenchmarkLargeModNat-4                    114956             10484 ns/op
BenchmarkModInverseNat-4                   44448             26838 ns/op
BenchmarkLargeModInverseNat-4                271           4304645 ns/op
BenchmarkExpNat-4                            196           5948716 ns/op
BenchmarkLargeExpNat-4                         3         359713739 ns/op
BenchmarkSetBytesNat-4                    836378              1506 ns/op
PASS
ok      github.com/cronokirby/safenum   44.859s
[safenum] → go test -bench=. -tags math_big_pure_go                                                                                                                              
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        5669407               252.4 ns/op
BenchmarkModAddBig-4                      946286              1171 ns/op
BenchmarkLargeModAddBig-4                 215196              5212 ns/op
BenchmarkMulBig-4                         222129              4557 ns/op
BenchmarkModMulBig-4                      215752              5561 ns/op
BenchmarkLargeModMulBig-4                 125840              9571 ns/op
BenchmarkModBig-4                        1275968              1004 ns/op
BenchmarkLargeModBig-4                    234193              5070 ns/op
BenchmarkModInverseBig-4                  877096              1414 ns/op
BenchmarkLargeModInverseBig-4              53406             23734 ns/op
BenchmarkExpBig-4                           8257            141085 ns/op
BenchmarkLargeExpBig-4                        21          51494745 ns/op
BenchmarkSetBytesBig-4                   4290268               292.7 ns/op
BenchmarkAddNat-4                        5286901               272.6 ns/op
BenchmarkModAddNat-4                       56404             21761 ns/op
BenchmarkLargeModAddNat-4                 109947             11317 ns/op
BenchmarkMulNat-4                          77644             15774 ns/op
BenchmarkModMulNat-4                       21037             56950 ns/op
BenchmarkLargeModMulNat-4                  16245             73966 ns/op
BenchmarkModNat-4                          57388             20625 ns/op
BenchmarkLargeModNat-4                    111604             10840 ns/op
BenchmarkModInverseNat-4                   45824             25909 ns/op
BenchmarkLargeModInverseNat-4                181           6707925 ns/op
BenchmarkExpNat-4                            192           6192261 ns/op
BenchmarkLargeExpNat-4                         3         415365259 ns/op
BenchmarkSetBytesNat-4                    613538              1669 ns/op
PASS
ok      github.com/cronokirby/safenum   39.011s
```

# 020e34e3436d885500318e1777e7364a1c3c393d (2021-03-27)

Limb by limb reduction

``` 
[safenum] → go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        7193640               174.0 ns/op
BenchmarkModAddBig-4                     1129430              1049 ns/op
BenchmarkLargeModAddBig-4                 637372              1735 ns/op
BenchmarkMulBig-4                         470226              2635 ns/op
BenchmarkModMulBig-4                      313644              3456 ns/op
BenchmarkLargeModMulBig-4                 257732              4091 ns/op
BenchmarkModBig-4                        1342916               891.9 ns/op
BenchmarkLargeModBig-4                    669978              1662 ns/op
BenchmarkModInverseBig-4                  772887              1358 ns/op
BenchmarkLargeModInverseBig-4             102363             11236 ns/op
BenchmarkExpBig-4                           8823            136720 ns/op
BenchmarkLargeExpBig-4                        40          25947153 ns/op
BenchmarkSetBytesBig-4                   5259864               209.5 ns/op
BenchmarkAddNat-4                        6865567               177.1 ns/op
BenchmarkModAddNat-4                       56781             21101 ns/op
BenchmarkLargeModAddNat-4                  26634             44499 ns/op
BenchmarkMulNat-4                         147716              7994 ns/op
BenchmarkModMulNat-4                       23832             49599 ns/op
BenchmarkLargeModMulNat-4                  12249             96907 ns/op
BenchmarkModNat-4                          56899             20651 ns/op
BenchmarkLargeModNat-4                     26866             45090 ns/op
BenchmarkModInverseNat-4                   44586             26605 ns/op
BenchmarkLargeModInverseNat-4                265           4357764 ns/op
BenchmarkExpNat-4                            200           5926589 ns/op
BenchmarkLargeExpNat-4                         2         622324074 ns/op
BenchmarkSetBytesNat-4                    786177              1445 ns/op
PASS
ok      github.com/cronokirby/safenum   40.069s

[safenum] → go test -bench=. -tags math_big_pure_go
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        5884750               222.7 ns/op
BenchmarkModAddBig-4                      888044              1178 ns/op
BenchmarkLargeModAddBig-4                 218853              5074 ns/op
BenchmarkMulBig-4                         256408              4895 ns/op
BenchmarkModMulBig-4                      202692              5761 ns/op
BenchmarkLargeModMulBig-4                 111362              9603 ns/op
BenchmarkModBig-4                        1289430               915.5 ns/op
BenchmarkLargeModBig-4                    238866              4949 ns/op
BenchmarkModInverseBig-4                  892677              1415 ns/op
BenchmarkLargeModInverseBig-4              60750             20164 ns/op
BenchmarkExpBig-4                           8632            140500 ns/op
BenchmarkLargeExpBig-4                        22          51087158 ns/op
BenchmarkSetBytesBig-4                   4668199               243.2 ns/op
BenchmarkAddNat-4                        4858521               255.4 ns/op
BenchmarkModAddNat-4                       56320             21361 ns/op
BenchmarkLargeModAddNat-4                  24938             48347 ns/op
BenchmarkMulNat-4                          78342             15215 ns/op
BenchmarkModMulNat-4                       20944             55907 ns/op
BenchmarkLargeModMulNat-4                  10000            111708 ns/op
BenchmarkModNat-4                          56568             20852 ns/op
BenchmarkLargeModNat-4                     23888             48265 ns/op
BenchmarkModInverseNat-4                   44446             26053 ns/op
BenchmarkLargeModInverseNat-4                177           6559094 ns/op
BenchmarkExpNat-4                            202           5873289 ns/op
BenchmarkLargeExpNat-4                         2         717422832 ns/op
BenchmarkSetBytesNat-4                    686104              1476 ns/op
PASS
ok      github.com/cronokirby/safenum   39.768s
```

# b66ac649d9ca1b1b394a7706cea6495b107dcb9c (2021-03-25)

```
[safenum] → go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        7233490               168.4 ns/op
BenchmarkModAddBig-4                     1000000              1051 ns/op
BenchmarkLargeModAddBig-4                 599292              1772 ns/op
BenchmarkMulBig-4                         458648              2567 ns/op
BenchmarkModMulBig-4                      324234              3573 ns/op
BenchmarkLargeModMulBig-4                 285141              4197 ns/op
BenchmarkModBig-4                        1239144              1043 ns/op
BenchmarkLargeModBig-4                    711975              1732 ns/op
BenchmarkModInverseBig-4                  778612              1438 ns/op
BenchmarkLargeModInverseBig-4              91058             13763 ns/op
BenchmarkExpBig-4                           7298            139626 ns/op
BenchmarkLargeExpBig-4                        42          25930457 ns/op
BenchmarkSetBytesBig-4                   4770934               240.3 ns/op
BenchmarkAddNat-4                        6586465               166.4 ns/op
BenchmarkModAddNat-4                       27013             44123 ns/op
BenchmarkLargeModAddNat-4                   2725            438992 ns/op
BenchmarkMulNat-4                         140364              7858 ns/op
BenchmarkModMulNat-4                       12786             94944 ns/op
BenchmarkLargeModMulNat-4                   1135            893462 ns/op
BenchmarkModNat-4                          27309             41986 ns/op
BenchmarkLargeModNat-4                      2733            438626 ns/op
BenchmarkModInverseNat-4                   25102             48908 ns/op
BenchmarkLargeModInverseNat-4                237           4906628 ns/op
BenchmarkExpNat-4                             98          11851448 ns/op
BenchmarkLargeExpNat-4                         1        5541042938 ns/op
BenchmarkSetBytesNat-4                    788446              1696 ns/op
PASS
ok      github.com/cronokirby/safenum   41.362s

[safenum] → go test -bench=. -tags math_big_pure_go
goos: linux
goarch: amd64
pkg: github.com/cronokirby/safenum
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
BenchmarkAddBig-4                        6008720               206.7 ns/op
BenchmarkModAddBig-4                     1000000              1141 ns/op
BenchmarkLargeModAddBig-4                 226074              4899 ns/op
BenchmarkMulBig-4                         261199              5061 ns/op
BenchmarkModMulBig-4                      159643              6593 ns/op
BenchmarkLargeModMulBig-4                 119554              9764 ns/op
BenchmarkModBig-4                        1298679              1030 ns/op
BenchmarkLargeModBig-4                    226292              5268 ns/op
BenchmarkModInverseBig-4                  810531              1395 ns/op
BenchmarkLargeModInverseBig-4              58944             21170 ns/op
BenchmarkExpBig-4                           8494            136181 ns/op
BenchmarkLargeExpBig-4                        22          50598450 ns/op
BenchmarkSetBytesBig-4                   5431112               199.1 ns/op
BenchmarkAddNat-4                        6781236               177.2 ns/op
BenchmarkModAddNat-4                       33860             36012 ns/op
BenchmarkLargeModAddNat-4                   1794            664556 ns/op
BenchmarkMulNat-4                          79315             14716 ns/op
BenchmarkModMulNat-4                       14188             83209 ns/op
BenchmarkLargeModMulNat-4                    872           1320999 ns/op
BenchmarkModNat-4                          34676             35121 ns/op
BenchmarkLargeModNat-4                      1726            674276 ns/op
BenchmarkModInverseNat-4                   29527             40186 ns/op
BenchmarkLargeModInverseNat-4                164           7174169 ns/op
BenchmarkExpNat-4                            120           9723271 ns/op
BenchmarkLargeExpNat-4                         1        8705644301 ns/op
BenchmarkSetBytesNat-4                    794876              1570 ns/op
PASS
ok      github.com/cronokirby/safenum   46.227s
```