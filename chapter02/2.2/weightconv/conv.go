package weightconv

func PToK(p Pound) Kilogram { return Kilogram(p * 0.453592) }

func KToP(k Kilogram) Pound { return Pound(k / 0.453592) }
