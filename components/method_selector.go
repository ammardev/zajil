package components

import (
	"net/http"

	"github.com/charmbracelet/lipgloss"
)

var (
    methods = []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodOptions}
)


const methodSelectorInnerWidth = 9

type MethodSelector struct {
    selectedMethod int
    style lipgloss.Style
    Width int
}

func NewMethodSelector() MethodSelector {
    return MethodSelector{
        selectedMethod: 0,
        Width: methodSelectorInnerWidth + borderPadding,
        style: lipgloss.NewStyle().
            Width(methodSelectorInnerWidth).
            Height(1).
            AlignHorizontal(lipgloss.Center).
            Border(lipgloss.RoundedBorder()),

    }
}

func (selector MethodSelector) GetMethod() string {
    return methods[selector.selectedMethod]
}

func (selector *MethodSelector) NextMethod() {
    selector.selectedMethod++

    if selector.selectedMethod == len(methods) {
        selector.selectedMethod = 0
    }
}

func (selector *MethodSelector) PreviousMethod() {
    selector.selectedMethod--

    if selector.selectedMethod < 0 {
        selector.selectedMethod = len(methods) - 1
    }
}

func (selector MethodSelector) Render() string {
    return selector.style.Render(selector.GetMethod())
}
