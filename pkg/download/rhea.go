package download

//Rhea downloads all required files for allbase from the Rhea database.
func Rhea(writePath string) {
	go File("https://ftp.expasy.org/databases/rhea/rdf/rhea.rdf.gz", writePath)

	// get Rhea to curated uniprot mappings - relatively small.
	go File("https://ftp.expasy.org/databases/rhea/tsv/rhea2uniprot_sprot.tsv", writePath)

	// get Rhea to chaotic uniprot mappings - larger than sprot but still relatively small.
	go File("https://ftp.expasy.org/databases/rhea/tsv/rhea2uniprot_trembl.tsv.gz", writePath)
}
