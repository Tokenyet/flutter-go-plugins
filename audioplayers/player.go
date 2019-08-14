package audioplayers

import (
	"bytes"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type channel struct {
	url  string
	str  beep.StreamSeekCloser
	ctrl beep.Ctrl
	fmt  beep.Format
}

func (c *channel) read() (io.ReadCloser, error) {
	//r, err := c.download()
	return os.Open(c.url)
}

func createChannel(url string) (*channel, error) {
	c := channel{url: url}
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
	c.ctrl = beep.Ctrl{Streamer: streamer}
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
	speaker.Play(beep.Seq(c.str))
}

func (c *channel) pause() {
}
