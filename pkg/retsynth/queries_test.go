package retsynth

import (
	"database/sql"
	"os"
	"reflect"
	"sort"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setup() {
	os.Setenv("RETSYNTH_DB_PATH", "../../data/dev/retsynth/minimal.db")
}

func TestConnectDB(t *testing.T) {
	setup()
	data := ConnectDB()
	if data == nil {
		t.Error("Error connecting to database")
	}
}

func TestGetUniqueMetabolicClusters(t *testing.T) {
	setup()
	tests := []struct {
		name string
		want int
	}{
		{
			name: "TestGetUniqueMetabolicClusters",
			want: 44,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUniqueMetabolicClusters(); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("GetUniqueMetabolicClusters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetModelIDsFromCluster(t *testing.T) {
	setup()
	type args struct {
		cluster string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetModelIDsFromCluster",
			args: args{
				cluster: "1",
			},
			want: []string{"632.718"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetModelIDsFromCluster(tt.args.cluster); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetModelIDsFromCluster() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllModelIDs(t *testing.T) {
	setup()
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "TestGetAllModelIDs",
			want: []string{
				"632.719",
				"632.718",
				"633.136",
				"633.137",
				"633.139",
				"633.141",
				"633.316",
				"633.408",
				"633417.11",
				"636.15",
				"636.38",
				"637971.3",
				"638.17",
				"639200.3",
				"639310.5",
				"64187.227",
				"64187.293",
				"64187.526",
				"64187.527",
				"64187.535",
				"64187.536",
				"64187.539",
				"64187.541",
				"64187.548",
				"64187.549",
				"64187.553",
				"64187.554",
				"64187.555",
				"64187.556",
				"64187.557",
				"64187.559",
				"64187.561",
				"64187.562",
				"64187.563",
				"64187.564",
				"64187.565",
				"64187.566",
				"64187.567",
				"64187.568",
				"64187.569",
				"64187.571",
				"64187.572",
				"64187.573",
				"64187.574",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAllModelIDs()
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllModelIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllModels(t *testing.T) {
	setup()
	model := []Model{
		{
			ID:       "632.719",
			FileName: "Yersinia_pestis_strain_M2085_Complete",
		},
	}
	print(model)
	tests := []struct {
		name string
		want []Model
	}{
		{
			name: "TestGetAllModels",
			want: []Model{
				{ID: "639310.5", FileName: "Olleya_aquimaris_strain_DAU311_Complete"},
				{ID: "632.719", FileName: "Yersinia_pestis_strain_M2085_Complete"},
				{ID: "64187.555", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_Ug11_Complete"},
				{ID: "633.141", FileName: "Yersinia_pseudotuberculosis_strain_FDAARGOS_580_strain_Not_applicable_Complete"},
				{ID: "64187.556", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_T19_Complete"},
				{ID: "636.38", FileName: "Edwardsiella_tarda_strain_Colony44_Complete"},
				{ID: "64187.527", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_PXO61_Complete"},
				{ID: "64187.563", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_CFBP1951_Complete"},
				{ID: "64187.553", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_CIX2374_Complete"},
				{ID: "64187.549", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_PXO142_Complete"},
				{ID: "633417.11", FileName: "Aeromonas_taiwanensis_strain_Colony382_Complete"},
				{ID: "632.718", FileName: "Yersinia_pestis_strain_M-1770_Complete"},
				{ID: "64187.557", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_Dak16_Complete"},
				{ID: "64187.293", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_MAI1_Complete"},
				{ID: "633.139", FileName: "Yersinia_pseudotuberculosis_strain_FDAARGOS_582_strain_Not_applicable_Complete"},
				{ID: "64187.561", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_CFBP7323_Complete"},
				{ID: "633.408", FileName: "Yersinia_pseudotuberculosis_598_Complete"},
				{ID: "64187.548", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_ICMP3125_Complete"},
				{ID: "64187.568", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_CFBP7320_Complete"},
				{ID: "639200.3", FileName: "Sphaerotilus_natans_subsp._sulfidivorans_strain_D-507_Complete"},
				{ID: "64187.541", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_XM9_Complete"},
				{ID: "64187.574", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_PXO404_Complete"},
				{ID: "64187.227", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_XF89b_Complete"},
				{ID: "633.316", FileName: "Yersinia_pseudotuberculosis_strain_598_Complete"},
				{ID: "64187.571", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_CFBP1949_Complete"},
				{ID: "64187.565", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_CFBP7337_Complete"},
				{ID: "64187.536", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_PX086_Complete"},
				{ID: "637971.3", FileName: "Lactobacillus_koreensis_strain_26-25_Complete"},
				{ID: "64187.573", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_PXO421_Complete"},
				{ID: "64187.539", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_IX-280_Complete"},
				{ID: "64187.569", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_CFBP7319_Complete"},
				{ID: "638.17", FileName: "Arsenophonus_nasoniae_strain_FIN_Complete"},
				{ID: "64187.526", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_PXO61_Complete"},
				{ID: "64187.564", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_CFBP1948_Complete"},
				{ID: "64187.562", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_CFBP7321_Complete"},
				{ID: "633.136", FileName: "Yersinia_pseudotuberculosis_strain_FDAARGOS_583_strain_Not_applicable_Complete"},
				{ID: "64187.567", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_CFBP7322_Complete"},
				{ID: "64187.554", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_CIX298_Complete"},
				{ID: "64187.572", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_PXO513_Complete"},
				{ID: "64187.566", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_CFBP7324_Complete"},
				{ID: "64187.535", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_PX079_Complete"},
				{ID: "633.137", FileName: "Yersinia_pseudotuberculosis_strain_FDAARGOS_579_strain_Not_applicable_Complete"},
				{ID: "64187.559", FileName: "Xanthomonas_oryzae_pv._oryzae_strain_CFBP7340_Complete"},
				{ID: "636.15", FileName: "Edwardsiella_tarda_strain_KC-Pc-HB1_Complete"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := GetAllModels()
			sort.Slice(got[:], func(i, j int) bool {
				return got[i].ID < got[j].ID
			})
			sort.Slice(tt.want[:], func(i, j int) bool {
				return tt.want[i].ID < tt.want[j].ID
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllModels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOrganismName(t *testing.T) {
	setup()
	type args struct {
		organismID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "TestGetOrganismName",
			args: args{
				organismID: "633417.11",
			},
			want: sql.NullString{String: "Aeromonas_taiwanensis_strain_Colony382_Complete", Valid: true},
		},
		{
			name: "TestGetOrganismName",
			args: args{
				organismID: "83333",
			},
			want: sql.NullString{String: "", Valid: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOrganismName(tt.args.organismID); got != tt.want {
				t.Errorf("GetOrganismName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOrganismID(t *testing.T) {
	setup()
	type args struct {
		organismName string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "TestGetOrganismID",
			args: args{
				organismName: "Escherichia coli",
			},
			want: sql.NullString{String: "", Valid: false},
		},
		{
			name: "TestGetOrganismID",
			args: args{
				organismName: "Aeromonas_taiwanensis_strain_Colony382_Complete",
			},
			want: sql.NullString{String: "633417.11", Valid: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOrganismID(tt.args.organismName); got != tt.want {
				t.Errorf("GetOrganismID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCompoundID(t *testing.T) {
	setup()
	type args struct {
		compoundName string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "TestGetCompoundID",
			args: args{
				compoundName: "ATP",
			},
			want: sql.NullString{String: "C00002_c0", Valid: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompoundID(tt.args.compoundName); got != tt.want {
				t.Errorf("GetCompoundID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLikeCompoundID(t *testing.T) {
	setup()
	type args struct {
		compoundName string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "TestGetLikeCompoundID",
			args: args{
				compoundName: "ATP",
			},
			want: sql.NullString{String: "C00002_c0", Valid: true},
		},
		{
			name: "TestGetLikeCompoundID",
			args: args{
				compoundName: "Krishna%",
			},
			want: sql.NullString{String: "", Valid: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLikeCompoundID(tt.args.compoundName); got != tt.want {
				t.Errorf("GetLikeCompoundID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCompoundIDFromInchi(t *testing.T) {
	setup()
	type args struct {
		inchi string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "TestGetCompoundIDFromInchi",
			args: args{
				inchi: "InChI=1S/C6H12O6/c7-1-3(9)5(11)6(12)4(10)2-8/h3,5,7-10H,1-2H2,(H,8,9)(H,11,12)/t3-,5+/m0/s1",
			},
			want: sql.NullString{String: "", Valid: false},
		},
		{
			name: "TestGetCompoundIDFromInchi",
			args: args{
				inchi: "InChI=1S/Mg/q+2",
			},
			want: sql.NullString{String: "C00305_e0", Valid: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompoundIDFromInchi(tt.args.inchi); got != tt.want {
				t.Errorf("GetCompoundIDFromInchi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCompoundInchi(t *testing.T) {
	setup()
	type args struct {
		compoundID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "TestGetCompoundInchi",
			args: args{
				"C00305_e0",
			},
			want: sql.NullString{String: "InChI=1S/Mg/q+2", Valid: true},
		},
		{
			name: "TestGetCompoundInchi",
			args: args{
				compoundID: "gibberish_id",
			},
			want: sql.NullString{String: "", Valid: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompoundInchi(tt.args.compoundID); got != tt.want {
				t.Errorf("GetCompoundInchi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCompoundName(t *testing.T) {
	type args struct {
		compoundID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "TestGetCompoundName",
			args: args{
				compoundID: "cpd00002",
			},
			want: sql.NullString{String: "ATP", Valid: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompoundName(tt.args.compoundID); got != tt.want {
				t.Errorf("GetCompoundName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCompoundNameFromInchi(t *testing.T) {
	type args struct {
		inchi string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "TestGetCompoundNameFromInchi",
			args: args{
				inchi: "InChI=1S/C6H12O6/c7-1-3(9)5(11)6(12)4(10)2-8/h3,5,7-10H,1-2H2,(H,8,9)(H,11,12)/t3-,5+/m0/s1",
			},
			want: sql.NullString{String: "Glucose", Valid: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompoundNameFromInchi(tt.args.inchi); got != tt.want {
				t.Errorf("GetCompoundNameFromInchi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCompoundCompartment(t *testing.T) {
	type args struct {
		compoundID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "TestGetCompoundCompartment",
			args: args{
				compoundID: "cpd00002",
			},
			want: sql.NullString{String: "c", Valid: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompoundCompartment(tt.args.compoundID); got != tt.want {
				t.Errorf("GetCompoundCompartment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionName(t *testing.T) {
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "TestGetReactionName",
			args: args{
				reactionID: "rxn00001",
			},
			want: sql.NullString{String: "ATP + H2O <=> ADP + phosphate", Valid: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionName(tt.args.reactionID); got != tt.want {
				t.Errorf("GetReactionName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionIDsFromCompound(t *testing.T) {
	type args struct {
		compoundID string
		isProduct  bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionIDsFromCompound(tt.args.compoundID, tt.args.isProduct); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReactionIDsFromCompound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionSpecies(t *testing.T) {
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionSpecies(tt.args.reactionID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReactionSpecies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactantCompoundIDs(t *testing.T) {
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactantCompoundIDs(tt.args.reactionID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReactantCompoundIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionsWithProduct(t *testing.T) {
	type args struct {
		compoundID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionsWithProduct(tt.args.compoundID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReactionsWithProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProductCompundIDs(t *testing.T) {
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetProductCompundIDs(tt.args.reactionID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductCompundIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetModelCompounds(t *testing.T) {
	type args struct {
		modelID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetModelCompounds(tt.args.modelID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetModelCompounds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllCompoundIDs(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllCompoundIDs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCompoundIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllCompounds(t *testing.T) {
	tests := []struct {
		name string
		want []Compound
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllCompounds(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCompounds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllCompoundInchistrings(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllCompoundInchistrings(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCompoundInchistrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetModelReactions(t *testing.T) {
	type args struct {
		modelID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetModelReactions(tt.args.modelID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetModelReactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllReactions(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllReactions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllReactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionReversibility(t *testing.T) {
	type args struct {
		reactionID string
		modelID    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionReversibility(tt.args.reactionID, tt.args.modelID); got != tt.want {
				t.Errorf("GetReactionReversibility() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionReversibilityGlobal(t *testing.T) {
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionReversibilityGlobal(tt.args.reactionID); got != tt.want {
				t.Errorf("GetReactionReversibilityGlobal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionGeneAssociations(t *testing.T) {
	type args struct {
		reactionID string
		modelID    string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionGeneAssociations(tt.args.reactionID, tt.args.modelID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReactionGeneAssociations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionProteinAssociations(t *testing.T) {
	type args struct {
		reactionID string
		modelID    string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionProteinAssociations(tt.args.reactionID, tt.args.modelID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReactionProteinAssociations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStoichiometry(t *testing.T) {
	type args struct {
		reactionID string
		compoundID string
		isProduct  bool
	}
	tests := []struct {
		name string
		args args
		want sql.NullFloat64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStoichiometry(tt.args.reactionID, tt.args.compoundID, tt.args.isProduct); got != tt.want {
				t.Errorf("GetStoichiometry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionCatalysts(t *testing.T) {
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionCatalysts(tt.args.reactionID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReactionCatalysts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCompartmentID(t *testing.T) {
	type args struct {
		compartmentName string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompartmentID(tt.args.compartmentName); got != tt.want {
				t.Errorf("GetCompartmentID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionSolvents(t *testing.T) {
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionSolvents(tt.args.reactionID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReactionSolvents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionTemperature(t *testing.T) {
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullFloat64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionTemperature(tt.args.reactionID); got != tt.want {
				t.Errorf("GetReactionTemperature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionPressure(t *testing.T) {
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullFloat64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionPressure(tt.args.reactionID); got != tt.want {
				t.Errorf("GetReactionPressure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionTime(t *testing.T) {
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullFloat64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionTime(tt.args.reactionID); got != tt.want {
				t.Errorf("GetReactionTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionYield(t *testing.T) {
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullFloat64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionYield(tt.args.reactionID); got != tt.want {
				t.Errorf("GetReactionYield() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionReference(t *testing.T) {
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionReference(tt.args.reactionID); got != tt.want {
				t.Errorf("GetReactionReference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionsByType(t *testing.T) {
	type args struct {
		reactionType string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionsByType(tt.args.reactionType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReactionsByType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionType(t *testing.T) {
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionType(tt.args.reactionID); got != tt.want {
				t.Errorf("GetReactionType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllReactionKEGGIDs(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllReactionKEGGIDs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllReactionKEGGIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionKEGGID(t *testing.T) {
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionKEGGID(tt.args.reactionID); got != tt.want {
				t.Errorf("GetReactionKEGGID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCompoundKEGGID(t *testing.T) {
	type args struct {
		compoundID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompoundKEGGID(tt.args.compoundID); got != tt.want {
				t.Errorf("GetCompoundKEGGID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllCompoundKEGGIDs(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllCompoundKEGGIDs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCompoundKEGGIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllChemicalFormulas(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllChemicalFormulas(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllChemicalFormulas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetChemicalFormula(t *testing.T) {
	type args struct {
		compoundID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetChemicalFormula(tt.args.compoundID); got != tt.want {
				t.Errorf("GetChemicalFormula() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCASNumber(t *testing.T) {
	type args struct {
		compoundID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCASNumber(tt.args.compoundID); got != tt.want {
				t.Errorf("GetCASNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCompoundIDByFormula(t *testing.T) {
	type args struct {
		formula string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompoundIDByFormula(tt.args.formula); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCompoundIDByFormula() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCompoundNameBySearchTerm(t *testing.T) {
	type args struct {
		searchTerm string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompoundNameBySearchTerm(tt.args.searchTerm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCompoundNameBySearchTerm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetModelIDByFileName(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetModelIDByFileName(tt.args.fileName); got != tt.want {
				t.Errorf("GetModelIDByFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllFBAModelIDs(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllFBAModelIDs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllFBAModelIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
