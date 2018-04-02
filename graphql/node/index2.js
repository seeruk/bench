const gqllanguage = require("graphql/language");
const gqllexer = require("graphql/language/lexer");
const hrtime = process.hrtime;

let hrstart = hrtime();

for (let i = 0; i < 5000000; i++) {
    const lexer = gqllanguage.createLexer(new gqllanguage.Source("query 0.001 foo { name model foo bar baz qux }"), {});

    do {
        let foo = lexer.token;
        lexer.advance();
    } while(lexer.token.kind !== gqllexer.TokenKind.EOF);
}

let hrend = hrtime(hrstart);

console.info("Execution time (hr): %ds %dms", hrend[0], hrend[1]/1000000);
