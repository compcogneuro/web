+++
Name = "OReilly (2026) Search"
Title = "Everything is Search, and Gradient Descent is the only Efficient Search Mechanism"
Authors = "Randall C. O'Reilly"
Affiliations = "Astera Institute; Department of Psychology<br>Center for Neuroscience, University of California Davis<br>correspondence: oreilly@ucdavis.edu"
Abstract = "."
Date = "2026-06-07"
Version = "1"
Categories = ["Papers"]
bibfile = "ccnlab.json"
+++

Every cognitively-relevant type of computation can be cast as form of search:

* _Problem solving_ can be defined as search through _problem space_, as in many classical symbolic artificial intelligence (AI) systems ([[@NewellSimon72]]; [[@Newell90]]).

* _Planning_ is search through _action space_ to accomplish a desired outcome, which is the focus of many approaches in model-based reinforcement learning.

* _Learning_ is search through _representation space_, to find the best basis space for representing inputs, that supports the desired computational processes and behavioral outputs.

* _Inference_ is search through learned representations to find the best _interpretation_ of a given input stimulus relative to the stable semantics of the environment.

* _Evolution_ is search through _phenotype space_, e.g., as in genetic algorithms.

This is not an empty tautology, because it provides a unified understanding about the central problem faced in all of these domains: Each of these search spaces is plagued by the curse of dimensionality, and effective solutions for real-world, large-scale problems must somehow tackle the exponential explosion problem. 

Indeed, the field of computational complexity provides a unified framework based fundamentally on the process of search, that formally quantifies the computational cost of different types of problems through the use of universal Turing machines ([[@AroraBarak09]]). The core distinction here is in terms of _polynomial_ (P) versus exponential time costs, where polynomial means that the compute time is _on the order of_ (expressed as $O(f(n))$) a _constant_ exponential factor of the dimensionality of the problem space (e.g., $O(n^2)$ where the polynomial factor is 2). By contrast, exponential time means that $n$ is instead in the exponent, e.g., $O(2^n)$, which is where the exponential explosion problem occurs as $n$ increases, making these problems _intractable_.

Computational complexity theory defines an equivalence class across a wide range of problems, known as _NP complete_, where NP stands for _nondeterministic polynomial_ time. The terminology is not particularly intuitive, but the NP complete class defines a large set of _decision problems_ that require _exhaustive, enumerative_ (i.e, _brute force_) search over an exponentially-large space of possible solutions, where each decision evaluates a binary (yes or no) function, which itself runs in polynomial time (which accounts for the _P_ in _NP_). 

Formally, the NP class is defined in terms of a _nondeterministic Turing machine_ (which is where the _N_ in _NP_ comes from) which effectively performs a brute-force search process over a space of deterministic Turing machines, each of which runs its evaluation function in polynomial time. The more general _NP hard_ designation is not restricted to binary decision problems -- the _complete_ equivalence class only applies to such problems that can all be reduced to being effectively notational variants of each other. Most cognitively relevant computations such as Bayesian inference and optimal decision making are at least in the NP-hard class ([[@vanRooijWrightWareham12]]).

The framing of computational problems as search provides a crucial, intuitive insight into why some problems are exponentially costly, and others are not. In effect, if there is no efficient way to narrow the search space, then there is no other alternative to performing the brute force search through all possible states. If there _were_ an efficient, systematic way to narrow the search process, then it could potentially move from the NP complete class to the tractable P class.

An illustrative case is the classic _travelling salesman problem_ (TSP), which involves finding the minimal distance route between a set of $n$ randomly-distributed cities, which is in the NP complete class. However, TSP can be solved in polynomial time _if_ there is an underlying metric space (e.g., Euclidean distances) and the optimality constraint is relaxed, so that a solution only need be within some factor of the shortest distance of a route between the given cities ([[@Arora98]]). Interestingly, either of these constraints alone is insufficient. The resulting algorithm indeed involves an efficient way of narrowing the search space by dividing it into different subregions and applying a novel search constraint on those regions.

{id="figure_search" style="height:20em"}
![The two main computational frameworks for understanding the computational complexity of search as a general computational process. **A** In a Turing Machine, TM, there is a single focus of computational processing at each point in time, which can involve one of a small set of finite operations (read / write from a long-term "tape" storage, or operations on elements in active memory). If there is no way to systematically narrow the search space, it must serially enumerate all relevant combinations to find the best solution, resulting in combinatorially exponential (i.e., NP) worst-case behavior (e.g., the target ends up being the last combination considered, like finding a needle in a haystack). **B** In a dedicated-weight-gradient-search (DWGS) algorithm like a neural network, each unit is dedicated to a specific feature value, and synaptic weights capture all interactions. One update iteration involves gradient computations across all such units, which effectively incorporates _all_ of the different weighted contstraints simultaneously, in choosing the next direction to search. The graded, smooth nature of the search space is essential for these incremental steps to provide an efficient narrowing of the search space. Under such conditions, a tractably small number of iterations will generally produce reasonable search results.](media/fig_search_turing_vs_nnet.png)

In this context, we can understand the _stochastic gradient descent_ procedure in neural network models (e.g., [[@RumelhartHintonWilliams86]]) as a uniquely powerful way of searching exponentially-large combinatorial feature spaces ([[#figure_search]]), because the gradient-descent process provides a highly efficient way of shaping the direction of the search process toward more optimal solutions. As in the TSP problem, the success of this gradient descent procedure depends on operating within a system and environment that has some kind of underlying _metric_ properties within a space that is fundamentally _smooth_ enough, such that incremental graded steps are able to accumulate over iterations, and move ultimately toward a search solution. Furthermore, this process can only produce _local minima_ in the relevant optimization metric, i.e., approximately-optimal solutions, not the absolute global optimal solution.

The success of extremely large (billions of parameters) neural-network based artificial-intelligence (AI) models depends entirely on the efficacy of this gradient-descent process. The importance of this efficient search process can also be seen in the negative case, in terms of the inability to scale up earlier symbolic AI models, for which a smooth metric gradient is undefined: symbols and the propositional structures in which they are used have discrete, "sharp" edges and many discontinuities. Furthermore, although the variational inference procedure enables gradient descent to be used on probabilistic (e.g., Bayesian) models ([[@BleiKucukelbirMcAuliffe17]]), there are still limitations in such models that prevent the kind of deep scaling possible in simpler, more linear neural network models.

It is also critical to emphasize the _no free lunch_ proscription against a generally optimal learning algorithm ([[@VapnikChervonenkis71]]; [[@GemanGeman84]]): gradient descent only works when gradients are actually useful for search, which represents a kind of _bias_ in the algorithm, that can easily be violated and thus result in very bad performance. Empirically, it is nevertheless evident that gradient descent in neural networks is effective for a wide range of relevant problems, and recent work is making progress analyzing why this is so ([[@LeePanageasPiliourasEtAl19]]; [[@DauphinPascanuGulcehreEtAl14]]).

In the context of a neural network, the gradient-descent search process can operate at two different levels, which correspond to different effective timescales of search:

* Iterative gradient descent over synaptic weights (i.e., learning), as in the error backpropagation algorithm, which configures these weights to shape the internal representations of the network over the course of many different learning trials, to optimize the learning objective function. In modern AI models, this is often the ability to predict the next element in a sequence.

* Iterative gradient descent over _activation states_ within an individual learning trial, which corresponds at least in part to various types of search listed at the outset, including inference, planning, and problem-solving. This can be performed using backpropagation to activation states, or via a _constraint satisfaction_ settling process as in the original _Hopfield_ network or _Boltzmann Machine_ ([[@Hopfield82]]; [[@HopfieldTank85]]; [[@AckleyHintonSejnowski85]]).

    This more rapid time scale of search process is not typically used in current large AI models, which are based on strictly feedforward architectures that do not support this form of search. Instead, this functionality is captured in part by the attention mechanism in transformers ([[@VaswaniShazeerParmarEtAl17]]) and cascading effects across multiple deep layers ([[@McClellandHillRudolphEtAl20]]). Nevertheless, it is likely that this inner-loop of constraint satisfaction search is critical for human intelligence, and, interestingly, many theories of consciousness involve a central element of this kind of recurrent, interactive processing (e.g., [[@Lamme06]]; [[@TononiKoch15]]).

In summary, the main point of this paper is simply to provide the broad conceptual framework articulated above, with the empirically-backed conclusion that stochastic gradient descent in weighted neural-network architectures may represent a uniquely powerful form of search, which can be seen as the central process involved in all major cognitive functions. This perspective also highlights the potential importance of the shorter time scale constraint-satisfaction search process to support inference, planning and problem-solving, which is missing in current large-scale AI models.

In the remaining sections, current research on gradient descent in learning and constraint satisfaction processing is reviewed, including new experimental results that point to important benefits of biologically-based neural network properties, including discrete spiking, surround inhibition, and the restriction to positive-only weights and activation states.

# Acknowledgments

This work has benefited significantly from input from Alex Petrov and other members of the analogy discussion group, hosted by Jonathan Cohen at Princeton.

