+++
Categories = ["Neuroscience", "Computation"]
bibfile = "ccnlab.json"
+++

The **subsumption** architecture for robotics advanced by [[@^Brooks86]] provides a valuable way of understanding the results of [[evolution]] in the brain, where new systems and mechanisms have been bolted on top of existing ones to support new levels of functionality. Two important features of this organization are:

* The new system can take advantage of all the existing functionality, and therefore provide a more specific type of function that doesn't need to handle all of problems that have already been solved.

* The new system typically needs to have ways of controlling and modulating elements of the existing functionality, so that it can exert its own control over behavior. For example, a more considered pathway of responding needs to suppress more automatic reflex pathways encoded in existing brain systems.

From a computer science, programming perspective, this is like [monkey patching](https://en.wikipedia.org/wiki/Monkey_patch) (wikipedia link) new functionality on top of an existing codebase. Although people often end up _refactoring_ an entire codebase once they realize how all the new functionality could be better accomplished using a single coherent design, it is difficult for evolution to perform such a massive refactoring, because it always has to ship highly functional code in order to survive. Therefore, it tends to make much more incremental changes.

Because evolution often operates by an incremental, additive process, the overall result can be a convoluted mess of these layers upon layers of monkey-patched systems. On the other hand, if everything has appropriate adaptive learning and tuning mechanisms, the emergent behavior can be quite robust.

