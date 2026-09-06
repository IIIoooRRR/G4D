package api_test

import (
	"fmt"
	"testing"

	"github.com/IIIoooRRR/G4D/api"
)

const (
	path1 = "/channel/"
	path2 = "messages/"
)

var GarbageString string

func BenchmarkGetURI(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	id := "1234567"
	msgId := "1234567876543"
	for i := 0; i < b.N; i++ {
		GarbageString = api.GetURI(path1, id, path2, msgId)
	}
}
func BenchmarkSprintF(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	id := "1234567"
	msgId := "1234567876543"
	for i := 0; i < b.N; i++ {
		GarbageString = fmt.Sprintf("/channel/%s/messages/%s", id, msgId)
	}
}
func BenchmarkGetURL(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	host := "https://discord.com"
	path := "/channels/123456789012345678/messages"
	for i := 0; i < b.N; i++ {
		GarbageString = api.GetURI(host, path)
	}
}
func BenchmarkConcatination(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	host := "https://discord.com"
	path := "/channels/123456789012345678/messages"
	for i := 0; i < b.N; i++ {
		GarbageString = host + path
	}
}
