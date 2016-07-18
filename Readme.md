Go Regexp Parse-Tree viewer
===========================

Create dot-graph for golang regexp package regular expression parse-trees:

e.g.

```
$ anareg '^abc.*'
digraph G {
node0 [label="OpConcat: `\Aabc(?-s:.)*`"];
node1 [label="OpBeginText: `\A`"];
node0 -> node1
node2 [label="OpLiteral: `abc`"];
node0 -> node2
node3 [label="OpStar: `(?-s:.)*`"];
node0 -> node3
node4 [label="OpAnyCharNotNL: `(?-s:.)`"];
node3 -> node4
}
```

use with dot-tool (graphviz) to convert graph into pdf or svg format.
