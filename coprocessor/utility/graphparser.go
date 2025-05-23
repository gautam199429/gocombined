package utilities

import (
	"regexp"
	"strings"
	"time"

	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

type FieldsMap map[string]string

type EntitlementIdMap map[string]string

var parsedSchema *ast.Schema

func ParseSchema() (FieldsMap, EntitlementIdMap, error) {
	const cacheKey = "fieldsMap"
	if cached, found := GetFromCache(cacheKey); found {
		if fields, ok := cached.(FieldsMap); ok {
			return fields, nil, nil
		}
	}
	body, err := DownloadSchemaAsString()
	if err != nil {
		return nil, nil, err
	}
	doc, err := gqlparser.LoadSchema(&ast.Source{Input: body})
	if err != nil {
		return nil, nil, err
	}
	parsedSchema = doc
	entitlementIdMap, err := ExtractEntitlementIdentifiers(body)
	if err != nil {
		return nil, nil, err
	}
	allFieldMap := make(FieldsMap)
	for typeName, def := range doc.Types {
		if validateString(typeName) && len(def.Fields) > 0 {
			fieldMap := make(map[string]string)
			for _, field := range def.Fields {
				if validateString(field.Name) {
					fieldMap[field.Name] = field.Type.String()
					allFieldMap[field.Name] = field.Type.String()
				}
			}
		}
	}
	SetInCache(cacheKey, allFieldMap, 24*time.Hour)
	return allFieldMap, entitlementIdMap, nil
}

func ExtractEntitlementIdentifiers(schema string) (EntitlementIdMap, error) {
	result := make(EntitlementIdMap)
	regex := regexp.MustCompile(`key:\s*"(.*?)"(?:[^{}]|{[^{}]*})*?node:\s*{\s*entitlementIdentifier:\s*"(.*?)"`)
	matches := regex.FindAllStringSubmatch(schema, -1)
	for _, match := range matches {
		key := match[1]
		entitlementIdentifier := match[2]
		result[key] = entitlementIdentifier
	}
	return result, nil
}

func ResolveRefIdNameFallback(typeName string) string {
	if typeDef, ok := parsedSchema.Types[typeName]; ok {
		for _, dir := range typeDef.Directives {
			if dir.Name == "key" {
				for _, arg := range dir.Arguments {
					if arg.Name == "fields" {
						return arg.Value.Raw
					}
				}
			}
		}
	}
	return ""
}

func validateString(str string) bool {
	return !strings.HasPrefix(str, "__")
}
