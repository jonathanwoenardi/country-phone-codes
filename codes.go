package country_phone_codes

// countryCodeToPhoneCodeMap is a map of country codes to its corresponding country phone codes.
// Country codes are ISO 3166-1 alpha-2.
// Source: https://www.iso.org/obp/ui.
// Country phone codes are defined by ITU-T in standard E.164.
// Source: https://www.itu.int/dms_pub/itu-t/opb/sp/T-SP-E.164D-11-2011-PDF-E.pdf.
// The map is produced by combining two sources mentioned above.
var countryCodeToPhoneCodeMap = map[string][]string{
	"AD": []string{"+376"},
	"AE": []string{"+971"},
	"AF": []string{"+93"},
	"AG": []string{"+1"},
	"AI": []string{"+1"},
	"AL": []string{"+355"},
	"AM": []string{"+374"},
	"AO": []string{"+244"},
	"AQ": []string{"+672"}, // 672 is for Australian Antarctic Territory
	"AR": []string{"+54"},
	"AS": []string{"+1"},
	"AT": []string{"+43"},
	"AU": []string{"+61"},
	"AW": []string{"+297"},
	"AX": []string{"+358"}, // sharing with Finland
	"AZ": []string{"+994"},
	"BA": []string{"+387"},
	"BB": []string{"+1"},
	"BD": []string{"+880"},
	"BE": []string{"+32"},
	"BF": []string{"+226"},
	"BG": []string{"+359"},
	"BH": []string{"+973"},
	"BI": []string{"+257"},
	"BJ": []string{"+229"},
	"BL": []string{"+590"}, // sharing with Guadeloupe
	"BM": []string{"+1"},
	"BN": []string{"+673"},
	"BO": []string{"+591"},
	"BQ": []string{"+599"}, // sharing with Curaçao
	"BR": []string{"+55"},
	"BS": []string{"+1"},
	"BT": []string{"+975"},
	"BV": []string{}, // no assigned country codes
	"BW": []string{"+267"},
	"BY": []string{"+375"},
	"BZ": []string{"+501"},
	"CA": []string{"+1"},
	"CC": []string{"+61"}, // sharing with Australia
	"CD": []string{"+243"},
	"CF": []string{"+236"},
	"CG": []string{"+242"},
	"CH": []string{"+41"},
	"CI": []string{"+225"},
	"CK": []string{"+682"},
	"CL": []string{"+56"},
	"CM": []string{"+237"},
	"CN": []string{"+86"},
	"CO": []string{"+57"},
	"CR": []string{"+506"},
	"CU": []string{"+53"},
	"CV": []string{"+238"},
	"CW": []string{"+599"},
	"CX": []string{"+61"}, // sharing with Australia
	"CY": []string{"+357"},
	"CZ": []string{"+420"},
	"DE": []string{"+49"},
	"DJ": []string{"+253"},
	"DK": []string{"+45"},
	"DM": []string{"+1"},
	"DO": []string{"+1"},
	"DZ": []string{"+213"},
	"EC": []string{"+593"},
	"EE": []string{"+372"},
	"EG": []string{"+20"},
	"EH": []string{"+212"}, // sharing with Morocco
	"ER": []string{"+291"},
	"ES": []string{"+34"},
	"ET": []string{"+251"},
	"FI": []string{"+358"},
	"FJ": []string{"+679"},
	"FK": []string{"+500"},
	"FM": []string{"+691"},
	"FO": []string{"+298"},
	"FR": []string{"+33"},
	"GA": []string{"+241"},
	"GB": []string{"+44"},
	"GD": []string{"+1"},
	"GE": []string{"+995"},
	"GF": []string{"+594"},
	"GG": []string{"+44"}, // sharing with United Kingdom
	"GH": []string{"+233"},
	"GI": []string{"+350"},
	"GL": []string{"+299"},
	"GM": []string{"+220"},
	"GN": []string{"+224"},
	"GP": []string{"+590"},
	"GQ": []string{"+240"},
	"GR": []string{"+30"},
	"GS": []string{"+500"}, // sharing with Falkland Islands
	"GT": []string{"+502"},
	"GU": []string{"+1"},
	"GW": []string{"+245"},
	"GY": []string{"+592"},
	"HK": []string{"+852"},
	"HM": []string{}, // no assigned country codes
	"HN": []string{"+504"},
	"HR": []string{"+385"},
	"HT": []string{"+509"},
	"HU": []string{"+36"},
	"ID": []string{"+62"},
	"IE": []string{"+353"},
	"IL": []string{"+972"},
	"IM": []string{"+44"}, // sharing with United Kingdom
	"IN": []string{"+91"},
	"IO": []string{"+246"},
	"IQ": []string{"+964"},
	"IR": []string{"+98"},
	"IS": []string{"+354"},
	"IT": []string{"+39"},
	"JE": []string{"+44"}, // sharing with United Kingdom
	"JM": []string{"+1"},
	"JO": []string{"+962"},
	"JP": []string{"+81"},
	"KE": []string{"+254"},
	"KG": []string{"+996"},
	"KH": []string{"+855"},
	"KI": []string{"+686"},
	"KM": []string{"+269"},
	"KN": []string{"+1"},
	"KP": []string{"+850"},
	"KR": []string{"+82"},
	"KW": []string{"+965"},
	"KY": []string{"+1"},
	"KZ": []string{"+7"},
	"LA": []string{"+856"},
	"LB": []string{"+961"},
	"LC": []string{"+1"},
	"LI": []string{"+423"},
	"LK": []string{"+94"},
	"LR": []string{"+231"},
	"LS": []string{"+266"},
	"LT": []string{"+370"},
	"LU": []string{"+352"},
	"LV": []string{"+371"},
	"LY": []string{"+218"},
	"MA": []string{"+212"},
	"MC": []string{"+377"},
	"MD": []string{"+373"},
	"ME": []string{"+382"},
	"MF": []string{"+590"}, // sharing with Guadeloupe
	"MG": []string{"+261"},
	"MH": []string{"+692"},
	"MK": []string{"+389"},
	"ML": []string{"+223"},
	"MM": []string{"+95"},
	"MN": []string{"+976"},
	"MO": []string{"+853"},
	"MP": []string{"+1"},
	"MQ": []string{"+596"},
	"MR": []string{"+222"},
	"MS": []string{"+1"},
	"MT": []string{"+356"},
	"MU": []string{"+230"},
	"MV": []string{"+960"},
	"MW": []string{"+265"},
	"MX": []string{"+52"},
	"MY": []string{"+60"},
	"MZ": []string{"+258"},
	"NA": []string{"+264"},
	"NC": []string{"+687"},
	"NE": []string{"+227"},
	"NF": []string{"+672"}, // part of Australian External Territories
	"NG": []string{"+234"},
	"NI": []string{"+505"},
	"NL": []string{"+31"},
	"NO": []string{"+47"},
	"NP": []string{"+977"},
	"NR": []string{"+674"},
	"NU": []string{"+683"},
	"NZ": []string{"+64"},
	"OM": []string{"+968"},
	"PA": []string{"+507"},
	"PE": []string{"+51"},
	"PF": []string{"+689"},
	"PG": []string{"+675"},
	"PH": []string{"+63"},
	"PK": []string{"+92"},
	"PL": []string{"+48"},
	"PM": []string{"+508"},
	"PN": []string{"+64"},
	"PR": []string{"+1"},
	"PS": []string{"+970"},
	"PT": []string{"+351"},
	"PW": []string{"+680"},
	"PY": []string{"+595"},
	"QA": []string{"+974"},
	"RE": []string{"+262"}, // part of French Departments and Territories in the Indian Ocean
	"RO": []string{"+40"},
	"RS": []string{"+381"},
	"RU": []string{"+7"},
	"RW": []string{"+250"},
	"SA": []string{"+966"},
	"SB": []string{"+677"},
	"SC": []string{"+248"},
	"SD": []string{"+249"},
	"SE": []string{"+46"},
	"SG": []string{"+65"},
	"SH": []string{"+247", "+290"}, // 247 is for Ascension, 290 is for Saint Helena and Tristan da Cuhna
	"SI": []string{"+386"},
	"SJ": []string{"+47"}, // sharing with Norway
	"SK": []string{"+421"},
	"SL": []string{"+232"},
	"SM": []string{"+378"},
	"SN": []string{"+221"},
	"SO": []string{"+252"},
	"SR": []string{"+597"},
	"SS": []string{"+211"},
	"ST": []string{"+239"},
	"SV": []string{"+503"},
	"SX": []string{"+1"},
	"SY": []string{"+963"},
	"SZ": []string{"+268"},
	"TC": []string{"+1"},
	"TD": []string{"+235"},
	"TF": []string{"+262"}, // part of French Departments and Territories in the Indian Ocean
	"TG": []string{"+228"},
	"TH": []string{"+66"},
	"TJ": []string{"+992"},
	"TK": []string{"+690"},
	"TL": []string{"+670"},
	"TM": []string{"+993"},
	"TN": []string{"+216"},
	"TO": []string{"+676"},
	"TR": []string{"+90"},
	"TT": []string{"+1"},
	"TV": []string{"+688"},
	"TW": []string{"+886"},
	"TZ": []string{"+255"},
	"UA": []string{"+380"},
	"UG": []string{"+256"},
	"UM": []string{}, // no assigned country codes
	"US": []string{"+1"},
	"UY": []string{"+598"},
	"UZ": []string{"+998"},
	"VA": []string{"+39", "+379"}, // 379 is reserved for future use
	"VC": []string{"+1"},
	"VE": []string{"+58"},
	"VG": []string{"+1"},
	"VI": []string{"+1"},
	"VN": []string{"+84"},
	"VU": []string{"+678"},
	"WF": []string{"+681"},
	"WS": []string{"+685"},
	"YE": []string{"+967"},
	"YT": []string{"+262"}, // part of French Departments and Territories in the Indian Ocean
	"ZA": []string{"+27"},
	"ZM": []string{"+260"},
	"ZW": []string{"+263"},
}

// nanpAreaCodeMap is a map of ISO 3166 country codes included in NANP to their coressponding country + area phone code.
// North American Numbering Plan (NANP) is a telephone numbering plan used in most North American countries.
// Source: https://nationalnanpa.com/enas/geoAreaCodeAlphabetReport.do.
// The map currently do not included regions from Canada or United States.
// Therefore, the map cannot be used to differentiate Canada and United States phone number.
var nanpAreaCodeToPhoneCodeMap = map[string][]string{
	"AG": []string{"+1268"},
	"AI": []string{"+1264"},
	"AS": []string{"+1684"},
	"BB": []string{"+1246"},
	"BM": []string{"+1441"},
	"BS": []string{"+1242"},
	"CA": []string{"+1"},
	"DM": []string{"+1767"},
	"DO": []string{"+1809", "+1829", "+1849"},
	"GD": []string{"+1473"},
	"GU": []string{"+1671"},
	"JM": []string{"+1658", "+1876"},
	"KN": []string{"+1869"},
	"KY": []string{"+1345"},
	"LC": []string{"+1758"},
	"MP": []string{"+1670"},
	"MS": []string{"+1664"},
	"PR": []string{"+1787", "+1939"},
	"SX": []string{"+1721"},
	"TC": []string{"+1649"},
	"TT": []string{"+1868"},
	"US": []string{"+1"},
	"VC": []string{"+1784"},
	"VG": []string{"+1284"},
	"VI": []string{"+1340"},
}
