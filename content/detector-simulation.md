+++
Categories = ["Simulations"]
bibfile = "ccnlab.json"
+++

<sim-detector>

<div>

## Introduction

This simulation shows how an individual [[neuron]] can act like a [[neuron detector|detector]], picking out specific patterns from its inputs and responding with varying degrees of selectivity to the match between its synaptic weights and the input activity pattern. See [web/sims/detector](https://github.com/compcogneuro/web/tree/main/sims/detector) for the source code.

We will see how a particular pattern of weights makes a simulated neuron respond more to some input patterns than others. By adjusting the level of excitability of the neuron, we can make the neuron respond only to the pattern that best fits its weights, or in a more graded manner to other patterns that are close to its weight pattern. This provides some insight into why the neuron activation function works the way it does.

## The network and input patterns

The `Network` has an `Input` layer that will have patterns of activation in the shape of different digits, and these input neurons are connected to the receiving neuron (`RecvNeuron`) via a set of weighted synaptic connections. We can view the pattern of weights (synaptic strengths) that this receiving unit has from the input, which should give us an idea about what this unit will detect.

* Select the `Wts` tab at the top of the list of network variables at the left of the [[#sim-detector:Network]] view, then click [[#sim-detector:Network/r.Wt]] as the value you want to display, and then click on the `RecvNeuron` to view its receiving weights.

You should now see the `Input` grid lit up in the pattern of an `8`. This is the weight pattern for the receiving unit for connections from the input units, with the weight value displayed in the corresponding sending (input) unit. Thus, when the input units have an activation pattern that matches this weight pattern, the receiving unit will be maximally activated. Input patterns that are close to the target `8` input will produce graded activations as a function of how close they are. Thus, this pattern of weights determines what the unit detects, as we will see. First, we will examine the patterns of inputs that will be presented to the network.

* In the data browser on the left side of the simulation, click on `Inputs` to open it up, and then click on `Digits` to see all of the input patterns.

The display that comes up shows all of the different *input patterns* that will be presented ("clamped") onto the `Input` layer, so we can see how the receiving unit responds. Each row of the display represents a single *trial* that will be presented to the network. As you can see, the input data in this case contains the digits from 0 to 9, represented in a simple font on a 5x7 grid of pixels (picture elements). Each pixel in a given event (digit) will drive the corresponding input unit in the network.

## Running the network

To see the receiving neuron respond to these input patterns, we will present them one-by-one, and determine why the neuron responds as it does given its weights. Thus, we need to view the activations again in the network window.

* Select the `Act` tab at the top of the network variables and click [[#sim-detector:Network/Act]] to view activations, then click the [[#sim-detector:Step]] button in the toolbar, which will step one `Trial` as indicated.

This activates the pattern of a `0` (zero) in the `Input`, and shows the 200 cycles (milliseconds) of **settling** that make up one [[theta rhythm|theta cycle]] trial, during which the state of the receiving neuron is iteratively updated according to the spiking neuron equations, just as the unit in the [[neuron simulation]] was updated over time.

The receiving unit showed an activity value of 0 because it never got activated above its firing threshold by the `0` input pattern. Before getting into the nitty-gritty of why the unit responded this way, let's proceed through the remaining digits and observe how it responds to other inputs.

* Press [[#sim-detector:Step]] for each of the other digits, until the number `8` shows up.

You should have seen the receiving unit finally activated when the digit `8` was presented, with an activation of zero for all the other digits. Thus, as expected, the receiving unit acts like an `8` detector: only when the input perfectly matches the input weights is there enough excitatory input to drive the receiving neuron above its firing threshold.

* You can use the "VCR" style buttons at the bottom of the `Network` to review each cycle of updating, to see the progression of activation over time. Selecting [[#sim-detector:Network/Raster]] in the network toolbar, and viewing [[#sim-detector:Network/Spike]], gives a nice picture of the individual spikes over time.

* Go ahead and do one more [[#sim-detector:Step]] to see what happens with `9`.

We can use a graph to view the pattern of receiving unit activation across the different input patterns.

* Click on the [[#sim-detector:Test Trial Plot]] tab.

The graph shows the excitatory net input ([[#sim-detector:Test Trial Plot/GeSyn]]) and the activation ([[#sim-detector:Test Trial Plot/Act]]) for the unit as a function of trial (and digit) number along the X axis. For `Act` you should see a flat line with a single peak at 8.

## Computing the excitatory conductance (net input)

Now, let's try to understand exactly why the unit responds as it does. The key to doing so is to understand the relationship between the pattern of weights and the input pattern.

* Go back to the `Wts` / [[#sim-detector:Network/r.Wt]] display in the network, and open the `Inputs/Digits` patterns again, so you can see both the weights and the digit patterns at the same time.

{id="question_overlap"}
> For each digit pattern, report the number of active units in the pattern where there is also a weight of 1 according to the `8` digit pattern shown in the `r.Wt` view in the network. In other words, report the *overlap* between the digit input activity and the weight pattern.

The number of inputs having a weight of 1 that you just calculated determines the total excitatory synaptic conductance `GeSyn` going into the receiving unit, which is also known as the **net input**. As covered in [[neuron#Computing input conductances]] and [[neuron dendrites#eq_get]], this is the average of the sending activation $x_i$ times the weight $w_i$ over all $N$ of the inputs, with a correction factor $\alpha$ for the expected activity level of the sending layer:

$$
GeSyn = \frac{1}{\alpha} \frac{1}{N} \sum_i x_i w_i
$$

* You can select the `Gmisc` tab in the network variables and click [[#sim-detector:Network/GeSyn]] to see these values in the network as you step through the inputs, and it is also plotted in the [[#sim-detector:Test Trial Plot]].

* Do [[#sim-detector:Init]] and [[#sim-detector:Step]] to see the `0` input again. If you hover over the `RecvNeuron` with your mouse, you should see that it has a value of `GeSyn = .3529..`.

To apply the above equation, you should have observed that `0` has 6 units in common with `8`, and $N = 35$ (7x5), so $\frac{1}{N} \sum_i x_i w_i$ is about .1714. Next, we need to apply the $\alpha$ correction factor, which is the expected activity level of the input layer, set here to be the activity level of the `8`: 17 of the 35 units are active. Thus, we should get:

$$
GeSyn = \frac{1}{17/35} \times \frac{6}{35} = \frac{6}{17} = .3529...
$$

which is exactly what the model shows. Notice that the two factors of 35 cancel out, so the net input reduces to just the *overlap divided by 17*, where 17 is the number of active units in the weight pattern. Thus the `8` itself, which matches the weights perfectly, gives $GeSyn = 17/17 = 1$, and you can multiply any of these `GeSyn` values by 17 to read the overlap off directly, confirming your answers above.

As a result of working through this net input calculation, you should now have a detailed understanding of how the net excitatory input to the neuron reflects the degree of match between the input pattern and the weights. You have also observed how the activation value can ignore much of the graded information present in this input signal, due to the presence of the **threshold**. This gives you a good sense for *why* neurons have these thresholds: it allows them to filter out all the "sub-threshold noise" and only communicate a clear, easily interpreted signal when it has detected what it is looking for.

* Click on [[#sim-detector:Test Trial Plot/Ge]] in the plot to compare it against `GeSyn`.

`Ge` is the *total* excitatory conductance, which adds the voltage-sensitive [[neuron channels#NMDA]] and [[neuron channels#VGCC]] channel contributions on top of the AMPA synaptic conductance in `GeSyn`. Because these channels are voltage sensitive, `Ge` is *not* a linear function of the input overlap: it is boosted for the more strongly-activating patterns, which is part of what supports [[stable activation]]. `GeSyn` is the pure net input measure that the equation above describes.

## Manipulating leak

Next, we will explore how we can change how much information is conveyed by the activation signal. We will manipulate the leak current [[#sim-detector:Gbar L]] ($\overline{g}_l$ in [[neuron#eq_gbar-l]]), which has a default value of 240. Note that this is much larger than the standard value of 20 used for a neuron in a normal network, because there is no [[inhibition]] at all in this network, so leak has to do all of the work of counteracting the excitatory input. With this leak, only the strongest (best fitting) input pattern (the `8`) activates the unit. Without a strong leak like this, the excitatory inputs for many of the other input patterns would put the receiving unit above threshold.

**IMPORTANT:** you must press [[#sim-detector:Init]] for changes in `Gbar L` to take effect!

* Reduce the [[#sim-detector:Gbar L]] value from 240 to 220, and do [[#sim-detector:Init]] then [[#sim-detector:Run]], and look at the [[#sim-detector:Test Trial Plot]]. Then try 200, 180 and 300.

{id="question_leak"}
> What happens to the pattern of receiving neuron activity over the different digits when you change `Gbar L` to 220, 200, 180 and 300 -- which input digits does it respond to in each case? In terms of the tug-of-war model between excitation and inhibition & leak, why does changing leak have this effect (a simple one-sentence answer is sufficient)?

{id="question_leak-info"}
> Why might it be beneficial for the neuron to have a lower level of leak (e.g., `Gbar L` = 200 or 180) compared to the original default value, in terms of the overall information that this neuron can convey about the input patterns it is "seeing"?

It is clearly important how responsive the neuron is to its inputs. However, there are tradeoffs associated with different levels of responsivity. The brain solves this kind of problem by using many neurons to code each input, so that some neurons can be more "high threshold" and others can be more "low threshold" types, providing their corresponding advantages and disadvantages in specificity and generality of response. As we will see in [[inhibition]], our tinkering with the value of the leak current is also largely replaced by the inhibitory input, which plays an important role in providing a dynamically adjusted level of inhibition for counteracting the excitatory net input. This ensures that neurons are generally in the right responsivity range for conveying useful information, and it makes each neuron's responsivity dependent on other neurons, which has many important consequences as one can imagine from the above explorations.

* Do [[#sim-detector:Defaults]] to restore the original parameters.

## Links

Next in [[Intro Book]]: [[Neocortex]]

</div>
