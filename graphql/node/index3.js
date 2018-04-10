const gqllanguage = require("graphql/language");
const gqllexer = require("graphql/language/lexer");
const hrtime = process.hrtime;


const query = `
	# Mutation for testing different token types.
	mutation {
		createPost(
			id: 1024
			title: "String Value"
			content: """Block string value isn't supported by all libs."""
			readTime: 2.742
		)
	}
`

const hrstart = hrtime();

for (let i = 0; i < 5000000; i++) {
  const parsed = gqllanguage.parse(query)
}

const hrend = hrtime(hrstart);

console.info("Execution time (hr): %ds %dms", hrend[0], hrend[1]/1000000);
