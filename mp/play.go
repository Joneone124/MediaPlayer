package mp

import (
	"fmt"
)

type Player interface {
	Play(source string)
}

func Play(source, mtype string) {
	var p Player

	switch mtype {
	case mtype == "MP3":
		p = &MP3Player{}
	case mtype == "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Println("Unsupported music type", mtype)
		return
	}
	p.Play(source)
}
