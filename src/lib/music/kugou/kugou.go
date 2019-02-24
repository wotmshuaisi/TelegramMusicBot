package kugou

import (
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/buger/jsonparser"
	"github.com/google/uuid"
	"github.com/wotmshuaisi/TelegramMusicBot/src/lib/music"
	"github.com/wotmshuaisi/TelegramMusicBot/src/lib/utils"
)

var geturlWg sync.WaitGroup

type handler struct {
	EndPoint  string
	Endpoint1 string
}

func (h *handler) getURL(id string, index int, l *[]*music.Item) {
	// wait group
	geturlWg.Add(1)
	defer geturlWg.Done()
	// get song detail
	b, err := utils.HTTPGetJSON(h.Endpoint1 + "app/i/getSongInfo.php?cmd=playInfo&hash=" + id)
	if err != nil {
		(*l)[index].URL = "ERROR"
		logrus.WithError(err).Errorf("bytes: %s", b)
		return
	}
	// set url
	(*l)[index].URL, err = jsonparser.GetString(b, "url")
	if err != nil {
		(*l)[index].URL = "ERROR"
		logrus.WithError(err).Errorf("id: %s, bytes: %s", id, b)
		return
	}
}

func (h *handler) fetchSongs(b []byte) []*music.Item {
	var l []*music.Item
	// var wg sync.WaitGroup
	// fetch list songs
	jsonparser.ArrayEach(b, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		// get id, title, performer
		id, err := jsonparser.GetString(value, "hash")
		title, err := jsonparser.GetString(value, "songname")
		duration, _ := jsonparser.GetInt(value, "duration")
		performer, _ := jsonparser.GetString(value, "singername")
		if err != nil {
			return
		}
		uid, _ := uuid.NewUUID()
		// append audio item
		l = append(l, &music.Item{
			ID:        uid.String(),
			Title:     title,
			Performer: performer,
			Duration:  int(duration),
		})
		// set url
		go h.getURL(id, len(l)-1, &l)
	}, "info")
	// waiting for all taks done
	geturlWg.Wait()
	return l
}

func (h *handler) List(text string) (*[]*music.Item, error) {
	b, err := utils.HTTPGetJSON(h.EndPoint + "api/v3/search/song?format=json&keyword=" + text + "&page=0&pagesize=15&showtype=1")
	if err != nil {
		return nil, err
	}
	b, _, _, err = jsonparser.Get(b, "data")
	l := h.fetchSongs(b)
	return &l, nil
}

func (h *handler) Get(id string) (*music.Item, error) {
	panic("not implemented")
}

// NewAPI return kuwo API handler
func NewAPI() music.API {
	return &handler{
		EndPoint:  "http://mobilecdn.kugou.com/",
		Endpoint1: "http://m.kugou.com/",
	}
}
