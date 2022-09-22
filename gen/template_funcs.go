package gen

import (
	"errors"
	"golang.org/x/net/html"
	"strings"
)

var FuncMap = map[string]interface{}{
	"toLower":      strings.ToLower,
	"toUpper":      strings.ToUpper,
	"toGetterName": ToGetterName,
	"toSetterName": ToSetterName,
	"toSelector":   ToSelector,
	"toClassName":  ToClassName,
	"escapeHtml":   EscapeHtml,
	"escapeQuote":  EscapeQuote,
	"toImport":     ToImport,
	"dict":         Dict,
}

func ToGetterName(name string) string {
	return "get" + strings.Title(name)
}
func ToSetterName(name string) string {
	return "Set" + strings.Title(name)
}
func ToSelector(name string) string {
	if name == "isEmail" {
		name = "isEmailAndGmail"
	}
	var first = string(name[2])
	if strings.ToLower(string(name[3])) == string(name[3]) {
		first = strings.ToLower(string(name[2]))
	}

	return first + name[3:] + "Validator"
}
func ToClassName(name string) string {
	if name == "isEmail" {
		name = "isEmailAndGmail"
	}
	return strings.ToUpper(string(name[2])) + name[3:] + "ValidatorDirective"
}
func ToImport(name string) string {
	return name[:len(name)-3]
}

func EscapeHtml(name string) string {
	return html.EscapeString(name)
}

func EscapeQuote(name string) string {
	return strings.Replace(name, "'", "\\'", 1)
}

func Dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}
