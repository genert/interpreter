## EcmaScript 6 interpreter in Go

WIP

Limits:
- Supports only ASCII characters instead of the full Unicode range.

Comes with REPL (Read Eval Print Loop) and top down recursive descent parser ("Pratt parser"). 

Parser takes sources code as input and produce a data structure which represents the source code.
When parser builds up the data structure, the input is analyzed, checking that it conforms to the expected structure - a process called syntactic analysis.
