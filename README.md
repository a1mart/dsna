# Data Structures and Algorithms

| **Name**                      | **Type**      | **Avg Time Complexity**                                | **Space Complexity**   | **Use Cases**                                   |
| ----------------------------- | ------------- | ------------------------------------------------------ | ---------------------- | ----------------------------------------------- |
| **AVL Tree**                  | Balanced BST  | Search: O(log n), Insert: O(log n), Delete: O(log n)   | O(n)                   | Self-balancing tree for ordered data            |
| **Multiset (Bag)**            | Container     | Insert: O(1) avg (hash), O(log n) (tree), Lookup: O(1) | O(n)                   | Duplicates allowed, frequency count             |
| **Bloom Filter**              | Probabilistic | Insert: O(k), Lookup: O(k)                             | O(m)                   | Membership test with false positives            |
| **B+ Tree**                   | Search Tree   | Search: O(log n), Insert: O(log n), Delete: O(log n)   | O(n)                   | Database indexing                               |
| **B Tree**                    | Search Tree   | Search: O(log n), Insert: O(log n), Delete: O(log n)   | O(n)                   | Disk-based indexing                             |
| **Binary Search Tree**        | Tree          | Search: O(log n)*, Insert: O(log n)*                   | O(n)                   | Ordered data (*unbalanced can degrade to O(n)*) |
| **Circular Linked List**      | Linked List   | Insert/Delete: O(1), Search: O(n)                      | O(n)                   | Round-robin scheduling                          |
| **Count-Min Sketch**          | Probabilistic | Update: O(k), Query: O(k)                              | O(m)                   | Streaming frequency estimation                  |
| **Deque**                     | Queue         | Push/Pop: O(1) both ends                               | O(n)                   | Double-ended operations                         |
| **Dijkstra's Algorithm**      | Graph Algo    | O((V + E) log V) with priority queue                   | O(V + E)               | Shortest path                                   |
| **Disjoint Set (Union-Find)** | DS            | Find: O(α(n)), Union: O(α(n))                          | O(n)                   | Dynamic connectivity                            |
| **Doubly Linked List**        | Linked List   | Insert/Delete: O(1), Search: O(n)                      | O(n)                   | Bidirectional navigation                        |
| **Floyd-Warshall Algorithm**  | Graph Algo    | O(V³)                                                  | O(V²)                  | All-pairs shortest path                         |
| **Graph**                     | Abstract DS   | Depends on representation (Adj list vs Adj matrix)     | O(V+E) or O(V²)        | Network, relationships                          |
| **Hash Table**                | DS            | Insert/Search: O(1) avg                                | O(n)                   | Fast lookup                                     |
| **Heap**                      | DS            | Insert: O(log n), Extract-min/max: O(log n)            | O(n)                   | Priority queues                                 |
| **Stack**                     | DS            | Push/Pop: O(1)                                         | O(n)                   | LIFO operations                                 |
| **Queue**                     | DS            | Enqueue/Dequeue: O(1)                                  | O(n)                   | FIFO operations                                 |
| **SyncMap**                   | Concurrency   | O(1) avg                                               | O(n)                   | Thread-safe maps in Go                          |
| **SyncPool**                  | Concurrency   | O(1) avg                                               | O(n)                   | Object pooling in Go                            |
| **Trie**                      | Tree          | Insert/Search: O(L) (L = length of word)               | O(ALPHABET\_SIZE \* n) | Prefix search                                   |
| **Fenwick Tree (BIT)**        | DS            | Update: O(log n), Query: O(log n)                      | O(n)                   | Range sum queries                               |
| **Kd Tree**                   | Tree          | Search: O(log n) avg                                   | O(n)                   | Nearest neighbor search                         |
| **Linked List**               | DS            | Insert/Delete: O(1), Search: O(n)                      | O(n)                   | Dynamic memory allocation                       |
| **Segment Tree**              | DS            | Update: O(log n), Query: O(log n)                      | O(n)                   | Range queries                                   |
| **Set**                       | DS            | Insert/Search: O(1) avg                                | O(n)                   | Unique elements                                 |
| **Suffix Array**              | String Algo   | O(n log n) for construction                            | O(n)                   | Pattern matching                                |
| **Suffix Tree**               | String Algo   | O(n) construction with Ukkonen’s algorithm             | O(n)                   | Fast substring search                           |
| **Ternary Search Tree**       | Tree          | Search: O(L), Insert: O(L)                             | O(n)                   | Memory-efficient trie                           |


## Legend
- **n** = Number of elements
- **V** = Number of vertices (for graphs)
- **E** = Number of edges (for graphs)
- **k** = Number of hash functions (for Bloom filter & Count-Min Sketch)