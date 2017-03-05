package main

// DbUpdater updates the cache file.
type DbUpdater struct {
	h Hasher
	w Walker
}

func newDbUpdater(h Hasher, w Walker) *DbUpdater {
	return &DbUpdater{h: h, w: w}
}

// Update loads the cached values, removes entries for deleted files,
// adds digests of new files, and saves the database to disk.
func (u *DbUpdater) Update(db *Db, root string) {
	files := u.w.Walk(root)

	db.load()
	u.removeGone(db, files)
	u.addNew(db, files)
	db.save()
}

func (u *DbUpdater) removeGone(db *Db, files []string) {
	for filename := range db.list() {
		if !u.sliceContains(files, filename) {
			db.remove(filename)
		}
	}
}

func (u *DbUpdater) addNew(db *Db, files []string) {
	m := db.list()
	for _, filename := range files {
		if _, ok := m[filename]; !ok {
			hash := u.h.Hash(filename)
			db.add(filename, hash)
		}
	}
}

func (u *DbUpdater) sliceContains(slice []string, s string) bool {
	for _, value := range slice {
		if value == s {
			return true
		}
	}
	return false
}
