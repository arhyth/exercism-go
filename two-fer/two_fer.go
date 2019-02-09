/*
package twofer is a giver that exposes the ShareWith function
*/
package twofer

import "fmt"

// ShareWith accepts a name string and returns a cliche string
func ShareWith(name string) string {
	var c string
	if name == "" {
		c = "you"
	} else {
		c = name
	}
	return fmt.Sprintf("One for %s, one for me.", c)
}
