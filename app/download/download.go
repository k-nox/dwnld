package download

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/k-nox/dwnld/app/config"
	"github.com/lrstanley/go-ytdlp"
)

type Downloader struct {
	ctx context.Context
}

func (d *Downloader) Startup(ctx context.Context) {
	d.ctx = ctx
	d.ensureInstalled()
}

func (d *Downloader) ensureInstalled() {
	ytdlp.MustInstall(d.ctx, &ytdlp.InstallOptions{})
}

func (d *Downloader) Download(url string, opts config.Download) error {
	if !opts.TargetResoltion.Valid() {
		return fmt.Errorf("requested resolution %s is not allowed currently", opts.TargetResoltion)
	}

	format := fmt.Sprintf("res:%s,vcodec:h264,acodec:m4a", opts.TargetResoltion)
	cmd := ytdlp.New().FormatSort(format)

	cmd.Output(filepath.Join(opts.OutputDirectory, opts.OutputTemplate))

	if opts.EmbedSubtitles {
		cmd.EmbedSubs()
	}

	if opts.WriteInfoJSON {
		cmd.WriteInfoJSON()
	}

	_, err := cmd.Run(d.ctx, url)
	if err != nil {
		return fmt.Errorf("error running download cmd: %w", err)
	}

	return nil
}
