package lengthconv

func FToM(f Feet) Meters { return Meters(f * 0.3048) }

func MToF(m Meters) Feet { return Feet(m / 0.3048) }
