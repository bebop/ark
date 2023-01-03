package retsynth

// Models is a struct that contains the information about a model
type Model struct {
	ID       string `db:"ID"`        //ID is the model ID
	FileName string `db:"file_name"` //FileName is the name of the model file
}

// FBA Models is a struct that contains the information about a Flux Balanance Analysis model
type FBAModel struct {
	ID       string `db:"ID"` //ID is the FBA model ID
	FileName string `db:"file_name"`
}

// Compounds is a struct that contains the information about a compound
type Compound struct {
	ID              string `db:"ID"`              //ID is the compound ID
	Name            string `db:"name"`            //Name is the compound name
	Compartments    string `db:"compartments"`    //Compartments is the compartment of the compound
	KeggID          string `db:"kegg_id"`         //KeggID is the Kegg ID of the compound
	ChemicalFormula string `db:"chemicalformula"` //ChemicalFormula is the chemical formula of the compound
	InchiString     string `db:"inchistring"`     //InchiString is the Inchi string of the compound
	CASNumber       string `db:"casnumber"`       //CASNumber is the CAS number of the compound
}

// Compartments is a struct that contains the information about a compartment
type Compartment struct {
	ID   string `db:"ID"`   //ID is the compartment ID
	Name string `db:"name"` //Name is the compartment name
}

// ModelCompounds is a struct that contains the information about a  compound model
type ModelCompound struct {
	CompundID string `db:"cpd_ID"`   //CompundID is the compound ID
	ModelID   string `db:"model_ID"` //ModelID is the model ID
}

// Reactions is a struct that contains the information about a reaction
type Reaction struct {
	ID     string `db:"ID"`      //ID is the reaction ID
	Name   string `db:"name"`    //Name is the reaction name
	KeggID string `db:"kegg_id"` //KeggID is the Kegg ID of the reaction
	Type   string `db:"type"`    //Type is the type of the reaction
}

// ModelReactions is a struct that contains the information about a reaction model
type ModelReaction struct {
	ReactionID   string `db:"reaction_ID"` //ReactionID is the reaction ID
	ModelID      string `db:"model_ID"`    //ModelID is the model ID
	IsReversible bool   `db:"is_rev"`      //IsReversible is the reversibility of the reaction
}

// ReactionCompunds is a struct that contains the information about a reaction compound
type ReactionCompound struct {
	ReactionID   string `db:"reaction_ID"`  //ReactionID is the reaction ID
	CompoundID   string `db:"cpd_ID"`       //CompoundID is the compound ID
	IsProduct    bool   `db:"is_prod"`      //IsProduct is the product of the reaction
	Stochiometry int    `db:"stochiometry"` //Stochiometry is the stochiometry of the reaction
	FileNumer    int    `db:"filenum"`      //FileNumer is the file number of the reaction
}

// ReactionReversibility is a struct that contains the information about a reaction reversibility
type ReactionReversibility struct {
	ReactionID   string `db:"reaction_ID"`   //ReactionID is the reaction ID
	IsReversible bool   `db:"is_reversible"` //IsReversible is the reversibility of the reaction
}

// ReactionGene is a struct that contains the information about the reaction gene
type ReactionGene struct {
	ReactionID string `db:"reaction_ID"` //ReactionID is the reaction ID
	ModelID    string `db:"model_ID"`    //ModelID is the model ID
	GeneID     string `db:"gene_ID"`     //GeneID is the gene ID
}

// ReactionProtien is a struct that contains the information about the reaction protien
type ReactionProtien struct {
	ReactionID string `db:"reaction_ID"` //ReactionID is the reaction ID
	ModelID    string `db:"model_ID"`    //ModelID is the model ID
	ProtienID  string `db:"protien_ID"`  //ProtienID is the protien ID
}

// Cluster is a struct that contains the information about a cluster
type Cluster struct {
	ClusterNum string `db:"cluster_num"`
	ClusterID  string `db:"ID"`
}

// OringalDBCompooundIDs is a struct that contains the information about the original database compound IDs
type OringalDBCompooundIDs struct {
	CompoundID string `db:"ID"`       //CompoundID is the compound ID
	InchiID    string `db:"inchi_id"` //InchiID is the Inchi ID
}