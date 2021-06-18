%{
package yacc_parseJson

func setScannerData(yylex interface{}, data interface{}) {
	yylex.(*Scanner).data = data
}

%}

%union {
    empty   struct{}
    b       bool
    s       string
    i       interface{}
    a       []interface{}
    m       map[string]interface{}
}

// "{" open curly brackets; LC
// "}" close curly brackets; RC
// "[" open square brackets; LS
// "]" close square brackets; RS

%token<empty> COMMA COLON LS RS LC RC
%token<empty> TRUE FALSE NULL
%token<s> STRING
%token<i> NUMBER

%type<b> true false
%type<i> value number null
%type<s> str
%type<m> object object_data
%type<a> array array_data

%%

json:
  value
  {
    setScannerData(yylex, $1)
  }

value:
  str
  {
    $$ = $1
  }
| number
  {
    $$ = $1
  }
| true
  {
    $$ = $1
  }
| false
  {
    $$ = $1
  }
| null
  {
    $$ = $1
  }
| object
  {
    $$ = $1
  }
| array
  {
    $$ = $1
  }

object:
  LC object_data RC
  {
    $$ = $2
  }
| LC RC
  {
    // no-op
  }

object_data:
  STRING COLON value // "key": "value"
  {
    $$ = map[string]interface{}{$1:$3};
  }
| object_data COMMA STRING COLON value
  {
    $$ = $1
    $$[$3] = $5
  }

array:
  LS array_data RS
  {
    $$ = $2
  }
| LS RS
  {
    // no-op
  }

array_data:
  value
  {
    $$ = []interface{}{$1}
  }
| array_data COMMA value
  {
    $$ = append($1, $3)
  }

str:
  STRING
  {
    $$ = $1
  }

number:
  NUMBER
  {
    $$ = $1
  }

true:
  TRUE
  {
    $$ = true
  }

false:
  FALSE
  {
    $$ = false
  }

null:
  NULL
  {
    $$ = nil
  }