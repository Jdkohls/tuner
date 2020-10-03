package keybind

import (
	"os/exec"

	"github.com/Pauloo27/go-mpris"
	"github.com/Pauloo27/tuner/player"
	"github.com/Pauloo27/tuner/utils"
	"github.com/eiannone/keyboard"
)

type Keybind struct {
	Handler              func(cmd *exec.Cmd, mpv *player.MPV)
	KeyName, Description string
}

var (
	ByKey    = map[keyboard.Key]Keybind{}
	ByChar   = map[rune]Keybind{}
	keybinds []Keybind
)

func RegisterDefaultKeybinds() {
	killMpv := Keybind{
		Description: "Stop the player",
		KeyName:     "Esc",
		Handler: func(cmd *exec.Cmd, mpv *player.MPV) {
			_ = cmd.Process.Kill()
		},
	}

	ByKey[keyboard.KeyEsc] = killMpv
	killMpv.KeyName = "CtrlC"
	ByKey[keyboard.KeyCtrlC] = killMpv

	ByKey[keyboard.KeySpace] = Keybind{
		Description: "Play/Pause song",
		KeyName:     "Space",
		Handler: func(cmd *exec.Cmd, mpv *player.MPV) {
			mpv.PlayPause()
		},
	}

	ByChar['9'] = Keybind{
		Description: "Decrease the volume",
		KeyName:     "9",
		Handler: func(cmd *exec.Cmd, mpv *player.MPV) {
			volume, err := mpv.Player.GetVolume()
			utils.HandleError(err, "Cannot get MPV volume")
			err = mpv.Player.SetVolume(volume - 0.05)
			utils.HandleError(err, "Cannot set MPV volume")
		},
	}

	ByChar['0'] = Keybind{
		Description: "Increase the volume",
		KeyName:     "0",
		Handler: func(cmd *exec.Cmd, mpv *player.MPV) {
			volume, err := mpv.Player.GetVolume()
			utils.HandleError(err, "Cannot get MPV volume")
			err = mpv.Player.SetVolume(volume + 0.05)
			utils.HandleError(err, "Cannot set MPV volume")
		},
	}

	ByChar['?'] = Keybind{
		Description: "Toggle keybind list",
		KeyName:     "?",
		Handler: func(cmd *exec.Cmd, mpv *player.MPV) {
			mpv.ShowHelp = !mpv.ShowHelp
			mpv.Update()
		},
	}

	ByChar['l'] = Keybind{
		Description: "Toggle loop",
		KeyName:     "L",
		Handler: func(cmd *exec.Cmd, mpv *player.MPV) {
			loop, err := mpv.Player.GetLoopStatus()
			utils.HandleError(err, "Cannot get mpv loop status")
			newLoopStatus := mpris.LoopNone
			if loop == mpris.LoopNone {
				newLoopStatus = mpris.LoopTrack
			}
			err = mpv.Player.SetLoopStatus(newLoopStatus)
			utils.HandleError(err, "Cannot set loop status")
		},
	}

	ByChar['p'] = Keybind{
		Description: "Toggle lyric",
		KeyName:     "P",
		Handler: func(cmd *exec.Cmd, mpv *player.MPV) {
			if len(mpv.LyricLines) == 0 {
				go mpv.FetchLyric()
			}
			mpv.ShowLyric = !mpv.ShowLyric
			mpv.Update()
		},
	}

	ByChar['w'] = Keybind{
		Description: "Scroll lyric up",
		KeyName:     "W",
		Handler: func(cmd *exec.Cmd, mpv *player.MPV) {
			if mpv.LyricIndex > 0 {
				mpv.LyricIndex = mpv.LyricIndex - 1
				mpv.Update()
			}
		},
	}

	ByChar['s'] = Keybind{
		Description: "Scroll lyric down",
		KeyName:     "S",
		Handler: func(cmd *exec.Cmd, mpv *player.MPV) {
			if mpv.LyricIndex < len(mpv.LyricLines) {
				mpv.LyricIndex = mpv.LyricIndex + 1
				mpv.Update()
			}
		},
	}
}

func HandlePress(c rune, key keyboard.Key, cmd *exec.Cmd, mpv *player.MPV) {
	if bind, ok := ByKey[key]; ok {
		bind.Handler(cmd, mpv)
	} else if bind, ok := ByChar[c]; ok {
		bind.Handler(cmd, mpv)
	}
}

func ListBinds() []Keybind {
	if keybinds != nil {
		return keybinds
	}
	keybinds = []Keybind{}
	for _, bind := range ByKey {
		keybinds = append(keybinds, bind)
	}
	for _, bind := range ByChar {
		keybinds = append(keybinds, bind)
	}
	return keybinds
}
