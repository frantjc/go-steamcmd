package steamcmd

import (
	"context"
	"fmt"
	"io"
)

type WorkshopDownloadItem struct {
	AppID           int
	PublishedFileID int
}

var _ Cmd = new(WorkshopDownloadItem)

func (c *WorkshopDownloadItem) String() string {
	return fmt.Sprintf("%d/%d", c.AppID, c.PublishedFileID)
}

func (*WorkshopDownloadItem) check(flags *promptFlags) error {
	if !flags.loggedIn {
		return fmt.Errorf("cannot workshop_download_item before login")
	}

	return nil
}

func (c *WorkshopDownloadItem) args() ([]string, error) {
	if c == nil || c.AppID == 0 || c.PublishedFileID == 0 {
		return nil, fmt.Errorf("workshop_download_item requires app ID and published file ID")
	}

	args := []string{"workshop_download_item", fmt.Sprint(c.AppID), fmt.Sprint(c.PublishedFileID)}

	return args, nil
}

func (c *WorkshopDownloadItem) readOutput(ctx context.Context, r io.Reader) error {
	return base.readOutput(ctx, r)
}

func (c *WorkshopDownloadItem) modify(_ *promptFlags) error {
	return nil
}
