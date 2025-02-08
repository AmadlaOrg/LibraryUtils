package location

import "github.com/AmadlaOrg/LibraryUtils/env"

const (
	HeryStoragePath = env.HeryStoragePath
)

type AbsPaths struct {
	Storage    string // e.g.: /home/user/.hery/
	Catalog    string // e.g.: /home/user/.hery/collection/
	Collection string // e.g.: /home/user/.hery/collection/amadla/
	Entities   string // e.g.: /home/user/.hery/collection/amadla/entity/
	Cache      string // e.g.: /home/user/.hery/collection/amadla/amadla.cache
}
