Formatting a string in go. for humans.

- current implementation is naive. PRs are welcome.

Usage
```go
package main

import (
  "fmt"
  "github.com/jossef/format"
)

func main() {
  formattedString := format.String(`hello "{name}". is lizard? {isLizard}`, format.Items{"name": "Mr Dude", "isLizard": false})
  fmt.Println(formattedString)
}
```  

[online demo](https://repl.it/@jossef/format)