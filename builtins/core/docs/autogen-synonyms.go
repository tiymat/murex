package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!and`:            `and`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!global`:         `global`,
	`!set`:            `set`,
	`!event`:          `event`,
	`(`:               `brace-quote`,
	`echo`:            `out`,
	`!or`:             `or`,
	`!if`:             `if`,
	`!catch`:          `catch`,
	`rx`:              `rx`,
	`and`:             `and`,
	`get`:             `get`,
	`brace-quote`:     `brace-quote`,
	`ttyfd`:           `ttyfd`,
	`g`:               `g`,
	`try`:             `try`,
	`global`:          `global`,
	`swivel-datatype`: `swivel-datatype`,
	`swivel-table`:    `swivel-table`,
	`out`:             `out`,
	`f`:               `f`,
	`tread`:           `tread`,
	`catch`:           `catch`,
	`trypipe`:         `trypipe`,
	`append`:          `append`,
	`err`:             `err`,
	`export`:          `export`,
	`event`:           `event`,
	`murex-docs`:      `murex-docs`,
	`read`:            `read`,
	`alter`:           `alter`,
	`set`:             `set`,
	`if`:              `if`,
	`getfile`:         `getfile`,
	`tout`:            `tout`,
	`pt`:              `pt`,
	`or`:              `or`,
	`prepend`:         `prepend`,
	`post`:            `post`,
	`>`:               `>`,
	`>>`:              `>>`,
}
