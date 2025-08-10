package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"github.com/urfave/cli/v3"
	"github.com/louis0113/github-activity/models"
)

var gitStats []models.GitStat

func unmarshalPayload[T any](payload json.RawMessage) (*T, error) {
	var value T
	if err := json.Unmarshal(payload, &value); err != nil {
		return nil, err
	}
	return &value, nil
}

func printEventInfo(stat models.GitStat, payload any) {
	switch p := payload.(type) {
	case *models.PLCreate:

    if p.Ref != "" {
		fmt.Printf("[%s] %s created %s '%s' in repository %s\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Actor.Login,
			p.RefType,
			p.Ref,
			stat.Repo.Name)
    } else {
		fmt.Printf("[%s] %s created %s %s\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Actor.Login,
			p.RefType,
			stat.Repo.Name)
    }

	case *models.PLPush:
		fmt.Printf("[%s] %s pushed %d commits to %s in repository %s\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Actor.Login,
			p.Size,
			p.Ref,
			stat.Repo.Name)
	case *models.PLDelete:
		fmt.Printf("[%s] %s deleted %s '%s' in repository %s\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Actor.Login,
			p.RefType,
			p.Ref,
			stat.Repo.Name)
	case *models.PLFork:
		fmt.Printf("[%s] %s forked repository '%s' to '%s'\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Actor.Login,
			stat.Repo.Name,
			p.Forkee.Name)
	case *models.PLGollum:
		for _, page := range p.Pages {
			fmt.Printf("[%s] %s %s wiki page '%s' in repository %s\n",
				stat.CreatedAt.Format("2006-01-02 15:04"),
				stat.Actor.Login,
				page.Action,
				page.Title,
				stat.Repo.Name)
		}
	case *models.PLIssues:
		fmt.Printf("[%s] %s %s issue '%s' (#%d) in repository %s\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Actor.Login,
			p.Action,
			p.Issue.Title,
			p.Issue.Number,
			stat.Repo.Name)
	case *models.PLIssueComment:
		fmt.Printf("[%s] %s %s on issue '%s' (#%d) in repository %s\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Actor.Login,
			p.Action,
			p.Issue.Title,
			p.Issue.Number,
			stat.Repo.Name)
	case *models.PLMember:
		fmt.Printf("[%s] %s %s user '%s' to repository %s\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Actor.Login,
			p.Action,
			p.Member.Login,
			stat.Repo.Name)
	case *models.PLPublic:
		fmt.Printf("[%s] %s made repository '%s' public\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Actor.Login,
			stat.Repo.Name)
	case *models.PLPullRequest:
		fmt.Printf("[%s] %s %s pull request '%s' (#%d) in repository %s\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Actor.Login,
			p.Action,
			p.PullRequest.Title,
			p.PullRequest.Number,
			stat.Repo.Name)
	case *models.PLPullRequestReview:
		fmt.Printf("[%s] %s %s a pull request review on '%s' (#%d) in repository %s\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Actor.Login,
			p.Action,
			p.PullRequest.Title,
			p.PullRequest.Number,
			stat.Repo.Name)
	case *models.PLPullRequestReviewComment:
		fmt.Printf("[%s] %s %s a review comment on pull request '%s' (#%d) in repository %s\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Actor.Login,
			p.Action,
			p.PullRequest.Title,
			p.PullRequest.Number,
			stat.Repo.Name)
	case *models.PLRelease:
		fmt.Printf("[%s] %s %s release '%s' in repository %s\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Actor.Login,
			p.Action,
			p.Release.TagName,
			stat.Repo.Name)
	case *models.PLSponsorship:
		fmt.Printf("[%s] %s %s sponsorship for user '%s'\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Actor.Login,
			p.Action,
			p.Sponsorable.Login)
	case *models.PLWatch:
		fmt.Printf("[%s] %s %s repository in repository %s\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Actor.Login,
			p.Action,
			stat.Repo.Name)
	default:
		fmt.Printf("[%s] Unhandled event type: %s\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Type)
	}
}

func handleEvent(stat models.GitStat) {
	switch stat.Type {
	case "CreateEvent":
		payload, err := unmarshalPayload[models.PLCreate](stat.Payload)
		if err != nil {
			log.Printf("Error decoding CreateEvent payload: %v\n", err)
			return
		}
		printEventInfo(stat, payload)
	case "PushEvent":
		payload, err := unmarshalPayload[models.PLPush](stat.Payload)
		if err != nil {
			log.Printf("Error decoding PushEvent payload: %v\n", err)
			return
		}
		printEventInfo(stat, payload)
	case "DeleteEvent":
		payload, err := unmarshalPayload[models.PLDelete](stat.Payload)
		if err != nil {
			log.Printf("Error decoding DeleteEvent payload: %v\n", err)
			return
		}
		printEventInfo(stat, payload)
	case "ForkEvent":
		payload, err := unmarshalPayload[models.PLFork](stat.Payload)
		if err != nil {
			log.Printf("Error decoding ForkEvent payload: %v\n", err)
			return
		}
		printEventInfo(stat, payload)
	case "GollumEvent":
		payload, err := unmarshalPayload[models.PLGollum](stat.Payload)
		if err != nil {
			log.Printf("Error decoding GollumEvent payload: %v\n", err)
			return
		}
		printEventInfo(stat, payload)
	case "IssuesEvent":
		payload, err := unmarshalPayload[models.PLIssues](stat.Payload)
		if err != nil {
			log.Printf("Error decoding IssuesEvent payload: %v\n", err)
			return
		}
		printEventInfo(stat, payload)
	case "IssueCommentEvent":
		payload, err := unmarshalPayload[models.PLIssueComment](stat.Payload)
		if err != nil {
			log.Printf("Error decoding IssueCommentEvent payload: %v\n", err)
			return
		}
		printEventInfo(stat, payload)
	case "MemberEvent":
		payload, err := unmarshalPayload[models.PLMember](stat.Payload)
		if err != nil {
			log.Printf("Error decoding MemberEvent payload: %v\n", err)
			return
		}
		printEventInfo(stat, payload)
	case "PublicEvent":
		payload, err := unmarshalPayload[models.PLPublic](stat.Payload)
		if err != nil {
			log.Printf("Error decoding PublicEvent payload: %v\n", err)
			return
		}
		printEventInfo(stat, payload)
	case "PullRequestEvent":
		payload, err := unmarshalPayload[models.PLPullRequest](stat.Payload)
		if err != nil {
			log.Printf("Error decoding PullRequestEvent payload: %v\n", err)
			return
		}
		printEventInfo(stat, payload)
	case "PullRequestReviewEvent":
		payload, err := unmarshalPayload[models.PLPullRequestReview](stat.Payload)
		if err != nil {
			log.Printf("Error decoding PullRequestReviewEvent payload: %v\n", err)
			return
		}
		printEventInfo(stat, payload)
	case "PullRequestReviewCommentEvent":
		payload, err := unmarshalPayload[models.PLPullRequestReviewComment](stat.Payload)
		if err != nil {
			log.Printf("Error decoding PullRequestReviewCommentEvent payload: %v\n", err)
			return
		}
		printEventInfo(stat, payload)
	case "ReleaseEvent":
		payload, err := unmarshalPayload[models.PLRelease](stat.Payload)
		if err != nil {
			log.Printf("Error decoding ReleaseEvent payload: %v\n", err)
			return
		}
		printEventInfo(stat, payload)
	case "SponsorshipEvent":
		payload, err := unmarshalPayload[models.PLSponsorship](stat.Payload)
		if err != nil {
			log.Printf("Error decoding SponsorshipEvent payload: %v\n", err)
			return
		}
		printEventInfo(stat, payload)
	case "WatchEvent":
		payload, err := unmarshalPayload[models.PLWatch](stat.Payload)
		if err != nil {
			log.Printf("Error decoding WatchEvent payload: %v\n", err)
			return
		}
		printEventInfo(stat, payload)
	default:
		fmt.Printf("[%s] Unhandled event type: %s\n",
			stat.CreatedAt.Format("2006-01-02 15:04"),
			stat.Type)
	}
}

func main() {
	cmd := &cli.Command{
		Name:  "github-activity",
		Usage: "A simple program for show your recent github activity",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.NArg() != 1 {
				return cli.Exit("Your username is required as the only argument", 86)
			}

			username := cmd.Args().First()
			end := fmt.Sprintf("https://api.github.com/users/%s/events", username)

			resp, err := http.Get(end)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			err = json.Unmarshal(body, &gitStats)
			if err != nil {
				log.Fatal(err)
			}

			for _, stat := range gitStats {
				handleEvent(stat)
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
