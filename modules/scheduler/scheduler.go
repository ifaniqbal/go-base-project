package scheduler

type Scheduler struct{}

func (c Scheduler) Run() error {
	return nil
}

func Default() *Scheduler {
	return &Scheduler{}
}
