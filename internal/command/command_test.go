package command

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/wayming/dpat/internal/random"
)

func testToUpper(number int, strLen int, channel chan command, work *sync.WaitGroup) {
	defer work.Done()

	fmt.Println("testToUpper Start")
	var upperReceiver upperStringReceiver
	for i := 0; i < number; i++ {
		// testUpperCmd := &ToUpperCmd{random.RandString(strLen), upperReceiver}
		// testUpperCmd.Execute()
		// fmt.Println(testUpperCmd.otext)
		channel <- &ToUpperCmd{random.RandString(strLen), upperReceiver}
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("testToUpper End")
}

func testToLower(number int, strLen int, channel chan command, work *sync.WaitGroup) {
	defer work.Done()

	fmt.Println("testToLower Start")
	var lowerReceiver lowerStringReceiver
	for i := 0; i < number; i++ {
		// testLowerCmd := &ToLowerCmd{random.RandString(strLen), lowerReceiver}
		// testLowerCmd.Execute()
		// fmt.Println(testLowerCmd.otext)
		channel <- &ToUpperCmd{random.RandString(strLen), lowerReceiver}
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("testToLower End")
}

func cmdProcess(channel chan command) {
	fmt.Println("Begin processing")
	for cmd := range channel {
		cmd.Execute()
	}
	fmt.Println("Finish processing")
}

// TestPattern test pattern
func TestPattern(t *testing.T) {
	fmt.Println("command package")

	var wg sync.WaitGroup
	ch := make(chan command)
	for i := 0; i < 100; i++ {
		go testToUpper(1000, 10, ch, &wg)
		go testToLower(1000, 15, ch, &wg)
		wg.Add(2)
	}

	go cmdProcess(ch)

	wg.Wait()
	fmt.Println("WaitGroup Completed")

	close(ch)
	fmt.Println("Channel Closed")

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}
