package smoke_tests_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestSmokeTests(t *testing.T) {

	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("smoke-test-report.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Smoke Tests", []Reporter{junitReporter})
}
