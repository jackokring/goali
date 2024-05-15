# PLAN

Seems a good use of graphics is a simple game. Some kind of slow adventure? Ah, it's all gone text grammar logic again.

## Tile Based

Definitely the easiest. Subtile things can be tacked on later.

## AI

Use a constrained vocabulary encoding and a small coefficient LLM for game conversation. Perhaps a larger online model fallback with re-expression reduction of the model to the smaller encoding. Perhaps "Say `expression` differently without using the words `list`." GAN prompt automation kind of things. Limiting the words to named game nouns/adjectives and possible verbs/adverbs understood with a spattering of, auxiliary "pre-proto-logical", grammar words.

