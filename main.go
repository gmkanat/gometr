package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Measurable
	Ping() error
	GetID() string
	Health() bool
}

type Checker struct {
	items []Checkable
}

func (c *Checker) Add(items ...Checkable) {
	c.items = append(c.items, items...)
}

func (c *Checker) String() string {
	var ids []string
	for _, item := range c.items {
		ids = append(ids, item.GetID())
	}
	return strings.Join(ids, " ")
}

func (c *Checker) Check() {
	for _, item := range c.items {
		if !item.Health() {
			fmt.Println(item.GetID(), "does not work")
		}
	}
}

type GoMetrClient struct {
	URL        string
	TimeoutSec int
}

func (g *GoMetrClient) Ping() error {
	// Add code to ping the server and return error if necessary
	return nil
}

func (g *GoMetrClient) GetID() string {
	return g.URL
}

func (g *GoMetrClient) Health() bool {
	hc := g.getHealth()
	if hc.Status {
		return true
	}
	return false
}

func (g *GoMetrClient) GetMetrics() string {
	// Add code to get metrics from the server and return as a string
	return ""
}

type HealthCheck struct {
	ServiceID string
	Status    bool
	Error     error
}

func (g *GoMetrClient) getHealth() HealthCheck {
	var hc HealthCheck
	hc.ServiceID = g.URL
	hc.Status, hc.Error = g.checkHealth()
	return hc
}

func (g *GoMetrClient) checkHealth() (bool, error) {
	if strings.HasSuffix(strconv.Itoa(g.TimeoutSec), "0") {
		return true, nil
	}
	return false, nil
}

func main() {
	var checker Checker
	checker.Add(&GoMetrClient{URL: "https://example1.com", TimeoutSec: 10})
	checker.Add(&GoMetrClient{URL: "https://example2.com", TimeoutSec: 15})
	checker.Add(&GoMetrClient{URL: "https://example3.com", TimeoutSec: 20})
	checker.Check()
}
