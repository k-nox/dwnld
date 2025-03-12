package app

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/lrstanley/go-ytdlp"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const defaultOutputTemplate = "%(title)s [%(id)s].%(ext)s"

type Downloader struct {
	ctx context.Context
}

func (d *Downloader) ensureInstalled() {
	ytdlp.MustInstall(d.ctx, &ytdlp.InstallOptions{})
}

func (d *Downloader) Download(url string, outputDir string, outputTempl string) error {
	cmd := ytdlp.New().FormatSort("res,vcodec:h264,acodec:m4a")

	if outputTempl == "" {
		outputTempl = defaultOutputTemplate
	}

	cmd.Output(filepath.Join(outputDir, outputTempl))

	_, err := cmd.Run(d.ctx, url)
	if err != nil {
		return err
	}

	return nil
}

func (d *Downloader) ChooseDirectory() (string, error) {
	dir, err := runtime.OpenDirectoryDialog(d.ctx, runtime.OpenDialogOptions{
		Title:                "Choose directory to save downloads",
		CanCreateDirectories: true,
	})
	if err != nil {
		return "", fmt.Errorf("error choosing directory: %w", err)
	}
	return dir, nil
}
