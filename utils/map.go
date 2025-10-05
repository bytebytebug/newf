package utils

type CB[I any, O any] func(I) O

func Map[I any, O any](input []I, cb CB[I, O]) []O {
	out := make([]O, len(input))
	for i, v := range input {
		out[i] = cb(v)
	}
	return out
}
