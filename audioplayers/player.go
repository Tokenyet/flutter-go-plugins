package audioplayers

import (
	"bytes"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type channel struct {
	id   string
	url  string
	str  beep.StreamSeekCloser
	ctrl *beep.Ctrl
	fmt  beep.Format
}

func isUrl(test string) bool {
	u, err := url.ParseRequestURI(test)
	return err == nil && u.Scheme != ""
}

func (c *channel) read() (io.ReadCloser, error) {
	if isUrl(c.url) {
		fmt.Printf("Downloading %s\n", c.url)
		return c.download()
	}
	return os.Open(c.url)
}

func createChannel(id, url string) (*channel, error) {
	c := channel{id: id, url: url}
	r, err := c.read()

	if err != nil {
		return nil, err
	}

	err = c.decode(r)
	return &c, err
}

func (c *channel) decode(r io.ReadCloser) error {
	streamer, format, err := mp3.Decode(r)

	if err != nil {
		return err
	}

	c.str = streamer
	c.ctrl = &beep.Ctrl{Streamer: streamer}
	c.fmt = format

	return nil
}

func (c *channel) download() (io.ReadCloser, error) {
	resp, err := http.Get(c.url)
	if err != nil {
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	return ioutil.NopCloser(r), nil
}

func (c *channel) play() {
	c.ctrl.Paused = false
}

func (c *channel) pause() {
	c.ctrl.Paused = true
}
