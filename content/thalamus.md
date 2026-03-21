+++
Categories = ["Neuroscience"]
bibfile = "ccnlab.json"
+++

The **thalamus** plays essential roles in multiple different brain circuits, and is thus discussed mostly in the context of those other circuits. This page provides a big-picture organizational and functional view of the thalamus, with pointers to those other pages. Understanding the functions of the thalamus provides a powerful unifying and bridging tool for understanding neural function more broadly, because once you understand how one domain works, the same concepts can be extended to other domains.

From an [[evolution|evolutionary]], comparative perspective, the thalamus consistently provides interconnectivity between primary sensory areas and the _pallium_, which is the outer-most region of the brain, i.e., the _telecephelon_. In mammals, the [[neocortex]] is a major component of the pallium, while other vertibrates have smaller, simpler cortex-like structures in their pallium that play the same functional role as the neocortex ([[@Butler08]]). The functions highlighted below are specific to mammals.

* Relaying with attentional modulation the primary sensory and motor inputs to the neocortex, such as the LGN (lateral geniculate nucleus) for vision. There are extensive top-down projections from sensory cortical areas back to these primary sensory thalamic areas, supporting attentional modulation of these sensory inputs ([[@PoltoratskiMaierNewtonEtAl19]]; [[@Briggs20]]). Primary motor cortex receives both a direct primary motor input from the [[cerebellum]] and inputs from somatosensory thalamus, which play an analogous role on the motor side, teaching motor cortex how to "speak in the language of the spinal cord".

* In posterior [[neocortex]], the _pulvinar_ nucleus of the thalamus supports [[predictive learning]] by alternatively representing a top-down prediction and a bottom-up outcome. Similar connectivity is present in the _mediodorsal (MD)_ nucleus that interconnects with [[prefrontal cortex]] and lower [[motor]] cortex areas, suggesting a similar predictive learning role. Likewise, anterior thalamic areas connected to the systems processing [[space]] can support this form of learning.

* There are multiple different thalamic nuclei including the MD, VA and VM that are interconnected with frontal [[neocortex]], particularly [[prefrontal cortex]] (PFC) and the [[basal ganglia]] (BG), which play critical roles in goal-driven behavior and [[motor]] control as captured in the [[Rubicon]] model.

* The intralaminar (IL) nuclei provide an extensive feedback pathway for the basal ganglia.

* The paraventricular thalamic nuclei interconnect with the goal-driven PFC, hippocampus and associated [[emotion]] processing subcortical areas, and provide overall modulation of goal-driving processing.

## Thalamic circuitry

{id="figure_pulv-conns" style="height:15em"}
![Connectivity between the neocortex and the pulvinar nucleus of the thalamus, in the case of primary and secondary visual areas. Visual input from the retina provides driver inputs to the lateral geniculate nucleus (LGN) of the thalamus, which then sends glutamatergic excitatory projections primarily into layer 4 of primary visual cortex (V1). The pulvinar thalamic nucleus receives strong driver inputs from deep layer 5b neurons in V1, which can have as few as one driver synapse per thalamic relay neuron (TRN). There are numerous and relatively weaker projections from layer 6 (VI) neurons in secondary visual cortex (V2). Diagram based on Sherman & Guillery (2006).](media/fig_pulvinar_connectivity.png)

[[#figure_pulv-conns]], based on [[@ShermanGuillery06]], shows the thalamic connectivity for primary and secondary visual cortex in the mammalian brain. The main features of this connectivity are common across most thalamic areas.

There is only one predominant type of thalamic neuron in the different thalamic nuclei, the _thalamic relay neuron_ (TRN). The _thalamic reticular nucleus_ (TRN), which is an outer shell wrapping around the body of the thalamus, is the one exception, having GABAergic [[inhibition|inhibitory]] neurons, which we discuss below. TRNs release excitatory glutamate neurotransmitter into the [[neocortex]], and receive two qualitatively distinct types of inputs:

* **Drivers** are strong inputs, often relatively few in number (as few as a single input). They can have large extended _glomeruli_ with multiple synaptic connections and even a "dedicated" inhibitory intgerneuron encapsulated in a single anatomical structure. These can originate from subcortical areas (e.g., retinal input to the LGN for vision) or cortical areas (e.g., the deep layer 5b IB intrinsic bursting neurons that are drivers in the [[predictive learning]] performed by the pulvinar). Some thalamic areas have no identified driver inputs, and are typically those that receive extensive inhibitory inputs from the basal ganglia output pathways.

    The nature of the driver inputs is one of the most important clues to thalamic (and cortical) function. For example, cortical drivers from intrinsic bursting neurons provide the temporal dynamics needed to support predictive learning via the [[temporal derivative]] mechanism, whereas subcortical sensory drivers do not appear to have this property. However, subcortical drivers in the motor system naturally have the requisite delays, as discussed in the [[motor#Frontal cortex learning via subcortical efferent copy signals|motor]] page.

* **Non-drivers** are the rest of the inputs, which are like most other synaptic connections in the brain. [[@ShermanGuillery06]] and many others label these as _modulatory_ inputs, but that is perhaps too strong of a functional assumption, given their proposed function in the context of [[predictive learning]]. For example, there is ample evidence that these inputs can drive sustained firing of TRNs, as reviewed in [[@OReillyRussinZolfagharEtAl21]].

The TRNs notably lack any local excitatory interconnectivity, and only have a form of pooled inhibition mediated by the TRN reticular neurons. Therefore, they are uniquely suited for providing a faithful representation of the inputs they receive, serving as a kind of _projection screen_ that does not impose significant distortion of its own.

Although the relatively inert nature of the thalamic circuitry has led to its characterization as a _relay_ system, the presence of two distinctive types of inputs suggests a more powerful kind of function could be be supported. The [[predictive learning]] framework for example shows how the distinctive temporal aspects of the driver inputs to the pulvinar supports a [[temporal derivative]] computation that effectively computes an error gradient between the driver and non-driver inputs, across time.

The nature of the connectivity from TRNs into the cortex provides another important insight into the function of the thalamus. There are two primary dimensions along which these projections vary, which were originally captured by the **core** vs. **matrix** categories of [[@^Jones98]], but which are more accurately categorized as dimensions ([[@ClascaRubio-GarridoJabaudon12]]; [[@ShermanUsrey24]]):

* Laminar targets: the classical _core_ type projects mainly to the middle layers of the cortex, primarily layer 4 (e.g., as shown in [[#figure_pulv-conns]]), while the classical _matrix_ type projects to layers 1 and 5, but avoids the middle layers. However, there are actually different proportions of these connections across different thalamic neurons, rather than sharply defined categories.

    Functionally, projections into the middle layers provides direct input information into the cortex, as in the prototypical sensory input into primary sensory areas. Projections into layers 1 and 5 can drive two main functions: regulation of overall cortical activity levels via the VIP inhibitory interneurons in layer 1 (this was the original interpretation of the matrix type), and maintenance of sustained activity (primarily in [[prefrontal cortex]]) via bidirectional excitatory loops through layer 5b PT (pyramidal tract) neurons, which also have extensive dendritic tufts in layer 1.

* Area focus: the classical _core_ type projects to only one cortical area, typically in a relatively circumscribed, _focal_ manner, while the classical _matrix_ type projects in a broader, diffuse to multiple different areas, consistent with a more modulatory function.

{id="figure_map-phillips" style="height:30em"}
![Map of thalamic areas in the rodent according to gene expression patterns, which identifies 3+1 major clusters. Blue are primary sensory and other areas that receive strong subcortical driver inputs. Red are higher-order areas that either receive cortical drivers or no drivers. Orange are the intralaminar and paraventricular nuclei that have more specialized connectivity with the basal ganglia and other subcortical areas. The reunions nucleus, which interconnects the prefrontal cortex with the hippocampus, has its own unique status. From Phillips et al., 2019.](media/fig_thalamus_types_phillips_etal_19.png)

[[#figure_map-phillips]] shows a map of different thalamic areas in rodents, which provides a broad set of 3+1 categories based on patterns of gene expression, which is useful for informing our understanding of function across different thalamic areas. This data does not provide more fine-grained distinctions in terms of patterns of connectivity, but they do generally correspond to different types of driver inputs, as we'll see below.

{id="figure_types-calasca" style="height:30em"}
![Different types of thalamic projections into the cortex, varying in the laminar targets and the area focus, according to Clasca et al., 2012. The C-type is the classical first-order sensory core-type, projecting to the middle lamina in a focal manner, while the M-type (multiareal) is the classical matrix type projecting to layer 1 and 5, across multiple areas. The M-type (focal) combines features from classical core and matrix types, with projections that target middle layers in addition to layer 1, with more focal area-wise connectivity. Meanwhile the IL (intralaminar) pattern (i.e., PF = parafascicular thalamus) is specific to neurons in these nuclei, targeting only deep layers and also extensively the striatum.](media/fig_thalamus_types_clasca_etal_12.png)

[[#figure_types-calasca]] shows an extension of the classical core and matrix categories by [[@^ClascaRubio-GarridoJabaudon12]], with two distinct M-types, one of which corresponds to the classical matrix type, while the M-type (focal) combines features of core and matrix. There is also the unique projection pattern of the intralaminar (IL) nuclei, which includes the PF (parafasicular) and CM (centromedian) nuclei, which play a critical role in the function of the basal ganglia (see [[basal ganglia#Parafascicular feedback loops]]), and have extensive connectivity into the striatum along with projections into the deep cortical layers.

Instead of retaining a core-vs-matrix categorical approach, we prefer to use functional, feature-based terminology in describing the connectivity profiles of neurons in different thalamic areas. A given TRN cell will in general have multiple of these properties.

+ **Input:** focal projections into the middle layers, driving informational input to a cortical areas.

* **Maintenance:** focal projections to layers 1 and 5, supporting active maintenance in layer 5b PT neurons.

+ **Modulatory:** diffuse projections to layers 1 and 5. This is not exclusive to classical arousal-modulating matrix-type neurons, as basal-ganglia modulated motor areas including VA and VM have this pattern, to provide more dynamic, fine-grained modulation of motor cortical areas.

* **Single:** projects to a single area.

* **Multi:** projects to multiple areas. Note that this can be orthogonal to the focal vs. diffuse nature of the projection, for example MDm (mediodorsal medial) has focal but multi-area projections.

## Functional organization of the thalamus

{id="figure_map-clasca" style="height:30em"}
![Map of thalamic areas in the rodent according to the four types shown in the prior figure, using the same colors. From Clasca et al., 2012.](media/fig_thalamus_areas_clasca_etal_12.png)

{id="figure_map-rovo" style="height:50em"}
![Different types of thalamic drivers in the primate thalamus, which provides key insight into functions of these thalamic areas and the cortical areas they interconnectd with. Large = strong, small = weak. vGLUT2 = subcortical drivers, vGLUT1 = cortical drivers; From Rovo et al., 2012.](media/fig_rovo_ulbert_ascady_12.png)

[[#figure_map-clasca]] provides a map of the rodent thalamus using the [[@^ClascaRubio-GarridoJabaudon12]] categories, and [[#figure_map-rovo]] provides a map of the macaque monkey thalamus, labeling the type of driver synapse present in each area ([[@RovoUlbertAcsady12]]). These are functionally interpretable as follows:

* **Subcortical** drivers: Strong = large vGLUT2 (green).
* **Cortical** drivers: Strong = large vGLUT1 (red), Weaker = small vGLUT1 (pink).
* **None** (white). These typically receive inhibitory inputs from the basal ganglia.
* Other combinations are discussed below where applicable.

By putting these two maps together (and keeping in mind the data in [[#figure_map-phillips]]), we can identify the general functionality of different thalamic areas.

It is critical to recognize that although each thalamic nucleus has a predominant type of TRN connectivity pattern that we identify it with, there is almost always some diversity, such that for example classical matrix-type neurons are intermixed within the classical core-type that predominate in the LGN ([[@ShermanUsrey24]]). Functionally, these matrix-type neurons can cover a lot of territory with their broad connectivity patterns, so fewer are required to do the job. In general, we assume that modulation of overall excitability and arousal levels in the cortex via such neurons is an additional job performed by most thalamic areas, but dedicated areas that perform only this function are unnecessary.

* **Primary sensory** (Driver: Subcortical; Projection: Input, Single): **LGN** (lateral geniculate nucleus, vision, V1); **MGNv** (medial ventral GN; auditory, A1); **MGNm** (medial medial GN; vestibular and others); **VPL**, **VPM** (ventral posteriolateral, lateral, medial; somatosensory, S1). These are most of the green areas in the driver figure, with the exception of VL and AV, having large, strong glutamatergic drivers from subcortical sources (e.g., the retina for LGN).

* **Primary motor** (Driver: Subcortical; Projection: Input, Single): **VL** (ventral lateral). This is a primary, first-order type of thalamic area, which instead of being sensory, is primary motor (M1). The strong excitatory subcortical drivers onto these neurons arise from ascending spinal pathways that converge onto the deep [[cerebellum|cerebellar]] nuclei (DCN) and then project from there into the VL. These neurons likely provide training signals to primary motor cortex, in the form of efferent copy signals reflecting what the spinal motor system actually did, versus the commands that were sent down to it. Anatomically, VL is adjacent to the somatosensory VPL, VPM nuclei which project topographically to primary somatosensory cortex, just across the central sulcus from M1. See below for further discussion of the [[#motor thalamus]] areas.

* **Primary space and time** (Driver: Subcortical; Projection: Input, Multiple): **AV** (anteroventral) and **AD** (anterodorsal). Like the VL, this is an interesting case of a subcortically-driven, input-projecting area that is not a standard sensory input. These areas receive driver inputs from the mammillary bodies (medial to AV and lateral to AD), which convey important [[space|spatial]] and [[time|timing]] signals, that are then projected into the (pre-post) subiculum and retrosplenial cortex, which are the primary spatial representation areas of the cortex. The lateral mammillary bodies projecting to the AD nucleus are the primary source of _head direction_ cells, that integrate vestibular and motor signals to keep track of which way the head (and body more generally) is facing in an _allocentric_ (world-based) coordinate system. The medial mammillary bodies projecting to the AV convey the [[theta rhythm]] signals generated in the _ventral tegmental nucleus (VTN)_, which also drive neurons in the medial septum, to coordinate processing throughout the medial temporal lobe and [[hippocampus]].

<!---  TODO: more about nature of projections etc. -->

* **Secondary perceptual** (Driver: Cortical; Projection: Input, Modulatory, Single & Multi): **Pulvinar** is the main nucleus of this type, with topographic projections into all of the higher-order sensory areas. The driver map shows these areas in red and pink, where the red areas have strong glutamatergic drivers, but unlike primary areas, these come from the layer 5b cortical drivers, not subcortical drivers. The pulvinar as such is not labeled as such in the rodent, but the lateral posterior (**LP**) is thought to be a homologous area based on connectivity ([[@LeowZhouSullivanEtAl22]]), and it has the same focal connectivity (LPMC, LPMR). The yellow areas in the driver map represents areas of pulvinar that receive both cortical and subcortical inputs, specifically from the superior colliculus, which is known to project into the pulvinar.

* **Secondary motor** (Driver: None; Projection: Modulatory, Maintenance, Input, Multi & Single): **VA**, **VM**, **VAM**, **VAL** (ventral anterior, medial and lateral) are all colored white in the driver figure, because they have no strong drivers. Instead, they receive convergent input from a range of different subcortical sources, along with extensive inhibition from the basal ganglia output nuclei (SNr, GPi). These are all of the red dot areas in the connectivity figure, representing the main population of M-type multiareal neurons. Based on this pattern of connectivity, these are likely to have very broad modulatory effects on the corresponding higher-order motor areas, allowing the BG to modulate overall excitability in these motor cortex areas via disinhibitory gating of the BG output pathways. See [[#motor thalamus]] below for further discussion.

<!---  TanibuchiKitanoJinnai09 -- VAmc = DLPFC+SNr, MD lat = SNr-gated IFG PFC  * Secondary spatial (? ?): **AM** (anteromedial). ?? -->

* **Secondary PFC, goal-gating** (Driver: None; Projection: Input, Maintenance, Modulatory, Multi): **MDm** (mediodorsal, medial). Based on the detailed examination of MD connectivity by [[@^KuramotoPanFurutaEtAl17]], and the white driver color of the most medial areas of MD, this subregion of MD does not have any strong drivers, and largely receives inhibitory projections from the BG output nuclei and especially from the ventral pallidum (VP). According to the [[Rubicon]] model and the extensive data reviewed in [[prefrontal cortex]], this MDm area drives disinhibitory gating of distributed goal representations across multiple PFC areas. [[@^KuramotoPanFurutaEtAl17]] report that each MDm neuron interconnects with multiple areas, and most terminate in central layers 3,4 not layer 1. This represents a distinct type of connectivity, that could be labeled C-type (multiareal).

<!--- TODO: CollinsAnastasiadesMarlinEtAl18  and PhillipsKambiRedinbaughEtAl21 PhillipsAfrasiabiKambiEtAl25 LiZhengWangEtAl24 -->

* **Secondary PFC, predictive learning and goal updating** (Driver: Cortical; Projection: Input, Maintenance, Single): **MDc**, **MDl** (mediodorsal central, lateral). These areas of MD are effectively like the pulvinar for higher-order sensory areas, in terms of receiving strong cortical driver inputs (red), while having M-type focal connectivity patterns. Based on the strong theoretical and empirical basis for pulvinar supporting predictive learning, this same function is likely supported for the goal-driven areas of OFC and PL PFC that these MD areas project to.

<!--- TODO: do these areas receive inhibition from bg? Tom thinks so.. look in Kuramoto -->

* **Intralaminar** (**IL**: **PF**, **CM**, etc) (Driver: None; Projection: IL-type to striatum, cortical deep layers). These nuclei provide extensive feedback to the [[basal ganglia]] and also to deep layers of the motor neocortex, and play critical specialized roles in BG function and learning.

<!--- * LD: ??  spatial right? -->

[[@GiguereGoldman-Rakic88]] is primate study showing patchy organization in primates.


## Frontal thalamus

There are three major functions for thalamic areas interconnected with frontal areas:

* Training of primary motor cortex to speak the language of the spinal cord (VL, POm / PULm). These receive strong driver input from spinal motor signals (VL) or somatosensory signals (POm / PULm), and project strong driver inputs into the middle lamina of motor areas.

* Corticocortical [[predictive learning]], analogous to that supported by the pulvinar for posterior cortex, so higher layers can learn to predict the motor-related activity of lower layers. This is supported by subsets of neurons within the MD and AM nuclei. This helps learning of longer sequences of motor actions, for example.

* [[Basal ganglia]] gating signals across multiple timescales, to trigger dynamic updating of frontal representations, for goal engagement, motor plan, and motor action updating. This is supported by other subsets of neurons within MD and AM, and the VA (which corresponds with VM in rodent; [[@IlinskyKultas-Ilinsky87]]) nucleus, including the special VAmc nucleus that is specific to primates.

### Rodent

{id="figure_prjn-kuramoto" style="height:25em"}
![Summary of projections from the different ventral motor thalamic areas in rodents, where VM is the rodent version of VA, and VAmc does not exist. VL can be defined as the specific subset of neurons receiving cerebellar nucleus (CN) driver inputs (inluding caudodorsal VA), which have focal core-like projections as shown on the right. The VM neurons (which include rostroventral VA), consistent with VA and VAmc in macaque, receive extensive BG output (SNr, GPi) and project in a more diffuse matrix-like pattern to superficial layers. From Kuramoto et al., 2015.](media/fig_thalamus_prjn_vm_vl_kuramoto_etal_15.png)

We begin with the simpler case of the rodent, which has each of these different functions clearly represented by different thalamic nuclei. The primate case then adds more complexity that can be understood as building upon this same basic foundation.

[[#figure_prjn-kuramoto]] (from [[@KuramotoOhnoFurutaEtAl15]]) shows the simpler situation in rodents, where one can define two ventral thalamic areas according to their driver inputs, that then correspond roughly to corresponding areas in the macaque monkey: VL, and VM. VL (which also includes caudodorsal VA) receives strong drivers from the cerebellar nuclei, and projects focally to middle lamina in motor areas (i.e., core-like), enabling motor areas to learn the language of the spinal cord. VM (which also includes rostroventral VA) receives BG disinhibitory outputs (SNr, GPi), and projects to superficial lamina in a widely-branching manner (i.e., matrix-like), and allows BG to send a broad gating signal to all of motor cortex. 

{id="figure_md-pfc" style="height:20em"}
![Thalamocortical connectivity for different regions of the MD thalamus in rats. MDm (medial) is the main goal-driven gating area, which receives disinhibitory input from ventral pallidum (VP), and projects to the core goal areas (IL, PL, and OFC areas). MDc (central) projects to more OFC areas and may be important for updating during goal approach. MDl (lateral) targets the complementary ACC and PL goal areas, along with motor plan areas in (FrA = frontal association higher level motor, M2 = secondary motor). From Kuramoto et al., 2017.](media/fig_md_pfc_conns_kuramoto_etal_17.png)

[[#figure_md-pfc]] shows the different regions of the MD thalamus interconnecting with areas of the PFC in rats (from [[@KuramotoPanFurutaEtAl17]]), with the primary goal-gating functionality supported by MDm (medial), which projects to all of the goal areas (IL, PL, OFC, and FrA) and receives disinhibition from VP (ventral pallidum), which is the output pathway from ventral striatum goal gating areas as described in [[basal ganglia]]. The other two areas of MD (c = central, l = lateral) provide complementary, more specific gating pathways for OFC (central) and motor, ACC, and PL (lateral), potentially allowing each of these coordinated components of the goal representation to be updated separately.

{id="figure_frontal-areas" style="height:40em"}
![Thalamic nuclei in macaque monkeys showing the projections received from SNr (BG) as a function of the frontal areas that these cortical areas project to (inset below legend), with A = PSv, ventral of the principal sulcus (e.g., IFG in human PFC); B = dlPFC (areas 9, 46); C = motor areas (M1, SMA, PMd). VAmc projects widely to all PFC areas, and has significant SNr input, while MDpc and MDmf mainly interconect with the PSv area associated with caudal SNr inputs. AM = anteromedial. Note that MD has a mix of neurons with no SNr input (typically with cortical drivers, like pulvinar), and those with SNr or VP (ventral pallidum, for MDmc) inputs. From Tanibuchi et al., 2009.](media/fig_thalamus_frontal_tanibuchi_etal_09.png)

### Primate

{id="figure_phillips-cons" style="height:40em"}
![Similarity of patterns of thalamic connectivity with prefrontal cortex (PFC) in the macaque, which typically also reflects direct interconnectivity strength. Red and orange areas are ventral and medial PFC goal-driven areas (i.e, the core of the [[Rubicon]] model), including OFC (13, 14, 11; specific outcomes), IL (25; abstract outcomes), and PL (32; integrated reward-cost utility: note broader interconnectivity, especially with 24, 25 and 9); Yellow (10) is polar frontal cortex, unique to primates, important for episodic memory and outer loops of processing; Green (47,12) is vlPFC (ventrolateral) for control over the temporal lobe (language, objects etc); Cyan (9, 46d, 45a) is dlPFC (dorsolateral) for control over parietal lobe for spatial and temporal organization of action (broad plans); Blue (8) is FEF (frontal eye fields) which has similar connectivity to PM (premotor), SMA (supplementary motor area) and primary motor cortex (M1). Area 24 is ACC (anterior cingulate cortex) which represents the value (mostly costs) of actions, and has very broad interconnectivity with motor areas, but also goal areas, especially 32 (PL). Adapted from Phillips et al., 2019.](media/fig_thalamus_pfc_simat_phillips_etal_19.png)

{id="table_motor-thal" title="Motor thalamus areas"}
| Area      | Function     | Drivers          | Modulators        | BG   | Strong outs    | Weak outs | Lamina   | Other ins            | Other outs |
|-----------|--------------|------------------|-------------------|------|----------------|-----------|----------|----------------------|------------|
| **Learning** |           |                  |                   |      |                |           |          |                      |            |
| PULm / POm | S1->M1 Learn | S1, Spinal      | M1, M2            | --   | M1, M2, SMA..  |           | Mid?     | ?                    | ?          |
| VL        | CN->M1,M2 Learn | CN            | M1,M2             | --   | M1, M2         |           | Mid      |                      |            | 
| MD        | Predictive learning | (see below) |                 | --   |                |           |          |                      |            |
| **Gating** |  **(focal)** |                 |                   |      |                |           |          |                      |            |
| AM        | MTL <-> goal | ACC, OFC, IL, PL | d/vlPFC, lOFC     | mGPi | 10, 11, 12, 46 | PL, ACC   | Mid      | CA3, aSub, BLA, mGPi | hippo?     |
| MDm/mc    | goal engage  | OFC, IL, PL      | vlPFC, ACC        | VP   | OFC            | PL, IL    | S/D, Mid | IT polar, EC, PRh, Sub, BLA |     |
| MDc/pc    | dl/plan WM   | 10, dlPFC        | ACC, PL, PM, SMA  | SNr  | dlPFC (WM)     | PM, SMA   | Mid      | parietal             | parietal   |
| MDl/mf/dc | eye/act motor| 8 (FEF), PM, SMA |                   | SNr  | 8, PM, SMA     |           | Mid      | SC, parietal attn    |            |
| **Gating** | **(diffuse)** |                |                   |      |                |           |          |                      |            |
| VA        | BG->motor    |                  |                   | GPi  | M1, M2, SMA, DM|           |          |                      |            |
| VAmc / VM | BG->dl/motor | dl, PL, ACC      |                   | SNr  | dlPFC, M2, M1  |           | Mid + S/D| parietal, temp pole  | parietal   |

[[#table_motor-thal]] describes motor thalamus areas in primate using brodmann area numbers, from [[@PhillipsKambiRedinbaughEtAl21]] and [[@TimbieBarbas15]]

* MD has both BG-gated cells for BG-gated updating, and also intermixed neurons that do not receive BG inputs (SNr or GPi) and can play a predictive learning function. MD mostly targets middle layers ("core style"), although with largely non-standard calretinin marker neurons, instead of standard PV core-style (CB = matrix style, only found in MDm). Cortical driver source inputs are PT layer 5b (deep) ([[@XiaoZikopoulosBarbas09]]). For the BG gating, the following types appear to be separately supported. 
    + MDm: Goal representations in goal areas (PL, ACC, IL, OFC).
    + MDc: Working memory / plan representations in dlPFC.
    + MDl: Motor plans at finer time scales.
* AM also has BG gating of goal areas, but appears to be more interconnected with MTL / hippocampal systems. TODO: find paper just on AM. Driving 10 and 46, and getting direct CA3, aSub inputs. Cortical driver source inputs are PT layer 5b (deep) ([[@XiaoZikopoulosBarbas09]]).

    + MDm (MDmc) is the principal goal engage system, triggered by VP gating signals, updating OFC, PL, IL, also with other hippocampal and BLA inputs (todo: explore those in [[Rubicon]] model.)
    + MDc (MDpc) is heavily SNr innervated (refs) and drives working memory updating in dlPFC and frontopolar (10). This is a core-style projection. It absolutely has all the properties needed for BG-mediated WM updating ala PBWM, right? todo: read primary refs on this. [[@KuramotoIwaiYamanakaEtAl17]] on MD only shows rats, and thus misses the critical PFC nature of this in primates.
    + MDl (MDmf/MDdc) has strong FEF 8a driver inputs, and in primates is likely drive broader motor gating (updating motor plans based on new eye position?) triggered by new saccade. In rodents it is more strongly PL interconnected.

TODO: good summary on VL, VA, POm -- not in [[@PhillipsKambiRedinbaughEtAl21]]
    
* VA: "Its thalamocortical neurons are under the control of the inhibitory projection from the GPi (Ilinsky and Kultas-Ilinsky, 1987; Ilinsky et al., 1993; Jones, 2007)."  

* VAmc has highly overlapping inputs and very broad outputs. is main pathway for SNr/BG gating in primates; rodents have that through MD; unclear why that is not true also in primates given above connectivity. Matrix-like layer I outputs, inputs from goal and dlPFC areas -- update motor cortex when new plans are activated basically. Thalamic prjns arise from 5a (superficial, IT) not PT.
    + 6 CT targets PV -> Mid layers (core) = classic CT predictive learning; 5a targets CB -> S/D (matrix) = BG gating?? So also a mix of different types.
    + VAmc vs. MDc: very broad inputs, but targets lamina 3,4!? lamina looks like core, but projection *inputs* are very broad. outputs??
    

* AM, MD get significant CT layer 6 inputs (only 20% layer 5 drivers), whereas VL,VA,VAmc all have nearly 50% layer 5 inputs ([[@XiaoZikopoulosBarbas09]]).  PT layer 5b (deep) ([[@XiaoZikopoulosBarbas09]]) are more focal, whereas 5b (superficial, IT) are broader
    
## Anterior thalamic areas


