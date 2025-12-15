Ascii-art-web
Description

Ascii-art-web is a Go project that runs an HTTP server and provides a web GUI to generate ASCII art from text.

The user enters text in a form, selects a banner font (standard, shadow, or thinkertoy), and the server returns the ASCII-art rendering of that text.

The rendering logic is based on the ASCII-art banner approach where each printable ASCII character (from space 32 to ~ 126) maps to 8 lines of drawing loaded from a banner file 

main

.

Authors
<Alioune Sall>
<Emilia Chedot>
<Thiago Vargues>

Usage: how to run
Requirements

Go installed

Banner files available in the project root:

standard.txt

shadow.txt

thinkertoy.txt

HTML templates in: ./templates/

Run the server

From the project directory:

go run .


Or build and run:

go build -o ascii-art-web
./ascii-art-web


Then open in your browser:

http://localhost:8080

How to use in the browser

Type text in the input field.

Select a banner.

Click Generate.

Notes about newlines:

The input supports the two-character sequence \n, which is converted to a real newline before rendering (same behavior as the CLI version) 

main

 

README

.

Implementation details: algorithm
HTTP layer (web)

GET /

Returns the main HTML page (template).

POST /ascii-art

Reads form values:

text (user input)

banner (standard/shadow/thinkertoy)

Validates input and banner choice.

Loads the selected banner file.

Renders the result and displays it (either on the same page or on a result page, depending on your implementation).

Banner loading

The banner file is read and split by newline. Each printable ASCII character uses a block of 9 lines in the file:

Line 0: separator (empty)

Lines 1–8: the 8 visual lines of the character

The loader builds a map in the form:

map[rune][]string


Where each rune stores its 8 drawing lines 

main

.

Rendering

Normalize the input:

Replace literal \n with real newlines 

main

.

Split into logical lines by newline 

main

.

For each logical line:

If empty: output a blank line (except at the very end) 

main

.

Otherwise: print 8 rows.

Each row is built by concatenating the corresponding row of each character’s banner representation, using a buffer builder approach 

main

.

Characters outside the supported ASCII range are ignored 

main

.

Tests (optional)

A helper script exists in the CLI project version (audit_tests.sh) to run multiple input cases (including \n) 

audit_tests

. For the web version, you can test endpoints using your browser or tools like curl (if you add those commands yourself).

If you paste your current web main.go (server version), I can tailor the README to your exact port, template names, and whether the POST renders on / or /ascii-art (so it matches your code 1:1).