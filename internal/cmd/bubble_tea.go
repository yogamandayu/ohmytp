package cmd

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
	"log"
	"strings"
)

const listHeight = 14

type Item string

func (i Item) FilterValue() string { return "" }

type BubbleTea struct {
	Model       BubbleTeaModel
	Style       BubbleTeaStyle
	Command     func(args string) error
	MapArgsItem map[string]string
}

type BubbleTeaModel struct {
	List   list.Model
	Choice string
	IsQuit bool
}

type BubbleTeaStyle struct {
	Title        lipgloss.Style
	Item         lipgloss.Style
	Pagination   lipgloss.Style
	SelectedItem lipgloss.Style
	Help         lipgloss.Style
	QuitText     lipgloss.Style
}

func NewBubbleTea() BubbleTea {
	b := BubbleTea{
		Style: BubbleTeaStyle{
			Title:        lipgloss.NewStyle().MarginLeft(2),
			Item:         lipgloss.NewStyle().PaddingLeft(4),
			SelectedItem: lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170")),
			Pagination:   list.DefaultStyles().PaginationStyle.PaddingLeft(4),
			Help:         list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1),
			QuitText:     lipgloss.NewStyle().Margin(1, 0, 2, 4),
		},
	}

	var items []list.Item

	const defaultWidth = 20

	l := list.New(items, ItemDelegate{
		ItemStyle:         b.Style.Item,
		SelectedItemStyle: b.Style.SelectedItem,
	}, defaultWidth, listHeight)
	l.Title = "What program you want to run?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = b.Style.Title
	l.Styles.PaginationStyle = b.Style.Pagination
	l.Styles.HelpStyle = b.Style.Help
	b.Model.List = l

	return b
}

func (b *BubbleTea) SetListItemAndArgs(l map[string]string) {
	var items []list.Item
	for item := range l {
		items = append(items, Item(item))
	}
	b.MapArgsItem = l
	b.Model.List.SetItems(items)
}

func (b *BubbleTea) SetCommand(command func(args string) error) {
	b.Command = command
}

func (b BubbleTea) Init() tea.Cmd {
	return nil
}

func (b BubbleTea) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.Model.List.SetWidth(msg.Width)
		return b, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			b.Model.IsQuit = true
			return b, tea.Quit

		case "enter":
			i, ok := b.Model.List.SelectedItem().(Item)
			if ok {
				b.Model.Choice = string(i)
			}
			return b, tea.Quit
		}
	}

	var cmd tea.Cmd
	b.Model.List, cmd = b.Model.List.Update(msg)
	return b, cmd
}

func (b BubbleTea) View() string {
	if b.Model.Choice != "" {
		err := b.Command(b.MapArgsItem[b.Model.Choice])
		if err != nil {
			log.Fatalf("Unable to run CLI command, err: %v", err)
		}
	}
	if b.Model.IsQuit {
		return b.Style.QuitText.Render("Program exit")
	}
	return "\n" + b.Model.List.View()
}

type ItemDelegate struct {
	ItemStyle         lipgloss.Style
	SelectedItemStyle lipgloss.Style
}

func (d ItemDelegate) Height() int                             { return 1 }
func (d ItemDelegate) Spacing() int                            { return 0 }
func (d ItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}
	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := d.ItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return d.SelectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	_, _ = fmt.Fprint(w, fn(str))
}
