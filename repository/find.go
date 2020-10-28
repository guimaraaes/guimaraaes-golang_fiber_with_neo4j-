package repository

import "github.com/guimaraaes/golang_fiber_with_neo4j/repository/utils"

func FindR(model interface{}, info map[string]interface{}) ([]string, string) {
	node, properties := utils.Label_Properties_Node(model, info)
	// fmt.Println(node)
	// fmt.Println(properties)
	var query string
	var m map[string]interface{} = nil
	if info == nil {
		query = "MATCH (m) WHERE $node in labels(m) " +
			"CALL apoc.path.subgraphAll(m, {maxLevel:0}) " +
			"YIELD nodes WITH [node in nodes | node {.*, label:labels(node)[0]}] as nodes " +
			"RETURN apoc.convert.toJson(nodes[0])"
		m = map[string]interface{}{"node": node}
	} else {
		query = "MATCH (m {" + properties + "  }) WHERE $node in labels(m) " +
			"CALL apoc.path.subgraphAll(m, {maxLevel:0}) " +
			"YIELD nodes WITH [node in nodes | node {.*, label:labels(node)[0]}] as nodes " +
			"RETURN apoc.convert.toJson(nodes[0])"
		m = map[string]interface{}{"node": node}
	}
	c, err := utils.QueryCall(query, m)
	return c, err
}