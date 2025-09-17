+++
Categories = ["Computation", "Activation", "Learning"]
bibfile = "ccnlab.json"
+++

**Search** is perhaps the single most general unifying concept in all of [[computation]].

* _Problem solving_ can be defined as search through _problem space_, as in many classical symbolic [[artificial intelligence]] (AI) systems defined it ([[@NewellSimon72]]; [[Newell91]]).

* _Planning_ is search through _action space_ to accomplish a desired outcome, which is the focus of many approaches in [[reinforcement learning#model-based]] reinforcement learning.

* _Learning_ is search through _representation space_, to find the best [[linear algebra|basis]] for representing inputs, that supports the desired computational processes and behavioral outputs.

* _Evolution_ is search through _phenotype space_, e.g., as in the widely-used [[genetic algorithm]].

This is not an empty tautology, because it provides a unified understanding about a central problem faced in all of these domains:

> Each of these spaces is plauged by the [[curse of dimensionality]], and effective solutions for real-world, large-scale problems must somehow tackle the exponential explosion problem.

Indeed, the field of [computational complexity theory](https://en.wikipedia.org/wiki/Computational_complexity_theory) (wikipedia) provides a unified framework based on an equivalence class across a wide range of problems, known as _NP complete_ that are not solvable in polynomial time. The NP means _nondeterministic polynomial_ time, which means that the solution can be verified quickly (in polynomial time) but actually generating the solution takes much longer, e.g., exponential time.

## Enumerative search is cursed

If the key element of the algorithm used in any of these domains involves an _enumerative_ search through each state, which involves enumerating specific combinations of relevant values along the different dimensions, then that algorithm will suffer from the curse of dimensionality, and will exponentially-quickly become intractable as the state space increases.

For example, the simplest _brute force_ algorithm for any such problem involves a giant set of nested `for` loops through all values along each dimension, enumerating each possible combination, and then evaluating some function computed on this state vector. As anyone who has experienced such nested `for` loops will recognize, these bog down very quickly as the number of such levels is increased. 

Although it is intuitively appealing to refer to this as a _serial_ search process, the problems persist even if you use a parallel implementation wrapped around this enumerative process. Any amount of parallelism will just divide the exponentially-growing factor by a fixed constant on the order of the parallel hardware involved. Even if this is the brain with 100 billion neurons, that factor will quickly pale in comparison to the combinatorial explosion of something with only 100 different dimensions (e.g., $2^100 / 1.0e11 = 1.0e19$, for 100 binary-valued dimensions).

## Dimensionally based search is necessary

Thus, the only plausible way to search in high-dimensional spaces is to use an algorithm that operates directly on the dimensions themselves, and not on their combinatorial enumerations, while still somehow capturing the combinatorial interactions across the dimensions that are relevant to the problem at hand. This may sound like magic, but in fact we have a number of very well-established examples that demonstrate the power of such algorithms to tackle hard, very high-dimensional problems in surprisingly efficient ways.

The paradigmatic example for our purposes is the _stochastic gradient descent_ process used in [[error backpropagation]] learning, which computes the _first order_ partial derivatives (i.e., _gradients_) of the synaptic weights relative to an overall error function (aka _objective_ function). The computational cost of backpropagation is linear on the order of the number of synaptic weights and yet has been shown to be highly effective in optimizing neural networks to accomplish a very wide array of different tasks, including of course the [[large language models]] (LLMs) powering the widely-used ChatGPT like AI models.

Interestingly, many people considered using something like the backpropagation algorithm over the years, but didn't even bother because they assumed that this simple gradient-based algorithm would just produce horribly sub-optimal solutions (i.e., _local minima_). While there is no guarantee that gradient descent will produce optimal solutions, in practice, with the appropriate heuristic techniques (e.g., the _AdaMax_ mechanism; [[@KingmaBa14]]) and network configurations, it certainly does well enough to be of immense practical importance.

This provides an important object lesson for the computational complexity theory approach, which is traditionally framed in terms of optimal solutions and worst-case scenarios, instead of considering the potential for dimensionally based, gradient descent algorithms to produce "good enough" solutions in the "real world" scenarios that actually matter. It is undoubtedly very difficult to formally analyze such vague constraints, but nevertheless, this is obviously what we care about in considering how the brain actually functions.

<!--- TODO: find better lit on this stuff. Initial google search not that promising: https://scholar.google.com/scholar?hl=en&as_sdt=0%2C5&q=%22backpropagation%22+%22computational+complexity%22+%22np%22+%22nondeterministic+polynomial%22&btnG= -->

## Learning as search

Perhaps the biggest divide between classical symbolic AI and the "new" neural-network based approaches is that neural networks enable a dimensional, gradient-based learning strategy that allows learning to search through high-dimensional representational spaces, now reaching into the billions of parameters, which represents a practically infinite combinatorial space (with exponential size well beyond any human comprehension, exceeding the number of atoms on the universe by exponentially-large exponential factors!).

By contrast, the discrete symbolic nature of classical AI frameworks, where the models are essentially [[Turing machine]]s with various AI-based mechanisms, is incompatible with such a gradient-based learning mechanism, and thus any attempt to learn in such models immediately collapses in the face of the curse of dimensionality.

The _stochastic_ nature of the stochastic gradient descent process derives from the use of small subsamples of the total problem space to compute these gradients, whereas using the full space ends up getting "stymied" by all the constraints imposed across this entire space. Considerable research is now illuminating why it actually works as implausibly well as it does (e.g., [[@NakkiranKaplunBansalEtAl21]]; [[@Shwartz-ZivTishby17]]; [[@VidalBrunaGiryesEtAl17]]).

## Constraint satisfaction as search

The concept of [[constraint satisfaction]] provides another high-level unified way of understanding a wide range of computational processes, and this generality is not accidental from the present search-based perspective. In this case, it is a search over possible states that solve or optimize a set of constraints. Many problems can be formulated in this way, including planning and problem solving, so constraint satisfaction provides a more general way of understanding the essential computations in these domains.

Here too the stochastic, gradient-based approach to search is the only one capable of handling high-dimensional state spaces, and it also is generally the most effective approach, and does not suffer too much from local minima problems ([[@HoosTsang06]]). One class of such algorithms is in fact a neural network, going back to the pioneering approach of [[@HopfieldTank85]], who showed that iterative updating of a bidirectionally-connected Hopfield network could provide good solutions to the travelling salesman problem, which is a classic example in the constraint satisfaction domain.

In the [[Axon]] framework, the use of the [[GeneRec]] approximation to [[error backpropagation]] via [[bidirectional connectivity]] and [[temporal derivative]] learning rules results in the ability to perform constraint satisfaction by effectively performing [[error backpropagation#backpropagation to activations]]. Thus, Axon networks are continuously generating [[optimized representations]] that reflect the results of this constraint-satisfaction process, and therefore simultaneously performing search through representation space dynamically in activation space, and over the course of learning through synaptic weight adjustments.

## Emergent serial processing

Even though serial, Turing-machine like symbolic processing is cursed for searching high-dimensional spaces, serial, symbolic computation is much more _flexible_, because serial processes allow arbitrary _combinations_ of dimensions to be processed in different ways for each case. This essential advantage of serial processing is why the [[Turing machine]] is a universal computational device, whereas parallel, distributed computation generally is not, and instead must be adapted to specific problems (e.g., via a learning process).

Thus, the most effective balance of gradient-based, dimensional processing and serial processing is with a foundation of the former that tackles the high-dimensional nature of the real-world, with an emergent serial process operating on the resulting lower-dimensional, abstracted summary of the situation. This low-dimensional abstract space should then be amenable to more flexible serial processing along the lines of a [[Turing machine]] architecture.

We believe that this overall configuration is essential for intelligent behavior, and it provides a compelling description of many aspects of human cognition. In terms of the longstanding debate between symbolic vs. "subsymbolic" approaches to AI, this represents a synthesis of the two, with symbolic processing emerging out of fundamentally subsymbolic, neural-network hardware.

Interestingly, [[large language models]] have this same overall configuration as well, because they are extensively trained to predict the behavior of computer programs, and thus essentially learn to behave like a Turing machine ([[@YangCampbellHuangEtAl25]]). Indeed, research has shown that excluding the computer programming aspects of the standard training corpus significantly impacts the overall cognitive flexibility of the resulting system.

<!---  TODO CITES!!. -->

## Representational issues

There are important implications for the nature of [[distributed representations]] in relation to the efficacy of dimensional, gradient-based algorithms. Specifically, the difference between [[combinatorial vs conjunctive]] representations is key. A combinatorial representation has individual representational units representing each relevant dimension independently (e.g., each separate color and each separate shape), whereas a conjunctive representation encodes specific combinations across these dimensions (typically using graded and distributed representations).

The purely combinatorial representation makes it difficult to simultaneously represent different unique combinations of features (e.g., a green square and a red triangle, versus a green triangle and a red square). This will cause problems for dimensionally based algorithms that operate on these representations. However, conjunctive representations _bind_ together different combinations of features, so that operations on the elements of such representations can be sensitive to these different combinations.

This is relevant to binding issues in the visual system, which has been described as using both parallel and serial processes, depending on the target properties. We can efficiently perform parallel search for basic visual features such as color and orientation, but increasingly complex combinations of such features require increasingly serial spatial [[attention]] to narrow the search space ([[@Treisman77]]; [[@Wolfe10]]; [[@HerdOReilly05]]).

## Bayesian models and search

The widely-explored Bayesian models of cognition ([[@TenenbaumKempGriffithsEtAl11]]; [[@ChaterOaksfordHahnEtAl10]]) are often challenged by the curse of dimensionality, because they are largely based on symbolic-level representations that are not amenable to parallel search mechanisms. In general, the probability distributions used in these models become intractable very quickly, unless strong simplifying assumptions are made, e.g., that each variable is either statistically independent or mutually exclusive from the others. The _particle filter_ ([[@DoucetFreitasGordon01]]) or _MCMC_ (Markov chain monte carlo) sampling ([[@Neal93]]) methods use parallel threads of sequential search through high-dimensional probability spaces, and the sequential nature of these methods imposes the expected limitations of serial search more generally ([[@Sanborn17]]; [[@GershmanBeck17]]).

For these reasons, there are no examples of large-scale Bayesian models: they are all smaller-scale models that demonstrate important principles about how inference under uncertainty might operate. There are some interesting potential connections with the behavior of discrete spiking neurons ([[@GershmanBeck17]]; [[@PougetBeckMaEtAl13]]; [[@McKeeCrandellChaudhuriEtAl22]]), which suggest a way to connect these more abstract models with more parallel underlying neural mechanisms that do scale to high-dimensional spaces.

## Gradients in evolution

[[Evolution]] in the natural world involves massively parallel search across "replicated" organisms, each living their lives in parallel, but the ontological development of each such organism is serial, which allows for considerable flexibility in the expression of the shared genomic programs that accumulate across individuals. The seriality of trial-and-error exploration creates particular challenges for [[reinforcement learning]].


