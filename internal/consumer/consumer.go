package consumer

type Consumer struct{}

func (c Consumer) Run() error {
	return nil
}

func Default() *Consumer {
	return &Consumer{}
}
