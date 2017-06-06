# go-clush

go-clush is a simple golang library wrapping the clush (clustershell) command.

## Example

```
package main

import (
    "fmt"

    "github.com/Devatoria/go-clush"
)

func main() {
    ret, _ := clush.RunOnGroup("webservers", "puppet agent -t")
    fmt.Println("Stdout: ", ret.Stdout)
    fmt.Println("Stderr: ", ret.Stderr)
 }
```
