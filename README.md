# ASCII-ART-WEB

## Description
Ascii-art-web is a web application that provides a graphical user interface for generating ASCII art using different banners. The server allows users to input text, select a banner style, and submit a request to generate ASCII art. The generated ASCII art can be displayed on the same page as per our implementation.



File structure



## Usage
### Prerequisites
    + Go 1.22.2 
### Running the Server

1. Clone the Repository
   ```bash
   git clone https://learn.zone01kisumu.ke/git/wyonyango/ascii-art-web
   cd ascii-art-web
   ```
2. To run the program.
    ```bash
    go run .
    ```
The server will start and listen on port 8081 by default.
You can access it by navigating to `http://localhost:8081` in your web browser.

## Implementation details: algorithm
1. Server Setup:
    - The server is implemented using Go's net/http package.
    - HTML templates are used to render the user interface and display the results.

2. HTTP Endpoints:
    - GET /: Serves the main page of the application.
        - Uses Go templates to render the HTML form.
    - POST /ascii-art: Handles form submissions.
        - Accepts text input and banner selection.
        - Generates ASCII art using the provided text and banner style.
        - Returns the generated ASCII art either on the same page.
3. Error Handling:
    - Returns appropriate HTTP status codes:
        + 200 OK: For successful requests.
        + 404 Not Found: For missing resources (e.g., templates, banners).
        + 400 Bad Request: For invalid form submissions.
        + 500 Internal Server Error: For unexpected errors.

## File System

```
|-banners
|   |-standard.txt
|   |-shadow.txt
|   |_thinkertoy.txt
|-static
|-artgen
|   |-map.go
|   |_printingascii.go
|-ascii-art
|   |_formHandler.go
|-templates
|   |-style.css
|   |_form.html
|-main.go
|-README.md
```


## Authors
+ Wycliffe Onyango
+ Fred Gitonga

## Issues

If you encounter any issues or have suggestions for improvements, please submit an issue or propose changes via the repository.