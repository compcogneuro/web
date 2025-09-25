+++
Categories = ["Computation", "Activation", "Learning"]
bibfile = "ccnlab.json"
+++

**Search** is perhaps the single most general unifying concept in all of [[computation]].

* _Problem solving_ can be defined as search through _problem space_, as in many classical symbolic [[artificial intelligence]] (AI) systems defined it ([[@NewellSimon72]]; [[Newell90]]).

* _Planning_ is search through _action space_ to accomplish a desired outcome, which is the focus of many approaches in [[reinforcement learning#model-based]] reinforcement learning.

* _Learning_ is search through _representation space_, to find the best [[linear algebra|basis]] for representing inputs, that supports the desired computational processes and behavioral outputs.

* _Inference_ is search through learned representations to find the best _interpretation_ of a given input stimulus relative to the stable semantics of the environment.

* _Evolution_ is search through _phenotype space_, e.g., as in the widely-used [[genetic algorithm]].

This is not an empty tautology, because it provides a unified understanding about a central problem faced in all of these domains:

> Each of these spaces is plauged by the [[curse of dimensionality]], and effective solutions for real-world, large-scale problems must somehow tackle the exponential explosion problem.

Indeed, the field of [computational complexity theory](https://en.wikipedia.org/wiki/Computational_complexity_theory) (wikipedia) provides a unified framework based on an equivalence class across a wide range of problems, known as _NP complete_ that are not solvable in polynomial time. The _NP_ means _nondeterministic polynomial_ time, which means that the solution can be verified quickly (in polynomial time) but actually generating the solution takes much longer, e.g., exponential time. Any problem in this NP-complete set is considered _intractable_ and must therefore be avoided, and yet, most cognitively-relevant computations such as Bayesian inference and optimal decision making are NP-hard ([[@vanRooijWrightWareham12]]).

## Enumerative search is cursed

If the algorithm used in a search problem domain involves any element of _enumerative_ search through states (i.e., enumerating specific combinations of relevant values along the different dimensions), then that algorithm will likely suffer from the curse of dimensionality, exponentially-quickly becoming intractable as the state space increases.

For example, the simplest _brute force_ algorithm for any such problem involves a giant set of nested `for` loops through all values along each dimension, enumerating each possible combination, and then evaluating some function computed on this state vector. As anyone who has experienced such nested `for` loops will recognize, these bog down very quickly as the number of such levels is increased. 

Although it is intuitively appealing to refer to this as a _serial_ search process, the problems persist even if you use a parallel implementation wrapped around this enumerative process. Any amount of parallelism will just divide the exponentially-growing factor by a fixed constant on the order of the parallel hardware involved. Even if this is the brain with 100 billion neurons, that factor will quickly pale in comparison to the combinatorial explosion of something with only 100 different dimensions (e.g., $2^{100} / 1.0e11 \approx 1.0e19$, for 100 binary-valued dimensions).

Furthermore, it is always possible to implement a parallel algorithm on a serial computer, and indeed this is often done. Thus, the issues below about parallel vs. serial refer to the underlying _algorithmic_ structure of the computation, not how it might actually be implemented in hardware. Again, these hardware-level considerations only contribute constant factors that quickly become irrelevant in the face of high-dimensional combinatorial explosion.

## Polynomial-time (P) search in neural networks

The only plausible way to search in high-dimensional spaces is to use an algorithm that operates directly at the level of the dimensional values themselves, and not on their combinatorial enumerations, while still somehow capturing enough of the combinatorial interactions across the dimensions that are relevant to the problem at hand. Such an algorithm would have a computational cost that is _polynomial_ on the order of the number of dimensions.

For example, in a neural network with each unit representing a value on a given dimension, the computational cost is a function of the number of synaptic weights interconnecting the neurons, which is $n^2$ in the worst case, making the computational cost polynomial with an exponent of 2. If this neural network is implemented in dedicated neural hardware that operates in parallel at the level of the weights, as in the brain, then the computational time remains effectively constant, which is what we clearly observe in comparing tiny mouse brains with giant human brains.

{id="figure_search" style="height:20em"}
![The two main computational frameworks for understanding the computational complexity of search as a general computational process. In a Turing Machine as illustrated on the left, there is a single focus of computational processing at each point in time, which can involve one of a small set of finite operations (read / write from a long-term "tape" storage, or operations on elements in active memory). Thus it must serially enumerate all relevant combinations to find the best solution, resulting in NP worst-case behavior (e.g., the target ends up being the last combination considered, like finding a needle in a haystack). In a dedicated-parallel algorithm like a neural network (on the right), each unit is dedicated to a specific value, and synaptic weights capture all interactions. One update iteration involves gradient computations across all such units, which effectively searches the entire representational space in parallel (algorithmically). A fixed number of iterations will produce reasonable search results, resulting in polynomial computational complexity.](media/fig_search_turing_vs_nnet.png)

A neural network captures the interactions among the different values via the synaptic weights interconnecting the neurons ([[#figure_search]]). [[Bidirectional connectivity|Bidirectionally-connected]] networks can also leverage the compounding effects of multiple iterations of gradient-based updates over time (e.g., over 200 ms for the [[theta rhythm]]), to compute and communicate interactions among distributed combinations of variables. Thus, instead of enumerating all of the different combinations across dimensions, these synaptic weights (and iterative interactions) have to capture everything that is relevant. It may seem implausible that this is sufficient relative to full exhaustive enumeration, but the remarkable success of [[abstract neural network]] algorithms demonstrates that this is indeed a highly efficient, pragmatic way of managing the tradeoff between computational complexity and computational power.

A major reason why neural networks work as well as they do is that these synaptic weights can be trained to encode the relevant combinatorial interactions, using algorithms that are themselves polynomially efficient. Indeed, the paradigmatic example of efficient search is the _stochastic gradient descent_ process used in [[error backpropagation]] learning, which computes the _first order_ partial derivatives (i.e., _gradients_) of the synaptic weights relative to an overall error function (the _objective_ function). The computational cost of backpropagation is linear on the order of the number of synaptic weights, and yet has been shown to be highly effective in optimizing neural networks to accomplish a wide array of different tasks, including of course the [[large language models]] (LLMs) powering the widely-used ChatGPT like AI models.

Interestingly, many people considered using something like the backpropagation algorithm over the years (it is really just a very standard gradient-based optimization technique based on elementary calculus), but didn't even bother because they assumed that this simple gradient-based algorithm would just produce horribly sub-optimal solutions (i.e., _local minima_). While there is no guarantee that gradient descent will produce optimal solutions, in practice, with the appropriate heuristic techniques (e.g., the _AdaMax_ mechanism; [[@KingmaBa14]]) and network configurations, it certainly does well enough to be of immense practical importance.

This heuristic algorithmic approach that eschews optimality constraints in favor of "good enough" solutions (i.e., _satisficing_; [[@Simon56]]]) contrasts with the computational complexity approach, which is traditionally framed in terms of optimal solutions and worst-case scenarios involving the run-times of [[Turing machine]] programs. Thus, the success of error backpropagation does not undermine the NP-complete framework and its central hypothesis about the fundamental intractability of that class of algorithms, but it does suggest that more "relaxed" ways of analyzing the relative costs and capabilities of various algorithms would be useful.

To attempt to more precisely label the broader category of neural-network like algorithms, we refer to them as _dedicated-parallel, gradient-based, weighted, and adaptive_ algorithms:

* Dedicated representational elements capture the entire relevant state being searched, instead of flexibly operating on individual state components that can be arbitrarily combined during the computation. This supports computation across the _entire state_ in one algorithmically-parallel step, i.e., **dedicated-parallel** computation.

* Gradients (partial derivatives) are used to efficiently search the entire dedicated-parallel state in one algorithmic step. A partial derivative automatically computes the effects of all the other state elements with respect to changes taken on one local element of the state (as in the [[error backpropagation]] algorithm), and it is typically recursively computable, supporting local parallel computation, as demonstrated by the [[GeneRec]] algorithm. The recent explosion in neural network algorithm development has all been powered by the computational efficacy of the _autograd_ framework.

    Iterative updating via the gradients proceeds for a finite number of iterations, with the final result being sufficiently optimized within some reasonable threshold. This forecloses the worst-case search scenario of exponential iterations.

* Graded weighting factors are used to capture relationships across the representational elements, which are thus compatible with the dedicated-parallel, gradient-based computation. This places a strong constraint on the nature of the combinatorics that the system can represent, and is what truncates the combinatorial explosion to very low order (e.g., $n^2$).

    Much of the exploratory engineering in designing a neural network model involves configuring the topology of these weighted connections to capture the relevant interactions, which represents a kind of outer-loop of hyperparameter searching, so the computational cost of this should not be overlooked. From a biological perspective, this is what evolution does (see below).

* These weighting factors are adapted (using efficient gradient-based mechanisms) to optimize the dedicated representational space so that it properly captures the relevant relationships across the state variables (i.e., learning). Once learned, the efficient propagation of information through the dedicated-parallel network automatically expresses the relevant state interactions, to support functions such as inference, problem-solving, and planning.

This attempt to generalize the properties of [[abstract neural network]] models is similar to the original _parallel distributed processing (PDP)_ terminology of [[@RumelhartMcClelland86]]. It is unclear if it meaningfully expands the class of useful such algorithms beyond existing neural network algorithms, but  perhaps it is helpful in more clearly delineating the algorithmic properties of these networks relative to other systems.

<!--- TODO: find better lit on this stuff. Initial google search not that promising: https://scholar.google.com/scholar?hl=en&as_sdt=0%2C5&q=%22backpropagation%22+%22computational+complexity%22+%22np%22+%22nondeterministic+polynomial%22&btnG= -->

## Learning as search

The ability of neural network models to implement efficient polynomial-time search through representational space represents perhaps the biggest divide between them and classical symbolic AI models. The discrete symbolic nature of classical AI frameworks, where the models are essentially [[Turing machine]]s with various AI-based affordances, is incompatible with gradient-based, synaptically-mediated learning mechanisms. Therefore, any attempt to learn in such models ends up requiring enumerative combinatorial search, which fails in the face of combinatorial explosion. Again, this remains true even if parallel hardware can be used, as it is a property of the underlying algorithms.

There is now a growing literature illuminating why stochastic gradient descent works as implausibly well as it does (e.g., [[@NakkiranKaplunBansalEtAl21]]; [[@Shwartz-ZivTishby17]]; [[@VidalBrunaGiryesEtAl17]]).
The _stochastic_ nature of the stochastic gradient descent process, which is critical to its success and consistent with the broader importance of noise in search processes, is an interesting emergent process. Instead of being something that is externally injected, it derives from the use of small subsamples of the total problem space to compute the gradients. If the full space is used to compute the gradients, this ends up getting "stymied" by all the constraints imposed across this entire space, and learning typically does not succeed.

## Constraint satisfaction as search

[[Constraint satisfaction]] provides another high-level unified way of understanding a wide range of computational processes, and this generality is not accidental from the present search-based perspective:  Constraint satisfaction is effectively the search over possible states that solve or optimize (approximately!) a set of constraints. Many problems can be formulated in this way, including planning and problem solving, so constraint satisfaction provides a more general way of understanding the essential computations in these domains.

The stochastic, gradient-based approach to constraint satisfaction can efficiently scale to high-dimensional state spaces because it has polynomial time complexity, and generally does not suffer from local minima problems ([[@HoosTsang06]]). One class of such algorithms is in fact a neural network, going back to the pioneering approach of [[@HopfieldTank85]], who showed that iterative updating of a bidirectionally-connected Hopfield network could provide good solutions to the travelling salesman problem, which is a classic example in the constraint satisfaction domain. While external noise is often used in constraint satisfaction algorithms, spiking neural networks naturally exhibit sufficient emergent noise.

Interestingly, the stochastic gradient descent process in error backpropagation learning is directly related to the same process as applied to constraint satisfaction, which can be understood as iterative [[error backpropagation#backpropagation to activations]] in a recurrent ([[bidirectional connectivity|bidirectionally-connected]]) network ([[@Almeida87]]; [[@Pineda88]]). This is naturally accomplished in the bidirectionally-connected [[Axon]] framework, which also implements a version of the [[GeneRec]] approximation to [[error backpropagation]].

Thus, Axon networks are continuously generating [[optimized representations]] that reflect the results of the constraint-satisfaction process, and therefore simultaneously performing search through representation space dynamically in activation space, and over the course of learning through synaptic weight adjustments.

## Emergent serial processing

Even though serial, Turing-machine like symbolic processing is cursed for searching high-dimensional spaces, serial, symbolic computation is much more _flexible_, because serial processes allow arbitrary _combinations_ of dimensions to be processed in different ways for each case. This essential advantage of serial processing is why the [[Turing machine]] is a universal computational device, whereas parallel, distributed computation generally is not, and instead must be adapted to specific problems (e.g., via a learning process).

Thus, the most effective balance of gradient-based, polynomial-time processing and serial processing is with a foundation of the former that tackles the high-dimensional nature of the real-world, with an emergent serial process operating on the resulting lower-dimensional, abstracted summary of the situation. This low-dimensional abstract space should then be amenable to more flexible serial processing along the lines of a Turing machine architecture.

We believe that this overall configuration is essential for intelligent behavior, and it provides a compelling description of many aspects of human cognition. In terms of the longstanding debate between symbolic vs. "subsymbolic" approaches to AI, this represents a synthesis of the two, with symbolic processing emerging out of fundamentally subsymbolic, neural-network hardware ([[@AndersonLebiere98]]; [[@JilkLebiereOReillyEtAl08]]).

Interestingly, [[large language models]] have this same overall configuration as well, because they are extensively trained to predict the behavior of computer programs, and thus essentially learn to behave like a Turing machine ([[@YangCampbellHuangEtAl25]]). Indeed, research has shown that excluding the computer programming aspects of the standard training corpus significantly impacts the overall cognitive flexibility of the resulting system.

<!---  TODO CITES!!. -->

## Representational issues

There are important implications for the nature of [[distributed representations]] in relation to the efficacy of polynomial-time, gradient-based neural network algorithms. Specifically, the difference between [[combinatorial vs conjunctive]] representations is critical. A _combinatorial_ representation has individual representational units representing each relevant dimension independently (e.g., each separate color and each separate shape), whereas a _conjunctive_ representation encodes specific combinations across these dimensions. Note that the full combinatorial explosion where each conjunctive combination is separately represented can be avoided by using graded and distributed representations across the high-dimensional space of conjunctions, known as _mixed selectivity_ ([[@FusiMillerRigotti16]]).

{id="figure_binding-problem" style="height:20em"}
![The binding problem in the case of combinatorial representations of individual features that are combined independently to represent distinct objects. Because these independent features do not encode the bindings between shape and color, the rest of the brain doesn't know what was actually present.](media/fig_binding_prob.png)

The purely combinatorial representation makes it difficult to simultaneously represent different unique combinations of features (e.g., a blue square and a red triangle, versus a blue triangle and a red square; [[#figure_binding-problem]]). This will cause problems for algorithms that operate directly on these representations. By contrast, conjunctive representations _bind_ together different combinations of features, so that operations on the elements of such representations can be sensitive to these different combinations.

These representational issues are nicely illustrated in the case of binding problems in the context of visual search, which has been described as using both parallel and serial processes, depending on the search display properties. For example, people can efficiently perform parallel search for basic visual features such as color and orientation, which "pop out" automatically and quickly, e.g., when looking for a red target among a collection of green distracters. As the target and distracters have more complex combinations of such features, increasingly serial spatial [[attention]] is engaged to narrow the search space ([[@Treisman77]]; [[@Wolfe10]]; [[@HerdOReilly05]]).

## Bayesian models and search

The widely-explored Bayesian models of cognition ([[@TenenbaumKempGriffithsEtAl11]]; [[@ChaterOaksfordHahnEtAl10]]) are often challenged by the curse of dimensionality, because they are largely based on symbolic-level representations that are not amenable to parallel search mechanisms. In general, the probability distributions used in these models become intractable very quickly, unless strong simplifying assumptions are made, e.g., that each variable is either statistically independent or mutually exclusive from the others. The _particle filter_ ([[@DoucetFreitasGordon01]]) or _MCMC_ (Markov chain monte carlo) sampling ([[@Neal93]]) methods use parallel threads of sequential search through high-dimensional probability spaces, and the sequential nature of these methods imposes the expected limitations of serial search more generally ([[@Sanborn17]]; [[@GershmanBeck17]]).

For these reasons and despite considerable effort and some progress, there are no practical examples of large-scale Bayesian models ([[@ZhuChenHuEtAl17]]): they instead tend to be smaller-scale models that demonstrate important principles about how inference under uncertainty might operate. There are some interesting potential connections with the behavior of discrete spiking neurons ([[@GershmanBeck17]]; [[@PougetBeckMaEtAl13]]; [[@McKeeCrandellChaudhuriEtAl22]]), which suggest a way to connect these more abstract models with more parallel underlying neural mechanisms that do scale to high-dimensional spaces.

## Gradients in evolution

[[Evolution]] in the natural world involves massively parallel search across "replicated" organisms, each living their lives in parallel, but the ontological development of each such organism is serial, which allows for considerable flexibility in the expression of the shared genomic programs that accumulate across individuals. However, this seriality of the within-lifetime trial-and-error exploration creates particular challenges for [[reinforcement learning]].

Because evolution operates on developmental programs encoded in DNA, the nature of this developmental process plays a critical role in shaping the efficiency of the evolutionary search process. To the extent that the developmental programs involve graded, gradient-based processes that allow for incremental graded searching (e.g., making a limb slightly longer or shorter, or a lobe slightly bigger), it may be possible for evolution to become more efficient than an exhaustive, enumerative search process. This would then be inconsistent with arguments that evolution is an intractable NP-complete search process (e.g., [[@RichBlokpoeldeHaanEtAl21]]). 

Nevertheless, it is important to recognize that even the mind-bending "computational power" of the evolutionary process on Earth can easily be dwarfed by the combinatorial explosion of even relatively low-dimensional systems (e.g., $2^{100} \approx 1.0e30$, which is still very large when divided by a billion years).

From the retrospective perspective of how evolution actually proceeded thus far, the progression from individual cells to organisms composed of many cells provided an important bias on the nature of processing in the brain, by making it organized around the interactions among a large number of distinct neural processing elements. This naturally favors distributed, polynomial-time processing algorithms, which fortuitously seem to provide efficient solutions to an entire spectrum of computational problems that organisms must solve.

Remarkably, the evolutionary process also settled on a body and brain blueprint (_bauplan_) for vertebrates about 500 million years ago, which has proven remarkably durable and adaptable. The human brain is structured according to this same blueprint that emerged in jawed fishes, with the same core functionality of many essential brain areas including the [[basal ganglia]], [[cerebellum]], and [[hypthalamus]] all in place. These brain systems have adapted and integrated with the [[neocortex]], which is unique to mammals starting around 200 million years ago, and is essential for our higher cognitive functions. Thus, somehow, the evolutionary search process seems to have been efficient in finding the neural algorithms that work.

