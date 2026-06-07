+++
Categories = ["Neuroscience"]
bibfile = "ccnlab.json"
+++

The ultimate goal of the [[Axon]] project is to understand the function of the mammalian **brain**. As [[#figure_bauplan]] shows, the overall configuration and elements of the brain (_bauplan_) are already present in jawed fishes, which diverged from the mammalian line roughly 500 million years ago in [[evolution]]ary time. This perspective helps one appreciate that the brain is fundamentally organized to use sensory signals (which tend to be in the dorsal half of the neural tube) to drive adaptive [[motor]] behavior, conditioned on various internal states (hunger, etc). These internal [[emotion]]al drives are located in the anterior, ventral portion of the neural tube, as are the various coordinating motor output pathways.

{id="figure_bauplan" style="height:40em"}
![The overall organization and major components of the mammalian brain were present in the jawed fishes (_gnathostome_) which diverged from the mammalian evolutionary line roughly 500 million years ago. **A** shows the major components of the gnathostome, from [[@^Cisek21]]. Cbm = cerebellum; Str / Pd / SNr = basal ganglia; Tectum = superior colliculus in mammals; VLPall, MPall are pallidum that evolves into the neocortex in mammals. **B** shows the major components of the primate brain, which retains all the major subcortical systems (basal ganglia and cerebellum most importantly), and adds the neocortex, which is critical for the adaptive and flexible nature of human thought. From [[@^GrillnerElManira20]].](media/fig_evolution_brain_template_gnathostome_vs_primate.png)

The major components of the brain present in this evolutionary bauplan are:

* The [[motor#Spinal cord]] contains extensive motor-control circuits that directly incorporate online sensory information and provide a lower-dimensional space for driving adaptive motor control by higher areas. The midbrain motor control areas provide additional higher levels of control, consistent with the [[subsumption]] architecture.

* Various _motivational_ and _regulatory_ systems, including the [[hypothalamus]], pituitary, and [[amygdala]], which monitor the overall state of the organism and drive behavior to satisfy current needs.

* The [[cerebellum]], which is in the dorsal, sensory portion, and is thus primarily a massive parallel processor of sensory signals, which facilitates adaptive motor control through [[predictive learning]]. Specifically, it learns to anticipate the sensory consequences of motor actions, and to thus accomplish two key functions: _adaptive filtering_ which removes these self-generated sensory signals, to provide a more _stable_ and _invariant_ sensory world on which to plan actions; and _forward models_ which anticipate the sensory consequences far enough in advance to enable other motor actions to respond appropriately, for example to drive eye movements that automatically compensate for head movements (i.e., the [[cerebellum#vestibulo-ocular reflex]]).

* The [[Basal ganglia]] which learn to control motor actions based on [[dopamine]] neuromodulatory signals that reflect deviations from expected reward outcomes, as captured in [[reinforcement learning]] algorithms. 

* The **Tectum**, which is the [[superior colliculus]], in mammals, which has a number of evolutionarily-shaped sensory-motor behavior mappings, e.g., for which stimuli to approach vs avoid. This is where may _instinctual_ forms of behavior originate.

* The **Pallium** which becomes the [[neocortex]] in mammals, including older parts thereof, such as the [[hippocampus]] and the olfactory bulb. This provides a higher level of sensory-motor learning, as a complement to the relatively fixed behavior patterns generated in the tectum and lower areas. This high level of plasticity is nevertheless grounded by extensive training signals from subcortical areas (via the [[thalamus]]), which teach the neocortex to communicate effectively with the subcortical areas, and more generally provide critical learning and control signals.

[[Axon]] has separate models of each of these components, along with more complex models that integrate multiple interacting elements to accomplish robust, adaptive behavior. The [[Rubicon]] framework provides an outer-loop of control of all of these elements, based on ventral and medial [[prefrontal cortex]] areas important for goal-driven, top-down control to satisfy active needs.

