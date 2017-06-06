# go-clush

go-clush is a simple golang library wrapping the clush (clustershell) command.

## Example

```go
package main

import (
    "fmt"

    "github.com/Devatoria/go-clush"
)

func main() {
    // Will run "clush -g webservers puppet agent -t"
    ret, _ := clush.RunOnGroup("webservers", "puppet agent -t")
    fmt.Println("Stdout: ", ret.Stdout)
    fmt.Println("Stderr: ", ret.Stderr)
 }
```
