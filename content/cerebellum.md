+++
Categories = ["Neuroscience", "Cognition"]
bibfile = "ccnlab.json"
+++

The **cerebellum** is a system of highly specialized neural elements unlike those found in any other part of the brain, the precise function of which has been tantalizingly elusive since the seminal publication of many of its unique features ([[@EcclesItoSzentagothai67]]). Although it is widely considered to be a motor control system, it is anatomically and evolutionarily a sensory system, which untangles intertwined sensory signals to provide a kind of [[linear algebra|orthogonalized basis space]] upon which to drive motor actions.

This untangling function is known as an **adaptive filter** ([[@Fujita82a]]; [[@WidrowStearns85]]), which is a form of [[predictive learning]] that filters out what is expected or predicted based on other sensory signals, so that the _residual_ signal represents a "pure" and otherwise unexplained sensory signal.

When you move any part of your body, especially the head or eyes, you create massive changes across all of your senses. If you did not somehow subtract these **self-motion** sensory signals in your perceptual system, you would never be able to act sensibly in the world ([[@Cullen23]]). Everything would be in a constant state of turbulent motion, and you would never be able to coordinate a proper motor action based on all of these moving targets.

The evolutionary history of the cerebellum is consistent with this adaptive filter function, based on proto-cerebellar circuits and cerebellar analogs in other animals ([[@BodznickMontgomeryCarey99]]; [[@MontgomeryPerks19]]; [[@BellHanSawtell08]]; [[@Cisek21]]). The most ancient part of the mammalian cerebellum is essentially an integral component of the primary brainstem [[vestibular]] sensory nucleus, where it learns to subtract the effects of self-generated motion signals at all levels of the body, to produce a _pure_ vestibular signal that represents any remaining discrepancy from what would be expected based on the self-motion generated signals.

This pure vestibular signal is then something that can be acted upon effectively. If there is an additional sense of motion after all of your own actions have been accounted for, you know that you are slipping or falling or being pushed around in some way. Furthermore, this pure signal tells you exactly which way you should act in order to counteract the residual motion. When combined with other sensory inputs from vision, touch and proprioception, you can also figure out what the likely cause of the disturbance is, and thus condition your motor response accordingly.

{id="figure_env" style="height:20em"}
![The temporal envelope surrounding a descending motor command, with temporal delays offsetting from the start of the command. A cascade of ascending sensorimotor signals from the spinal cord tracks the unfodling of the motor action (reflections of actual muscle signals, proprioceptive muscle spindle fiber signals, etc). The vestibular signal reflects the head and body motion in terms of velocity and acceleration in three dimensions.](media/fig_cerebellum_motor_envelope.png)

As this example makes clear, the adaptive filtering by the cerebellum provides an  essential basis to enable sensible motor actions to be generated. A major challenge that the cerebellum must overcome is managing the many different **temporal delays** associated with the effects of motor actions and the sensory transduction thereof. When you decide to move your head, the **efferent copy** of that **descending motor command** provides the earliest signal about what is going to happen, setting off a complex cascade of subsequent signals that all need to be accounted for ([[#figure_env]]).

Your spinal cord and skeletal muscles do a lot of further work to actually implement the motor command in an efficient way, so you need to also get the actual **ascending motor signals** back up from the lower-level spinal motor pathways and the associated **proprioceptive** signals about the current levels of muscle stretching from the muscle _spindle fibers_. These signals unfold over time as the motor action takes place, even if the descending motor command was relatively ballistic.

Then, as your head starts to move, you will start getting signals from your vestibular system reflecting the motion. The specific spatiotemporal pattern of these signals will depend on the exact motor parameters, so this is a further challenging prediction problem. As shown in [[#figure_env]], the adaptive filtering job performed by the cerebellum can be organized around the entire _temporal envelope_ defined by the initial descending motor command.

{id="figure_cancel" style="height:20em"}
![The prediction of the sensory consequences of a motor action can be subtracted from the original sensory signal to cancel the sensory signal. If there is an external perturbation of the sensory signal that does not fit the prediction, then there is a residual signal that clearly reveals the exact parameters of the perturbation, untangled from the sensory signals generated by the motor command.](media/fig_cerebellum_sensory_cancellation.png)

The core computation performed by the adaptive filter is to subtract the learned prediction of the sensory consequences of a motor action from the actual sensory inputs ([[#figure_cancel]]). If the actual sensory inputs match the learned prediction, then the result is _cancellation_ of the sensory signal. However, if there is a perturbation of the sensory signal from something else unexpectedly going on (a slip, a push, etc), then this will show up clearly as a _residual_ signal that has not been canceled by the prediction.

{id="figure_cereb" style="height:20em"}
![Cerebellar components that implement the adaptive filter computation. In the cerebellar nucleus (CN), grouped inhibitory (CNi) and excitatory (CNe) neurons operate on a specific sensory channel, with sample activity over time as shown. The CNi learns to predict the sensory input activity, and successful prediction results in stable activity over the CNe output (projection) neuron, with excitatory sensory input balanced by corresponding inhibition from both the CNi and the PC (Purkinje cell). The CNe has a high (~40 Hz) tonic baseline firing rate, as does the PC. The IO (inferior olive) compares the CNi prediction vs. the sensory input, and fires a burst of spikes if there is an excitatory input without a corresponding inhibitory input. This leads to complex spike firing in the PC, and can drive plasticity in the CNi.](media/fig_cerebellum_as_adaptive_filter.png)

The key elements of the cerebellar circuitry that can implement this adaptive filter process are shown in [[#figure_cereb]]. There are many other ideas in the literature about what the cerebellum might be doing, which share elements in common with this framework. Indeed, even though each element has been discussed in the literature, the specific details of the current model have not otherwise all been combined into one coherent framework. These issues are discussed below, after providing an overview of the framework.

* Each specific sensory channel (e.g., one particular canal in the vestibular system) has an anatomically organized pathway that specifically learns to predict it. The specific sensory input drives excitatory input into _glomeruli_ (relatively large anatomical structures where multiple axons and dendrites converge) on **inferior olive (IO)** neurons.

    Each IO neuron has about 50 such glomeruli, each of which receives 2 excitatory and 2 inhibitory inputs ([[@DeZeeuwSimpsonHoogenraadEtAl98]]). Each glomerulus connects about 5-6 different nearby IO neurons together, with gap junctions that cause them to be synchronized. Thus, there are about 5-6 IO neurons processing each specific sensory input projection, with likely redundancy in sampling the same specific sensory channel across cerebellar _microzones_.
    
    The error detection logic in the IO glomeruli operates by virtue of inhibitory inputs resetting dendritic conductances, such that excitatory inputs are ineffective within a ~50 ms window, but are strongly amplified outside of that window, so that even a few spikes outside of the window can drive supra-threshold burst activity in the IO neuron ([[@LoyolaHooglandHoedemakerEtAl23]]).

* An associated **cerebellar nucleus (CN)** inhibitory (GABAergic) neuron (_CNi_) projects inhibition into the same IO glomerulus, representing the learned prediction of the specific sensory channel. By detecting the coincidence of excitatory and inhibitory inputs within a narrow time window, the IO glomerulus can compute the _prediction error_ ([[@Ito13]]; [[@DeZeeuwSimpsonHoogenraadEtAl98]]).

* The IO prediction error signal drives learning in the **Purkinje cells (PC)**, which are the primary output neurons of the _cerebellar cortex_.  Each IO neuron sends _climbing fiber (CF)_ axonal output that goes up to the PCs, and also back down to the CN. Each IO neuron sends CF output to about 10 PC neurons, with about 300 synapses per PC neuron ([[@Ito84]]).

    The IO neuron only fires at most two spikes when an error is detected, due to intrinsic currents, but these are sufficient to then drive powerful _complex spike_ activity patterns in the PC neurons that drive learning in these neurons ([[@ItoKano82]]; [[@CoesmansWeberDeZeeuwEtAl04]]). IO neurons also fire single spikes spontaneously at around 1 Hz, but these spontaneous spikes are generally less impactful than the error-driven ones ([[@TitleyKislinSimmonsEtAl19]]; [[@ZangDeSchutter19]]; [[@GuoUusisaari25]]).

* The PC neurons send GABA inhibition to the CN, targeting the CNi and excitatory projection neurons (_CNe_), which convey the primary output of the entire cerebellar system. Both the PC and CNe neurons fire at a high tonic baseline level (roughly 40-50 Hz), due to intrinsic channel properties, which allows for significant bidirectional modulation of their firing rates. Thus, the IO-driven learning in the PC can drive inhibitory inputs onto the CNe neurons that balances out the additional excitatory drive from the sensory input, thereby cancelling the sensory signal.

    The CNi neurons need to learn to predict the specific sensory input in order to make the IO computation an accurate prediction error signal. Consistent with this account, direct evidence shows that this CNi projection is responsible for reducing IO firing and complex spikes in PC neurons as learning proceeds ([[@KimKrupaThompson98]]; [[@Ito01]]). Furthermore, this pathway was responsible for a _blocking_ effect in conditioned learning, which is a signature of the elimination of prediction error signals in a delta-rule like mechanism ([[@Fanselow98]]; [[@KimKrupaThompson98]]). This learning may depend on input from the PC into the CNi neurons ([[@PughRaman08]]).

* The inputs for driving the PC (and other neurons in the cerebellar cortex) and the CNi neurons are conveyed by the **mossy fiber (MF)** pathway originating in the **pontine nuclei (PN)** and other sensory / motor sources. This broad sample of time-varying inputs allows the PC and CNi neurons to generate a prediction of the time-varying specific sensory input.

    The cerebellar cortex also contains a very large number of tiny **granule cells (GC)**, that receive MF inputs and send excitation to the PCs, via distinctive _parallel fiber_ axons, with each PC receiving as many as 150,000 such parallel fiber inputs. The large number of these GCs result in the cerebellum having about 70% of the total number of neurons in the human brain. These GCs, along with inhibitory _Golgi_ cells and other specialized cell types in some areas, are thought to provide a very high-dimensional time-varying representation that is particularly useful for rapidly learning novel temporal firing patterns in the PC neurons ([[@Marr69]]; [[@BuonomanoMauk94]]; [[@MaukBuonomano04]]).

This model shows how the elaborate and specialized nature of the cerebellar circuitry seems ideal for performing the adaptive filter function, consistent with  a major ongoing hypothesis ([[@Fujita82a]]; [[@MiallWolpert96]]; [[@DeanPorrillEkerotEtAl10]]; [[@TanakaIshikawaLeeEtAl20]]). 

The presence of two distinct pathways for canceling the sensory signal, a direct pathway within the CN and an indirect pathway through the cerebellar cortex, raises the question as to why two separate pathways would be required, and how they might perform complementary functions. Taking inspiration from the seminal ideas of David Marr for both the [[hippocampus]] ([[@Marr71]]) and the cerebellum ([[@Marr69]]), one hypothesis is that the use of a very high-dimensional representational space in the cerebellar cortex (via the GCs) allows the system to rapidly learn new spatiotemporal firing patterns without suffering from extensive interference with prior learning.

Furthermore, the CN system can slowly learn the most reliable and stable forms of sensory predictions, while the cerebellar cortex can rapidly learn more context-dependent, idiosyncratic predictions, which may then decay away over time ([[@HerzfeldHallTringidesEtAl20]]; [[@KassardjianTanChungEtAl05]]). There is evidence that the CN pathway does not have as accurate of timing relative to the cortical pathway ([[@GarciaSteeleMauk99]]), which is consistent with the high-dimensional representation formed in the cerebellar cortex. Because each PC integrates input from roughly 150,000 granule cell inputs, it can form novel, arbitrary associations across many more diverse stimulus inputs than the more limited dendritic arbors of CNi neurons.

It is now well established that rapid episodic learning is a major function of the hippocampal circuitry, and [[@^HerzfeldHallTringidesEtAl20]] provide consistent evidence for this idea in the cerebellum. They showed extensive data on the timecourse of learning in monkeys that fits much better with a model having two learning systems, one fast and one slow, versus a single learning system. 

Interestingly, their analysis suggests that the fast learning system (through the cerebellar cortex) should also serve to train the slower system, which is possible via PC inputs onto the CNi prediction neurons, as we discuss next.

## Learning rules

There is biological evidence for [[synaptic plasticity]] in just about every synapse in the cerebellar system ([[@HanselLindenDAngelo01]]; [[@Carey11]]; [[@GaovanBeugenDeZeeuw12]]). The most important synapses based on the logic of [[#figure_cereb]] are:

* CNi predictor inputs.
* PC inputs.

### CNi learning

Starting with the CNi prediction neuron, a simple [[error-backpropagation#delta-rule]] function is optimal from a mathematical perspective:

{id="eq_delta-w" title="Delta rule"}
$$
\Delta w = (s - y) x
$$

where _y_ is the prediction activity, _s_ is the sensory value, and _x_ is the sending activity. 

This equation requires bidirectional learning based on the mismatch between the sensory value and prediction: if the prediction is too low, then the synaptic weights need to be increased (LTP: long-term potentiation), and decreased (LTD: long-term depression) in the opposite case.

If the only driver of learning in the CNi is the error signal conveyed by the activity of the IO via its collateral projection back to the CN, it is difficult to see how bidirectional plasticity could be achieved, because it only fires at most a doublet of two spikes when there is an error ([[@TitleyKislinSimmonsEtAl19]]), and definitely fires much too infrequently in general to support a meaningful bidirectional modulation in its firing rate (even if some other mechanisms might make it somewhat more informative than a purely binary signal; [[@ZangDeSchutter19]]).

Fortunately, there is direct evidence for bidirectional synaptic plasticity in CN neurons ([[@PughRaman08]]; [[@PughRaman06]]; [[@ZhangLinden06]]; [[@McElvainBagnallSakatosEtAl10]]). However, these experiments were conducted in slices, which creates interpretational difficulties. In particular, the _in vitro_ slice preparation exhibits rebound bursting that is likely not present in awake behaving animals (i.e., [[in activo]]), and this bursting was a focus of these studies.

In the absence of definitive data, one possible learning mechanism that could implement [[#eq_delta-w]] is to have the specific sensory input synapse onto the soma (or proximal dendrites) of the CNi neuron, and drive learning based on the contrast between this somatic activity versus that occuring in the dendrites. The results and theorizing summarized by [[@^PughRaman08]] are potentially consistent with this idea, demonstrating distinct contributions from somatic and dendritic activity, based on differential localization of Ca++ channels.

There is a general consensus in the field that learning in the PC neurons is necessary for the _systems consolidation_ of learning in the CN ([[@RaymondMedina18]]). However, each of the relevant studies has issues that leave this point in doubt. Perhaps the strongest-seeming data comes from [[@^GarciaSteeleMauk99]], who lesioned the cerebral cortex post-learning in a conditioning paradigm, and found that this lesion prevented subsequent acquisition to a new CS. However, this is could actually have been due to a blocking effect, as was found by [[@^KimKrupaThompson98]].

Another study by [[@^ZeeuwHanselBianEtAl98]] used a genetic knockout to disable a critical protein kinase involved in [[synaptic plasticity]], selectively in the PC neurons. While they found that this blocked learning in a Visual Optic Reflex (VOR) task in both the PCs and the CN, they only used a short amount of training (1 hr) which was shown in another study to only affect learning in the cerebral cortex ([[@KassardjianTanChungEtAl05]]).

Consistent with the fast vs. slow dissociation, [[@^KassardjianTanChungEtAl05]] only found cortically-independent learning in the CN after 3 days of training. Furthermore, [[@^ZeeuwHanselBianEtAl98]] noted that the standard VOR, and considerable motor behavior, was intact in their knockout mice, providing strong evidence that plasticity in the PC is _not_ important for the slow learning over developmental time in the CN.

Another study by [[@^ShutohOhkiKitazawaEtAl06]] used a range of manipulations, and did show that the inactivation of the IO affected longer-term VOR adaptation, and a blocker of learning-related chemical cascades in the cerebellar cortex impaired learning in the CN.

Overall, it seems reasonable to conclude that the CNi can learn independent of the cerebellar cortex, and that this learning is significantly slower than learning in the PC neurons. The PC projections to CNi neurons may also facilitate learning ([[@HerzfeldHallTringidesEtAl20]]), while the IO collaterals may play a critical modulatory role. This is clearly an area where significant additional research is needed, but the strong computational imperative for [[#eq_delta-w]] provides a strong testable set of predictions in the meantime.

### PC learning

The role of the IO climbing fiber activity in shaping learning in the PC has been much more extensively studied, consistent with the general focus on the cerebellar cortex since the original highly-influential theories of [[@^Marr69]], [[@^Albus71]], and [[@^Ito72]]. Recent data confirms that this climbing fiber input is critical for PC plasticity ([[@SilvaRamirez-BuriticaPritchettEtAl24]]).

If the IO is the driver of PC learning, then the lack of bidirectionality in this signal creates challenges for accurately learning to predict a sensory signal, which again requires something like the bidirectional delta-rule [[#eq_delta-w]]. However, given the evidence suggesting that the PC learning is fast and more transitory in nature, it is reasonable that a unidirectional IO signal indicating the presence of an unexpected sensory signal can be complemented with a slower decay process that prevents all of the synapses from only going in one direction. This is consistent with a number of specific theories (cites).

The specific logic for IO error detection summarized above is asymmetrically organized to detect excess excitatory sensory spiking, so only positive error terms will be signalled by IO spiking, corresponding to LTP. 

<!--- TODO: actual learning in PC more formally -->

* CNe learns too -- could learn to be "flat" based on temporal changes.

## Standard accounts

<!--- TODO -->

In general, much more attention has been paid to the circuits in the cerebellar cortex vs. those in the CN ([[@KebschullCasoniConsalezEtAl24]]; [[@HerzfeldHallTringidesEtAl20]]), and the specific way in which the CN and IO can work together to drive the critical prediction error signal has perhaps been under-appreciated, given how uniquely well-suited its highly unusual biological properties are for this function ([[@DeZeeuwSimpsonHoogenraadEtAl98]]). Indeed, many discussions of cerebellar function assume without much further specificity that the IO conveys some form of externally-generated error signal.

    While most theories and considerable data support the notion of the IO as representing errors, based on the seminal theories of [[@^Marr69]] and [[@^Albus71]], it is not clear that this specific computational account of how the IO glomeruli can directly compute a coincidence-based prediction error signal has been directly articulated. The specific role of the CNi as providing the learned prediction into the IO


There are various other ways in which learning and processing in the cerebellar circuit could be configured, even within the more specific framework assumed here.

* CNi also has the subtraction applied to it, not just CNe. This is actually what most theorists appear to assume. One advantage is that the error signal driving the IO would be simpler: if there is a purturbation in CNi activity of any sort, then that counts as an error, without the need for a fine-grained comparison operation like that proposed for the glomeruli.

    However, that idea is inconsistent with several biological and functional constraints. First, the IO neurons definitely receive direct excitatory projections from the raw sensory signals that the cerebellum is learning to cancel. Therefore, this signal needs to somehow be actively cancelled directly in the IO. The elaborate glomeruli structures seem to perfectly designed for this comparison to not continue with this hypothesis. Second, the CNi neurons would have to cancel themselves, despite neural delays, and this creates significant difficulties.
    
    ? more?



There is extensive research on the latter, since the discovery of LTD at the GC to PC synapses (i.e., the _parallel fibers_) by [[@^ItoSakuraiTongroach82]]. The original logic for this LTD-based learning however does not make sense for the adaptive filter learning logic. The original motor-learning logic is that if an error occurred, whatever the PC was doing must have been incorrect, and therefore it should reduce its input weights.



## Biological details

If there is a mismatch between the excitatory and inhibitory inputs, then the IO neurons will fire a burst of spikes via the **climbing fibers** that wrap around a corresponding **Purkinje cell** (PC) in the cerebellar cortex. This then causes the PC to fire a **complex spike** which drives rapid learning throughout the cerebellar cortex and in the DCN and IO neurons themselves. The only output of the PC, and thus the entire complex circuitry of the cerebellar cortex, is directly onto the DCN neurons, in the form of strong inhibition.

* key data from [[@DeZeeuwBerrebi95]] on PC inputs to CN, and likely independence of CNi vs. CNe -- key prediction and alternative models.

* show data from [[@TanakaIshikawaKakei19]], [[@IshikawaTomatsuIzawaEtAl16]], 2014, [[@ZobeiriCullen24]] on actual firing data in CN, PC neurons etc. DC neurons do not go silent!  every neuron learns what the firing of every other neuron means, so the cancelation is all relative. Weird result from [[@ZobeiriCullen24]] along these lines.

## Other theories of cerebellar function

Brief..

## Hierarchical cascades of predictive controllers

The residual signals provide control knobs for higher levels of control! Key example of VOR vs. saccades etc!  Saccade is an error signal from perspective of VOR. Subsumption / override and neural mechanisms supporting that.

output of lower level is input to next higher level -- higher level learns to predict the residuals in lower level, using broader / higher-level context that explains the perturbations.

Actual motor control signals come from BG and cortex, using sensory signals from cerebellum.

Key Powers et al point: you can do all of motor control in sensory space!

What does cerebellum need to handle this?


