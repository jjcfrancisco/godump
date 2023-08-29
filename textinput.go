package main

import (
	//"fmt"
	//"strings"

	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
)

// TextInput components
const (
    env = iota
    database
    hostname
    port
    username
    password
)

func adbTextInputs() []textinput.Model {

	    var inputs []textinput.Model = make([]textinput.Model, 6)

        inputs[env] = textinput.New()
        inputs[env].Placeholder = "dev"
        inputs[env].Focus()
        inputs[env].CharLimit = 156
        inputs[env].Width = 20
        inputs[env].Prompt = ""

        inputs[database] = textinput.New()
        inputs[database].Placeholder = "postgres"
        inputs[database].CharLimit = 156
        inputs[database].Width = 20
        inputs[database].Prompt = ""
        //inputs[database].Validate = dbValidator
    
        inputs[hostname] = textinput.New()
        inputs[hostname].Placeholder = "localhost"
        inputs[hostname].CharLimit = 156
        inputs[hostname].Width = 20
        inputs[hostname].Prompt = ""
        //inputs[hostname].Validate = hostValidator
    
        inputs[port] = textinput.New()
        inputs[port].Placeholder = "5432"
        inputs[port].CharLimit = 156
        inputs[port].Width = 20
        inputs[port].Prompt = ""
    
        inputs[username] = textinput.New()
        inputs[username].Placeholder = "postgres"
        inputs[username].CharLimit = 156
        inputs[username].Width = 20
        inputs[username].Prompt = ""
    
        inputs[password] = textinput.New()
        inputs[password].Placeholder = "mysecretpassword"
        inputs[password].CharLimit = 156
        inputs[password].Width = 20
        inputs[password].Prompt = ""

        return inputs
}

func edbTextInputs(ic *inputConf) []textinput.Model {

	    var inputs []textinput.Model = make([]textinput.Model, 6)

		inputs[env] = textinput.New()
        inputs[env].Focus()
        inputs[env].CharLimit = 156
        inputs[env].Width = 20
        inputs[env].Prompt = ""
		inputs[env].SetValue(ic.env)

        inputs[database] = textinput.New()
        inputs[database].Placeholder = "postgres"
        inputs[database].CharLimit = 156
        inputs[database].Width = 20
        inputs[database].Prompt = ""
		inputs[database].SetValue(ic.database)
        //inputs[database].Validate = dbValidator
    
        inputs[hostname] = textinput.New()
        inputs[hostname].Placeholder = "localhost"
        inputs[hostname].CharLimit = 156
        inputs[hostname].Width = 20
        inputs[hostname].Prompt = ""
		inputs[hostname].SetValue(ic.hostname)
        //inputs[hostname].Validate = hostValidator
    
        inputs[port] = textinput.New()
        inputs[port].Placeholder = "5432"
        inputs[port].CharLimit = 156
        inputs[port].Width = 20
        inputs[port].Prompt = ""
		inputs[port].SetValue(ic.port)
    
        inputs[username] = textinput.New()
        inputs[username].Placeholder = "postgres"
        inputs[username].CharLimit = 156
        inputs[username].Width = 20
        inputs[username].Prompt = ""
		inputs[username].SetValue(ic.user)
    
        inputs[password] = textinput.New()
        inputs[password].Placeholder = "mysecretpassword"
        inputs[password].CharLimit = 156
        inputs[password].Width = 20
        inputs[password].Prompt = ""
		inputs[password].SetValue(ic.password)

		return inputs

}

func inputSearch() textinput.Model {

	ti := textinput.New()
	ti.Placeholder = "my_file.dump"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
    //ti.Prompt = ""

    return ti

}

func dbValidator(s string) error {
    
    if len(s) == 0 {
            return fmt.Errorf("Must be at least one character")
    }

    return nil
}
func hostValidator(s string) error {
    
    if len(s) == 0 {
        return fmt.Errorf("Must be at least one character")
    }

    return nil
} 
func portValidator(s string) {}
func nameValidator(s string) {}
func passValidator(s string) {}

//func ccnValidator(s string) error {
//	// Credit Card Number should a string less than 20 digits
//	// It should include 16 integers and 3 spaces
//	if len(s) > 16+3 {
//		return fmt.Errorf("CCN is too long")
//	}
//
//	if len(s) == 0 || len(s)%5 != 0 && (s[len(s)-1] < '0' || s[len(s)-1] > '9') {
//		return fmt.Errorf("CCN is invalid")
//	}
//
//	// The last digit should be a number unless it is a multiple of 4 in which
//	// case it should be a space
//	if len(s)%5 == 0 && s[len(s)-1] != ' ' {
//		return fmt.Errorf("CCN must separate groups with spaces")
//	}
//
//	// The remaining digits should be integers
//	c := strings.ReplaceAll(s, " ", "")
//	_, err := strconv.ParseInt(c, 10, 64)
//
//	return err
//}
