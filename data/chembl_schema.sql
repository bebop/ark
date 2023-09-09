CREATE TABLE irac_classification (
	irac_class_id BIGINT NOT NULL, 
	active_ingredient VARCHAR(500) NOT NULL, 
	level1 VARCHAR(1) NOT NULL, 
	level1_description VARCHAR(2000) NOT NULL, 
	level2 VARCHAR(3) NOT NULL, 
	level2_description VARCHAR(2000) NOT NULL, 
	level3 VARCHAR(6) NOT NULL, 
	level3_description VARCHAR(2000) NOT NULL, 
	level4 VARCHAR(8) NOT NULL, 
	irac_code VARCHAR(3) NOT NULL, 
	CONSTRAINT irac_classification_pk PRIMARY KEY (irac_class_id), 
	CONSTRAINT uk_irac_class_l4 UNIQUE (level4)
);
CREATE TABLE chembl_id_lookup (
	chembl_id VARCHAR(20) NOT NULL, 
	entity_type VARCHAR(50) NOT NULL, 
	entity_id BIGINT NOT NULL, 
	status VARCHAR(10) NOT NULL, 
	last_active INTEGER, 
	CONSTRAINT chembl_id_lookup_pk PRIMARY KEY (chembl_id), 
	CONSTRAINT ck_chembl_id_lookup_status CHECK (status in ('ACTIVE','INACTIVE','OBS')), 
	CONSTRAINT chembl_id_lookup_uk UNIQUE (entity_type, entity_id)
);
CREATE TABLE atc_classification (
	who_name VARCHAR(2000), 
	level1 VARCHAR(10), 
	level2 VARCHAR(10), 
	level3 VARCHAR(10), 
	level4 VARCHAR(10), 
	level5 VARCHAR(10) NOT NULL, 
	level1_description VARCHAR(2000), 
	level2_description VARCHAR(2000), 
	level3_description VARCHAR(2000), 
	level4_description VARCHAR(2000), 
	CONSTRAINT pk_atc_code PRIMARY KEY (level5)
);
CREATE TABLE usan_stems (
	usan_stem_id BIGINT NOT NULL, 
	stem VARCHAR(100) NOT NULL, 
	subgroup VARCHAR(100) NOT NULL, 
	annotation VARCHAR(2000), 
	stem_class VARCHAR(100), 
	major_class VARCHAR(100), 
	who_extra SMALLINT, 
	CONSTRAINT pk_usan_stems PRIMARY KEY (usan_stem_id), 
	CONSTRAINT ck_usan_stems_class CHECK (stem_class in ('Suffix','Prefix','Infix')), 
	CONSTRAINT ck_usan_stems_mc CHECK (major_class in ('GPCR','NR','PDE','kinase','ion channel','protease')), 
	CONSTRAINT ck_usan_stems_who CHECK (who_extra in (0,1)), 
	CONSTRAINT uk_usan_stems_stemsub UNIQUE (stem, subgroup)
);
CREATE TABLE organism_class (
	oc_id BIGINT NOT NULL, 
	tax_id BIGINT, 
	l1 VARCHAR(200), 
	l2 VARCHAR(200), 
	l3 VARCHAR(200), 
	CONSTRAINT pk_orgclass_oc_id PRIMARY KEY (oc_id), 
	CONSTRAINT uk_orgclass_tax_id UNIQUE (tax_id)
);
CREATE TABLE data_validity_lookup (
	data_validity_comment VARCHAR(30) NOT NULL, 
	description VARCHAR(200), 
	CONSTRAINT sys_c00117330 PRIMARY KEY (data_validity_comment)
);
CREATE TABLE assay_type (
	assay_type VARCHAR(1) NOT NULL, 
	assay_desc VARCHAR(250), 
	CONSTRAINT pk_assaytype_assay_type PRIMARY KEY (assay_type)
);
CREATE TABLE confidence_score_lookup (
	confidence_score SMALLINT NOT NULL, 
	description VARCHAR(100) NOT NULL, 
	target_mapping VARCHAR(30) NOT NULL, 
	CONSTRAINT confidence_score_lookup_pk PRIMARY KEY (confidence_score)
);
CREATE TABLE curation_lookup (
	curated_by VARCHAR(32) NOT NULL, 
	description VARCHAR(100) NOT NULL, 
	CONSTRAINT pf_curlu_cur_by PRIMARY KEY (curated_by)
);
CREATE TABLE source (
	src_id INTEGER NOT NULL, 
	src_description VARCHAR(500), 
	src_short_name VARCHAR(20), 
	CONSTRAINT pk_source_src_id PRIMARY KEY (src_id)
);
CREATE TABLE relationship_type (
	relationship_type VARCHAR(1) NOT NULL, 
	relationship_desc VARCHAR(250), 
	CONSTRAINT pk_reltype_relationship_type PRIMARY KEY (relationship_type)
);
CREATE TABLE target_type (
	target_type VARCHAR(30) NOT NULL, 
	target_desc VARCHAR(250), 
	parent_type VARCHAR(25), 
	CONSTRAINT pk_targtype_target_type PRIMARY KEY (target_type)
);
CREATE TABLE variant_sequences (
	variant_id BIGINT NOT NULL, 
	mutation VARCHAR(2000), 
	accession VARCHAR(25), 
	version BIGINT, 
	isoform BIGINT, 
	sequence TEXT, 
	organism VARCHAR(200), 
	tax_id BIGINT, 
	CONSTRAINT pk_varseq_variant_id PRIMARY KEY (variant_id), 
	CONSTRAINT uk_varseq_mut_acc UNIQUE (mutation, accession)
);
CREATE TABLE bioassay_ontology (
	bao_id VARCHAR(11) NOT NULL, 
	label VARCHAR(100) NOT NULL, 
	CONSTRAINT bioassay_ontology_pk PRIMARY KEY (bao_id)
);
CREATE TABLE action_type (
	action_type VARCHAR(50) NOT NULL, 
	description VARCHAR(200) NOT NULL, 
	parent_type VARCHAR(50), 
	CONSTRAINT action_type_pk PRIMARY KEY (action_type)
);
CREATE TABLE frac_classification (
	frac_class_id BIGINT NOT NULL, 
	active_ingredient VARCHAR(500) NOT NULL, 
	level1 VARCHAR(2) NOT NULL, 
	level1_description VARCHAR(2000) NOT NULL, 
	level2 VARCHAR(2) NOT NULL, 
	level2_description VARCHAR(2000), 
	level3 VARCHAR(6) NOT NULL, 
	level3_description VARCHAR(2000), 
	level4 VARCHAR(7) NOT NULL, 
	level4_description VARCHAR(2000), 
	level5 VARCHAR(8) NOT NULL, 
	frac_code VARCHAR(4) NOT NULL, 
	CONSTRAINT frac_classification_pk PRIMARY KEY (frac_class_id), 
	CONSTRAINT uk_frac_class_l5 UNIQUE (level5)
);
CREATE TABLE activity_smid (
	smid BIGINT NOT NULL, 
	CONSTRAINT pk_actsamid PRIMARY KEY (smid)
);
CREATE TABLE component_sequences (
	component_id BIGINT NOT NULL, 
	component_type VARCHAR(50), 
	accession VARCHAR(25), 
	sequence TEXT, 
	sequence_md5sum VARCHAR(32), 
	description VARCHAR(200), 
	tax_id BIGINT, 
	organism VARCHAR(150), 
	db_source VARCHAR(25), 
	db_version VARCHAR(10), 
	CONSTRAINT pk_targcomp_seqs_compid PRIMARY KEY (component_id), 
	CONSTRAINT ck_targcomp_seqs_src CHECK (db_source in ('SWISS-PROT','TREMBL','Manual')), 
	CONSTRAINT ck_targcomp_seqs_taxid CHECK (tax_id > 0), 
	CONSTRAINT ck_targcomp_seqs_type CHECK (component_type in ('PROTEIN','DNA','RNA')), 
	CONSTRAINT uk_targcomp_seqs_acc UNIQUE (accession)
);
CREATE TABLE protein_classification (
	protein_class_id BIGINT NOT NULL, 
	parent_id BIGINT, 
	pref_name VARCHAR(500), 
	short_name VARCHAR(50), 
	protein_class_desc VARCHAR(410) NOT NULL, 
	definition VARCHAR(4000), 
	class_level BIGINT NOT NULL, 
	CONSTRAINT prot_class_pk PRIMARY KEY (protein_class_id), 
	CONSTRAINT ck_prot_class_level CHECK (class_level >= 0 and class_level <= 10)
);
CREATE TABLE bio_component_sequences (
	component_id BIGINT NOT NULL, 
	component_type VARCHAR(50) NOT NULL, 
	description VARCHAR(200), 
	sequence TEXT, 
	sequence_md5sum VARCHAR(32), 
	tax_id BIGINT, 
	organism VARCHAR(150), 
	CONSTRAINT pk_biocomp_seqs_compid PRIMARY KEY (component_id)
);
CREATE TABLE go_classification (
	go_id VARCHAR(10) NOT NULL, 
	parent_go_id VARCHAR(10), 
	pref_name VARCHAR(200), 
	class_level SMALLINT, 
	aspect VARCHAR(1), 
	path VARCHAR(1000), 
	CONSTRAINT go_classification_pk PRIMARY KEY (go_id)
);
CREATE TABLE assay_classification (
	assay_class_id BIGINT NOT NULL, 
	l1 VARCHAR(100), 
	l2 VARCHAR(100), 
	l3 VARCHAR(1000), 
	class_type VARCHAR(50), 
	source VARCHAR(50), 
	CONSTRAINT pk_assay_class PRIMARY KEY (assay_class_id), 
	CONSTRAINT uk_assay_class_l3 UNIQUE (l3)
);
CREATE TABLE structural_alert_sets (
	alert_set_id BIGINT NOT NULL, 
	set_name VARCHAR(100) NOT NULL, 
	priority SMALLINT NOT NULL, 
	CONSTRAINT pk_str_alert_set_id PRIMARY KEY (alert_set_id), 
	CONSTRAINT uk_str_alert_name UNIQUE (set_name)
);
CREATE TABLE products (
	dosage_form VARCHAR(200), 
	route VARCHAR(200), 
	trade_name VARCHAR(200), 
	approval_date DATETIME, 
	ad_type VARCHAR(5), 
	oral SMALLINT, 
	topical SMALLINT, 
	parenteral SMALLINT, 
	black_box_warning SMALLINT, 
	applicant_full_name VARCHAR(200), 
	innovator_company SMALLINT, 
	product_id VARCHAR(30) NOT NULL, 
	nda_type VARCHAR(10), 
	CONSTRAINT pk_products_id PRIMARY KEY (product_id), 
	CONSTRAINT ck_products_adtype CHECK (ad_type in ('OTC','RX','DISCN')), 
	CONSTRAINT ck_products_bbw CHECK (black_box_warning in (0,1)), 
	CONSTRAINT ck_products_inn CHECK (innovator_company in (0,1)), 
	CONSTRAINT ck_products_nda CHECK (NDA_TYPE in ('N','A')), 
	CONSTRAINT ck_products_oral CHECK (oral in (0,1)), 
	CONSTRAINT ck_products_par CHECK (parenteral in (0,1)), 
	CONSTRAINT ck_products_top CHECK (topical in (0,1))
);
CREATE TABLE patent_use_codes (
	patent_use_code VARCHAR(8) NOT NULL, 
	definition VARCHAR(500) NOT NULL, 
	CONSTRAINT patent_use_codes_pk PRIMARY KEY (patent_use_code), 
	CONSTRAINT ck_patent_use_code CHECK (patent_use_code like ('U-%'))
);
CREATE TABLE research_stem (
	res_stem_id BIGINT NOT NULL, 
	research_stem VARCHAR(20), 
	CONSTRAINT pk_res_stem_id PRIMARY KEY (res_stem_id), 
	CONSTRAINT uk_res_stem UNIQUE (research_stem)
);
CREATE TABLE hrac_classification (
	hrac_class_id BIGINT NOT NULL, 
	active_ingredient VARCHAR(500) NOT NULL, 
	level1 VARCHAR(2) NOT NULL, 
	level1_description VARCHAR(2000) NOT NULL, 
	level2 VARCHAR(3) NOT NULL, 
	level2_description VARCHAR(2000), 
	level3 VARCHAR(5) NOT NULL, 
	hrac_code VARCHAR(2) NOT NULL, 
	CONSTRAINT hrac_classification_pk PRIMARY KEY (hrac_class_id), 
	CONSTRAINT uk_hrac_class_l3 UNIQUE (level3)
);
CREATE TABLE protein_family_classification (
	protein_class_id BIGINT NOT NULL, 
	protein_class_desc VARCHAR(810) NOT NULL, 
	l1 VARCHAR(100) NOT NULL, 
	l2 VARCHAR(100), 
	l3 VARCHAR(100), 
	l4 VARCHAR(100), 
	l5 VARCHAR(100), 
	l6 VARCHAR(100), 
	l7 VARCHAR(100), 
	l8 VARCHAR(100), 
	CONSTRAINT protein_class_pk PRIMARY KEY (protein_class_id), 
	CONSTRAINT uk_protclass_desc UNIQUE (protein_class_desc), 
	CONSTRAINT uk_protclass_levels UNIQUE (l1, l2, l3, l4, l5, l6, l7, l8)
);
CREATE TABLE domains (
	domain_id BIGINT NOT NULL, 
	domain_type VARCHAR(20) NOT NULL, 
	source_domain_id VARCHAR(20) NOT NULL, 
	domain_name VARCHAR(20), 
	domain_description VARCHAR(500), 
	CONSTRAINT pk_domain_id PRIMARY KEY (domain_id), 
	CONSTRAINT ck_domain_type CHECK (domain_type in ('Pfam-A','Pfam-B'))
);
CREATE TABLE version (
	name VARCHAR(20) NOT NULL, 
	creation_date DATETIME, 
	comments VARCHAR(2000), 
	CONSTRAINT pk_version_name PRIMARY KEY (name)
);
CREATE TABLE activity_stds_lookup (
	std_act_id BIGINT NOT NULL, 
	standard_type VARCHAR(250) NOT NULL, 
	definition VARCHAR(500), 
	standard_units VARCHAR(100) NOT NULL, 
	normal_range_min NUMERIC(24, 12), 
	normal_range_max NUMERIC(24, 12), 
	CONSTRAINT pk_actstds_stdactid PRIMARY KEY (std_act_id), 
	CONSTRAINT uk_actstds_typeunits UNIQUE (standard_type, standard_units)
);
CREATE TABLE molecule_dictionary (
	molregno BIGINT NOT NULL, 
	pref_name VARCHAR(255), 
	chembl_id VARCHAR(20) NOT NULL, 
	max_phase SMALLINT NOT NULL, 
	therapeutic_flag SMALLINT NOT NULL, 
	dosed_ingredient SMALLINT NOT NULL, 
	structure_type VARCHAR(10) NOT NULL, 
	chebi_par_id BIGINT, 
	molecule_type VARCHAR(30), 
	first_approval INTEGER, 
	oral SMALLINT NOT NULL, 
	parenteral SMALLINT NOT NULL, 
	topical SMALLINT NOT NULL, 
	black_box_warning SMALLINT NOT NULL, 
	natural_product SMALLINT NOT NULL, 
	first_in_class SMALLINT NOT NULL, 
	chirality SMALLINT NOT NULL, 
	prodrug SMALLINT NOT NULL, 
	inorganic_flag SMALLINT NOT NULL, 
	usan_year INTEGER, 
	availability_type SMALLINT, 
	usan_stem VARCHAR(50), 
	polymer_flag SMALLINT, 
	usan_substem VARCHAR(50), 
	usan_stem_definition VARCHAR(1000), 
	indication_class VARCHAR(1000), 
	withdrawn_flag SMALLINT NOT NULL, 
	withdrawn_year INTEGER, 
	withdrawn_country VARCHAR(1000), 
	withdrawn_reason VARCHAR(1000), 
	withdrawn_class VARCHAR(500), 
	CONSTRAINT pk_moldict_molregno PRIMARY KEY (molregno), 
	CONSTRAINT fk_moldict_chembl_id FOREIGN KEY(chembl_id) REFERENCES chembl_id_lookup (chembl_id) ON DELETE CASCADE, 
	CONSTRAINT ck_moldict_app CHECK (first_approval < 2050 and first_approval > 1900), 
	CONSTRAINT ck_moldict_bbw CHECK (black_box_warning in (-1,0,1)), 
	CONSTRAINT ck_moldict_chi CHECK (chirality in (-1,0,1,2)), 
	CONSTRAINT ck_moldict_dosed CHECK (dosed_ingredient in (0,1)), 
	CONSTRAINT ck_moldict_fic CHECK (first_in_class in (-1,0,1)), 
	CONSTRAINT ck_moldict_inor CHECK (inorganic_flag in (-1,0,1)), 
	CONSTRAINT ck_moldict_np CHECK (natural_product in (-1,0,1)), 
	CONSTRAINT ck_moldict_oral CHECK (oral in (0,1)), 
	CONSTRAINT ck_moldict_par CHECK (parenteral in (0,1)), 
	CONSTRAINT ck_moldict_phase CHECK (max_phase in (0,1,2,3,4)), 
	CONSTRAINT ck_moldict_polyflag CHECK (polymer_flag IN (0, 1, null)), 
	CONSTRAINT ck_moldict_pro CHECK (prodrug in (-1,0,1)), 
	CONSTRAINT ck_moldict_strtype CHECK (structure_type in ('NONE','MOL','SEQ','BOTH')), 
	CONSTRAINT ck_moldict_theraflag CHECK (therapeutic_flag IN (0, 1)), 
	CONSTRAINT ck_moldict_top CHECK (topical in (0,1)), 
	CONSTRAINT ck_moldict_usanyear CHECK (usan_year > 1900 and usan_year < 2050), 
	CONSTRAINT ck_moldict_withd CHECK (WITHDRAWN_FLAG in (0,1)), 
	CONSTRAINT uk_moldict_chemblid UNIQUE (chembl_id)
);
CREATE TABLE defined_daily_dose (
	atc_code VARCHAR(10) NOT NULL, 
	ddd_units VARCHAR(200), 
	ddd_admr VARCHAR(1000), 
	ddd_comment VARCHAR(2000), 
	ddd_id BIGINT NOT NULL, 
	ddd_value NUMERIC, 
	CONSTRAINT pk_ddd_id PRIMARY KEY (ddd_id), 
	CONSTRAINT fk_ddd_atccode FOREIGN KEY(atc_code) REFERENCES atc_classification (level5) ON DELETE CASCADE
);
CREATE TABLE cell_dictionary (
	cell_id BIGINT NOT NULL, 
	cell_name VARCHAR(50) NOT NULL, 
	cell_description VARCHAR(200), 
	cell_source_tissue VARCHAR(50), 
	cell_source_organism VARCHAR(150), 
	cell_source_tax_id BIGINT, 
	clo_id VARCHAR(11), 
	efo_id VARCHAR(12), 
	cellosaurus_id VARCHAR(15), 
	cl_lincs_id VARCHAR(8), 
	chembl_id VARCHAR(20), 
	cell_ontology_id VARCHAR(10), 
	CONSTRAINT pk_celldict_cellid PRIMARY KEY (cell_id), 
	CONSTRAINT fk_celldict_chembl_id FOREIGN KEY(chembl_id) REFERENCES chembl_id_lookup (chembl_id) ON DELETE CASCADE, 
	CONSTRAINT ck_cell_dict_lincs CHECK (CL_LINCS_ID like ('LCL-%')), 
	CONSTRAINT uk_celldict UNIQUE (cell_name, cell_source_tax_id), 
	CONSTRAINT uk_cell_chembl_id UNIQUE (chembl_id)
);
CREATE TABLE docs (
	doc_id BIGINT NOT NULL, 
	journal VARCHAR(50), 
	year INTEGER, 
	volume VARCHAR(50), 
	issue VARCHAR(50), 
	first_page VARCHAR(50), 
	last_page VARCHAR(50), 
	pubmed_id BIGINT, 
	doi VARCHAR(100), 
	chembl_id VARCHAR(20) NOT NULL, 
	title VARCHAR(500), 
	doc_type VARCHAR(50) NOT NULL, 
	authors VARCHAR(4000), 
	abstract TEXT, 
	patent_id VARCHAR(20), 
	ridx VARCHAR(200) NOT NULL, 
	src_id INTEGER NOT NULL, 
	CONSTRAINT pk_docs_doc_id PRIMARY KEY (doc_id), 
	CONSTRAINT fk_docs_chembl_id FOREIGN KEY(chembl_id) REFERENCES chembl_id_lookup (chembl_id) ON DELETE CASCADE, 
	CONSTRAINT fk_docs_src_id FOREIGN KEY(src_id) REFERENCES source (src_id) ON DELETE CASCADE, 
	CONSTRAINT ck_docs_chemblid CHECK (chembl_id like ('CHEMBL%')), 
	CONSTRAINT ck_docs_doctype CHECK (doc_type in ('PUBLICATION','BOOK','DATASET','PATENT')), 
	CONSTRAINT ck_docs_year CHECK (year < 2050 and year > 1900), 
	CONSTRAINT uk_docs_chemblid UNIQUE (chembl_id)
);
CREATE TABLE target_dictionary (
	tid BIGINT NOT NULL, 
	target_type VARCHAR(30), 
	pref_name VARCHAR(200) NOT NULL, 
	tax_id BIGINT, 
	organism VARCHAR(150), 
	chembl_id VARCHAR(20) NOT NULL, 
	species_group_flag SMALLINT NOT NULL, 
	CONSTRAINT pk_targdict_tid PRIMARY KEY (tid), 
	CONSTRAINT fk_targdict_chembl_id FOREIGN KEY(chembl_id) REFERENCES chembl_id_lookup (chembl_id) ON DELETE CASCADE, 
	CONSTRAINT fk_targdict_target_type FOREIGN KEY(target_type) REFERENCES target_type (target_type) ON DELETE CASCADE, 
	CONSTRAINT ck_targdict_species CHECK (species_group_flag in (0,1)), 
	CONSTRAINT uk_targdict_chemblid UNIQUE (chembl_id)
);
CREATE TABLE tissue_dictionary (
	tissue_id BIGINT NOT NULL, 
	uberon_id VARCHAR(15), 
	pref_name VARCHAR(200) NOT NULL, 
	efo_id VARCHAR(20), 
	chembl_id VARCHAR(20) NOT NULL, 
	bto_id VARCHAR(20), 
	caloha_id VARCHAR(7), 
	CONSTRAINT pk_tissue_dict_tissue_id PRIMARY KEY (tissue_id), 
	CONSTRAINT fk_tissue_chembl_id FOREIGN KEY(chembl_id) REFERENCES chembl_id_lookup (chembl_id) ON DELETE CASCADE, 
	CONSTRAINT ck_tissue_uberon_id CHECK (uberon_id like ('UBERON:%')), 
	CONSTRAINT uk_tissue_chembl_id UNIQUE (chembl_id), 
	CONSTRAINT uk_tissue_dict_uberon_efo UNIQUE (uberon_id, efo_id), 
	CONSTRAINT uk_tissue_pref_name UNIQUE (pref_name)
);
CREATE TABLE activity_supp (
	as_id BIGINT NOT NULL, 
	rgid BIGINT NOT NULL, 
	smid BIGINT, 
	type VARCHAR(250) NOT NULL, 
	relation VARCHAR(50), 
	value NUMERIC, 
	units VARCHAR(100), 
	text_value VARCHAR(1000), 
	standard_type VARCHAR(250), 
	standard_relation VARCHAR(50), 
	standard_value NUMERIC, 
	standard_units VARCHAR(100), 
	standard_text_value VARCHAR(1000), 
	comments VARCHAR(4000), 
	CONSTRAINT pk_actsupp_as_id PRIMARY KEY (as_id), 
	CONSTRAINT fk_act_smids FOREIGN KEY(smid) REFERENCES activity_smid (smid) ON DELETE CASCADE, 
	CONSTRAINT uk_actsupp_rgid_type UNIQUE (rgid, type)
);
CREATE TABLE component_class (
	component_id BIGINT NOT NULL, 
	protein_class_id BIGINT NOT NULL, 
	comp_class_id BIGINT NOT NULL, 
	CONSTRAINT pk_comp_class_id PRIMARY KEY (comp_class_id), 
	CONSTRAINT fk_comp_class_compid FOREIGN KEY(component_id) REFERENCES component_sequences (component_id) ON DELETE CASCADE, 
	CONSTRAINT fk_comp_class_pcid FOREIGN KEY(protein_class_id) REFERENCES protein_classification (protein_class_id) ON DELETE CASCADE, 
	CONSTRAINT uk_comp_class UNIQUE (component_id, protein_class_id)
);
CREATE TABLE protein_class_synonyms (
	protclasssyn_id BIGINT NOT NULL, 
	protein_class_id BIGINT NOT NULL, 
	protein_class_synonym VARCHAR(1000), 
	syn_type VARCHAR(20), 
	CONSTRAINT pk_protclasssyn_synid PRIMARY KEY (protclasssyn_id), 
	CONSTRAINT fk_protclasssyn_protclass_id FOREIGN KEY(protein_class_id) REFERENCES protein_classification (protein_class_id) ON DELETE CASCADE, 
	CONSTRAINT ck_protclasssyn_syntype CHECK (syn_type in ('CHEMBL','CONCEPT_WIKI','UMLS','CW_XREF','MESH_XREF')), 
	CONSTRAINT uk_protclasssyn UNIQUE (protein_class_id, protein_class_synonym, syn_type)
);
CREATE TABLE structural_alerts (
	alert_id BIGINT NOT NULL, 
	alert_set_id BIGINT NOT NULL, 
	alert_name VARCHAR(100) NOT NULL, 
	smarts VARCHAR(4000) NOT NULL, 
	CONSTRAINT pk_str_alert_id PRIMARY KEY (alert_id), 
	CONSTRAINT fk_str_alert_set_id FOREIGN KEY(alert_set_id) REFERENCES structural_alert_sets (alert_set_id) ON DELETE CASCADE, 
	CONSTRAINT uk_str_alert_smarts UNIQUE (alert_set_id, alert_name, smarts)
);
CREATE TABLE product_patents (
	prod_pat_id BIGINT NOT NULL, 
	product_id VARCHAR(30) NOT NULL, 
	patent_no VARCHAR(20) NOT NULL, 
	patent_expire_date DATETIME NOT NULL, 
	drug_substance_flag SMALLINT NOT NULL, 
	drug_product_flag SMALLINT NOT NULL, 
	patent_use_code VARCHAR(10), 
	delist_flag SMALLINT NOT NULL, 
	submission_date DATETIME, 
	CONSTRAINT pk_prod_pat_id PRIMARY KEY (prod_pat_id), 
	CONSTRAINT fk_prod_pat_product_id FOREIGN KEY(product_id) REFERENCES products (product_id) ON DELETE CASCADE, 
	CONSTRAINT fk_prod_pat_use_code FOREIGN KEY(patent_use_code) REFERENCES patent_use_codes (patent_use_code) ON DELETE CASCADE, 
	CONSTRAINT ck_patents_delistflag CHECK (delist_flag IN (0, 1)), 
	CONSTRAINT ck_patents_prodflag CHECK (drug_product_flag IN (0, 1)), 
	CONSTRAINT ck_patents_subsflag CHECK (drug_substance_flag IN (0, 1)), 
	CONSTRAINT uk_prod_pat UNIQUE (product_id, patent_no, patent_expire_date, patent_use_code)
);
CREATE TABLE component_synonyms (
	compsyn_id BIGINT NOT NULL, 
	component_id BIGINT NOT NULL, 
	component_synonym VARCHAR(500), 
	syn_type VARCHAR(20), 
	CONSTRAINT pk_compsyn_synid PRIMARY KEY (compsyn_id), 
	CONSTRAINT fk_compsyn_compid FOREIGN KEY(component_id) REFERENCES component_sequences (component_id) ON DELETE CASCADE, 
	CONSTRAINT ck_compsyn_syntype CHECK (syn_type in ('GENE_SYMBOL','GENE_SYMBOL_OTHER','UNIPROT','MANUAL','OTHER','EC_NUMBER')), 
	CONSTRAINT uk_compsyn UNIQUE (component_id, component_synonym, syn_type)
);
CREATE TABLE research_companies (
	co_stem_id BIGINT NOT NULL, 
	res_stem_id BIGINT, 
	company VARCHAR(100), 
	country VARCHAR(50), 
	previous_company VARCHAR(100), 
	CONSTRAINT pk_resco_co_stem_id PRIMARY KEY (co_stem_id), 
	CONSTRAINT fk_resco_res_stem_id FOREIGN KEY(res_stem_id) REFERENCES research_stem (res_stem_id) ON DELETE CASCADE, 
	CONSTRAINT uk_resco_stem_co UNIQUE (res_stem_id, company)
);
CREATE TABLE component_domains (
	compd_id BIGINT NOT NULL, 
	domain_id BIGINT, 
	component_id BIGINT NOT NULL, 
	start_position BIGINT, 
	end_position BIGINT, 
	CONSTRAINT pk_compd_id PRIMARY KEY (compd_id), 
	CONSTRAINT fk_compd_compid FOREIGN KEY(component_id) REFERENCES component_sequences (component_id) ON DELETE CASCADE, 
	CONSTRAINT fk_compd_domainid FOREIGN KEY(domain_id) REFERENCES domains (domain_id) ON DELETE CASCADE, 
	CONSTRAINT ck_compd_end CHECK (end_position > 0), 
	CONSTRAINT ck_compd_start CHECK (start_position > 0), 
	CONSTRAINT uk_compd_start UNIQUE (domain_id, component_id, start_position)
);
CREATE TABLE component_go (
	comp_go_id BIGINT NOT NULL, 
	component_id BIGINT NOT NULL, 
	go_id VARCHAR(10) NOT NULL, 
	CONSTRAINT pk_comp_go PRIMARY KEY (comp_go_id), 
	CONSTRAINT fk_comp_id FOREIGN KEY(component_id) REFERENCES component_sequences (component_id) ON DELETE CASCADE, 
	CONSTRAINT fk_go_id FOREIGN KEY(go_id) REFERENCES go_classification (go_id) ON DELETE CASCADE, 
	CONSTRAINT uk_comp_go UNIQUE (component_id, go_id)
);
CREATE TABLE molecule_irac_classification (
	mol_irac_id BIGINT NOT NULL, 
	irac_class_id BIGINT NOT NULL, 
	molregno BIGINT NOT NULL, 
	CONSTRAINT molecule_irac_classificationpk PRIMARY KEY (mol_irac_id), 
	CONSTRAINT fk_irac_class_id FOREIGN KEY(irac_class_id) REFERENCES irac_classification (irac_class_id) ON DELETE CASCADE, 
	CONSTRAINT fk_irac_molregno FOREIGN KEY(molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE, 
	CONSTRAINT uk_mol_irac_class UNIQUE (irac_class_id, molregno)
);
CREATE TABLE molecule_atc_classification (
	mol_atc_id BIGINT NOT NULL, 
	level5 VARCHAR(10) NOT NULL, 
	molregno BIGINT NOT NULL, 
	CONSTRAINT pk_molatc_mol_atc_id PRIMARY KEY (mol_atc_id), 
	CONSTRAINT fk_molatc_level5 FOREIGN KEY(level5) REFERENCES atc_classification (level5) ON DELETE CASCADE, 
	CONSTRAINT fk_molatc_molregno FOREIGN KEY(molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE
);
CREATE TABLE assays (
	assay_id BIGINT NOT NULL, 
	doc_id BIGINT NOT NULL, 
	description VARCHAR(4000), 
	assay_type VARCHAR(1), 
	assay_test_type VARCHAR(20), 
	assay_category VARCHAR(20), 
	assay_organism VARCHAR(250), 
	assay_tax_id BIGINT, 
	assay_strain VARCHAR(200), 
	assay_tissue VARCHAR(100), 
	assay_cell_type VARCHAR(100), 
	assay_subcellular_fraction VARCHAR(100), 
	tid BIGINT, 
	relationship_type VARCHAR(1), 
	confidence_score SMALLINT, 
	curated_by VARCHAR(32), 
	src_id INTEGER NOT NULL, 
	src_assay_id VARCHAR(50), 
	chembl_id VARCHAR(20) NOT NULL, 
	cell_id BIGINT, 
	bao_format VARCHAR(11), 
	tissue_id BIGINT, 
	variant_id BIGINT, 
	aidx VARCHAR(200) NOT NULL, 
	CONSTRAINT pk_assays_assay_id PRIMARY KEY (assay_id), 
	CONSTRAINT fk_assays_assaytype FOREIGN KEY(assay_type) REFERENCES assay_type (assay_type) ON DELETE CASCADE, 
	CONSTRAINT fk_assays_cell_id FOREIGN KEY(cell_id) REFERENCES cell_dictionary (cell_id) ON DELETE CASCADE, 
	CONSTRAINT fk_assays_chembl_id FOREIGN KEY(chembl_id) REFERENCES chembl_id_lookup (chembl_id) ON DELETE CASCADE, 
	CONSTRAINT fk_assays_confscore FOREIGN KEY(confidence_score) REFERENCES confidence_score_lookup (confidence_score) ON DELETE CASCADE, 
	CONSTRAINT fk_assays_cur_by FOREIGN KEY(curated_by) REFERENCES curation_lookup (curated_by) ON DELETE CASCADE, 
	CONSTRAINT fk_assays_doc_id FOREIGN KEY(doc_id) REFERENCES docs (doc_id) ON DELETE CASCADE, 
	CONSTRAINT fk_assays_reltype FOREIGN KEY(relationship_type) REFERENCES relationship_type (relationship_type) ON DELETE CASCADE, 
	CONSTRAINT fk_assays_src_id FOREIGN KEY(src_id) REFERENCES source (src_id) ON DELETE CASCADE, 
	CONSTRAINT fk_assays_tid FOREIGN KEY(tid) REFERENCES target_dictionary (tid) ON DELETE CASCADE, 
	CONSTRAINT fk_assays_tissue_id FOREIGN KEY(tissue_id) REFERENCES tissue_dictionary (tissue_id) ON DELETE CASCADE, 
	CONSTRAINT fk_assays_variant_id FOREIGN KEY(variant_id) REFERENCES variant_sequences (variant_id) ON DELETE CASCADE, 
	CONSTRAINT fk_chembl_bao_format FOREIGN KEY(bao_format) REFERENCES bioassay_ontology (bao_id) ON DELETE CASCADE, 
	CONSTRAINT ck_assays_category CHECK (assay_category in ('screening','panel','confirmatory','summary','other')), 
	CONSTRAINT ck_assays_chemblid CHECK (chembl_id like ('CHEMBL%')), 
	CONSTRAINT ck_assays_testtype CHECK (assay_test_type in ('In vivo','In vitro','Ex vivo')), 
	CONSTRAINT uk_assays_chemblid UNIQUE (chembl_id)
);
CREATE TABLE compound_records (
	record_id BIGINT NOT NULL, 
	molregno BIGINT, 
	doc_id BIGINT NOT NULL, 
	compound_key VARCHAR(250), 
	compound_name VARCHAR(4000), 
	src_id INTEGER NOT NULL, 
	src_compound_id VARCHAR(150), 
	cidx VARCHAR(200) NOT NULL, 
	CONSTRAINT pk_cmpdrec_record_id PRIMARY KEY (record_id), 
	CONSTRAINT fk_cmpdrec_doc_id FOREIGN KEY(doc_id) REFERENCES docs (doc_id) ON DELETE CASCADE, 
	CONSTRAINT fk_cmpdrec_molregno FOREIGN KEY(molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE, 
	CONSTRAINT fk_cmpdrec_src_id FOREIGN KEY(src_id) REFERENCES source (src_id) ON DELETE CASCADE
);
CREATE TABLE binding_sites (
	site_id BIGINT NOT NULL, 
	site_name VARCHAR(200), 
	tid BIGINT, 
	CONSTRAINT pk_bindsite_id PRIMARY KEY (site_id), 
	CONSTRAINT fk_bindsite_tid FOREIGN KEY(tid) REFERENCES target_dictionary (tid) ON DELETE CASCADE
);
CREATE TABLE target_relations (
	tid BIGINT NOT NULL, 
	relationship VARCHAR(20) NOT NULL, 
	related_tid BIGINT NOT NULL, 
	targrel_id BIGINT NOT NULL, 
	CONSTRAINT target_relations_pk PRIMARY KEY (targrel_id), 
	CONSTRAINT fk_targrel_reltid FOREIGN KEY(related_tid) REFERENCES target_dictionary (tid) ON DELETE CASCADE, 
	CONSTRAINT fk_targrel_tid FOREIGN KEY(tid) REFERENCES target_dictionary (tid) ON DELETE CASCADE, 
	CONSTRAINT ck_targrel_rel CHECK (relationship in ('EQUIVALENT TO', 'OVERLAPS WITH', 'SUBSET OF', 'SUPERSET OF'))
);
CREATE TABLE compound_structures (
	molregno BIGINT NOT NULL, 
	molfile TEXT, 
	standard_inchi VARCHAR(4000), 
	standard_inchi_key VARCHAR(27) NOT NULL, 
	canonical_smiles VARCHAR(4000), 
	CONSTRAINT pk_cmpdstr_molregno PRIMARY KEY (molregno), 
	CONSTRAINT fk_cmpdstr_molregno FOREIGN KEY(molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE, 
	CONSTRAINT uk_cmpdstr_stdinch UNIQUE (standard_inchi), 
	CONSTRAINT uk_cmpdstr_stdinchkey UNIQUE (standard_inchi_key)
);
CREATE TABLE biotherapeutics (
	molregno BIGINT NOT NULL, 
	description VARCHAR(2000), 
	helm_notation VARCHAR(4000), 
	CONSTRAINT pk_biother_molregno PRIMARY KEY (molregno), 
	CONSTRAINT fk_biother_molregno FOREIGN KEY(molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE
);
CREATE TABLE compound_structural_alerts (
	cpd_str_alert_id BIGINT NOT NULL, 
	molregno BIGINT NOT NULL, 
	alert_id BIGINT NOT NULL, 
	CONSTRAINT pk_cpd_str_alert_id PRIMARY KEY (cpd_str_alert_id), 
	CONSTRAINT fk_cpd_str_alert_id FOREIGN KEY(alert_id) REFERENCES structural_alerts (alert_id) ON DELETE CASCADE, 
	CONSTRAINT fk_cpd_str_alert_molregno FOREIGN KEY(molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE, 
	CONSTRAINT uk_cpd_str_alert UNIQUE (molregno, alert_id)
);
CREATE TABLE molecule_hierarchy (
	molregno BIGINT NOT NULL, 
	parent_molregno BIGINT, 
	active_molregno BIGINT, 
	CONSTRAINT pk_molhier_molregno PRIMARY KEY (molregno), 
	CONSTRAINT fk_molhier_active_molregno FOREIGN KEY(active_molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE, 
	CONSTRAINT fk_molhier_molregno FOREIGN KEY(molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE, 
	CONSTRAINT fk_molhier_parent_molregno FOREIGN KEY(parent_molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE
);
CREATE TABLE molecule_synonyms (
	molregno BIGINT NOT NULL, 
	syn_type VARCHAR(50) NOT NULL, 
	molsyn_id BIGINT NOT NULL, 
	res_stem_id BIGINT, 
	synonyms VARCHAR(200), 
	CONSTRAINT pk_cmpdsyns_synid PRIMARY KEY (molsyn_id), 
	CONSTRAINT fk_cmpdsyns_molregno FOREIGN KEY(molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE, 
	CONSTRAINT fk_cmpdsyns_resstem FOREIGN KEY(res_stem_id) REFERENCES research_stem (res_stem_id) ON DELETE CASCADE, 
	CONSTRAINT uk_cmpdsyns UNIQUE (molregno, syn_type, synonyms)
);
CREATE TABLE molecule_hrac_classification (
	mol_hrac_id BIGINT NOT NULL, 
	hrac_class_id BIGINT NOT NULL, 
	molregno BIGINT NOT NULL, 
	CONSTRAINT molecule_hrac_classificationpk PRIMARY KEY (mol_hrac_id), 
	CONSTRAINT fk_hrac_class_id FOREIGN KEY(hrac_class_id) REFERENCES hrac_classification (hrac_class_id) ON DELETE CASCADE, 
	CONSTRAINT fk_hrac_molregno FOREIGN KEY(molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE, 
	CONSTRAINT uk_mol_hrac_class UNIQUE (hrac_class_id, molregno)
);
CREATE TABLE target_components (
	tid BIGINT NOT NULL, 
	component_id BIGINT NOT NULL, 
	targcomp_id BIGINT NOT NULL, 
	homologue SMALLINT NOT NULL, 
	CONSTRAINT pk_targcomp_id PRIMARY KEY (targcomp_id), 
	CONSTRAINT fk_targcomp_compid FOREIGN KEY(component_id) REFERENCES component_sequences (component_id) ON DELETE CASCADE, 
	CONSTRAINT fk_targcomp_tid FOREIGN KEY(tid) REFERENCES target_dictionary (tid) ON DELETE CASCADE, 
	CONSTRAINT ck_targcomp_hom CHECK (homologue in (0,1,2)), 
	CONSTRAINT uk_targcomp_tid_compid UNIQUE (tid, component_id)
);
CREATE TABLE compound_properties (
	molregno BIGINT NOT NULL, 
	mw_freebase NUMERIC(9, 2), 
	alogp NUMERIC(9, 2), 
	hba INTEGER, 
	hbd INTEGER, 
	psa NUMERIC(9, 2), 
	rtb INTEGER, 
	ro3_pass VARCHAR(3), 
	num_ro5_violations SMALLINT, 
	cx_most_apka NUMERIC(9, 2), 
	cx_most_bpka NUMERIC(9, 2), 
	cx_logp NUMERIC(9, 2), 
	cx_logd NUMERIC(9, 2), 
	molecular_species VARCHAR(50), 
	full_mwt NUMERIC(9, 2), 
	aromatic_rings INTEGER, 
	heavy_atoms INTEGER, 
	qed_weighted NUMERIC(3, 2), 
	mw_monoisotopic NUMERIC(11, 4), 
	full_molformula VARCHAR(100), 
	hba_lipinski INTEGER, 
	hbd_lipinski INTEGER, 
	num_lipinski_ro5_violations SMALLINT, 
	CONSTRAINT pk_cmpdprop_molregno PRIMARY KEY (molregno), 
	CONSTRAINT fk_cmpdprop_molregno FOREIGN KEY(molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE, 
	CONSTRAINT ck_cmpdprop_aromatic CHECK (aromatic_rings >= 0), 
	CONSTRAINT ck_cmpdprop_bpka CHECK (CX_MOST_BPKA>=0), 
	CONSTRAINT ck_cmpdprop_fullmw CHECK (full_mwt > 0), 
	CONSTRAINT ck_cmpdprop_hba CHECK (hba >= 0), 
	CONSTRAINT ck_cmpdprop_hba_lip CHECK (hba_lipinski >= 0), 
	CONSTRAINT ck_cmpdprop_hbd CHECK (hbd >= 0), 
	CONSTRAINT ck_cmpdprop_hbd_lip CHECK (hbd_lipinski >= 0), 
	CONSTRAINT ck_cmpdprop_heavy CHECK (heavy_atoms >= 0), 
	CONSTRAINT ck_cmpdprop_lip_ro5 CHECK (num_lipinski_ro5_violations in (0,1,2,3,4)), 
	CONSTRAINT ck_cmpdprop_mwfree CHECK (mw_freebase > 0), 
	CONSTRAINT ck_cmpdprop_psa CHECK (psa >= 0), 
	CONSTRAINT ck_cmpdprop_qed CHECK (qed_weighted >= 0), 
	CONSTRAINT ck_cmpdprop_ro3 CHECK (ro3_pass in ('Y','N')), 
	CONSTRAINT ck_cmpdprop_ro5 CHECK (num_ro5_violations in (0,1,2,3,4)), 
	CONSTRAINT ck_cmpdprop_rtb CHECK (rtb >= 0), 
	CONSTRAINT ck_cmpdprop_species CHECK (molecular_species in ('ACID','BASE','ZWITTERION','NEUTRAL'))
);
CREATE TABLE molecule_frac_classification (
	mol_frac_id BIGINT NOT NULL, 
	frac_class_id BIGINT NOT NULL, 
	molregno BIGINT NOT NULL, 
	CONSTRAINT molecule_frac_classificationpk PRIMARY KEY (mol_frac_id), 
	CONSTRAINT fk_frac_class_id FOREIGN KEY(frac_class_id) REFERENCES frac_classification (frac_class_id) ON DELETE CASCADE, 
	CONSTRAINT fk_frac_molregno FOREIGN KEY(molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE, 
	CONSTRAINT uk_mol_frac_class UNIQUE (frac_class_id, molregno)
);
CREATE TABLE drug_mechanism (
	mec_id BIGINT NOT NULL, 
	record_id BIGINT NOT NULL, 
	molregno BIGINT, 
	mechanism_of_action VARCHAR(250), 
	tid BIGINT, 
	site_id BIGINT, 
	action_type VARCHAR(50), 
	direct_interaction SMALLINT, 
	molecular_mechanism SMALLINT, 
	disease_efficacy SMALLINT, 
	mechanism_comment VARCHAR(2000), 
	selectivity_comment VARCHAR(1000), 
	binding_site_comment VARCHAR(1000), 
	variant_id BIGINT, 
	CONSTRAINT molecule_mechanism_pk PRIMARY KEY (mec_id), 
	CONSTRAINT fk_drugmec_actiontype FOREIGN KEY(action_type) REFERENCES action_type (action_type) ON DELETE CASCADE, 
	CONSTRAINT fk_drugmec_molregno FOREIGN KEY(molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE, 
	CONSTRAINT fk_drugmec_rec_id FOREIGN KEY(record_id) REFERENCES compound_records (record_id) ON DELETE CASCADE, 
	CONSTRAINT fk_drugmec_site_id FOREIGN KEY(site_id) REFERENCES binding_sites (site_id) ON DELETE CASCADE, 
	CONSTRAINT fk_drugmec_tid FOREIGN KEY(tid) REFERENCES target_dictionary (tid) ON DELETE CASCADE, 
	CONSTRAINT fk_drugmec_varid FOREIGN KEY(variant_id) REFERENCES variant_sequences (variant_id) ON DELETE CASCADE, 
	CONSTRAINT ck_drugmec_direct CHECK (direct_interaction in (0,1)), 
	CONSTRAINT ck_drugmec_efficacy CHECK (disease_efficacy in (0,1)), 
	CONSTRAINT ck_drugmec_molecular CHECK (molecular_mechanism in (0,1))
);
CREATE TABLE drug_indication (
	drugind_id BIGINT NOT NULL, 
	record_id BIGINT NOT NULL, 
	molregno BIGINT, 
	max_phase_for_ind SMALLINT, 
	mesh_id VARCHAR(20) NOT NULL, 
	mesh_heading VARCHAR(200) NOT NULL, 
	efo_id VARCHAR(20), 
	efo_term VARCHAR(200), 
	CONSTRAINT drugind_pk PRIMARY KEY (drugind_id), 
	CONSTRAINT drugind_molregno_fk FOREIGN KEY(molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE, 
	CONSTRAINT drugind_rec_fk FOREIGN KEY(record_id) REFERENCES compound_records (record_id) ON DELETE CASCADE, 
	CONSTRAINT drugind_phase_ck CHECK (MAX_PHASE_FOR_IND=0 OR MAX_PHASE_FOR_IND=1 OR MAX_PHASE_FOR_IND=2 OR MAX_PHASE_FOR_IND=3 OR MAX_PHASE_FOR_IND=4), 
	CONSTRAINT drugind_uk UNIQUE (record_id, mesh_id, efo_id)
);
CREATE TABLE assay_parameters (
	assay_param_id BIGINT NOT NULL, 
	assay_id BIGINT NOT NULL, 
	type VARCHAR(250) NOT NULL, 
	relation VARCHAR(50), 
	value NUMERIC, 
	units VARCHAR(100), 
	text_value VARCHAR(4000), 
	standard_type VARCHAR(250), 
	standard_relation VARCHAR(50), 
	standard_value NUMERIC, 
	standard_units VARCHAR(100), 
	standard_text_value VARCHAR(4000), 
	comments VARCHAR(4000), 
	CONSTRAINT pk_assay_param PRIMARY KEY (assay_param_id), 
	CONSTRAINT fk_assay_param_assayid FOREIGN KEY(assay_id) REFERENCES assays (assay_id) ON DELETE CASCADE, 
	CONSTRAINT uk_assay_param UNIQUE (assay_id, type)
);
CREATE TABLE activities (
	activity_id BIGINT NOT NULL, 
	assay_id BIGINT NOT NULL, 
	doc_id BIGINT, 
	record_id BIGINT NOT NULL, 
	molregno BIGINT, 
	standard_relation VARCHAR(50), 
	standard_value NUMERIC, 
	standard_units VARCHAR(100), 
	standard_flag SMALLINT, 
	standard_type VARCHAR(250), 
	activity_comment VARCHAR(4000), 
	data_validity_comment VARCHAR(30), 
	potential_duplicate SMALLINT, 
	pchembl_value NUMERIC(4, 2), 
	bao_endpoint VARCHAR(11), 
	uo_units VARCHAR(10), 
	qudt_units VARCHAR(70), 
	toid INTEGER, 
	upper_value NUMERIC, 
	standard_upper_value NUMERIC, 
	src_id INTEGER, 
	type VARCHAR(250) NOT NULL, 
	relation VARCHAR(50), 
	value NUMERIC, 
	units VARCHAR(100), 
	text_value VARCHAR(1000), 
	standard_text_value VARCHAR(1000), 
	CONSTRAINT pk_act_activity_id PRIMARY KEY (activity_id), 
	CONSTRAINT fk_act_assay_id FOREIGN KEY(assay_id) REFERENCES assays (assay_id) ON DELETE CASCADE, 
	CONSTRAINT fk_act_bao_endpoint FOREIGN KEY(bao_endpoint) REFERENCES bioassay_ontology (bao_id) ON DELETE CASCADE, 
	CONSTRAINT fk_act_doc_id FOREIGN KEY(doc_id) REFERENCES docs (doc_id) ON DELETE CASCADE, 
	CONSTRAINT fk_act_molregno FOREIGN KEY(molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE, 
	CONSTRAINT fk_act_record_id FOREIGN KEY(record_id) REFERENCES compound_records (record_id) ON DELETE CASCADE, 
	CONSTRAINT fk_act_src_id FOREIGN KEY(src_id) REFERENCES source (src_id) ON DELETE CASCADE, 
	CONSTRAINT fk_data_val_comm FOREIGN KEY(data_validity_comment) REFERENCES data_validity_lookup (data_validity_comment) ON DELETE CASCADE, 
	CONSTRAINT ck_potential_dup CHECK (POTENTIAL_DUPLICATE IN (0,1)), 
	CONSTRAINT ck_stand_flag CHECK (standard_flag in (0,1)), 
	CONSTRAINT ck_stand_relation CHECK (standard_relation in ('>','<','=','~','<=','>=','<<','>>'))
);
CREATE TABLE drug_warning (
	warning_id BIGINT NOT NULL, 
	record_id BIGINT, 
	molregno BIGINT, 
	warning_type VARCHAR(20), 
	warning_class VARCHAR(100), 
	warning_description VARCHAR(4000), 
	warning_country VARCHAR(1000), 
	warning_year INTEGER, 
	CONSTRAINT sys_c00117264 PRIMARY KEY (warning_id), 
	CONSTRAINT fk_warning_record_id FOREIGN KEY(record_id) REFERENCES compound_records (record_id) ON DELETE CASCADE
);
CREATE TABLE metabolism (
	met_id BIGINT NOT NULL, 
	drug_record_id BIGINT, 
	substrate_record_id BIGINT, 
	metabolite_record_id BIGINT, 
	pathway_id BIGINT, 
	pathway_key VARCHAR(50), 
	enzyme_name VARCHAR(200), 
	enzyme_tid BIGINT, 
	met_conversion VARCHAR(200), 
	organism VARCHAR(100), 
	tax_id BIGINT, 
	met_comment VARCHAR(1000), 
	CONSTRAINT pk_rec_met_id PRIMARY KEY (met_id), 
	CONSTRAINT fk_recmet_drug_recid FOREIGN KEY(drug_record_id) REFERENCES compound_records (record_id) ON DELETE CASCADE, 
	CONSTRAINT fk_recmet_met_recid FOREIGN KEY(metabolite_record_id) REFERENCES compound_records (record_id) ON DELETE CASCADE, 
	CONSTRAINT fk_recmet_sub_recid FOREIGN KEY(substrate_record_id) REFERENCES compound_records (record_id) ON DELETE CASCADE, 
	CONSTRAINT fk_recmet_tid FOREIGN KEY(enzyme_tid) REFERENCES target_dictionary (tid) ON DELETE CASCADE, 
	CONSTRAINT uk_recmet UNIQUE (drug_record_id, substrate_record_id, metabolite_record_id, pathway_id, enzyme_name, enzyme_tid, tax_id)
);
CREATE TABLE biotherapeutic_components (
	biocomp_id BIGINT NOT NULL, 
	molregno BIGINT NOT NULL, 
	component_id BIGINT NOT NULL, 
	CONSTRAINT pk_biocomp_id PRIMARY KEY (biocomp_id), 
	CONSTRAINT fk_biocomp_compid FOREIGN KEY(component_id) REFERENCES bio_component_sequences (component_id) ON DELETE CASCADE, 
	CONSTRAINT fk_biocomp_molregno FOREIGN KEY(molregno) REFERENCES biotherapeutics (molregno) ON DELETE CASCADE, 
	CONSTRAINT uk_biocomp UNIQUE (molregno, component_id)
);
CREATE TABLE assay_class_map (
	ass_cls_map_id BIGINT NOT NULL, 
	assay_id BIGINT NOT NULL, 
	assay_class_id BIGINT NOT NULL, 
	CONSTRAINT pk_assay_cls_map PRIMARY KEY (ass_cls_map_id), 
	CONSTRAINT fk_ass_cls_map_assay FOREIGN KEY(assay_id) REFERENCES assays (assay_id) ON DELETE CASCADE, 
	CONSTRAINT fk_ass_cls_map_class FOREIGN KEY(assay_class_id) REFERENCES assay_classification (assay_class_id) ON DELETE CASCADE, 
	CONSTRAINT uk_ass_cls_map UNIQUE (assay_id, assay_class_id)
);
CREATE TABLE formulations (
	product_id VARCHAR(30) NOT NULL, 
	ingredient VARCHAR(200), 
	strength VARCHAR(300), 
	record_id BIGINT NOT NULL, 
	molregno BIGINT, 
	formulation_id BIGINT NOT NULL, 
	CONSTRAINT pk_formulations_id PRIMARY KEY (formulation_id), 
	CONSTRAINT fk_formulations_molregno FOREIGN KEY(molregno) REFERENCES molecule_dictionary (molregno) ON DELETE CASCADE, 
	CONSTRAINT fk_formulations_productid FOREIGN KEY(product_id) REFERENCES products (product_id) ON DELETE CASCADE, 
	CONSTRAINT fk_formulations_recid FOREIGN KEY(record_id) REFERENCES compound_records (record_id) ON DELETE CASCADE, 
	CONSTRAINT uk_formulations UNIQUE (product_id, record_id)
);
CREATE TABLE site_components (
	sitecomp_id BIGINT NOT NULL, 
	site_id BIGINT NOT NULL, 
	component_id BIGINT, 
	domain_id BIGINT, 
	site_residues VARCHAR(2000), 
	CONSTRAINT pk_sitecomp_id PRIMARY KEY (sitecomp_id), 
	CONSTRAINT fk_sitecomp_compid FOREIGN KEY(component_id) REFERENCES component_sequences (component_id) ON DELETE CASCADE, 
	CONSTRAINT fk_sitecomp_domainid FOREIGN KEY(domain_id) REFERENCES domains (domain_id) ON DELETE CASCADE, 
	CONSTRAINT fk_sitecomp_siteid FOREIGN KEY(site_id) REFERENCES binding_sites (site_id) ON DELETE CASCADE, 
	CONSTRAINT uk_sitecomp UNIQUE (site_id, component_id, domain_id)
);
CREATE TABLE indication_refs (
	indref_id BIGINT NOT NULL, 
	drugind_id BIGINT NOT NULL, 
	ref_type VARCHAR(50) NOT NULL, 
	ref_id VARCHAR(4000) NOT NULL, 
	ref_url VARCHAR(4000) NOT NULL, 
	CONSTRAINT indication_refs_pk PRIMARY KEY (indref_id), 
	CONSTRAINT indref_drugind_fk FOREIGN KEY(drugind_id) REFERENCES drug_indication (drugind_id) ON DELETE CASCADE, 
	CONSTRAINT indref_uk UNIQUE (drugind_id, ref_type, ref_id)
);
CREATE TABLE mechanism_refs (
	mecref_id BIGINT NOT NULL, 
	mec_id BIGINT NOT NULL, 
	ref_type VARCHAR(50) NOT NULL, 
	ref_id VARCHAR(200), 
	ref_url VARCHAR(400), 
	CONSTRAINT pk_mechanism_refs PRIMARY KEY (mecref_id), 
	CONSTRAINT fk_mechanism_refs_mecid FOREIGN KEY(mec_id) REFERENCES drug_mechanism (mec_id) ON DELETE CASCADE, 
	CONSTRAINT ck_mechanism_ref_type CHECK (ref_type in ('PMDA','ISBN','IUPHAR','DOI','EMA','PubMed','USPO','DailyMed','FDA','Expert','Other','InterPro','Wikipedia','UniProt','KEGG','PMC','ClinicalTrials','PubChem','Patent','BNF','HMA')), 
	CONSTRAINT uk_mechanism_refs UNIQUE (mec_id, ref_type, ref_id)
);
CREATE TABLE activity_properties (
	ap_id BIGINT NOT NULL, 
	activity_id BIGINT NOT NULL, 
	type VARCHAR(250) NOT NULL, 
	relation VARCHAR(50), 
	value NUMERIC, 
	units VARCHAR(100), 
	text_value VARCHAR(1000), 
	standard_type VARCHAR(250), 
	standard_relation VARCHAR(50), 
	standard_value NUMERIC, 
	standard_units VARCHAR(100), 
	standard_text_value VARCHAR(1000), 
	comments VARCHAR(4000), 
	result_flag SMALLINT NOT NULL, 
	CONSTRAINT pk_actprop_ap_id PRIMARY KEY (ap_id), 
	CONSTRAINT fk_activity_property FOREIGN KEY(activity_id) REFERENCES activities (activity_id) ON DELETE CASCADE, 
	CONSTRAINT uk_actprop_id_type UNIQUE (activity_id, type)
);
CREATE TABLE warning_refs (
	warnref_id BIGINT NOT NULL, 
	warning_id BIGINT, 
	ref_type VARCHAR(50), 
	ref_id VARCHAR(4000), 
	ref_url VARCHAR(4000), 
	CONSTRAINT sys_c00117259 PRIMARY KEY (warnref_id), 
	CONSTRAINT fk_warnref_warn_id FOREIGN KEY(warning_id) REFERENCES drug_warning (warning_id) ON DELETE CASCADE
);
CREATE TABLE predicted_binding_domains (
	predbind_id BIGINT NOT NULL, 
	activity_id BIGINT, 
	site_id BIGINT, 
	prediction_method VARCHAR(50), 
	confidence VARCHAR(10), 
	CONSTRAINT pk_predbinddom_predbind_id PRIMARY KEY (predbind_id), 
	CONSTRAINT fk_predbinddom_act_id FOREIGN KEY(activity_id) REFERENCES activities (activity_id) ON DELETE CASCADE, 
	CONSTRAINT fk_predbinddom_site_id FOREIGN KEY(site_id) REFERENCES binding_sites (site_id) ON DELETE CASCADE, 
	CONSTRAINT ck_predbinddom_conf CHECK (confidence in ('high','medium','low')), 
	CONSTRAINT ck_predbinddom_method CHECK (prediction_method in ('Manual','Single domain','Multi domain'))
);
CREATE TABLE metabolism_refs (
	metref_id BIGINT NOT NULL, 
	met_id BIGINT NOT NULL, 
	ref_type VARCHAR(50) NOT NULL, 
	ref_id VARCHAR(200), 
	ref_url VARCHAR(400), 
	CONSTRAINT pk_metref_id PRIMARY KEY (metref_id), 
	CONSTRAINT fk_metref_met_id FOREIGN KEY(met_id) REFERENCES metabolism (met_id) ON DELETE CASCADE, 
	CONSTRAINT uk_metref UNIQUE (met_id, ref_type, ref_id)
);
CREATE TABLE activity_supp_map (
	actsm_id BIGINT NOT NULL, 
	activity_id BIGINT NOT NULL, 
	smid BIGINT NOT NULL, 
	CONSTRAINT pk_actsm_id PRIMARY KEY (actsm_id), 
	CONSTRAINT fk_act_smid FOREIGN KEY(smid) REFERENCES activity_smid (smid) ON DELETE CASCADE, 
	CONSTRAINT fk_supp_act FOREIGN KEY(activity_id) REFERENCES activities (activity_id) ON DELETE CASCADE
);
CREATE TABLE ligand_eff (
	activity_id BIGINT NOT NULL, 
	bei NUMERIC(9, 2), 
	sei NUMERIC(9, 2), 
	le NUMERIC(9, 2), 
	lle NUMERIC(9, 2), 
	CONSTRAINT pk_ligeff_actid PRIMARY KEY (activity_id), 
	CONSTRAINT fk_ligeff_actid FOREIGN KEY(activity_id) REFERENCES activities (activity_id) ON DELETE CASCADE, 
	CONSTRAINT ck_ligeff_bei CHECK (bei > 0), 
	CONSTRAINT ck_ligeff_sei CHECK (sei > 0)
);
CREATE INDEX idx_moldict_pref_name ON molecule_dictionary (pref_name);
CREATE INDEX idx_moldict_max_phase ON molecule_dictionary (max_phase);
CREATE INDEX idx_moldict_ther_flag ON molecule_dictionary (therapeutic_flag);
CREATE UNIQUE INDEX idx_moldict_chembl_id ON molecule_dictionary (chembl_id);
CREATE UNIQUE INDEX organism_class_pk ON organism_class (oc_id);
CREATE INDEX idx_assays_doc_id ON assays (doc_id);
CREATE INDEX idx_assays_desc ON assays (description);
CREATE INDEX idx_assays_src_id ON assays (src_id);
CREATE UNIQUE INDEX idx_assays_chembl_id ON assays (chembl_id);
CREATE INDEX tmp_bao_format ON assays (bao_format);
CREATE INDEX idx_assay_assay_id ON assays (assay_type);
CREATE INDEX idx_docs_pmid ON docs (pubmed_id);
CREATE INDEX bmx_doc_iss ON docs (issue);
CREATE UNIQUE INDEX pk_doc_doc_id ON docs (doc_id);
CREATE INDEX bmx_doc_vol ON docs (volume);
CREATE INDEX bmx_doc_jrnl ON docs (journal);
CREATE INDEX bmx_doc_year ON docs (year);
CREATE UNIQUE INDEX pk_rt_rt ON relationship_type (relationship_type);
CREATE INDEX idx_td_pname ON target_dictionary (pref_name);
CREATE INDEX idx_td_org ON target_dictionary (organism);
CREATE INDEX idx_td_taxid ON target_dictionary (tax_id);
CREATE INDEX idx_td_t_type ON target_dictionary (target_type);
CREATE INDEX idx_td_chembl_id ON target_dictionary (chembl_id);
CREATE UNIQUE INDEX pk_tt_tt ON target_type (target_type);
CREATE UNIQUE INDEX tissue_dictionary_pk ON tissue_dictionary (tissue_id);
CREATE UNIQUE INDEX pk_comp_rec_recid ON compound_records (record_id);
CREATE INDEX idx_comp_rec_srccpid ON compound_records (src_compound_id);
CREATE INDEX fk_comp_rec_docid ON compound_records (doc_id);
CREATE INDEX idx_comp_rec_ckey ON compound_records (compound_key);
CREATE INDEX idx_comp_rec_cidx ON compound_records (cidx);
CREATE INDEX idx_comp_rec_srcid ON compound_records (src_id);
CREATE INDEX fk_comp_rec_molregno ON compound_records (molregno);
CREATE INDEX idx_actsupp_type ON activity_supp (type);
CREATE INDEX idx_actsupp_std_type ON activity_supp (standard_type);
CREATE INDEX idx_actsupp_rel ON activity_supp (relation);
CREATE INDEX idx_actsupp_std_val ON activity_supp (standard_value);
CREATE INDEX idx_actsupp_val ON activity_supp (value);
CREATE INDEX idx_actsupp_std_text ON activity_supp (standard_text_value);
CREATE INDEX idx_actsupp_units ON activity_supp (units);
CREATE INDEX idx_actsupp_std_rel ON activity_supp (standard_relation);
CREATE INDEX idx_actsupp_std_units ON activity_supp (standard_units);
CREATE INDEX idx_actsupp_text ON activity_supp (text_value);
CREATE UNIQUE INDEX protclass_pk ON protein_classification (protein_class_id);
CREATE INDEX idx_cmpdstr_smiles ON compound_structures (canonical_smiles);
CREATE UNIQUE INDEX compound_structures_pk ON compound_structures (molregno);
CREATE INDEX idx_cmpdstr_stdinchi ON compound_structures (standard_inchi);
CREATE INDEX idx_cmpdstr_stdkey ON compound_structures (standard_inchi_key);
CREATE UNIQUE INDEX bio_component_seqs_pk ON bio_component_sequences (component_id);
CREATE UNIQUE INDEX drug_indication_pk ON drug_indication (drugind_id);
CREATE UNIQUE INDEX mechanism_refs_uk ON mechanism_refs (mec_id, ref_type, ref_id);
CREATE UNIQUE INDEX mechanism_refs_pk ON mechanism_refs (mecref_id);
CREATE INDEX idx_assay_param_text ON assay_parameters (text_value);
CREATE INDEX idx_assay_param_std_val ON assay_parameters (standard_value);
CREATE INDEX idx_assay_param_rel ON assay_parameters (relation);
CREATE INDEX idx_assay_param_std_units ON assay_parameters (standard_units);
CREATE INDEX idx_assay_param_std_rel ON assay_parameters (standard_relation);
CREATE INDEX idx_assay_param_std_type ON assay_parameters (standard_type);
CREATE INDEX idx_assay_param_std_text ON assay_parameters (standard_text_value);
CREATE INDEX idx_assay_param_type ON assay_parameters (type);
CREATE INDEX idx_assay_param_units ON assay_parameters (units);
CREATE INDEX idx_assay_param_val ON assay_parameters (value);
CREATE UNIQUE INDEX assay_classification_pk ON assay_classification (assay_class_id);
CREATE UNIQUE INDEX structural_alert_set_pk ON structural_alert_sets (alert_set_id);
CREATE INDEX idx_actprop_resflag ON activity_properties (result_flag);
CREATE INDEX idx_actprop_val ON activity_properties (standard_value);
CREATE INDEX idx_act_prop_text ON activity_properties (standard_text_value);
CREATE INDEX idx_actprop_type ON activity_properties (standard_type);
CREATE INDEX idx_actprop_units ON activity_properties (standard_units);
CREATE INDEX idx_actprop_relation ON activity_properties (standard_relation);
CREATE INDEX idx_act_text ON activities (text_value);
CREATE INDEX fk_act_doc_id ON activities (doc_id);
CREATE INDEX idx_act_std_text ON activities (standard_text_value);
CREATE INDEX idx_act_type ON activities (type);
CREATE INDEX idx_act_std_type ON activities (standard_type);
CREATE INDEX fk_act_record_id ON activities (record_id);
CREATE INDEX idx_act_src_id ON activities (src_id);
CREATE INDEX idx_act_std_unit ON activities (standard_units);
CREATE INDEX idx_act_units ON activities (units);
CREATE INDEX idx_act_rel ON activities (relation);
CREATE INDEX idx_act_val ON activities (value);
CREATE INDEX idx_act_upper ON activities (upper_value);
CREATE INDEX idx_acc_relation ON activities (standard_relation);
CREATE INDEX idx_act_std_upper ON activities (standard_upper_value);
CREATE INDEX idx_act_pchembl ON activities (pchembl_value);
CREATE INDEX fk_act_molregno ON activities (molregno);
CREATE INDEX idx_act_std_val ON activities (standard_value);
CREATE INDEX fk_act_assay_id ON activities (assay_id);
CREATE INDEX idx_molhier_parent ON molecule_hierarchy (parent_molregno);
CREATE UNIQUE INDEX pk_actsmid ON activity_supp_map (actsm_id);
CREATE INDEX idx_cp_alogp ON compound_properties (alogp);
CREATE INDEX idx_cp_ro5 ON compound_properties (num_ro5_violations);
CREATE UNIQUE INDEX pk_com_molreg ON compound_properties (molregno);
CREATE INDEX idx_cp_hbd ON compound_properties (hbd);
CREATE INDEX idx_cp_hba ON compound_properties (hba);
CREATE INDEX idx_cp_mw ON compound_properties (mw_freebase);
CREATE INDEX idx_cp_rtb ON compound_properties (rtb);
CREATE INDEX idx_cp_psa ON compound_properties (psa);
