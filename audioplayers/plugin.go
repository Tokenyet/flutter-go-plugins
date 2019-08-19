package audioplayers

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"time"
)

const channelName = "xyz.luan/audioplayers"

type speakerFunc func()

// Not thread safe, for while
var loaded = false

// TODO: mutex me??
func load(sr beep.SampleRate) {
	speaker.Init(sr, sr.N(time.Second/10))
}

func doInSpeaker(f speakerFunc) {
	speaker.Lock()
	defer speaker.Unlock()
	f()
}

func readArg(args interface{}, key string) string {
	return args.(map[interface{}]interface{})[key].(string)
}

type AudioplayersPlugin struct {
	channels map[string]*channel
}

func NewAudioplayersPlugin() *AudioplayersPlugin {
	return &AudioplayersPlugin{map[string]*channel{}}
}

func (p *AudioplayersPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
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

func (p *AudioplayersPlugin) getChannel(args interface{}) (*channel, error) {
	url := readArg(args, "url")
	id := readArg(args, "playerId")

	return p.findOrCreateChannel(id, url)
}

func (p *AudioplayersPlugin) handlePlay(arguments interface{}) (interface{}, error) {
	fmt.Printf("Play\n")
	fmt.Printf("All %v\n", arguments)

	channel, err := p.getChannel(arguments)
	if err != nil {
		return nil, err
	}

	if !loaded {
		load(channel.fmt.SampleRate)
	}

	speaker.Play(beep.Mix(p.getStreams()...))

	return int32(1), nil
}

func (p *AudioplayersPlugin) getStreams() []beep.Streamer {
	ret := make([]beep.Streamer, 0, len(p.channels))
	for _, c := range p.channels {
		ret = append(ret, c.ctrl)
	}
	return ret
}

func (p *AudioplayersPlugin) findOrCreateChannel(id, url string) (*channel, error) {
	channel, ok := p.channels[id]
	if ok {
		return channel, nil
	}

	channel, err := createChannel(id, url)
	if err != nil {
		return nil, err
	}

	p.channels[id] = channel
	return channel, nil
}

func (p *AudioplayersPlugin) handleResume(arguments interface{}) (interface{}, error) {
	fmt.Printf("Resume %v\n", arguments)

	channel, err := p.getChannel(arguments)
	if err != nil {
		return nil, err
	}
	doInSpeaker(func() {
		channel.ctrl.Paused = false
	})
	return int32(1), nil
}

func (p *AudioplayersPlugin) handlePause(arguments interface{}) (interface{}, error) {
	fmt.Printf("Pause %v\n", arguments)
	channel, err := p.getChannel(arguments)
	if err != nil {
		return nil, err
	}
	doInSpeaker(func() {
		channel.ctrl.Paused = true
	})
	return int32(1), nil
}

func (p *AudioplayersPlugin) handleStop(arguments interface{}) (interface{}, error) {
	fmt.Printf("Stop %v\n", arguments)
	channel, err := p.getChannel(arguments)
	if err != nil {
		return nil, err
	}
	doInSpeaker(func() {
		channel.ctrl.Paused = true
		channel.str.Seek(0)
	})
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
