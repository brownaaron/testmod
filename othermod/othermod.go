package othermod

import "fmt" 

// Hi returns a friendly greeting
func Hi(name string) string {
   return fmt.Sprintf("Hi other, %s", name)
}
