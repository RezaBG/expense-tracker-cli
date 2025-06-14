# GitHub Activity CLI (Golang)

Simple command-line tool built with Go to fetch recent public GitHub activity for any user.

---

## Features

- Fetch recent activity for any GitHub username
- Supports activity types: Push, Pull Request, Issues, Watch, etc.
- Displays activity in human-readable format
- Handles invalid usernames or API errors gracefully
- Uses GitHub public API
- Fully written in Go using only standard library

---

## Usage

### Build & Run

```bash
# Clone the repo
git clone https://github.com/your-username/github-activity-cli.git
cd github-activity-cli

# Run
go run . <username>

# Example:
go run . kamranahmedse
```

### Output Example

- Pushed code to kamranahmedse/developer-roadmap
- Opened a pull request on kamranahmedse/developer-roadmap
- Starred kamranahmedse/developer-roadmap

## Project Submission

This project was built as part of the [Github Activity CLI project challenge on roadmap.sh](https://roadmap.sh/projects/github-user-activity).
