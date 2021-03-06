const gqllanguage = require("graphql/language");
const gqllexer = require("graphql/language/lexer");

const query = "query \"\\u4e16\"  \"a\\r\\nb\" foo { name model }"
const lexer = gqllanguage.createLexer(new gqllanguage.Source(query), {});

console.log("Query: '" + query + "'")

do {
  let foo = lexer.token;
  lexer.advance();
  console.log(foo);
} while(lexer.token.kind !== gqllexer.TokenKind.EOF);
