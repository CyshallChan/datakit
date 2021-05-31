package webaudio

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/chromedp/cdproto/cdp"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// GraphObjectID an unique ID for a graph object (AudioContext, AudioNode,
// AudioParam) in Web Audio API.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-GraphObjectId
type GraphObjectID string

// String returns the GraphObjectID as string value.
func (t GraphObjectID) String() string {
	return string(t)
}

// ContextType enum of BaseAudioContext types.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-ContextType
type ContextType string

// String returns the ContextType as string value.
func (t ContextType) String() string {
	return string(t)
}

// ContextType values.
const (
	ContextTypeRealtime ContextType = "realtime"
	ContextTypeOffline  ContextType = "offline"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ContextType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ContextType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ContextType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ContextType(in.String()) {
	case ContextTypeRealtime:
		*t = ContextTypeRealtime
	case ContextTypeOffline:
		*t = ContextTypeOffline

	default:
		in.AddError(errors.New("unknown ContextType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ContextType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ContextState enum of AudioContextState from the spec.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-ContextState
type ContextState string

// String returns the ContextState as string value.
func (t ContextState) String() string {
	return string(t)
}

// ContextState values.
const (
	ContextStateSuspended ContextState = "suspended"
	ContextStateRunning   ContextState = "running"
	ContextStateClosed    ContextState = "closed"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ContextState) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ContextState) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ContextState) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ContextState(in.String()) {
	case ContextStateSuspended:
		*t = ContextStateSuspended
	case ContextStateRunning:
		*t = ContextStateRunning
	case ContextStateClosed:
		*t = ContextStateClosed

	default:
		in.AddError(errors.New("unknown ContextState value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ContextState) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// NodeType enum of AudioNode types.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-NodeType
type NodeType string

// String returns the NodeType as string value.
func (t NodeType) String() string {
	return string(t)
}

// ChannelCountMode enum of AudioNode::ChannelCountMode from the spec.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-ChannelCountMode
type ChannelCountMode string

// String returns the ChannelCountMode as string value.
func (t ChannelCountMode) String() string {
	return string(t)
}

// ChannelCountMode values.
const (
	ChannelCountModeClampedMax ChannelCountMode = "clamped-max"
	ChannelCountModeExplicit   ChannelCountMode = "explicit"
	ChannelCountModeMax        ChannelCountMode = "max"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ChannelCountMode) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ChannelCountMode) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ChannelCountMode) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ChannelCountMode(in.String()) {
	case ChannelCountModeClampedMax:
		*t = ChannelCountModeClampedMax
	case ChannelCountModeExplicit:
		*t = ChannelCountModeExplicit
	case ChannelCountModeMax:
		*t = ChannelCountModeMax

	default:
		in.AddError(errors.New("unknown ChannelCountMode value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ChannelCountMode) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ChannelInterpretation enum of AudioNode::ChannelInterpretation from the
// spec.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-ChannelInterpretation
type ChannelInterpretation string

// String returns the ChannelInterpretation as string value.
func (t ChannelInterpretation) String() string {
	return string(t)
}

// ChannelInterpretation values.
const (
	ChannelInterpretationDiscrete ChannelInterpretation = "discrete"
	ChannelInterpretationSpeakers ChannelInterpretation = "speakers"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ChannelInterpretation) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ChannelInterpretation) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ChannelInterpretation) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ChannelInterpretation(in.String()) {
	case ChannelInterpretationDiscrete:
		*t = ChannelInterpretationDiscrete
	case ChannelInterpretationSpeakers:
		*t = ChannelInterpretationSpeakers

	default:
		in.AddError(errors.New("unknown ChannelInterpretation value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ChannelInterpretation) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ParamType enum of AudioParam types.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-ParamType
type ParamType string

// String returns the ParamType as string value.
func (t ParamType) String() string {
	return string(t)
}

// AutomationRate enum of AudioParam::AutomationRate from the spec.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-AutomationRate
type AutomationRate string

// String returns the AutomationRate as string value.
func (t AutomationRate) String() string {
	return string(t)
}

// AutomationRate values.
const (
	AutomationRateARate AutomationRate = "a-rate"
	AutomationRateKRate AutomationRate = "k-rate"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t AutomationRate) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t AutomationRate) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *AutomationRate) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch AutomationRate(in.String()) {
	case AutomationRateARate:
		*t = AutomationRateARate
	case AutomationRateKRate:
		*t = AutomationRateKRate

	default:
		in.AddError(errors.New("unknown AutomationRate value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *AutomationRate) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ContextRealtimeData fields in AudioContext that change in real-time.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-ContextRealtimeData
type ContextRealtimeData struct {
	CurrentTime              float64 `json:"currentTime"`              // The current context time in second in BaseAudioContext.
	RenderCapacity           float64 `json:"renderCapacity"`           // The time spent on rendering graph divided by render quantum duration, and multiplied by 100. 100 means the audio renderer reached the full capacity and glitch may occur.
	CallbackIntervalMean     float64 `json:"callbackIntervalMean"`     // A running mean of callback interval.
	CallbackIntervalVariance float64 `json:"callbackIntervalVariance"` // A running variance of callback interval.
}

// BaseAudioContext protocol object for BaseAudioContext.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-BaseAudioContext
type BaseAudioContext struct {
	ContextID             GraphObjectID        `json:"contextId"`
	ContextType           ContextType          `json:"contextType"`
	ContextState          ContextState         `json:"contextState"`
	RealtimeData          *ContextRealtimeData `json:"realtimeData,omitempty"`
	CallbackBufferSize    float64              `json:"callbackBufferSize"`    // Platform-dependent callback buffer size.
	MaxOutputChannelCount float64              `json:"maxOutputChannelCount"` // Number of output channels supported by audio hardware in use.
	SampleRate            float64              `json:"sampleRate"`            // Context sample rate.
}

// AudioListener protocol object for AudioListener.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-AudioListener
type AudioListener struct {
	ListenerID GraphObjectID `json:"listenerId"`
	ContextID  GraphObjectID `json:"contextId"`
}

// AudioNode protocol object for AudioNode.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-AudioNode
type AudioNode struct {
	NodeID                GraphObjectID         `json:"nodeId"`
	ContextID             GraphObjectID         `json:"contextId"`
	NodeType              cdp.NodeType          `json:"nodeType"`
	NumberOfInputs        float64               `json:"numberOfInputs"`
	NumberOfOutputs       float64               `json:"numberOfOutputs"`
	ChannelCount          float64               `json:"channelCount"`
	ChannelCountMode      ChannelCountMode      `json:"channelCountMode"`
	ChannelInterpretation ChannelInterpretation `json:"channelInterpretation"`
}

// AudioParam protocol object for AudioParam.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-AudioParam
type AudioParam struct {
	ParamID      GraphObjectID  `json:"paramId"`
	NodeID       GraphObjectID  `json:"nodeId"`
	ContextID    GraphObjectID  `json:"contextId"`
	ParamType    ParamType      `json:"paramType"`
	Rate         AutomationRate `json:"rate"`
	DefaultValue float64        `json:"defaultValue"`
	MinValue     float64        `json:"minValue"`
	MaxValue     float64        `json:"maxValue"`
}
