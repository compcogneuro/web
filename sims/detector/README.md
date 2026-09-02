*[Back to simulations](https://compcogneuro.org/simulations)*

# Introduction

This simulation shows how an individual neuron can act like a detector, picking out specific patterns from its inputs and responding with varying degrees of selectivity to the match between its synaptic weights and the input activity pattern.

We will see how a particular pattern of weights makes a simulated neuron respond more to some input patterns than others. By adjusting the level of excitability of the neuron, we can make the neuron respond only to the pattern that best fits its weights, or in a more graded manner to other patterns that are close to its weight pattern. This provides some insight into why the spiking neuron activation function works the way it does.

# The Network and Input Patterns

We begin by examining the `Network` tab, showing the Detector network. The network has an `Input` layer that will have patterns of activation in the shape of different digits, and these input neurons are connected to the receiving neuron (`RecvNeuron`) via a set of weighted synaptic connections. We can view the pattern of weights (synaptic strengths) that this receiving unit has from the input, which should give us an idea about what this unit will detect.

* Select the `Wts` tab at the top of the list of network variables at the left of the Network view, then click `r.Wt` as the value you want to display, and then click on the `RecvNeuron` to view its receiving weights.

You should now see the `Input` grid lit up in the pattern of an `8`. This is the weight pattern for the receiving unit for connections from the input units, with the weight value displayed in the corresponding sending (input) unit. Thus, when the input units have an activation pattern that matches this weight pattern, the receiving unit will be maximally activated. Input patterns that are close to the target `8` input will produce graded activations as a function of how close they are. Thus, this pattern of weights determines what the unit detects, as we will see. First, we will examine the patterns of inputs that will be presented to the network.

* Click on the `Inputs/Digits` item in the data browser on the left side of the window to see all of the input patterns.

The display that comes up shows all of the different *input patterns* that will be presented ("clamped") onto the `Input` layer, so we can see how the receiving unit responds. Each row of the display represents a single *trial* that will be presented to the network. As you can see, the input data in this case contains the digits from 0 to 9, represented in a simple font on a 5x7 grid of pixels (picture elements). Each pixel in a given event (digit) will drive the corresponding input unit in the network.

# Running the Network

To see the receiving neuron respond to these input patterns, we will present them one-by-one, and determine why the neuron responds as it does given its weights. Thus, we need to view the activations again in the network window.

* Select the `Act` tab in the network variables and click `Act` to view activations, then click the `Step` button in the toolbar at the top of the window, which will step one `Trial` as indicated.

This activates the pattern of a `0` (zero) in the `Input`, and shows the 200 cycles (milliseconds) of **settling** that make up one **theta cycle** trial, during which the state of the receiving neuron is iteratively updated according to the spiking neuron equations (just as the unit in the `neuron` simulation was updated over time).

The receiving unit showed an activity value of 0 because it never got activated above its firing threshold by the `0` input pattern. Before getting into the nitty-gritty of why the unit responded this way, let's proceed through the remaining digits and observe how it responds to other inputs.

* Press `Step` for each of the other digits, until the number `8` shows up.

You should have seen the receiving unit finally activated when the digit `8` was presented, with an activation of zero for all the other digits. Thus, as expected, the receiving unit acts like an `8` detector: only when the input perfectly matches the input weights is there enough excitatory input to drive the receiving neuron above its firing threshold.

* You can use the "VCR" style buttons at the bottom of the `Network` to review each cycle of updating, to see the progression of activation over time. Selecting `Raster` in the Network toolbar, and viewing the `Spike` variable, gives a nice picture of the individual spikes over time.

* Go ahead and do one more `Step` to see what happens with `9`.

We can use a graph to view the pattern of receiving unit activation across the different input patterns.

* Click on the `Test Trial Plot` tab.

The graph shows the excitatory net input (`GeSyn`) and the activation (`Act`) for the unit as a function of trial (and digit) number along the X axis. For `Act` you should see a flat line with a single peak at 8.

# Computing Excitatory Conductance `GeSyn` (Net Input)

Now, let's try to understand exactly why the unit responds as it does. The key to doing so is to understand the relationship between the pattern of weights and the input pattern.

* Go back to the `Wts`/`r.Wt` display in the Network, then open the `Inputs/Digits` patterns again, and make sure you can see the weights as well. The idea is to compare the weights with the digit patterns.

> **Question 2.8:** For each digit pattern, report the number of active units in the pattern where there is also a weight of 1 according to the `8` digit pattern shown in the `r.Wt` view in the Network. In other words, report the *overlap* between the digit input activity and the weight pattern.

<sim-question id="2.8">

The number of inputs having a weight of 1 that you just calculated should correspond to the total excitatory synaptic input `GeSyn`, also called the **net input**, going into the receiving unit, which is a function of the average of the sending activation `Act` times the weight `Wt` over all the units, with a correction factor for the expected activity level in the layer, `Alpha`:

```
GeSyn = (1 / Alpha) * [Sum(Act * Wt) / N]
```

You can select the `Gmisc` tab in the network variables and click `GeSyn` to see these values as you step through the inputs, and it is also plotted in the `Test Trial Plot`.

* Do `Init` and `Step` to see the `0` input again. If you hover over the `RecvNeuron` with your mouse, you should see it has a value of `GeSyn = .3529..`. To apply the above equation, you should have observed that `0` has 6 units in common with `8`, and `N` = 35 (7*5), so that is about .1714. Next, we need to apply the `Alpha` correction factor, which we set to be the activity level of the `8`, which is 17 of the 35 units active. Thus, we should get:

```
GeSyn = (1 / (17 / 35)) * (6 / 35) = 6 / 17 = .3529...
```

which is exactly what the model shows. The two factors of `N` = 35 cancel out, so the net input reduces to just the *overlap divided by 17*, the number of active units in the weight pattern. Thus the `8` itself, which matches the weights perfectly, gives `GeSyn = 17 / 17 = 1`, and you can multiply any of these `GeSyn` values by 17 to read off the overlap directly, confirming your answers to Question 2.8.

As a result of working through the `GeSyn` net input calculation, you should now have a detailed understanding of how the net excitatory input to the neuron reflects the degree of match between the input pattern and the weights. You have also observed how the activation value can ignore much of the graded information present in this input signal, due to the presence of the **threshold**. This gives you a good sense for *why* neurons have these thresholds: it allows them to filter out all the "sub-threshold noise" and only communicate a clear, easily interpreted signal when it has detected what it is looking for.

Note that the `Ge` variable, which is also available in the network and the plot, is the *total* excitatory conductance, which includes the additional NMDA and VGCC channel contributions on top of the AMPA synaptic conductance in `GeSyn`. Because the NMDA channel is voltage sensitive, `Ge` is *not* a linear function of the input overlap: it is boosted for the more strongly-activating patterns. `GeSyn` is the pure net input measure.

# Manipulating Leak

Next, we will explore how we can change how much information is conveyed by the activation signal. We will manipulate the leak current (`GbarL`), which has a default value of 240. Note that this is much larger than the standard value of 20 used for a neuron in a normal network, because there is no inhibition at all in this network, so leak has to do all of the work of counteracting the excitatory input. With this leak, only the strongest (best fitting) input pattern (the 8) activates the unit. Without a strong leak like this, the excitatory inputs for many of the other input patterns would put the receiving unit above threshold.

**IMPORTANT:** you must press `Init` for changes in `GbarL` to take effect!

* Reduce the `Gbar L` value from 240 to 220, and do `Init` then `Run` and look at the `Test Trial Plot`. Then try 200, 180 and 300.

> **Question 2.9:** What happens to the pattern of receiving neuron activity over the different digits when you change GbarL to 220, 200, 180 and 300 -- which input digits does it respond to for each case? In terms of the tug-of-war model between excitatory and inhibition & leak (i.e., GbarL = leak), why does changing leak have this effect (a simple one-sentence answer is sufficient)?

<sim-question id="2.9">

> **Question 2.10:** Why might it be beneficial for the neuron to have a lower level of leak (e.g., GbarL = 200 or 180) compared to the original default value, in terms of the overall information that this neuron can convey about the input patterns it is "seeing"?

<sim-question id="2.10">

It is clearly important how responsive the neuron is to its inputs. However, there are tradeoffs associated with different levels of responsivity. The brain solves this kind of problem by using many neurons to code each input, so that some neurons can be more "high threshold" and others can be more "low threshold" types, providing their corresponding advantages and disadvantages in specificity and generality of response. As we will see in the next chapter, our tinkering with the value of the leak current `GbarL` is also largely replaced by the inhibitory input, which plays an important role in providing a dynamically adjusted level of inhibition for counteracting the excitatory net input. This ensures that neurons are generally in the right responsivity range for conveying useful information, and it makes each neuron's responsivity dependent on other neurons, which has many important consequences as one can imagine from the above explorations.
