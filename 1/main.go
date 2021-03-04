package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	instances := []*NATInstance{
		&NATInstance{
			Id:   "1",
			Zone: "us-west1-a",
		},
		&NATInstance{
			Id:   "2",
			Zone: "us-west1-b",
		},
		&NATInstance{
			Id:   "3",
			Zone: "us-west1-c",
		},
	}

	subnets := []*Subnet{
		&Subnet{
			Id:   "1",
			Zone: "us-west1-a",
		},
		&Subnet{
			Id:   "2",
			Zone: "us-west1-b",
		},
		&Subnet{
			Id:   "3",
			Zone: "us-west1-b",
		},
		&Subnet{
			Id:   "4",
			Zone: "us-west1-c",
		},
	}
	// default allocate
	res := allocate(instances, subnets)

	// re-check
	for _, i := range res {
		// If instance has no subnet, RandomAllocate
		if len(i.Subnets) == 0 {
			i.randomAllocate(subnets)
		}
	}
	printInstances(res)
}

type Subnet struct {
	Id   string
	Zone string
}

type NATInstance struct {
	Id      string
	Zone    string
	Subnets []*Subnet
}

func printInstances(instances []*NATInstance) {
	for _, i := range instances {
		fmt.Printf("Instance (%v - %v):\n", i.Id, i.Zone)
		for _, s := range i.Subnets {
			fmt.Printf("\tsubnet (%v - %v)\n", s.Id, s.Zone)
		}
	}
}

// Allocator allocate subnet
type Allocator interface {
	allocate()
	randomAllocate()
}

// Allocate subet for instance
func (n *NATInstance) allocate(s []*Subnet) {
	for _, i := range s {
		if i.Zone == n.Zone {
			// add subnet to instance Subnet
			n.Subnets = append(n.Subnets, i)
			continue
		}
	}
}

// RandomAllocate subet for instance when instance has no subnet
func (n *NATInstance) randomAllocate(s []*Subnet) {

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(s) + 1)

	// add subnet to instance Subnet
	n.Subnets = append(n.Subnets, s[num])
}

// allocate Subnets to Instances
func allocate(instances []*NATInstance, subnets []*Subnet) []*NATInstance {

	// implement this function
	for _, i := range instances {
		i.allocate(subnets)
	}
	return instances
}
