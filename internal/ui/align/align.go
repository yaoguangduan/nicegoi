package align

type Align string

const BoxCenter = Align("center")
const BoxStart = Align("start")
const BoxEnd = Align("end")
const BoxBaseline = Align("baseline")

type Justify string

const RowCenter = Justify("center")
const RowLeft = Justify("left")
const RowRight = Justify("right")
const RowSpaceBetween = Justify("space-between")
const RowSpaceAround = Justify("space-around")
