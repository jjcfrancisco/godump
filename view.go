package main

import (
    "fmt"
    "strings"
)

var spacer = " "
var quitMsg = "press q to quit"
var escMsg = "press esc to go back"
var enterMsg = "press enter to continue"

func buildView(m model) string {

	s := strings.Builder{}
    s.WriteString(fmt.Sprintf("\n%s\n\n", m.title))

    switch m.current {
    case "add-db", "edit-single-db":

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
        s.WriteString(fmt.Sprintf("%s%s\n", spacer, enterMsg))
        s.WriteString(fmt.Sprintf("%s%s\n", spacer, escMsg))

    case "type-filename":

        s.WriteString(" Name for the file:")
	    s.WriteString("\n")
        s.WriteString(fmt.Sprintf("\n %s", m.search.View()))
        s.WriteString("\n\n\n")
        s.WriteString(" *Default dumps directory is ~/Dumps ")
        s.WriteString("\n\n")
        s.WriteString(fmt.Sprintf("%s%s\n", spacer, enterMsg))
        s.WriteString(fmt.Sprintf("%s%s\n", spacer, escMsg))

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
		//Allows feedback msg
		if len(m.feedback) > 0 {
			s.WriteString("\n" + "*" + m.feedback + "\n\n")
		} else {
			s.WriteString("\n\n")
		}
        s.WriteString(fmt.Sprintf("%s%s\n", spacer, escMsg))
	    s.WriteString(fmt.Sprintf("%s%s\n", spacer, quitMsg))
    }
    
    return s.String()

}
