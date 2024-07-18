package test

import "fmt"
import "testing"
import "time"
import "strconv"

func TestChannel(t *testing.T) {
    
    channel := make(chan string)
    defer close(channel)
    
    go func() {
        time.Sleep(2 * time.Second)
        channel <- "Ahmad Beni Rusli"
        fmt.Println("Finishing...")
    }()
    
    data := <- channel
    fmt.Println(data)
    
    // channel <- "Beni"
//     data := <- channel
//     fmt.Println(<- channel)
    
    
}

func TestChannel2(t *testing.T) {
    channel := make(chan int)
    
    go func() {
        channel <- 10  //Memasukan data ke dalam channel
    }()
    go func() {
        fmt.Println(<- channel)
    }()
    time.Sleep(3 * time.Second)
}

func pow(params int) chan int {
    channel := make(chan int)
    go func(it int) {
        channel <- it * it
    }(params)
    return channel
}

func TestAsyncFunction(t *testing.T) {
    a := pow(20)
    b := pow(10)
    fmt.Println(<- a, <-b)
}

//=== SEND AND RECEIVE ONLY CHANNEL

//send only
func doSomething(c chan<- string) {
    go func(){
        c <- "Hello World"
    }()
}
//receive only
func printChannel(c <-chan string) {
    fmt.Println(<- c)
}

func TestSendAndReceiveOnlyChannel(t *testing.T) {
    channels := make(chan string)
    doSomething(channels)
    printChannel(channels)
}

//buffered channel

func TestBufferedChannel(t *testing.T) {
    chans := make(chan string, 3)
    chans <- "Ahmad Beni Rusli"
    chans <- "Xaviera Putri"
    chans <- "Via Fitriana"
    //chans <- "Suko"  Deadlock because out of capacity 
    fmt.Println(<- chans)
    fmt.Println(<- chans)
    fmt.Println(<- chans)
}

// Select keyword

func TestSelect(t *testing.T) {
    c1 := make(chan int)
    c2 := make(chan int)
    doneChannel := make(chan int)
    
    go func() {
        time.Sleep(4 * time.Second)
        c1 <- 1
    }()
    
    go func() {
        time.Sleep(3 * time.Second)
        c2 <- 2
    }()
    
    go func() {
        time.Sleep(5 * time.Second)
        doneChannel <- 0
    }()
   
    counter := 0
    for {
        select {
        case data := <- c1:
            fmt.Println(data)
            counter++
        case data := <- c2:
            fmt.Println(data)
            counter++
        }
        if counter == 2 {
			break
		}
    }
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go doSomething(channel1)
	go doSomething(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}
}