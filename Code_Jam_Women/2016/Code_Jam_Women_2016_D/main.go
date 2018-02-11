package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var splitAlphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var impossibleCombinations = [][]string{
	{"AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU", "AV", "AW", "AX", "AY", "AZ", "BA", "CA", "DA", "EA", "FA", "GA", "HA", "IA", "JA", "KA", "LA", "MA", "NA", "OA", "PA", "QA", "RA", "SA", "TA", "UA", "VA", "WA", "XA", "YA", "ZA"},
	{"BA", "BC", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BK", "BL", "BM", "BN", "BO", "BP", "BQ", "BR", "BS", "BT", "BU", "BV", "BW", "BX", "BY", "BZ", "AB", "CB", "DB", "EB", "FB", "GB", "HB", "IB", "JB", "KB", "LB", "MB", "NB", "OB", "PB", "QB", "RB", "SB", "TB", "UB", "VB", "WB", "XB", "YB", "ZB"},
	{"CA", "CB", "CD", "CE", "CF", "CG", "CH", "CI", "CJ", "CK", "CL", "CM", "CN", "CO", "CP", "CQ", "CR", "CS", "CT", "CU", "CV", "CW", "CX", "CY", "CZ", "AC", "BC", "DC", "EC", "FC", "GC", "HC", "IC", "JC", "KC", "LC", "MC", "NC", "OC", "PC", "QC", "RC", "SC", "TC", "UC", "VC", "WC", "XC", "YC", "ZC"},
	{"DA", "DB", "DC", "DE", "DF", "DG", "DH", "DI", "DJ", "DK", "DL", "DM", "DN", "DO", "DP", "DQ", "DR", "DS", "DT", "DU", "DV", "DW", "DX", "DY", "DZ", "AD", "BD", "CD", "ED", "FD", "GD", "HD", "ID", "JD", "KD", "LD", "MD", "ND", "OD", "PD", "QD", "RD", "SD", "TD", "UD", "VD", "WD", "XD", "YD", "ZD"},
	{"EA", "EB", "EC", "ED", "EF", "EG", "EH", "EI", "EJ", "EK", "EL", "EM", "EN", "EO", "EP", "EQ", "ER", "ES", "ET", "EU", "EV", "EW", "EX", "EY", "EZ", "AE", "BE", "CE", "DE", "FE", "GE", "HE", "IE", "JE", "KE", "LE", "ME", "NE", "OE", "PE", "QE", "RE", "SE", "TE", "UE", "VE", "WE", "XE", "YE", "ZE"},
	{"FA", "FB", "FC", "FD", "FE", "FG", "FH", "FI", "FJ", "FK", "FL", "FM", "FN", "FO", "FP", "FQ", "FR", "FS", "FT", "FU", "FV", "FW", "FX", "FY", "FZ", "AF", "BF", "CF", "DF", "EF", "GF", "HF", "IF", "JF", "KF", "LF", "MF", "NF", "OF", "PF", "QF", "RF", "SF", "TF", "UF", "VF", "WF", "XF", "YF", "ZF"},
	{"GA", "GB", "GC", "GD", "GE", "GF", "GH", "GI", "GJ", "GK", "GL", "GM", "GN", "GO", "GP", "GQ", "GR", "GS", "GT", "GU", "GV", "GW", "GX", "GY", "GZ", "AG", "BG", "CG", "DG", "EG", "FG", "HG", "IG", "JG", "KG", "LG", "MG", "NG", "OG", "PG", "QG", "RG", "SG", "TG", "UG", "VG", "WG", "XG", "YG", "ZG"},
	{"HA", "HB", "HC", "HD", "HE", "HF", "HG", "HI", "HJ", "HK", "HL", "HM", "HN", "HO", "HP", "HQ", "HR", "HS", "HT", "HU", "HV", "HW", "HX", "HY", "HZ", "AH", "BH", "CH", "DH", "EH", "FH", "GH", "IH", "JH", "KH", "LH", "MH", "NH", "OH", "PH", "QH", "RH", "SH", "TH", "UH", "VH", "WH", "XH", "YH", "ZH"},
	{"IA", "IB", "IC", "ID", "IE", "IF", "IG", "IH", "IJ", "IK", "IL", "IM", "IN", "IO", "IP", "IQ", "IR", "IS", "IT", "IU", "IV", "IW", "IX", "IY", "IZ", "AI", "BI", "CI", "DI", "EI", "FI", "GI", "HI", "JI", "KI", "LI", "MI", "NI", "OI", "PI", "QI", "RI", "SI", "TI", "UI", "VI", "WI", "XI", "YI", "ZI"},
	{"JA", "JB", "JC", "JD", "JE", "JF", "JG", "JH", "JI", "JK", "JL", "JM", "JN", "JO", "JP", "JQ", "JR", "JS", "JT", "JU", "JV", "JW", "JX", "JY", "JZ", "AJ", "BJ", "CJ", "DJ", "EJ", "FJ", "GJ", "HJ", "IJ", "KJ", "LJ", "MJ", "NJ", "OJ", "PJ", "QJ", "RJ", "SJ", "TJ", "UJ", "VJ", "WJ", "XJ", "YJ", "ZJ"},
	{"KA", "KB", "KC", "KD", "KE", "KF", "KG", "KH", "KI", "KJ", "KL", "KM", "KN", "KO", "KP", "KQ", "KR", "KS", "KT", "KU", "KV", "KW", "KX", "KY", "KZ", "AK", "BK", "CK", "DK", "EK", "FK", "GK", "HK", "IK", "JK", "LK", "MK", "NK", "OK", "PK", "QK", "RK", "SK", "TK", "UK", "VK", "WK", "XK", "YK", "ZK"},
	{"LA", "LB", "LC", "LD", "LE", "LF", "LG", "LH", "LI", "LJ", "LK", "LM", "LN", "LO", "LP", "LQ", "LR", "LS", "LT", "LU", "LV", "LW", "LX", "LY", "LZ", "AL", "BL", "CL", "DL", "EL", "FL", "GL", "HL", "IL", "JL", "KL", "ML", "NL", "OL", "PL", "QL", "RL", "SL", "TL", "UL", "VL", "WL", "XL", "YL", "ZL"},
	{"MA", "MB", "MC", "MD", "ME", "MF", "MG", "MH", "MI", "MJ", "MK", "ML", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV", "MW", "MX", "MY", "MZ", "AM", "BM", "CM", "DM", "EM", "FM", "GM", "HM", "IM", "JM", "KM", "LM", "NM", "OM", "PM", "QM", "RM", "SM", "TM", "UM", "VM", "WM", "XM", "YM", "ZM"},
	{"NA", "NB", "NC", "ND", "NE", "NF", "NG", "NH", "NI", "NJ", "NK", "NL", "NM", "NO", "NP", "NQ", "NR", "NS", "NT", "NU", "NV", "NW", "NX", "NY", "NZ", "AN", "BN", "CN", "DN", "EN", "FN", "GN", "HN", "IN", "JN", "KN", "LN", "MN", "ON", "PN", "QN", "RN", "SN", "TN", "UN", "VN", "WN", "XN", "YN", "ZN"},
	{"OA", "OB", "OC", "OD", "OE", "OF", "OG", "OH", "OI", "OJ", "OK", "OL", "OM", "ON", "OP", "OQ", "OR", "OS", "OT", "OU", "OV", "OW", "OX", "OY", "OZ", "AO", "BO", "CO", "DO", "EO", "FO", "GO", "HO", "IO", "JO", "KO", "LO", "MO", "NO", "PO", "QO", "RO", "SO", "TO", "UO", "VO", "WO", "XO", "YO", "ZO"},
	{"PA", "PB", "PC", "PD", "PE", "PF", "PG", "PH", "PI", "PJ", "PK", "PL", "PM", "PN", "PO", "PQ", "PR", "PS", "PT", "PU", "PV", "PW", "PX", "PY", "PZ", "AP", "BP", "CP", "DP", "EP", "FP", "GP", "HP", "IP", "JP", "KP", "LP", "MP", "NP", "OP", "QP", "RP", "SP", "TP", "UP", "VP", "WP", "XP", "YP", "ZP"},
	{"QA", "QB", "QC", "QD", "QE", "QF", "QG", "QH", "QI", "QJ", "QK", "QL", "QM", "QN", "QO", "QP", "QR", "QS", "QT", "QU", "QV", "QW", "QX", "QY", "QZ", "AQ", "BQ", "CQ", "DQ", "EQ", "FQ", "GQ", "HQ", "IQ", "JQ", "KQ", "LQ", "MQ", "NQ", "OQ", "PQ", "RQ", "SQ", "TQ", "UQ", "VQ", "WQ", "XQ", "YQ", "ZQ"},
	{"RA", "RB", "RC", "RD", "RE", "RF", "RG", "RH", "RI", "RJ", "RK", "RL", "RM", "RN", "RO", "RP", "RQ", "RS", "RT", "RU", "RV", "RW", "RX", "RY", "RZ", "AR", "BR", "CR", "DR", "ER", "FR", "GR", "HR", "IR", "JR", "KR", "LR", "MR", "NR", "OR", "PR", "QR", "SR", "TR", "UR", "VR", "WR", "XR", "YR", "ZR"},
	{"SA", "SB", "SC", "SD", "SE", "SF", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SP", "SQ", "SR", "ST", "SU", "SV", "SW", "SX", "SY", "SZ", "AS", "BS", "CS", "DS", "ES", "FS", "GS", "HS", "IS", "JS", "KS", "LS", "MS", "NS", "OS", "PS", "QS", "RS", "TS", "US", "VS", "WS", "XS", "YS", "ZS"},
	{"TA", "TB", "TC", "TD", "TE", "TF", "TG", "TH", "TI", "TJ", "TK", "TL", "TM", "TN", "TO", "TP", "TQ", "TR", "TS", "TU", "TV", "TW", "TX", "TY", "TZ", "AT", "BT", "CT", "DT", "ET", "FT", "GT", "HT", "IT", "JT", "KT", "LT", "MT", "NT", "OT", "PT", "QT", "RT", "ST", "UT", "VT", "WT", "XT", "YT", "ZT"},
	{"UA", "UB", "UC", "UD", "UE", "UF", "UG", "UH", "UI", "UJ", "UK", "UL", "UM", "UN", "UO", "UP", "UQ", "UR", "US", "UT", "UV", "UW", "UX", "UY", "UZ", "AU", "BU", "CU", "DU", "EU", "FU", "GU", "HU", "IU", "JU", "KU", "LU", "MU", "NU", "OU", "PU", "QU", "RU", "SU", "TU", "VU", "WU", "XU", "YU", "ZU"},
	{"VA", "VB", "VC", "VD", "VE", "VF", "VG", "VH", "VI", "VJ", "VK", "VL", "VM", "VN", "VO", "VP", "VQ", "VR", "VS", "VT", "VU", "VW", "VX", "VY", "VZ", "AV", "BV", "CV", "DV", "EV", "FV", "GV", "HV", "IV", "JV", "KV", "LV", "MV", "NV", "OV", "PV", "QV", "RV", "SV", "TV", "UV", "WV", "XV", "YV", "ZV"},
	{"WA", "WB", "WC", "WD", "WE", "WF", "WG", "WH", "WI", "WJ", "WK", "WL", "WM", "WN", "WO", "WP", "WQ", "WR", "WS", "WT", "WU", "WV", "WX", "WY", "WZ", "AW", "BW", "CW", "DW", "EW", "FW", "GW", "HW", "IW", "JW", "KW", "LW", "MW", "NW", "OW", "PW", "QW", "RW", "SW", "TW", "UW", "VW", "XW", "YW", "ZW"},
	{"XA", "XB", "XC", "XD", "XE", "XF", "XG", "XH", "XI", "XJ", "XK", "XL", "XM", "XN", "XO", "XP", "XQ", "XR", "XS", "XT", "XU", "XV", "XW", "XY", "XZ", "AX", "BX", "CX", "DX", "EX", "FX", "GX", "HX", "IX", "JX", "KX", "LX", "MX", "NX", "OX", "PX", "QX", "RX", "SX", "TX", "UX", "VX", "WX", "YX", "ZX"},
	{"YA", "YB", "YC", "YD", "YE", "YF", "YG", "YH", "YI", "YJ", "YK", "YL", "YM", "YN", "YO", "YP", "YQ", "YR", "YS", "YT", "YU", "YV", "YW", "YX", "YZ", "AY", "BY", "CY", "DY", "EY", "FY", "GY", "HY", "IY", "JY", "KY", "LY", "MY", "NY", "OY", "PY", "QY", "RY", "SY", "TY", "UY", "VY", "WY", "XY", "ZY"},
	{"ZA", "ZB", "ZC", "ZD", "ZE", "ZF", "ZG", "ZH", "ZI", "ZJ", "ZK", "ZL", "ZM", "ZN", "ZO", "ZP", "ZQ", "ZR", "ZS", "ZT", "ZU", "ZV", "ZW", "ZX", "ZY", "AZ", "BZ", "CZ", "DZ", "EZ", "FZ", "GZ", "HZ", "IZ", "JZ", "KZ", "LZ", "MZ", "NZ", "OZ", "PZ", "QZ", "RZ", "SZ", "TZ", "UZ", "VZ", "WZ", "XZ", "YZ"},
}

func main() {
	input := make([]string, 0)
	output := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numberOfCases, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	for i := 0; i < numberOfCases; i++ {
		scanner.Scan()
		text := scanner.Text()
		scanner.Scan()
		text += " " + scanner.Text()
		input = append(input, text)

	}
	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	for i, v := range input {
		output = append(output, getSolution(i+1, v))
	}

	for _, v := range output {
		fmt.Println(v)
	}

}

func getSolution(caseNumber int, input string) string {
	output := "Case #" + strconv.Itoa(caseNumber) + ": "

	inputs := strings.Split(input, " ")

	nbOfWords, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic(err)
	}

	passwords := inputs[1:]

	if nbOfWords != len(passwords) {
		panic("Error in the input")
	}

	output += getPossibleAlphabet(nbOfWords, passwords, alphabet)

	return output
}

func getPossibleAlphabet(nbOfWords int, passwords []string, initialOutput string) string {

	result := initialOutput

	if checkImpossibleCombinations(nbOfWords, passwords) {
		return "IMPOSSIBLE"
	}

	for i, v := range passwords {
		if len(v) == 1 {
			result = "IMPOSSIBLE"
			break
		}
		if lettersRepeated(v) {
			passwords = append(passwords[:i], passwords[i+1:]...)
			continue
		}
		if !strings.Contains(result, v) {
			continue
		}

		result = checkNewCombinations(v, passwords, result)
		if result == "IMPOSSIBLE" {
			break
		}
	}

	return result
}

func lettersRepeated(word string) bool {
	repeatedLetters := false

	for i := 0; i < len(alphabet); i++ {
		if strings.Count(word, alphabet[i:i+1]) > 1 {
			repeatedLetters = true
			break
		}
	}

	return repeatedLetters
}

func checkNewCombinations(password string, passwordArray []string, incorrectSolution string) string {

	resultFound := false
	counter := 0
	result := ""
	resultStrArray := splitAlphabet
	for !resultFound && counter <= 10000 {
		resultStrArray = shuffle(resultStrArray)
		result = strings.Join(resultStrArray, "")

		resultFound = true

		for _, v := range passwordArray {
			if strings.Contains(result, v) {
				resultFound = false
				break
			}
		}
		counter++
	}

	if !resultFound {
		result = "IMPOSSIBLE"
	}

	return result

}

func shuffle(src []string) []string {
	final := make([]string, len(src))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}
	return final
}

func checkImpossibleCombinations(nbOfWords int, passwords []string) bool {

	impossible := false

	if nbOfWords >= 50 {
		for _, combination := range impossibleCombinations {
			impossible = true
			for _, neededForCombination := range combination {
				if !contains(passwords, neededForCombination) {
					impossible = false
					break
				}
			}
			if impossible {
				break
			}
		}
	}

	return impossible

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
