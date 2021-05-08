package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"k8s.io/helm/pkg/chartutil"
)

func getFunctions() template.FuncMap {
	f := sprig.TxtFuncMap()

	// from Helm!
	extra := template.FuncMap{
		"toToml":   chartutil.ToToml,
		"toYaml":   chartutil.ToYaml,
		"fromYaml": chartutil.FromYaml,
		"toJson":   chartutil.ToJson,
		"fromJson": chartutil.FromJson,
	}

	for k, v := range extra {
		f[k] = v
	}

	return f
}

func executeSingleTemplate(values map[string]interface{}, tplString string, isStrict bool) (string, error) {
	buf := bytes.NewBuffer(nil)
	tpl := template.New("mytmpl").Funcs(getFunctions())
	if isStrict {
		tpl.Option("missingkey=error")
	}

	tpl, err := tpl.Parse(tplString)
	if err != nil {
		return "", fmt.Errorf("Error parsing template(s): %v", err)
	}

	err = tpl.Execute(buf, values)
	if err != nil {
		return "", fmt.Errorf("Failed to parse standard input: %v", err)
	}

	// Work around to remove the "<no value>" go templates add.
	return strings.Replace(buf.String(), "<no value>", "", -1), nil
}

type TplRequest struct {
	Values   string
	Template string
}

func render(w http.ResponseWriter, req *http.Request) {
	var tplreq TplRequest
	err := json.NewDecoder(req.Body).Decode(&tplreq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonMap := make(map[string]interface{})
	errI := json.Unmarshal([]byte(tplreq.Values), &jsonMap)
	if err != nil {
		panic(errI)
	}

	fmt.Println(tplreq.Values)
	r, err := executeSingleTemplate(jsonMap, tplreq.Template, false)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, r)
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	fs := http.FileServer(http.Dir("./webui-vue/dist"))
	http.Handle("/", fs)

	http.HandleFunc("/render", render)
	addr, err := determineListenAddress()
	if err != nil {
		fmt.Println(err)
	}

	errListen := http.ListenAndServe(addr, nil)
	if errListen != nil {
		fmt.Println(errListen)
	}

}
