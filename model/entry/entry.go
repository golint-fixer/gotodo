package entry

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/marthjod/gotodo/model/context"
	"github.com/marthjod/gotodo/model/priority"
	"github.com/marthjod/gotodo/model/project"
	color "github.com/wsxiaoys/terminal/color"
)

var (
	projectsRE   = regexp.MustCompile(project.MatchRegexp)
	contextsRE   = regexp.MustCompile(context.MatchRegexp)
	priorityRE   = regexp.MustCompile(priority.MatchRegexp)
	whitespaceRE = regexp.MustCompile("  +")
)

// Entry represents a single todo.txt file entry
type Entry struct {
	Priority    priority.Priority `json:"priority"`
	Description string            `json:"description"`
	Projects    []project.Project `json:"projects"`
	Contexts    []context.Context `json:"contexts"`
	Done        bool              `json:"done"`
}

// ByPriority implements sort.Interface for []Entry based on the Priority field.
type ByPriority []Entry

func (a ByPriority) Len() int           { return len(a) }
func (a ByPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPriority) Less(i, j int) bool { return a[i].Priority < a[j].Priority }

// Read converts a string into an Entry
func Read(line string) Entry {
	var (
		e        = Entry{}
		prio     string
		contexts []string
		projects []string
	)

	if strings.HasPrefix(line, "x") {
		e.Done = true
		line = line[1:]
		// if !includingDone {
		// 	return Entry{}
		// }
	}

	prio = priorityRE.FindString(line)

	if contexts = contextsRE.FindAllString(line, -1); contexts == nil {
		contexts = []string{}
	}

	if projects = projectsRE.FindAllString(line, -1); projects == nil {
		projects = []string{}
	}

	e.Contexts = context.GetContexts(contexts...)
	e.Projects = project.GetProjects(projects...)
	e.Priority = priority.GetPriority(prio)

	line = clearLine(line, contexts...)
	line = clearLine(line, projects...)
	line = clearLine(line, prio)

	e.Description = strings.TrimSpace(whitespaceRE.ReplaceAllLiteralString(line, " "))

	return e
}

func (e *Entry) String() string {
	var concat = []string{}

	if e.Done {
		concat = append(concat, "x")
	}

	if e.Priority != priority.None {
		concat = append(concat, fmt.Sprintf("(%s)", e.Priority))
	}

	concat = append(concat, e.Description)

	for _, p := range e.Projects {
		concat = append(concat, fmt.Sprintf("+%s", p))
	}

	for _, c := range e.Contexts {
		concat = append(concat, fmt.Sprintf("@%s", c))
	}

	return strings.Join(concat, " ")
}

// ColorString returns colorized string representation of Entry.
func (e *Entry) ColorString() string {
	var concat = []string{}

	if e.Done {
		concat = append(concat, "x")
	}

	if e.Priority != priority.None {
		switch e.Priority {
		case priority.A:
			concat = append(concat, color.Sprintf("@{r}(%s)", e.Priority))
		case priority.B:
			concat = append(concat, color.Sprintf("@{y}(%s)", e.Priority))
		default:
			concat = append(concat, fmt.Sprintf("(%s)", e.Priority))
		}
	}

	concat = append(concat, e.Description)

	for _, p := range e.Projects {
		concat = append(concat, fmt.Sprintf("+%s", p))
	}

	for _, c := range e.Contexts {
		concat = append(concat, fmt.Sprintf("@%s", c))
	}

	return strings.Join(concat, " ")
}

func clearLine(line string, s ...string) string {
	for _, sub := range s {
		line = strings.Replace(line, sub, "", -1)
	}
	return line
}
