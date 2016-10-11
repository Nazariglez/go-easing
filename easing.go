/**
 * Created by nazarigonzalez on 11/10/16.
 */

package easing

import "math"

type Easing func(t float64) float64

const pid2 float64 = math.Pi / 2
const pi2 float64 = math.Pi * 2

func Interpolate(from, to, totalTime, elapsedTime float64, easing Easing) float64 {
	return from + ((to - from) * easing(elapsedTime/totalTime))
}

func Linear() Easing {
	return func(t float64) float64 {
		return t
	}
}

func InQuad() Easing {
	return func(t float64) float64 {
		return t * t
	}
}

func OutQuad() Easing {
	return func(t float64) float64 {
		return t * (2 - t)
	}
}

func InOutQuad() Easing {
	return func(t float64) float64 {
		t *= 2

		if t < 1 {
			return 0.5 * t * t
		}

		t--
		return -0.5 * (t*(t-2) - 1)
	}
}

func InCubic() Easing {
	return func(t float64) float64 {
		return t * t * t
	}
}

func OutCubic() Easing {
	return func(t float64) float64 {
		t--
		return t*t*t + 1
	}
}

func InOutCubic() Easing {
	return func(t float64) float64 {
		t *= 2

		if t < 1 {
			return 0.5 * t * t * t
		}

		t -= 2
		return 0.5 * (t*t*t + 2)
	}
}

func InQuart() Easing {
	return func(t float64) float64 {
		return t * t * t * t
	}
}

func OutQuart() Easing {
	return func(t float64) float64 {
		t--

		return 1 - t*t*t*t
	}
}

func InOutQuart() Easing {
	return func(t float64) float64 {
		t *= 2

		if t < 1 {
			return 0.5 * t * t * t * t
		}

		t -= 2
		return 0.5 * (t*t*t*t - 2)
	}
}

func InQuint() Easing {
	return func(t float64) float64 {
		return t * t * t * t * t
	}
}

func OutQuint() Easing {
	return func(t float64) float64 {
		t--
		return t*t*t*t*t + 1
	}
}

func InOutQuint() Easing {
	return func(t float64) float64 {
		t *= 2

		if t < 1 {
			return 0.5 * t * t * t * t * t
		}

		t -= 2
		return 0.5 * (t*t*t*t*t + 2)
	}
}

func InSine() Easing {
	return func(t float64) float64 {
		return 1 - math.Cos(t*pid2)
	}
}

func OutSine() Easing {
	return func(t float64) float64 {
		return math.Sin(t * pid2)
	}
}

func InOutSine() Easing {
	return func(t float64) float64 {
		return 0.5 * (1 - math.Cos(math.Pi*t))
	}
}

func InExpo() Easing {
	return func(t float64) float64 {
		if t == 0 {
			return 0
		}

		return math.Pow(1024, t-1)
	}
}

func OutExpo() Easing {
	return func(t float64) float64 {
		if t == 1 {
			return 1
		}

		return 1 - math.Pow(2, -10*t)
	}
}

func InOutExpo() Easing {
	return func(t float64) float64 {
		if t == 0 {
			return 0
		} else if t == 1 {
			return 1
		}

		t *= 2

		if t < 1 {
			return 0.5 * (math.Pow(1024, t-1))
		}

		return 0.5 * (-math.Pow(2, -10*(t-1)) + 2)
	}
}

func InCirc() Easing {
	return func(t float64) float64 {
		return 1 - math.Sqrt(1-t*t)
	}
}

func OutCirc() Easing {
	return func(t float64) float64 {
		t--
		return math.Sqrt(1 - (t * t))
	}
}

func InOutCirc() Easing {
	return func(t float64) float64 {
		t *= 2

		if t < 1 {
			return -0.5 * (math.Sqrt(1-t*t) - 1)
		}

		return 0.5 * (math.Sqrt(1-(t-2)*(t-2)) + 1)
	}
}

func InElastic() Easing {
	return calculateInElastic(0.1, 0.4)
}

func InElasticCustom(a, p float64) Easing {
	return calculateInElastic(a, p)
}

func calculateInElastic(a, p float64) Easing {
	return func(t float64) float64 {
		if t == 0 {
			return 0
		}

		if t == 1 {
			return 1
		}

		s := 0.0

		if a < 1 {
			a = 1
			s = p / 4
		} else {
			s = p * math.Asin(1/a) / pi2
		}

		return (a * math.Pow(2, 10*(t-1)) * math.Sin(((t-1)-s)*pi2/p))
	}
}

func OutElastic() Easing {
	return calculateOutElastic(0.1, 0.4)
}

func OutElasticCustom(a, p float64) Easing {
	return calculateOutElastic(a, p)
}

func calculateOutElastic(a, p float64) Easing {
	return func(t float64) float64 {
		if t == 0 {
			return 0
		}

		if t == 1 {
			return 1
		}

		s := 0.0

		if a < 1 {
			a = 1
			s = p / 4
		} else {
			s = p * math.Asin(1/a) / pi2
		}

		return (a*math.Pow(2, -10*t)*math.Sin((t-s)*pi2/p) + 1)
	}
}

func InOutElastic() Easing {
	return calculateInOutElastic(0.1, 0.4)
}

func InOutElasticCustom(a, p float64) Easing {
	return calculateInOutElastic(a, p)
}

func calculateInOutElastic(a, p float64) Easing {
	return func(t float64) float64 {
		if t == 0 {
			return 0
		}

		if t == 1 {
			return 1
		}

		s := 0.0

		if a < 1 {
			a = 1
			s = p / 4
		} else {
			s = p * math.Asin(1/a) / pi2
		}

		t *= 2

		if t < 1 {
			return -0.5 * (a * math.Pow(2, 10*(t-1)) * math.Sin(((t-1)-s)*pi2/p))
		}

		return a*math.Pow(2, -10*(t-1))*math.Sin(((t-1)-s)*pi2/p)*0.5 + 1
	}
}

func InBack() Easing {
	return calculateInBack(1.70158)
}

func InBackCustom(v float64) Easing {
	return calculateInBack(v)
}

func calculateInBack(v float64) Easing {
	return func(t float64) float64 {
		return t * t * ((v+1)*t - v)
	}
}

func OutBack() Easing {
	return calculateOutBack(1.70158)
}

func OutBackCustom(v float64) Easing {
	return calculateOutBack(v)
}

func calculateOutBack(v float64) Easing {
	return func(t float64) float64 {
		t--
		return t*t*((v+1)*t+v) + 1
	}
}

func InOutBack() Easing {
	return calculateInOutBack(1.70158)
}

func InOutBackCustom(v float64) Easing {
	return calculateInOutBack(v)
}

func calculateInOutBack(v float64) Easing {
	return func(t float64) float64 {
		v *= 1.525
		t *= 2

		if t < 1 {
			return 0.5 * (t * t * ((v+1)*t - v))
		}

		return t*t*((v+1)*t+v) + 1
	}
}

//todo fix bounce ease

func OutBounce() Easing {
	return func(t float64) float64 {
		if t < 1/2.75 {
			return 7.5625 * t * t
		} else if t < 2/2.75 {
			t -= 1.5 / 2.75
			return 7.5625*t*t + 0.75
		} else if t < 2.5/2.75 {
			t -= 2.25 / 2.75
			return 7.5625*t*t + 0.9375
		}

		t -= 2.625 / 2.75
		return 7.5625 * t * t * 0.984375
	}
}

//allocate here to avoid call each time
var internalOutBounce Easing = OutBounce()

func InBounce() Easing {
	return func(t float64) float64 {
		return 1 - internalOutBounce(1-t)
	}
}

var internalInBounce Easing = InBounce()

func InOutBounce() Easing {
	return func(t float64) float64 {
		if t < 0.5 {
			return internalInBounce(t*2) * 0.5
		}

		return internalOutBounce(t*2-1)*0.5 + 0.5
	}
}
