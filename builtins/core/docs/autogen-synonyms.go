package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`echo`:            `out`,
	`!or`:             `or`,
	`!catch`:          `catch`,
	`!global`:         `global`,
	`!event`:          `event`,
	`(`:               `brace-quote`,
	`!and`:            `and`,
	`!if`:             `if`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!set`:            `set`,
	`f`:               `f`,
	`try`:             `try`,
	`export`:          `export`,
	`swivel-table`:    `swivel-table`,
	`err`:             `err`,
	`tout`:            `tout`,
	`rx`:              `rx`,
	`and`:             `and`,
	`getfile`:         `getfile`,
	`out`:             `out`,
	`pt`:              `pt`,
	`catch`:           `catch`,
	`append`:          `append`,
	`prepend`:         `prepend`,
	`murex-docs`:      `murex-docs`,
	`>>`:              `>>`,
	`ttyfd`:           `ttyfd`,
	`alter`:           `alter`,
	`get`:             `get`,
	`g`:               `g`,
	`or`:              `or`,
	`if`:              `if`,
	`global`:          `global`,
	`event`:           `event`,
	`swivel-datatype`: `swivel-datatype`,
	`>`:               `>`,
	`set`:             `set`,
	`post`:            `post`,
	`brace-quote`:     `brace-quote`,
	`read`:            `read`,
	`tread`:           `tread`,
	`trypipe`:         `trypipe`,
}
