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

// Tests if the retrieved data has a random subset of the expected data.
// The sparse data is queried using 'ORDER BY RANDOM() LIMIT 10' in the SQL query.
func sparseEquals[T comparable](data []T, sparse []T) bool {

	for _, sparsevalue := range sparse {
		exist := false
		for _, datavalue := range data {
			if sparsevalue == datavalue {
				exist = true
			}
		}
		if !exist {
			return false
		}
	}
	return true
}

type Ordered interface {
	int | float64 | ~string
}

func unsortedEqual[T Ordered](data []T, expected []T) bool {

	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })

	return reflect.DeepEqual(data, expected)
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
			name: "TestGetCompoundName",
			args: args{
				compoundID: "C00305_e0",
			},
			want: sql.NullString{String: "Mg", Valid: true},
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
			name: "TestGetCompoundNameFromInchi",
			args: args{
				inchi: "InChI=1S/Mg/q+2",
			},
			want: sql.NullString{String: "Mg", Valid: true},
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
			name: "TestGetCompoundCompartment",
			args: args{
				compoundID: "C00001_c0",
			},
			want: sql.NullString{String: "c0", Valid: true},
		},
		{
			name: "TestGetCompoundCompartment",
			args: args{
				compoundID: "gibberish_id",
			},
			want: sql.NullString{String: "", Valid: false},
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
	setup()
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
				reactionID: "rxn10163_c0",
			},
			want: sql.NullString{String: "Tetradecanoate transport via proton symport", Valid: true},
		},
		{
			name: "TestGetReactionName",
			args: args{
				reactionID: "gibberish_id",
			},
			want: sql.NullString{String: "", Valid: false},
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
	setup()
	type args struct {
		compoundID string
		isProduct  bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetReactionIDsFromCompound",
			args: args{
				compoundID: "C00001_c0",
				isProduct:  false,
			},
			want: []string{
				"R08158_c0",
				"R00383_c0",
				"bio10_64187.565",
				"R01100_c0",
				"rxn08840_c0",
				"rxn08202_c0",
				"R01466_c0",
				"R05366_c0",
				"R08190_c0",
				"R02133_c0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionIDsFromCompound(tt.args.compoundID, tt.args.isProduct); !sparseEquals(got, tt.want) {
				t.Errorf("GetReactionIDsFromCompound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionSpecies(t *testing.T) {
	setup()
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetReactionSpecies",
			args: args{
				reactionID: "R03067_c0",
			},
			want: []string{
				"633.137",
				"64187.567",
				"632.719",
				"64187.527",
				"64187.564",
				"64187.564",
				"64187.539",
				"64187.571",
				"64187.563",
				"64187.559",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionSpecies(tt.args.reactionID); !sparseEquals(got, tt.want) {
				t.Errorf("GetReactionSpecies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactantCompoundIDs(t *testing.T) {
	setup()
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetReactantCompoundIDs",
			args: args{
				reactionID: "R00546_c0",
			},
			want: []string{
				"C00001_c0",
				"C00001_c0",
				"C02715_c0",
				"C02715_c0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactantCompoundIDs(tt.args.reactionID); !unsortedEqual(got, tt.want) {
				t.Errorf("GetReactantCompoundIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionsWithProduct(t *testing.T) {
	setup()
	type args struct {
		compoundID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetReactionsWithProduct",
			args: args{
				compoundID: "R00546_c0",
			},
			want: []string{
				"C00001_c0",
				"C00001_c0",
				"C02715_c0",
				"C02715_c0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactantCompoundIDs(tt.args.compoundID); !unsortedEqual(got, tt.want) {
				t.Errorf("GetReactantCompoundIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProductCompundIDs(t *testing.T) {
	setup()
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetProductCompundIDs",
			args: args{
				reactionID: "R00546_c0",
			},
			want: []string{
				"C00049_c0",
				"C00049_c0",
				"C00060_c0",
				"C00060_c0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetProductCompundIDs(tt.args.reactionID); !unsortedEqual(got, tt.want) {
				t.Errorf("GetProductCompundIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetModelCompounds(t *testing.T) {
	setup()
	type args struct {
		modelID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetModelCompounds",
			args: args{
				modelID: "632.718",
			},
			want: []string{
				"cpd15700_c0",
				"C05315_c0",
				"C00249_e0",
				"C00655_c0",
				"C06366_c0",
				"C03741_c0",
				"C01177_c0",
				"cpd15726_c0",
				"C15926_c0",
				"C00501_c0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetModelCompounds(tt.args.modelID); !sparseEquals(got, tt.want) {
				t.Errorf("GetModelCompounds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllCompoundIDs(t *testing.T) {
	setup()
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "TestGetAllCompoundIDs",
			want: []string{
				"C00001_c0",
				"C00002_c0",
				"C00003_c0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllCompoundIDs(); !sparseEquals(got, tt.want) {
				t.Errorf("GetAllCompoundIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllCompounds(t *testing.T) {
	setup()
	tests := []struct {
		name string
		want []Compound
	}{
		{
			name: "TestGetAllCompounds",
			want: []Compound{
				{
					ID:              "C05172_c0",
					Name:            sql.NullString{String: "Selenophosphate", Valid: true},
					Compartment:     sql.NullString{String: "c0", Valid: true},
					KeggID:          sql.NullString{String: "C05172", Valid: true},
					ChemicalFormula: sql.NullString{String: "H3O3PSe", Valid: true},
					CASNumber:       sql.NullString{String: "", Valid: false},
					InchiString:     sql.NullString{String: `InChI=1S/H3O3PSe/c1-4(2,3)5/h(H3,1,2,3,5)`, Valid: true},
				},
				{
					ID:              "C05448_c0",
					Name:            sql.NullString{String: "3alpha,7alpha,24-Trihydroxy-5beta-cholestanoyl-CoA", Valid: true},
					Compartment:     sql.NullString{String: "c0", Valid: true},
					KeggID:          sql.NullString{String: "C05448", Valid: true},
					ChemicalFormula: sql.NullString{String: "C48H80N7O20P3S", Valid: true},
					CASNumber:       sql.NullString{String: "", Valid: false},
					InchiString:     sql.NullString{String: `InChI=1S/C48H80N7O20P3S/c1-25(29-8-9-30-36-31(12-15-48(29,30)6)47(5)14-11-28(56)19-27(47)20-33(36)58)7-10-32(57)26(2)45(63)79-18-17-50-35(59)13-16-51-43(62)40(61)46(3,4)22-72-78(69,70)75-77(67,68)71-21-34-39(74-76(64,65)66)38(60)44(73-34)55-24-54-37-41(49)52-23-53-42(37)55/h23-34,36,38-40,44,56-58,60-61H,7-22H2,1-6H3,(H,50,59)(H,51,62)(H,67,68)(H,69,70)(H2,49,52,53)(H2,64,65,66)/t25-,26-,27+,28-,29-,30+,31+,32-,33-,34-,36+,38-,39-,40+,44-,47+,48-/m1/s1`, Valid: true},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllCompounds(); !sparseEquals(got, tt.want) {
				t.Errorf("GetAllCompounds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllCompoundInchistrings(t *testing.T) {
	setup()
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "TestGetAllCompoundInchistrings",
			want: []string{
				`InChI=1S/C41H68N7O17P3S/c1-4-5-6-7-8-9-10-11-12-13-14-15-16-17-18-19-20-21-32(50)69-25-24-43-31(49)22-23-44-39(53)36(52)41(2,3)27-62-68(59,60)65-67(57,58)61-26-30-35(64-66(54,55)56)34(51)40(63-30)48-29-47-33-37(42)45-28-46-38(33)48/h8-9,11-12,14-15,28-30,34-36,40,51-52H,4-7,10,13,16-27H2,1-3H3,(H,43,49)(H,44,53)(H,57,58)(H,59,60)(H2,42,45,46)(H2,54,55,56)/b9-8-,12-11-,15-14-/t30-,34-,35-,36+,40-/m1/s1`,
				`InChI=1S/C6H14O12P2/c7-1-2(8)6(18-20(14,15)16)4(10)3(9)5(1)17-19(11,12)13/h1-10H,(H2,11,12,13)(H2,14,15,16)/t1-,2-,3-,4+,5?,6?/m1/s1`,
				`InChI=1S/C6H12N3O4P/c7-5(3-13-14(10,11)12)1-6-2-8-4-9-6/h2,4-5H,1,3,7H2,(H,8,9)(H2,10,11,12)/t5-/m0/s1`,
				`InChI=1S/C3H4N2O4/c4-3(9)5-1(6)2(7)8/h(H,7,8)(H3,4,5,6,9)`,
				`InChI=1S/C9H8O6/c10-6(4-5-8(12)13)2-1-3-7(11)9(14)15/h1-5,11H,(H,12,13)(H,14,15)/b2-1-,5-4+,7-3+`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllCompoundInchiStrings(); !sparseEquals(got, tt.want) {
				t.Errorf("GetAllCompoundInchistrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetModelReactions(t *testing.T) {
	setup()
	type args struct {
		modelID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetModelReactions",
			args: args{
				modelID: "632.719",
			},
			want: []string{
				"R01875_c0",
				"R00956_c0",
				"R01714_c0",
				"R01059_c0",
				"R00704_c0",
				"R07281_c0",
				"R00253_c0",
				"R02971_c0",
				"R02941_c0",
				"R05234_c0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetModelReactions(tt.args.modelID); !sparseEquals(got, tt.want) {
				t.Errorf("GetModelReactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllReactions(t *testing.T) {
	setup()
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "TestGetAllReactions",
			want: []string{
				"R04857_c0",
				"rxn13782_c0",
				"rxn05516_c0",
				"rxn08846_c0",
				"R01271_c0",
				"rxn05200_c0",
				"R03018_c0",
				"R01876_c0",
				"R02285_c0",
				"R02376_c0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllReactions(); !sparseEquals(got, tt.want) {
				t.Errorf("GetAllReactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionReversibility(t *testing.T) {
	setup()
	type args struct {
		reactionID string
		modelID    string
	}
	tests := []struct {
		name string
		args args
		want sql.NullBool
	}{
		{
			name: "TestGetReactionReversibility",
			args: args{
				reactionID: "R03067_c0",
				modelID:    "632.718",
			},
			want: sql.NullBool{Bool: false, Valid: true},
		},
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
	setup()
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullBool
	}{
		{
			name: "TestGetReactionReversibilityGlobal",
			args: args{
				reactionID: "R03067_c0",
			},
			want: sql.NullBool{Bool: false, Valid: true},
		},
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
	setup()
	type args struct {
		reactionID string
		modelID    string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetReactionGeneAssociations",
			args: args{
				reactionID: "R03067_c0",
				modelID:    "632.718",
			},
			want: []string{
				"632.718.peg.1395",
				"632.718.peg.1396",
				"632.718.peg.3649",
				"632.718.peg.1396",
				"632.718.peg.3649",
				"632.718.peg.1395",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionGeneAssociations(tt.args.reactionID, tt.args.modelID); !unsortedEqual(got, tt.want) {
				t.Errorf("GetReactionGeneAssociations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionProteinAssociations(t *testing.T) {
	setup()
	type args struct {
		reactionID string
		modelID    string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetReactionProteinAssociations",
			args: args{
				reactionID: "R03067_c0",
				modelID:    "632.718",
			},
			want: []string{
				"(EC 2.5.1.15)",
				"(EC 2.5.1.19)",
				"(EC 2.5.1.15)",
				"(EC 2.5.1.19)",
			},
		},
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
	setup()
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
		{
			name: "TestGetStoichiometry",
			args: args{
				reactionID: "R00546_c0",
				compoundID: "C00001_c0",
				isProduct:  false,
			},
			want: sql.NullFloat64{Float64: 1, Valid: true},
		},
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
	setup()
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetReactionCatalysts",
			args: args{
				reactionID: "test_reaction_id",
			},
			want: []string{
				"test_catalyst_id",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionCatalysts(tt.args.reactionID); !sparseEquals(got, tt.want) {
				t.Errorf("GetReactionCatalysts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCompartmentID(t *testing.T) {
	setup()
	type args struct {
		compartmentName string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "TestGetCompartmentID",
			args: args{
				compartmentName: "Cytosol",
			},
			want: sql.NullString{String: "c0", Valid: true},
		},
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
	setup()
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetReactionSolvents",
			args: args{
				reactionID: "test_reaction_id",
			},
			want: []string{
				"test_solvent_id",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionSolvents(tt.args.reactionID); !unsortedEqual(got, tt.want) {
				t.Errorf("GetReactionSolvents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionTemperature(t *testing.T) {
	setup()
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullFloat64
	}{
		{
			name: "TestGetReactionTemperature",
			args: args{
				reactionID: "test_reaction0",
			},
			want: sql.NullFloat64{Float64: 25.0, Valid: true},
		},
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
	setup()
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullFloat64
	}{
		{
			name: "TestGetReactionPressure",
			args: args{
				reactionID: "test_reaction0",
			},
			want: sql.NullFloat64{Float64: 99.0, Valid: true},
		},
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
	setup()
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullFloat64
	}{
		{
			name: "TestGetReactionTime",
			args: args{
				reactionID: "test_reaction0",
			},
			want: sql.NullFloat64{Float64: 44.0, Valid: true},
		},
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
	setup()
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullFloat64
	}{
		{
			name: "TestGetReactionYield",
			args: args{
				reactionID: "test_reaction0",
			},
			want: sql.NullFloat64{Float64: 0.5, Valid: true},
		},
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
	setup()
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "TestGetReactionReference",
			args: args{
				reactionID: "test_reaction0",
			},
			want: sql.NullString{String: "test_reference", Valid: true},
		},
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
	setup()
	type args struct {
		reactionType string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetReactionsByType",
			args: args{
				reactionType: "bio",
			},
			want: []string{
				"R00214_c0",
				"rxn05456_c0",
				"R05599_c0",
				"R01103_c0",
				"R00946_c0",
				"R04425_c0",
				"rxn05667_c0",
				"R04091_c0",
				"rxn09978_c0",
				"bio10_636.15",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionsByType(tt.args.reactionType); !sparseEquals(got, tt.want) {
				t.Errorf("GetReactionsByType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionType(t *testing.T) {
	setup()
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "TestGetReactionType",
			args: args{
				reactionID: "R03067_c0",
			},
			want: sql.NullString{String: "bio", Valid: true},
		},
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
	setup()
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "TestGetAllReactionKEGGIDs",
			want: []string{
				"R05320",
				"R00925",
				"R01012",
				"R01558",
				"R08306",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllReactionKEGGIDs(); !sparseEquals(got, tt.want) {
				t.Errorf("GetAllReactionKEGGIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetReactionKEGGID(t *testing.T) {
	setup()
	type args struct {
		reactionID string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "Test GetReactionKEGGID",
			args: args{
				reactionID: "R00001_c0",
			},
			want: sql.NullString{
				String: "",
				Valid:  false,
			},
		},
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
			name: "Test GetCompoundKEGGID",
			args: args{
				compoundID: "C00001_c0",
			},
			want: sql.NullString{String: "C00001", Valid: true},
		},
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
	setup()
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "Test GetAllCompoundKEGGIDs",
			want: []string{
				"C05812",
				"C00811",
				"None",
				"C00668",
				"C01842",
				"None",
				"C00255",
				"None",
				"C00077",
				"C01096",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllCompoundKEGGIDs(); !sparseEquals(got, tt.want) {
				t.Errorf("GetAllCompoundKEGGIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllChemicalFormulas(t *testing.T) {

	setup()
	got := GetAllChemicalFormulas()
	print(got)
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "Test GetAllChemicalFormulas",
			want: []string{
				"C2HCl3",
				"C20H12O",
				"C37H66N7O17P3S",
				"C55H77CoN15O11",
				"C24H41N8O17P3S",
				"C10H17N3O6S",
				"C6H12O7",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllChemicalFormulas(); !sparseEquals(got, tt.want) {
				t.Errorf("GetAllChemicalFormulas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetChemicalFormula(t *testing.T) {
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
			name: "Test GetChemicalFormula",
			args: args{
				compoundID: "C00001_c0",
			},
			want: sql.NullString{String: "H2O", Valid: true},
		},
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
			name: "Test GetCASNumber",
			args: args{
				compoundID: "C00001_c0",
			},
			want: sql.NullString{String: "7732-18-5", Valid: true},
		},
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
	setup()
	type args struct {
		formula string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test GetCompoundIDByFormula",
			args: args{
				formula: "C6H12O6",
			},
			want: []string{
				"C00267_c0",
				"C00267_e0",
				"C00137_c0",
				"C00221_c0",
				"C00936_e0",
				"C01720_c0",
				"C02782_c0",
				"C10906_e0",
				"C00962_c0",
				"C00124_c0",
				"C00936_c0",
				"C00124_e0",
				"C10906_c0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompoundIDByFormula(tt.args.formula); !unsortedEqual(got, tt.want) {
				t.Errorf("GetCompoundIDByFormula() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCompoundNameBySearchTerm(t *testing.T) {
	setup()
	type args struct {
		searchTerm string
	}
	tests := []struct {
		name string
		args args
		want []Compound
	}{
		{
			name: "Test GetCompoundNameBySearchTerm",
			args: args{
				searchTerm: "glucose",
			},
			want: []Compound{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompoundBySearchTerm(tt.args.searchTerm); !sparseEquals(got, tt.want) {
				t.Errorf("GetCompoundNameBySearchTerm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetModelIDByFileName(t *testing.T) {
	setup()
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want sql.NullString
	}{
		{
			name: "Test GetModelIDByFileName",
			args: args{
				fileName: "Aeromonas_taiwanensis_strain_Colony382_Complete",
			},
			want: sql.NullString{String: "633417.11", Valid: true},
		},
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
	setup()
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "Test GetAllFBAModelIDs",
			want: []string{
				"64187.571",
				"64187.573",
				"632.718",
				"64187.553",
				"64187.536",
				"64187.564",
				"636.15",
				"639200.3",
				"64187.566",
				"64187.548",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllFBAModelIDs(); !sparseEquals(got, tt.want) {
				t.Errorf("GetAllFBAModelIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCompoundBySearchTerm(t *testing.T) {
	setup()
	type args struct {
		searchTerm string
	}
	tests := []struct {
		name string
		args args
		want []Compound
	}{
		{
			name: "Test GetCompoundBySearchTerm",
			args: args{
				searchTerm: "glucose",
			},
			want: []Compound{
				{
					ID:              "C00842_c0",
					Name:            sql.NullString{String: "dTDPglucose", Valid: true},
					Compartment:     sql.NullString{String: "c0", Valid: true},
					KeggID:          sql.NullString{String: "C00842", Valid: true},
					ChemicalFormula: sql.NullString{String: "C16H26N2O16P2", Valid: true},
					CASNumber:       sql.NullString{String: "2196-62-5", Valid: true},
					InchiString:     sql.NullString{String: `InChI=1S/C16H26N2O16P2/c1-6-3-18(16(25)17-14(6)24)10-2-7(20)9(31-10)5-30-35(26,27)34-36(28,29)33-15-13(23)12(22)11(21)8(4-19)32-15/h3,7-13,15,19-23H,2,4-5H2,1H3,(H,26,27)(H,28,29)(H,17,24,25)/t7-,8+,9+,10+,11+,12-,13+,15+/m0/s1`, Valid: true},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCompoundBySearchTerm(tt.args.searchTerm); !sparseEquals(got, tt.want) {
				t.Errorf("GetCompoundBySearchTerm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOrganismBySearchTerm(t *testing.T) {
	setup()
	type args struct {
		searchTerm string
	}
	tests := []struct {
		name string
		args args
		want []Model
	}{
		{
			name: "Test GetOrganismBySearchTerm",
			args: args{
				searchTerm: "aero",
			},
			want: []Model{
				{
					"633417.11",
					"Aeromonas_taiwanensis_strain_Colony382_Complete",
				},
				{
					"639200.3",
					"Sphaerotilus_natans_subsp._sulfidivorans_strain_D-507_Complete",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOrganismBySearchTerm(tt.args.searchTerm); !sparseEquals(got, tt.want) {
				t.Errorf("GetOrganismBySearchTerm() = %v, want %v", got, tt.want)
			}
		})
	}
}
