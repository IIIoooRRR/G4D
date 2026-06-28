package g4d

func (b *Bot) InitProcessors(p Definition, quantity uint, limitSize uint) {
	for i := 0; i < int(quantity); i++ {
		go p(b, limitSize)
	}
}
