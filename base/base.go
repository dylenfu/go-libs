package base

type Tee struct {
	Num  uint
	Name  string
}

func (tee *Tee)  GetTeeNum() uint {
	return tee.Num
}

func (tee Tee) GetTeeName() string {
	return tee.Name
}

func ChannelDemo() {
	messages := make(chan string)

	msg := "hi"
	go func() {
		messages <- msg
	}()

	select {
	case msg := <-messages:
		println("received message", msg)
	}

}
