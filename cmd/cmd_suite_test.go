package cmd

import (
	"testing"

	. "github.com/bdittmer/blocky/log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sirupsen/logrus/hooks/test"
)

var (
	fatal      bool
	loggerHook *test.Hook
)

func TestCmd(t *testing.T) {
	BeforeSuite(func() {
		Log().ExitFunc = func(int) { fatal = true }

		loggerHook = test.NewGlobal()
		Log().AddHook(loggerHook)
	})
	AfterSuite(func() {
		Log().ExitFunc = nil
		loggerHook.Reset()
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "Command Suite")
}
