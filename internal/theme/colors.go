package theme

import "github.com/charmbracelet/lipgloss"

const (
	ElectricBlue = "#0080FF"
	Purple       = "#9D7BFF"
	LightBlue    = "#49b9ff"
	Magenta      = "#FF00FF"
	SlateBlue    = "#7B68EE"
)

var (
	// Logo
	LogoColor = lipgloss.Color(ElectricBlue)

	// Headings
	PaneHeadingColor = lipgloss.Color(ElectricBlue)

	// Borders
	DefaultBorderColor = lipgloss.Color(Purple)
	ActiveBorderColor  = lipgloss.Color(LightBlue)

	// Sidebar
	SidebarItemColor       = lipgloss.Color(Purple)
	SidebarItemActiveColor = lipgloss.Color(LightBlue)

	// Posts
	PostTitleColor         = lipgloss.Color(Purple)
	PostTitleSelectedColor = lipgloss.Color(LightBlue)
	PostBorderColor        = lipgloss.Color(Purple)
	PostBorderActiveColor  = lipgloss.Color(LightBlue)
	SubredditColor         = lipgloss.Color(Magenta)
	MetaTextColor          = lipgloss.Color(SlateBlue)

	// Preview
	PreviewTitleColor     = lipgloss.Color(LightBlue)
	PreviewSubredditColor = lipgloss.Color(Magenta)
	PreviewMetaColor      = lipgloss.Color(SlateBlue)

	// Votes
	VoteDefaultColor    = lipgloss.Color(SlateBlue)
	VoteUpActiveColor   = lipgloss.Color(LightBlue)
	VoteDownActiveColor = lipgloss.Color(Magenta)
	VoteCountColor      = lipgloss.Color(Purple)
	VoteHintColor       = lipgloss.Color(SlateBlue)

	// Search
	SearchBorderColor = lipgloss.Color(LightBlue)
	SearchIconColor   = lipgloss.Color(ElectricBlue)
	SearchCursorColor = lipgloss.Color(LightBlue)
	SearchHintColor   = lipgloss.Color(SlateBlue)
	NoResultsColor    = lipgloss.Color(SlateBlue)

	// Settings
	SettingsLabelColor       = lipgloss.Color("252")
	SettingsInputBorderColor = lipgloss.Color(Purple)
	SettingsInputActiveColor = lipgloss.Color(LightBlue)
	SettingsCursorColor      = lipgloss.Color(LightBlue)
	SettingsPlaceholderColor = lipgloss.Color(SlateBlue)
	SettingsHintColor        = lipgloss.Color(SlateBlue)

	// Help Modal
	HelpModalTitleColor  = lipgloss.Color(ElectricBlue)
	HelpModalBorderColor = lipgloss.Color(LightBlue)
	HelpModalKeyColor    = lipgloss.Color(LightBlue)
	HelpModalDescColor   = lipgloss.Color("252")
	HelpModalHeaderColor = lipgloss.Color(Purple)
	HelpModalBgColor     = lipgloss.Color("#1a1a1a")
)
