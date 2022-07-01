package main

import (
	"github.com/Masterminds/semver/v3"
	"testing"
)

func BenchmarkSemVer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sv, _ := semver.NewVersion("1.2.3-beta.1+build345")
		sv2, _ := semver.NewVersion("2")
		sv.GreaterThan(sv2)
	}
}
