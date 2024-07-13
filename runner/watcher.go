package runner

import (
	"github.com/qwenode/gogo/ee"
	"github.com/qwenode/gogo/ff"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"

	"github.com/howeyc/fsnotify"
)

func watchFolder(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fatal(err)
	}
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if isWatchedFile(ev.Name) {
					watcherLog("sending event %s", ev)
					startChannel <- ev.String()
				}
			case err := <-watcher.Error:
				watcherLog("error: %s", err)
			}
		}
	}()

	watcherLog("Watching %s", path)
	err = watcher.Watch(path)

	if err != nil {
		fatal(err)
	}
}

func watch() {
	rootPath := root()
	onlyRun := ee.GetBool("only_run")
	if onlyRun {
		log.Warn().Msg("ONLY RUN MODE")
		rootPath = filepath.Join(ff.GetWorkDirectory(), tmpPath())
	}
	filepath.Walk(
		rootPath, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() && (!isTmpDir(path) || onlyRun) {
				if len(path) > 1 && strings.HasPrefix(filepath.Base(path), ".") {
					return filepath.SkipDir
				}

				if isIgnoredFolder(path) && !onlyRun {
					watcherLog("Ignoring %s", path)
					return filepath.SkipDir
				}

				watchFolder(path)
			}

			return err
		},
	)
}
