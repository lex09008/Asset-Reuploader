package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/lex09008/Asset-Reuploader/internal/app/config"
	"github.com/lex09008/Asset-Reuploader/internal/color"
	"github.com/lex09008/Asset-Reuploader/internal/console"
	"github.com/lex09008/Asset-Reuploader/internal/files"
	"github.com/lex09008/Asset-Reuploader/internal/roblox"
)

var cookieFile = config.Get("cookie_file")

func main() {
	console.ClearScreen()

	fmt.Println("Authenticating cookie...")

	cookie, readErr := files.Read(cookieFile)
	cookie = strings.TrimSpace(cookie)

	c, clientErr := roblox.NewClient(cookie)
	console.ClearScreen()

	if readErr != nil || clientErr != nil {
		if readErr != nil && !os.IsNotExist(readErr) {
			color.Error.Println(readErr)
		}

		if clientErr != nil && cookie != "" {
			color.Error.Println(clientErr)
		}

		getCookie(c)
	}

	if err := files.Write(cookieFile, c.Cookie); err != nil {
		color.Error.Println("Failed to save cookie: ", err)
	}

	fmt.Println("localhost started. Waiting to start reuploading.")
	if err := serve(c); err != nil {
		log.Fatal(err)
	}
}

func getCookie(c *roblox.Client) {
	for {
		i, err := console.LongInput("ROBLOSECURITY: ")
		console.ClearScreen()
		if err != nil {
			color.Error.Println(err)
			continue
		}

		fmt.Println("Authenticating cookie...")
		err = c.SetCookie(i)
		console.ClearScreen()
		if err != nil {
			color.Error.Println(err)
			continue
		}

		files.Write(cookieFile, i)
		break
	}
}
