package audioplayers

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"time"
)

const channelName = "xyz.luan/audioplayers"

// Not thread safe, for while
var loaded = false

// TODO: mutex me??
func load(sr beep.SampleRate) {
	speaker.Init(sr, sr.N(time.Second/10))
}

type AudioplayersPlugin struct {
	channels map[string]*channel
}

func NewAudioplayersPlugin() *AudioplayersPlugin {
	return &AudioplayersPlugin{map[string]*channel{}}
}

func (p *AudioplayersPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	fmt.Printf("Loading ...\n\n\n")
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("play", p.handlePlay)
	channel.HandleFunc("resume", p.handleResume)
	channel.HandleFunc("pause", p.handlePause)
	channel.HandleFunc("stop", p.handleStop)
	channel.HandleFunc("release", p.handleRelease)
	channel.HandleFunc("seek", p.handleSeek)
	channel.HandleFunc("setVolume", p.handleSetVolume)
	channel.HandleFunc("setUrl", p.handleSetUrl)
	channel.HandleFunc("setReleaseMode", p.handleSetReleaseMode)
	return nil
}

func (p *AudioplayersPlugin) handlePlay(arguments interface{}) (interface{}, error) {
	fmt.Printf("Play\n")
	fmt.Printf("All %v\n", arguments)

	url := arguments.(map[interface{}]interface{})["url"].(string)
	fmt.Printf("url %s\n", url)

	channel, err := p.getChannel(url)
	if err != nil {
		return nil, err
	}

	if !loaded {
		load(channel.fmt.SampleRate)
	}

	channel.play()
	return int32(1), nil
}

func (p *AudioplayersPlugin) getChannel(url string) (*channel, error) {
	channel, ok := p.channels[url]
	if ok {
		return channel, nil
	}

	channel, err := createChannel(url)
	if err != nil {
		return nil, err
	}

	p.channels[url] = channel
	return channel, nil
}

func (p *AudioplayersPlugin) handleResume(arguments interface{}) (interface{}, error) {
	fmt.Printf("Resume %v\n", arguments)
	return int32(1), nil
}

func (p *AudioplayersPlugin) handlePause(arguments interface{}) (interface{}, error) {
	fmt.Printf("Pause %v\n", arguments)
	return int32(1), nil
}

func (p *AudioplayersPlugin) handleStop(arguments interface{}) (interface{}, error) {
	fmt.Printf("Stop %v\n", arguments)
	return int32(1), nil
}

func (p *AudioplayersPlugin) handleRelease(arguments interface{}) (interface{}, error) {
	fmt.Printf("Release %v\n", arguments)
	return int32(1), nil
}

func (p *AudioplayersPlugin) handleSeek(arguments interface{}) (interface{}, error) {
	fmt.Printf("Seek %v\n", arguments)
	return int32(1), nil
}

func (p *AudioplayersPlugin) handleSetVolume(arguments interface{}) (interface{}, error) {
	fmt.Printf("setVolume %v\n", arguments)
	return int32(1), nil
}

func (p *AudioplayersPlugin) handleSetUrl(arguments interface{}) (interface{}, error) {
	fmt.Printf("setUrl %v\n", arguments)
	return int32(1), nil
}

func (p *AudioplayersPlugin) handleSetReleaseMode(arguments interface{}) (interface{}, error) {
	fmt.Printf("setReleaseMode %v\n", arguments)
	return int32(1), nil
}
