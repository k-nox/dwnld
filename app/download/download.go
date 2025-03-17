package download

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/k-nox/dwnld/app/config"
	"github.com/lrstanley/go-ytdlp"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const defaultOutputTemplate = "%(title)s [%(id)s].%(ext)s"

type Downloader struct {
	ctx         context.Context
	defaultOpts config.Download
}

func Startup(ctx context.Context, d *Downloader, opts config.Download) {
	d.ctx = ctx
	d.defaultOpts = opts
	d.ensureInstalled()
}

func (d *Downloader) ensureInstalled() {
	ytdlp.MustInstall(d.ctx, &ytdlp.InstallOptions{})
}

func (d *Downloader) Defaults() config.Download {
	return d.defaultOpts
}

func (d *Downloader) Download(url string, opts config.Download) error {
	if !opts.TargetResoltion.Valid() {
		return fmt.Errorf("requested resolution %s is not allowed currently", opts.TargetResoltion)
	}

	format := fmt.Sprintf("res:%s,vcodec:h264,acodec:m4a", opts.TargetResoltion)
	cmd := ytdlp.New().FormatSort(format)

	if opts.OutputTemplate == "" {
		opts.OutputTemplate = defaultOutputTemplate
	}

	cmd.Output(filepath.Join(opts.OutputDirectory, opts.OutputTemplate))

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
