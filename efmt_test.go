package efmt

import (
	"testing"
)


// Baseline testing
func TestFloat2Str(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
    tests := []struct {
        name string
        x       float64
        sfmt     string
        units   string
        mega    bool
        want string
    }{
        {"No prefix", 1.2346, "5.3","Sec",false,"1.235 Sec "},
        {"High Edge", 1000.0, "5.3","Sec",false,"1.000 KSec"},
        {"Mega Test", 1024.0, "5.3","Sec",true,"1.000 KSec"},
        {"nano Test", 1e-9, "5.3","Sec",false,"1.000 nSec"},
        {"giga Test", 1.05e+9, "5.3","Sec",false,"1.050 GSec"},
        {"yotta Test", 1e+24, "5.3","Sec",false,"1.000 YSec"},
        {"yocto Test", 1.05e-24, "5.3","Sec",false,"1.050 ySec"},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Float2Str(tt.x,tt.sfmt,tt.units,tt.mega); got != tt.want {
                t.Errorf("Float2Str() = \"%v\", want \"%v\"", got, tt.want)
            }
        })
    }


}
func TestInt2Str(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
    tests := []struct {
        name string
        x       int64
        sfmt     string
        units   string
        mega    bool
        want string
    }{
        {"No prefix", 987, "3.0","Sec",false,"987 Sec "},
        {"K prefix", 9870, "3.0","Sec",false," 10 KSec"},
        {"K.2 prefix", 98700, "5.2","Sec",false,"98.70 KSec"},

    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Int2Str(tt.x,tt.sfmt,tt.units,tt.mega); got != tt.want {
                t.Errorf("Float2Str() = \"%v\", want \"%v\"", got, tt.want)
            }
        })
    }


}

func BenchmarkFloat2Str1(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Float2Str(2.01,"6.2","Units",false)
    }
}
func BenchmarkFloat2Str2(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Float2Str(201000000.0,"6.2","Units",false)
    }
}

func benchmarkFloat2Str(val float64, b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Float2Str(val,"6.2","Units",false)
    }
}
