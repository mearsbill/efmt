// Copyright  (C)  2018 William Mears
// Simple package to perform numberical conversion to strings
//
// wmears@imears.com
//
package efmt

import (
    "fmt"
)

// converts float with units to string...  AND SUPPORT Mega= 1024 * 1024
// ex: efmt.Float2Str(.0005,"7.2","Sec",false) => " 500.00 uSec"
// ex: efmt.Float2Str(1010 ,"7.2","Sec",false) => "   1.01 KSec"
func Float2Str(val float64, sfmt string, units string,megaNotMillion bool) (string) {
    unitsTable := []byte("<yzafpnum KMGTPEZY>")  // each position worth 1000
    unitsZero := ((len(unitsTable)+1)/2)-1
    // first normalize the float to be between 1.000 and 999.99
    unitsIndex := 0 // the "center of the range"
    unitsValue := 1000.0
    if megaNotMillion { unitsValue = 1024.0}
    switch {
    case val > unitsValue:
        for val > unitsValue {
            unitsIndex++
            val /= unitsValue
        }
    case val < 1.0:
        for val <1.0 {
            unitsIndex--
            val *= unitsValue
        }
    }
    fmtStr := fmt.Sprintf("%%%sf",sfmt)
    mantissaStr := fmt.Sprintf(fmtStr,val)
//    fmt.Printf("format=%s MANTISSA = %s, unitsIndex:%d \n",fmtStr,mantissaStr,unitsIndex)
    // special case while keeping string length constant
    if unitsIndex == 0 {
        return fmt.Sprintf("%s %s ",mantissaStr,units)
    } else {
        idx := unitsZero + unitsIndex;
        if idx < 0 { idx=0 }
        if idx >len(unitsTable)-1 {idx = len(unitsTable)-1}
        return fmt.Sprintf("%s %c%s",mantissaStr,unitsTable[idx],units)
    }
}

func IntToStr(val int64, sfmt string, units string,megaNotMillion bool) (string) {
    return Float2Str(float64(val),sfmt,units,megaNotMillion)
}
