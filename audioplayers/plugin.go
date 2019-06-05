package audioplayers
// https://github.com/faiface/beep/wiki/Hello,-Beep!

import (
        "fmt"
        //"time"
        "os"
        "log"
	      flutter "github.com/go-flutter-desktop/go-flutter"
        "github.com/go-flutter-desktop/go-flutter/plugin"

        "github.com/faiface/beep"
        "github.com/faiface/beep/mp3"
        //"github.com/faiface/beep/speaker"
)

const channelName = "xyz.luan/audioplayers"

type AudioplayersPlugin struct{}

var _ flutter.Plugin = &AudioplayersPlugin{}

var url string
var loadedStreamer beep.Streamer
var loadedFormat beep.Format

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
        fmt.Printf("Url %s", url)

        //speaker.Init(loadedFormat.SampleRate, loadedFormat.SampleRate.N(time.Second/10))

        //done := make(chan bool)
        //speaker.Play(beep.Seq(loadedStreamer, beep.Callback(func() {
        //        done <- true
        //})))

        //<-done

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
        url = string(arguments.(map[interface{}]interface{})["url"].(string))
        fmt.Printf("setUrl %s", url)
        //isLocal := bool(arguments.(map[interface{}]interface{})["isLocal"].(bool))

        file, err := os.Open(url)
        if err != nil {
                log.Fatal(err)
        }

        streamer, format, err := mp3.Decode(file)

        loadedStreamer = streamer
        loadedFormat = format

        if err != nil {
                log.Fatal(err)
        }

        //defer streamer.Close()
        return int32(1), err
}

func handleSetReleaseMode(arguments interface{}) (reply interface{}, err error) {
        fmt.Printf("setReleaseMode")
        return int32(1), nil
}
