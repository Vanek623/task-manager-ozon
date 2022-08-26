package counters

import (
	"expvar"
	"fmt"
)

type CounterName int

const (
	Success CounterName = iota
	Fail
	Incoming
	Outbound
)

type Counters struct {
	success  *expvar.Int
	fail     *expvar.Int
	incoming *expvar.Int
	outbound *expvar.Int
}

func New(name string) *Counters {
	return &Counters{
		success:  expvar.NewInt(makeName(name, successName)),
		fail:     expvar.NewInt(makeName(name, failName)),
		incoming: expvar.NewInt(makeName(name, incomingName)),
		outbound: expvar.NewInt(makeName(name, outboundName)),
	}
}

func (c *Counters) Inc(counter CounterName) {
	switch counter {
	case Success:
		c.success.Add(1)
	case Fail:
		c.fail.Add(1)
	case Incoming:
		c.incoming.Add(1)
	case Outbound:
		c.outbound.Add(1)
	}
}

func (c *Counters) Value(counter CounterName) int64 {
	switch counter {
	case Success:
		return c.success.Value()
	case Fail:
		return c.fail.Value()
	case Incoming:
		return c.incoming.Value()
	case Outbound:
		return c.outbound.Value()
	}

	return 0
}

const (
	successName  = "success"
	failName     = "fail"
	incomingName = "incoming"
	outboundName = "outbound"
)

func makeName(group, name string) string {
	return fmt.Sprintf("%s_%s", group, name)
}
