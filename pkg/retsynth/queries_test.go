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
func sparseExist[T comparable](data []T, sparse []T) bool {

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
			want: []string{"rxn00001_c0", "rxn00002_c0", "rxn00003_c0", "rxn00004_c0", "rxn00005_c0", "rxn00006_c0", "rxn00007_c0", "rxn00008_c0", "rxn00009_c0", "rxn00010_c0", "rxn00011_c0", "rxn00012_c0", "rxn00013_c0", "rxn00014_c0", "rxn00015_c0", "rxn00016_c0", "rxn00017_c0", "rxn00018_c0", "rxn00019_c0", "rxn00020_c0", "rxn00021_c0", "rxn00022_c0", "rxn00023_c0", "rxn00024_c0", "rxn00025_c0", "rxn00026_c0", "rxn00027_c0", "rxn00028_c0", "rxn00029_c0", "rxn00030_c0", "rxn00031_c0", "rxn00032_c0", "rxn00033_c0", "rxn00034_c0", "rxn00035_c0", "rxn00036_c0", "rxn00037_c0", "rxn00038_c0", "rxn00039_c0", "rxn00040_c0", "rxn00041_c0", "rxn00042_c0", "rxn00043_c0", "rxn00044_c0", "rxn00045_c0", "rxn00046_c0", "rxn00047_c0", "rxn00048_c0", "rxn00049_c0", "rxn00050_c0", "rxn00051_c0", "rxn00052_c0", "rxn00053_c0", "rxn00054_c0", "rxn00055_c0", "rxn00056_c0", "rxn00057_c0", "rxn00058_c0", "rxn00059_c0", "rxn00060_c0", "rxn00061_c0", "rxn00062_c0"},
		},
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
			if got := GetReactionSpecies(tt.args.reactionID); !sparseExist(got, tt.want) {
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
				reactionID: "R03067_c0",
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
			got := GetReactantCompoundIDs(tt.args.reactionID)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
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
				compoundID: "R03067_c0",
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
			got := GetReactantCompoundIDs(tt.args.compoundID)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
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
				"C00001_c0",
				"C00001_c0",
				"C02715_c0",
				"C02715_c0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetProductCompundIDs(tt.args.reactionID)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
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
				modelID: "iJO1366",
			},
			want: []string{
				"C00001_c0",
				"C00002_c0",
				"C00003_c0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetModelCompounds(tt.args.modelID)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
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
			if got := GetAllCompoundIDs(); !sparseExist(got, tt.want) {
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
				{},
				{},
				{},
				{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllCompounds(); !sparseExist(got, tt.want) {
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
				"",
				"",
				"",
				"",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllCompoundInchistrings(); !sparseExist(got, tt.want) {
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
				modelID: "iJO1366",
			},
			want: []string{
				"R03067_c0",
				"R03068_c0",
				"R03069_c0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetModelReactions(tt.args.modelID)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
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
				"R03067_c0",
				"R03068_c0",
				"R03069_c0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllReactions(); !sparseExist(got, tt.want) {
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
				"b0002",
				"b0003",
				"b0004",
				"b0005",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetReactionGeneAssociations(tt.args.reactionID, tt.args.modelID)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
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
				"b0002",
				"b0003",
				"b0004",
				"b0005",
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
				reactionID: "R03067_c0",
				compoundID: "cpd00001_c0",
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
				reactionID: "R03067_c0",
			},
			want: []string{
				"b0002",
				"b0003",
				"b0004",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetReactionCatalysts(tt.args.reactionID)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
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
				compartmentName: "c0",
			},
			want: sql.NullString{String: "632.718.1", Valid: true},
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
				reactionID: "R03067_c0",
			},
			want: []string{
				"cpd00001_c0",
				"cpd00002_c0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetReactionSolvents(tt.args.reactionID)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
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
				reactionID: "R03067_c0",
			},
			want: sql.NullFloat64{Float64: 37, Valid: true},
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
				reactionID: "R03067_c0",
			},
			want: sql.NullFloat64{Float64: 1, Valid: true},
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
				reactionID: "R03067_c0",
			},
			want: sql.NullFloat64{Float64: 1, Valid: true},
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
				reactionID: "R03067_c0",
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
				reactionID: "R03067_c0",
			},
			want: sql.NullString{String: "10.1021/ja00003a001", Valid: true},
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
				reactionType: "synthesis",
			},
			want: []string{"R03067_c0", "R03067_c1", "R03067_c2", "R03067_c3", "R03067_c4", "R03067_c5", "R03067_c6", "R03067_c7", "R03067_c8", "R03067_c9", "R03067_c10", "R03067_c11", "R03067_c12", "R03067_c13", "R03067_c14", "R03067_c15", "R03067_c16", "R03067_c17", "R03067_c18", "R03067_c19", "R03067_c20", "R03067_c21", "R03067_c22", "R03067_c23", "R03067_c24", "R03067_c25", "R03067_c26", "R03067_c27", "R03067_c28", "R03067_c29", "R03067_c30", "R03067_c31", "R03067_c32", "R03067_c33", "R03067_c34", "R03067_c35", "R03067_c36", "R03067_c37", "R03067_c38", "R03067_c39", "R03067_c40", "R03067_c41", "R03067_c42", "R03067_c43", "R03067_c44", "R03067_c45", "R03067_c46", "R03067_c47", "R03067_c48", "R03067_c49", "R03067_c50", "R03067_c51", "R03067_c52", "R03067_c53", "R03067_c54", "R03067_c55", "R03067_c56", "R03067_c57", "R03067_c58", "R03067_c59", "R03067_c60", "R03067_c61", "R03067_c62", "R03067_c63", "R03067_c64", "R03067_c65", "R03067_c66", "R03067_c67", "R03067_c68", "R03067_c69", "R03067"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReactionsByType(tt.args.reactionType); !sparseExist(got, tt.want) {
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
			want: sql.NullString{String: "synthesis", Valid: true},
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
				"R00001_c0",
				"R00002_c0",
				"R00003_c0",
				"R00004_c0",
				"R00005_c0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllReactionKEGGIDs(); !sparseExist(got, tt.want) {
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
				String: "R00001",
				Valid:  true,
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
			if got := GetAllCompoundKEGGIDs(); !sparseExist(got, tt.want) {
				t.Errorf("GetAllCompoundKEGGIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllChemicalFormulas(t *testing.T) {
	setup()
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "Test GetAllChemicalFormulas",
			want: []string{
				"C00001_c0",
				"C00002_c0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllChemicalFormulas(); !sparseExist(got, tt.want) {
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
			want: sql.NullString{String: "C6H12O6", Valid: true},
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
			want: sql.NullString{String: "50-00-0", Valid: true},
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
				"C00031_c0",
				"C00031_c1",
				"C00031_c2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetCompoundIDByFormula(tt.args.formula)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
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
		want []string
	}{
		{
			name: "Test GetCompoundNameBySearchTerm",
			args: args{
				searchTerm: "glucose",
			},
			want: []string{
				"D-glucose",
				"D-glucose-6-phosphate",
				"D-glucose-6-phosphate 1-dehydrogenase",
				"D-glucose-6-phosphate aldolase",
				"D-glucose-6-phosphate aldolase (EC",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetCompoundNameBySearchTerm(tt.args.searchTerm)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
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
				fileName: "iJO1366.xml",
			},
			want: sql.NullString{String: "iJO1366", Valid: true},
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
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "Test GetAllFBAModelIDs",
			want: []string{
				"iJO1366",
				"iAF1260",
				"iAF692",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllFBAModelIDs(); !sparseExist(got, tt.want) {
				t.Errorf("GetAllFBAModelIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
