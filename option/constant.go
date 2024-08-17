package option

type Align string

const BoxCenter = Align("center")
const BoxStart = Align("start")
const BoxEnd = Align("end")
const BoxBaseline = Align("baseline")

type Justify string

const RowCenter = Justify("center")
const RowStart = Justify("start")
const RowEnd = Justify("end")
const RowSpaceBetween = Justify("space-between")
const RowSpaceAround = Justify("space-around")

type Placement string

const Top = Placement("top")
const Left = Placement("left")
const Right = Placement("right")
const Bottom = Placement("bottom")

type Theme string

const Primary = Theme("primary")
const Success = Theme("success")
const Default = Theme("default")
const Danger = Theme("danger")
const Warning = Theme("warning")

type Shape string

const Rectangle = Shape("rectangle")
const Square = Shape("square")
const Circle = Shape("circle")
const Round = Shape("round")
const Mark = Shape("mark") // just for tag widget

type Size string

const Large = Size("large")
const Small = Size("small")
const Medium = Size("medium")

type Variant string

const Base = Variant("base")
const Outline = Variant("outline")
const Dashed = Variant("dashed")
const Text = Variant("text")

type TagVariant string

const TagVarDark = TagVariant("dark")
const TagVarLight = TagVariant("light")
const TagVarOutline = TagVariant("outline")
const TagVarLightOutline = TagVariant("light-outline")

type DateTimeMode string

const Date = DateTimeMode("date")
const Year = DateTimeMode("year")
const Month = DateTimeMode("month")

type ProgressState string

const (
	ProgressActive  ProgressState = "active"
	ProgressError   ProgressState = "error"
	ProgressWarning ProgressState = "warning"
	ProgressSuccess ProgressState = "success"
)
