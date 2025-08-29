+++
Categories = ["Neuroscience", "Cognition"]
bibfile = "ccnlab.json"
+++

The **cerebellum** is a system of highly specialized neural elements unlike those found in any other part of the brain, the precise function of which has been tantalizingly elusive. Although is is widely considered to be a motor control system, our model of it is actually almost entirely sensory in nature: it is a system that untangles intertwined sensory signals, to provide a kind of [[linear algebra|orthogonalized basis space]] upon which to drive motor actions.

This function is known as an **adaptive filter** (e.g., a _Kahlman_ filter; [[@LaurensAngelaki18]]), which is a form of [[predictive learning]] that filters out what is expected or predicted based on other sensory signals, so that the _residual_ signal represents a "pure" and otherwise unexplained sensory signal.

When you move any part of your body, especially the head or eyes, you create massive changes across all of your senses. If you did not somehow subtract these **self-motion** sensory signals in your perceptual system, you would never be able to act sensibly in the world ([[@Cullen23]]). Everything would be in a constant state of turbulent motion, and you'd never be able to coordinate a proper motor action based on all of these moving targets.

The evolutionary history of the cerebellum is consistent with this adaptive filter function, based on proto-cerebellar circuits and cerebellar analogs in other animals ([[@BodznickMontgomeryCarey99]]; [[@MontgomeryPerks19]]; [[@BellHanSawtell08]]; [[@Cisek21]]). The most ancient part of the mammalian cerebellum is essentially an integral component of the primary brainstem [[vestibular]] sensory nucleus, where it learns to subtract the effects of self-generated motion signals at all levels of the body, to produce a _pure_ vestibular signal that represents any remaining discrepancy from what would be expected based on the self-motion generated signals.

This pure vestibular signal is then something that can be acted upon effectively. If there is an additional sense of motion after all of your own actions have been accounted for, you know that you are slipping or falling or being pushed around in some way. Furthermore, this pure signal tells you exactly which way you should act in order to counteract the residual motion. When combined with other sensory inputs from vision, touch and proprioception, you can also figure out what the likely cause of the disturbance is, and thus condition your motor response accordingly.

<!--- TODO: timeline figure of the chain of events -->

As this example makes clear, the adaptive filtering by the cerebellum provides an  essential basis to enable sensible motor actions to be generated. This job requires dealing with different **temporal delays** associated with the effects of motor actions and the sensory transduction thereof. When you decide to move your head, the **efferent copy** of that **descending motor command** provides the earliest signal about what is going to happen, setting off a complex cascade of subsequent signals that all need to be accounted for (figure X).

Your spinal cord and skeletal muscles do a lot of further work to actually implement the motor command in an efficient way, so you need to also get the actual **ascending motor signals** back up from the lower-level spinal motor pathways and the associated **proprioceptive** signals about the current levels of muscle stretching from the muscle _spindle fibers_, etc. These signals unfold over time as the motor action takes place, even if the descending motor command was relatively ballistic. As you learn how your body actually works, you should be able to accurately predict all of these signals, and any discrepancies thus represent unexplained disturbances that are thus available as signals to drive further actions.

Then, as your head starts to move, you will start getting signals from your vestibular system reflecting the motion. The specific spatiotemporal pattern of these signals will depend on the exact motor parameters, so this is a further challenging prediction problem.

As shown in figure X, the adaptive filtering job performed by the cerebellum can be organized around the entire _temporal envelope_ defined by the initial descending motor command. Interestingly, one of the most specialized elements of the cerebellar circuitry, the neurons of the **inferior olivary nucleus** (IO), have specialized mechanisms that are perfectly configured to group together a chunk of cerebellar hardware within a roughly 150-250 msec time window, which should encompass the entire envelope around an individual motor action.

Specifically the IO neurons have intrinsic oscillatory dynamics that coordinate their firing over time, and electrical gap junction connections that cause these dynamics to be synchronized across a population of IO cells. Furthermore, the IO neurons have specialized _glomeruli_ on their dendrites that organize a combination of excitatory and inhibitory inputs (specifically 2 sets of each), which appear ideally configured to perform the critical subtraction at the heart of adaptive filtering.

<!--- TODO: IO specific figure showing subtraction -->

The inhibitory inputs come directly from **deep cerebellar nuclei** (DCN) neurons, and represent the _prediction_, while the excitatory inputs come from a combination of other DCN excitatory inputs and sensory inputs, and represent the actual _outcome_ in the [[predictive learning]] framework.

If there is a mismatch between the excitatory and inhibitory inputs, then the IO neurons will fire a burst of spikes via the **climbing fibers** that wrap around a corresponding **Purkinje cell** (PC) in the cerebellar cortex. This then causes the PC to fire a **complex spike** which drives rapid learning throughout the cerebellar cortex and in the DCN and IO neurons themselves. The only output of the PC, and thus the entire complex circuitry of the cerebellar cortex, is directly onto the DCN neurons, in the form of strong inhibition.

Critically, the DCN and IO elements of the system are really the heart of what the cerebellum does: the DCN neurons provide the entire "visible" output of the cerebellum that projects out to the rest of the brain. These DCN neurons receive the various sensory signals shown in Figure-X, and learn to turn these sensory signals into accurate predictions, that then drive the inhibitory input to the IO neurons.

<!--- (but how!? they can't themselves impose delays. what are they actually doing!?).  -->

It contains more than half of the neurons in the brain, but most of those are tiny granule cells, which David Marr recognized as potentially having a similar function to the granule cells in the dentate gyrus of the [[hippocampus]] ([[@Marr69]]; [[@Marr71]]). That part of the picture has stood the test of time: the cerebellum and the hippocampus are both learning systems that take advantage of very high-dimensional spaces to quickly memorize things.

