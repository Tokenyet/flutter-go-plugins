package audioplayers

import (
        "fmt"
	      flutter "github.com/go-flutter-desktop/go-flutter"
        "github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "xyz.luan/audioplayers"

type AudioplayersPlugin struct{}

var _ flutter.Plugin = &AudioplayersPlugin{}

func (p *AudioplayersPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
        channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
        channel.HandleFunc("play", handlePlay)
        channel.HandleFunc("resume", handleResume)
        channel.HandleFunc("pause", handlePause)
        channel.HandleFunc("stop", handleStop)
        channel.HandleFunc("release", handleRelease)
        channel.HandleFunc("seek", handleSeek)
        channel.HandleFunc("setVolume", handleSetVolume)
        channel.HandleFunc("setUrl", handleSetUrl)
        channel.HandleFunc("setReleaseMode", handleSetReleaseMode)
        return nil
}

func handlePlay(arguments interface{}) (reply interface{}, err error) {
        fmt.Printf("Play")
        return int32(1), nil
}

func handleResume(arguments interface{}) (reply interface{}, err error) {
        fmt.Printf("Resume")
        return int32(1), nil
}

func handlePause(arguments interface{}) (reply interface{}, err error) {
        fmt.Printf("Pause")
        return int32(1), nil
}

func handleStop(arguments interface{}) (reply interface{}, err error) {
        fmt.Printf("Stop")
        return int32(1), nil
}

func handleRelease(arguments interface{}) (reply interface{}, err error) {
        fmt.Printf("Release")
        return int32(1), nil
}

func handleSeek(arguments interface{}) (reply interface{}, err error) {
        fmt.Printf("Seek")
        return int32(1), nil
}

func handleSetVolume(arguments interface{}) (reply interface{}, err error) {
        fmt.Printf("setVolume")
        return int32(1), nil
}

func handleSetUrl(arguments interface{}) (reply interface{}, err error) {
        url := []byte(arguments.(map[interface{}]interface{})["url"].(string))
        isLocal := bool(arguments.(map[interface{}]interface{})["isLocal"].(bool))

        fmt.Printf("setUrl")
        fmt.Printf("%s", url)
        fmt.Printf("%b", isLocal)

        return int32(1), nil
}

func handleSetReleaseMode(arguments interface{}) (reply interface{}, err error) {
        fmt.Printf("setReleaseMode")
        return int32(1), nil
}
