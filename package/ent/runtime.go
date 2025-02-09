// Code generated by ent, DO NOT EDIT.

package ent

import (
	"ekko/package/ent/answersubmission"
	"ekko/package/ent/question"
	"ekko/package/ent/scenario"
	"ekko/package/ent/scenariocandidate"
	"ekko/package/ent/submissionattempt"
	"ekko/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	answersubmissionMixin := schema.AnswerSubmission{}.Mixin()
	answersubmissionMixinFields0 := answersubmissionMixin[0].Fields()
	_ = answersubmissionMixinFields0
	answersubmissionFields := schema.AnswerSubmission{}.Fields()
	_ = answersubmissionFields
	// answersubmissionDescCreatedAt is the schema descriptor for created_at field.
	answersubmissionDescCreatedAt := answersubmissionMixinFields0[1].Descriptor()
	// answersubmission.DefaultCreatedAt holds the default value on creation for the created_at field.
	answersubmission.DefaultCreatedAt = answersubmissionDescCreatedAt.Default.(func() time.Time)
	// answersubmissionDescUpdatedAt is the schema descriptor for updated_at field.
	answersubmissionDescUpdatedAt := answersubmissionMixinFields0[2].Descriptor()
	// answersubmission.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	answersubmission.DefaultUpdatedAt = answersubmissionDescUpdatedAt.Default.(func() time.Time)
	// answersubmission.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	answersubmission.UpdateDefaultUpdatedAt = answersubmissionDescUpdatedAt.UpdateDefault.(func() time.Time)
	questionMixin := schema.Question{}.Mixin()
	questionMixinFields0 := questionMixin[0].Fields()
	_ = questionMixinFields0
	questionFields := schema.Question{}.Fields()
	_ = questionFields
	// questionDescCreatedAt is the schema descriptor for created_at field.
	questionDescCreatedAt := questionMixinFields0[1].Descriptor()
	// question.DefaultCreatedAt holds the default value on creation for the created_at field.
	question.DefaultCreatedAt = questionDescCreatedAt.Default.(func() time.Time)
	// questionDescUpdatedAt is the schema descriptor for updated_at field.
	questionDescUpdatedAt := questionMixinFields0[2].Descriptor()
	// question.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	question.DefaultUpdatedAt = questionDescUpdatedAt.Default.(func() time.Time)
	// question.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	question.UpdateDefaultUpdatedAt = questionDescUpdatedAt.UpdateDefault.(func() time.Time)
	scenarioMixin := schema.Scenario{}.Mixin()
	scenarioMixinFields0 := scenarioMixin[0].Fields()
	_ = scenarioMixinFields0
	scenarioFields := schema.Scenario{}.Fields()
	_ = scenarioFields
	// scenarioDescCreatedAt is the schema descriptor for created_at field.
	scenarioDescCreatedAt := scenarioMixinFields0[1].Descriptor()
	// scenario.DefaultCreatedAt holds the default value on creation for the created_at field.
	scenario.DefaultCreatedAt = scenarioDescCreatedAt.Default.(func() time.Time)
	// scenarioDescUpdatedAt is the schema descriptor for updated_at field.
	scenarioDescUpdatedAt := scenarioMixinFields0[2].Descriptor()
	// scenario.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	scenario.DefaultUpdatedAt = scenarioDescUpdatedAt.Default.(func() time.Time)
	// scenario.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	scenario.UpdateDefaultUpdatedAt = scenarioDescUpdatedAt.UpdateDefault.(func() time.Time)
	scenariocandidateMixin := schema.ScenarioCandidate{}.Mixin()
	scenariocandidateMixinFields0 := scenariocandidateMixin[0].Fields()
	_ = scenariocandidateMixinFields0
	scenariocandidateFields := schema.ScenarioCandidate{}.Fields()
	_ = scenariocandidateFields
	// scenariocandidateDescCreatedAt is the schema descriptor for created_at field.
	scenariocandidateDescCreatedAt := scenariocandidateMixinFields0[1].Descriptor()
	// scenariocandidate.DefaultCreatedAt holds the default value on creation for the created_at field.
	scenariocandidate.DefaultCreatedAt = scenariocandidateDescCreatedAt.Default.(func() time.Time)
	// scenariocandidateDescUpdatedAt is the schema descriptor for updated_at field.
	scenariocandidateDescUpdatedAt := scenariocandidateMixinFields0[2].Descriptor()
	// scenariocandidate.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	scenariocandidate.DefaultUpdatedAt = scenariocandidateDescUpdatedAt.Default.(func() time.Time)
	// scenariocandidate.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	scenariocandidate.UpdateDefaultUpdatedAt = scenariocandidateDescUpdatedAt.UpdateDefault.(func() time.Time)
	submissionattemptMixin := schema.SubmissionAttempt{}.Mixin()
	submissionattemptMixinFields0 := submissionattemptMixin[0].Fields()
	_ = submissionattemptMixinFields0
	submissionattemptFields := schema.SubmissionAttempt{}.Fields()
	_ = submissionattemptFields
	// submissionattemptDescCreatedAt is the schema descriptor for created_at field.
	submissionattemptDescCreatedAt := submissionattemptMixinFields0[1].Descriptor()
	// submissionattempt.DefaultCreatedAt holds the default value on creation for the created_at field.
	submissionattempt.DefaultCreatedAt = submissionattemptDescCreatedAt.Default.(func() time.Time)
	// submissionattemptDescUpdatedAt is the schema descriptor for updated_at field.
	submissionattemptDescUpdatedAt := submissionattemptMixinFields0[2].Descriptor()
	// submissionattempt.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	submissionattempt.DefaultUpdatedAt = submissionattemptDescUpdatedAt.Default.(func() time.Time)
	// submissionattempt.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	submissionattempt.UpdateDefaultUpdatedAt = submissionattemptDescUpdatedAt.UpdateDefault.(func() time.Time)
}
