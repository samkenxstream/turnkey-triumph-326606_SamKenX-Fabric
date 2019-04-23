/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package ledgerconfig

import (
	"testing"

	ledgertestutil "github.com/hyperledger/fabric/core/ledger/testutil"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetTotalLimitDefault(t *testing.T) {
	setUpCoreYAMLConfig()
	defaultValue := GetTotalQueryLimit()
	assert.Equal(t, 100000, defaultValue) //test default config is 100000
}

func TestGetTotalLimitUnset(t *testing.T) {
	viper.Reset()
	defaultValue := GetTotalQueryLimit()
	assert.Equal(t, 10000, defaultValue) //test default config is 1000
}

func TestGetTotalLimit(t *testing.T) {
	setUpCoreYAMLConfig()
	defer ledgertestutil.ResetConfigToDefaultValues()
	viper.Set("ledger.state.totalQueryLimit", 5000)
	updatedValue := GetTotalQueryLimit()
	assert.Equal(t, 5000, updatedValue) //test config returns 5000
}

func setUpCoreYAMLConfig() {
	//call a helper method to load the core.yaml
	ledgertestutil.SetupCoreYAMLConfig()
}
