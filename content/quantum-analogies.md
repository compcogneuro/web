+++
Categories = ["Computation"]
bibfile = "ccnlab.json"
+++

There are various interesting applications of concepts derived from **quantum physics** to neuroscience, which are highlighted here.

## Classical probabilities vs quantum probabilities

There have been various attempts to understand at a more abstract, principled level the fundamental differences between the standard "classical" notions of probability, versus the kinds of probability functions that describe the quantum realm (e.g., [[@Zurek05]]). These differences provide useful analogies for understanding the dynamics and representations of neural systems. In particular, [[bidirectional connectivity|bidirectionally connected]] networks that support [[constraint satisfaction]] processing dynamics behave somewhat like quantum systems, where there can be strong _interactions_ among different parts of the network, causing the overall system to behave in more chaotic and unpredictable ways, relative to feedforward networks without the mutual interdependencies caused by bidirectional connectivity.

A useful starting point for understanding these differences is in terms of [[linear algebra#non-negative factorization]] vs. sign-unconstrained factorization (i.e., principal components analysis or PCA). Read the section at the above link (and the earlier part of that page on PCA if otherwise unfamiliar with that) before proceeding.

In short, classical probabilities are like non-negative factors (see Kolmogorov's [probability axioms](https://en.wikipedia.org/wiki/Probability_axioms)): they can only be 0 or positive, and thus they define a purely _additive_ set of factors, which must ultimately sum to 1 across all relevant possibilities. As with the case of non-negative factorization, these probabilities align directly and uniquely with the discrete, _mutually-exclusive_ outcomes that they are associated with (heads vs. tails in a coin flip, for example).

By contrast, quantum probabilities are defined in terms of a [Hilbert space](https://www.intoquantum.pub/p/an-introduction-to-the-hilbert-space) in terms of orthogonal basis vectors, with probability amplitudes that are _complex_ numbers (i.e., a combination of a real and an imaginary number). The fact that complex numbers are two-dimensional (2D) entities mathematically implies that they are in general **non-commutative**, for the same reason that matrix multiplication is non-commutative: values can "mix" across dimensions, and the order of these mixing processes actually _matters_: $AB \neq BA$ in general.

Physically, this non-commutative property of quantum physics is associated with Heisenberg's _uncertainty principle_: measuring information along one dimension (e.g., position) necessarily makes the other dimension (e.g., momentum or velocity) _less_ certain. In other words, measuring something about a quantum system necessarily _changes_ the system being measured, so that the _order_ of such measurement operations matters (and commutativity is all about the order of such operations).

To understand all this more concretely, we need a brief digression on complex numbers.

The imaginary number: $i = \sqrt{-1}$ seems very mysterious, but really it is just an efficient way to represent an orthogonal additional dimension relative to real numbers, thereby providing an efficient way to represent a 2D value with a complex number (real + imaginary). Furthermore, because of the -1 in there, complex numbers "intrinsically" have positive and negative components, and thus violate the non-negativity constraint.

This built-in combination of positive and negative makes it natural to represent rotations and oscillations in terms of multiplications of complex numbers, because when you multiply two imaginary numbers they turn back into a real number:

$$
\sqrt{-1} \sqrt{-1} = \sqrt{-1}^2 = -1
$$

Meanwhile, the cross-products from multiplying complex numbers, where you multiply the real and imaginary components:

$$
(a + bi)(c + di) = ac - bd + (ad + cb)i
$$

then sucks in the real part into the imaginary world. Thus, the two dimensions of a complex number really like to rotate around into each other. To make this happen "smoothly" instead of in one big jump from imaginary to real, fractional exponential values can be used, which is the essence behind [Euler's formula](https://en.wikipedia.org/wiki/Euler%27s_formula):

$$
e^{ix}= \cos x + i \sin x
$$

Which makes it clear that the complex number on the right-hand side describes a circle in a 2D plane, that rotates around as a function of the angle _x_. One amazing implication of this relationship between the natural exponential function and the role of _x_ as an angle, is that if you plug in $x=\pi$ into this equation, then the result is -1 ($\cos \pi = -1$ and $\sin \pi = 0$), i.e.,:

$$
e^{i\pi} + 1 = 0
$$

which is one of the coolest equations in mathematics!

The final key element of quantum probabilities is that the equivalent of the constraint of all probabilities summing to 1 (i.e., the _norm_) is instead represented in terms of the _radius_ of the resulting vector in the complex hypersphere always being of length 1. In other words, the pure state of any quantum system is constrained to lie on the surface of a hypersphere (e.g., the [Bloch sphere](https://en.wikipedia.org/wiki/Bloch_sphere) for a 2 dimensional system). As in trigonometry, this radius is related to the _hypotenuse_ in the complex 2D space, and thus involves the Pythagorean theorem:

$$
r^2 = x^2 + y^2
$$

In the world of complex numbers, the squaring operation is performed by multiplying by the _complex conjugate_ (the complex number with the opposite sign).

Critically, this radius-based constraint _automatically_ introduces inter-dependencies between the different dimensions of a quantum system. Just as with the rotation operation, the system is constrained to have a _zero sum_ character: if you move in one direction relative to a given basis vector, then you automatically must compensate for that along the other basis vector, for the radius to remain a fixed length.

In effect, you can never _create_ or _destroy_ quantum probabilities -- they just rotate around on a sphere, which changes the alignment of different outcomes according to the basis vectors being used. This also means that the choice of such basis vectors is relatively less important: there isn't a privileged connection between a probability and each such basis vector

A corollary is that the quantum probabilities must remain fundamentally interconnected ("entangled") with each other, due to the nature of the quantum norm as a hypotenuse: when one changes, the other must necessarily follow suit, to preserve this "collective" norm.

Thus, the essence of the connections between non-negative factorization versus sign-unconstrained factorization (e.g., PCA), and classical probability versus quantum probability are:

* In both classical probabilities and non-negative factorization, there _is_ a privileged connection between the probability and the event it describes, and there is a strong mutual-exclusivity constraint such that each probability must apply uniquely to an orthogonal event (i.e., a [[distributed representations#localist representations|localist]] underlying representation).

* In quantum probabilities and sign-unconstrained factorization (e.g., PCA), there is _not_ a privileged connection between the probability and the event it describes: the probabilities are free to rotate around on a sphere defined by the basis vectors, and _any orthogonal set of basis vectors is as good as any other_.

For example, [[@^Zurek05]] leverages this principle with the concept of _einvariance_ to demonstrate the fundamental differences between classical probability and quantum probability, in relation to the defining "principle of indifference" that has been used to motivate classical probability theory. Specifically, probabilities quantify the lack of knowledge ("indifference") about the outcomes of various events. In effect, he shows that there is a much deeper level of "indifference" present in quantum systems, due to the lack of a privileged basis space.

At a broader level, this fundamental "indifference" or arbitrariness of the quantum world is what gives rise to these ideas that _reality_ at the quantum level is not a well-defined concept: it is all fundamentally arbitrary and interconnected, and any time you try to measure something, you inevitably end up rotating everything around and thus fundamentally changing the system. And you'll never be able to recover what state the system was in _before_ you interfered with it, so at some level, it doesn't make sense to attribute any kind of physical reality to that prior state.

These properties of the quantum realm can be seen as arising from a fundamentally wave-based system: waves automatically impose positive and negative _interference effects_ as they propagate and interact with other waves. Furthermore, waves are in a constant state of motion, constantly oscillating, which is essentially rotating through state space. Therefore, there really is no stable reality to a wave: it is a fundamentally dynamic, ephemeral thing. And you can't interact with the wave medium (e.g., water) without creating waves of your own, which will then propagate out and interfere with everything else.

Furthermore, the quantum wave function, e.g., the Schrödinger equation, describes a fundamentally _conservative_ wave, where the complex square of the wave (i.e., its total "probability") remains strictly fixed over time. This implies the hypotenuse-like zero-sum behavior and consequent interactions among even distal areas of the wave, even though the wave equation itself can be computed entirely in terms of locally-available state variables.

By contrast, classical probability is concerned with strictly isolated, localized things (like hard little "particles" -- e.g., balls in an urn), that do not mutually interfere or interact like waves do. The standard _Copenhagen_ interpretation of quantum physics (due to Bohr and Heisenberg) explains this transition from the quantum to the classical in terms of the sudden, instantaneous _collapse_ of the wave function at moment when a _measurement_ occurs. However, this explanation is fraught with all manner of logical inconsistencies, and is responsible for much of the confusion surrounding the nature of the quantum realm.

By contrast, the _de Broglie Bohm_ interpretation of quantum physics involves hard particles "surfing" quantum wave functions, and is entirely compatible with the standard quantum physics framework (see [wavereality.org](https://wavereality.org) for more details). Thus, the wave-particle duality at the heart of quantum mechanics is embraced in this synthetic picture, and the problematic notion of instantaneous wave function collapse is entirely avoided. Under this perspective, the quantum world dominates at the relatively short spatial and time scales where the wave functions dominate the dynamics, while the classical world emerges in terms of phenomena that reflect more of the particle-like properties.

## Application to neuroscience

In relation to neuroscience, both the classical and quantum pictures apply in various ways, perhaps metaphorically corresponding to the de Broglie Bohm framework where both are always at work, and different situations emphasise these different aspects to different degrees.

First, in connection with [[linear algebra#non-negative factorization]], the fact that neurons obey these non-negativity constraints imposes significant constraints on the way that they represent information, and on the effective degrees of freedom in these representations, in comparison to [[abstract neural network]]s (ANNs) that lack such non-negativity constraints. Specifically, non-negativity should automatically make the neural representations more directly _interpretable_ in terms of external variables in the outside world that are represented in the brain. This can help to explain why electrical recordings of neurons throughout the brain often show direct correlations with stimuli and behavioral responses of interest. 

Although we often take these correlations for granted, such correlations are generally much more difficult to extract from ANNs, which have been criticized as being "black boxes" that are very difficult to understand. Thus, one potential advantage of using more biologically-based models that incorporate these full non-negativity constraints should be in producing more interpretable behavior in terms of underlying neural representations.

However, as noted earlier, [[bidirectional connectivity|bidirectionally connected]] networks that support [[constraint satisfaction]] attractor dynamics behave much more like quantum systems, with the ability for relatively small differences in some of the constraints to strongly influence the resulting global state that the network settles into. This is the characteristic of dynamic systems at the edge of _criticality_, which exhibit more _chaotic_ behavior ([[@Langton90]]; [[@MaassNatschlagerMarkram02]]; [[@VerstraetenSchrauwenDHaeneEtAl07]]; [[@Carroll20]]). Thus, such systems should exhibit more quantum wave-like effects, where overt behavior can be more strongly influenced by environmental context factors, and the sequential ordering of events can have a relatively large effect (e.g., [[@BruzaWangBusemeyer15]]).

Interestingly, the role of [[language]] in human cognition may play a role somewhat analogous to that of particles in the quantum world: inside an individual's own brain, the world is all wavy and chaotic, but the need to reduce all of that down to a discrete symbolic utterance that can be communicated to another human being generates a hard, discrete, particle-like constraint onto the system. It is often the case that one only recognizes the limits of one's own understanding when trying to communicate it to another; this is why teaching and writing papers is so synergistic with the overall scientific research mission. It forces one to clarify, localize, and formalize understanding (especially if you happen to not be dispositionally predisposed to such forms of thinking in the first place).

In short, whereas there has been a fundamental confusion in the context of the Copenhagen interpretation about the role of human [[conscious-awareness|consciousness]] in the measurement process of quantum systems, perhaps the nature of human consciousness is actually well-captured _metaphorically_ by the nature of the quantum world. It is fundamentally unitary and chaotic and dynamic, but also supports a discrete level of symbolic interpretation riding on top of those bidirectional waves of neural activity. The interactions between these two complementary aspects of the whole system are likely synergistic, with each required for the full flowering of the unique capabilities of human intelligence.

Perhaps unfortunately, many have tried to go beyond this metaphorical connection, to attempt to establish a direct physical connection between quantum phenomena and consciousness. This seems unlikely, given the macroscopic, hot, and chaotic nature of the human brain. Nevertheless, perhaps the metaphorical connections can still be somewhat beneficial in deepening our understanding of the overall system.

