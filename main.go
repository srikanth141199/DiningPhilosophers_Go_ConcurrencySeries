package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name                string
	rightFork, leftFork int
}

// list of all philosophers
var philosophers = []Philosopher{
	{name: "Socrates", rightFork: 0, leftFork: 4},
	{name: "Plato", rightFork: 1, leftFork: 0},
	{name: "Aristotle", rightFork: 2, leftFork: 1},
	{name: "Confucius", rightFork: 3, leftFork: 2},
	{name: "Nietzsche", rightFork: 4, leftFork: 3},
}

var hunger = 3 // how many times a philosopher will eat
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

var orderMutex sync.Mutex
var orderFinished []string

func main() {
	fmt.Println("Hello, World!")

	// print a welcome message
	fmt.Println("Welcome to the Dining Philosophers problem!")
	fmt.Println("---------------------------------------------")
	fmt.Println("The table is empty")

	time.Sleep(sleepTime)

	// start a meal
	dine()

	// print out finished message
	fmt.Println("All philosophers have finished eating!")
	fmt.Println("Order of philosophers finishing eating:")
	for i, name := range orderFinished {
		fmt.Printf("%d. %s\n", i+1, name)
	}
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks
	forks := make(map[int]*sync.Mutex)

	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(p Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()
	//seat the philosopher at the table
	fmt.Printf("%s is seated at the table\n", p.name)
	seated.Done()
	seated.Wait()

	//eat 3 times

	for i := hunger; i > 0; i-- {
		//get the lock on both forks

		if p.leftFork > p.rightFork {
			forks[p.leftFork].Lock()
			fmt.Printf("%s picked up left fork %d\n", p.name, p.leftFork)
			forks[p.rightFork].Lock()
			fmt.Printf("%s picked up right fork %d\n", p.name, p.rightFork)
		} else {
			forks[p.rightFork].Lock()
			fmt.Printf("%s picked up right fork %d\n", p.name, p.rightFork)
			forks[p.leftFork].Lock()
			fmt.Printf("%s picked up left fork %d\n", p.name, p.leftFork)
		}

		fmt.Printf("%s is eating\n", p.name)
		time.Sleep(eatTime)

		fmt.Printf("%s is thinking\n", p.name)
		time.Sleep(thinkTime)

		forks[p.rightFork].Unlock()
		fmt.Printf("%s put down right fork %d\n", p.name, p.rightFork)

		forks[p.leftFork].Unlock()
		fmt.Printf("%s put down left fork %d\n", p.name, p.leftFork)

	}

	fmt.Printf("%s is done eating and leaves the table\n", p.name)

	orderMutex.Lock()
	orderFinished = append(orderFinished, p.name)
	orderMutex.Unlock()
}
