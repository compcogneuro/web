+++
Categories = ["Learning", "Computation"]
bibfile = "ccnlab.json"
+++

The **stability-plasticity dilemma** is a fundamental, intuitive constraint on the process of [[learning]]: if you learn about something new, that typically creates some form of _interference_ with or loss of previously learned information. Therefore, a given system typically has to optimize either rapid new learning (plasticity) or stable retention of prior information.

In a very simple example case, imagine a single input neuron connected to a single output neuron with a single synaptic weight. Any change in that single weight value is automatically going to change the behavior of the system relative to how it behaved previously. An obvious solution would be to create new neurons for each new thing that needs to be learned: but even in this case, as the number of such neurons multiplies, it can become increasingly difficult to then activate the correct output neuron in the face of increasing competition from all the others.

An early demonstration of this problem was given by [[@^McCloskeyCohen89]] in training an [[abstract neural network]] (using [[error backpropagation]]) on standard [[episodic memory]] tasks given to human participants. They found that the network exhibited **catastrophic interference** -- massive levels of interference well beyond what the human participants exhibited.

One solution to this problem is to introduce two different types of networks, one that is biased toward stability, and another that is biased toward rapid plasticity, which is effectively the principle behind the _complementary learning systems_ framework ([[@McClellandMcNaughtonOReilly95]]; [[@RudyOReilly99]]; [[@NormanOReilly03]]; [[@OReillyBhattacharyyaHowardEtAl14]]). In this framework, the [[hippocampus]] is optimized for rapid learning while minimizing interference, by using very _sparse_ levels of neural activity, which thus minimizes the re-use of synaptic weights across different memories. Meanwhile, the [[neocortex]] is optimized for slow, stable learning that integrates across many different experiences, which allows it to extract the statistical regularities that are common across experiences.

Various other solutions to this dilemma have been proposed ([[@CarpenterGrossberg87]]; [[@French99]]). Interestingly, the phenomenon of behavioral timescale synaptic plasticity (BTSP) ([[@Magee26]]) provides a related complementary learning systems approach, where a set of output neurons exhibit rapid plasticity to read out more slowly learned internal representations that capture the "deep" structure of the environment (see [[synaptic-plasticity#temporal eligibility traces]] for more details).

