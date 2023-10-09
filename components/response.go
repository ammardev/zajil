package components

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alecthomas/chroma/quick"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const responseViewStatusBarOuterHeight = 2

type ResponseView struct {
    textinput.Model
    Width int
    style lipgloss.Style
    statusBarStyle lipgloss.Style
    statusBarRightStyle lipgloss.Style
    statusStyle lipgloss.Style
    body string
    status int
    ttfb int
    viewport viewport.Model
}
var a lipgloss.Style

func NewResponseView() ResponseView {
    v := viewport.New(5, 5)

    response := ResponseView{
        Model: textinput.New(),
        style: lipgloss.NewStyle().Border(lipgloss.RoundedBorder()),
        statusBarStyle: lipgloss.NewStyle().
            Padding(0, 2, 0, 2).
            Border(lipgloss.RoundedBorder(), false, false, true, false),
        statusStyle: lipgloss.NewStyle().
            AlignHorizontal(lipgloss.Center).
            Background(lipgloss.Color("42")).
            Foreground(lipgloss.Color("232")).
            Bold(true),
        statusBarRightStyle: lipgloss.NewStyle().
            AlignHorizontal(lipgloss.Right),
        viewport: v,
    }
    response.Model.Prompt = ""

    return response
}

func (view *ResponseView) SetResponse(response *http.Response, ttfb int) {
    body, _ := io.ReadAll(response.Body)
    view.status = response.StatusCode
    view.ttfb = ttfb

    formattingBuffer := new(bytes.Buffer)
    json.Indent(formattingBuffer, body, "│ ", "\t")

    highlightingBuffer := new(bytes.Buffer)
    quick.Highlight(highlightingBuffer, formattingBuffer.String(), "json", "terminal256", "solarized-dark256")

    view.viewport.SetContent("│ " + highlightingBuffer.String())
}

func (view *ResponseView) Resize(width, height int) {
    view.style.Width(width - borderPadding)
    view.style.Height(height - borderPadding)

    view.viewport.Width = width - borderPadding
    view.viewport.Height = height - borderPadding - responseViewStatusBarOuterHeight

    view.statusBarStyle.Width(width - borderPadding)
    view.statusBarRightStyle.Width(width - borderPadding - 10)
}

func (view *ResponseView) HandleEvents(msg tea.Msg) tea.Cmd {
    var cmd tea.Cmd
    view.viewport, cmd = view.viewport.Update(msg)

    return cmd
}

func (view ResponseView) Render() string {
    statusBar := view.statusBarStyle.Render("Enter a url and then press `enter` to send the request")
    if view.status > 0 {
        statusBar = view.statusBarStyle.Render(
            view.statusStyle.Render(fmt.Sprint(view.status)),
            view.statusBarRightStyle.Render(fmt.Sprint("TTFB: ", view.ttfb, "ms")),
        )
    }

    return view.style.Render(
        statusBar,
        view.viewport.View(),
    )
}
