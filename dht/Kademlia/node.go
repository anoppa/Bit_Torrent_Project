package dht

import (
	"bytes"
	"math/big"
	"net"
	"strconv"
)

// NetworkNode representation of the Node in the network
type NetworkNode struct {
	//ID with be a binary nuber of 20 bytes
	ID []byte
	//IP direction of the node
	IP net.IP
	//port number
	port int
}

//most general representation of the dht Node(it would include other metadatas like las_seen_time)
type node struct {
	*NetworkNode
}

// NewNetworkNode constructor of NetworkNode
func NewNetworkNode(ip string, port string) *NetworkNode {
	p, _ := strconv.Atoi(port)
	return &NetworkNode{
		port: p,
		IP:   net.ParseIP(ip),
	}
}

func (n *node) CompactInfo() []byte {
	info := make([]byte, len(n.ID))
	copy(info, n.ID)
	info = append(info, n.IP...)
	info = append(info, uint8(n.port))
	return info
}

func newNodeFromCompactInfo(info []byte) *node {
	n := &node{}
	id := make([]byte, 20)
	copy(id, info[:20])
	ipB := make([]byte, 16)
	copy(ipB, info[20:36])
	ip := net.IP(ipB)
	portB := make([]byte, 2)
	copy(portB, info[36:])
	port := uint16(portB[0])<<8 + uint16(portB[1])
	n.ID = id
	n.IP = ip
	n.port = int(port)
	return n
}

// NewNode constructor of node
func newNode(netNode *NetworkNode) *node {
	n := &node{}
	n.NetworkNode = netNode
	return n
}

//compare two NetworkNodes
func equalsNodes(n1 *NetworkNode, n2 *NetworkNode, letIDNil bool) bool {
	if n1 == nil || n2 == nil {
		return false
	}
	if !letIDNil {
		if n1.ID == nil || n2.ID == nil {
			return false
		}
		if bytes.Compare(n1.ID, n2.ID) != 0 {
			return false
		}
	}
	if n1.port != n2.port {
		return false
	}
	if !n1.IP.Equal(n2.IP) {
		return false
	}
	return true
}

//struct to sort a list of nodes by the distance to a comparator
//sorted by x_or distance
type nodeList struct {
	//nodes in list
	Nodes []*NetworkNode

	//ID comparator
	Comparator []byte
}

//x_or distance of two nodes
func getDistance(id1 []byte, id2 []byte) *big.Int {
	id1_ := new(big.Int).SetBytes(id1)
	id2_ := new(big.Int).SetBytes(id2)
	dist := new(big.Int).Xor(id1_, id2_)
	return dist
}

//swap two Nodes in Nodelist
func (l *nodeList) Swap(i, j int) {
	l.Nodes[i], l.Nodes[j] = l.Nodes[j], l.Nodes[i]
}

//return true if node i is nearest to comparator
func (l *nodeList) Less(i, j int) bool {
	iDist := getDistance(l.Nodes[i].ID, l.Comparator)
	jDist := getDistance(l.Nodes[j].ID, l.Comparator)

	if iDist.Cmp(jDist) < 0 {
		return true
	}
	return false
}

//remove a node from a listNode
func (l *nodeList) RemoveNode(n *NetworkNode) {
	for i := 0; i < l.Len(); i++ {
		if bytes.Compare(l.Nodes[i].ID, n.ID) == 0 {
			l.Nodes = append(l.Nodes[:i], l.Nodes[i+1:]...)
			return
		}
	}
}

//Append to l each node in nodes list not in l
func (l *nodeList) AppendUniqueNetworkNodes(nodes []*NetworkNode) {
	for _, n := range nodes {
		exist := false
		for _, aux := range l.Nodes {
			if bytes.Compare(n.ID, aux.ID) == 0 {
				exist = true
				break
			}
		}
		if !exist {
			l.Nodes = append(l.Nodes, n)
		}
	}
}

// AppendUnique NetworkNodes of nodes that are not in l
func (l *nodeList) AppendUnique(nodes []*node) {
	for _, n := range nodes {
		exists := false
		for _, aux := range l.Nodes {
			if bytes.Equal(aux.ID, n.ID) {
				exists = true
				break
			}
		}
		if !exists {
			l.Nodes = append(l.Nodes, n.NetworkNode)
		}
	}
}

//Len return number of nodes in list
func (l *nodeList) Len() int {
	return len(l.Nodes)
}
