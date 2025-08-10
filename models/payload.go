package models

import "time"

type User struct {
	Login string `json:"login"`
}

type Repository struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Commit struct {
	Sha      string `json:"sha"`
	Author   Author `json:"author"`
	Message  string `json:"message"`
	Distinct bool   `json:"distinct"`
	Url      string `json:"url"`
}

type Author struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Issue struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type IssueComment struct {
	Body string `json:"body"`
}

type PullRequest struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

type Release struct {
	TagName    string `json:"tag_name"`
	Name       string `json:"name"`
	Draft      bool   `json:"draft"`
	Prerelease bool   `json:"prerelease"`
}

type PLCreate struct {
	Ref      string `json:"ref"`
	RefType  string `json:"ref_type"`
	Branch   string `json:"master_branch"`
	Desc     string `json:"description"`
	Pusher   string `json:"pusher_type"`
}

type PLPush struct {
	RepoID   uint64    `json:"repository_id"`
	PushID   uint64    `json:"push_id"`
	Size     uint64    `json:"size"`
	DistSize uint64    `json:"distinct_size"`
	Ref      string    `json:"ref"`
	Head     string    `json:"head"`
	Before   string    `json:"before"`
	Commits  []Commit  `json:"commits"`
	Public   bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}

type PLDelete struct {
	Ref     string `json:"ref"`
	RefType string `json:"ref_type"`
}

type PLFork struct {
	Forkee Repository `json:"forkee"`
}

type PLGollum struct {
	Pages []struct {
		PageName string `json:"page_name"`
		Title    string `json:"title"`
		Action   string `json:"action"`
	} `json:"pages"`
}

type PLIssues struct {
	Action string `json:"action"`
	Issue  Issue  `json:"issue"`
}

type PLIssueComment struct {
	Action string       `json:"action"`
	Issue  Issue        `json:"issue"`
	Comment IssueComment `json:"comment"`
}

type PLMember struct {
	Action string `json:"action"`
	Member User   `json:"member"`
}

type PLPublic struct {
	Action string `json:"action"`
}

type PLPullRequest struct {
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
}

type PLPullRequestReview struct {
	Action      string      `json:"action"`
	Review      struct {
		State string `json:"state"`
	} `json:"review"`
	PullRequest PullRequest `json:"pull_request"`
}

type PLPullRequestReviewComment struct {
	Action string       `json:"action"`
	Comment IssueComment `json:"comment"`
	PullRequest PullRequest `json:"pull_request"`
}

type PLRelease struct {
	Action  string  `json:"action"`
	Release Release `json:"release"`
}

type PLSponsorship struct {
	Action      string `json:"action"`
	Sponsorable User   `json:"sponsoreable"`
	Sponsor     User   `json:"sponsor"`
}

type PLWatch struct {
	Action string `json:"action"`
}
