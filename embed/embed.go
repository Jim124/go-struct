package embed

import "github.com/go-struct/saver"

type Embed interface {
	saver.Saver
	Display()
}
