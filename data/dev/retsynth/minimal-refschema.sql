--
-- File generated with SQLiteStudio v3.3.3 on Thu Jan 19 01:45:25 2023
--
-- Text encoding used: UTF-8
--
PRAGMA foreign_keys = off;
BEGIN TRANSACTION;

-- Table: cluster
CREATE TABLE cluster (
    cluster_num TEXT,
    ID          TEXT NOT NULL,
    PRIMARY KEY (
        ID
    )
);


-- Table: compartments
CREATE TABLE compartments (
    ID   TEXT NOT NULL,
    name TEXT,
    PRIMARY KEY (
        ID
    )
);


-- Table: compound
CREATE TABLE compound (
    ID              TEXT NOT NULL,
    name            TEXT,
    compartment     TEXT,
    kegg_id         TEXT,
    chemicalformula TEXT,
    casnumber       TEXT,
    inchistring     TEXT,
    PRIMARY KEY (
        ID
    )
);


-- Table: fba_models
CREATE TABLE fba_models (
    ID        TEXT,
    file_name TEXT,
    PRIMARY KEY (
        ID
    )
);


-- Table: model
CREATE TABLE model (
    ID        TEXT,
    file_name TEXT,
    PRIMARY KEY (
        ID
    )
);


-- Table: model_compound
CREATE TABLE model_compound (
    cpd_ID   TEXT,
    model_ID TEXT
);


-- Table: model_reaction
CREATE TABLE model_reaction (
    reaction_ID TEXT,
    model_ID    TEXT,
    is_rev      BIT (1) 
);


-- Table: original_db_cpdIDs
CREATE TABLE original_db_cpdIDs (
    ID       TEXT,
    inchi_id TEXT
);


-- Table: reaction
CREATE TABLE reaction (
    ID      TEXT NOT NULL,
    name    TEXT,
    kegg_id TEXT,
    type    TEXT,
    PRIMARY KEY (
        ID
    )
);


-- Table: reaction_catalysts
CREATE TABLE reaction_catalysts (
    reaction_ID  TEXT,
    catalysts_ID TEXT,
    name         TEXT
);


-- Table: reaction_compound
CREATE TABLE reaction_compound (
    reaction_ID   TEXT,
    cpd_ID        TEXT,
    is_prod       BIT (1),
    stoichiometry INT,
    filenum       INT
);


-- Table: reaction_gene
CREATE TABLE reaction_gene (
    reaction_ID TEXT,
    model_ID    TEXT,
    gene_ID     TEXT
);


-- Table: reaction_protein
CREATE TABLE reaction_protein (
    reaction_ID TEXT,
    model_ID    TEXT,
    protein_ID  TEXT
);


-- Table: reaction_reversibility
CREATE TABLE reaction_reversibility (
    reaction_ID   TEXT,
    is_reversible BIT (1),
    PRIMARY KEY (
        reaction_ID
    )
);


-- Table: reaction_solvents
CREATE TABLE reaction_solvents (
    reaction_ID TEXT,
    solvents_ID TEXT,
    name        TEXT
);


-- Table: reaction_spresi_info
CREATE TABLE reaction_spresi_info (
    reaction_ID TEXT PRIMARY KEY,
    temperature REAL,
    pressure    REAL,
    total_time  REAL,
    yield       REAL,
    reference   TEXT
);


-- Index: cluster_ind
CREATE INDEX cluster_ind ON cluster (
    cluster_num
);


-- Index: compound_ind
CREATE INDEX compound_ind ON compound (
    ID
);


-- Index: fba_models_ind
CREATE INDEX fba_models_ind ON fba_models (
    ID
);


-- Index: model_ind
CREATE INDEX model_ind ON model (
    ID
);


-- Index: modelcompound_ind1
CREATE INDEX modelcompound_ind1 ON model_compound (
    model_ID
);


-- Index: modelcompound_ind2
CREATE INDEX modelcompound_ind2 ON model_compound (
    cpd_ID
);


-- Index: modelreaction_ind1
CREATE INDEX modelreaction_ind1 ON model_reaction (
    model_ID
);


-- Index: modelreaction_ind2
CREATE INDEX modelreaction_ind2 ON model_reaction (
    reaction_ID
);


-- Index: original_db_cpdIDs_ind
CREATE INDEX original_db_cpdIDs_ind ON original_db_cpdIDs (
    ID,
    inchi_id
);


-- Index: reaction_ind
CREATE INDEX reaction_ind ON reaction (
    ID
);


-- Index: reaction_reversibility_ind
CREATE INDEX reaction_reversibility_ind ON reaction_reversibility (
    reaction_ID
);


-- Index: reactioncompound_ind1
CREATE INDEX reactioncompound_ind1 ON reaction_compound (
    reaction_ID,
    cpd_ID,
    is_prod
);


-- Index: reactioncompound_ind2
CREATE INDEX reactioncompound_ind2 ON reaction_compound (
    cpd_ID,
    is_prod
);


-- Index: reactiongene_ind
CREATE INDEX reactiongene_ind ON reaction_gene (
    reaction_ID,
    model_ID
);


-- Index: reactionprotein_ind
CREATE INDEX reactionprotein_ind ON reaction_protein (
    reaction_ID,
    model_ID
);


COMMIT TRANSACTION;
PRAGMA foreign_keys = on;
