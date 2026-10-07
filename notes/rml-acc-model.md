# Reinforcement Meta Learner (RML) model

These are raw notes (mostly extracts, minor commentary) on a series of papers by Massimo Silvetti, Eliana Vassena and colleagues on an ACC-centric model of meta-learning and RL.

## SilvettiVassenaAbrahamseEtAl18

Summary: not unreasonable model of dACC -- minor quibble about precise division of labor w/ motor / PFC etc but not that big of a deal.

From discussion:

> We propose that the dACC provides RL signals to the LC, about the statistical structure of the environment; in turn, the LC processes those signals to select optimal learning rate by approximating a Bayesian learner.

Note: is this too powerful a job for the LC itself to handle? what is the computation here?

> Finally, effort control is itself modulated by the same mechanisms optimizing learning rate for action selection. This aspect provides near optimal meta-flexibility to cognitive control, a novelty that merges cognitive control with Bayesian learning.

Note: it does not seem that there is an explicit computation of cost / benefit utility, but rather the "catecholamines" (DA?) are ramped up by reward but also have cost limits?

> Although we described them separately, in the RML, learning rate, effort estimation and reward-related processes are integrated and mutually dependent. For example, dynamic control of learning rate (λ) is based on RL signals from dACC modules. Learning rate modulation influences both decision-making for action selection and for boosting control (b). Boosting control modulates in parallel both LC and VTA, modulating both performance (NE) and learning (DA). Catecholamine modulation changes behavioural performance, influencing action selection and environmental feedback, thus influencing LC control over learning rate.

### Equations

p(a|s) = softmax(v(s,a) - C(s,a) / NE, tau)

v = reward value, C = cost -- so NE directly discounts costs.

separate learning of a "bosting value":

\delta v(s,b) = \lambda (DA_B,t - v(s,b))

basically just incremental learning to the DA_B,t value which is reward discounted by boost cost.

p(b|s) = softmax(v_b(s,b), tau)

b = boost value (integer between 1-10) -- weird.

NE = b

DA = r_t(R + \mu b) + b(1-\mu) \rho max_a(v(s', a))

so the boost b is distributed between boosting primary rewards, and modulating the TD Q-value which is what the 2nd term is (max 1-step future value for best action).

r_t is binary for when primary reward is present or not. so absence of reward is not penalized!

learning rates are adapted by a running-average variance estimator (in the LC), 

so, overall, fairly simple model.

## other quotes

> Like in earlier RL models, the dACC in the RML computes the values of specific stimuli and actions to achieve adaptive  behavior. However–and unlike in earlier models–dACC internal dynamics is modulated by catecholamines via recurrent interaction between the dACC itself and the brainstem nuclei.

> This double loop structure is aimed at optimizing performance (i.e., maximizing reward) while minimizing two different types of costs: the costs of motor actions (external loop; e.g. the metabolic cost of climbing a stair), and the boosting costs of neuromodulators release (internal loop; e.g. the cost of neurotransmitters depletion)

> In the RML, the dACC plays the role of a performance monitoring system, which compares expectations about environmental states and executed actions with environmental outcomes (cf. [22]). Discrepancies between expectations and outcomes generate prediction error (PE) signals (Figure G in S2 File), which are used to update the expectations themselves [3].

> One dACC module (dACCAct in Fig 1) receives environmental states and selects actions directed toward the external environment (part of the external loop).

> A second dACC module (dACCBoost in Fig 1) receives environmental states and consequently modulates (that is, boosts) the release of catecholamines from the brainstem nuclei LC and VTA (part of the internalloop). Catecholamines, in turn, control the internal dynamics of the dACC in real time (i.e. while the RML is interacting with the environment), by modulating the magnitude of reward signals (by VTA module) and the amount of effort(by LC module) that the RML exerts to execute a task.

> Hence, the interaction between dACC and LC allows disentangling uncertainty due to noise from uncertainty due to actual environmental changes [30,31] promoting plasticity (high learning rate) when new information must be acquired (condition Vol), and stability (low learning rate) when acquired information must be protected from noise (conditions Stat and Stat2). 

>  LC performs approximate Bayesian analysis on those signals to compute optimal learning rate. For this reason, the dACC is more responsive to overall environmental uncertainty (expressed by average PE), while LC selectively responds to volatility

> Increased NE influences decision-making in the dACCAct (effect of NE on action cost estimation in decision-making process, Eq 2 in Methods), facilitating effortful actions, while increased DA affects learning in the dACCAct (Eqs 1 and 7A in Methods), increasing the reward signal related to effortful actions. At the same time, boosting catecholamines has a cost (Eq 6B in Methods), so that the higher b, the higher was the reward discount


## SilvettiLasaponaraDaddaouaEtAl23

**A Reinforcement Meta-Learning framework of executive function and information demand**

Same RML as above, but adds a "suprisal" factor to the DA equation, as a function of the expected unsigned prediction error for outcome..

Basically, attention boosts come from RML model -- things that boost RPE get more boosts and they get attention. the modified RML with a surprisal factor fit the data even better.

## SilvestriniMusslickBerryEtAl23

**An Integrative Effort: Bridging Motivational Intensity Theory and Recent Neurocomputational and Neuronal Models of Effort and Control Allocation**

Overall, basically shows that these different frameworks (MIT, EVC, RML, NEAC) are largely similar in their predictions, which all stem from very basic calculus of benefit - cost with increased effort expended as long as this is positive, to overcome the costs, but then dropping quickly once costs exceed effort.

>  motivational intensity theory (MIT; Brehm, 1975; Brehm et al., 1983; Brehm & Self, 1989). From different angles, these theories seek to explain the mechanisms underlying the motivation to engage in effortful behavior.

> the core idea of MIT is that individuals avoid wasting resources and therefore calibrate their effort considering the dif culty and the importance of the task at hand.

> Mounting empirical evidence suggests that the exertion of cognitive control is associated with an intrinsic cost (Botvinick & Braver, 2015; Kool et al., 2010, 2017; Manohar et al., 2015; Westbrook & Braver, 2015). In line with MIT, the models presented in this review all seek to explain how people allocate effort, by describing the decision-making processes underlying the allocation of cognitive control (Shenhav et al., 2013) and by linking these to neural systems that compute such trade-offs (Silvetti et al., 2018) and implement control-demanding behavior (Sarter et al., 2006; Silvetti et al., 2018).

Key issue: what is the cost? opportunity? interference? or just time on task? and comparative: you gotta do something, so pick the best option, so always comparing -- cost is always a factor. but it is really time & effort, right?  And uncertainty. and probability of error. These are all the ACC cost factors.

but not metabolic per-se!

> Together, these ndings supported the prediction of MIT that potential motivation directly determines effort only when task dif culty is unspecified.

> EVC: The cost may take different functional forms but is assumed to scale monotonically with the intensity of the control signal.2

> Previous instantiations of EVC theory quantify perceived task dif culty as the expected probability of performing accurately on a task for a xed amount of control (see Figure 2A; Musslick et al., 2015; Musslick, Cohen, & Shenhav, 2018; Musslick et al., 2019). 

> For instance, the maximum tolerable dif culty of a task is predicted to be a function of opportunity costs: Simulated EVC agents become more likely to disengage from the primary task if the subjective value of the next best alternative task increases (Figure 2D).

> Thus, the RML implements optimal resource allocation along two key dimensions: selecting which are the best actions to be performed (toward the environment) and selecting whether it is worth boosting noradrenaline and dopamine release to achieve the goal (modulating its internal parameters, i.e., meta-learning).

> When confronted with these tasks, RML learns to choose between available options based on the estimation of expected reward and on the cost tied to the required effort, in terms of catecholaminergic boost necessary to successfully complete the task. This is where the meta-learning aspect is critical: the internal boosting signal is necessary to engage in hard tasks, but is also costly, and this cost is factored in into the net-value computation. This net value measured in RML-dACC modules ts with previously observed net-value signal coding in dACC (Chong et al., 2017). For example, the net-value signal in the RML-dACC modules is lower for a hard task compared to an easy task when reward is high (Silvetti et al., 2018).

> In practice, this means that when a task is easy, more boosting is not necessary. When difficulty increases, more boosting is required to successfully complete the task.

> As visible in the simulation depicted in Figure 2B, the EVC theory predicts that the amount of control invested increases with expected task dif culty until a maximum tolerable task difficulty is reached and drops sharply if the task becomes more dif cult. The maximum worthwhile control intensity (i.e., success importance) determines the maximum tolerable dif culty in a similar fashion as MIT. Thus, MIT and EVC theory make similar predictions about the nonmonotonic relationship between effort investment and task difficulty.

## SilvettiAlexanderVergutsEtAl14 -- basic review paper on ACC upon which theory is based

Summary: too much reliance on single-unit responses -- can find everything everywhere. Need population-level data (fMRI, neuropixels etc) to really see what a given area is uniquely contributing. Anyway, it is a nice review of different theories..

* Error likelihood: In a revision of the classical error detection theory, Brown and Braver (2005) proposed that ACC activity reflects the estimated probability of committing an error. In contrast to the conflict model, the Error Likelihood model predicts that ACC activity increases with the likelihood that, within a given context, a behavioral error will be committed, regardless of whether or not an error actually occurs.

* Consistent with the error likelihood model, increased ACC activity was observed in correctly solved trials associated with a higher probability of error, even on those trials in which a subject did not receive the change signal. However, this pattern was reversed in error trials: ACC activity following response errors was greater for task conditions in which the likelihood of committing an error was low, as compared to high-error likelihood conditions (Brown and Braver, 2005). 

Key point:

> This latter finding, incompatible with both the error likelihood and conflict accounts of ACC, would not be explained computationally for several years, as described in later sections.

* Energization: Deriving from the longstanding clinical observation of akinetic mutism following ACC lesions (Németh et al., 1988), a further theory on ACC cognitive functions is that it energizes the cognitive system when effort needs to be exerted (e.g. Kouneiher et al., 2009). Despite the imprecision of these concepts, it is an empirical finding that situations requiring high effort robustly activate ACC (e.g. Sohn et al., 2007; Krebs et al., 2012). 

* RL framework: These neurons coded for global value of expected reward, i.e. their discharge rate was modulated by both reward probability and expected reward magnitude. Furthermore, ACC cells incorporate timing information, such that their activity increased as the time of anticipated reward drew closer (Shidara and Richmond, 2002).

> BUT THIS IS COMING FROM OFC! ACC as a population code is more about action-value than outcome value

* The functional organization of ACC according to a rostro-caudal gradient suggests that the most rostral and caudal portions of the ACC (BA 24 and 32) perform the role of Actor, providing reward-based motor commands

* Value of action - cost: One of the main features of ACC is that its neural populations estimate the value of a specific action or stimulus after cost discounting, i.e. integrating the information about both reward magnitude/probability and estimated costs (Kennerley et al., 2011). 

* PRO model & negative valence prediction error: Alexander and Brown propose that specifically this second component (ωN) is responsible for most of the classical cognitive neuroscience results on ACC functions.

- Randy

