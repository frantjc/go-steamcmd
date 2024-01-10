package steamcmd

import "context"

func Start(ctx context.Context) (Prompt, error) {
	return Command("steamcmd").Start(ctx)
}
