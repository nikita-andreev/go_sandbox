package serverconnector_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "go_sandbox/serverconnector"
	"testing"
)

func TestServerConnector(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Connector Suite")
}

var _ = Describe("My first test", func() {
	It("QQQ", func() { Expect(A(3)).To(Equal(3)) })
})
