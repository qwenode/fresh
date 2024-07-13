package runner

import (
	"github.com/qwenode/gogo/ee"
	"github.com/qwenode/gogo/ff"
	"github.com/rs/zerolog/log"
	"path/filepath"
	"runtime"
	"time"
)

func root() string {
	return ee.GetString("watch_path")
}

func tmpPath() string {
	return ee.GetString("temp_path")
}

func buildName() string {
	return ee.GetString("build_name")
}
func buildPath() string {
	p := filepath.Join(ff.GetWorkDirectory(), tmpPath(), buildName())
	if runtime.GOOS == "windows" && filepath.Ext(p) != ".exe" {
		p += ".exe"
	}
	log.Info().Str("File", p).Msg("Build to")
	return p
}

func buildErrorsFileName() string {
	return ee.GetString("build_log")
}

func buildErrorsFilePath() string {
	return filepath.Join(tmpPath(), buildErrorsFileName())
}

func configPath() string {
	return ee.GetString("config_path")
}

func buildDelay() time.Duration {

	return time.Duration(ee.GetInt("build_delay"))
}
