package annex5

import (
	"strconv"

	"github.com/devsstudio/gosunat/enums"
	"github.com/devsstudio/gosunat/helpers"
	"github.com/devsstudio/gosunat/tables/table10"
	"github.com/devsstudio/gosunat/tables/table11"
	"github.com/devsstudio/gosunat/validators"

	"github.com/asaskevich/govalidator"
)

const (
	TICKET_1 = "1"
	TICKET_2 = "2"
	TICKET_3 = "3"
	TICKET_4 = "4"
	TICKET_5 = "5"
)

const (
	BVME_1 = "1"
	BVME_2 = "2"
	BVME_5 = "5"
)

type Annex5Rule struct {
	Sale     *RuleDetails `jsOn:"sale"`
	Purchase *RuleDetails `jsOn:"purchase"`
	Other    RuleDetails  `jsOn:"other"`
}

type RuleDetails struct {
	On     []enums.AccountingFile `jsOn:"on,omitempty"`
	States []int                  `jsOn:"states"`
}

func getRules() map[table10.Table10]Annex5Rule {
	return map[table10.Table10]Annex5Rule{
		table10.DOC_00: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1, enums.AF_FILE_14_2},
				States: []int{0, 1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_2, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_01: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1, enums.AF_FILE_14_2},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_02: {
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1},
				States: []int{0, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_03: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1, enums.AF_FILE_14_2},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_04: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_05: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_06: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_07: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1, enums.AF_FILE_14_2},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_08: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1, enums.AF_FILE_14_2},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_09: {
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_10: {
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_11: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_12: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1, enums.AF_FILE_14_2},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_13: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_14: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_15: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_16: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_17: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_18: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_19: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_20: {
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_21: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1},
				States: []int{0, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_22: {
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_23: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{0, 1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_24: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_25: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{0, 1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_26: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_27: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{0, 1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_28: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{0, 1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_29: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{0, 1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_30: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{0, 1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_31: {
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_32: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{0, 1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_33: {
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_34: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{0, 1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_35: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{0, 1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_36: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_37: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_40: {
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_41: {
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_42: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{0, 1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_43: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_44: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{0, 1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1},
				States: []int{0, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_45: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{0, 1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_46: {
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_48: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{0, 1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_49: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{0, 1, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1},
				States: []int{0, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_50: {
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_51: {
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_52: {
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_53: {
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_54: {
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_55: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_56: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{0, 1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_87: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_88: {
			Sale: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_14, enums.AF_FILE_14_1},
				States: []int{1, 2, 8, 9},
			},
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_89: {
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_91: {
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8_2},
				States: []int{0, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_96: {
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8, enums.AF_FILE_8_1, enums.AF_FILE_8_3},
				States: []int{0, 1, 6, 7, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_97: {
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8_2},
				States: []int{0, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
		table10.DOC_98: {
			Purchase: &RuleDetails{
				On:     []enums.AccountingFile{enums.AF_FILE_8_2},
				States: []int{0, 9},
			},
			Other: RuleDetails{
				States: []int{1, 8, 9},
			},
		},
	}
}

func ValidateDocumentCode(fileCode string, documentCode table10.Table10) bool {
	rules := getRules()

	// Verificar si la regla existe
	if rule, exists := rules[documentCode]; exists {
		// Validar según sea venta o compra
		if validators.IsSale(fileCode) {
			return rule.Sale != nil
		} else if validators.IsPurchase(fileCode) {
			return rule.Purchase != nil
		}
		return true
	}

	return false
}

func ValidateDocument(fileCode string, documentCode table10.Table10, state int) bool {
	rules := getRules()

	// Verificar si la regla existe
	if rule, exists := rules[documentCode]; exists {
		// Validar si es venta
		if validators.IsSale(fileCode) {
			return rule.Sale != nil &&
				helpers.ContainsAsString(rule.Sale.On, fileCode) &&
				helpers.Contains(rule.Sale.States, state)
		}

		// Validar si es compra
		if validators.IsPurchase(fileCode) {
			return rule.Purchase != nil &&
				helpers.ContainsAsString(rule.Purchase.On, fileCode) &&
				helpers.Contains(rule.Purchase.States, state)
		}

		// Validar si es otra categoría
		return helpers.Contains(rule.Other.States, state)
	}

	// Si la regla no existe, devolver false
	return false
}

func ValidateSerie(fileCode string, documentCode table10.Table10, documentSerie string, electronic bool) bool {

	ticketSlice := []string{TICKET_1, TICKET_2, TICKET_3, TICKET_4, TICKET_5}
	bvmeTicketTypes := []string{BVME_1, BVME_2, BVME_5}

	switch documentCode {
	case table10.DOC_00:
		return govalidator.IsAlphanumeric(documentSerie) && len(documentSerie) >= 0 && len(documentSerie) <= 20
	case table10.DOC_01:
		if electronic {
			return documentSerie == "E001" || validators.AlphanumericStartWith(documentSerie, "F", 4, 4)
		}
		return govalidator.IsInt(documentSerie) && len(documentSerie) == 4
	case table10.DOC_02:
		if validators.IsSale(fileCode) {
			return false
		}
		if electronic {
			return documentSerie == "E001"
		}
		return govalidator.IsInt(documentSerie) && len(documentSerie) == 4
	case table10.DOC_03:
		if electronic {
			return documentSerie == "EB01" || validators.AlphanumericStartWith(documentSerie, "B", 4, 4)
		}
		return govalidator.IsInt(documentSerie) && len(documentSerie) == 4
	case table10.DOC_04:
		if electronic {
			return documentSerie == "E001"
		}
		return govalidator.IsInt(documentSerie) && len(documentSerie) == 4
	case table10.DOC_05:
		return helpers.Contains(ticketSlice, documentSerie)
	case table10.DOC_06:
		intSerie, _ := strconv.Atoi(documentSerie)
		return govalidator.IsInt(documentSerie) && len(documentSerie) == 4 && intSerie > 0
	case table10.DOC_07, table10.DOC_08:
		if electronic {
			return helpers.Contains([]string{"E001", "EB01"}, documentSerie) || validators.AlphanumericStartWith(documentSerie, "F", 4, 4) || validators.AlphanumericStartWith(documentSerie, "B", 4, 4)
		}
		return govalidator.IsInt(documentSerie) && len(documentSerie) == 4
	case table10.DOC_09:
		if validators.IsSaleOrPurchase(fileCode) {
			return false
		}
		return govalidator.IsAlphanumeric(documentSerie) && len(documentSerie) == 4
	case table10.DOC_10:
		if validators.IsSale(fileCode) {
			return false
		} else {
			return documentSerie == "1683"
		}
	case table10.DOC_11:
		return govalidator.IsAlphanumeric(documentSerie) && validators.IsLength(documentSerie, 0, 20)
	case table10.DOC_12:
		return govalidator.IsAlphanumeric(documentSerie) && validators.IsLength(documentSerie, 1, 20)
	case table10.DOC_13, table10.DOC_14, table10.DOC_15, table10.DOC_16, table10.DOC_17, table10.DOC_18, table10.DOC_19:
		return govalidator.IsAlphanumeric(documentSerie) && validators.IsLength(documentSerie, 0, 20)
	case table10.DOC_20:
		if validators.IsSaleOrPurchase(fileCode) {
			return false
		} else {
			return govalidator.IsAlphanumeric(documentSerie) && validators.IsLength(documentSerie, 4, 4)
		}
	case table10.DOC_21:
		return govalidator.IsAlphanumeric(documentSerie) && validators.IsLength(documentSerie, 0, 20)
	case table10.DOC_22:
		if validators.IsSale(fileCode) {
			return false
		} else {
			return helpers.Contains([]string{"0820"}, documentSerie)
		}
	case table10.DOC_23:
		return govalidator.IsInt(documentSerie) && validators.IsLength(documentSerie, 4, 4)
	case table10.DOC_24:
		return govalidator.IsAlphanumeric(documentSerie) && validators.IsLength(documentSerie, 0, 20)
	case table10.DOC_25:
		return govalidator.IsInt(documentSerie) && validators.IsLength(documentSerie, 4, 4)
	case table10.DOC_26, table10.DOC_27, table10.DOC_28, table10.DOC_29, table10.DOC_30:
		return govalidator.IsAlphanumeric(documentSerie) && validators.IsLength(documentSerie, 0, 20)
	case table10.DOC_31:
		if validators.IsSaleOrPurchase(fileCode) {
			return false
		} else {
			return govalidator.IsAlphanumeric(documentSerie) && validators.IsLength(documentSerie, 4, 4)
		}
	case table10.DOC_32:
		return govalidator.IsAlphanumeric(documentSerie) && validators.IsLength(documentSerie, 0, 20)
	case table10.DOC_33:
		if validators.IsSaleOrPurchase(fileCode) {
			return false
		} else {
			return govalidator.IsAlphanumeric(documentSerie) && validators.IsLength(documentSerie, 4, 4)
		}
	case table10.DOC_34, table10.DOC_35, table10.DOC_36:
		return govalidator.IsInt(documentSerie) && validators.IsLength(documentSerie, 4, 4)
	case table10.DOC_37:
		return govalidator.IsAlphanumeric(documentSerie) && validators.IsLength(documentSerie, 0, 20)
	case table10.DOC_40, table10.DOC_41:
		if validators.IsSaleOrPurchase(fileCode) {
			return false
		} else {
			return govalidator.IsAlphanumeric(documentSerie) && validators.IsLength(documentSerie, 4, 4)
		}
	case table10.DOC_42, table10.DOC_43, table10.DOC_44, table10.DOC_45:
		return govalidator.IsAlphanumeric(documentSerie) && validators.IsLength(documentSerie, 0, 20)
	case table10.DOC_46:
		if validators.IsSale(fileCode) {
			return false
		} else {
			return govalidator.IsInt(documentSerie) && validators.IsLength(documentSerie, 4, 4)
		}
	case table10.DOC_48:
		return validators.IsIntGreaterThan(documentSerie, 0) && validators.IsLength(documentSerie, 4, 4)
	case table10.DOC_49:
		return govalidator.IsAlphanumeric(documentSerie) && validators.IsLength(documentSerie, 0, 20)
	case table10.DOC_50, table10.DOC_51, table10.DOC_52, table10.DOC_53, table10.DOC_54:
		if validators.IsSale(fileCode) {
			return false
		} else {
			return table11.ValidateCode(documentSerie)
		}
	case table10.DOC_55:
		if validators.IsSale(fileCode) {
			return helpers.Contains(bvmeTicketTypes, documentSerie)
		} else if validators.IsPurchase(fileCode) {
			return helpers.Contains(bvmeTicketTypes, documentSerie) && !helpers.Contains([]string{"5"}, documentSerie)
		} else {
			return false
		}
	case table10.DOC_56:
		return validators.IsIntGreaterThan(documentSerie, 0) && validators.IsLength(documentSerie, 4, 4)
	case table10.DOC_87, table10.DOC_88:
		return govalidator.IsAlphanumeric(documentSerie) && validators.IsLength(documentSerie, 0, 20)
	case table10.DOC_89:
		if validators.IsSale(fileCode) {
			return false
		} else {
			return validators.IsIntGreaterThan(documentSerie, 0) && validators.IsLength(documentSerie, 4, 4)
		}
	case table10.DOC_91, table10.DOC_96, table10.DOC_97, table10.DOC_98:
		if validators.IsSale(fileCode) {
			return false
		} else {
			return govalidator.IsAlphanumeric(documentSerie) && validators.IsLength(documentSerie, 0, 20)
		}
	default:
		return false
	}
}

func ValidateCorrelative(fileCode string, documentCode table10.Table10, documentCorrelative string, electronic bool) bool {
	switch documentCode {
	case table10.DOC_00:
		return govalidator.IsAlphanumeric(documentCorrelative) && validators.IsLength(documentCorrelative, 1, 20)
	case table10.DOC_01:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 8)
	case table10.DOC_02:
		if validators.IsSale(fileCode) {
			return false
		}
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 7)
	case table10.DOC_03:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 8)
	case table10.DOC_04:
		if electronic {
			return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 8)
		}
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 7)
	case table10.DOC_05:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 11)
	case table10.DOC_06, table10.DOC_07, table10.DOC_08:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 8)
	case table10.DOC_09:
		return govalidator.IsAlphanumeric(documentCorrelative) && validators.IsLength(documentCorrelative, 1, 8)
	case table10.DOC_10:
		if validators.IsSale(fileCode) {
			return false
		}
		return govalidator.IsInt(documentCorrelative) && validators.IsLength(documentCorrelative, 1, 20)
	case table10.DOC_11:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 15)
	case table10.DOC_12:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 20)
	case table10.DOC_13, table10.DOC_14, table10.DOC_15:
		return govalidator.IsAlphanumeric(documentCorrelative) && validators.IsLength(documentCorrelative, 1, 20) && validators.IsNotAllZeroes(documentCorrelative)
	case table10.DOC_16, table10.DOC_17, table10.DOC_18, table10.DOC_19:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 20)
	case table10.DOC_20:
		return govalidator.IsAlphanumeric(documentCorrelative) && validators.IsLength(documentCorrelative, 1, 8)
	case table10.DOC_21, table10.DOC_24, table10.DOC_26:
		return govalidator.IsAlphanumeric(documentCorrelative) &&
			validators.IsLength(documentCorrelative, 1, 20) &&
			validators.IsNotAllZeroes(documentCorrelative)
	case table10.DOC_22:
		if validators.IsSale(fileCode) {
			return false
		}
		return govalidator.IsInt(documentCorrelative) && validators.IsLength(documentCorrelative, 1, 20)
	case table10.DOC_23, table10.DOC_25:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 7)
	case table10.DOC_27, table10.DOC_28, table10.DOC_29, table10.DOC_30:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 20)
	case table10.DOC_31:
		return govalidator.IsAlphanumeric(documentCorrelative) && validators.IsLength(documentCorrelative, 1, 8)
	case table10.DOC_32:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 20)
	case table10.DOC_33:
		return govalidator.IsAlphanumeric(documentCorrelative) && validators.IsLength(documentCorrelative, 1, 8)
	case table10.DOC_34, table10.DOC_35:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 7)
	case table10.DOC_36:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 8)
	case table10.DOC_37:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 20)
	case table10.DOC_40, table10.DOC_41:
		return govalidator.IsAlphanumeric(documentCorrelative) && validators.IsLength(documentCorrelative, 1, 8)
	case table10.DOC_42, table10.DOC_43, table10.DOC_44, table10.DOC_45:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 20)
	case table10.DOC_46:
		if validators.IsSale(fileCode) {
			return false
		} else {
			return govalidator.IsInt(documentCorrelative) && validators.IsLength(documentCorrelative, 1, 20)
		}
	case table10.DOC_48:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 7)
	case table10.DOC_49:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 20)
	case table10.DOC_50, table10.DOC_51, table10.DOC_52, table10.DOC_53:
		if validators.IsSale(fileCode) {
			return false
		} else {
			return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 6)
		}
	case table10.DOC_54:
		if validators.IsSale(fileCode) {
			return false
		} else {
			return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 20)
		}
	case table10.DOC_55, table10.DOC_56:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 11)
	case table10.DOC_87, table10.DOC_88:
		return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 20)
	case table10.DOC_89:
		if validators.IsSale(fileCode) {
			return false
		} else {
			return validators.IsIntGreaterThan(documentCorrelative, 0) && validators.IsLength(documentCorrelative, 1, 7)
		}
	case table10.DOC_91, table10.DOC_96, table10.DOC_97, table10.DOC_98:
		if validators.IsSale(fileCode) {
			return false
		} else {
			return govalidator.IsAlphanumeric(documentCorrelative) &&
				validators.IsLength(documentCorrelative, 1, 20) &&
				validators.IsNotAllZeroes(documentCorrelative)
		}
	default:
		return false
	}
}
