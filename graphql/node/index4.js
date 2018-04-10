const gqllanguage = require("graphql/language");
const gqllexer = require("graphql/language/lexer");

const query = `foo(message: """
    Hello
  World
   What
Is
  The
    Indentation
               Here?
""")`

const lexer = gqllanguage.createLexer(new gqllanguage.Source(query), {});

console.log("Query: '" + query + "'")

do {
  let foo = lexer.token;
  lexer.advance();
  console.log(foo);
} while(lexer.token.kind !== gqllexer.TokenKind.EOF);
