package template

import (
	"fmt"
	"strings"
)

type GoModCfg struct {
	Module    string
	GoVersion string
}

func GoMod(cfg GoModCfg) (string, error) {
	cfg, err := fixGoVersion(cfg)
	if err != nil {
		return "", err
	}
	name := fmt.Sprintf("gomod:%v", cfg)
	return apply(name, gomodTemplate, cfg)
}

func fixGoVersion(cfg GoModCfg) (GoModCfg, error) {
	parsed := strings.Split(cfg.GoVersion, ".")
	if len(parsed) < 2 {
		return cfg, fmt.Errorf("invalid Go version[%s]", cfg.GoVersion)
	}
	if len(parsed) == 2 {
		return cfg, nil
	}
	cfg.GoVersion = parsed[0] + "." + parsed[1]
	return cfg, nil
}

const gomodTemplate = `module {{.Module}}

go {{.GoVersion}}`
