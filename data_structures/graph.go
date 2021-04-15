package main

import "fmt"

var VertexAlreadyExistsError = fmt.Errorf("Vertex already exists.")
var VertexDoesNotExistError = fmt.Errorf("Vertex does not exist.")
var EdgeAlreadyExistsError = fmt.Errorf("Edge already exists.")

// each vertex in the graph is uniquely identified by a key.
// any type satisfying the keyable interface can be stored
// in the graph
type Keyable interface {
	Key() string
}

type vertex struct {
	value Keyable
	edges map[string]*vertex
}

func (v *vertex) getKey() string {
	return v.value.Key()
}

func (v *vertex) getEdges() []*vertex {
	edges := make([]*vertex, 0, len(v.edges))
	for _, edge := range v.edges {
		edges = append(edges, edge)
	}

	return edges
}

func newVertex(value Keyable) *vertex {
	return &vertex{value: value, edges: map[string]*vertex{}}
}

type Graph struct {
	Vertices map[string]*vertex
	directed bool
}

// Constructor for undirected graph
func NewDirectedGraph() *Graph {
	return &Graph{Vertices: map[string]*vertex{}, directed: true}
}

// Constructor for directed graph
func NewUndirectedGraph() *Graph {
	return &Graph{Vertices: map[string]*vertex{}, directed: false}
}

// Method adds a new vertex storing the given value and identified
// by the values key to the graph
func (g *Graph) AddVertex(value Keyable) error {
	key := value.Key()
	if _, ok := g.Vertices[key]; ok {
		return VertexAlreadyExistsError
	}

	newVertex := newVertex(value)
	g.Vertices[key] = newVertex
	return nil
}

// Method adds a new edge to the graph between the vertices with the
// given keys
func (g *Graph) AddEdge(k1, k2 string) error {
	vertex1, ok := g.Vertices[k1]
	if !ok {
		return VertexDoesNotExistError
	}

	vertex2, ok := g.Vertices[k2]
	if !ok {
		return VertexDoesNotExistError
	}

	if _, ok := vertex1.edges[vertex2.value.Key()]; ok {
		return EdgeAlreadyExistsError
	}

	vertex1.edges[vertex2.value.Key()] = vertex2
	if !g.directed {
		vertex2.edges[vertex1.value.Key()] = vertex1
	}
	return nil
}

// Depth first search algorithm
func DFS(graph *Graph, startKey, targetKey string) bool {
	visited := map[string]bool{}
	stack := []*vertex{}

	// try to get the start vertex
	// if it doesn't exist, return false
	vertex, ok := graph.Vertices[startKey]
	if !ok {
		return false
	}

	// add first node to the stack and begin depth first search
	stack = append(stack, vertex)
	for len(stack) > 0 {
		// pop the current vertex from the stack
		currentVertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// if already visited, continue to next vertex
		if visited, ok := visited[currentVertex.value.Key()]; ok && visited {
			continue
		}
		// add current node to visited
		visited[currentVertex.value.Key()] = true

		// check if the vertex is the goal vertex
		if currentVertex.value.Key() == targetKey {
			return true
		}

		// otherwise get the edges of the current vertes
		// and add to the stack
		edges := currentVertex.getEdges()
		stack = append(stack, edges...)
	}

	return false
}

// Breadth first search algorithm
func BFS(graph *Graph, startKey, targetKey string) bool {
	visited := map[string]bool{}
	queue := []*vertex{}

	// try to get the start vertex
	// if it doesn't exist, return false
	vertex, ok := graph.Vertices[startKey]
	if !ok {
		return false
	}

	queue = append(queue, vertex)
	for len(queue) > 0 {
		// get the current vertex from the front of the queue
		currentVertex := queue[0]
		queue = queue[1:]

		// if already visited, continue to next vertex
		if visited, ok := visited[currentVertex.value.Key()]; ok && visited {
			continue
		}
		// add current node to visited
		visited[currentVertex.value.Key()] = true

		// check if the vertex is the goal vertex
		if currentVertex.value.Key() == targetKey {
			return true
		}

		// otherwise get the edges of the current vertes
		// and append to the queue
		edges := currentVertex.getEdges()
		queue = append(queue, edges...)
	}

	return false
}

// Person implements the Keyable interface
type Person struct {
	name string
	age  int
}

func (p *Person) Key() string {
	return p.name
}

func NewPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

func main() {
	michael := NewPerson("Michael", 23)
	ben := NewPerson("Ben", 22)
	john := NewPerson("John", 23)

	friendGraph := NewUndirectedGraph()

	friendGraph.AddVertex(michael)
	friendGraph.AddVertex(ben)
	friendGraph.AddVertex(john)

	friendGraph.AddEdge("Michael", "John")
	friendGraph.AddEdge("Michael", "Ben")

	fmt.Println(BFS(friendGraph, "John", "Ben"))
}
