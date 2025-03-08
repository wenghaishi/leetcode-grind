const airports = "PHX BKK OKC JFK LAX MEX EZE HEL LOS LAP LIM".split(" ");

const routes = [
  ["PHX", "LAX"],
  ["PHX", "JFK"],
  ["JFK", "OKC"],
  ["JFK", "HEL"],
  ["JFK", "LOS"],
  ["MEX", "LAX"],
  ["MEX", "BKK"],
  ["MEX", "LIM"],
  ["MEX", "EZE"],
  ["LIM", "BKK"],
];

const myList = new Map();

function addNode(ap) {
  myList.set(ap, []);
}

// Add edge, undirected
function addEdge(origin, destination) {
  myList.get(origin).push(destination);
  myList.get(destination).push(origin);
}

airports.forEach((ap) => addNode(ap));
routes.forEach((r) => addEdge(...r));

function bfs(start) {
  const queue = [start];
  const visited = new Set();

  while (queue.length > 0) {
    const airport = queue.shift();
    const edges = myList.get(airport);
    console.log(airport);

    if (edges.includes("BKK")) {
      console.log("Found bangkok");
    } else {
      edges.forEach((edge) => {
        if (!visited.has(edge)) {
          queue.push(edge);
          visited.add(edge);
        }
      });
    }
  }
}

function dfs(start, visited = new Set()) {
  console.log(start);
  visited.add(start);
  const edges = myList.get(start);

  for (const edge of edges) {
    if (edge === "BKK") {
      console.log("found BKK");
    } else if (!visited.has(edge)) {
      dfs(edge, visited);
    }
  }
}

// bfs("PHX");
dfs("PHX");
