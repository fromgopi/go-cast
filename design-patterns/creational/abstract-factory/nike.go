package main

type nike struct {
}

func (s *nike) makeShoe() iShoe {
	return &nike{
		shoe: shoe{
			logo: "nike",
			size: 14,
		},
	}
}
