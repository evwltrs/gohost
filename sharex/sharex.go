package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	f, err := os.Create("config.sxcu")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString("{\n  \"Version\": \"13.6.1\",\n  \"Name\": \"gohost\",\n  \"DestinationType\": \"FileUploader\",\n  \"RequestMethod\": \"POST\",\n  \"RequestURL\": \"" + os.Getenv("HOSTNAME") + "/api/upload\",\n  \"Headers\": {\n    \"api_key\": \"" + os.Getenv("API_KEY") + "\"\n  },\n  \"Body\": \"MultipartFormData\",\n  \"FileFormName\": \"file\",\n  \"URL\": \"$json:url$\"\n}")
	if err != nil {
		fmt.Println(err)
		err := f.Close()
		if err != nil {
			return
		}
		return
	}
	fmt.Println("config.sxcu created")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
