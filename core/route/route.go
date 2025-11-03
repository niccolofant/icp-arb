package route

import "fmt"

type Route struct {
	id       string
	steps    []*Step
	isActive bool
}

func New(
	steps []*Step,
	isActive bool,
) (*Route, error) {
	if len(steps) < 2 {
		return nil, fmt.Errorf("route must have at least 2 steps")
	}

	id := ""
	for i, step := range steps {
		if i > 0 {
			id = fmt.Sprintf("%s->%s", id, step.ID())
		} else {
			id = step.ID()
		}
	}

	return &Route{
		id:       id,
		steps:    steps,
		isActive: isActive,
	}, nil
}

func (r Route) ID() string {
	return r.id
}

func (r Route) Steps() []*Step {
	return r.steps
}

func (r Route) IsActive() bool {
	return r.isActive
}

func (r Route) String() string {
	str := ""
	for i, step := range r.Steps() {
		if i > 0 {
			str = fmt.Sprintf("%s->%s", str, step)
		} else {
			str = step.String()
		}
	}

	return str
}
