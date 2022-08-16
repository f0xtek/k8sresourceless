package logger

import (
	"strings"

	"github.com/rs/zerolog/log"
)

type NoResourceMsg struct {
	CpuMissing, MemMissing          bool
	PodNamespace, PodName, PodOwner string
	Content                         string
}

func (n *NoResourceMsg) format() {
	var msg string
	if n.CpuMissing && n.MemMissing {
		msg = "no cpu or memory resource requests defined"
	} else if n.CpuMissing {
		msg = "no cpu resource requests defined"
	} else if n.MemMissing {
		msg = "no memory resource requests defined"
	}
	n.Content = strings.ReplaceAll(msg, "\"", "")
}

func (n *NoResourceMsg) Log() {
	n.format()
	log.Info().Str("namespace", n.PodNamespace).Str("name", n.PodName).Str("owner", n.PodOwner).Msg(n.Content)
}
