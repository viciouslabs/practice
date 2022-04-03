# Disjoint Sets

### Terminology

- **Parent node**: the direct parent node of a vertex. For example, in Figure 5, the parent node of vertex 3 is 1, the parent node of vertex 2 is 0, and the parent node of vertex 9 is 9.

- **Root node**: a node without a parent node; it can be viewed as the parent node of itself. For example, in Figure 5, the root node of vertices 3 and 2 is 0. As for 0, it is its own root node and parent node. Likewise, the root node and parent node of vertex 9 is 9 itself. Sometimes the root node is referred to as the head node.

- **The find function** finds the root node of a given vertex. For example, in Figure 5, the output of the find function for vertex 3 is 0.

- **The union function** unions two vertices and makes their root nodes the same. In Figure 5, if we union vertex 4 and vertex 5, their root node will become the same, which means the union function will modify the root node of vertex 4 or vertex 5 to the same root node.


## Implementations

### Implementation with Quick Find:
In this case, the time complexity of the find function will be O(1)O(1). However, the union function will take more time with the time complexity of O(N)O(N).

- When initializing a union-find constructor, we need to create an array of size NN with the values equal to the corresponding array indices; this requires linear time.
- Each call to find will require O(1)O(1) time since we are just accessing an element of the array at the given index.
- Each call to union will require O(N)O(N) time because we need to traverse through the entire array and update the root vertices for all the vertices of the set that is going to be merged into another set.
- The connected operation takes O(1)O(1) time since it involves the two find calls and the equality check operation.

### Implementation with Quick Union:
Compared with the Quick Find implementation, the time complexity of the union function is better. Meanwhile, the find function will take more time in this case.

- The same as in the quick find implementation, when initializing a union-find constructor, we need to create an array of size NN with the values equal to the corresponding array indices; this requires linear time.
- For the find operation, in the worst-case scenario, we need to traverse every vertex to find the root for the input vertex. The maximum number of operations to get the root vertex would be no more than the tree's height, so it will take O(N)O(N) time.
- The union operation consists of two find operations which (only in the worst-case) will take O(N)O(N) time, and two constant time operations, including the equality check and updating the array value at a given index. Therefore, the union operation also costs O(N)O(N) in the worst-case.
- The connected operation also takes O(N)O(N) time in the worst-case since it involves two find calls.

### Implementation with Quick Union By Rank:
The find operation in Quick Union is O(N). This is converted to O(H) where H represents the height of the unified tree. The implementation is very similar to the quick union. However, in the union we choose the parent node, and the child node by comparing the height of the two node. With this decision we try and keep the height of the tree to be the same, or increase it by 1, in the case of two sets having the same height. This idea of balancing the tree at the time of union came about to be known as the Quick Union by Rank.

- For the union-find constructor, we need to create two arrays of size NN each.
- For the find operation, in the worst-case scenario, when we repeatedly union components of equal rank, the tree height will be at most \log(N) + 1log(N)+1, so the find operation requires O(\log N)O(logN) time.
- For the union and connected operations, we also need O(\log N)O(logN) time since these operations are dominated by the find operation.