package sense

import (
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ToStringByPrecise(bigNum *big.Int, decimals *big.Int) string {
	result := ""
	destStr := bigNum.String()
	destLen := uint64(len(destStr))
	precise := decimals.Uint64()
	if precise >= destLen { // add "0.000..." at former of destStr
		var i uint64 = 0
		prefix := "0."
		for ; i < precise-destLen; i++ {
			prefix += "0"
		}
		result = prefix + destStr
	} else { // add "."
		pointIndex := destLen - precise
		result = destStr[0:pointIndex] + "." + destStr[pointIndex:]
	}
	result = removeZeroAtTail(result)
	return result
}

// delete no need "0" at last of result
func removeZeroAtTail(str string) string {
	i := len(str) - 1
	for ; i >= 0; i-- {
		if str[i] != '0' {
			break
		}
	}
	str = str[:i+1]
	// delete "." at last of result
	if str[len(str)-1] == '.' {
		str = str[:len(str)-1]
	}
	return str
}

func ToIntByPrecise(str string, decimals *big.Int) *big.Int {
	result := new(big.Int)
	splits := strings.Split(str, ".")
	precise := decimals.Uint64()
	if len(splits) == 1 { // doesn't contain "."
		var i uint64 = 0
		for ; i < precise; i++ {
			str += "0"
		}
		intValue, ok := new(big.Int).SetString(str, 10)
		if ok {
			result.Set(intValue)
		}
	} else if len(splits) == 2 {
		value := new(big.Int)
		ok := false
		floatLen := uint64(len(splits[1]))
		if floatLen <= precise { // add "0" at last of str
			parseString := strings.Replace(str, ".", "", 1)
			var i uint64 = 0
			for ; i < precise-floatLen; i++ {
				parseString += "0"
			}
			value, ok = value.SetString(parseString, 10)
		} else { // remove redundant digits after "."
			splits[1] = splits[1][:precise]
			parseString := splits[0] + splits[1]
			value, ok = value.SetString(parseString, 10)
		}
		if ok {
			result.Set(value)
		}
	}

	return result
}

// go test -v github.com/dylenfu/go-libs/sense -run TestToIntPrecials
func TestToIntPrecials(t *testing.T) {
	testStr := []string{
		"1111",
		"111.2333",
		"11.0000000000000000000000000000000000000",
		"0238462326847234.61637644",
		"39119032.1111111111111111111111111111199999999999999999999999",
		"132716.0000002734824800000000",
		"-1111",
		"-111.2333",
		"-11.0000000000000000000000000000000000000",
		"-0238462326847234.61637644",
		"-39119032.1111111111111111111111111111199999999999999999999999",
		"-132716.0000002734824800000000",
	}
	testDecimals := []*big.Int{
		big.NewInt(0),
		big.NewInt(10),
		big.NewInt(18),
		big.NewInt(50),
	}
	expectStr := [12][4]string{
		{"1111", "1111", "1111", "1111"},
		{"111", "111.2333", "111.2333", "111.2333"},
		{"11", "11", "11", "11"},
		{"238462326847234", "238462326847234.61637644", "238462326847234.61637644", "238462326847234.61637644"},
		{"39119032", "39119032.1111111111", "39119032.111111111111111111",
			"39119032.11111111111111111111111111111999999999999999999999"},
		{"132716", "132716.0000002734", "132716.00000027348248", "132716.00000027348248"},
		{"-1111", "-1111", "-1111", "-1111"},
		{"-111", "-111.2333", "-111.2333", "-111.2333"},
		{"-11", "-11", "-11", "-11"},
		{"-238462326847234", "-238462326847234.61637644", "-238462326847234.61637644", "-238462326847234.61637644"},
		{"-39119032", "-39119032.1111111111", "-39119032.111111111111111111",
			"-39119032.11111111111111111111111111111999999999999999999999"},
		{"-132716", "-132716.0000002734", "-132716.00000027348248", "-132716.00000027348248"},
	}
	for i, str := range testStr {
		for j, decimals := range testDecimals {
			parsedInt := ToIntByPrecise(str, decimals)
			//t.Logf("parsed int is %d", parsedInt)
			parsedStr := ToStringByPrecise(parsedInt, decimals)
			//t.Logf("parsed str is \"%s\",", parsedStr)
			t.Logf("init str:%s, decimals:%d, parsed int:%d, parsed string:%s", str, decimals, parsedInt, parsedStr)
			assert.Equal(t, expectStr[i][j], parsedStr)
		}
	}
}
