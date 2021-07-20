# Composition Revisions

* Owner: Nic Cope (@negz)
* Reviewers: Crossplane Maintainers
* Status: DRAFT

## Background

In Crossplane _Composition_ allows platform teams to define and offer bespoke
infrastructure APIs to the teams of application developers they support. These
APIs are known as _Composite Resources_ (XRs). Crossplane powers each XR by
composing one or more _Managed Resources_ (MRs). When an XR is created
Crossplane uses a `Composition` to determine which MRs are required to satisfy
the XR. Note that Composition is used in two ways here; _Composition_ is the
name of the feature, while "a `Composition`" is one of the Crossplane resources
that configures the feature.

Platform engineers can use Crossplane to define an arbitrary number of XR types,
and an arbitrary number of `Compositions`. Each `Composition` declares that it
satisfies a particular type of XR, in the sense that the `Composition` tells
Crossplane what resources should be composed to _satisfy_ the XR's desired
state. Any XR can be satisfied by one `Composition` at any point in time. There
is a one-to-many relationship between a `Composition` and the XRs that it
satisfies.

![XR to Composition relationship][xr-to-composition]

Note that in the above diagram the `example-a` and `example-b` `CompositeWidget`
XRs are both satisfied by one `Composition`; `large`. Meanwhile `example-c` is
satisfied by a different `Composition`.

Today it is possible to update a `Composition` in place, but doing so is risky.
_All_ XRs that use said `Composition` will be updated instantaneously. These
updates will often be surprising, because the party making the update and the
parties affected by the update will typically be different people. That is,
typically a platform engineer would update the `Composition` and that update
would instantly cause changes to various XRs provisioned and owned by app teams.

Ideally it would be possible for an updated `Composition` to be introduced then
rolled out in a controlled fashion to the various XRs it satisfies. It should be
possible to do this in a fashion that enables the separation of concerns; i.e.
to support one team (typically the platform team) introducing a new
`Composition` and potentially another team (e.g. an app team) choosing when
their XR should start consuming that `Composition`.

## Goals

Functionality wise, this design intends to:

* Allow a `Composition` that is in use to be updated in a measured fashion.
* Respect the separation of concerns; don't assume that the person introducing
  the new `Composition` is the same person who will update the XRs that consume
  it.

It must be possible to introduce this functionality in a measured, backward
compatible way. Crossplane's behaviour and v1 APIs should not change for anyone
who does not opt into this new functionality.

## Proposal

```yaml
apiVersion: database.example.org/v1alpha1
kind: PostgreSQLInstance
metadata:
  name: my-db
  namespace: default
spec:
  parameters:
    storageGB: 20
  compositionSelector:
    matchLabels:
      provider: gcp
  compositionRef:
    name: example-gcp
  # A new optional field, typically set automatically to the latest revision of
  # the chosen composition.
  compositionRevisionRef:
    name: example-gcp-fk2ks
  # Automatic or Manual. Defaults to Automatic.
  revisionActivationPolicy: Automatic
```

TODO

* Can we feature flag revision support without a ton of code duplication?
* Should we allow XRDs to enforce a revision activation policy, e.g. to require
  automatic activation of an enforced composition?

## Prior Art

* Crossplane [package revisions][packages-v2]
* Metacontroller [controller revisions][metacontroller-controller-revisions]
* Kubernetes [controller revisions][kubernetes-controller-revisions]

## Alternatives Considered

* Treat `Compositions` as create-only; require composed resources to be updated
  directly in-place after creation.
* Prefer forking and introducing new `Compositions` rather than updating them.

[packages-v2]: design-doc-packages-v2.md
[metacontroller-controller-revisions]: https://metacontroller.github.io/metacontroller/api/controllerrevision.html
[kubernetes-controller-revisions]: https://kubernetes.io/docs/tasks/manage-daemon/rollback-daemon-set/#understanding-daemonset-revisions
[xr-to-composition]: images/xr-to-composition.svg