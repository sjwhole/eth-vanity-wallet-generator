package main

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

func generateWallet(startString, endString string, counter *uint64, wg *sync.WaitGroup, found *uint32) {
	defer wg.Done()
	for {
		if atomic.LoadUint32(found) == 1 {
			return
		}

		privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
		if err != nil {
			continue
		}

		publicKey := privateKey.PublicKey
		address := crypto.PubkeyToAddress(publicKey).Hex()

		atomic.AddUint64(counter, 1)

		if strings.HasPrefix(strings.ToLower(address), "0x"+startString) && strings.HasSuffix(strings.ToLower(address), endString) {
			atomic.StoreUint32(found, 1)

			fmt.Println("\nFound matching address!")
			fmt.Printf("Address: %s\n", address)
			fmt.Printf("Private Key: %x\n", crypto.FromECDSA(privateKey))

			return
		}
	}
}

func printStats(counter *uint64, difficulty float64, fiftyPercentProb float64, found *uint32) {
	var prevValue uint64 = 0
	for {
		time.Sleep(1 * time.Second)
		currentValue := atomic.LoadUint64(counter)
		speed := currentValue - prevValue
		prevValue = currentValue

		currentProbability := 1 - math.Pow(1-1/difficulty, float64(currentValue))

		fmt.Printf("\n--- Statistics ---\n")
		fmt.Printf("Difficulty: %.0f\n", difficulty)
		fmt.Printf("Generated: %d addresses\n", currentValue)
		fmt.Printf("50%% Probability: %.0f addresses\n", fiftyPercentProb)
		fmt.Printf("Speed: %d addr/s\n", speed)
		fmt.Printf("Current Probability: %.6f%%\n", currentProbability*100)
		fmt.Printf("------------------\n")

		// Check if a matching address has been found
		if atomic.LoadUint32(found) == 1 {
			return
		}
	}
}

func main() {
	// Input the desired starting and ending strings
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the starting string (without '0x'): ")
	startStringInput, _ := reader.ReadString('\n')
	startString := strings.ToLower(strings.TrimSpace(startStringInput))

	fmt.Print("Enter the ending string: ")
	endStringInput, _ := reader.ReadString('\n')
	endString := strings.ToLower(strings.TrimSpace(endStringInput))

	fmt.Printf("Enter the number of goroutines to use (Available CPUs: %d): ", runtime.NumCPU())
	var numGoroutines int
	fmt.Scanf("%d", &numGoroutines)

	lengthOfMatch := len(startString) + len(endString)
	difficulty := math.Pow(16, float64(lengthOfMatch))

	fiftyPercentProb := math.Log(1-0.5) / math.Log(1-1/difficulty)

	var counter uint64 = 0
	var found uint32 = 0

	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go generateWallet(startString, endString, &counter, &wg, &found)
	}

	go printStats(&counter, difficulty, fiftyPercentProb, &found)

	wg.Wait()
}
