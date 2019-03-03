package testmod

import "fmt" 

// Hi returns a friendly greeting
func Hi(name string) string {
   return fmt.Sprintf("v1.1.0 Hi, %s", name)
}
