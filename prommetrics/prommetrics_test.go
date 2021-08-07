package prommetrics

import (
	"testing"
)

func TestHelp(t *testing.T) {
	help, _ := Generate("backup_date", "backup", "/home/test/test.tar", "2020130202141414")
	if help != "#TYPE backup_date gauge\n" {
		t.Error("help format error", help)
	}

}

func TestMetric(t *testing.T) {
	_, metric := Generate("backup_date", "backup", "/home/test/test.tar", "2020130202141414")
	if metric != "backup_date{backup=\"/home/test/test.tar\"} 2020130202141414\n" {
		t.Error("metric format error", metric)
	}

}
