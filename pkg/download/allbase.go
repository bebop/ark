package download

// Allbase literally downloads all the base data needed to build a standard allbase deployment
// the amount of data is dummy high to casually test on your personal machine. Run at your own risk.
func Allbase() {
	writePath := "../data/build"

	// Typically I'd write these functions to return errors but since I'm using go routines
	// the blocking nature of using channels to report errors would either make the
	// concurrency of go routines moot or make it so the returned errors were not returned until
	// all of the go routines were done which in this case kind of makes reporting errors a bit useless.

	// The solution here is that all of the functions called by the go routines will just log fatal errors.

	// I suppose it may be of some use to report when go routines are finished for the user's sake but that isn't a priority for
	// this pull request.

	// get Rhea - ~300MB total.
	go Rhea(writePath)

	// get CHEMBL Sqlite file - ~300MB compressed.
	go Chembl(writePath)

	// get curated sprot uniprot - ~1GB compressed.
	go File("https://ftp.uniprot.org/pub/databases/uniprot/current_release/knowledgebase/complete/uniprot_sprot.xml.gz", writePath)

	// get chaotic trembl uniprot - ~160GB compressed.
	go File("https://ftp.uniprot.org/pub/databases/uniprot/current_release/knowledgebase/complete/uniprot_trembl.xml.gz", writePath)

	// gets all of annotated genbank - Not sure how big it is as of writing this but it's a lot.
	go Genbank(writePath)
}
