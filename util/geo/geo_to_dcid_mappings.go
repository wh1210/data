// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package geo

// CountryCodeToDCID maps ISO 3166-1 Alpha-2 codes to Data Commons DCID's.
var CountryCodeToDCID = map[string]string{
	"AD": "country/AND",
	"AE": "country/ARE",
	"AF": "country/AFG",
	"AG": "country/ATG",
	"AI": "country/AIA",
	"AL": "country/ALB",
	"AM": "country/ARM",
	"AN": "country/ANT",
	"AO": "country/AGO",
	"AQ": "country/ATA",
	"AR": "country/ARG",
	"AS": "country/ASM",
	"AT": "country/AUT",
	"AU": "country/AUS",
	"AW": "country/ABW",
	"AZ": "country/AZE",
	"BA": "country/BIH",
	"BB": "country/BRB",
	"BD": "country/BGD",
	"BE": "country/BEL",
	"BF": "country/BFA",
	"BG": "country/BGR",
	"BH": "country/BHR",
	"BI": "country/BDI",
	"BJ": "country/BEN",
	"BM": "country/BMU",
	"BN": "country/BRN",
	"BO": "country/BOL",
	"BR": "country/BRA",
	"BS": "country/BHS",
	"BT": "country/BTN",
	"BV": "country/BVT",
	"BW": "country/BWA",
	"BY": "country/BLR",
	"BZ": "country/BLZ",
	"CA": "country/CAN",
	"CC": "country/CCK",
	"CD": "country/COD",
	"CF": "country/CAF",
	"CG": "country/COG",
	"CH": "country/CHE",
	"CI": "country/CIV",
	"CK": "country/COK",
	"CL": "country/CHL",
	"CM": "country/CMR",
	"CN": "country/CHN",
	"CO": "country/COL",
	"CR": "country/CRI",
	"CU": "country/CUB",
	"CV": "country/CPV",
	"CX": "country/CXR",
	"CY": "country/CYP",
	"CZ": "country/CZE",
	"DE": "country/DEU",
	"DJ": "country/DJI",
	"DK": "country/DNK",
	"DM": "country/DMA",
	"DO": "country/DOM",
	"DZ": "country/DZA",
	"EC": "country/ECU",
	"EE": "country/EST",
	"EG": "country/EGY",
	"EH": "country/ESH",
	// NOTE: EL is not an alpha2 code.
	// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Imperfect_implementations
	"EL": "country/GRC",
	"ER": "country/ERI",
	"ES": "country/ESP",
	"ET": "country/ETH",
	"FI": "country/FIN",
	"FJ": "country/FJI",
	"FK": "country/FLK",
	"FM": "country/FSM",
	"FO": "country/FRO",
	"FR": "country/FRA",
	"GA": "country/GAB",
	"GB": "country/GBR",
	"GD": "country/GRD",
	"GE": "country/GEO",
	"GF": "country/GUF",
	"GG": "country/GGY",
	"GH": "country/GHA",
	"GI": "country/GIB",
	"GL": "country/GRL",
	"GM": "country/GMB",
	"GN": "country/GIN",
	"GP": "country/GLP",
	"GQ": "country/GNQ",
	"GR": "country/GRC",
	"GS": "country/SGS",
	"GT": "country/GTM",
	"GU": "country/GUM",
	"GW": "country/GNB",
	"GY": "country/GUY",
	"HK": "country/HKG",
	"HM": "country/HMD",
	"HN": "country/HND",
	"HR": "country/HRV",
	"HT": "country/HTI",
	"HU": "country/HUN",
	"ID": "country/IDN",
	"IE": "country/IRL",
	"IL": "country/ISR",
	"IM": "country/IMN",
	"IN": "country/IND",
	"IO": "country/IOT",
	"IQ": "country/IRQ",
	"IR": "country/IRN",
	"IS": "country/ISL",
	"IT": "country/ITA",
	"JE": "country/JEY",
	"JM": "country/JAM",
	"JO": "country/JOR",
	"JP": "country/JPN",
	"KE": "country/KEN",
	"KG": "country/KGZ",
	"KH": "country/KHM",
	"KI": "country/KIR",
	"KM": "country/COM",
	"KN": "country/KNA",
	"KP": "country/PRK",
	"KR": "country/KOR",
	"KW": "country/KWT",
	"KY": "country/CYM",
	"KZ": "country/KAZ",
	"LA": "country/LAO",
	"LB": "country/LBN",
	"LC": "country/LCA",
	"LI": "country/LIE",
	"LK": "country/LKA",
	"LR": "country/LBR",
	"LS": "country/LSO",
	"LT": "country/LTU",
	"LU": "country/LUX",
	"LV": "country/LVA",
	"LY": "country/LBY",
	"MA": "country/MAR",
	"MC": "country/MCO",
	"MD": "country/MDA",
	"ME": "country/MNE",
	"MG": "country/MDG",
	"MH": "country/MHL",
	"MK": "country/MKD",
	"ML": "country/MLI",
	"MM": "country/MMR",
	"MN": "country/MNG",
	"MO": "country/MAC",
	"MP": "country/MNP",
	"MQ": "country/MTQ",
	"MR": "country/MRT",
	"MS": "country/MSR",
	"MT": "country/MLT",
	"MU": "country/MUS",
	"MV": "country/MDV",
	"MW": "country/MWI",
	"MX": "country/MEX",
	"MY": "country/MYS",
	"MZ": "country/MOZ",
	"NA": "country/NAM",
	"NC": "country/NCL",
	"NE": "country/NER",
	"NF": "country/NFK",
	"NG": "country/NGA",
	"NI": "country/NIC",
	"NL": "country/NLD",
	"NO": "country/NOR",
	"NP": "country/NPL",
	"NR": "country/NRU",
	"NU": "country/NIU",
	"NZ": "country/NZL",
	"OM": "country/OMN",
	"PA": "country/PAN",
	"PE": "country/PER",
	"PF": "country/PYF",
	"PG": "country/PNG",
	"PH": "country/PHL",
	"PK": "country/PAK",
	"PL": "country/POL",
	"PM": "country/SPM",
	"PN": "country/PCN",
	"PR": "country/PRI",
	"PS": "country/PSE",
	"PT": "country/PRT",
	"PW": "country/PLW",
	"PY": "country/PRY",
	"QA": "country/QAT",
	"RE": "country/REU",
	"RO": "country/ROU",
	"RS": "country/SRB",
	"RU": "country/RUS",
	"RW": "country/RWA",
	"SA": "country/SAU",
	"SB": "country/SLB",
	"SC": "country/SYC",
	"SD": "country/SDN",
	"SE": "country/SWE",
	"SG": "country/SGP",
	"SH": "country/SHN",
	"SI": "country/SVN",
	"SJ": "country/SJM",
	"SK": "country/SVK",
	"SL": "country/SLE",
	"SM": "country/SMR",
	"SN": "country/SEN",
	"SO": "country/SOM",
	"SR": "country/SUR",
	"ST": "country/STP",
	"SV": "country/SLV",
	"SY": "country/SYR",
	"SZ": "country/SWZ",
	"TC": "country/TCA",
	"TD": "country/TCD",
	"TF": "country/ATF",
	"TG": "country/TGO",
	"TH": "country/THA",
	"TJ": "country/TJK",
	"TK": "country/TKL",
	"TL": "country/TLS",
	"TM": "country/TKM",
	"TN": "country/TUN",
	"TO": "country/TON",
	"TR": "country/TUR",
	"TT": "country/TTO",
	"TV": "country/TUV",
	"TW": "country/TWN",
	"TZ": "country/TZA",
	"UA": "country/UKR",
	"UG": "country/UGA",
	// NOTE: UK is not an alpha2 code.
	// See https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2#Imperfect_implementations
	"UK": "country/GBR",
	"UM": "country/UMI",
	"US": "country/USA",
	"UY": "country/URY",
	"UZ": "country/UZB",
	"VA": "country/VAT",
	"VC": "country/VCT",
	"VE": "country/VEN",
	"VG": "country/VGB",
	"VI": "country/VIR",
	"VN": "country/VNM",
	"VU": "country/VUT",
	"WF": "country/WLF",
	"WS": "country/WSM",
	"YE": "country/YEM",
	"YT": "country/MYT",
	"ZA": "country/ZAF",
	"ZM": "country/ZMB",
	"ZW": "country/ZWE",
}

// USStateCodeToDCID maps US state two-letter identifiers to Data Commons DCIDs.
// TODO: Add mapping from standard ISO 3166-2 codes, and other instances of
// identifiers used to DCIDs.
var USStateCodeToDCID = map[string]string{
	"AL": "geoId/01",
	"AK": "geoId/02",
	"AZ": "geoId/04",
	"AR": "geoId/05",
	"CA": "geoId/06",
	"CO": "geoId/08",
	"CT": "geoId/09",
	"DE": "geoId/10",
	"DC": "geoId/11",
	"FL": "geoId/12",
	"GA": "geoId/13",
	"HI": "geoId/15",
	"ID": "geoId/16",
	"IL": "geoId/17",
	"IN": "geoId/18",
	"IA": "geoId/19",
	"KS": "geoId/20",
	"KY": "geoId/21",
	"LA": "geoId/22",
	"ME": "geoId/23",
	"MD": "geoId/24",
	"MA": "geoId/25",
	"MI": "geoId/26",
	"MN": "geoId/27",
	"MS": "geoId/28",
	"MO": "geoId/29",
	"MT": "geoId/30",
	"NE": "geoId/31",
	"NV": "geoId/32",
	"NH": "geoId/33",
	"NJ": "geoId/34",
	"NM": "geoId/35",
	"NY": "geoId/36",
	"NC": "geoId/37",
	"ND": "geoId/38",
	"OH": "geoId/39",
	"OK": "geoId/40",
	"OR": "geoId/41",
	"PA": "geoId/42",
	"RI": "geoId/44",
	"SC": "geoId/45",
	"SD": "geoId/46",
	"TN": "geoId/47",
	"TX": "geoId/48",
	"UT": "geoId/49",
	"VT": "geoId/50",
	"VA": "geoId/51",
	"WA": "geoId/53",
	"WV": "geoId/54",
	"WI": "geoId/55",
	"WY": "geoId/56",
	"AS": "geoId/60",
	"GU": "geoId/66",
	"MP": "geoId/69",
	"PR": "geoId/72",
	"UM": "geoId/74",
	"VI": "geoId/78",
}
