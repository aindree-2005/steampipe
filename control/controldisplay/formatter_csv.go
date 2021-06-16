package controldisplay

import (
	"bytes"
	"context"
	"encoding/csv"
	"io"
	"strings"

	"github.com/spf13/viper"
	"github.com/turbot/steampipe/constants"
	"github.com/turbot/steampipe/control/execute"
)

type CSVFormatter struct {
	csvWriter *csv.Writer
}

func (j *CSVFormatter) Format(_ context.Context, tree *execute.ExecutionTree) (io.Reader, error) {
	resultColumns := newResultColumns(tree)
	renderer := newGroupCsvRenderer()
	outBuffer := bytes.NewBufferString("")
	data := renderer.Render(tree)

	j.csvWriter = csv.NewWriter(outBuffer)
	j.csvWriter.Comma = []rune(viper.GetString(constants.ArgSeparator))[0]

	j.csvWriter.Write(resultColumns.AllColumns)
	j.csvWriter.WriteAll(data)
	return strings.NewReader(outBuffer.String()), nil
}
