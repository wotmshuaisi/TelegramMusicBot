package netease

import (
	"github.com/wotmshuaisi/TelegramMusicBot/src/lib/music"
)

type handler struct {
	EndPoint string
}

func (h *handler) List(text string) (*[]music.Item, error) {
	panic("not implemented")
}

// NewAPI return Netease API handler
func NewAPI(endpoint string) music.API {
	return &handler{
		EndPoint: endpoint,
	}
}
