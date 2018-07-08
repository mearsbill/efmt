// Copyright (c) 2018 William G. Mears
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
// Author Contact  opensource@imears.com

// This package provides extened formating of strings ... including...
// converts float to scientific notation with prefixes ...  AND SUPPORTS Mega= 1024 * 1024
// ex: efmt.Float2Str(.0005,"7.2","Sec",false) => " 500.00 uSec"
// ex: efmt.Float2Str(1010 ,"7.2","Sec",false) => "   1.01 KSec"

package efmt

import (
    "fmt"
)

func Float2Str(val float64, sfmt string, units string,megaNotMillion bool) (string) {
    unitsTable := []byte("<yzafpnum KMGTPEZY>")  // each position worth 1000
    unitsZero := ((len(unitsTable)+1)/2)-1
    // first normalize the float to be between 1.000 and 999.99
    unitsIndex := 0 // the "center of the range"
    unitsValue := 1000.0
    if megaNotMillion { unitsValue = 1024.0}
    switch {
    case val >= unitsValue:
        for val >= unitsValue {
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

// Converts ints to strings using scientifc notations with prefix.
func Int2Str(val int64, sfmt string, units string,megaNotMillion bool) (string) {
    return Float2Str(float64(val),sfmt,units,megaNotMillion)
}
