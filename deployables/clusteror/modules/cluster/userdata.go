package cluster

import (
	"bytes"
	"context"
	b64 "encoding/base64"
	"fmt"
	"path/filepath"
	"text/template"
)

var tmplPath, _ = filepath.Abs("../../deployables/clusteror/modules/cluster/template/user-data.tmpl")
var tmpl, tmplParseErr = template.New("user-data.tmpl").ParseFiles(tmplPath)

type CompileUserDataOptions struct {
	AssociatedIP string
	JoinToken    string
	PublicKey    string
}

func compileUserData(ctx context.Context, options *CompileUserDataOptions) (string, error) {
	if tmplParseErr != nil {
		panic(fmt.Errorf("failed to create userdata template:%s", tmplParseErr))
	}
	var userdata bytes.Buffer
	tmpl.Execute(&userdata, nil)
	encoded := b64.StdEncoding.EncodeToString(userdata.Bytes())
	return string(encoded), nil
}
