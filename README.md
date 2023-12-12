# ark

## ark is a database for engineering organisms

* **Comprehensive:** ark is a single service that can search, store, and cross reference sequences and pathways from genbank, rhea, chembl, and uniprot

* **Modern:** ark serves JSON and is written in Go. ark parses legacy file formats so you don't have to.

* **Stable (soon):** ark will be well tested and designed to be used in industrial, academic, and hobbyist settings

## Install

ark is still in hot development and not production ready. We currently only ship a pre-release dev branch for contributors.

 `git clone https://github.com/TimothyStiles/ark && cd ark && go test -v ./...`

## Community

* **[Discord](https://discord.gg/Hc8Ncwt):** Chat about ark and join us for game nights on our discord server!

## Contributing

* **[Code of conduct](CODE_OF_CONDUCT.md):** Please read the full text so you can understand what we're all about and remember to be excellent to each other!

* **[Contributor's guide](CONTRIBUTING.md):** Please read through it before you start hacking away and pushing contributions to this fine codebase.

## Sponsor

* **[Sponsor](https://github.com/sponsors/TimothyStiles):** 🤘 Thanks for your support 🤘

## License

* [MIT](LICENSE)

* Copyright (c) 2022 Timothy Stiles
<!-- We're building a modern solution to the biotech data access problem. Too many siloed databases using terrible dataformats that can't crossref eachother. ark is still in hot development but as it stands now the plan is to have a single service that can search and cross reference sequences and pathways from genbank, rhea, chembl, and uniprot. All sequences get hashes so they have a unique identifier. ark should in the end be able to spit out JSON for every sequence it has and every pathway it can make. Keoni should write down how somewhere but search is going to be dummy fast and easy to use through an API. -->
<!-- ## Specification
When completed ark will have:

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
* [ ] Improved data streaming. Currently ark downloads THEN inserts.
* [ ] REST API endpoints to:
  * [ ] Query for metabolic pathways:
    * [x] Breadth first search
    * [ ] Depth first search
    * [ ] A* search
  * [ ] Query for sequences across genbank and uniprot:
    * [ ] Super fast search using seqhash indentifiers.
    * [ ] Alignment (will be implemented in Poly).
      * [ ] BWA
      * [ ] minimap2
  * [ ] Insert user provided sequences. -->
