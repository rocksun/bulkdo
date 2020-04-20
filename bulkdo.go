package bulkdo

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
)

func readItems(in io.Reader) ([]map[string]string, error) {
	reader := csv.NewReader(in)
	rows := make([]map[string]string, 0)
	var header []string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if header == nil {
			header = record
		} else {
			dict := map[string]string{}
			for i := range header {
				dict[header[i]] = record[i]
			}
			rows = append(rows, dict)
		}

	}

	return rows, nil
}

func parseCommands(tplReader io.Reader, items []map[string]string) ([]string, error) {
	data, readErr := ioutil.ReadAll(tplReader)
	if readErr != nil {
		return nil, readErr
	}
	t, parseErr := template.New("Commands").Parse(string(data))
	if parseErr != nil {
		return nil, parseErr
	}

	commands := make([]string, 0)
	for _, item := range items {
		v := make(map[string]map[string]string, 0)
		v["v"] = item

		var b bytes.Buffer

		exeErr := t.Execute(&b, v)
		if exeErr != nil {
			return nil, exeErr
		}

		commands = append(commands, b.String())

	}

	return commands, nil
}

func execCommands(cmds []string) ([]string, error) {
	outs := make([]string, 0)
	for _, cmd := range cmds {
		file, err := ioutil.TempFile("", "bulkdo.*.bat")
		if err != nil {
			return nil, err
		}
		defer os.Remove(file.Name())

		file.WriteString(cmd)
		command := exec.Command("cmd.exe", "/c", file.Name())
		out, exeErr := command.CombinedOutput()
		if exeErr != nil {
			return nil, exeErr
		}
		outs = append(outs, string(out))

	}
	return outs, nil
}
