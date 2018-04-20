package docs

// Digest stores a 1 line summary of each builtins
var Digest map[string]string = map[string]string{
	`swivel-table`:    `Rotates a table by 90 degrees`,
	`get`:             `Makes a standard HTTP request and returns the result as a JSON object`,
	`try`:             `Handles errors inside a block of code`,
	`event`:           `Event driven programming for shell scripts`,
	`tout`:            `Print a string to the STDOUT and set it's data-type`,
	`>`:               `Writes STDIN to disk - overwriting contents if file already exists`,
	`read`:            `'read' a line of input from the user and store as a variable`,
	`if`:              `Conditional statement to execute different blocks of code depending on the result of the condition`,
	`>>`:              `Writes STDIN to disk - appending contents if file already exists`,
	`swivel-datatype`: `Converts tabulated data into a map of values for serialised data-types such as JSON and YAML`,
	`unset`:           `Deallocates an environmental variable (aliased to '!export')`,
	`g`:               `Glob pattern matching for file system objects (eg *.txt)`,
	`rx`:              `Regexp pattern matching for file system objects (eg '.*\.txt')`,
	`tread`:           `'read' a line of input from the user and store as a user defined typed variable`,
	`catch`:           `Handles the exception code raised by 'try' or 'trypipe'`,
	`append`:          `Add data to the end of an array`,
	`out`:             `'echo' a string to the STDOUT with a trailing new line character`,
	`pt`:              `Pipe telemetry. Writes data-types and bytes written`,
	`f`:               `Lists objects (eg files) in the current working directory`,
	`alter`:           `Change a value within a structured data-type and pass that change along the pipeline without altering the original source input`,
	`murex-docs`:      `Displays the man pages for _murex_ builtins`,
	`getfile`:         `Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.`,
	`trypipe`:         `Checks state of each function in a pipeline and exits block on error`,
	`prepend`:         `Add data to the start of an array`,
	`post`:            `HTTP POST request with a JSON-parsable return`,
	`brace-quote`:     `Write a string to the STDOUT without new line`,
	`err`:             `Print a line to the STDERR`,
	`ttyfd`:           `Returns the TTY device of the parent.`,
	`set`:             `Define a variable and set it's value`,
}
