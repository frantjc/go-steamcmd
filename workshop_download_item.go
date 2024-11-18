package steamcmd

import (
	"context"
	"fmt"
	"io"
)

type workshopDownloadItem struct {
	AppID           int
	PublishedFileID int
}

var _ cmd = new(workshopDownloadItem)

func (c *workshopDownloadItem) String() string {
	return fmt.Sprintf("%d/%d", c.AppID, c.PublishedFileID)
}

func (*workshopDownloadItem) Check(flags *promptFlags) error {
	if !flags.loggedIn {
		return fmt.Errorf("cannot workshop_download_item before login")
	}

	return nil
}

func (c *workshopDownloadItem) Args() ([]string, error) {
	if c == nil || c.AppID == 0 || c.PublishedFileID == 0 {
		return nil, fmt.Errorf("workshop_download_item requires app ID and published file ID")
	}

	args := []string{"workshop_download_item", fmt.Sprint(c.AppID), fmt.Sprint(c.PublishedFileID)}

	return args, nil
}

func (c *workshopDownloadItem) ReadOutput(ctx context.Context, r io.Reader) error {
	return readOutput(ctx, r, 0)
}

func (c *workshopDownloadItem) Modify(_ *promptFlags) error {
	return nil
}
