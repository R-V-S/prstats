package github

import (
	"context"
	"os"

	"github.com/R-V-S/prstats/constructs/abstract/prstat"
	// "github.com/fatih/color"
	"github.com/google/go-github/v29/github"
	// json "github.com/nwidger/jsoncolor"
	"golang.org/x/oauth2"
)

// FetchPRStats ...
func FetchPRStats(prStatAggregate *prstat.Aggregate, ticketid string) error {
	githubToken := os.Getenv("GITHUB_API_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	issues, _, err := client.Search.Issues(context.Background(), ticketid+" in:title org:DailyFeats type:pr", &github.SearchOptions{TextMatch: true})
	if err != nil {
		return err
	}

	prStatAggregate.PRCount = len(issues.Issues)
	org := "DailyFeats"     // TODO: get from response
	repo := "maxwell-titan" // TODO: get from response
	reviewCommentCount := 0
	for _, issue := range issues.Issues {
		prNumber := issue.Number
		pullRequest, _, err := client.PullRequests.Get(context.Background(), org, repo, *prNumber)
		if err != nil {
			return err
		}
		// reviews, _, err := client.PullRequests.ListReviews(context.Background(), "DailyFeats", "maxwell-titan", *prNumber, &github.ListOptions{})
		reviewCommentCount = reviewCommentCount + *pullRequest.ReviewComments
	}
	prStatAggregate.AvgReviewCommentCount = reviewCommentCount / prStatAggregate.PRCount

	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	// create custom formatter
	// f := json.NewFormatter()

	// set custom colors
	// f.StringColor = color.New(color.FgHiGreen, color.Bold)
	// f.TrueColor = color.New(color.FgWhite, color.Bold)
	// f.FalseColor = color.New(color.FgRed)
	// f.NumberColor = color.New(color.FgWhite)
	// f.NullColor = color.New(color.FgWhite, color.Bold)

	// marshal v with custom formatter,
	// dst contains colorized output
	// dst, err := json.MarshalIndentWithFormatter(issues, "", "  ", f)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// print colorized output to stdout
	// fmt.Println(string(dst))

	return err
}
