package app

import (
	"context"

	"github.com/lrstanley/go-ytdlp"
)

type Downloader struct {
	ctx context.Context
}

func (d *Downloader) ensureInstalled() {
	ytdlp.MustInstall(d.ctx, &ytdlp.InstallOptions{})
}

func (d *Downloader) Download(url string) error {
	_, err := ytdlp.New().FormatSort("res:1080").Run(d.ctx, url)
	if err != nil {
		return err
	}

	return nil
}
