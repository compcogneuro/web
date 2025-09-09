+++
Categories = ["Neuroscience"]
bibfile = "ccnlab.json"
+++

The **theta rhythm** is a roughly 5 hz oscillation found in various areas of the brain, with a characteristic period of 200 ms (1000 ms / 5 hz = 200 ms). There are many diverse sources of neuroscience evidence suggesting that there are neural mechanisms specifically organized around processing at this timescale, across various mammalian species, including two prominent examples:

* In primates including humans, the median visual fixation duration is strongly peaked around 200 ms ([[@DevillezGuyaderCurranEtAl20]]): we tend to look at things for about 200 ms at a time (cites).

* The theta cycle in the rodent [[hippocampus]] has been extensively studied, and it drives a complex series of dynamics throughout the hippocampal formation (cites).

This timescale plays a critical role in the [[Axon]] framework, for organizing the process of [[predictive learning]] and the [[kinase algorithm]] learning mechanisms. While the biological nature of the theta cycle exhibits considerable flexibility and can be continuously synchronized with salient events taking place in the environment, we typically take advantage of the much greater simplicity of simply imposing this 200 ms time window onto the simulated environments that we use for training and testing the models.

<!--- TODO: more? -->

## Brainstem origins: VTN, MMB

The _ventral tegmental nucleus (VTN)_ (of Gudden) contains neurons with intracellular currents that generate  theta rhythm oscillations intrinsically, likely providing the source of this rhythm for the medial temporal lobe pathway converging on the hippocampus. These VTN neurons project to the medial mammillary bodies (MMB) which provide strong driving inputs to the anteroventral (AV) nucleus of the [[thalamus]], which then project into the subiculum of the hippocampus.

<!--- TODO: figure, more details about MMB contributions, top-down vs. bottom-up reset etc. -->

