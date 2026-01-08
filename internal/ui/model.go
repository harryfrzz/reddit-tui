package ui

import (
	"reddit-tui/internal/data"
	"reddit-tui/internal/icons"
	"reddit-tui/internal/models"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	SidebarItems  []string
	Posts         []models.Post
	SidebarCursor int
	PostsCursor   int
	ActivePane    string
	Width         int
	Height        int
	PostsScroll   int
	PreviewScroll int
	SearchQuery   string
	SearchResults []models.Post
	AllPosts      []models.Post
	IsSearching   bool
}

func InitialModel() Model {
	posts, err := data.LoadSamplePosts()
	if err != nil {
		posts = []models.Post{}
	}

	return Model{
		SidebarItems: []string{
			icons.Home + " Home",
			icons.Popular + " Popular",
			icons.Explore + " Explore",
			icons.Settings + " Settings",
			icons.Login + " Login/Auth",
		},
		Posts:         posts,
		AllPosts:      posts,
		SidebarCursor: 0,
		PostsCursor:   0,
		ActivePane:    "sidebar",
		Width:         80,
		Height:        24,
		PostsScroll:   0,
		PreviewScroll: 0,
		SearchQuery:   "",
		SearchResults: []models.Post{},
		IsSearching:   false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			switch m.ActivePane {
			case "sidebar":
				m.ActivePane = "posts"
			case "posts":
				m.ActivePane = "preview"
			case "preview":
				m.ActivePane = "sidebar"
			}
		case "u":
			// Upvote the current post (only in preview pane)
			if m.ActivePane == "preview" && m.PostsCursor >= 0 {
				if m.IsSearching && m.PostsCursor < len(m.SearchResults) {
					m.SearchResults[m.PostsCursor].ToggleUpvote()
				} else if !m.IsSearching && m.PostsCursor < len(m.Posts) {
					m.Posts[m.PostsCursor].ToggleUpvote()
				}
			}
		case "d":
			// Downvote the current post (only in preview pane)
			if m.ActivePane == "preview" && m.PostsCursor >= 0 {
				if m.IsSearching && m.PostsCursor < len(m.SearchResults) {
					m.SearchResults[m.PostsCursor].ToggleDownvote()
				} else if !m.IsSearching && m.PostsCursor < len(m.Posts) {
					m.Posts[m.PostsCursor].ToggleDownvote()
				}
			}
		case "enter":
			// When in sidebar, select the section
			if m.ActivePane == "sidebar" {
				if m.SidebarCursor == 2 { // Explore is at index 2
					m.IsSearching = true
					m.PostsCursor = 0
					m.PostsScroll = 0
				} else {
					m.IsSearching = false
					m.Posts = m.AllPosts
					m.PostsCursor = 0
					m.PostsScroll = 0
				}
				m.ActivePane = "posts"
			}
		case "esc":
			// Exit search mode
			if m.IsSearching && m.ActivePane == "posts" {
				m.SearchQuery = ""
				m.SearchResults = []models.Post{}
			}
		case "backspace":
			// Handle backspace in search mode
			if m.IsSearching && m.ActivePane == "posts" && len(m.SearchQuery) > 0 {
				m.SearchQuery = m.SearchQuery[:len(m.SearchQuery)-1]
				m.performSearch()
				m.PostsCursor = 0
				m.PostsScroll = 0
			}
		case "up", "k":
			if m.ActivePane == "sidebar" {
				if m.SidebarCursor > 0 {
					m.SidebarCursor--
				}
			} else if m.ActivePane == "posts" {
				// Only navigate if not in search input mode
				if !m.IsSearching {
					if m.PostsCursor > 0 {
						m.PostsCursor--
						m.PreviewScroll = 0
						if m.PostsCursor < m.PostsScroll {
							m.PostsScroll = m.PostsCursor
						}
					}
				} else {
					// In search mode, navigate through results
					if m.PostsCursor > 0 {
						m.PostsCursor--
						m.PreviewScroll = 0
						if m.PostsCursor < m.PostsScroll {
							m.PostsScroll = m.PostsCursor
						}
					}
				}
			} else if m.ActivePane == "preview" {
				if m.PreviewScroll > 0 {
					m.PreviewScroll--
				}
			}
		case "down", "j":
			if m.ActivePane == "sidebar" {
				if m.SidebarCursor < len(m.SidebarItems)-1 {
					m.SidebarCursor++
				}
			} else if m.ActivePane == "posts" {
				// Determine which list to use
				postsList := m.Posts
				if m.IsSearching {
					postsList = m.SearchResults
				}
				if !m.IsSearching {
					if m.PostsCursor < len(postsList)-1 {
						m.PostsCursor++
						m.PreviewScroll = 0
						visiblePosts := (m.Height - 3 - 4) / 4
						if visiblePosts < 1 {
							visiblePosts = 1
						}
						if m.PostsCursor >= m.PostsScroll+visiblePosts {
							m.PostsScroll = m.PostsCursor - visiblePosts + 1
						}
					}
				} else {
					// In search mode, navigate through results
					if m.PostsCursor < len(postsList)-1 {
						m.PostsCursor++
						m.PreviewScroll = 0
						visiblePosts := (m.Height - 3 - 4) / 4
						if visiblePosts < 1 {
							visiblePosts = 1
						}
						if m.PostsCursor >= m.PostsScroll+visiblePosts {
							m.PostsScroll = m.PostsCursor - visiblePosts + 1
						}
					}
				}
			} else if m.ActivePane == "preview" {
				m.PreviewScroll++
			}
		default:
			if m.IsSearching && m.ActivePane == "posts" {
				if len(msg.String()) == 1 {
					r := []rune(msg.String())[0]
					if r >= 32 && r < 127 { // Printable ASCII
						m.SearchQuery += msg.String()
						m.performSearch()
						m.PostsCursor = 0
						m.PostsScroll = 0
					}
				}
			}
		}
	}

	return m, nil
}

func (m *Model) performSearch() {
	if m.SearchQuery == "" {
		m.SearchResults = []models.Post{}
		return
	}

	query := strings.ToLower(m.SearchQuery)
	m.SearchResults = []models.Post{}

	for _, post := range m.AllPosts {
		if strings.Contains(strings.ToLower(post.Title), query) ||
			strings.Contains(strings.ToLower(post.Subreddit), query) ||
			strings.Contains(strings.ToLower(post.Author), query) {
			m.SearchResults = append(m.SearchResults, post)
		}
	}
}
