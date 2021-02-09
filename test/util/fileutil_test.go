package util

import (
	"co/iiq/i/notification-server/src/util"
	"testing"
)

func TestReadPropertiesFile(t *testing.T) {
	props, err := util.ReadPropertiesFile("../resources/test.properties")
	if err != nil {
		t.Error("Error while reading properties file")
	}

	if err == nil && (props["host"] != "localhost" || props["proxyHost"] != "test" || props["protocol"] != "https://" || props["chunk"] != "") {
		t.Error("Error properties not loaded correctly")
	}
}
