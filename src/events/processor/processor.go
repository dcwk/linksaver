package processor

import (
	"github.com/dcwk/linksaver/src/base/telegram"
)

type Processor struct {
	tg     *telegram.Client
	offset int
}
