package main

import (
    "fmt"
    "strings"
)

func buildView(m model) string {

	s := strings.Builder{}
    s.WriteString(fmt.Sprintf("\n%s\n\n", m.title))

    switch m.current {
    case "add-db":

        //if m.textInputs[env].Err != nil {
	    //    s.WriteString(fmt.Sprintf(" Environment %s  -> %s", m.textInputs[env].View(), m.textInputs[env].Err))
        //} else {
	    //    s.WriteString(fmt.Sprintf(" Environment %s", m.textInputs[env].View()))
        //}

	    s.WriteString(fmt.Sprintf(" Environment %s", m.textInputs[env].View()))
	    s.WriteString(fmt.Sprintf("\n Database %s", m.textInputs[database].View()))
	    s.WriteString(fmt.Sprintf("\n Hostname %s", m.textInputs[hostname].View()))
	    s.WriteString(fmt.Sprintf("\n Port %s", m.textInputs[port].View()))
	    s.WriteString(fmt.Sprintf("\n Username %s", m.textInputs[username].View()))
	    s.WriteString(fmt.Sprintf("\n Password %s", m.textInputs[password].View()))
        s.WriteString("\n\n")
        s.WriteString(" press enter to continue\n")
        s.WriteString(" press esc to go back\n")
	    s.WriteString(" press q to quit\n")

    default:

        for i, v := range m.menu {

            if m.cursor == i {
	        	s.WriteString(" - ")
	        } else {
	        	s.WriteString("   ")
	        }
	        s.WriteString(v)
	        s.WriteString("\n")

        }
        s.WriteString("\n\n")
	    s.WriteString(" press esc to go back\n")
	    s.WriteString(" press q to quit\n")
    }
    
    return s.String()

}
