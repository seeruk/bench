const gqllanguage = require("graphql/language");
const gqllexer = require("graphql/language/lexer");
const hrtime = process.hrtime;

let hrstart = hrtime();

for (let i = 0; i < 1000000; i++) {
    const lexer = gqllanguage.createLexer(new gqllanguage.Source("query foo { name model }"), {});

    do {
        let foo = lexer.token;
        lexer.advance();
    } while(lexer.token.kind !== gqllexer.TokenKind.EOF);
}

let hrend = hrtime(hrstart);

console.info("Execution time (hr): %ds %dms", hrend[0], hrend[1]/1000000);

hrstart = hrtime();

const lexer = gqllanguage.createLexer(new gqllanguage.Source("query foo { name model }"), {});

do {
    let foo = lexer.token;
    lexer.advance();
} while(lexer.token.kind !== gqllexer.TokenKind.EOF);

hrend = hrtime(hrstart);

console.info("Execution time (hr): %ds %dms", hrend[0], hrend[1]/1000000);