package testlib

import (
	"encoding/json"
	"fmt"
	"sync"
)

type metrics struct {
	message int
	event   int
	lock    *sync.Mutex
}

func newMetrics() *metrics {
	return &metrics{
		message: 0,
		event:   0,
		lock:    new(sync.Mutex),
	}
}

func (m *metrics) IncrEvent() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.event += 1
}

func (m *metrics) IncrMessage() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.message += 1
}

func (m *metrics) String() string {
	m.lock.Lock()
	defer m.lock.Unlock()
	return fmt.Sprintf("{message: %d, event: %d}", m.message, m.event)
}

type ReportStore struct {
	metrics map[string]*metrics
	lock    *sync.Mutex
}

func NewReportStore() *ReportStore {
	return &ReportStore{
		metrics: make(map[string]*metrics),
		lock:    new(sync.Mutex),
	}
}

func (r *ReportStore) NewMessage(testcase string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if _, ok := r.metrics[testcase]; !ok {
		r.metrics[testcase] = newMetrics()
	}
	metrics := r.metrics[testcase]
	metrics.IncrMessage()
}

func (r *ReportStore) NewEvent(testcase string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if _, ok := r.metrics[testcase]; !ok {
		r.metrics[testcase] = newMetrics()
	}
	metrics := r.metrics[testcase]
	metrics.IncrEvent()
}

func (r *ReportStore) String() string {
	r.lock.Lock()
	defer r.lock.Unlock()

	kv := make(map[string]string)
	for t, m := range r.metrics {
		kv[t] = m.String()
	}
	res, _ := json.Marshal(kv)
	return string(res)
}
