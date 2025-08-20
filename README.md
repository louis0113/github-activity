# github-activity

A simple program to show recent GitHub activity.

## Overview

`github-activity` is a command-line tool written in Go that retrieves and displays a user's recent public activity from GitHub's API. It provides a human-readable summary of various GitHub events such as pushes, forks, issues, and more.

## Installation

### Prerequisites

  * `Go (version 1.22.2 or later)`

### Steps

1.  Clone the repository:
    ```bash
    git clone https://github.com/louis0113/github-activity.git
    cd github-activity
    ```
2.  Build the executable:
    ```bash
    go build -o github-activity
    ```

## Usage

Run the program from your terminal with your GitHub username as the only argument:

```bash
./github-activity <your-username>
```

### Example Output

The output provides a timeline of events with the timestamp, actor, and a description of the action.

  * `[2025-08-20 19:17] octocat pushed 2 commits to main in repository octocat/Spoon-Knife`
  * `[2025-08-20 19:16] octocat created repository octocat/Hello-World`
  * `[2025-08-20 19:15] octocat opened issue 'Found a bug' (#1) in repository octocat/Hello-World`

## Supported Events

The program can handle and display information for a variety of GitHub events, including:

  * **CreateEvent**: When a branch, tag, or repository is created.
  * **PushEvent**: When one or more commits are pushed to a repository.
  * **DeleteEvent**: When a branch or tag is deleted.
  * **ForkEvent**: When a repository is forked.
  * **GollumEvent**: When a Wiki page is created or updated.
  * **IssuesEvent**: When an issue is opened, closed, or otherwise modified.
  * **IssueCommentEvent**: When a comment is made on an issue.
  * **MemberEvent**: When a user is added to a repository.
  * **PublicEvent**: When a private repository is made public.
  * **PullRequestEvent**: When a pull request is opened, closed, or synchronized.
  * **PullRequestReviewEvent**: When a review is submitted on a pull request.
  * **PullRequestReviewCommentEvent**: When a comment is made on a pull request review.
  * **ReleaseEvent**: When a release is published or updated.
  * **SponsorshipEvent**: When a sponsorship is created or cancelled for a user.
  * **WatchEvent**: When a user stars a repository.

If an event type is not recognized, it will be marked as "Unhandled".

## Dependencies

The project uses the following Go modules:

  * `github.com/urfave/cli/v3 v3.3.8` 

## License

This project is licensed under the MIT License.

-----

## File Descriptions

  * `github-activity.go`: The main application file. It contains the logic for fetching data from the GitHub API, parsing the JSON response, and printing the formatted output. It uses the `urfave/cli/v3` package to handle command-line arguments.
  * `gitstat.go`: Defines the core data structures for a GitHub event, including `GitStat`, `Actor`, and `Repo`.
  * `payload.go`: Defines the data structures for the specific payloads of each GitHub event type, such as `PLCreate`, `PLPush`, `PLIssues`, etc..
  * `go.mod`: Specifies the module path and the required dependencies for the project.
  * `go.sum`: Contains cryptographic hashes for the dependencies to ensure their integrity.
  * `LICENSE.md`: The license file for the project.
