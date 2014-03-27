package karatsuba

import (
	"math/big"
	"testing"
)

func TestSplit(t *testing.T) {
	for _, testData := range listTests {
		res := split(testData.input, 2)

		if res[0].Cmp(testData.expected[0]) != 0 {
			t.Fatalf("left is not equal %v == %v -> %v", res[0], testData.expected[0], res[0].Cmp(testData.expected[0]))
		}

		if res[1].Cmp(testData.expected[1]) != 0 {
			t.Fatalf("right is not equal %v == %v -> %v", res[1], testData.expected[1], res[1].Cmp(testData.expected[1]))
		}
	}
}

var listTests = []struct {
	input    *big.Int
	expected []*big.Int
}{
	{
		bigint(100),
		[]*big.Int{bigint(25), bigint(0)},
	},
	{
		bigint(101),
		[]*big.Int{bigint(25), bigint(1)},
	},
	{
		bigint(102),
		[]*big.Int{bigint(25), bigint(2)},
	},
	{
		bigint(103),
		[]*big.Int{bigint(25), bigint(3)},
	},
	{
		bigint(104),
		[]*big.Int{bigint(26), bigint(0)},
	},
	{
		bigint(105),
		[]*big.Int{bigint(26), bigint(1)},
	},
	{
		bigint(106),
		[]*big.Int{bigint(26), bigint(2)},
	},
	{
		bigint(107),
		[]*big.Int{bigint(26), bigint(3)},
	},
	{
		bigint(108),
		[]*big.Int{bigint(27), bigint(0)},
	},
	{
		bigint(109),
		[]*big.Int{bigint(27), bigint(1)},
	},
	{
		bigint(110),
		[]*big.Int{bigint(27), bigint(2)},
	},
	{
		bigint(111),
		[]*big.Int{bigint(27), bigint(3)},
	},
	{
		bigint(112),
		[]*big.Int{bigint(28), bigint(0)},
	},
	{
		bigint(113),
		[]*big.Int{bigint(28), bigint(1)},
	},
	{
		bigint(114),
		[]*big.Int{bigint(28), bigint(2)},
	},
	{
		bigint(115),
		[]*big.Int{bigint(28), bigint(3)},
	},
	{
		bigint(116),
		[]*big.Int{bigint(29), bigint(0)},
	},
	{
		bigint(117),
		[]*big.Int{bigint(29), bigint(1)},
	},
	{
		bigint(118),
		[]*big.Int{bigint(29), bigint(2)},
	},
	{
		bigint(119),
		[]*big.Int{bigint(29), bigint(3)},
	},
	{
		bigint(120),
		[]*big.Int{bigint(30), bigint(0)},
	},
	{
		bigint(121),
		[]*big.Int{bigint(30), bigint(1)},
	},
	{
		bigint(122),
		[]*big.Int{bigint(30), bigint(2)},
	},
	{
		bigint(123),
		[]*big.Int{bigint(30), bigint(3)},
	},
	{
		bigint(124),
		[]*big.Int{bigint(31), bigint(0)},
	},
	{
		bigint(125),
		[]*big.Int{bigint(31), bigint(1)},
	},
	{
		bigint(126),
		[]*big.Int{bigint(31), bigint(2)},
	},
	{
		bigint(127),
		[]*big.Int{bigint(31), bigint(3)},
	},
	{
		bigint(128),
		[]*big.Int{bigint(32), bigint(0)},
	},
	{
		bigint(129),
		[]*big.Int{bigint(32), bigint(1)},
	},
	{
		bigint(130),
		[]*big.Int{bigint(32), bigint(2)},
	},
	{
		bigint(131),
		[]*big.Int{bigint(32), bigint(3)},
	},
	{
		bigint(132),
		[]*big.Int{bigint(33), bigint(0)},
	},
	{
		bigint(133),
		[]*big.Int{bigint(33), bigint(1)},
	},
	{
		bigint(134),
		[]*big.Int{bigint(33), bigint(2)},
	},
	{
		bigint(135),
		[]*big.Int{bigint(33), bigint(3)},
	},
	{
		bigint(136),
		[]*big.Int{bigint(34), bigint(0)},
	},
	{
		bigint(137),
		[]*big.Int{bigint(34), bigint(1)},
	},
	{
		bigint(138),
		[]*big.Int{bigint(34), bigint(2)},
	},
	{
		bigint(139),
		[]*big.Int{bigint(34), bigint(3)},
	},
	{
		bigint(140),
		[]*big.Int{bigint(35), bigint(0)},
	},
	{
		bigint(141),
		[]*big.Int{bigint(35), bigint(1)},
	},
	{
		bigint(142),
		[]*big.Int{bigint(35), bigint(2)},
	},
	{
		bigint(143),
		[]*big.Int{bigint(35), bigint(3)},
	},
	{
		bigint(144),
		[]*big.Int{bigint(36), bigint(0)},
	},
	{
		bigint(145),
		[]*big.Int{bigint(36), bigint(1)},
	},
	{
		bigint(146),
		[]*big.Int{bigint(36), bigint(2)},
	},
	{
		bigint(147),
		[]*big.Int{bigint(36), bigint(3)},
	},
	{
		bigint(148),
		[]*big.Int{bigint(37), bigint(0)},
	},
	{
		bigint(149),
		[]*big.Int{bigint(37), bigint(1)},
	},
	{
		bigint(150),
		[]*big.Int{bigint(37), bigint(2)},
	},
	{
		bigint(151),
		[]*big.Int{bigint(37), bigint(3)},
	},
	{
		bigint(152),
		[]*big.Int{bigint(38), bigint(0)},
	},
	{
		bigint(153),
		[]*big.Int{bigint(38), bigint(1)},
	},
	{
		bigint(154),
		[]*big.Int{bigint(38), bigint(2)},
	},
	{
		bigint(155),
		[]*big.Int{bigint(38), bigint(3)},
	},
	{
		bigint(156),
		[]*big.Int{bigint(39), bigint(0)},
	},
	{
		bigint(157),
		[]*big.Int{bigint(39), bigint(1)},
	},
	{
		bigint(158),
		[]*big.Int{bigint(39), bigint(2)},
	},
	{
		bigint(159),
		[]*big.Int{bigint(39), bigint(3)},
	},
	{
		bigint(160),
		[]*big.Int{bigint(40), bigint(0)},
	},
	{
		bigint(161),
		[]*big.Int{bigint(40), bigint(1)},
	},
	{
		bigint(162),
		[]*big.Int{bigint(40), bigint(2)},
	},
	{
		bigint(163),
		[]*big.Int{bigint(40), bigint(3)},
	},
	{
		bigint(164),
		[]*big.Int{bigint(41), bigint(0)},
	},
	{
		bigint(165),
		[]*big.Int{bigint(41), bigint(1)},
	},
	{
		bigint(166),
		[]*big.Int{bigint(41), bigint(2)},
	},
	{
		bigint(167),
		[]*big.Int{bigint(41), bigint(3)},
	},
	{
		bigint(168),
		[]*big.Int{bigint(42), bigint(0)},
	},
	{
		bigint(169),
		[]*big.Int{bigint(42), bigint(1)},
	},
	{
		bigint(170),
		[]*big.Int{bigint(42), bigint(2)},
	},
	{
		bigint(171),
		[]*big.Int{bigint(42), bigint(3)},
	},
	{
		bigint(172),
		[]*big.Int{bigint(43), bigint(0)},
	},
	{
		bigint(173),
		[]*big.Int{bigint(43), bigint(1)},
	},
	{
		bigint(174),
		[]*big.Int{bigint(43), bigint(2)},
	},
	{
		bigint(175),
		[]*big.Int{bigint(43), bigint(3)},
	},
	{
		bigint(176),
		[]*big.Int{bigint(44), bigint(0)},
	},
	{
		bigint(177),
		[]*big.Int{bigint(44), bigint(1)},
	},
	{
		bigint(178),
		[]*big.Int{bigint(44), bigint(2)},
	},
	{
		bigint(179),
		[]*big.Int{bigint(44), bigint(3)},
	},
	{
		bigint(180),
		[]*big.Int{bigint(45), bigint(0)},
	},
	{
		bigint(181),
		[]*big.Int{bigint(45), bigint(1)},
	},
	{
		bigint(182),
		[]*big.Int{bigint(45), bigint(2)},
	},
	{
		bigint(183),
		[]*big.Int{bigint(45), bigint(3)},
	},
	{
		bigint(184),
		[]*big.Int{bigint(46), bigint(0)},
	},
	{
		bigint(185),
		[]*big.Int{bigint(46), bigint(1)},
	},
	{
		bigint(186),
		[]*big.Int{bigint(46), bigint(2)},
	},
	{
		bigint(187),
		[]*big.Int{bigint(46), bigint(3)},
	},
	{
		bigint(188),
		[]*big.Int{bigint(47), bigint(0)},
	},
	{
		bigint(189),
		[]*big.Int{bigint(47), bigint(1)},
	},
	{
		bigint(190),
		[]*big.Int{bigint(47), bigint(2)},
	},
	{
		bigint(191),
		[]*big.Int{bigint(47), bigint(3)},
	},
	{
		bigint(192),
		[]*big.Int{bigint(48), bigint(0)},
	},
	{
		bigint(193),
		[]*big.Int{bigint(48), bigint(1)},
	},
	{
		bigint(194),
		[]*big.Int{bigint(48), bigint(2)},
	},
	{
		bigint(195),
		[]*big.Int{bigint(48), bigint(3)},
	},
	{
		bigint(196),
		[]*big.Int{bigint(49), bigint(0)},
	},
	{
		bigint(197),
		[]*big.Int{bigint(49), bigint(1)},
	},
	{
		bigint(198),
		[]*big.Int{bigint(49), bigint(2)},
	},
	{
		bigint(199),
		[]*big.Int{bigint(49), bigint(3)},
	},
}
