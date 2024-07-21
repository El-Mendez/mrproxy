package requestTabs

import (
	"fmt"
	request2 "mrproxy/shared"
	"strings"
)

func renderRequest(request *request2.Request) string {
	doc := strings.Builder{}

	doc.WriteString(request.Method)
	doc.WriteString(" ")
	doc.WriteString(request.Query)
	doc.WriteString("\n\n")

	for _, header := range request.Headers {
		doc.WriteString(header.Key)
		doc.WriteString(": ")
		doc.WriteString(fmt.Sprintf("%v", header.Val))
		doc.WriteString("\n")
	}

	doc.WriteString("\n\n")
	renderJson(&doc, 0, request.Body, false)
	doc.WriteString("\n\n\n")
	return doc.String()
}

func renderJson(doc *strings.Builder, indentation int, value interface{}, initialIndentation bool) {
	if initialIndentation {
		doc.WriteString(strings.Repeat(" ", indentation*2))
	}

	if value == nil {
		doc.WriteString("null")
		return
	}

	switch v := value.(type) {
	case string:
		doc.WriteString(v)
	case float64:
		doc.WriteString(fmt.Sprintf("%f", v))
	case bool:
		doc.WriteString(fmt.Sprintf("%v", v))
	case []interface{}:
		doc.WriteString("[")
		for i := 0; i < len(v); i++ {
			renderJson(doc, indentation+1, v[i], true)
			doc.WriteString(",")
		}
		doc.WriteString(strings.Repeat(" ", indentation*2))
		doc.WriteString("]")
	case []request2.JsonField:
		doc.WriteString("{")
		for _, field := range v {
			doc.WriteString(strings.Repeat(" ", indentation*2+2))
			doc.WriteString(field.Key)
			doc.WriteString(": ")
			renderJson(doc, indentation+1, field.Val, false)
			doc.WriteString(",")
		}
		doc.WriteString(strings.Repeat(" ", indentation*2))
		doc.WriteString("}")
	}

}
