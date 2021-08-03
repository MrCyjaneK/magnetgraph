# Hello!

This is an example post that can be spread over the magnetgraph network.

But what actually is a magnetgraph? It's a simple standard for publishing content that will not change. Under the hood it's using [BitTorrent](https://en.wikipedia.org/wiki/BitTorrent) protocol, and a no external-data policy.

Therefore, if you would like to embed let's say this cute cat

![Look at this cute cat UwU](https://cataas.com/cat)

But wait? Where is the cat, in the source code you can clearly see that the cat image _is_ there, but the problem is: the cat image you have requested is provided under this URL: https://cataas.com/cat, and that's a problem, this image change on every request, so that would be bad if you would write article about a cat being cute, while the cat has suicidal thoughts?

Yup - that would be bad, But we have a solution for that, you can embed the cat image inside the article, for example in the `images/` directory.

![Look at this well-seeded picture of a cat](images/cat.jpg)

Now you can continue talking about the cat, while everybody is sure that nobody has ever changed this picture into something that can change. So yup - once you post something, it will be here forever, and will never ever disappear.

-------------------

If you would like to check the syntax, open file [syntax.md](syntax.md)