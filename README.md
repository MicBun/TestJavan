# TestJavan
TestJavan

This app use swagger for api documentation

---
### How to use with Go Build
1. Clone this Repository using `git clone`
2. Check config/db.go for the commented out code and adjust accordingly
3. use CMD Command on this project root
    1. `go build -tags netgo -ldflags '-s -w' -o app`
    2. `./app`
4. access http://localhost:8080/swagger/index.html#/
5. The app is ready to use.
---
### How to use with Go Run
1. Clone this Repository using `git clone`
2. Check config/db.go for the commented out code and adjust accordingly
3. use CMD Command on this project root
    1. `go run main.go`
4. access http://localhost:8080/swagger/index.html#/
---
For more information, please contact me LinkedIn: https://www.linkedin.com/in/MicBun