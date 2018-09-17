package smoke_tests_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}

var _ = Describe("Shopping cart", func() {
	Context("initially", func() {
		It("has 3 namespaces", func() {
			// TODO
		})
		Specify("has 10 pods in kube-system namespace", func() {
			// TODO
		})
	})

	Context("When I add a new new namespace called k8s-test", func() {
		Context("the k8s cluster", func() {
			It("has 1 more namespace", func() {
				// TODO
			})
			Specify("Then namespace k8s-test should be accessible listed", func() {
				// TODO
			})
		})
	})

	Context("When I add a new nignx load balancer service to k8s-test namespace", func() {
		Context("the k8s cluster", func() {
			It("has 1 more pod in k8s-test", func() {
				// TODO
			})
			It("has 1 more deployment in k8s-test", func() {
				// TODO
			})
			It("has 1 more service in k8s-test", func() {
				// TODO
			})
			Specify("Then the nginx service should be accessible", func() {
				// TODO
			})
		})
	})
})
