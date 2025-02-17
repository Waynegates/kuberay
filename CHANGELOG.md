# Change Log

All notable changes to this project will be documented in this file.
 
The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [v0.2.0](https://github.com/ray-project/kuberay/tree/v0.2.0) (2022-03-13)

### Features

* Support envFrom in rayclusters deployed with Helm ([#183](https://github.com/ray-project/kuberay/pull/183), @ebr)
* Helm: support imagePullSecrets for ray clusters ([#182](https://github.com/ray-project/kuberay/pull/182), @ebr)
* Support scheduling constraints in Helm-deployed clusters ([#181](https://github.com/ray-project/kuberay/pull/181), @ebr)
* Helm: ensure RBAC rules are up to date with the latest autogenerated manifest ([#175](https://github.com/ray-project/kuberay/pull/175), @ebr)
* add resource command ([#170](https://github.com/ray-project/kuberay/pull/170), @zhuangzhuang131419)
* Use container to generate proto files ([#160](https://github.com/ray-project/kuberay/pull/160), @Jeffwan)
* Support in-tree autoscaler ([#163](https://github.com/ray-project/kuberay/pull/163), @Jeffwan)
* [CLI] check viper error ([#172](https://github.com/ray-project/kuberay/pull/172), @chenk008)
* [Feature]Add subcommand `--version` ([#166](https://github.com/ray-project/kuberay/pull/166), @chenk008)
* [Feature] Add flag `watch-namespace` ([#165](https://github.com/ray-project/kuberay/pull/165), @chenk008)
* Support enableIngress for RayCluster ([#38](https://github.com/ray-project/kuberay/pull/38), @Jeffwan)
* Add CRD verb permission in helm ([#144](https://github.com/ray-project/kuberay/pull/144), @chenk008)
* Add quick start deployment manifests ([#132](https://github.com/ray-project/kuberay/pull/132), @Jeffwan)
* Add CLI to kuberay ([#135](https://github.com/ray-project/kuberay/pull/135), @wolfsniper2388)
* Ray Operator: Upgrade to Go v1.17 ([#128](https://github.com/ray-project/kuberay/pull/128), @haoxins)
* Add deploy manifests for apiserver ([#119](https://github.com/ray-project/kuberay/pull/119), @Jeffwan)
* Implement resource manager and gRPC services ([#127](https://github.com/ray-project/kuberay/pull/127), @Jeffwan)
* Generate go clients and swagger files ([#126](https://github.com/ray-project/kuberay/pull/126), @Jeffwan)
* [service] Init backend service project ([#113](https://github.com/ray-project/kuberay/pull/113), @Jeffwan)
* Add gRPC service definition and gRPC gateway ([#112](https://github.com/ray-project/kuberay/pull/112), @Jeffwan)
* [proto] Add core api definitions as protobuf message ([#93](https://github.com/ray-project/kuberay/pull/93), @Jeffwan)
* Use ray start block in Pod's entrypoint ([#77](https://github.com/ray-project/kuberay/pull/77), @chenk008)
* Add generated clientsets, informers and listers ([#97](https://github.com/ray-project/kuberay/pull/97), @Jeffwan)
* Add codegen scripts and make required api changes ([#96](https://github.com/ray-project/kuberay/pull/96), @harryge00)
* Reorganize api folder for code generation ([#91](https://github.com/ray-project/kuberay/pull/91), @harryge00)

### Bug fixes

* Fix serviceaccount typo in operator role ([#188](https://github.com/ray-project/kuberay/pull/188), @Jeffwan)
* Fix cli typo ([#173](https://github.com/ray-project/kuberay/pull/173), @chenk008)
* [Bug]Leader election need lease permission ([#169](https://github.com/ray-project/kuberay/pull/169), @chenk008)
* refactor: rename kubray -> kuberay ([#145](https://github.com/ray-project/kuberay/pull/145), @tekumara)
* Fix the Helm chart's image name ([#130](https://github.com/ray-project/kuberay/pull/130), @haoxins)
* fix typo in the helm chart templates ([#129](https://github.com/ray-project/kuberay/pull/129), @haoxins)
* fix issue that modifies the list while iterating through it ([#125](https://github.com/ray-project/kuberay/pull/125), @wilsonwang371)
* Add helm ([#109](https://github.com/ray-project/kuberay/pull/109), @zhuangzhuang131419)
* Update samples yaml ([#102](https://github.com/ray-project/kuberay/pull/102), @ryantd)
* fix missing template objectmeta ([#95](https://github.com/ray-project/kuberay/pull/95), @chenk008)
* fix typo in Readme ([#81](https://github.com/ray-project/kuberay/pull/81), @denkensk)

### Testing

* kuberay compatibility test with ray ([#157](https://github.com/ray-project/kuberay/pull/157), @wilsonwang371)
* Setup ci for apiserver ([#162](https://github.com/ray-project/kuberay/pull/162), @Jeffwan)
* Enable gofmt and move goimports to linter job ([#158](https://github.com/ray-project/kuberay/pull/158), @Jeffwan)
* add more debug info for bug-150: goimport issue ([#151](https://github.com/ray-project/kuberay/pull/151), @wilsonwang371)
* add nightly docker build workflow ([#141](https://github.com/ray-project/kuberay/pull/141), @wilsonwang371)
* enable goimport and add new makefile target to only build image without test ([#123](https://github.com/ray-project/kuberay/pull/123), @wilsonwang371)
* [Feature]add docker build stage to ci workflow ([#122](https://github.com/ray-project/kuberay/pull/122), @wilsonwang371)
* Pass --timeout option to golangci-lint ([#116](https://github.com/ray-project/kuberay/pull/116), @Jeffwan)
* Add linter job for github workflow ([#79](https://github.com/ray-project/kuberay/pull/79), @feilengcui008)

### Docs and Miscs

* Add Makefile for cli project ([#192](https://github.com/ray-project/kuberay/pull/192), @Jeffwan)
* Manifests and docs improvement for prerelease  ([#191](https://github.com/ray-project/kuberay/pull/191), @Jeffwan)
* Add documentation for autoscaling feature ([#189](https://github.com/ray-project/kuberay/pull/189), @Jeffwan)
* docs: Fix typo in best practice ([#190](https://github.com/ray-project/kuberay/pull/190), @nakamasato)
* add kuberay on kind jupyter notebook ([#147](https://github.com/ray-project/kuberay/pull/147), @wilsonwang371)
* Add KubeRay release guideline ([#161](https://github.com/ray-project/kuberay/pull/161), @Jeffwan)
* Add troubleshooting guide for ray version mismatch ([#154](https://github.com/ray-project/kuberay/pull/154), @scarlet25151)
* Explanation and Best Practice for workers-head Reconnection ([#142](https://github.com/ray-project/kuberay/pull/142), @nostalgicimp)
* [docs] Folder name change to kuberay-operator ([#143](https://github.com/ray-project/kuberay/pull/143), @asm582)
* Improve the Helm charts docs ([#131](https://github.com/ray-project/kuberay/pull/131), @haoxins)
* add auto-scale doc ([#108](https://github.com/ray-project/kuberay/pull/108), @akanso)
* Add core API and backend service design doc ([#98](https://github.com/ray-project/kuberay/pull/98), @Jeffwan)
* [Feature] add more options in bug template ([#121](https://github.com/ray-project/kuberay/pull/121), @wilsonwang371)
* Rename service module to apiserver ([#118](https://github.com/ray-project/kuberay/pull/118), @Jeffwan)


## [v0.1.0](https://github.com/ray-project/kuberay/tree/v0.1.0) (2021-10-16)

### Feature

* Check duplicate services explicitly ([#72](https://github.com/ray-project/kuberay/pull/72), @Jeffwan)
* Expose reconcile concurrency as a command flag ([#67](https://github.com/ray-project/kuberay/pull/67), @feilengcui008)
* Ignore reconcile cluster being deleted ([#63](https://github.com/ray-project/kuberay/pull/63), @feilengcui008)
* Add issue and pr templates ([#44](https://github.com/ray-project/kuberay/pull/44), @chaomengyuan)
* Create root level .gitignore file ([#37](https://github.com/ray-project/kuberay/pull/37), @Jeffwan)
* Remove BAZEL build in ray-operator project ([#32](https://github.com/ray-project/kuberay/pull/32), @chenk008)
* Upgrade Kubebuilder to 3.0.0 and optimize Github workflow ([#31](https://github.com/ray-project/kuberay/pull/31), @Jeffwan)
* Update v1alpha1 RayCluster CRD and controllers ([#22](https://github.com/ray-project/kuberay/pull/22), @Jeffwan)
* Deprecate msft operator and rename to ray-operator ([#20](https://github.com/ray-project/kuberay/pull/20), @Jeffwan)
* Deprecate ByteDance operator and move to unified one ([#19](https://github.com/ray-project/kuberay/pull/19), @Jeffwan)
* Deprecate antgroup ray operator and move to unified implementation ([#18](https://github.com/ray-project/kuberay/pull/18), @chenk008)
* Upgrade to go 1.15 ([#12](https://github.com/ray-project/kuberay/pull/12), @tgaddair)
* Remove unused generated manifest from kubebuilder ([#11](https://github.com/ray-project/kuberay/pull/11), @Jeffwan)
* Clean up kustomization manifests ([#10](https://github.com/ray-project/kuberay/pull/10), @Jeffwan)
* Add RayCluster v1alpha1 controller ([#8](https://github.com/ray-project/kuberay/pull/8), @Jeffwan)
* Scaffolding out Bytedance's ray operator project  ([#7](https://github.com/ray-project/kuberay/pull/7), @Jeffwan)
* allow deletion of workers ([#5](https://github.com/ray-project/kuberay/pull/5), @akanso)
* Ray Autoscaler integrate with Ray K8s Operator ([#2](https://github.com/ray-project/kuberay/pull/2), @Qstar)
* Add license ([#3](https://github.com/ray-project/kuberay/pull/3), @akanso)
* Operator with Design 1B ([#1](https://github.com/ray-project/kuberay/pull/1), @akanso)

### Bugs

* Fix flaky tests by retrying 409 conflict error ([#73](https://github.com/ray-project/kuberay/pull/73), @Jeffwan)
* Fix issues in heterogeneous sample ([#45](https://github.com/ray-project/kuberay/pull/45), @anencore94)
* Fix incorrect manifest setting and remove unused manifests ([#34](https://github.com/ray-project/kuberay/pull/34), @Jeffwan)
* Fix status update issue and redis port formatting issue ([#16](https://github.com/ray-project/kuberay/pull/16), @Jeffwan)
* Fix leader election failure and crd too long issue ([#9](https://github.com/ray-project/kuberay/pull/9), @Jeffwan)

### Misc

* chore: Add github workflow and go report badge ([#58](https://github.com/ray-project/kuberay/pull/58), @Jeffwan)
* Update README.md ([#52](https://github.com/ray-project/kuberay/pull/52), @Jeffwan)
* Add community profile documentation ([#49](https://github.com/ray-project/kuberay/pull/49), @Jeffwan)
* Rename go module to kuberay ([#50](https://github.com/ray-project/kuberay/pull/50), @Jeffwan)
