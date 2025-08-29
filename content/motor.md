+++
Categories = ["Neuroscience", "Cognition"]
bibfile = "ccnlab.json"
+++

The **motor** system in the brain involves many different areas, and in many ways, the entire brain exists in service of producing motor output. For example, Daniel Wolpert has made the point that the sea squirt eats its own brain once it no longer needs to move around. More specific discussion of various aspects of motor control can be found in [[reinforcement learning]], [[basal ganglia]], [[cerebellum]], and the [[Rubicon]] framework. This page provides a high-level overview and fills in a few missing pieces.

For a more detailed treatment from experts in this area, see [[@^ArberCosta22]].

## Frontal cortex learning via subcortical efferent copy signals

Unlike most of the rest of the brain, the [[neocortex]] is thought to gain almost all of its functionality through learning taking place during an animal's lifetime. This poses an important challenge for the motor system: how can the motor cortical areas learn to "speak the language" of the spinal cord, to provide meaningful motor control signals?

Our hypothesis is that the same mechanisms used in [[predictive learning]] of sensory and other representations in the cortex, supported by the pulvinar and MD nuclei in the [[thalamus]], could also be at work on the motor side in frontal cortex areas. As reviewed in the thalamus page, the VL (ventrolateral) nucleus receives strong focal driver inputs from the deep cerebellar nuclei (DCN), which in turn receives extensive ascending projections from the spinal motor pathways.

There are natural delays between the time when the frontal cortex drives a descending motor command and the time when the efferent copy signals from these ascending pathways reflect the actual low-level motor commands that were actually generated (with significant contributions and shaping by a number of subcortical areas including the BG, cerebellum, and midbrain & brainstem lower motor areas). This delay provides the [[temporal derivative]] dynamics needed to drive predictive learning, where the motor cortex is effectively learning to predict what the motor system will do in response to the commands it sends down.

{id="figure_big-loop" style="height:30em"}
![The "big loop" of motor processing, which could support a form of predictive learning to train frontal cortex. The thalamus (VL = ventrolateral specifically) receives strong driver inputs from the DCN, and standard non-driver inputs from motor cortex. The difference over time between these two inputs can drive the temporal derivative error-driven learning mechanism. From Arber & Costa, 2022](media/fig_motor_big_loop_arber_costa_22.png)

This hypothesis is consistent with the "big loop" of descending and ascending motor signals shown in [[#figure_big-loop]] from [[@ArberCosta22]].

