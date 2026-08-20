package certification

import "testing"

func TestVoltForge16(t *testing.T) {
	index := NewReportRestoreIndex()
	index.Restore("PPS", "phone-1")
	if got := index.Load("PPS"); len(got) != 1 || got[0] != "phone-1" {
		t.Fatalf("got %#v", got)
	}
}
