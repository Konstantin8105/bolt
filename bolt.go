package bolt

import (
	"fmt"
	"math"
)

// Bolt - base property of bolt
type Bolt struct {
	bc Class
	bd BoltDiameter
}

// New - create a new bolt
func New(bc Class, bd BoltDiameter) Bolt {
	return Bolt{bc: bc, bd: bd}
}

// Fyb - return Fyb stress.
// unit: Pa
func (b Bolt) Fyb() Fyb {
	return Fyb{BoltClass: b.bc}
}

// Fub - return Fub stress.
// unit: Pa
func (b Bolt) Fub() Fub {
	return Fub{BoltClass: b.bc}
}

// D - diameter of bolt.
// unit: meter
func (b Bolt) D() BoltDiameter {
	return b.bd
}

// Do - diameter of bolt hole.
// unit: meter
func (b Bolt) Do() HoleDiameter {
	return HoleDiameter{Dia: b.bd}
}

// Cl - class of bolt
func (b Bolt) Cl() Class {
	return b.bc
}

// As - area of As
func (b Bolt) As() AreaAs {
	return AreaAs{Dia: b.bd}
}

// Class is class of bolt
type Class string

// Typical names of bolt classes
const (
	G4p6  Class = "4.6"
	G4p8  Class = "4.8"
	G5p6  Class = "5.6"
	G5p8  Class = "5.8"
	G6p8  Class = "6.8"
	G8p8  Class = "8.8"
	G10p9 Class = "10.9"
)

// GetBoltClassList - list of all allowable bolt classes
func GetBoltClassList() []Class {
	return []Class{G4p6, G4p8, G5p6, G5p8, G6p8, G8p8, G10p9}
}

func (bc Class) String() string {
	return fmt.Sprintf("Cl%s", string(bc))
}

// BoltDiameter is diameter of bolt
// unit: meter
type BoltDiameter float64

// Typical bolt diameters
const (
	D12 BoltDiameter = 12.e-3
	D16 BoltDiameter = 16.e-3
	D20 BoltDiameter = 20.e-3
	D24 BoltDiameter = 24.e-3
	D30 BoltDiameter = 30.e-3
	D36 BoltDiameter = 36.e-3
	D42 BoltDiameter = 42.e-3
	D48 BoltDiameter = 48.e-3
)

// GetBoltDiameterList - list of all allowable bolt classes
func GetBoltDiameterList() []BoltDiameter {
	return []BoltDiameter{D12, D16, D20, D24, D30, D36, D42, D48}
}

func (bd BoltDiameter) String() string {
	return fmt.Sprintf("HM%.0f", float64(bd)*1e3)
}

type HoleDiameter struct {
	Dia BoltDiameter
}

var holeDiameter = map[BoltDiameter]Diameter{
	D12: 13e-3,
	D16: 18e-3,
	D20: 22e-3,
	D24: 26e-3,
	D30: 33e-3,
	D36: 39e-3,
	D42: 45e-3,
	D48: 51e-3,
}

func (hd HoleDiameter) Value() Diameter {
	return holeDiameter[hd.Dia]
}
func (hd HoleDiameter) String() string {
	return fmt.Sprintf("For bolt %s hole is %s", hd.Dia, hd.Value())
}

// Diameter - dimension of diameter
type Diameter float64

func (dia Diameter) String() string {
	return fmt.Sprintf("Ø%s", Dimension(float64(dia)))
}

// Table of Fyb.
// unit: Pa
var fyb = map[Class]Stress{
	G4p6:  240.e6,
	G4p8:  320.e6,
	G5p6:  300.e6,
	G5p8:  400.e6,
	G6p8:  480.e6,
	G8p8:  640.e6,
	G10p9: 900.e6,
}

// Table of Fub.
// unit: Pa
var fub = map[Class]Stress{
	G4p6:  400.e6,
	G4p8:  400.e6,
	G5p6:  500.e6,
	G5p8:  500.e6,
	G6p8:  600.e6,
	G8p8:  800.e6,
	G10p9: 1000.e6,
}

// Fyb - stress of bolt in according to table 3.1. EN1993-1-8.
// unit: Pa
type Fyb struct {
	Stress
	BoltClass Class
}

// Value - return value of Fyb
func (f Fyb) Value() Stress {
	return fyb[f.BoltClass]
}

func (f Fyb) String() string {
	return fmt.Sprintf("In according to table 3.1 EN1993-1-8 value Fyb is %s", f.Value())
}

// Fub - stress of bolt in according to table 3.1. EN1993-1-8.
// unit: Pa
type Fub struct {
	Stress
	BoltClass Class
}

// Value - return value of Fub
func (f Fub) Value() Stress {
	return fub[f.BoltClass]
}

func (f Fub) String() string {
	return fmt.Sprintf("In according to table 3.1 EN1993-1-8 value Fub is %s", f.Value())
}

// Stress - struct of float64 for Stress values.
// unit: Pa
type Stress float64

func (s Stress) String() string {
	return fmt.Sprintf("%.1f MPa", float64(s)*1.e-6)
}

type BoltPinch struct {
	Dia BoltDiameter
}

var boltPinch = map[BoltDiameter]Dimension{
	D12: 1.75e-3,
	D16: 2.00e-3,
	D20: 2.50e-3,
	D24: 3.00e-3,
	D30: 3.50e-3,
	D36: 4.00e-3,
	D42: 4.50e-3,
	D48: 5.00e-3,
}

// Value - return value of bolt pinch
func (bp BoltPinch) Value() Dimension {
	return boltPinch[bp.Dia]
}

// Dimension - type for linear dimension sizes (height, thk, width)
type Dimension float64

func (d Dimension) String() string {
	return fmt.Sprintf("%.1f mm", float64(d)*1.e3)
}

// Area - type of area.
// unit - sq.meter
type Area float64

func (a Area) String() string {
	return fmt.Sprintf("%.1f mm\u00B2", float64(a)*1.e6)
}

// AreaAs tension stress area of the bolt
type AreaAs struct {
	Dia BoltDiameter
}

// Value - return value of Area As
func (as AreaAs) Value() Area {
	bp := BoltPinch{Dia: as.Dia}
	p := float64(bp.Value())
	dia := float64(as.Dia)
	return Area(math.Pi / 4. * math.Pow(dia-0.935229*p, 2.0))
}

func (as AreaAs) String() string {
	return fmt.Sprintf("Tension stress area of the bolt %s is %s", as.Dia, as.Value())
}

/*

	private double BOLT_AREA_As(int Dia)
	{
	    return (new General()).CONST_M_PI/4.* pow(Dia-0.935229*BOLT_PITCH(Dia),2.);
	}

	private double BOLT_AREA_A(int Dia)
	{
	    return (new General()).CONST_M_PI/4.* pow(Dia,2.);
	}

	private double EN1993_1_8_TABLE_3_4_FtRd(double Pub, double As, double gamma_M2)
	{
	    double k2 = 0.9;
	    return k2 * Pub * As / gamma_M2;
	}

	private double EN1993_1_8_TABLE_3_4_FvRd(double Pub, double As, double gamma_M2, BOLT_CLASS _BS)
	{
	    double alphaV = 0.0;
	    switch(_BS)
	    {
	        case g4_6 : alphaV = 0.6; break;
	        case g4_8 : alphaV = 0.5; break;
	        case g5_6 : alphaV = 0.6; break;
	        case g5_8 : alphaV = 0.5; break;
	        case g6_8 : alphaV = 0.5; break;
	        case g8_8 : alphaV = 0.6; break;
	        case g10_9: alphaV = 0.5; break;
	        default:
	        	alphaV = 0.0;
	    }
	    return alphaV * Pub * As / gamma_M2;
	}

	private double EN1993_1_8_TABLE_3_3_e1_min(int DiameterBolt)
	{
	    return BOLT_Do(DiameterBolt)*1.2;
	}

	private double EN1993_1_8_TABLE_3_3_e2_min(int DiameterBolt)
	{
	    return BOLT_Do(DiameterBolt)*1.2;
	}

	private double EN1993_1_8_TABLE_3_3_e3_min(int DiameterBolt)
	{
	    return BOLT_Do(DiameterBolt)*1.5;
	}

	private double EN1993_1_8_TABLE_3_3_e4_min(int DiameterBolt)
	{
	    return BOLT_Do(DiameterBolt)*1.5;
	}

	private double EN1993_1_8_TABLE_3_3_p1_min(int DiameterBolt)
	{
	    return BOLT_Do(DiameterBolt)*2.2;
	}

	private double EN1993_1_8_TABLE_3_3_p2_min(int DiameterBolt)
	{
	    return BOLT_Do(DiameterBolt)*2.4;
	}


	private int DiameterBolt;
	private int DiameterHole;
	private BOLT_CLASS BS;
	private double gamma_M2;

	private double A;
	private double As;
	private double F_v_Rd;
	private double F_t_Rd;
	//private double B_p_Rd;

	Bolt(int _Dia, BOLT_CLASS _BS)
	{
		DiameterBolt   	= _Dia;
		DiameterHole 	= BOLT_Do(DiameterBolt);//diameter+0.003;

		BS = _BS;

		gamma_M2 = 1.25;

		//B_p_Rd = 1e30;
		A  = BOLT_AREA_A (DiameterBolt);
		As = BOLT_AREA_As(DiameterBolt);

		F_t_Rd = EN1993_1_8_TABLE_3_4_FtRd(EN1993_1_8_TABLE_3_1_Fub(BS), As, gamma_M2);
		F_v_Rd = EN1993_1_8_TABLE_3_4_FvRd(EN1993_1_8_TABLE_3_1_Fub(BS), As, gamma_M2,BS);
	}

	String Output()
	{
		String out = new String();
		out = "";
	    out += "Bolt: " + PrintfDia(DiameterBolt) + "\n";
	    out += "Class of bolt: " + PrintfBS(BS) + "\n";
	    out += "Diameter of hole: " + DiameterHole + " mm\n";
	    out += "Partial safety factor for joint:\n";
	    out += String.format("gammaM2 = %.2f\n",gamma_M2);
	    out += "\n";
	    out += "The gross cross-section area of bolt:\n";
	    out += String.format("A  = %.1f sq.mm\n",A );
	    out += "The tensile stress area of the bolt:\n";
	    out += String.format("As = %.1f sq.mm\n",As);
	    out += "\n";
	    out += "Nominal value of the yield strenght(table 3.1 EN1993-1-8):\n";
	    out += String.format("Fyb  = %.1f MPA\n",EN1993_1_8_TABLE_3_1_Fyb(BS) );
	    out += "Nominal value of the ultimate tensile strenght(table 3.1 EN1993-1-8):\n";
	    out += String.format("Fub  = %.1f MPA\n",EN1993_1_8_TABLE_3_1_Fub(BS) );
	    out += "\n";
	    out += "Shear resistance per shear plane(table 3.4 EN1993-1-8):\n";
	    out += String.format("Fv,Rd = %.1fkN\n",F_v_Rd*1e-3);
	    out += "Tension resistance(table 3.4 EN1993-1-8):\n";
	    out += String.format("Ft,Rd = %.1fkN\n",F_t_Rd*1e-3);
	    out += "\n";
	    out += "Minimal spacing, end and edge distances(table 3.3 EN1993-1-8):\n";
	    out += String.format("End  distance e1  = %.1f mm\n",EN1993_1_8_TABLE_3_3_e1_min(DiameterBolt) );
	    out += String.format("Edge distance e2  = %.1f mm\n",EN1993_1_8_TABLE_3_3_e2_min(DiameterBolt) );
	    out += String.format("Distance e3 in slotted holes = %.1f mm\n",EN1993_1_8_TABLE_3_3_e3_min(DiameterBolt) );
	    out += String.format("Distance e4 in slotted holes = %.1f mm\n",EN1993_1_8_TABLE_3_3_e4_min(DiameterBolt) );
	    out += String.format("Spacing p1 = %.1f mm\n",EN1993_1_8_TABLE_3_3_p1_min(DiameterBolt) );
	    out += String.format("Spacing p2 = %.1f mm\n",EN1993_1_8_TABLE_3_3_p2_min(DiameterBolt) );

	    // *
	    if(B_p_Rd < 1e20)
	    	out += "B_p_Rd = %.1f kN\n",B_p_Rd*1e-3);* //
	    return out;
	};
};
*/
