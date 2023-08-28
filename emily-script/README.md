# emily-script

This sub-folder includes my experiments with creating an own scripting language `emily-script` (`em`). I took [«Writing an interpreter in go»](https://interpreterbook.com/) by Thorsten Ball as a guide to understand the underlying logic of a scripting language and its challenges.

## Structure

```mermaid
flowchart
    subgraph Interpreter
        subgraph Input
            code[Source Code]
        end
        subgraph REPL
            input[User Input REPL]
        end
        subgraph Lexer
            parse[Parsing]
            lex[Lexer]
        end
        subgraph AST
            ast[Abstract Syntax Tree AST]
        end
        subgraph Evaluation
            eval[Evaluation]
        end
    end

    code --> |tokenizing| parse --> lex --> |parsing| ast --> eval
    input --> |tokenizing| parse
```

### Token

Defining tokens for the `emily-script` language. These are later used for parsing and lexing.

Some of the most important tokens:
- ILLEGAL: undefined tokens 
- EOF: end of file
- IDENT: identifier
- INT: integers
- ASSIGN: equals (=)
- COMMA: ,
- SEMICOLON: ;
- LPAREN: (
- RPAREN: )

### Lexer

A Lexer parses the input code char by char and identifies the associated token.

### Repl

REPL (read eval print loop) reads input, evaluates it and prints the result/tokens in a loop.

### Ast

AST (abstract syntax tree) is a data structure to represent the input code in a logical tree.

**Differences `statements` vs `expression`:**
> `expressions` produce values but `statements` not

```mermaid
flowchart LR
    subgraph AST 
        S1[Statements]
    end
    subgraph AST:Statement
        S2[Statement]
        N[Statement Name]
        V[Statement Value]
    end
    I[AST:Identifier]
    E[AST:Expression]
    S1 --> S2 --> N --> I
    S2 --> V --> E
```

Examples:

```em
// doesnt produce values -> statement
let test = 10;
return x

// produce values -> expression
5;
add(1, 2);
```
