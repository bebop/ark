# allbase

## All your basepair are belong to us

We're building a modern solution to the biotech data access problem. Too many siloed databases using terrible dataformats that can't crossref eachother.

allbase is still in hot development but as it stands now the plan is to have a single service that can search and cross reference sequences and pathways from genbank, rhea, chembl, and uniprot. All sequences get hashes so they have a unique identifier. Allbase should in the end be able to spit out JSON for every sequence it has and every pathway it can make. Keoni should write down how somewhere but search is going to be dummy fast and easy to use through an API.

TODO:

### Next release
- Integrate litestream (https://litestream.io/)
- Incremental updating
- Hashing for genbank files
- Hashing for full uniprot entries
- contious integration and deployment
