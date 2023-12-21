package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var userInput string

func main() {
	
	//Load .env
	err := godotenv.Load()
	if err != nil{
		log.Fatal(err)
	}

	//initializing the gemini model 
	ctx := context.Background()
	
	fmt.Printf("This is Gemini AI ! ! \n")

	//NOTE: API_KEY is important
	client, err := genai.NewClient(ctx,option.WithAPIKey(os.Getenv("API_KEY")) )
	if err !=nil{
		log.Fatal("Error making new client", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")

	//Taking input from the user

	// _ ,err = fmt.Scan(&userInput)
	// if err!= nil{
	// 	log.Fatal("Error taking the input")
	// }
	
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print(">> ")
		userInput, _ := reader.ReadString('\n')
		
		if userInput == "exit"{
			break
		}
		resp, err := model.GenerateContent(ctx, genai.Text(userInput))
		if err!= nil{
			fmt.Print(err)
		}
		printResponse(resp)
	}
	
}

func printResponse(resp *genai.GenerateContentResponse){
	for _, cand := range resp.Candidates{
		if resp.Candidates != nil{
			for _,part := range cand.Content.Parts{
				fmt.Println(">>> ",part)
			}
		}
	}
}