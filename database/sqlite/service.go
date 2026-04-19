package sqlite

// New to set up the entity Cache service
func New(dbAbsPath string) Database {
	return &databaseImpl{
		dbAbsPath: dbAbsPath,
		queries: &Queries{
			CreateTable: []Query{},
			DropTable:   []Query{},
			Insert:      []Query{},
			Update:      []Query{},
			Delete:      []Query{},
			Select:      []Query{},
		},
	}
}
