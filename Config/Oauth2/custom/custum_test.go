package custom

import "testing"

func TestCustomOauth_ReadConfig(t *testing.T) {
	testconfig := NewCustom()
	testconfig.ReadConfig("../oauth2.yaml")

}
