# Reddit TUI

**Reddit TUI** is a lightweight, fast, and fully-featured terminal-based Reddit client written in **Go**. It brings the "Front Page of the Internet" directly to your command line, offering a distraction-free, keyboard-driven experience for power users.

Whether you are browsing on a remote server via SSH, looking to save system resources, or simply prefer the efficiency of the terminal, Reddit TUI provides a seamless interface to browse subreddits, read discussions, and interact with the community.

---

## Features

* **Blazing Fast:** Optimized for performance with minimal resource footprint.
* **Vim-Style Navigation:** Navigate posts and comments without ever touching your mouse.
* **Secure Authentication:** Log in via OAuth2 to access your subscribed subreddits and feed.

---

## Installation

### Prerequisites

* **Go 1.21+** (Ensure you have a recent version of Go installed)
* Terminal with TrueColor support (recommended).

### Install via Go

You can install `reddit-tui` directly using the `go install` command:

```bash
go install [github.com/harryfrzz/reddit-tui@latest](https://github.com/harryfrzz/reddit-tui@latest)

```

---

### Keybindings

**Reddit TUI** uses intuitive, Vim-inspired keybindings for navigation.

| Key | Action |
| --- | --- |
| `j` / `↓` | Move selection down |
| `k` / `↑` | Move selection up |
| `l` / `Enter` | Open post / Expand comment |
| `h` / `Esc` | Go back / Collapse comment |
| `u` | Upvote |
| `d` | Downvote |
| `q` | Quit application |

> **Note:** Keybindings can be customized in the configuration file.

---

## Tech Stack

* **Language:** Go
* **TUI Library:** Bubbletea / Lipgloss
* **API Wrapper:** Go Standard Library / Reddit API

---

## Contributing

Contributions are welcome! If you have a feature request or want to report a bug, please open an issue.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---
# Reddit TUI – TODO / Roadmap

## Core Features
- [ ] Comment Reply Support
- [ ] Post Submission
- [ ] Subreddit Search
- [ ] Post Sorting Options
- [ ] User Profile View
- [ ] Saved Posts and Comments
- [ ] Inbox and Notifications

## UI / UX
- [ ] Theme System
- [ ] Preview Mode
- [ ] Image Preview Support
- [ ] Improved Markdown Rendering
- [ ] Status Bar Enhancements
- [ ] Smooth Transitions
- [ ] Loading Skeletons

## Customization
- [ ] Config File Support
- [ ] Keybinding Remapping
- [ ] Keybinding Profiles
- [ ] Startup Arguments
- [ ] Default Subreddit Configuration

## Authentication and Accounts
- [ ] Multiple Account Support
- [ ] Token Auto Refresh
- [ ] Anonymous Read Mode

---

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=harryfrzz/reddit-tui&type=date&legend=top-left)](https://www.star-history.com/#harryfrzz/reddit-tui&type=date&legend=top-left)
---

## License

Distributed under the MIT License. See `LICENSE` for more information.
