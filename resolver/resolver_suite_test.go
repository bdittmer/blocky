package resolver_test

import (
	"testing"

	. "github.com/bdittmer/blocky/log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestResolver(t *testing.T) {
	ConfigureLogger("Warn", "text", true)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Resolver Suite")
}
