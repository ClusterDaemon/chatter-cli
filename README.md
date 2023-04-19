# Chatter-CLI

## Introduction

Chatter-CLI is a command-line chat application that leverages OpenAI models to provide users with an interactive and responsive conversational experience. The goal of this project is to create a versatile, user-friendly application that can be used in various contexts, such as a standalone command-line utility or in conjunction with other shell programs.

## Key Features

1. A command-line chat application that is simple and effective to integrate with CLI and pipeline workflows.
2. Stores message history locally for easy access and reuse.
3. Allows multiple chat sessions and starts new chats by default.
4. Supports sending files as messages with custom handling instructions.
5. Compatible with multiple OpenAI models for a tailored chat experience.

## Architecture

Chatter-CLI is written in Golang, and is designed to be used as workflow components. Its simple use of IO streams enables effective usage in CLI workflow environments, as well as contianerized workflow pipelines like Argo Workflows or Kubeflow.

Currently, state is stored as a JSON file. Support for other storage backends will come in the future.

### Data Flow

1. User provides input through the command-line interface, including text messages, files, and options for handling the data.
2. The input is processed and sent as a request to the GPT-4 API.
3. The GPT-4 API processes the request and generates a response.
4. The response is parsed and displayed to the user in the command-line interface.
5. The chat session, including message history, is stored locally for future reference and reuse.

### Error Handling

Chatter-CLI includes error handling for user input, file operations, and API requests. Errors are displayed to the user through the command-line interface, allowing them to take appropriate action to resolve the issue.

## Installation and Setup

### Requirements

- Go 1.16 or higher
- An OpenAI API key for GPT-4 API access

### Building and Running

1. Clone the Chatter-CLI repository.
2. Run `go build -o chatter-cli` to compile the project.
3. Set the `OPENAI_API_KEY` environment variable to your OpenAI API key.
4. Execute `./chatter-cli` to run the application.

## Future Enhancements

1. Improve error handling and user guidance for edge cases and incorrect input.
2. Add support for user-configurable settings, such as chat session storage location and preferred GPT-4 models.
3. Implement additional input options and handling methods for file attachments.
4. Extend the application to support other chat platforms and APIs.
5. Optimize performance and resource usage for large chat sessions and high-volume message processing.
6. Remove stateful components, opting for remote storage without depending on shared filesystems.

## Conclusion

Chatter-CLI is a versatile and user-friendly command-line chat application that leverages the GPT-4 API to provide an interactive conversational experience. The design focuses on ease of use, adaptability, and extensibility, making it a valuable tool for a wide range of use cases and environments. With continued development and user feedback, Chatter-CLI has the potential to become a go-to solution for command-line chat applications and GPT-4 API integration.

