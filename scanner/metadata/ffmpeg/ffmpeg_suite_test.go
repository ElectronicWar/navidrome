package ffmpeg

import (
	"testing"

	"github.com/navidrome/navidrome/log"
	"github.com/navidrome/navidrome/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFFMpeg(t *testing.T) {
	tests.Init(t, true)
	log.SetLevel(log.LevelCritical)
	RegisterFailHandler(Fail)
	RunSpecs(t, "FFMpeg Suite")
}
