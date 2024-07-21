package requestTabs

import (
	"bytes"
	"encoding/json"
	"github.com/alecthomas/chroma/v2/quick"
	"mrproxy/shared"
	"net/http"
	"strings"
)

func renderRequest(method string, query string, headers http.Header, body []byte) string {
	doc := strings.Builder{}

	doc.WriteString(method)
	doc.WriteString(" ")
	doc.WriteString(query)
	doc.WriteString("\n\n")

	for k, values := range headers {
		for _, value := range values {
			doc.WriteString(shared.HighlightedText.Render(k))
			doc.WriteString(": ")
			doc.WriteString(value)
			doc.WriteString("\n")
		}
	}

	doc.WriteString("\n\n")

	if body != nil {
		str, err := renderPrettyJson(body)
		if err == nil {
			doc.WriteString(str)
		} else {
			doc.WriteString(string(body))
		}
	} else {
		doc.WriteString("No body...")
	}

	doc.WriteString("\n\n\n")
	return doc.String()
}

func renderPrettyJson(obj []byte) (string, error) {
	var err error
	var data interface{}
	err = json.Unmarshal(obj, &data)
	if err != nil {
		return "", err
	}

	var prettyJson []byte
	prettyJson, err = json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	var buff bytes.Buffer
	err = quick.Highlight(&buff, string(prettyJson), "json", "terminal256", "autumn")
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
