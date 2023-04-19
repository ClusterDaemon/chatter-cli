package main

import (
	"bytes"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
    "flag"
	"github.com/google/uuid"
)

type Message struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type ChatSession struct {
	ID       string    `json:"id"`
	Messages []Message `json:"messages"`
}

type GPT4Request struct {
	Session  ChatSession `json:"session"`
	ChatText string      `json:"chat_text"`
}

type GPT4Response struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func generateSessionID() string {
	return uuid.New().String()
}

func sendMessage(session ChatSession, message, model string) (GPT4Response, error) {
	// Prepare the GPT-4 API request
	gpt4Request := GPT4Request{
		Session:  session,
		ChatText: message,
	}
	requestJSON, err := json.Marshal(gpt4Request)
	if err != nil {
		return GPT4Response{}, err
	}

	// Send the API request
	apiURL := fmt.Sprintf("https://api.openai.com/v1/engines/%s/chat", model)
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		return GPT4Response{}, err
	}
	defer resp.Body.Close()

	// Parse the API response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GPT4Response{}, err
	}

	var gpt4Response GPT4Response
	err = json.Unmarshal(body, &gpt4Response)
	if err != nil {
		return GPT4Response{}, err
	}

	return gpt4Response, nil
}


func loadChatSessions() (map[string]ChatSession, error) {
	activeSessions := make(map[string]ChatSession)

	data, err := ioutil.ReadFile(chatSessionsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return activeSessions, nil
		}
		return nil, err
	}

	err = json.Unmarshal(data, &activeSessions)
	if err != nil {
		return nil, err
	}

	return activeSessions, nil
}

func saveChatSessions(sessions map[string]ChatSession) error {
	data, err := json.MarshalIndent(sessions, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(chatSessionsFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

const chatSessionsFile = "chat_sessions.json"

var modelFlag = flag.String("model", "davinci-codex", "The AI model to use (e.g., 'davinci-codex', 'curie-codex')")
var modelShortFlag = flag.String("m", "", "The AI model to use (short flag, same as --model)")
var sessionFlag = flag.String("session", "", "The chat session ID to use or resume")
var sessionShortFlag = flag.String("s", "", "The chat session ID to use or resume (short flag, same as --session)")

func main() {
	flag.Parse()

	model := *modelFlag
	if *modelShortFlag != "" {
		model = *modelShortFlag
	}
	sessionID := *sessionFlag
	if *sessionShortFlag != "" {
	    sessionID = *sessionShortFlag
	}

	var currentSession ChatSession
	var activeSessions map[string]ChatSession

	activeSessions, err := loadChatSessions()
	if err != nil {
		fmt.Println("Error loading chat sessions:", err)
		return
	}

	if sessionID != "" {
	    if session, ok := activeSessions[sessionID]; ok {
	        currentSession = session
	    } else {
	        currentSession = ChatSession{
	            ID:       sessionID,
	            Messages: []Message{},
	        }
	        activeSessions[sessionID] = currentSession
	    }
	} else {
	    // Generate a new session ID if none is provided
	    newSessionID := generateSessionID()
	    currentSession = ChatSession{
	        ID:       newSessionID,
	        Messages: []Message{},
	    }
	    activeSessions[newSessionID] = currentSession
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		// Read user input
		fmt.Print(">")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		response, err := sendMessage(currentSession, input, model)

		// Update the current session with the new message
		newMessage := Message{
			ID:   response.ID,
			Text: response.Text,
		}

		currentSession.Messages = append(currentSession.Messages, newMessage)
		activeSessions[currentSession.ID] = currentSession

		// Save the updated chat sessions to the file
		err = saveChatSessions(activeSessions)
		if err != nil {
			fmt.Println("Error saving chat sessions:", err)
			return
		}
	}
}
