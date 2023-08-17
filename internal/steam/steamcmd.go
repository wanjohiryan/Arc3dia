package steamcmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

//TODO: support Steam Guard

func NewSteamCMD() {
	// Get the Steam login credentials from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Steam Username: ")
	username, _ := reader.ReadString('\n')
	fmt.Print("Steam Password: ")
	password, _ := reader.ReadString('\n')

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	// Create a new SteamCMD command
	cmd := exec.Command("steamcmd", "+login", username, password)

	// Create a pipe to capture the command output
	cmdReader, cmdWriter := io.Pipe()
	cmd.Stdout = cmdWriter

	// Create a WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Start the SteamCMD process
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// Start a goroutine to read the command output continuously
	wg.Add(1)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(cmdReader)
		for scanner.Scan() {
			output := scanner.Text()
			fmt.Println(output) // Print the output or handle it as desired
		}
	}()

	// Wait for the SteamCMD prompt
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		// Check if the Steam Guard prompt is displayed
		if strings.Contains(input, "Steam Guard code") {
			// Prompt the user to enter the Steam Guard code
			fmt.Print("Steam Guard Code: ")
			steamGuardCode, _ := reader.ReadString('\n')

			// Write the Steam Guard code to the SteamCMD process
			_, err := fmt.Fprintln(cmdWriter, steamGuardCode)
			if err != nil {
				log.Fatal(err)
			}
			// Check for the SteamCMD prompt to indicate successful login
		} else if strings.Contains(input, "Steam>") {
			// You can now continue with further commands in the SteamCMD session
			log.Println("Succesful login")
			break
			// // Write the Steam login credentials to the SteamCMD process
			// _, err := fmt.Fprintln(cmdWriter, "login", username, password)
			// if err != nil {
			// 	log.Fatal(err)
			// }
		}

		// Check for the SteamCMD prompt to indicate successful login
		// if strings.Contains(input, "Steam>") {
		// 	// You can now continue with further commands in the SteamCMD session
		// 	log.Println("Succesful login")
		// 	break
		// }
	}

	// Close the writer and wait for the SteamCMD process to finish
	cmdWriter.Close()
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	// Wait for the output goroutine to finish
	wg.Wait()
}
