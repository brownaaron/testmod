package testmod

import "fmt" 

// Hi returns a friendly greeting
func Hi(name string) string {
   return fmt.Sprintf("v1.0.2 Hi, %s", name)
}
