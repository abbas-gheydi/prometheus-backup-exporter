package config

import "testing"

var configuration Configurations

func TestReadConf(t *testing.T) {
	err := ReadConfs(&configuration, "config-example")
	if err != nil {
		t.Error("could not read config file")
	}

	if configuration.Server.Port != "88088"{
	t.Error ("could not get server port from config file")
	}

	if configuration.DailyBackups[0] != "/just/4/test"{
	t.Error("colud not get backup paths from config files")
	}

}



