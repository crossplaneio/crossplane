/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/crossplane/crossplane-runtime/pkg/meta"

	"github.com/crossplane/crossplane/apis/apiextensions/v1alpha1"
)

// NewCompositionRevision creates a new revision of the supplied Composition.
func NewCompositionRevision(c Composition) v1alpha1.CompositionRevision {
	cr := &v1alpha1.CompositionRevision{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: c.GetName() + "-",
		},
		Spec: NewCompositionRevisionSpec(c.Spec),
	}

	// TODO(negz): Determine and set revision number. This may require
	// moving this logic into a Composition controller, but we'd also like
	// to use it when the CompositionRevision feature flag is disabled.

	ref := meta.TypedReferenceTo(&c, CompositionGroupVersionKind)
	_ = meta.AddControllerReference(cr, meta.AsOwner(ref))

	return *cr
}

// NewCompositionRevisionSpec translates a composition's spec to a composition
// revision spec.
func NewCompositionRevisionSpec(cs CompositionSpec) v1alpha1.CompositionRevisionSpec {
	rs := v1alpha1.CompositionRevisionSpec{
		CompositeTypeRef: v1alpha1.TypeReference{
			APIVersion: cs.CompositeTypeRef.APIVersion,
			Kind:       cs.CompositeTypeRef.Kind,
		},
		PatchSets:                         make([]v1alpha1.PatchSet, len(cs.PatchSets)),
		Resources:                         make([]v1alpha1.ComposedTemplate, len(cs.Resources)),
		WriteConnectionSecretsToNamespace: cs.WriteConnectionSecretsToNamespace,
	}

	for i := range cs.PatchSets {
		rs.PatchSets[i] = NewCompositionRevisionPatchSet(cs.PatchSets[i])
	}

	for i := range cs.Resources {
		rs.Resources[i] = NewCompositionRevisionComposedTemplate(cs.Resources[i])
	}

	return rs
}

// NewCompositionRevisionPatchSet translates a composition's patch set to a
// composition revision patch set.
func NewCompositionRevisionPatchSet(ps PatchSet) v1alpha1.PatchSet {
	rps := v1alpha1.PatchSet{
		Name:    ps.Name,
		Patches: make([]v1alpha1.Patch, len(ps.Patches)),
	}

	for i := range ps.Patches {
		rps.Patches[i] = NewCompositionRevisionPatch(ps.Patches[i])
	}
	return rps
}

// NewCompositionRevisionComposedTemplate translates a composition's composed
// (resource) template to a composition composed template.
func NewCompositionRevisionComposedTemplate(ct ComposedTemplate) v1alpha1.ComposedTemplate {
	rct := v1alpha1.ComposedTemplate{
		Name: ct.Name,
		Base: ct.Base,
	}

	for i := range ct.Patches {
		rct.Patches[i] = NewCompositionRevisionPatch(ct.Patches[i])
	}

	for i := range ct.ConnectionDetails {
		rct.ConnectionDetails[i] = NewCompositionRevisionConnectionDetail(ct.ConnectionDetails[i])
	}

	for i := range ct.ReadinessChecks {
		rct.ReadinessChecks[i] = NewCompositionRevisionReadinessCheck(ct.ReadinessChecks[i])
	}

	return rct
}

// NewCompositionRevisionPatch translates a composition's patch to a
// composition revision patch.
func NewCompositionRevisionPatch(p Patch) v1alpha1.Patch {
	rp := v1alpha1.Patch{
		Type:          v1alpha1.PatchType(p.Type),
		FromFieldPath: p.FromFieldPath,
		Combine:       &v1alpha1.Combine{},
		ToFieldPath:   p.ToFieldPath,
		PatchSetName:  p.PatchSetName,
	}

	if p.Combine != nil {
		rp.Combine = &v1alpha1.Combine{
			Strategy:  v1alpha1.CombineStrategy(p.Combine.Strategy),
			Variables: make([]v1alpha1.CombineVariable, len(p.Combine.Variables)),
		}

		if p.Combine.String != nil {
			rp.Combine.String.Format = p.Combine.String.Format
		}

		for i := range p.Combine.Variables {
			rp.Combine.Variables[i].FromFieldPath = p.Combine.Variables[i].FromFieldPath
		}
	}

	for i := range p.Transforms {
		rp.Transforms[i] = NewCompositionRevisionTransform(p.Transforms[i])
	}

	if p.Policy != nil && p.Policy.FromFieldPath != nil {
		pol := v1alpha1.FromFieldPathPolicy(*p.Policy.FromFieldPath)
		rp.Policy.FromFieldPath = &pol
	}

	return rp
}

// NewCompositionRevisionTransform translates a compostion's transform to a
// composition revision transform.
func NewCompositionRevisionTransform(t Transform) v1alpha1.Transform {
	rt := v1alpha1.Transform{Type: v1alpha1.TransformType(t.Type)}
	if t.Math != nil {
		rt.Math = &v1alpha1.MathTransform{Multiply: t.Math.Multiply}
	}
	if t.Map != nil {
		rt.Map = &v1alpha1.MapTransform{Pairs: t.Map.Pairs}
	}
	if t.String != nil {
		rt.String = &v1alpha1.StringTransform{Format: t.String.Format}
	}
	if t.Convert != nil {
		rt.Convert = &v1alpha1.ConvertTransform{ToType: t.Convert.ToType}
	}
	return rt
}

// NewCompositionRevisionConnectionDetail translates a composition's connection
// detail to a composition revision connection detail.
func NewCompositionRevisionConnectionDetail(cd ConnectionDetail) v1alpha1.ConnectionDetail {
	return v1alpha1.ConnectionDetail{
		Name: cd.Name,
		Type: func() *v1alpha1.ConnectionDetailType {
			if cd.Type == nil {
				return nil
			}
			t := v1alpha1.ConnectionDetailType(*cd.Type)
			return &t
		}(),
		FromConnectionSecretKey: cd.FromConnectionSecretKey,
		FromFieldPath:           cd.FromFieldPath,
		Value:                   cd.Value,
	}
}

// NewCompositionRevisionReadinessCheck translates a composition's readiness
// check to a composition revision readiness check.
func NewCompositionRevisionReadinessCheck(rc ReadinessCheck) v1alpha1.ReadinessCheck {
	return v1alpha1.ReadinessCheck{
		Type:         v1alpha1.ReadinessCheckType(rc.Type),
		FieldPath:    rc.FieldPath,
		MatchString:  rc.MatchString,
		MatchInteger: rc.MatchInteger,
	}
}
