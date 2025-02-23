package stats

import (
	"testing"

	. "github.com/bdittmer/blocky/log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStats(t *testing.T) {
	ConfigureLogger("Warn", "text", true)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stats Suite")
}
