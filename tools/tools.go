// +build tools

package tools

import (
	_ "github.com/glerchundi/sqlboiler-crdb"
	_ "github.com/kyleconroy/sqlc/cmd/sqlc"
	_ "github.com/volatiletech/sqlboiler"
)

/*
  This file is to keep go modules from removing tools
  needed by go generate scripts. The approach described in
  https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
*/
