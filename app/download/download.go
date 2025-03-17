package download

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

func Startup(ctx context.Context, d *Downloader) {
	d.ctx = ctx
	d.ensureInstalled()
}

func (d *Downloader) ensureInstalled() {
	ytdlp.MustInstall(d.ctx, &ytdlp.InstallOptions{})
}

func (d *Downloader) Download(opts Options) error {
	if !opts.TargetResoltion.valid() {
		return fmt.Errorf("requested resolution %s is not allowed currently", opts.TargetResoltion)
	}

	format := fmt.Sprintf("res:%s,vcodec:h264,acodec:m4a", opts.TargetResoltion)
	cmd := ytdlp.New().FormatSort(format)

	if opts.OutputTempl == "" {
		opts.OutputTempl = defaultOutputTemplate
	}

	cmd.Output(filepath.Join(opts.OutputDir, opts.OutputTempl))

	_, err := cmd.Run(d.ctx, opts.URL)
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
