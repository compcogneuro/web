+++
Categories = ["Neuroscience"]
bibfile = "ccnlab.json"
+++

The **thalamus** plays essential roles in multiple different brain circuits, and is thus discussed mostly in the context of those other circuits. This page provides a big-picture organizational view of the thalamus, with pointers to those other pages.

From an evolutionary, comparative perspective, the thalamus consistently provides interconnectivity between primary sensory areas and the _pallium_, which is the outer-most region of the brain, i.e., the _telecephelon_. In mammals, the [[neocortex]] is a major component of the pallium, while other vertibrates have smaller, simpler cortex-like structures in their pallium that play the same functional role as the neocortex ([[@Butler08]]). The functions highlighted below are specific to mammals.

* In posterior [[neocortex]], the _pulvinar_ nucleus of the thalamus supports [[predictive learning]] by alternatively representing a top-down prediction and a bottom-up outcome.

* There are multiple different thalamic nuclei interconnected with frontal [[neocortex]], particularly [[prefrontal cortex]] and the [[basal ganglia]], which play critical roles in goal-driven behavior as captured in the [[Rubicon]] model.

<!--- TODO: reunions and hippocampus, etc. -->


## Thalamic circuitry

{id="figure_pulv-conns" style="height:15em"}
![Connectivity between the neocortex and the pulvinar nucleus of the thalamus, in the case of primary and secondary visual areas. Visual input from the retina provides driver inputs to the lateral geniculate nucleus (LGN) of the thalamus, which then sends glutamatergic excitatory projections primarily into layer 4 of primary visual cortex (V1). The pulvinar thalamic nucleus receives strong driver inputs from deep layer 5b neurons in V1, which can have as few as one driver synapse per thalamic relay neuron (TRN). There are numerous and relatively weaker projections from layer 6 (VI) neurons in secondary visual cortex (V2). Diagram based on Sherman & Guillery (2006).](media/fig_pulvinar_connectivity.png)

[[#figure_pulv-conns]], based on [[@ShermanGuillery06]], shows the thalamic connectivity for primary and secondary visual cortex in the mammalian brain. The main features of this connectivity are common across most thalamic areas.

There is only one predominant type of thalamic neuron in the different thalamic nuclei, the _thalamic relay neuron_ (TRN). The _thalamic reticular nucleus_ (TRN), which is an outer shell wrapping around the body of the thalamus, is the one exception, having GABAergic [[inhibition|inhibitory]] neurons, which we discuss below.  TRNs releases excitatory glutamate neurotransmitter into the [[neocortex]], and receive two qualitatively distinct types of inputs:

* **Drivers** are strong inputs, often relatively few in number (as few as a single input). They can have large extended _glomeruli_ with multiple synaptic connections encapsulated in a single anatomical structure.

* **Non-drivers** are the rest of the inputs, which are like most other synaptic connections in the brain. [[@ShermanGuillery06]] and many others label these as _modulatory_ inputs, but that is perhaps too strong of a functional assumption, given their proposed function in the context of [[predictive learning]]. For example, there is ample evidence that these inputs can drive sustained firing of TRNs, as reviewed in [[@OReillyRussinZolfagharEtAl21]].

The TRNs notably lack any local excitatory interconnectivity, and only have a form of pooled inhibition mediated by the TRN reticular neurons. Therefore, they are uniquely suited for providing a faithful representation of the inputs they receive, serving as a kind of _projection screen_ that does not impose significant distortion of its own.

Although the relatively inert nature of the thalamic circuitry has led to its characterization as a _relay_ system, the presence of two distinctive types of inputs suggests a more powerful kind of function could be be supported. The [[predictive learning]] framework for example shows how the distinctive temporal aspects of the driver inputs to the pulvinar supports a [[temporal derivative]] computation that effectively computes an error gradient between the driver and non-driver inputs, across time.

The nature of the connectivity from TRNs into the cortex provides another important insight into the function of the thalamus. There are two primary dimensions along which these projections vary, which were originally captured by the **core** vs. **matrix** categories of [[@^Jones98]], but which are more accurately categorized as dimensions ([[@ClascaRubio-GarridoJabaudon12]]; [[@ShermanUsrey24]]):

* Laminar targets: the classical _core_ type projects mainly to the middle layers of the cortex, primarily layer 4 (e.g., as shown in [[#figure_pulv-conns]]), while the classical _matrix_ type projects to layers 1 and 5, but avoids the middle layers. However, there are actually different proportions of these connections across different thalamic neurons, rather than sharply defined categories.

* Area focus: the classical _core_ type projects to only one cortical area, typically in a relatively circumscribed, _focal_ manner, while the classical _matrix_ type projects to multiple different areas.

{id="figure_thal-types" style="height:30em"}
![Different types of thalamic projections into the cortex, varying in the laminar targets and the area focus, according to Clasca et al., 2012. The C-type is the classical first-order sensory core-type, projecting to the middle lamina in a focal manner, while the M-type (multiareal) is the classical matrix type projecting to layer 1 and 5, across multiple areas. The M-type (focal) combines features from classical core and matrix types, with projections that target middle layers in addition to layer 1, with more focal area-wise connectivity. Meanwhile the IL (intralaminar) pattern (i.e., PF = parafascicular thalamus) is specific to neurons in these nuclei, targeting only deep layers and also extensively the striatum.](media/fig_thalamus_types_clasca_etal_12.png)

[[#figure_thal-types]] shows an extension of the classical core and matrix categories by [[@^ClascaRubio-GarridoJabaudon12]], with two distinct M-types, one of which corresponds to the classical matrix type, while the M-type (focal) combines features of core and matrix.

## Functional organization of the thalamus

{id="figure_thal-clasca" style="height:30em"}
![Map of thalamic areas in the rodent according to the four types shown in the prior figure, using the same colors. From Clasca et al., 2012.](media/fig_thalamus_areas_clasca_etal_12.png)

{id="figure_thal-rovo" style="height:50em"}
![Different types of thalamic drivers in the primate thalamus, which provides key insight into functions of these thalamic areas and the cortical areas they interconnectd with. From Rovo et al., 2012.](media/fig_rovo_ulbert_ascady_12.png)

[[#figure_thal-clasca]] provides a map of the rodent thalamus using the [[@^ClascaRubio-GarridoJabaudon12]] categories, and [[#figure_thal-rovo]] provides a map of the macaque monkey thalamus, labeling the type of driver synapse present in each area ([[@RovoUlbertAcsady12]]). By putting these two maps together, we can identify the general functionality of different thalamic areas, some of which have multiple different neuron types within them and thus likely have multiple different functional roles.

* Primary sensory (first-order, C-type): **LGN** (lateral geniculate nucleus, vision, V1); **MGNv** (medial ventral GN; auditory, A1); **VPL**, **VPM** (ventral posteriolateral, lateral, medial; somatosensory, S1). These are most of the green areas in the driver figure, with the exception of VL and AV, having large, strong glutamatergic drivers from subcortical sources (e.g., the retina for LGN).

* Primary motor (first-order, C-type): **VL** (ventral lateral). This is a primary, first-order type of thalamic area, which instead of being sensory, is primary motor (M1). The strong excitatory subcortical drivers onto these neurons arise from ascending spinal pathways that converge onto the deep cerebellar nuclei (DCN) and then project from there into the VL. These neurons likely provide training signals to primary motor cortex, in the form of efferent copy signals reflecting what the spinal motor system actually did, versus the commands that were sent down to it. Anatomically, VL is adjacent to the somatosensory VPL, VPM nuclei which project topographically to primary somatosensory cortex, just across the central sulcus from M1.

* Secondary perceptual (higher-order, M-type focal): **Pulvinar** is the main nucleus of this type, with topographic projections into all of the higher-order sensory areas. The driver map shows these areas in red and pink, where the red areas have strong glutamatergic drivers, but unlike primary areas, these come from the layer 5b cortical drivers, not subcortical drivers. The pulvinar as such is largely absent in the rodent, but the lateral posterior (**LP**) is thought to be a homologous area based on connectivity ([[@LeowZhouSullivanEtAl22]]), and it has the same M-type focal connectivity (LPMC, LPMR). The yellow areas in the driver map may represent areas of pulvinar that receive both cortical and subcortical inputs, specifically from the superior colliculus, which is known to project into the pulvinar.

* Secondary motor (higher-order, M-type multiareal): **VA**, **VM**, **VAM**, **VAL** (ventral anterior, medial and lateral) are all colored white in the driver figure, because they have no strong drivers. Instead, they receive convergent input from a range of different subcortical sources, along with extensive inhibition from the basal ganglia output nuclei (SNr, GPi). These are all of the red dot areas in the connectivity figure, representing the main population of M-type multiareal neurons. Based on this pattern of connectivity, these are likely to have very broad modulatory effects on the corresponding higher-order motor areas, allowing the BG to modulate overall excitability in these motor cortex areas via disinhibitory gating of the BG output pathways.

* Goal-gating: **MDm** (mediodorsal, medial; multiareal but also focal). Based on the detailed examination of MD connectivity by [[@^KuramotoPanFurutaEtAl17]], and the white driver color of the most medial areas of MD, this subregion of MD does not have any strong drivers, and largely receives inhibitory projections from the BG output nuclei and especially from the ventral pallidum (VP). According to the [[Rubicon]] model and the extensive data reviewed in [[prefrontal cortex]], this MDm area drives disinhibitory gating of distributed goal representations across multiple PFC areas. [[@^KuramotoPanFurutaEtAl17]] report that each MDm neuron interconnects with multiple areas, and most terminate in central layers 3,4 not layer 1. This represents a distinct type of connectivity, that could be labeled C-type (multiareal).

* v/mPFC predictive learning and goal updating: **MDc**, **MDl** (mediodorsal central, lateral, M-type focal). These areas of MD are effectively like the pulvinar for higher-order sensory areas, in terms of receiving strong cortical driver inputs (red), while having M-type focal connectivity patterns. Based on the strong theoretical and empirical basis for pulvinar supporting predictive learning, this same function is likely supported for the goal-driven areas of OFC and PL PFC that these MD areas project to.

<!--- TODO: do these areas receive inhibition from bg? Tom thinks so.. look in Kuramoto -->

* IL (PF, CM, etc). Easy

* AV: hippocampus. LD?


[[@GiguereGoldman-Rakic88]] is primate study showing patchy organization in primates.

