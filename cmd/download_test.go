package cmd

import "testing"

func TestGetGenbank(t *testing.T) {
	getGenbank()
}
func TestGetFile(t *testing.T) {
	getFile("https://ftp.expasy.org/databases/rhea/rdf/rhea.rdf.gz")
	getFile("https://ftp.expasy.org/databases/rhea/tsv/rhea2uniprot_sprot.tsv")
	getFile("https://ftp.expasy.org/databases/rhea/tsv/rhea2uniprot_trembl.tsv.gz")
}
