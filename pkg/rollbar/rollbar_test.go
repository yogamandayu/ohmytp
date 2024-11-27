package rollbar_test

import (
	"testing"

	"github.com/yogamandayu/ohmytp/consts"
	"github.com/yogamandayu/ohmytp/tests"
)

func TestRollbar(t *testing.T) {

	testSuite := tests.NewTestSuite()
	defer func() {
		t.Cleanup(testSuite.Clean)
	}()
	testSuite.LoadApp()
	testSuite.App.Rollbar.Message(consts.RollbarSeverityLevelInfo.String(), "Hello World!")
}
