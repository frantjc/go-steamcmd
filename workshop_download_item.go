package steamcmd

import (
	"fmt"
)

type WorkshopDownloadItem struct {
	AppID           int
	PublishedFileID int
}

var _ Command = new(WorkshopDownloadItem)

func (c WorkshopDownloadItem) String() string {
	return fmt.Sprintf("%d/%d", c.AppID, c.PublishedFileID)
}

func (WorkshopDownloadItem) Check(flags *Flags) error {
	if !flags.LoggedIn {
		return fmt.Errorf("steamcmd: cannot workshop_download_item before login")
	}

	return nil
}

func (c WorkshopDownloadItem) Args() ([]string, error) {
	if c.AppID == 0 || c.PublishedFileID == 0 {
		return nil, fmt.Errorf("steamcmd: workshop_download_item requires app ID and published file ID")
	}

	args := []string{"workshop_download_item", fmt.Sprint(c.AppID), fmt.Sprint(c.PublishedFileID)}

	return args, nil
}

func (c WorkshopDownloadItem) Modify(_ *Flags) error {
	return nil
}
