package greeting

import "fmt"

const testVersion = 3

func HelloWorld(s string) string {
    if s == "" {
        s = "World"
    }
    return fmt.Sprintf("Hello, %v!", s)
}
