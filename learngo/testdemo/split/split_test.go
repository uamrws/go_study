package split

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

/*
用来检测代码覆盖率的脚本
!#/bin/sh

set -e -u

go test -coverprofile cover.out -covermode atomic

perc=`go tool cover -func=cover.out | tail -n 1 | sed -Ee 's!^[^[:digit:]]+([[:digit:]]+(\.[[:digit:]]+)?)%$!\1!'`
res=`echo "$perc >= 80.0" | bc`
test "$res" -eq 1 && exit 0
echo "Insufficient coverage: $perc" >&2
exit 1
*/
func TestSplit(t *testing.T) {
	tests := map[string]struct {
		a, b string
		c    []string
	}{
		"normal": {"a,b", ",", []string{"a", "b"}},
		"multi":  {"a:b:c", ":", []string{"a", "b", "c"}},
		"double": {"a::b::c", "::", []string{"a", "b", "c"}},
		"triple": {"a:::b:::c:::", ":::", []string{"a", "b", "c", ""}},
		"none":   {"a:b:c", "", []string{"a:b:c"}},
		"other":  {"a:b:c", "*", []string{"a:b:c"}},
		"hanz":   {"中一国一", "一", []string{"中", "国", ""}},
	}

	for name, ts := range tests {
		t.Run(name, func(t *testing.T) {
			if actual := Split(ts.a, ts.b); !reflect.DeepEqual(actual, ts.c) {
				t.Errorf("Split(%#v, %#v) expect %#v,got %#v\n", ts.a, ts.b, ts.c, actual)
			}
		})

	}
}
func BenchmarkSplit(b *testing.B) {
	tests := map[string]struct {
		a, b string
		c    []string
	}{
		"normal": {"a,b", ",", []string{"a", "b"}},
		"multi":  {"a:b:c", ":", []string{"a", "b", "c"}},
		"double": {"a::b::c", "::", []string{"a", "b", "c"}},
		"triple": {"a:::b:::c:::", ":::", []string{"a", "b", "c", ""}},
		"none":   {"a:b:c", "", []string{"a:b:c"}},
		"other":  {"a:b:c", "*", []string{"a:b:c"}},
		"hanz":   {"中一国一", "一", []string{"中", "国", ""}},
	}
	b.ResetTimer()
	for name, ts := range tests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Split(ts.a, ts.b)
			}
		})
	}
}

func BenchmarkSplitParallel(b *testing.B) {
	tests := map[string]struct {
		a, b string
		c    []string
	}{
		"normal": {"a,b", ",", []string{"a", "b"}},
		"multi":  {"a:b:c", ":", []string{"a", "b", "c"}},
		"double": {"a::b::c", "::", []string{"a", "b", "c"}},
		"triple": {"a:::b:::c:::", ":::", []string{"a", "b", "c", ""}},
		"none":   {"a:b:c", "", []string{"a:b:c"}},
		"other":  {"a:b:c", "*", []string{"a:b:c"}},
		"hanz":   {"中一国一", "一", []string{"中", "国", ""}},
	}
	for _, ts := range tests {
		//设置cpu使用数
		//b.SetParallelism(1)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				Split(ts.a, ts.b)
			}
		})
	}
}
func ExampleSplit() {
	fmt.Println(Split("a,b", ","))
	// Output: [a b]
}
func TestMain(m *testing.M) {
	fmt.Println("a")
	exitCode := m.Run()
	fmt.Println("coverage", testing.Coverage())
	os.Exit(exitCode)
}
