package mp

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat int
	progress int
}

func (p *MP3Player) Play(source string)  {
	fmt.Println("Playing mp3 music ",source)
	p.progress = 0
	for p.progress < 100{
		//假装正在播放
		time.Sleep(100*time.Millisecond)
		fmt.Print(".");
		p.progress += 10
	}
	fmt.Println("\n Finieshed playing ",source)
}
