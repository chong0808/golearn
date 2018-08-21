package merkel

type Engine struct {
	stop  bool
	start bool
}

func (en *Engine) startAction() {
	en.start = true
	en.stop = false
}
func (en *Engine) stopAction() {
	en.start = false
	en.stop = true
}

type Car struct {
	wheelCount int
	Engine
}

func (c Car) numberOfWheel() int {
	return c.wheelCount
}

type Mercedes struct {
	Car
}

func (m *Mercedes) sayHiToMerkel() string {
	return "Merkel"
}
