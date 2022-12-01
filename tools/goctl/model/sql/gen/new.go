package gen

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"

	"github.com/zeromicro/go-zero/tools/goctl/model/sql/template"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func genNew(table Table, withCache, postgreSql bool) (string, error) {
	text, err := pathx.LoadTemplate(category, modelNewTemplateFile, template.New)
	if err != nil {
		return "", err
	}

	t := fmt.Sprintf(`"%s"`, wrapWithRawString(table.Name.Source(), postgreSql))
	//tName := fmt.Sprintf(`"%s"`, wrapWithRawString(table.Name.Source(), postgreSql))
	if postgreSql {
		t = "`" + fmt.Sprintf(`"%s"."%s"`, table.Db.Source(), table.Name.Source()) + "`"
		//tName = fmt.Sprintf(`"%s"."%s"`, table.Db.Source(), table.Name.Source())
	}
	rs := []rune(t)
	tName := string(rs[2 : len(t)-2])

	output, err := util.With("new").
		Parse(text).
		Execute(map[string]interface{}{
			"table":                 t,
			"tableName":             tName,
			"withCache":             withCache,
			"upperStartCamelObject": table.Name.ToCamel(),
			"lowerStartCamelObject": stringx.From(table.Name.ToCamel()).Untitle(),
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
