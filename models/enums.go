package models

// BundleType is a FHIR enum.
type BundleType string

const (
	BundleTypeDocument            BundleType = "document"
	BundleTypeMessage             BundleType = "message"
	BundleTypeTransaction         BundleType = "transaction"
	BundleTypeTransactionResponse BundleType = "transaction-response"
	BundleTypeBatch               BundleType = "batch"
	BundleTypeBatchResponse       BundleType = "batch-response"
	BundleTypeHistory             BundleType = "history"
	BundleTypeSearchset           BundleType = "searchset"
	BundleTypeCollection          BundleType = "collection"
)

// IsValid ...
func (e BundleType) IsValid() bool {
	switch e {
	case BundleTypeDocument, BundleTypeMessage, BundleTypeTransaction,
		BundleTypeTransactionResponse, BundleTypeBatch, BundleTypeBatchResponse,
		BundleTypeHistory, BundleTypeSearchset, BundleTypeCollection:
		return true
	}

	return false
}

// String ...
func (e BundleType) String() string {
	return string(e)
}

// SearchEntryMode is a FHIR enum.
type SearchEntryMode string

const (
	SearchEntryModeMatch   SearchEntryMode = "match"
	SearchEntryModeInclude SearchEntryMode = "include"
	SearchEntryModeOutcome SearchEntryMode = "outcome"
)

// IsValid ...
func (e SearchEntryMode) IsValid() bool {
	switch e {
	case SearchEntryModeMatch, SearchEntryModeInclude, SearchEntryModeOutcome:
		return true
	}

	return false
}

// String ...
func (e SearchEntryMode) String() string {
	return string(e)
}

// EpisodeOfCareStatus is a FHIR enum.
type EpisodeOfCareStatus string

const (
	EpisodeOfCareStatusPlanned        EpisodeOfCareStatus = "planned"
	EpisodeOfCareStatusWaitlist       EpisodeOfCareStatus = "waitlist"
	EpisodeOfCareStatusActive         EpisodeOfCareStatus = "active"
	EpisodeOfCareStatusOnhold         EpisodeOfCareStatus = "onhold"
	EpisodeOfCareStatusFinished       EpisodeOfCareStatus = "finished"
	EpisodeOfCareStatusCancelled      EpisodeOfCareStatus = "cancelled"
	EpisodeOfCareStatusEnteredInError EpisodeOfCareStatus = "entered_in_error"
)

// IsValid ...
func (e EpisodeOfCareStatus) IsValid() bool {
	switch e {
	case EpisodeOfCareStatusPlanned, EpisodeOfCareStatusWaitlist, EpisodeOfCareStatusActive,
		EpisodeOfCareStatusOnhold, EpisodeOfCareStatusFinished, EpisodeOfCareStatusCancelled, EpisodeOfCareStatusEnteredInError:
		return true
	}

	return false
}

// String ...
func (e EpisodeOfCareStatus) String() string {
	return string(e)
}

// EncounterStatus is a FHIR enum.
type EncounterStatus string

const (
	EncounterStatusPlanned        EncounterStatus = "planned"
	EncounterStatusArrived        EncounterStatus = "arrived"
	EncounterStatusTriaged        EncounterStatus = "triaged"
	EncounterStatusInProgress     EncounterStatus = "in_progress"
	EncounterStatusOnleave        EncounterStatus = "onleave"
	EncounterStatusFinished       EncounterStatus = "finished"
	EncounterStatusCancelled      EncounterStatus = "cancelled"
	EncounterStatusEnteredInError EncounterStatus = "entered_in_error"
	EncounterStatusUnknown        EncounterStatus = "unknown"
)

// IsValid ...
func (e EncounterStatus) IsValid() bool {
	switch e {
	case EncounterStatusPlanned, EncounterStatusArrived, EncounterStatusTriaged,
		EncounterStatusInProgress, EncounterStatusOnleave, EncounterStatusFinished,
		EncounterStatusCancelled, EncounterStatusEnteredInError, EncounterStatusUnknown:
		return true
	}

	return false
}

// String ...
func (e EncounterStatus) String() string {
	return string(e)
}

// EncounterLocationStatus is a FHIR enum.
type EncounterLocationStatus string

const (
	EncounterLocationStatusPlanned   EncounterLocationStatus = "planned"
	EncounterLocationStatusActive    EncounterLocationStatus = "active"
	EncounterLocationStatusReserved  EncounterLocationStatus = "reserved"
	EncounterLocationStatusCompleted EncounterLocationStatus = "completed"
)

// IsValid ...
func (e EncounterLocationStatus) IsValid() bool {
	switch e {
	case EncounterLocationStatusPlanned, EncounterLocationStatusActive,
		EncounterLocationStatusReserved, EncounterLocationStatusCompleted:
		return true
	}

	return false
}

// String ...
func (e EncounterLocationStatus) String() string {
	return string(e)
}

// ObservationStatus is a FHIR enum.
type ObservationStatus string

const (
	ObservationStatusRegistered     ObservationStatus = "registered"
	ObservationStatusPreliminary    ObservationStatus = "preliminary"
	ObservationStatusFinal          ObservationStatus = "final"
	ObservationStatusAmended        ObservationStatus = "amended"
	ObservationStatusCorrected      ObservationStatus = "corrected"
	ObservationStatusCancelled      ObservationStatus = "cancelled"
	ObservationStatusEnteredInError ObservationStatus = "entered_in_error"
	ObservationStatusUnknown        ObservationStatus = "unknown"
)

// IsValid ...
func (e ObservationStatus) IsValid() bool {
	switch e {
	case ObservationStatusRegistered, ObservationStatusPreliminary, ObservationStatusFinal,
		ObservationStatusAmended, ObservationStatusCorrected, ObservationStatusCancelled,
		ObservationStatusEnteredInError, ObservationStatusUnknown:
		return true
	}

	return false
}

// String ...
func (e ObservationStatus) String() string {
	return string(e)
}

// TimingRepeatWhenEnum is a FHIR enum.
type TimingRepeatWhenEnum string

const (
	TimingRepeatWhenEnumMorn      TimingRepeatWhenEnum = "MORN"
	TimingRepeatWhenEnumMornEarly TimingRepeatWhenEnum = "MORN_early"
	TimingRepeatWhenEnumMornLate  TimingRepeatWhenEnum = "MORN_late"
	TimingRepeatWhenEnumNoon      TimingRepeatWhenEnum = "NOON"
	TimingRepeatWhenEnumAft       TimingRepeatWhenEnum = "AFT"
	TimingRepeatWhenEnumAftEarly  TimingRepeatWhenEnum = "AFT_early"
	TimingRepeatWhenEnumAftLate   TimingRepeatWhenEnum = "AFT_late"
	TimingRepeatWhenEnumEve       TimingRepeatWhenEnum = "EVE"
	TimingRepeatWhenEnumEveEarly  TimingRepeatWhenEnum = "EVE_early"
	TimingRepeatWhenEnumEveLate   TimingRepeatWhenEnum = "EVE_late"
	TimingRepeatWhenEnumNight     TimingRepeatWhenEnum = "NIGHT"
	TimingRepeatWhenEnumPhs       TimingRepeatWhenEnum = "PHS"
	TimingRepeatWhenEnumHs        TimingRepeatWhenEnum = "HS"
	TimingRepeatWhenEnumWake      TimingRepeatWhenEnum = "WAKE"
	TimingRepeatWhenEnumC         TimingRepeatWhenEnum = "C"
	TimingRepeatWhenEnumCm        TimingRepeatWhenEnum = "CM"
	TimingRepeatWhenEnumCd        TimingRepeatWhenEnum = "CD"
	TimingRepeatWhenEnumCv        TimingRepeatWhenEnum = "CV"
	TimingRepeatWhenEnumAc        TimingRepeatWhenEnum = "AC"
	TimingRepeatWhenEnumAcm       TimingRepeatWhenEnum = "ACM"
	TimingRepeatWhenEnumAcd       TimingRepeatWhenEnum = "ACD"
	TimingRepeatWhenEnumAcv       TimingRepeatWhenEnum = "ACV"
	TimingRepeatWhenEnumPc        TimingRepeatWhenEnum = "PC"
	TimingRepeatWhenEnumPcm       TimingRepeatWhenEnum = "PCM"
	TimingRepeatWhenEnumPcd       TimingRepeatWhenEnum = "PCD"
	TimingRepeatWhenEnumPcv       TimingRepeatWhenEnum = "PCV"
)

// IsValid ...
func (e TimingRepeatWhenEnum) IsValid() bool {
	switch e {
	case TimingRepeatWhenEnumMorn, TimingRepeatWhenEnumMornEarly, TimingRepeatWhenEnumMornLate,
		TimingRepeatWhenEnumNoon, TimingRepeatWhenEnumAft, TimingRepeatWhenEnumAftEarly,
		TimingRepeatWhenEnumAftLate, TimingRepeatWhenEnumEve, TimingRepeatWhenEnumEveEarly,
		TimingRepeatWhenEnumEveLate, TimingRepeatWhenEnumNight, TimingRepeatWhenEnumPhs,
		TimingRepeatWhenEnumHs, TimingRepeatWhenEnumWake, TimingRepeatWhenEnumC, TimingRepeatWhenEnumCm,
		TimingRepeatWhenEnumCd, TimingRepeatWhenEnumCv, TimingRepeatWhenEnumAc, TimingRepeatWhenEnumAcm,
		TimingRepeatWhenEnumAcd, TimingRepeatWhenEnumAcv, TimingRepeatWhenEnumPc, TimingRepeatWhenEnumPcm,
		TimingRepeatWhenEnumPcd, TimingRepeatWhenEnumPcv:
		return true
	}

	return false
}

// String ...
func (e TimingRepeatWhenEnum) String() string {
	return string(e)
}

// UnitsOfTime is a FHIR enum.
type UnitsOfTime string

const (
	UnitsOfTimeS   UnitsOfTime = "s"
	UnitsOfTimeMin UnitsOfTime = "min"
	UnitsOfTimeH   UnitsOfTime = "h"
	UnitsOfTimeD   UnitsOfTime = "d"
	UnitsOfTimeWk  UnitsOfTime = "wk"
	UnitsOfTimeMo  UnitsOfTime = "mo"
	UnitsOfTimeA   UnitsOfTime = "a"
)

// IsValid ...
func (e UnitsOfTime) IsValid() bool {
	switch e {
	case UnitsOfTimeS, UnitsOfTimeMin, UnitsOfTimeH, UnitsOfTimeD,
		UnitsOfTimeWk, UnitsOfTimeMo, UnitsOfTimeA:
		return true
	}

	return false
}

// String ...
func (e UnitsOfTime) String() string {
	return string(e)
}

// QuantityComparatorEnum is a FHIR enum.
type QuantityComparatorEnum string

const (
	// QuantityComparatorEnumLessThan ...
	QuantityComparatorEnumLessThan QuantityComparatorEnum = "less_than"
	// QuantityComparatorEnumLessThanOrEqualTo ...
	QuantityComparatorEnumLessThanOrEqualTo QuantityComparatorEnum = "less_than_or_equal_to"
	// QuantityComparatorEnumGreaterThanOrEqualTo ...
	QuantityComparatorEnumGreaterThanOrEqualTo QuantityComparatorEnum = "greater_than_or_equal_to"
	// QuantityComparatorEnumGreaterThan ...
	QuantityComparatorEnumGreaterThan QuantityComparatorEnum = "greater_than"
)

// IsValid ...
func (e QuantityComparatorEnum) IsValid() bool {
	switch e {
	case QuantityComparatorEnumLessThan, QuantityComparatorEnumLessThanOrEqualTo, QuantityComparatorEnumGreaterThanOrEqualTo,
		QuantityComparatorEnumGreaterThan:
		return true
	}

	return false
}

// String ...
func (e QuantityComparatorEnum) String() string {
	switch e {
	case QuantityComparatorEnumLessThan:
		return "<"
	case QuantityComparatorEnumLessThanOrEqualTo:
		return "<="
	case QuantityComparatorEnumGreaterThanOrEqualTo:
		return ">="
	case QuantityComparatorEnumGreaterThan:
		return ">"
	default:
		return string(e)
	}
}

// AddressTypeEnum is a FHIR enum.
type AddressTypeEnum string

const (
	// AddressTypeEnumPostal ...
	AddressTypeEnumPostal AddressTypeEnum = "postal"
	// AddressTypeEnumPhysical ...
	AddressTypeEnumPhysical AddressTypeEnum = "physical"
	// AddressTypeEnumBoth ...
	AddressTypeEnumBoth AddressTypeEnum = "both"
)

// IsValid ...
func (e AddressTypeEnum) IsValid() bool {
	switch e {
	case AddressTypeEnumPostal, AddressTypeEnumPhysical, AddressTypeEnumBoth:
		return true
	}

	return false
}

// String ...
func (e AddressTypeEnum) String() string {
	return string(e)
}

// AddressUseEnum is a FHIR enum.
type AddressUseEnum string

const (
	// AddressUseEnumHome ...
	AddressUseEnumHome AddressUseEnum = "home"
	// AddressUseEnumWork ...
	AddressUseEnumWork AddressUseEnum = "work"
	// AddressUseEnumTemp ...
	AddressUseEnumTemp AddressUseEnum = "temp"
	// AddressUseEnumOld ...
	AddressUseEnumOld AddressUseEnum = "old"
	// AddressUseEnumBilling ...
	AddressUseEnumBilling AddressUseEnum = "billing"
)

// IsValid ...
func (e AddressUseEnum) IsValid() bool {
	switch e {
	case AddressUseEnumHome, AddressUseEnumWork, AddressUseEnumTemp, AddressUseEnumOld, AddressUseEnumBilling:
		return true
	}

	return false
}

// String ...
func (e AddressUseEnum) String() string {
	return string(e)
}

// AgeComparatorEnum is a FHIR enum.
type AgeComparatorEnum string

const (
	// AgeComparatorEnumLessThan ...
	AgeComparatorEnumLessThan AgeComparatorEnum = "less_than"
	// AgeComparatorEnumLessThanOrEqualTo ...
	AgeComparatorEnumLessThanOrEqualTo AgeComparatorEnum = "less_than_or_equal_to"
	// AgeComparatorEnumGreaterThanOrEqualTo ...
	AgeComparatorEnumGreaterThanOrEqualTo AgeComparatorEnum = "greater_than_or_equal_to"
	// AgeComparatorEnumGreaterThan ...
	AgeComparatorEnumGreaterThan AgeComparatorEnum = "greater_than"
)

// IsValid ...
func (e AgeComparatorEnum) IsValid() bool {
	switch e {
	case AgeComparatorEnumLessThan, AgeComparatorEnumLessThanOrEqualTo, AgeComparatorEnumGreaterThanOrEqualTo, AgeComparatorEnumGreaterThan:
		return true
	}

	return false
}

// String renders an age comparator enum as a string.
func (e AgeComparatorEnum) String() string {
	switch e {
	case AgeComparatorEnumLessThan:
		return "<"
	case AgeComparatorEnumLessThanOrEqualTo:
		return "<="
	case AgeComparatorEnumGreaterThanOrEqualTo:
		return ">="
	case AgeComparatorEnumGreaterThan:
		return ">"
	}

	return string(e)
}

// ContactPointSystemEnum is a FHIR enum.
type ContactPointSystemEnum string

const (
	// ContactPointSystemEnumPhone ...
	ContactPointSystemEnumPhone ContactPointSystemEnum = "phone"
	// ContactPointSystemEnumFax ...
	ContactPointSystemEnumFax ContactPointSystemEnum = "fax"
	// ContactPointSystemEnumEmail ...
	ContactPointSystemEnumEmail ContactPointSystemEnum = "email"
	// ContactPointSystemEnumPager ...
	ContactPointSystemEnumPager ContactPointSystemEnum = "pager"
	// ContactPointSystemEnumURL ...
	ContactPointSystemEnumURL ContactPointSystemEnum = "url"
	// ContactPointSystemEnumSms ...
	ContactPointSystemEnumSms ContactPointSystemEnum = "sms"
	// ContactPointSystemEnumOther ...
	ContactPointSystemEnumOther ContactPointSystemEnum = "other"
)

// IsValid ...
func (e ContactPointSystemEnum) IsValid() bool {
	switch e {
	case ContactPointSystemEnumPhone, ContactPointSystemEnumFax,
		ContactPointSystemEnumEmail, ContactPointSystemEnumPager,
		ContactPointSystemEnumURL, ContactPointSystemEnumSms, ContactPointSystemEnumOther:
		return true
	}

	return false
}

// String ...
func (e ContactPointSystemEnum) String() string {
	return string(e)
}

// ContactPointUseEnum is a FHIR enum.
type ContactPointUseEnum string

const (
	// ContactPointUseEnumHome ...
	ContactPointUseEnumHome ContactPointUseEnum = "home"
	// ContactPointUseEnumWork ...
	ContactPointUseEnumWork ContactPointUseEnum = "work"
	// ContactPointUseEnumTemp ...
	ContactPointUseEnumTemp ContactPointUseEnum = "temp"
	// ContactPointUseEnumOld ...
	ContactPointUseEnumOld ContactPointUseEnum = "old"
	// ContactPointUseEnumMobile ...
	ContactPointUseEnumMobile ContactPointUseEnum = "mobile"
)

// IsValid ...
func (e ContactPointUseEnum) IsValid() bool {
	switch e {
	case ContactPointUseEnumHome, ContactPointUseEnumWork,
		ContactPointUseEnumTemp, ContactPointUseEnumOld, ContactPointUseEnumMobile:
		return true
	}

	return false
}

// String ...
func (e ContactPointUseEnum) String() string {
	return string(e)
}

// DurationComparatorEnum is a FHIR enum.
type DurationComparatorEnum string

const (
	// DurationComparatorEnumLessThan ...
	DurationComparatorEnumLessThan DurationComparatorEnum = "less_than"
	// DurationComparatorEnumLessThanOrEqualTo ...
	DurationComparatorEnumLessThanOrEqualTo DurationComparatorEnum = "less_than_or_equal_to"
	// DurationComparatorEnumGreaterThanOrEqualTo ...
	DurationComparatorEnumGreaterThanOrEqualTo DurationComparatorEnum = "greater_than_or_equal_to"
	// DurationComparatorEnumGreaterThan ...
	DurationComparatorEnumGreaterThan DurationComparatorEnum = "greater_than"
)

// IsValid ...
func (e DurationComparatorEnum) IsValid() bool {
	switch e {
	case DurationComparatorEnumLessThan, DurationComparatorEnumLessThanOrEqualTo,
		DurationComparatorEnumGreaterThanOrEqualTo, DurationComparatorEnumGreaterThan:
		return true
	}

	return false
}

// String ...
func (e DurationComparatorEnum) String() string {
	switch e {
	case DurationComparatorEnumLessThan:
		return "<"
	case DurationComparatorEnumLessThanOrEqualTo:
		return "<="
	case DurationComparatorEnumGreaterThan:
		return ">"
	case DurationComparatorEnumGreaterThanOrEqualTo:
		return ">="
	}

	return string(e)
}

// HumanNameUseEnum is a FHIR enum.
type HumanNameUseEnum string

const (
	// HumanNameUseEnumUsual ...
	HumanNameUseEnumUsual HumanNameUseEnum = "usual"
	// HumanNameUseEnumOfficial ...
	HumanNameUseEnumOfficial HumanNameUseEnum = "official"
	// HumanNameUseEnumTemp ...
	HumanNameUseEnumTemp HumanNameUseEnum = "temp"
	// HumanNameUseEnumNickname ...
	HumanNameUseEnumNickname HumanNameUseEnum = "nickname"
	// HumanNameUseEnumAnonymous ...
	HumanNameUseEnumAnonymous HumanNameUseEnum = "anonymous"
	// HumanNameUseEnumOld ...
	HumanNameUseEnumOld HumanNameUseEnum = "old"
	// HumanNameUseEnumMaiden ...
	HumanNameUseEnumMaiden HumanNameUseEnum = "maiden"
)

// IsValid ...
func (e HumanNameUseEnum) IsValid() bool {
	switch e {
	case HumanNameUseEnumUsual, HumanNameUseEnumOfficial,
		HumanNameUseEnumTemp, HumanNameUseEnumNickname, HumanNameUseEnumAnonymous,
		HumanNameUseEnumOld, HumanNameUseEnumMaiden:
		return true
	}

	return false
}

// String ...
func (e HumanNameUseEnum) String() string {
	return string(e)
}

// IdentifierUseEnum is a FHIR enum.
type IdentifierUseEnum string

const (
	// IdentifierUseEnumUsual ...
	IdentifierUseEnumUsual IdentifierUseEnum = "usual"
	// IdentifierUseEnumOfficial ...
	IdentifierUseEnumOfficial IdentifierUseEnum = "official"
	// IdentifierUseEnumTemp ...
	IdentifierUseEnumTemp IdentifierUseEnum = "temp"
	// IdentifierUseEnumSecondary ...
	IdentifierUseEnumSecondary IdentifierUseEnum = "secondary"
	// IdentifierUseEnumOld ...
	IdentifierUseEnumOld IdentifierUseEnum = "old"
)

// IsValid ...
func (e IdentifierUseEnum) IsValid() bool {
	switch e {
	case IdentifierUseEnumUsual, IdentifierUseEnumOfficial, IdentifierUseEnumTemp,
		IdentifierUseEnumSecondary, IdentifierUseEnumOld:
		return true
	}

	return false
}

// String ...
func (e IdentifierUseEnum) String() string {
	return string(e)
}

// PatientGenderEnum is a FHIR enum.
type PatientGenderEnum string

const (
	// PatientGenderEnumMale ...
	PatientGenderEnumMale PatientGenderEnum = "male"
	// PatientGenderEnumFemale ...
	PatientGenderEnumFemale PatientGenderEnum = "female"
	// PatientGenderEnumOther ...
	PatientGenderEnumOther PatientGenderEnum = "other"
	// PatientGenderEnumUnknown ...
	PatientGenderEnumUnknown PatientGenderEnum = "unknown"
)

// IsValid ...
func (e PatientGenderEnum) IsValid() bool {
	switch e {
	case PatientGenderEnumMale, PatientGenderEnumFemale, PatientGenderEnumOther, PatientGenderEnumUnknown:
		return true
	}

	return false
}

// String ...
func (e PatientGenderEnum) String() string {
	return string(e)
}

// PatientContactGenderEnum is a FHIR enum.
type PatientContactGenderEnum string

const (
	// PatientContactGenderEnumMale ...
	PatientContactGenderEnumMale PatientContactGenderEnum = "male"
	// PatientContactGenderEnumFemale ...
	PatientContactGenderEnumFemale PatientContactGenderEnum = "female"
	// PatientContactGenderEnumOther ...
	PatientContactGenderEnumOther PatientContactGenderEnum = "other"
	// PatientContactGenderEnumUnknown ...
	PatientContactGenderEnumUnknown PatientContactGenderEnum = "unknown"
)

// IsValid ...
func (e PatientContactGenderEnum) IsValid() bool {
	switch e {
	case PatientContactGenderEnumMale, PatientContactGenderEnumFemale,
		PatientContactGenderEnumOther, PatientContactGenderEnumUnknown:
		return true
	}

	return false
}

// String ...
func (e PatientContactGenderEnum) String() string {
	return string(e)
}

// PatientLinkTypeEnum is a FHIR enum.
type PatientLinkTypeEnum string

const (
	// PatientLinkTypeEnumReplacedBy ...
	PatientLinkTypeEnumReplacedBy PatientLinkTypeEnum = "replaced_by"
	// PatientLinkTypeEnumReplaces ...
	PatientLinkTypeEnumReplaces PatientLinkTypeEnum = "replaces"
	// PatientLinkTypeEnumRefer ...
	PatientLinkTypeEnumRefer PatientLinkTypeEnum = "refer"
	// PatientLinkTypeEnumSeealso ...
	PatientLinkTypeEnumSeealso PatientLinkTypeEnum = "seealso"
)

// IsValid ...
func (e PatientLinkTypeEnum) IsValid() bool {
	switch e {
	case PatientLinkTypeEnumReplacedBy, PatientLinkTypeEnumReplaces, PatientLinkTypeEnumRefer, PatientLinkTypeEnumSeealso:
		return true
	}

	return false
}

// String ...
func (e PatientLinkTypeEnum) String() string {
	return string(e)
}

// NarrativeStatusEnum is a FHIR enum.
type NarrativeStatusEnum string

const (
	// NarrativeStatusEnumGenerated ...
	NarrativeStatusEnumGenerated NarrativeStatusEnum = "generated"
	// NarrativeStatusEnumExtensions ...
	NarrativeStatusEnumExtensions NarrativeStatusEnum = "extensions"
	// NarrativeStatusEnumAdditional ...
	NarrativeStatusEnumAdditional NarrativeStatusEnum = "additional"
	// NarrativeStatusEnumEmpty ...
	NarrativeStatusEnumEmpty NarrativeStatusEnum = "empty"
)

// IsValid ...
func (e NarrativeStatusEnum) IsValid() bool {
	switch e {
	case NarrativeStatusEnumGenerated, NarrativeStatusEnumExtensions, NarrativeStatusEnumAdditional, NarrativeStatusEnumEmpty:
		return true
	}

	return false
}

// String ...
func (e NarrativeStatusEnum) String() string {
	return string(e)
}

// ServiceRequestType is a custom categorization of service requests.
//
// Service requests are used to create referrals, lab orders etc. These enums are just
// helpful in differentiating.
type ServiceRequestType string

const (
	ReferralServiceRequestType ServiceRequestType = "REFERRAL_SERVICE_REQUEST"
	LabOrderServiceRequestType ServiceRequestType = "LAB_ORDER_SERVICE_REQUEST"
)

// String converts the service request data meaning to string.
func (s ServiceRequestType) String() string {
	switch s {
	case ReferralServiceRequestType:
		return "Patient Referral (Service Request)"
	case LabOrderServiceRequestType:
		return "Laboratory Order (Service Request)"
	default:
		return string(s)
	}
}

// ServiceRequestStatusEnum represents the possible statuses of a ServiceRequest.
type ServiceRequestStatusEnum string

const (
	ServiceRequestStatusDraft          ServiceRequestStatusEnum = "draft"
	ServiceRequestStatusActive         ServiceRequestStatusEnum = "active"
	ServiceRequestStatusOnHold         ServiceRequestStatusEnum = "on-hold"
	ServiceRequestStatusRevoked        ServiceRequestStatusEnum = "revoked"
	ServiceRequestStatusCompleted      ServiceRequestStatusEnum = "completed"
	ServiceRequestStatusEnteredInError ServiceRequestStatusEnum = "entered-in-error"
	ServiceRequestStatusUnknown        ServiceRequestStatusEnum = "unknown"
)

// ServiceRequestIntentEnum represents the possible intents of a ServiceRequest.
type ServiceRequestIntentEnum string

const (
	ServiceRequestIntentProposal      ServiceRequestIntentEnum = "proposal"
	ServiceRequestIntentPlan          ServiceRequestIntentEnum = "plan"
	ServiceRequestIntentDirective     ServiceRequestIntentEnum = "directive"
	ServiceRequestIntentOrder         ServiceRequestIntentEnum = "order"
	ServiceRequestIntentOriginalOrder ServiceRequestIntentEnum = "original-order"
	ServiceRequestIntentReflexOrder   ServiceRequestIntentEnum = "reflex-order"
	ServiceRequestIntentFillerOrder   ServiceRequestIntentEnum = "filler-order"
	ServiceRequestIntentInstanceOrder ServiceRequestIntentEnum = "instance-order"
	ServiceRequestIntentOption        ServiceRequestIntentEnum = "option"
)

// ServiceRequestPriorityEnum represents the possible priorities of a ServiceRequest.
type ServiceRequestPriorityEnum string

const (
	ServiceRequestPriorityRoutine ServiceRequestPriorityEnum = "routine"
	ServiceRequestPriorityUrgent  ServiceRequestPriorityEnum = "urgent"
	ServiceRequestPriorityAsap    ServiceRequestPriorityEnum = "asap"
	ServiceRequestPriorityStat    ServiceRequestPriorityEnum = "stat"
)

// ConsentStatusEnum a type enum tha represents a Consent Status field of consent resource.
type ConsentStatusEnum string

const (
	ConsentStatusActive   ConsentStatusEnum = "active"
	ConsentStatusInactive ConsentStatusEnum = "inactive"
)

// IsValid ...
func (c ConsentStatusEnum) IsValid() bool {
	switch c {
	case ConsentStatusActive, ConsentStatusInactive:
		return true
	}

	return false
}

// String converts status to string.
func (c ConsentStatusEnum) String() string {
	return string(c)
}

// ConsentProvisionTypeEnum a type enum tha represents a Consent Provision field of consent resource.
type ConsentProvisionTypeEnum string

const (
	ConsentProvisionTypeDeny   ConsentProvisionTypeEnum = "deny"
	ConsentProvisionTypePermit ConsentProvisionTypeEnum = "permit"
)

// IsValid ...
func (c ConsentProvisionTypeEnum) IsValid() bool {
	switch c {
	case ConsentProvisionTypeDeny, ConsentProvisionTypePermit:
		return true
	}

	return false
}

// String converts consent provision type to string.
func (c ConsentProvisionTypeEnum) String() string {
	return string(c)
}

// ConsentDataMeaningEnum represents the meaning of consent data.
type ConsentDataMeaningEnum string

const (
	ConsentDataMeaningInstance   ConsentDataMeaningEnum = "instance"
	ConsentDataMeaningRelated    ConsentDataMeaningEnum = "related"
	ConsentDataMeaningDependents ConsentDataMeaningEnum = "dependents"
	ConsentDataMeaningAuthoredBy ConsentDataMeaningEnum = "authoredby"
)

// IsValid checks if the consent data meaning is valid.
func (c ConsentDataMeaningEnum) IsValid() bool {
	switch c {
	case ConsentDataMeaningInstance, ConsentDataMeaningRelated,
		ConsentDataMeaningDependents, ConsentDataMeaningAuthoredBy:
		return true
	}

	return false
}

// String converts the consent data meaning to string.
func (c ConsentDataMeaningEnum) String() string {
	return string(c)
}
