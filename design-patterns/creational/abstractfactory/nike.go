package abstractfactory

type nike struct {
}

func (s *nike) makeShoe() iShoe {
	// shoe :=
	return &nike{
		shoe: shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *nike) makeShort() iShort {
	return &nikeShort{
		short: short{
			logo: "nike",
			size: 14,
		},
	}
}
