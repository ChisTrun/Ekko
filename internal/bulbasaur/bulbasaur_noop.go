package bulbasaur

import "context"

type Noop struct{}

func (b *Noop) FindUserByNameRequest(ctx context.Context, content string) []uint64 {
	return []uint64{}
}
