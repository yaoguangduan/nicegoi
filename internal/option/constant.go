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
