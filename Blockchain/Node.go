package Blockchain

/*
	NewNode contains json information about the new node
*/
type NewNode struct {
	Url string `json:"new_node_url"`
}

/*
	Nodes is an array of addresses of each node
*/
type Nodes []string
