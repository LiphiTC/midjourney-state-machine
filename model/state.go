package model

const (
	// "Imagine/Blend/Describe" commands
	ActionImagine  string = "imagine"
	ActionBlend    string = "blend"
	ActionDescribe string = "describe"

	ActionReroll string = "reroll"

	// "U1/U2/U3/U4" buttons, to separate your selected image
	ActionUpscale1 string = "upscale1"
	ActionUpscale2 string = "upscale2"
	ActionUpscale3 string = "upscale3"
	ActionUpscale4 string = "upscale4"

	// "Upscale(2x/4x)" buttons, to increase the size of your image
	ActionUpscale2x string = "upscale_2x"
	ActionUpscale4x string = "upscale_4x"

	// "V1/V2/V3/V4" buttons
	ActionVariation1 string = "variation1"
	ActionVariation2 string = "variation2"
	ActionVariation3 string = "variation3"
	ActionVariation4 string = "variation4"

	// "Vary(Strong)/Vary(Subtle)" buttons
	ActionVaryStrong string = "high_variation"
	ActionVarySubtle string = "low_variation"

	// "Vary(Region)" button
	ActionInpaint string = "inpaint"

	// "Make Square/Zoom Out 1.5x/Zoom Out 2x/Custom Zoom" buttons
	// does not check if the "Make Square" button exists
	ActionOutpaintMakeSquare string = "outpaint_1x"
	ActionOutpaint75         string = "outpaint_1.5x"
	ActionOutpaint50         string = "outpaint_2x"
	ActionOutpaintCustom     string = "outpaint_custom"

	ActionPanUp    string = "pan_up"
	ActionPanDown  string = "pan_down"
	ActionPanLeft  string = "pan_left"
	ActionPanRight string = "pan_right"

	// "1/2/3/4/imagine all" buttons, use "describe" to distinguish from imagine command
	// describe not supported now
	ActionDescribe1   string = "describe1"
	ActionDescribe2   string = "describe2"
	ActionDescribe3   string = "describe3"
	ActionDescribe4   string = "describe4"
	ActionDescribeAll string = "describe_all"

	// any pan actions in history
	PanStateVertical   string = "vertical"
	PanStateHorizontal string = "horizontal"
)

var ActionsUpscaleSeparate = []string{
	ActionUpscale1,
	ActionUpscale2,
	ActionUpscale3,
	ActionUpscale4,
}

var ActionsUpscaleIncrease = []string{
	ActionUpscale2x,
	ActionUpscale4x,
}

var ActionsVariation = []string{
	ActionVariation1,
	ActionVariation2,
	ActionVariation3,
	ActionVariation4,
}

var ActionsVary = []string{
	ActionVaryStrong,
	ActionVarySubtle,
}

var ActionsOutpaint = []string{
	ActionOutpaint75,
	ActionOutpaint50,
	ActionOutpaintCustom,
	ActionOutpaintMakeSquare,
}

var ActionsPan = []string{
	ActionPanUp,
	ActionPanDown,
	ActionPanLeft,
	ActionPanRight,
}

var ActionsDescribe = []string{
	ActionDescribe1,
	ActionDescribe2,
	ActionDescribe3,
	ActionDescribe4,
	// ActionDescribeAll,
}

// actions lists in 6 basic states, based on state graph in README.md
var State1Actions []string
var State2Actions []string
var State3Actions []string
var State4Actions []string
var State5Actions []string
var State6Actions []string
var State7Actions []string

func init() {
	State1Actions = append(State1Actions, ActionReroll)
	State1Actions = append(State1Actions, ActionsUpscaleSeparate...)
	State1Actions = append(State1Actions, ActionsVariation...)

	State2Actions = append(State2Actions, ActionReroll)
	State2Actions = append(State2Actions, ActionsDescribe...)

	State3Actions = append(State3Actions, ActionsVary...)
	State3Actions = append(State3Actions, ActionInpaint)
	State3Actions = append(State3Actions, ActionsUpscaleIncrease...)
	State3Actions = append(State3Actions, ActionsOutpaint...)
	State3Actions = append(State3Actions, ActionsPan...)

	State4Actions = append(State4Actions, ActionReroll)
	State4Actions = append(State4Actions, ActionsUpscaleSeparate...)

	State5Actions = append(State5Actions, ActionsOutpaint...)
	State5Actions = append(State5Actions, ActionPanUp)
	State5Actions = append(State5Actions, ActionPanDown)

	State6Actions = append(State6Actions, ActionsOutpaint...)
	State6Actions = append(State6Actions, ActionPanLeft)
	State6Actions = append(State6Actions, ActionPanRight)

	State7Actions = append(State7Actions, ActionsUpscaleIncrease...)
}
