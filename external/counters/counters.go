package counters

import (
	"expvar"
	"fmt"
)

// CounterName название счетчика
type CounterName int

const (
	// Success счетчик успешно обработанных входящих запросов
	Success CounterName = iota
	// Fail счетчик неудачно обработанных входящих запросов
	Fail
	// Incoming счетчик входящих запросов
	Incoming
	// Outbound счетчик исходящих запросов
	Outbound
)

const (
	successName  = "success"
	failName     = "fail"
	incomingName = "incoming"
	outboundName = "outbound"
)

var initInfo map[CounterName]string

func init() {
	initInfo = make(map[CounterName]string)

	initInfo[Success] = successName
	initInfo[Fail] = failName
	initInfo[Incoming] = incomingName
	initInfo[Outbound] = outboundName
}

// Counters группа счетчиков
type Counters struct {
	cs []*expvar.Int
}

// New создание группы счетчиков
func New(groupName string) *Counters {
	cs := make([]*expvar.Int, 0, len(initInfo))

	for _, v := range initInfo {
		cs = append(cs, expvar.NewInt(makeName(groupName, v)))
	}

	return &Counters{
		cs: cs,
	}
}

// Inc увеличить выбранный счетчик
func (c *Counters) Inc(counter CounterName) {
	c.cs[counter].Add(1)
}

// Value прочитать значение счетчика
func (c *Counters) Value(counter CounterName) int64 {
	c.cs[counter].Value()

	return 0
}

func makeName(group, name string) string {
	return fmt.Sprintf("%s:%s", group, name)
}
