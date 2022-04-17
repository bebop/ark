# allbase

## All your basepair are belong to us

We're building a modern solution to the biotech data access problem. Too many siloed databases using terrible dataformats that can't crossref eachother. Allbase is still in hot development but as it stands now the plan is to have a single service that can search and cross reference sequences and pathways from genbank, rhea, chembl, and uniprot. All sequences get hashes so they have a unique identifier. Allbase should in the end be able to spit out JSON for every sequence it has and every pathway it can make. Keoni should write down how somewhere but search is going to be dummy fast and easy to use through an API.
## Specification
When completed allbase will have:

* [x] Functions that pull data from Genbank, Rhea, CHEMBL, and Uniprot.
  * [ ] Cron jobs for daily updates from public DBs mentioned above.
* [ ] Deploy as:
  * [ ] Single server
  * [ ] Cluster
* [ ] CI/CD
  * [ ] 97%+ code coverage
  * [ ] mock testing
  * [ ] production testing
  * [ ] deploy on push to main
  * [ ] continuous db updates
* [ ] Annotate given sequence string
* [ ] Improved data streaming. Currently allbase downloads THEN inserts.
* [ ] REST API endpoints to:
  * [ ] Query for metabolic pathways:
    * [x] Breadth first search.
    * [ ] Depth first search.
    * [ ] A* search.
  * [ ] Query for sequences across genbank and uniprot:
    * [ ] Super fast search using seqhash indentifiers.
    * [ ] Alignment (will be implemented in Poly).
      * [ ] BWA
      * [ ] minimap2
  * [ ] Insert user provided sequences.
