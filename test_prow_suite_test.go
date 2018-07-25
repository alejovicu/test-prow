package test_prow_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTestProw(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestProw Suite")
}
