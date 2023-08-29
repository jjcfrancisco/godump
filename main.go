package main

import (
	"fmt"
	"os"
    
    //"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)



func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
    var cmd tea.Cmd
    var cmds []tea.Cmd = make([]tea.Cmd, len(m.textInputs))

    switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
        case "esc":
            return goBack(m), nil

		case "ctrl+c":
			return m, tea.Quit

		case "q":

			excep := exceptions(m)

			if excep {
				return m, nil
			} else {
				return m, tea.Quit
			}


		case "enter":
            m = newModel(m)
			return m, nil

		case "down":
			m.cursor++
            switch m.current {
                case "add-db", "edit-single-db":
                    if m.cursor >= len(m.textInputs) {
                        m.cursor = 0
                    }
                    m = nextTextarea(m, msg)
                default:
			        if m.cursor >= len(m.menu) {
			        	m.cursor = 0
			        }
            }

		case "up":
			m.cursor--
            switch m.current {
            case "add-db", "edit-single-db":
                if m.cursor < 0 {
                    m.cursor = len(m.textInputs) - 1
                }
                m = nextTextarea(m, msg)
            default:
			    if m.cursor < 0 {
			    	m.cursor = len(m.menu) - 1
			    }
            }

        case "ctrl+t":
			fmt.Println(m.current, m.previous)
	    }

	}
    
	for i := range m.textInputs {
		m.textInputs[i], cmds[i] = m.textInputs[i].Update(msg)
	}

    m.search, cmd = m.search.Update(msg)
    cmds = append(cmds, cmd)

    return m, tea.Batch(cmds...)
}

func (m model) View() string {

    s := buildView(m)

	return s
}

func main() {
    p := tea.NewProgram(model{menu: mmItems, title: mmTitle})

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.choice != "" {
        fmt.Printf("\n---\nYou chose %s!\n", m.choice)
	}
}
