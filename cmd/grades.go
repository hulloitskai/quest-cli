package cmd

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/cheynewallace/tabby"
	"github.com/manifoldco/promptui"
	"github.com/stevenxie/quest-cli/internal/interact"
	"github.com/stevenxie/uwquest"
	ess "github.com/unixpickle/essentials"
)

func registerGradesCmd(app *kingpin.Application) {
	gradesCmd = app.Command(
		"grades",
		"List Quest grades for a particular term.",
	)

	// Register flags.
	gradesCmd.Flag("term", "The name of the term to load grades for.").Short('t').
		StringVar(&gradesOpts.Term)
	gradesCmd.Flag("poll", "Repeatedly poll for new grades.").Short('p').
		BoolVar(&gradesOpts.Poll)
}

var (
	gradesCmd  *kingpin.CmdClause
	gradesOpts struct {
		Term string
		Poll bool
	}
)

const (
	gradesMinPollInterval = time.Minute
	gradesMaxPollInterval = 2 * time.Minute
)

func grades() error {
	c, err := interact.BuildClient()
	if err != nil {
		return err
	}

	var rng *rand.Rand
	if gradesOpts.Poll { // initialize only if gradeOpts.Poll is enabled
		src := rand.NewSource(time.Now().Unix())
		rng = rand.New(src)
	}

check:
	interact.Errln("Fetching terms...")
	terms, err := c.Terms()
	if err != nil {
		return ess.AddCtx("fetching terms", err)
	}
	if len(terms) == 0 {
		interact.Errln("No terms were found.")
		os.Exit(2)
	}

	// Filter terms based on opts, if applicable.
	if gradesOpts.Term != "" {
		var termsCopy []*uwquest.Term = terms
		terms = make([]*uwquest.Term, 0, len(termsCopy))

		for _, term := range termsCopy {
			var (
				name  = strings.ToLower(term.Name)
				query = strings.ToLower(gradesOpts.Term)
			)
			if strings.Contains(name, query) {
				terms = append(terms, term)
			}
		}

		if len(terms) == 0 {
			interact.Errf("No terms match the name '%s'\n", gradesOpts.Term)
			os.Exit(3)
		}
	}

	var target *uwquest.Term
	if len(terms) == 1 {
		target = terms[0]
	} else {
		items := make([]string, len(terms))
		for i, term := range terms {
			items[i] = term.Name
		}

		sel := promptui.Select{
			Label: "Select term",
			Items: items,
		}

		index, _, err := sel.Run()
		if err != nil {
			return ess.AddCtx("selecting term", err)
		}

		target = terms[index]
	}

	interact.Errf("Fetching grades for term: '%s'...\n", target.Name)
	grades, err := c.Grades(target.ID)
	if err != nil {
		return ess.AddCtx("fetching grades", err)
	}

	table := tabby.New()
	table.AddHeader("COURSE", "GRADE", "POINTS")
	for _, cg := range grades {
		var points string
		if cg.GradePoints == nil {
			points = "N/A"
		} else {
			points = strconv.FormatFloat(float64(*cg.GradePoints), 'f', 3, 32)
		}

		table.AddLine(cg.Name, cg.Grade, points)
	}
	interact.Errln()
	table.Print()
	interact.Errln()

	if gradesOpts.Poll {
		gradesOpts.Term = target.Name
		interval := time.Duration(rng.Int63n(int64(gradesMaxPollInterval-
			gradesMinPollInterval))) + gradesMinPollInterval

		for interval >= 0 {
			interact.Errf("\rChecking again in %.f seconds (press <ctrl-c> to stop).",
				interval.Seconds())
			time.Sleep(time.Second)
			interval -= time.Second
		}
		interact.Errln()
		goto check
	}
	return nil
}
