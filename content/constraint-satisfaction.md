+++
Categories = ["Computation", "Activation"]
bibfile = "ccnlab.json"
+++

**Constraint satisfaction** is one of the most important concepts for understanding the power of [[bidirectional connectivity]] in the [[neocortex]], and is central to the promise of the [[Axon]] approach, which is one of the few neural network models to incorporate extensive bidirectional connectivity in a way that tames its wild side while retaining its many benefits.

{id="figure_hopfield" style="height:30em"}
![The bidirectionally-connected network from Hopfield & Tank (1985) solving the traveling salesman problem (TSP). Cities are represented as rows (A-J) and the position of each city in a path is represented by columns. The network's solution in panel **(d)** is the city order DHIFGEAJCB (i.e., city D has the first position active, H is next, etc). The synaptic weights in the network encode how close each city is to the others, and there are inhibitory connections within the same column and row, so that each city is only represented once, and each position only has one city. An iterative _settling_ process updates the neuron activities (shown by the size of the square), across panels a-d. This process involves computing the gradients of each unit relative to the others, producing efficient search.](media/fig_hopfield_tank_85_tsp.png)

The constraint satisfaction problem (CSP) is defined as finding a set of values for N variables that satisfy a set of constraints defined over those variables ([[@Tsang14]]). A classic example is the N-queens problem, where you need to place N chess queens on a board such that no two queens should threaten each other. Another such problem is the _travelling salesman problem_ (TSP), analyzed by [[@^HopfieldTank85]] using a bidirectionally-connected Hopfield network, which involves finding a minimum distance route between _N_ cities ([[#figure_hopfield]]).

Thus, the CSP is essentially a problem of [[search]]ing over all possible states to find the one(s) that best fit the set of constraints imposed. As the number of states increases, the number of possible states explodes exponentially due to the [[curse of dimensionality]], 

A bidirectionally-connected neural network can implement this search process in a highly efficient manner, by integrating all of the constraints in a single gradient-based computational step over _dedicated-parallel_ representations, effectively performing a _stochastic gradient descent_ process over possible solution states. This is essentially the same strategy used in [[error-backpropagation]] learning, to search over the high-dimensional space of possible representations, as explained in [[search]]. Mathematically, it is effectively a process of [[error backpropagation#backpropagation to activations]].

Purely feedforward networks do not adapt their representations dynamically as they process the current set of inputs, and instead just generate a representation in one sweep, based on the learned weights. Thus, they are not optimizing these representations to find the most _satisfying_ way of interpreting the current situation. By contrast, the iterative back-and-forth interactions among bidirectionally-connected neurons ends up optimizing the active representation, which then provides the basis for subsequent learning.


