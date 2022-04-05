package kvalscanner

// Token represents a lexical token.
type Token int

// Tokens to scan for that satisfy the KVAL (Key Value Access Language) specification
const (
	ILLEGAL Token = iota // Spacial token, ILLEGAL: Token is an illegal character
	EOF                  // Spacial token, EOF: Token signals the end of input
	WS                   // Spacial token, WS: Token signals whitespace has been found

	LITERAL // LITERAL: String literal discovered

	BUCKEY // Other operator, BUCKKEY: >>>> Bucket to Key syntax for KVAL
	BUCBUC // Other operator, BUCBUC: >> Bucket to Bucket syntax for KVAL
	KEYVAL // Other operator, KEYVALL :: Key to Value syntax for KVAL
	ASSIGN // Other operator, ASSIGN: => Assignment operator for KVAL renames

	USCORE // Single character operator, USCORE: _ Return unknown Key or Value
	OPATT  // Single character operator, OPATT: { Open a regular expression pattern
	CPATT  // Single character operator, COATT: } Close a regular expression pattern

	INS // Keyword, INS: Insert capability of KVAL
	GET // Keyword, GET: Get capability of KVAL
	LIS // Keyword, LIS: LIS capability of KVAL
	DEL // Keyword, DEL: Delete capability of KVAL
	REN // Keyword, REN: Rename capability of KVAL

	REGEX // Regular expression, REGEX: {PATT} ANy regex pattern inside OPATT and CPATT
)

// ErrorLookup map is used to look up an error and provide a human readable response
var ErrorLookup = map[Token]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",        // Spacial token, EOF: Token signals the end of input
	WS:      "WHITESPACE", // Spacial token, WS: Token signals whitespace has been found

	LITERAL: "LITERAL", // LITERAL: String literal discovered

	BUCKEY: "BUCKET >>>> KEY",  // Other operator, BUCKKEY: >>>> Bucket to Key syntax for KVAL
	BUCBUC: "BUCKET >> BUCKET", // Other operator, BUCBUC: >> Bucket to Bucket syntax for KVAL
	KEYVAL: "KEY :: VALUE",     // Other operator, KEYVALL :: Key to Value syntax for KVAL
	ASSIGN: "ASSIGNMENT",       // Other operator, ASSIGN: => Assignment operator for KVAL renames

	USCORE: "UNDERSCORE",          // Single character operator, USCORE: _ Return unknown Key or Value
	OPATT:  "OPEN REGEX PATTERN",  // Single character operator, OPATT: { Open a regular expression pattern
	CPATT:  "CLOSE REGEX PATTERN", // Single character operator, COATT: } Close a regular expression pattern

	INS: "INSERT KEYWORD", // Keyword, INS: Insert capability of KVAL
	GET: "GET KEYWORD",    // Keyword, GET: Get capability of KVAL
	LIS: "LIS KEYWORD",    // Keyword, LIS: LIS capability of KVAL
	DEL: "DEL KEYWORD",    // Keyword, DEL: Delete capability of KVAL
	REN: "REN KEYWORD",    // Keyword, REN: Rename capability of KVAL

	REGEX: "REGEX PATTERN", // Regular expression, REGEX: {PATT} ANy regex pattern inside OPATT and CPATT
}

// KeywordMap values exported for KVAL Parser to verify keywords
// Lookup 'LIT' value in KeyWordMap and if found we have a KVAL key word,
// e.g. INS, GET, LIS, REN, DEL. If used correctly this map will help a parser
// take care of and case-sensitivity issues when users specify query strings.
var KeywordMap = map[string]int{
	"INS": 0x1, // INS: Insert keyword uppercase
	"ins": 0x1, // ins: Insert keyword lowercase
	"Ins": 0x1, // ins: Insert keyword mixed-case

	"GET": 0x2, // GET: Get keyword uppercase
	"get": 0x2, // get: Get keyword lowercase
	"Get": 0x2, // get: Get keyword mixed-case

	"LIS": 0x3, // LIS: List keyword uppercase
	"lis": 0x3, // lis: List keyword lowercase
	"Lis": 0x3, // lis: List keyword mixed-case

	"DEL": 0x4, // DEL: Delete keyword uppercase
	"del": 0x4, // del: Delete keyword lowercase
	"Del": 0x4, // del: Delete keyword mixed-case

	"REN": 0x5, // REN: Rename keyword uppercase
	"ren": 0x5, // ren: Rename keyword lowercase
	"Ren": 0x5, // ren: Rename keyword mixed-case
}
