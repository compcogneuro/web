+++
Categories = ["Computation"]
bibfile = "ccnlab.json"
+++

The widely-used **Bayes' theorem** provides a more efficient way of computing the _conditional probability_ of a hypothesis $h_j$ given some data $d$:

{id="eq_phd" title="Posterior"}
$$
P(h_j|d) = \frac{P(h_j,d)}{\sum_i P(h_i,d)}
$$

This is also known as the **posterior** probability of the hypothesis. This "pure" form of the expression is in terms of the **joint probabilities** $P(h_i,d)$, which is the probability that the given hypothesis is true and the given data is present. In this form, it is clear that the conditional probability is just a form of _normalization_ where you consider one specific hypothesis $h_j$ relative to all other possible hypotheses $h_i$, and this normalization ensures that the resulting probability sums to 1. If the data are continuous, then you can replace the sum with an _integral_.

The fundamental problem in dealing with probabilities is this pesky requirement that they sum to 1. This ultimately requires some way of accounting for the space of _all possible combinations of hypotheses and data_ (i.e., the normalizing denominator in [[#eq_phd]], also known as the **partition function**), which entails the [[curse of dimensionality]] as the size of these spaces gets larger. Thus, all but very small-scale applications of true probability-based computations are computationally _intractable_ ([[@ChaterTenenbaumYuille06]]; [[@vanRooijWrightWareham12]]). We return to this issue below ([[#approximate solutions]]), where we see that the powerful gradient descent process described in [[search]] can be used to overcome this difficulty.

Bayes' theorem is, mechanically, just a simple bit of algebra to re-express the posterior probability in terms of the **likelihood** conditional probability, $P(d | h_j)$, which tells you how likely the given observed data is, assuming the hypothesis $h_j$ is true. The key advantage of a likelihood is that it is theoretically _independent_ of all the other hypotheses, and therefore could be computed directly from the given hypothesis itself (we'll see that this actually involves a bit of wishful thinking, as you can see in [[#eq_phd]] that some kind of accounting for the other hypotheses must enter somewhere).

To derive Bayes' theorem, we can start by using the same conditional probability definition as [[#eq_phd]], in reverse:

{id="eq_pdh" title="Likelihood"}
$$
P(d | h_j) = \frac{P(d, h_j)}{\sum_k P(d_k, h_j)}
$$

Here you can see that instead of having to account for all the different hypotheses, you need to account for all the different possible data outcomes $d_k$. 

The goal of Bayes' theorem is to now express [[#eq_phd]] in terms of these likelihood terms. To make the math simpler, we introduce the **marginal** distribution as a different way of thinking about the denominator terms in these equations:

{id="eq_pd" title="Marginal"}
$$
P(d) = \sum_i P(h_i,d)
$$

The key idea here is that the overall probability of any given data value can be computed by adding up its probability of occuring under each of the different hypotheses. The term _marginal_ comes from the computation of these sums in the margins of a table listing all possible values of $h$ in one axis and all possible values of $d$ in the other axis (see [Wikipedia](https://en.wikipedia.org/wiki/Marginal_distribution)).

Thus, we can re-write the conditional probability in a more commonly-used form:

{id="eq_phd_2"}
$$
P(h_j|d) = \frac{P(h_j,d)}{P(d)}
$$

And a simple bit of algebra gives us a way of computing the joint probability in terms of the conditional probability:

{id="eq_phand"}
$$
P(h_j,d) = P(h_j|d) P(d)
$$

But this could have been done the other way around starting with [[#eq_pdh]], leading to this other definition of the same quantity (note that the joint probability is _commutative_ in the order of the arguments: $P(h_j,d) = P(d,h_j)$, whereas conditional probabilities are _not_: $P(h_j|d) \neq P(d | h_j)$:

{id="eq_phand"}
$$
P(h_j,d) = P(d|h_j) P(h_j)
$$

And thus we can substitute that into [[#eq_phd_2]] and get **Bayes' rule**:

{id="eq_bayes" title="Bayes' rule"}
$$
P(h_j|d) = \frac{P(d|h_j) P(h_j)}{P(d)}
$$

This equation now introduces the probability of the hypothesis itself being true _independent of any data_ $P(h_j)$, which is known as the **prior** probability of the hypothesis (it is _prior_ to seeing the data). This has a more intuitive interpretation conceptually, reflecting how plausible or generally reasonable the hypothesis might be. This is where one might want to introduce a penalty factor favoring simpler hypotheses over more complex ones, for example (i.e., _Occams's razor_).

Thus, the major attraction of Bayes' rule is that it transforms the posterior probability computation into something that involves more intuitively appealing elements:

{id="eq_concepts"}
$$
P(h_j|d) = \frac{\rm{likelihood} \; x \; \rm{prior}}{P(d)}
$$

However, you still end up with this pesky normalization factor of $P(d)$ in the denominator, which still entails the curse of dimensionality. Furthermore, computing the prior in a properly-normalized way also requires computing the probability of a particular hypothesis $h_j$ relative to all the other possible hypotheses, so you still depend on somehow covering that space.

You can get rid of the denominator by computing **odds ratios**, where the shared denominator cancels out:

{id="eq_odds" title="Odds ratio"}
$$
O_d(h_i|h_j) = \frac{P(h_i|d)}{P(h_j|d)} = \frac{P(d|h_i) P(h_i)}{P(d|h_j) P(h_j)}
$$

Of particular relevance to cognitive applications, the Bayesian formulation provides a natural way of thinking about _perception_ as a **generative** process, in terms of the likelihood of generating the observed data based on the internal hypotheses represented in the brain (or model thereof) ([[@RaoBallard99]]; [[@Friston05]]). Specifically, the **inference** process of what hypotheses are most likely to be "responsible" for the current sensory input data (i.e., the posterior $P(h_j|d)$ involves computing the likelihood for each hypothesis to have generated the data ($P(d|h_j)$) -- this is the fundamental transformation that Bayes' rule provides.

See [[predictive-learning#Bayesian predictive coding]] for more details.

## Approximate solutions

Because probabilities must be normalized, performing exact Bayesian computation quickly becomes intractable as the scale of the space of hypotheses and / or data increases. There are two main techniques that are widely used to combat this problem, each of which is not without its own limitations.

One technique is **sampling**, which involves running the model to generate samples from the generative process, and using these samples to infer the larger probability distribution. This is known as **Markov chain Monte Carlo (MCMC)** sampling ([[@Hastings70]]; [[@GelfandSmith90]]), of which **Gibbs sampling** ([[@GemanGeman84]]) and **particle filters** (sequential Monte Carlo) ([[@DelMoralDoucetJasra12]]) are examples, within the broader **Approximate Bayesian Computation (ABC)** approach ([[@MarinPudloRobertEtAl12]]; [[@RobertCasella04]]). This is the technique used in the [[Boltzmann machine]], for example.

The other technique is **variational inference (VI)** which uses gradient-based algorithms operating on a "proxy" for the true probability distribution, which is computationally tractable ([[@BleiKucukelbirMcAuliffe17]]; [[@VafaiiGalorYates25]]). As discussed in [[search]], gradient-based search in neural-network-like models with synaptic weights connecting representational nodes has proven to be an unexpectedly (and "unreasonably") efficient way of searching high-dimensional spaces, and VI leverages this core computation to get around the [[curse of dimensionality]] associated with using probabilistic models.

To see how VI works, we introduce some new notation, where ${\bf z}$ is a vector of internal parameters of a generative model (e.g., the weight values, or, more commonly in this case, the mean and standard deviation of a Gaussian random noise generator), and ${\bf x}$ is the current vector of input data (i.e., $d$ from above). Also, for some reason, the lower-case $p$ is preferred in this area, instead of $P$. In these terms, the Bayesian generative model is:

{id="eq_pzx" title="Model posterior"}
$$
p({\bf z} | {\bf x}) = \frac{p({\bf x}|{\bf z}) p({\bf z})}{p({\bf x})}
$$

Because we cannot compute $p({\bf x})$, this is not tractable to compute. However, we can introduce an approximation to this function that _is_ tractable to compute, e.g., because it doesn't have any normalization terms itself: $q_{\theta}({\bf z} | {\bf x})$. The only constraint imposed on this approximation function is that its parameters $\theta$ should be optimized to minimize the "difference" between $q$ and the true posterior probability that we want. 

The standard way to compute a difference between two probability distributions is the **Kullback-Liebler divergence** function (which is the log of the integrated odds ratio between the two distributions, over the relevant parameters, $\theta$ in this case):

{id="eq_kl" title="Kullback-Liebler divergence"}
$$
KL(q_{\theta}({\bf z} | {\bf x}) || p({\bf z} | {\bf x}) ) = E_{\theta} [log q_{\theta}({\bf z} | {\bf x}) ] - E_{\theta} [log p({\bf z} | {\bf x}) ]
$$ 

But again, this KL divergence _still_ involves the same intractable probability function $p({\bf z} | {\bf x})$. However, we substitute [[#eq_pzx]] into this KL equation:

{id="eq_kl" title="Expanded KL"}
$$
KL(q_{\theta}({\bf z} | {\bf x}) || p({\bf z} | {\bf x}) ) = E_{\theta} [log q_{\theta}({\bf z} | {\bf x}) ] - E_{\theta} [log p({\bf x} | {\bf z}) p({\bf z})] + log p({\bf x})]
$$ 

which reveals the intractable $p({\bf x})$ factor (which does not depend on $q$ and therefore does not require the expectation over the parameters of q. However, the other terms in this expression beside that one are in fact tractably computable, so in effect you can do some algebra and end up _maximizing_ the negative of these other terms, which provides a lower-bound on the actual target KL divergence. This lower bound is called the **evidence lower bound (ELBO)**:

{id="eq_elbo" title="Evidence lower bound"}
$$
\rm{ELBO}(q_{\theta}) =  E_{\theta} [log p({\bf x} | {\bf z}) p({\bf z})] - E_{\theta} [log q_{\theta}({\bf z} | {\bf x}) ]
$$

The key point here is that both of these terms are computationally tractable -- the first is just the non-normalized generative likelihood factor from the model, and the second is designed to be directly computable from its parameters, again without any intractable normalization factor. Clearly, the resulting behavior is going to be an approximation to the full probabilistic model, but the results from various VI models indicate that this is not a major factor. However, there are still potential computational efficiency issues when attempting to extend this approach into deeply-nested architectures ([[@VafaiiGalorYates25]]).

Interestingly, this variational inference framework and its connection to the gradient descent computation performed in error backpropagation has also allowed the derivation of a number of different learning algorithms starting from a more abstract, generalized probabilistic framework ([[@KhanRue23]]; [[@VastolaGershmanRajan26]]).

There are also other techniques for performing inference in unnormalized probabilistic models, for example by estimating the normalization denominator, and using contrastive functions that subtract away the shared normalization function ([[@GutmannHyvarinen13]]; [[@MatsudaHyvarinen19]]; [[@Hinton02]]).

