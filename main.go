package main

func main() {
	db := newDb("hashes.dat")

	u := newDbUpdater(SHA256{}, FilePathWalker{})
	u.Update(db, ".")

	df := &DupeFinder{}
	dupes := df.Find(db.list())

	p := &Printer{}
	p.PrintDupes(dupes)
}
