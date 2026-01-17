package ui

import (
	"reddit-tui/internal/data"
	"reddit-tui/internal/icons"
	"reddit-tui/internal/models"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	SidebarItems   []string
	Posts          []models.Post
	SidebarCursor  int
	PostsCursor    int
	ActivePane     string
	Width          int
	Height         int
	PostsScroll    int
	PreviewScroll  int
	SearchQuery    string
	SearchResults  []models.Post
	AllPosts       []models.Post
	IsSearching    bool
	ShowSettings   bool
	SettingsCursor int
	APIKey         string
	ClientSecret   string
	EditingField   int
	ShowHelp       bool
}

func InitialModel() Model {
	posts, err := data.LoadSamplePosts()
	if err != nil {
		posts = []models.Post{}
	}

	return Model{
		SidebarItems: []string{
			icons.Home + " Home",
			icons.Explore + " Explore",
			icons.Settings + " Settings",
		},
		Posts:          posts,
		AllPosts:       posts,
		SidebarCursor:  0,
		PostsCursor:    0,
		ActivePane:     "sidebar",
		Width:          80,
		Height:         24,
		PostsScroll:    0,
		PreviewScroll:  0,
		SearchQuery:    "",
		SearchResults:  []models.Post{},
		IsSearching:    false,
		ShowSettings:   false,
		SettingsCursor: 0,
		APIKey:         "",
		ClientSecret:   "",
		EditingField:   0,
		ShowHelp:       false,
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
		case "ctrl+c":
			return m, tea.Quit
		case "?":
			m.ShowHelp = !m.ShowHelp
		case "q":
			if m.EditingField != 0 || (m.IsSearching && m.ActivePane == "posts") {
				if m.ShowSettings && m.EditingField != 0 && m.ActivePane == "posts" {
					r := []rune(msg.String())[0]
					if r >= 32 && r < 127 {
						if m.EditingField == 1 {
							m.APIKey += msg.String()
						} else if m.EditingField == 2 {
							m.ClientSecret += msg.String()
						}
					}
				} else if m.IsSearching && m.ActivePane == "posts" {
					m.SearchQuery += msg.String()
					m.performSearch()
					m.PostsCursor = 0
					m.PostsScroll = 0
				}
			} else {
				return m, tea.Quit
			}
		case "tab":
			if m.ShowSettings {
				switch m.ActivePane {
				case "sidebar":
					m.ActivePane = "posts"
				case "posts":
					m.ActivePane = "sidebar"
				default:
					m.ActivePane = "sidebar"
				}
			} else {
				// Normal three-pane cycling
				switch m.ActivePane {
				case "sidebar":
					m.ActivePane = "posts"
				case "posts":
					m.ActivePane = "preview"
				case "preview":
					m.ActivePane = "sidebar"
				}
			}
		case "u":
			if m.ActivePane == "preview" && m.PostsCursor >= 0 {
				if m.IsSearching && m.PostsCursor < len(m.SearchResults) {
					m.SearchResults[m.PostsCursor].ToggleUpvote()
				} else if !m.IsSearching && m.PostsCursor < len(m.Posts) {
					m.Posts[m.PostsCursor].ToggleUpvote()
				}
			}
		case "d":
			if m.ActivePane == "preview" && m.PostsCursor >= 0 {
				if m.IsSearching && m.PostsCursor < len(m.SearchResults) {
					m.SearchResults[m.PostsCursor].ToggleDownvote()
				} else if !m.IsSearching && m.PostsCursor < len(m.Posts) {
					m.Posts[m.PostsCursor].ToggleDownvote()
				}
			}
		case "enter":
			if m.ActivePane == "sidebar" {
				switch m.SidebarCursor {
				case 0: // Home
					m.IsSearching = false
					m.ShowSettings = false
					m.Posts = m.AllPosts
					m.PostsCursor = 0
					m.PostsScroll = 0
					m.ActivePane = "posts"
				case 1: // Explore
					m.IsSearching = true
					m.ShowSettings = false
					m.PostsCursor = 0
					m.PostsScroll = 0
					m.ActivePane = "posts"
				case 2: // Settings
					m.ShowSettings = true
					m.IsSearching = false
					m.SettingsCursor = 0
					m.EditingField = 0
					m.ActivePane = "posts"
				}
			} else if m.ActivePane == "posts" && m.ShowSettings {
				if m.EditingField == 0 {
					m.EditingField = m.SettingsCursor + 1
				} else {
					m.EditingField = 0
				}
			}
		case "esc":
			if m.ShowSettings && m.EditingField != 0 {
				m.EditingField = 0
			} else if m.IsSearching && m.ActivePane == "posts" {
				m.SearchQuery = ""
				m.SearchResults = []models.Post{}
			}
		case "backspace":
			if m.ShowSettings && m.EditingField == 1 && len(m.APIKey) > 0 {
				m.APIKey = m.APIKey[:len(m.APIKey)-1]
			} else if m.ShowSettings && m.EditingField == 2 && len(m.ClientSecret) > 0 {
				m.ClientSecret = m.ClientSecret[:len(m.ClientSecret)-1]
			} else if m.IsSearching && m.ActivePane == "posts" && len(m.SearchQuery) > 0 {
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
			} else if m.ActivePane == "posts" && m.ShowSettings && m.EditingField == 0 {
				if m.SettingsCursor > 0 {
					m.SettingsCursor--
				}
			} else if m.ActivePane == "posts" {
				if !m.IsSearching {
					if m.PostsCursor > 0 {
						m.PostsCursor--
						m.PreviewScroll = 0
						if m.PostsCursor < m.PostsScroll {
							m.PostsScroll = m.PostsCursor
						}
					}
				} else {
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
			} else if m.ActivePane == "posts" && m.ShowSettings && m.EditingField == 0 {
				if m.SettingsCursor < 1 {
					m.SettingsCursor++
				}
			} else if m.ActivePane == "posts" {
				postsList := m.Posts
				if m.IsSearching {
					postsList = m.SearchResults
				}
				if !m.IsSearching {
					if m.PostsCursor < len(postsList)-1 {
						m.PostsCursor++
						m.PreviewScroll = 0
						visiblePosts := (m.Height - 3 - 4) / 5
						if visiblePosts < 1 {
							visiblePosts = 1
						}
						if m.PostsCursor >= m.PostsScroll+visiblePosts {
							m.PostsScroll = m.PostsCursor - visiblePosts + 1
						}
					}
				} else {
					if m.PostsCursor < len(postsList)-1 {
						m.PostsCursor++
						m.PreviewScroll = 0
						visiblePosts := (m.Height - 3 - 8) / 5
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
			if m.ShowSettings && m.EditingField != 0 && m.ActivePane == "posts" {
				if len(msg.String()) == 1 {
					r := []rune(msg.String())[0]
					if r >= 32 && r < 127 {
						if m.EditingField == 1 {
							m.APIKey += msg.String()
						} else if m.EditingField == 2 {
							m.ClientSecret += msg.String()
						}
					}
				}
			} else if m.IsSearching && m.ActivePane == "posts" {
				if len(msg.String()) == 1 {
					r := []rune(msg.String())[0]
					if r >= 32 && r < 127 {
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
