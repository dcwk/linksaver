package processor

import (
	"github.com/dcwk/linksaver/src/base/telegram"
	"github.com/dcwk/linksaver/src/events"
	e "github.com/dcwk/linksaver/src/infrastructure/error"
	"github.com/dcwk/linksaver/src/infrastructure/storage"
	"github.com/docker/docker/api/types/events"
)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

type Meta struct {
	ChatID   int
	UserName string
}

func New(client *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg:      client,
		storage: storage,
	}
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("Can't get events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	events := make([]events.Event, 0, len(updates))
	for _, update := range updates {
		events = append(events, event(update))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return events, nil
}

func event(update telegram.Update) events.Event {
	updateType := fetchType(update)
	event := events.Event{
		Type: updateType,
		Text: fetchText(update),
	}

	if updateType == events.Message {
		event.Meta = Meta{
			ChatID:   update.Message.Chat.Id,
			UserName: update.Message.From.Username,
		}
	}

	return event
}

func fetchText(update telegram.Update) string {
	if update.Message == nil {
		return ""
	}

	return update.Message.Text
}

func fetchType(update telegram.Update) events.Type {
	if update.Message == nil {
		return events.Unknown
	}

	return events.Message
}
