# ASCII-ART-WEB




## Description

File structure

```
|-banners
|   |-standard.txt
|   |-shadow.txt
|   |_thinkertoy.txt
|-static
|-ascii-art
|   |-formHandler.go
|   |-
|   |_
|-templates
|   |-css.css
|   |_index.html
|-main.go
|-README.md

```

## Implementation details: algorithm
+ Create the html template
    + Creating a form attribute to capture the user's input and choice of the bannerfile.
    + The form also provides an area to display the output(ascii presentation of the input)
+ Handle reponses when user feels the input form
    + Setting up an HTTP handler to process the form submission and parse the form data.
    + Error handling
+ Ascii Art generator
