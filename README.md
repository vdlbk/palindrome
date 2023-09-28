# palindrome

The idea was to associate every letter of an alphabet to a number. The number would be converted in base `K`, where `K` is the number of letters in the alphabet. 
We then want to find palindromic words that, converted in `base 10`, have also a palindromic number in `base 10`.

Example:    
The word `KYRGRYK` gives the number `5760110675` in `base 10` and `BQJ7JQB` in `base K`, using the Spanish alphabet with the letters in alphabetic orders. Though, it does not have a meaning in Spanish. Maybe in other languages.

Finding meaningful words or phrases looks almost impossible. But who knows?

| spanish letters | index | baseK |
|-----------------|-------|-------|
| _               | 0     | 0     |
| A               | 1     | 1     |
| B               | 2     | 2     |
| C               | 3     | 3     |
| D               | 4     | 4     |
| E               | 5     | 5     |
| F               | 6     | 6     |
| G               | 7     | 7     |
| H               | 8     | 8     |
| I               | 9     | 9     |
| J               | 10    | A     |
| K               | 12    | B     |
| L               | 11    | C     |
| M               | 13    | D     |
| N               | 14    | E     |
| Ñ               | 15    | F     |
| O               | 16    | G     |
| P               | 17    | H     |
| Q               | 18    | I     |
| R               | 19    | J     |
| S               | 20    | K     |
| T               | 21    | L     |
| U               | 22    | M     |
| V               | 23    | N     |
| W               | 24    | O     |
| X               | 25    | P     |
| Y               | 26    | Q     |
| Z               | 27    | R     |

## Result

| Number in base K | Number in base 10 | Word in Spanish Alphabet (Letter in alphabetic order) | Word in Spanish Alphabet (Vowel in even position) | Word in Spanish Alphabet (Order by most common letters) | Word in Spanish Alphabet (Vowels first) | Word in Spanish Alphabet (Order random) |
|-----------------|-------|-------| --- | --- | --- | --- |
| | | `_ABCDEFGHIJKLMNÑOPQRSTUVWXYZ` | `_BACDFEGHJIKLMNÑOPQRSTUVWXYZ` | `_EAOSRNIDLCTUMPBGVYQHFZJÑXKW` | `_EAOIUNSDLCTRMPBGVYQHFZJÑXKW` | `_EAFIUNSDLCTRMPBGJYQHOZVÑXKW` |
| 0 | 0 | | _ | _ | _ | _ | _ |
| 1 | 1 | A | B | E | E | E |
| 2 | 2 | B | A | A | A | A |
| 3 | 3 | C | C | O | O | F |
| 4 | 4 | D | D | S | I | I |
| 5 | 5 | E | F | R | U | U |
| 6 | 6 | F | E | N | N | N |
| 7 | 7 | G | G | I | S | S |
| 8 | 8 | H | H | D | D | D |
| 9 | 9 | I | J | L | L | L |
| B | 11 | K | K | T | T | T |
| M | 22 | U | U | Z | Z | Z |
| 88 | 232 | HH | HH | DD | DD | DD |
| GG | 464 | OO | OO | GG | GG | GG |
| OO | 696 | WW | WW | ÑÑ | ÑÑ | ÑÑ |
| 7A7 | 5775 | GJG | GIG | ICI | SCS | SCS |
| 828 | 6336 | HBH | HAH | DAD | DAD | DAD |
| D7D | 10401 | MGM | MGM | MIM | MSM | MSM |
| JCJ | 15251 | RLR | RLR | QUQ | QRQ | QRQ |
| LHL | 16961 | TPT | TPT | FVF | FVF | OJO |
| QEQ | 20802 | YNY | YNY | KPK | KPK | KPK |
| 4NBN4 | 2972792 | DVKVD | DVKVD | SJTJS | IJTJI | IVTVI |
| 97D79 | 5695965 | IGMGI | JGMGJ | LIMIL | LSMSL | LSMSL |
| A3O3A | 6231326 | JCWCJ | ICWCI | COÑOC | COÑOC | CFÑFC |
| AKJKA | 6601066 | JSRSJ | ISRSI | CHQHC | CHQHC | CHQHC |
| BB0BB | 7003007 | KK_KK | KK_KK | TT_TT | TT_TT | TT_TT |
| E8A8E | 8788878 | NHJHN | NHIHN | PDCDP | PDCDP | PDCDP |
| F4L4F | 9324239 | ÑDTDÑ | ÑDTDÑ | BSFSB | BIFIB | BIOIB |
| H282H | 10499401 | PBHBP | PAHAP | VADAV | VADAV | JADAJ |
| 9A00A9 | 161040161 | IJ__JI | JI__IJ | LC__CL | LC__CL | LC__CL |
| ONQQNO | 427777724 | WVYYVW | WVYYVW | ÑJKKJÑ | ÑJKKJÑ | ÑVKKVÑ |
| 112C211 | 500595005 | AABLBAA | BBALABB | EEAUAEE | EEARAEE | EEARAEE |
| 1D9R9D1 | 711757117 | AMIZIMA | BMJZJMB | EMLWLME | EMLWLME | EMLWLME |
| 1FILIF1 | 751585157 | AÑQTQÑA | BÑQTQÑB | EBYFYBE | EBYFYBE | EBYOYBE |
| 1GR3RG1 | 773939377 | AOZCZOA | BOZCZOB | EGWOWGE | EGWOWGE | EGWFWGE |
| 1OF0FO1 | 904171409 | AWÑ_ÑWA | BWÑ_ÑWB | EÑB_BÑE | EÑB_BÑE | EÑB_BÑE |
| 1R0Q0R1 | 947141749 | AZ_Y_ZA | BZ_Y_ZB | EW_K_WE | EW_K_WE | EW_K_WE |
| 35B8B53 | 1538668351 | CEKHKEC | CFKHKFC | ORTDTRO | OUTDTUO | FUTDTUF |
| BQJ7JQB | 5760110675 | KYRGRYK | KYRGRYK | TKQIQKT | TKQSQKT | TKQSQKT |
| CMM7MMC | 6175005716 | LUUGUUL | LUUGUUL | UZZIZZU | RZZSZZR | RZZSZZR |
| F41E14F | 7298118927 | ÑDANADÑ | ÑDBNBDÑ | BSEPESB | BIEPEIB | BIEPEIB |
| LOR6ROL | 10549494501 | TWZFZWT | TWZEZWT | FÑWNWÑF | FÑWNWÑF | OÑWNWÑO |
| N96L69N | 11242524211 | VIFTFIV | VJETEJV | JLNFNLJ | JLNFNLJ | VLNONLV |
| RD8E8DR | 13240004231 | ZMHNHMZ | ZMHNHMZ | WMDPDMW | WMDPDMW | WMDPDMW |
| RKCNCKR | 13363136331 | ZSLVLSZ | ZSLVLSZ | WHUJUHW | WHRJRHW | WHRVRHW |
| 5O5DD5O5 | 79124342197 | EWEMMEWE | FWFMMFWF | RÑRMMRÑR | UÑUMMUÑU | UÑUMMUÑU |
| 6H7EE7H6 | 89279097298 | FPGNNGPF | EPGNNGPE | NVIPPIVN | NVSPPSVN | NJSPPSJN |
| KD5HH5DK | 276220022672 | SMEPPEMS | SMFPPFMS | HMRVVRMH | HMUVVUMH | HMUJJUMH |
| 11A414A11 | 396183381693 | AAJDADJAA | BBIDBDIBB | EECSESCEE | EECIEICEE | EECIEICEE |

## Follow-ups
To improve the search, it needs to improve the way we find palindromic numbers. 
Also, we can try to optimize even better the computation accross threads.



