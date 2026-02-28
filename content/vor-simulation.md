+++
Name = "VOR simulation"
Categories = ["Neuroscience", "Simulations"]
bibfile = "ccnlab.json"
+++

<!--- <sim-vor> -->

<div>

## Introduction

This simulation provides an interactive model of the [[cerebellum#vestibulo-ocular reflex]] (VOR) and the contribution of the [[cerebellum]] **forward model** learning. See these links for essential background information before proceeding with this model.

### Existing computational models

#### Lisberger 1994

One of the earliest models, features the **HGVP** (horizontal gaze-velocity Purkinje) cells, which project inhibition onto the **FTN** (flocculus target neurons, in the vestibular nucleus), that then project to the ocular motor system, and are the functional equivalent of cerebellar nucleus neurons for the vestibular system. These neurons have a high tonic level of activity, are excitatory ([[@MatsunoKudohWatakabeEtAl16]]) and have two different subtypes respond in opposite directions: E_i = ipsilateral, E_c = contralateral eye movement-related activity ().

Under normal conditions, the ipsilateral neurons increase firing in response to head movements in the ipsilateral direction, while the contralateral ones exhibit the opposite pattern, providing a classic opponent-organized system. With VOR training, increases in VOR gain were associated mostly with large decreases in velocity sensitivity for the E_c opponent, and smaller increases in sensitivity for the E_i ipsilateral direction (i.e., reducing the inhibition a lot and increasing the excitation a bit). By contrast, reductions in VOR gain were associated with increases in E_c opponent activity _and_ E_i activity.

* E_i -> Ipsi = head rotating ipsi, *eye* rotating contra!!
* E_c -> contra = head rotating contra, *eye* rotating ipsi

GAIN = balance between these two! not either one alone. More gain -> E_i+, E_c-; Less gain -> E_i-, E_c+

Overall most of the HGVP Purkinje cells respond too late to drive the earliest modified component of the VOR, which must be driven by changes in the brain stem vestibular inputs to FTNs: the HGVP may contribute to plasticity but the FTN neurons are clearly the locus of most of the critical learning.

