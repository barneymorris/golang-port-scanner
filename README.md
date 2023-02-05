## What it is?
It's simple Golang program which show you all opened and listen ports and which allow you to kill some process on some specific port.

## How to use that?
1. Install Golang. I think v1.18 is enough.
2. Go to folder where you download that folder of code
3. Compile that: go build -o port-scanner ./main.go
4. Run it: ./port-scanner --help

| Command | Description |
|--|--|
| show | displays all opened and listen port. Usage: ./port-scanner show |
| kill | kill some process on specific port. --port flag is required

Feel free to use in any purposes. It's written in self-learning purposes.
