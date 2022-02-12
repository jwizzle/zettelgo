package zettels

import (
	"testing"
)

// Test creation, merging of configuration.
func TestCfg(t *testing.T) {
	cfg_1 := Cfg{
		Directory: "/path/of/cfg_1",
		Ignore_list: []string{
			".git",
		},
	}
	cfg_2 := Cfg{
		Directory: "/path/of/cfg_2",
		Ignore_list: []string{
			".blaat",
		},
	}

	cfg_1.Merge(cfg_2)

	if cfg_1.Directory != "/path/of/cfg_2" {
		t.Errorf("test_cfg: Dir not correctly updated: %v", cfg_1.Directory)
	}
	if len(cfg_1.Ignore_list) != len([]string{
		".git", ".blaat",
	}) {
		t.Errorf("test_cfg: Dir not correctly updated: %v", cfg_1.Directory)
	}
}
