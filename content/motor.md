+++
Categories = ["Neuroscience", "Cognition"]
bibfile = "ccnlab.json"
+++

The **motor** system in the brain involves many different areas, and in many ways, the entire brain exists in service of producing motor output. For example, Daniel Wolpert has made the point that the sea squirt eats its own brain once it no longer needs to move around. More specific discussion of various aspects of motor control can be found in [[reinforcement learning]], [[basal ganglia]], [[cerebellum]], and the [[Rubicon]] framework. This page provides a high-level overview of the challenges and solutions for motor control, and some details about the underlying physiology of muscles and how the spinal cord and brainstem provide a systematic basis for motor control.

<!--- For a more detailed treatment from experts in this area, see [[@^ArberCosta22]]. -->

## Dimensionality reduction, coordination, and muscle synergies

As in all aspects of neural function, the central problem in motor control is managing the [[curse of dimensionality]]: there are exponentially many possible combinations of muscle activations over time. How does the brain reduce this huge space down to the relatively small subset of muscle activations necessary for an animal's survival?

This is another instance of the fundamental [[search]] problem, and in the case of the motor system, [[evolution]] has done a lot of the work by building in complex **muscle synergies** into the spinal cord and brainstem, where individual **interneurons** (excitatory and inhibitory) activate **spatiotemporal patterns** of muscle activation. From this lower-dimensional _repertoire_ of elements (i.e., a set of [[linear algebra#basis space]]), complex patterns of motor behavior are constructed ([[@Bernstein67]]; [[@Bernstein96]]; [[@TreschSaltielBizzi99]]; [[@dAvellaSaltielBizzi03]]; [[@TingMcKay07]]; [[@LatashLevinScholzEtAl10]]; [[@BrutonODwyer18]]; [[@Latash18]]).

Another critical problem that is solved by these muscle synergies is the **coordination** across many different muscles that is required to accomplish any given motor action. The contraction of any given muscle creates a variety of physical consequences for other muscles, including altering the basic center of gravity of the entire organism. Thus, any given action must take all of these consequences into account, and ensure that all of the individual muscle contractions are indeed synergistic, and not working at cross purposes.

A variety of different terms have been used in the literature to refer to these synergies, including **reflexes** ([[@Sherrington10]]), **central pattern generators** ([[@GrillnerElManira20]]; [[@GrillnerZangger79]]), and _force fields_ ([[@GiszterMussa-IvaldiBizzi93]]).

{id="figure_synergies" style="height:30em"}
![Three muscle synergies extracted from frog kicking behavior. a) Activation of different muscles (vertical axis) over time (horizontal axis), accomplishing given overall motor action (moment arms, aligned by corresponding column order of the synergies): _HE_ hip extension; _HF_ hip flexion; _KE_ knee extension; _KF_ knee flexion; _AE_ ankle extension; _AF_ ankle flexion. b) Activation strengths over time for each synergy component ($C_{1-3}$) needed to reconstruct 6 different overall motor actions. These patterns differ primarily in magnitude, not temporal pattern. From d'Avella et al., 2003](media/fig_muscle_synergies_davella_etal_03.png)

As a concrete example, a detailed analysis of spinal-level muscle synergies in the frog leg system revealed a set of 4 synergies that could be combined with different activation strengths to explain a wide range of overall motor response patterns ([[@TreschSaltielBizzi99]]). A subsequent analysis allowing for contributions from the entire brain in intact frogs showed how the timing and activation modulation of 3 different spatiotemporal muscle synergies ([[#figure_synergies]]) can explain this same space of motor responses ([[@dAvellaSaltielBizzi03]]). 

Thus, the final motor behavior pattern depends on the integrated contributions of multiple levels of control, with the spinal muscle synergies implemented by interneurons providing the lowest level, and higher levels progressively adding their own direct synergies as well as broader coordination across the basic spinal elements ([[@GrillnerElManira20]]). For example, an analysis of correlated motor units in hand movement showed _last order_ (i.e., just upstream of the muscle activation) inputs from the pontomedullary reticular formation, the magnocellular red nucleus, and primary motor cortex (M1), in addition to the basic spinal cord interneurons ([[@XuMawaseSchieber24]]).

## Feedback control across levels

The other critical element of dimensionality reduction in the motor system arises from the interactions with the environment, which is also highly variable and thus represents a major additional source of variance and combinatorial explosion. There is considerable evidence that the muscle synergies implemented by spinal cord interneurons incorporate direct feedback control mechanisms, that automatically compensate for environmental perturbations ([[@WimalasenaPandarinathAuYong25]]; [[@ConwayHultbornKiehn87]]; [[@AngelGuertinJimenezEtAl96]]; [[@AlvarezFyffe07]]). 

Unfortunately, the ability to understand these feedback control mechanisms has been impaired by the difficulty in categorizing and mapping the connectivity of the spinal interneurons ([[@SenguptaBagnall23]]). Thus, extant circuit-based models tend to exclude such mechanisms ([[@RybakDoughertyShevtsova15]]; [[@McCreaRybak08]]). Nevertheless, more abstract state-space analyses of large populations of spinal interneurons have the potential to reveal the presence and dynamical implications of these feedback control mechanisms ([[@WimalasenaPandarinathAuYong25]]).

The rather underdeveloped understanding of feedback control in motor systems contrasts with the theoretical clarity of the mathematically-based hierarchical control known as **perceptual control theory (PCT)** ([[@Powers73]]; [[@Powers73a]]; [[@Cools85]]; [[@Yin14a]]; [[@BarterYin21]]). This framework is consistent with the principle of **equilibrium-point** motor control, which postulates that motor control signals specify a target length for each muscle, rather than dynamical variables such as force ([[@FeldmanLevin09]]; [[@GribbleOstrySanguinetiEtAl98]]).

While the simplest form of the equilibrium-point hypothesis is likely incorrect, the available evidence strongly suggests that muscle synergies incorporate their own internal feedback control dynamics that make them intrinsically robust, and adaptive, in the face of environmental perturbations.

## Frontal cortex learning via subcortical efferent copy signals

Unlike most of the rest of the brain, the [[neocortex]] is thought to gain almost all of its functionality through learning taking place during an animal's lifetime. This poses an important challenge for the motor system: how can the motor cortical areas learn to "speak the language" of the spinal cord, to provide meaningful motor control signals?

Our hypothesis is that the same mechanisms used in [[predictive learning]] of sensory and other representations in the cortex, supported by the pulvinar and MD nuclei in the [[thalamus]], could also be at work on the motor side in frontal cortex areas. As reviewed in the thalamus page, the VL (ventrolateral) nucleus receives strong focal driver inputs from the deep cerebellar nuclei (DCN), which in turn receives extensive ascending projections from the spinal motor pathways.

There are natural delays between the time when the frontal cortex drives a descending motor command and the time when the efferent copy signals from these ascending pathways reflect the actual low-level motor commands that were actually generated (with significant contributions and shaping by a number of subcortical areas including the BG, cerebellum, and midbrain & brainstem lower motor areas). This delay provides the [[temporal derivative]] dynamics needed to drive predictive learning, where the motor cortex is effectively learning to predict what the motor system will do in response to the commands it sends down.

{id="figure_big-loop" style="height:30em"}
![The "big loop" of motor processing, which could support a form of predictive learning to train frontal cortex. The thalamus (VL = ventrolateral specifically) receives strong driver inputs from the DCN, and standard non-driver inputs from motor cortex. The difference over time between these two inputs can drive the temporal derivative error-driven learning mechanism. From Arber & Costa, 2022](media/fig_motor_big_loop_arber_costa_22.png)

This hypothesis is consistent with the "big loop" of descending and ascending motor signals shown in [[#figure_big-loop]] from [[@ArberCosta22]].

## Hierarchical cascades of predictive controllers

The residual signals provide control knobs for higher levels of control! Key example of VOR vs. saccades etc!  Saccade is an error signal from perspective of VOR. Subsumption / override and neural mechanisms supporting that.

output of lower level is input to next higher level -- higher level learns to predict the residuals in lower level, using broader / higher-level context that explains the perturbations.

Actual motor control signals come from BG and cortex, using sensory signals from cerebellum.

Key Powers et al point: you can do all of motor control in sensory space!

What does cerebellum need to handle this?

## Descending muscle control signals

_Size principle_ [[@Henneman85]]; [[@LucaErim94]]: orderly activation of small to large according to total force as signalled by increasingly "strong" descending activity -- otherwise too high-dimensional. Smaller units have lower threshold, and fire at higher rates than larger units. This also optimizes fatigue effects. The size principle entails the _common-drive_ hypothesis, which is that a single unidimensional  descending control signal projects to all motor units for a given overall muscle, with the specific units activated according to the strength of the drive signal.

This characterizes general story, but in primates, there is more flexible control [[@MarshallGlaserTrautmannEtAl22]], associated also with a larger number of secondary motor areas in primates that have more focal, specialized projections to specific limbs etc [[@StrickDumRathelot21]].

Population coding of cortical neurons: [[@AflaloGraziano06a]] show that final hand posture accounts for most of the variance, but other factors such as speed, curvature of space, distance and force were also coded. This postural aspect, otherwise known as a _spatial synergy_ ([[@OverduindAvellaRohEtAl15]]), provides a simple model of muscle control where, regardless of the starting muscle configuration, the control signal specifies a _final configuration_ (i.e., pattern of total contraction) across all of the relevant muscles.

Where more complex sequences of motor actions are required, a temporal sequence of postural configurations is specified -- a sequence of "poses" -- otherwise known as a _spatiotemporal synergy ([[@OverduindAvellaRohEtAl15]])

[[@^MeyerSmithWright82]] synthesize psychophysical literature on speed and accuracy of motor movements, to develop a symmetric impulse control model that specifies the force and duration parameters as curves with an initial acceleration phase for the first half, followed by a symmetric deceleration phase in the second half. Both the force and time parameters of these curves can be controlled by people. There is evidence that ballistic movements are made below around 260 ms, with multiple iterations of visually-corrected movement updates happening after that, time permitting. See also [[@MeyerSmithKornblumEtAl90]].


